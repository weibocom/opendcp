/**
 *    Copyright (C) 2016 Weibo Inc.
 *
 *    This file is part of Opendcp.
 *
 *    Opendcp is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation; version 2 of the License.
 *
 *    Opendcp is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with Opendcp; if not, write to the Free Software
 *    Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301  USA
 */

package executor

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/astaxie/beego"

	"weibo.com/opendcp/orion/handler"
	"weibo.com/opendcp/orion/models"
	"weibo.com/opendcp/orion/service"
)

var (
	// Executor is the singleton FlowExecutor instance to run all tasks.
	Executor    = &FlowExecutor{}
	flowService = service.Flow
	logService  = service.Logs

	// Default retry option
	dftRetryOpt = &models.RetryOption{
		RetryTimes:  0,
		IgnoreError: false,
	}

	// lock for running jobs
	lock = &sync.Mutex{}

	// workers for each pool
	workers = make(map[int]*Worker)

	//chan to put nodeStates
	workNodeQueue *QueueNode
)

const (
	retryInterval = 5
)

// ExecOption contains options to execute a task.
type ExecOption struct {
	// MaxNum is the max num of nodes in one batch.
	MaxNum     int
	MaxFailNum int
}

// FlowExecutor dispatches tasks into queues.
type FlowExecutor struct {
}

//init queue
func Initial() {
	workNodeQueue = NewQueueNode()
	workNodeQueue.Start()
}

// Run creates a flow instance from the given template id, and nodes, and
// starts the flow instance.
func (exec *FlowExecutor) Run(flowImpl *models.FlowImpl, name string,
	option *ExecOption, nodes []*models.NodeState, context map[string]interface{}) error {

	fis, err := exec.Create(flowImpl, name, option, nodes, context)
	if err != nil {
		return err
	}

	for _, fi := range fis {
		ins := fi
		err = exec.Start(ins)
		if err != nil {
			return err
		}
	}

	return nil
}

// Create creates a flow instance with given template id, and nodes.
func (exec *FlowExecutor) Create(flowImpl *models.FlowImpl, name string, option *ExecOption,
	nodes []*models.NodeState, context map[string]interface{}) ([]*models.Flow, error) {

	beego.Info("Create new task from template[", flowImpl.Id, "]")

	nodesarray := make(map[int][]*models.NodeState)
	var instances []*models.Flow

	if option == nil {
		beego.Error("Option is nil")
		return nil, errors.New("option is nil")
	}

	if len(nodes) == 0 {
		beego.Error("nodes is empty")
		return nil, errors.New("nodes is empty")
	}

	// make sure all nodes are in the same pool
	for _, node := range nodes {
		pid := node.Pool.Id
		nodesarray[pid] = append(nodesarray[pid], node)
	}

	var stepOps []*models.StepOption
	err := json.Unmarshal([]byte(flowImpl.Steps), &stepOps)
	if err != nil {
		beego.Error("Bad step options: ", flowImpl.Steps, "[err]: ", err)
		return nil, errors.New("Bad step options: " + flowImpl.Steps + ", err:" + err.Error())
	}

	//merge override params with params
	overrideParams, ok := context["overrideParams"].(map[string]interface{})
	if !ok {
		beego.Error("bad overrideParams:", context["overrideParams"])
		return nil, errors.New("bad overrideParams!")
	}
	exec.MergeParams(stepOps, overrideParams)

	mergedBytes, _ := json.Marshal(stepOps)
	merged := string(mergedBytes)
	beego.Debug("Merged step options: " + merged)

	// create items for all nodes, set state = INIT
	opUser, ok := context["opUser"].(string)
	beego.Info("opUser Create new task ", opUser)
	if !ok {
		beego.Error("bad opUser:", context["opUser"])
		return nil, errors.New("bad opUser!")
	}

	for pid, node := range nodesarray {
		poolID := pid
		poolNode := node
		instance, err := exec.CreateFlowInstance(name, flowImpl, &models.Pool{Id: poolID}, poolNode[0].NodeType, stepOps, option, opUser)
		if err != nil {
			beego.Error("Fail to create flow instance", err)
			return nil, err
		}
		_, err = exec.CreateNodeStates(instance, poolNode)
		if err != nil {
			beego.Error("Fail to create node states for flow: ", instance.Name, err)
			flowService.DeleteBase(instance)
			return nil, err
		}
		instances = append(instances, instance)
	}

	return instances, nil
}

// Start starts an existing flow instance.
func (exec *FlowExecutor) Start(flow *models.Flow) error {
	lock.Lock()
	defer lock.Unlock()

	orign_flow_status := flow.Status

	if exec.isFlowInState(flow, models.STATUS_RUNNING) {
		logInfo := "Flow " + flow.Name + " is in state running: " + strconv.Itoa(flow.Status) + " do not start"
		logService.Info(flow.Id, logInfo)

		return errors.New(logInfo)
	}

	exec.SetFlowStatus(flow, models.STATUS_RUNNING)

	// run the flow
	job := func() error {
		logService.Info(flow.Id, "Run flow...")

		err := exec.RunFlow(flow, orign_flow_status)
		if err != nil {
			logService.Error(flow.Id, "Run flow error：", err)
			return err
		}

		return nil
	}

	// queue the job
	return exec.submit(flow.Pool.Id, job)
}

// submit submits a job(to run flow) into queue of a pool.
func (exec *FlowExecutor) submit(poolID int, job Job) error {

	worker := workers[poolID]
	if worker == nil {
		worker = NewWorker(poolID)
		workers[poolID] = worker
		worker.Start()
	}

	return worker.Submit(job)
}

// Pause paused a running flow by setting its status to STOPPED.
func (exec *FlowExecutor) Pause(flow *models.Flow) error {
	lock.Lock()
	defer lock.Unlock()

	beego.Info("Pause flow", flow.Name, flow.Id, "...")
	if !exec.isRunning(flow) {
		beego.Error("Pausing flow failed: flow", flow.Name, "is not running")
		return errors.New("Flow " + flow.Name + " is not running")
	}

	return exec.SetFlowStatus(flow, models.STATUS_STOPPED)
}
// Stop stopped a running flow by setting its status to SUCCESS.
func (exec *FlowExecutor) Stop(flow *models.Flow) error {
	lock.Lock()
	defer lock.Unlock()

	beego.Info("Stop flow", flow.Name, flow.Id, "...")

	err := exec.loadFlowStatus(flow)
	if err != nil {

		beego.Error("Stop failed: flow", flow.Name, ":", err.Error())
		return errors.New("Stop failed: flow" + flow.Name + ":" + err.Error())
	}


	if flow.Status == models.STATUS_SUCCESS || flow.Status == models.STATUS_FAILED {
		beego.Error("Stop failed: flow", flow.Name, "already finished")
		return errors.New("Flow " + flow.Name + " already finished")
	}
	return exec.SetFlowStatus(flow, models.STATUS_SUCCESS)
}

func (exec *FlowExecutor) SetFlowStatus(flow *models.Flow, status int) error {

	flow.Status = status
	beego.Debug("Set flow", flow.Name, "status =", flow.Status)
	flow.UpdatedTime = time.Now()

	return flowService.UpdateBase(flow)
}

func (exec *FlowExecutor) SetFlowStatusWithSpenTime(flow *models.Flow, spenTime float64, status int) error {

	flow.Status = status
	flow.RunTime = spenTime
	beego.Debug("Set flow", flow.Name, "status =", flow.Status)
	flow.UpdatedTime = time.Now()

	return flowService.UpdateBase(flow)
}

// load latest flow status from DB
func (exec *FlowExecutor) loadFlowStatus(flow *models.Flow) error {

	f := &models.Flow{Id: flow.Id}
	err := flowService.GetBase(f)
	if err != nil {
		return err
	}

	flow.Status = f.Status
	beego.Debug("Load latest status for flow:",
		flow.Id, flow.Name, ", status =", flow.Status)
	return nil
}

func (exec *FlowExecutor) isFlowInState(flow *models.Flow, state int) bool {

	err := exec.loadFlowStatus(flow)
	if err != nil {
		beego.Error("Update flow status fails: " + err.Error())
		return false
	}

	return flow.Status == state
}

func (exec *FlowExecutor) isRunning(flow *models.Flow) bool {
	return exec.isFlowInState(flow, models.STATUS_RUNNING)
}

// run the task by batches.
func (exec *FlowExecutor) RunFlow(flow *models.Flow, orign_flow_status int) error{

	logService.Info(flow.Id, fmt.Sprintf("Start running flow[%s,%d]", flow.Name, flow.Id))
	var (
		oknum, failednum int
	)
	defer func() {
		logService.Info(flow.Id, fmt.Sprintf("Finish running flow[%s,%d]", flow.Name, flow.Id))
	}()
	if !exec.isRunning(flow) {
		logService.Info(flow.Id, fmt.Sprintf("Flow %s %d state =%d not in running state, ignore", flow.Name, flow.Id, flow.Status))

		return nil
	}

	// load node states
	nodeStateList, oknodesNum, runnodeNum, err := exec.loadNodeStates(flow)
	if err != nil {
		logService.Error(flow.Id, "load node states happen error: "+err.Error())
		exec.rollBackFlowStatus(flow, oknodesNum, runnodeNum, orign_flow_status)
		return err
	}
	//if none node to do task then return and don't change the flow status
	logService.Debug(flow.Id, "load node length: "+strconv.Itoa(len(nodeStateList)))
	if len(nodeStateList) == 0 {
		exec.rollBackFlowStatus(flow, oknodesNum, runnodeNum, orign_flow_status)
		return err
	}

	// get all steps
	steps, stepOps, err := exec.getSteps(flow)
	if err != nil {
		logService.Error(flow.Id, "Get steps fails: "+err.Error())
		exec.rollBackFlowStatus(flow, oknodesNum, runnodeNum, orign_flow_status)
		return err
	}

	if len(steps) < 1 {
		logService.Error(flow.Id, "step length bellow 1 is: "+flow.Options)
		exec.rollBackFlowStatus(flow, oknodesNum, runnodeNum, orign_flow_status)
		return errors.New("step length bellow 1 is: " + flow.Options)

	}

	logService.Debug(flow.Id, fmt.Sprintf("Flow[%d] contains %d nodes", flow.Id, len(nodeStateList)))

	resultChannel := make(chan *models.NodeState, len(nodeStateList))
	for _, nodeState := range nodeStateList {
		toRunState := ToRunNodeState{
			resultChannel:  resultChannel,
			flow:           flow,
			steps:          steps,
			stepOptions:    stepOps,
			nodeState:      nodeState,
		}
		go workNodeQueue.Submit(toRunState)
	}

	resultFlowStatus := orign_flow_status
	if oknodesNum != 0 {
		resultFlowStatus = models.STATUS_SUCCESS
	}

	maxNodeStatesCostTime := 0.0

	for i := 0; i < len(nodeStateList); i++ {
		select {
		case nodeStatesResult := <-resultChannel:
			//if have one nodeState is success then flow is success
			if nodeStatesResult.Status == models.STATUS_SUCCESS {
				resultFlowStatus = models.STATUS_SUCCESS
				oknum++
			} else if nodeStatesResult.Status == models.STATUS_FAILED {
				failednum++
			} else if nodeStatesResult.Status == models.STATUS_STOPPED {
				resultFlowStatus = models.STATUS_STOPPED
			}
			if nodeStatesResult.RunTime > maxNodeStatesCostTime {
				maxNodeStatesCostTime = nodeStatesResult.RunTime
			}
		case <-time.After(time.Minute*15):
			failednum++
			logService.Debug(flow.Id, "RunAndCheck node timeout!")
		}
	}
	close(resultChannel)

	//if all nodeNodestate failed then the flow is failed
	if failednum == len(nodeStateList) && oknodesNum == 0 {
		resultFlowStatus = models.STATUS_FAILED
	}
	exec.SetFlowStatusWithSpenTime(flow, maxNodeStatesCostTime, resultFlowStatus)

	return  nil
}

// Run a batch of task.
func (exec *FlowExecutor) RunNodeState(flow *models.Flow, nodeState *models.NodeState,
	steps []*models.ActionImpl, stepOptions []*models.StepOption, resultChannel chan *models.NodeState) error {

	fid := flow.Id
	logService.Info(fid, fmt.Sprintf("Run Node, flow:%s flowId:%d nodeId:%d", flow.Name, flow.Id, nodeState.Id))

	defer func() {
		logService.Info(fid, fmt.Sprintf("Finish run node, flow:%s flowId:%d", flow.Name, flow.Id))
	}()

	//nodeRunnedTime := nodeState.RunTime
	startStepIndex := nodeState.StepNum

	var stepRunTimeArray []*models.StepRunTime

	//read run step list
	err := json.Unmarshal([]byte(nodeState.StepRunTime), &stepRunTimeArray)
	if err != nil {
		logService.Error(fid, "Fail to load StepRunTime:", nodeState.StepRunTime, ", err:", err)
	}
	//generate nodes runTime steps
	if len(stepRunTimeArray) == 0 {
		for _, step := range steps {
			stepRunTimeElement := &models.StepRunTime{
				Name:    step.Name,
				RunTime: 0.0,
			}
			stepRunTimeArray = append(stepRunTimeArray, stepRunTimeElement)
		}
	} else {
		stepRunTimeArray[startStepIndex].RunTime = 0.0
	}
	//update nodesState to running
	exec.UpdateNodeStatus(steps[startStepIndex].Name, startStepIndex, stepRunTimeArray, nodeState, models.STATUS_INIT)

	for i := startStepIndex; i < len(steps); i++ {
		step := steps[i]
		//read db to judge flow is stopped
		flow, _ := flowService.GetFlowWithRel(fid)
		if flow.Status == models.STATUS_STOPPED || flow.Status == models.STATUS_SUCCESS {
			logService.Warn(fid, "the step: "+step.Name+" begin stop!")
			err := exec.UpdateNodeStatus(step.Name, i, stepRunTimeArray, nodeState, flow.Status)
			if err != nil {
				logService.Error(fid, fmt.Sprintf("update node state db error: %s", err.Error()))
			}
			//put nodeState to chan
			resultChannel <- nodeState // Send nodeState to channel
			return nil
		}
		//check the nodeState has create vm
		if step.Name == "create_vm" && nodeState.Ip != "-" {
			logService.Error(fid, "the node has ip "+nodeState.Ip, "the step: "+step.Name+" error!")
			err := exec.UpdateNodeStatus(step.Name, i, stepRunTimeArray, nodeState, models.STATUS_FAILED)
			if err != nil {
				logService.Error(fid, fmt.Sprintf("update node state db error: %s", err.Error()))
			}
			//put nodeState to chan
			resultChannel <- nodeState // Send nodeState to channel
			return nil
		}
		handler := handler.GetHandler(step.Type)
		if handler == nil {
			logService.Error(fid, fmt.Sprintf("Handler not found for type %s", step.Type))
			nodeState.Status = models.STATUS_FAILED
			err := exec.UpdateNodeStatus(step.Name, i, stepRunTimeArray, nodeState, models.STATUS_FAILED)
			if err != nil {
				logService.Error(fid, fmt.Sprintf("update node state db error: %s", err.Error()))
			}
			//put nodeState to chan
			resultChannel <- nodeState // Send nodeState to channel
			return errors.New("handler not found for type[" + step.Type + "]")
		}
		// get param values
		stepOption := stepOptions[i]
		stepParams := stepOption.Values

		// get retry option
		retryOption := stepOption.Retry
		if retryOption == nil {
			retryOption = dftRetryOpt
		}

		needRunStepNodeState := make([]*models.NodeState, 0)
		needRunStepNodeState = append(needRunStepNodeState, nodeState)

		okNodes, _ := exec.RunStep(handler, step, i, needRunStepNodeState, stepParams, retryOption, stepRunTimeArray)

		if len(okNodes) == 0 {
			nodeState.Status = models.STATUS_FAILED
			logService.Warn(fid, fmt.Sprintf("node %d run fail at step %s", nodeState.Id, step.Name))
			//put nodeState to chan
			resultChannel <- nodeState // Send nodeState to channel
			return nil
		} else {
			nodeState.Status = models.STATUS_RUNNING
			logService.Warn(fid, fmt.Sprintf("node %d run success at step %s", nodeState.Id, step.Name))
		}
	}

	nodeState.Status = models.STATUS_SUCCESS
	err = exec.UpdateNodeStatus(steps[len(steps)-1].Name, len(steps)-1, stepRunTimeArray, nodeState, models.STATUS_SUCCESS)
	if err != nil {
		logService.Error(fid, fmt.Sprintf("update node state db error: %s", err.Error()))
	}
	//put nodeState to chan
	resultChannel <- nodeState // Send nodeState to channel

	logService.Info(fid, fmt.Sprintf("node %d run success all steps", nodeState.Id))

	return nil
}

// runStep runs one step of a batch
func (exec *FlowExecutor) RunStep(h handler.Handler, step *models.ActionImpl, stepIndex int,
	nstates []*models.NodeState, stepParams map[string]interface{},
	retryOption *models.RetryOption, stepRunTimeArray []*models.StepRunTime) ([]*models.NodeState, []*models.NodeState) {

	beginRunStepTime := time.Now()

	paramsBytes, _ := json.Marshal(stepParams)
	paramsJson := string(paramsBytes)

	fid := nstates[0].Flow.Id

	logService.Debug(fid, fmt.Sprintf("Start running step %s params: %s", step.Name, paramsJson))
	defer func() {
		logService.Debug(fid, fmt.Sprintf("Finish running step %s", step.Name))
	}()

	stepRunTimeArray[stepIndex].RunTime = time.Since(beginRunStepTime).Seconds()

	for _, node := range nstates {
		err := exec.UpdateNodeStatus(step.Name, stepIndex, stepRunTimeArray, node, models.STATUS_RUNNING)
		if err != nil {
			logService.Error(fid, "update runNode Step err: ", err)
		}
	}

	toRun := nstates

	var okNodes, errNodes []*models.NodeState
	for i := 0; i < retryOption.RetryTimes+1; i++ {
		// add interval for retry
		if i > 0 {
			time.Sleep(retryInterval * time.Second)
		}

		logService.Debug(fid, fmt.Sprintf("Run step %s for %d times", step.Name, i+1))

		result := h.Handle(step, stepParams, toRun, strconv.Itoa(fid))

		// retry if failed
		if result.Code == handler.CODE_ERROR {
			errNodes = toRun
			msg := fmt.Sprintf("Fail to run step [%s]: %s", step.Name, result.Msg)
			logService.Error(fid, msg)
			continue
		}

		// handle result, retry if failed
		results := result.Result
		if results == nil {
			errNodes = toRun
			msg := fmt.Sprintf("Node results is empty for [%s]", step.Name)
			logService.Error(fid, msg)
			continue
		}

		// update result by every node
		errNodes = make([]*models.NodeState, 0)
		for i, state := range toRun {
			//node := state.Node
			nr := results[i]
			if nr == nil {
				logService.Warn(fid, fmt.Sprintf("Result for node %d %s missing, set it as failed", state.Id, state.Ip))
				state.Status = models.STATUS_FAILED
				state.Log += step.Name + ":" + "<Missing result>\n"
				state.Steps = step.Name
				state.StepNum = stepIndex
				errNodes = append(errNodes, state)
			} else {
				logService.Debug(fid, fmt.Sprintf("Result for node [%d %s] is %d %s", state.Id, state.Ip, nr.Code, nr.Data))
				if nr.Code == models.STATUS_SUCCESS {
					state.Status = models.STATUS_RUNNING
					okNodes = append(okNodes, state)
				} else {
					state.Status = models.STATUS_FAILED
					errNodes = append(errNodes, state)
				}
				state.Log += step.Name + ":" + nr.Data + "\n"
				state.Steps = step.Name
				state.StepNum = stepIndex
			}
			stepRunTimeArray[stepIndex].RunTime = time.Since(beginRunStepTime).Seconds()
			err := exec.UpdateNodeStatus(step.Name, stepIndex, stepRunTimeArray, state, models.STATUS_RUNNING)
			if err != nil {
				logService.Error(fid, fmt.Sprintf("Fail to update state for node[%d %s]", state.Id, state.Ip), err)
			}
		}

		// retry if all nodes fails
		if len(errNodes) == len(toRun) {
			continue
		}

		// all successful
		if len(errNodes) == 0 {
			return okNodes, errNodes
		}
		toRun = errNodes
	}

	if retryOption.IgnoreError {
		// if ignore error is true, set all failed nodes to success
		for _, node := range errNodes {
			exec.UpdateNodeStatus(step.Name, stepIndex, stepRunTimeArray, node, models.STATUS_RUNNING)
		}
		okNodes = nstates
		errNodes = []*models.NodeState{}
	} else {
		for _, node := range errNodes {
			exec.UpdateNodeStatus(step.Name, stepIndex, stepRunTimeArray, node, models.STATUS_FAILED)
		}
	}

	return okNodes, errNodes
}

/**
* xxxxxxx
*
* 注意这里的models.ActionImpl  不对应DB里任何表..只是一个用于获取对应handler的结构
* 此结构里的声明是在models.在每个handler都有一个对应的初始化..
* 结构里的type字段 会在runBatch里handler.GetHandler时..将会被用到
*
*
* @access public
* @param flow 单个需要被执行的任务信息
* @param steps 由本FLOW中options的name字段关键词匹配过的handler信息..用于去获取对应handler..
* @param stepOptions 本FLOW中的options
 */

func (exec *FlowExecutor) getSteps(flow *models.Flow) ([]*models.ActionImpl, []*models.StepOption, error) {

	beego.Debug("Steps are", flow.Options)

	var stepOptions []*models.StepOption
	var steps []*models.ActionImpl
	err := json.Unmarshal([]byte(flow.Options), &stepOptions)
	if err != nil {
		beego.Error("Marshal steps: ", flow.Options, ", err:", err)
		return nil, nil, err
	}

	for _, stepOption := range stepOptions {
		name := stepOption.Name
		step := handler.GetActionImpl(name)
		if step == nil {
			beego.Error("Step [", name, "] not found")
			return nil, nil, errors.New("step [" + name + "] not found")
		}
		steps = append(steps, step)
	}

	return steps, stepOptions, nil
}

func (exec *FlowExecutor) rollBackFlowStatus(flow *models.Flow, oknodesNum int, runnodeNum int, orignStatus int) error {
	//if have run node then flow status is running
	if runnodeNum != 0 {
		err := exec.SetFlowStatus(flow, models.STATUS_RUNNING)
		return err
	}
	//if have not run node then if success node is none then roll previous status
	if oknodesNum == 0 {
		err := exec.SetFlowStatus(flow, models.STATUS_FAILED)
		return err
	} else {
		err := exec.SetFlowStatus(flow, orignStatus)
		return err
	}
}

func (exec *FlowExecutor) loadNodeStates(flow *models.Flow) ([]*models.NodeState, int, int, error) {
	beego.Debug("Load nodes & states for flow", flow.Id)
	states, err := flowService.GetNodeStatusByFlowId(flow.Id)
	if err != nil {
		return states, 0, 0, err
	}
	//select node status is failed,stopped and init
	var oknodes = 0
	var runnodes = 0
	filterNodeList := make([]*models.NodeState, 0)
	for _, state := range states {
		if state.Status == models.STATUS_FAILED && !strings.EqualFold(state.Steps, "create_vm") {
			filterNodeList = append(filterNodeList, state)
		} else if state.Status == models.STATUS_STOPPED {
			filterNodeList = append(filterNodeList, state)
		} else if state.Status == models.STATUS_INIT {
			filterNodeList = append(filterNodeList, state)
		} else if state.Status == models.STATUS_SUCCESS {
			oknodes++
		} else if state.Status == models.STATUS_RUNNING {
			runnodes++
		}
	}

	return filterNodeList, oknodes, runnodes, nil
}

/*
 * Create a flow instance in DB, set its state to INIT
 */
func (exec *FlowExecutor) CreateFlowInstance(name string, flowImpl *models.FlowImpl, pool *models.Pool,
	runType string, stepOps []*models.StepOption, option *ExecOption, opUser string) (*models.Flow, error) {

	bytes, err := json.Marshal(stepOps)
	if err != nil {
		beego.Error("Fail to json dump :", stepOps, err)
		return nil, err
	}
	optionValues := string(bytes)

	fi := &models.Flow{
		Name:        name,
		Options:     optionValues,
		Status:      models.STATUS_INIT,
		Impl:        flowImpl,
		Pool:        pool,
		StepLen:     option.MaxNum,
		OpUser:      opUser,
		FlowType:    runType,
		RunTime:     0.0,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	err = flowService.InsertBase(fi)

	return fi, err
}

// Create node states for every node in given task
func (exec *FlowExecutor) CreateNodeStates(flow *models.Flow, nodes []*models.NodeState) ([]*models.NodeState, error) {
	states := make([]*models.NodeState, len(nodes))
	for i, node := range nodes {
		state := &models.NodeState{
			Ip:          node.Ip,
			VmId:        node.VmId,
			Flow:        flow,
			Pool:        node.Pool,
			Status:      models.STATUS_INIT,
			Log:         "",
			Steps:       "[]",
			StepRunTime: "[]",
			StepNum:     0,
			RunTime:     0.0,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
			Deleted:     false,
			NodeType:    node.NodeType,
		}
		// TODO batch insert all states
		beego.Debug("Create node state for node: ip=", node.Ip)
		err := flowService.InsertBase(state)
		if err != nil {
			beego.Error("Fail to create state, error: ", err)
		}

		states[i] = state
	}

	return states, nil
}

func (exec *FlowExecutor) getNodeRunTime(stepRunTimeArray []*models.StepRunTime) float64 {

	nodeStateRunSpendTime := 0.0
	for _, stepTime := range stepRunTimeArray {
		nodeStateRunSpendTime += stepTime.RunTime
	}

	return nodeStateRunSpendTime
}

func (exec *FlowExecutor) UpdateNodeStatus(stepName string, stepIndex int,
	stepRunTimeArray []*models.StepRunTime, state *models.NodeState, stateCode int) error {

	nodeStateRunSpendTime := exec.getNodeRunTime(stepRunTimeArray)

	mergedBytes, _ := json.Marshal(stepRunTimeArray)
	merged := string(mergedBytes)
	state.Status = stateCode
	if stepName != "" {
		state.Steps = stepName
	}
	state.RunTime = nodeStateRunSpendTime
	state.StepNum = stepIndex
	state.StepRunTime = merged
	state.UpdatedTime = time.Now()

	err := flowService.UpdateNode(state)
	if err != nil {
		beego.Error("update node state", state.Ip, "status =", stateCode, "err :", err)
		return err
	}
	beego.Debug("Set node state", state.Ip, " step =", state.Steps, "status =", stateCode)
	return nil
}

func (exec *FlowExecutor) MergeParams(options []*models.StepOption,
	overrideParams map[string]interface{}) {

	opMap := make(map[string]*models.StepOption)
	for _, op := range options {
		v := op
		opMap[v.Name] = v
	}

	for k, v := range overrideParams {
		op := opMap[k]
		if op == nil {
			beego.Warn("step option of", k, "not found, skip")
			continue
		}

		values, ok := v.(map[string]interface{})
		if !ok {
			beego.Error("Bad override params:", k, "=", v)
			continue
		}

		old := op.Values
		for pk, pv := range values {
			beego.Debug("override param:", k, ": {", pk, ":", pv, "}")
			old[pk] = pv
		}
	}
}


