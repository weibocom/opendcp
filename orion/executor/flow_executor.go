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
	"sync"
	"time"

	"github.com/astaxie/beego"

	"weibo.com/opendcp/orion/handler"
	"weibo.com/opendcp/orion/models"
	"weibo.com/opendcp/orion/service"
	"net"
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
	checkTimeout = 28 //timeout minutes of wait node result

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

	fis, runNodes, err := exec.Create(flowImpl, name, option, nodes, context)
	if err != nil {
		return err
	}

	for _, fi := range fis {
		runNodes := runNodes[fi.Id]
		err = exec.Start(fi, runNodes)
		if err != nil {
			return err
		}
	}

	return nil
}

// Create creates a flow instance with given template id, and nodes.
func (exec *FlowExecutor) Create(flowImpl *models.FlowImpl, name string, option *ExecOption,
	nodes []*models.NodeState, context map[string]interface{}) ([]*models.Flow, map[int][]*models.NodeState, error) {

	beego.Info("Create new task from template[", flowImpl.Id, "]")

	var (
		nodesarray = make(map[int][]*models.NodeState)
		instances  []*models.Flow
		stepOps    []*models.StepOption
		newNodes   = make(map[int][]*models.NodeState)
	)

	if option == nil {
		beego.Error("Option is nil")
		return instances, newNodes, errors.New("option is nil")
	}

	if len(nodes) == 0 {
		beego.Error("nodes is empty")
		return instances, newNodes, errors.New("nodes is empty")
	}

	// make sure all nodes are in the same pool
	for _, node := range nodes {
		pid := node.Pool.Id
		nodesarray[pid] = append(nodesarray[pid], node)
	}

	err := json.Unmarshal([]byte(flowImpl.Steps), &stepOps)
	if err != nil {
		beego.Error("Bad step options: ", flowImpl.Steps, "[err]: ", err)
		return instances, newNodes, errors.New("Bad step options: " + flowImpl.Steps + ", err:" + err.Error())
	}

	//merge override params with params
	overrideParams, ok := context["overrideParams"].(map[string]interface{})
	if !ok {
		beego.Error("bad overrideParams:", context["overrideParams"])
		return instances, newNodes, errors.New("bad overrideParams!")
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
		return instances, newNodes, errors.New("bad opUser!")
	}

	for pid, node := range nodesarray {
		poolID := pid
		poolNode := node
		instance, err := exec.CreateFlowInstance(name, flowImpl, &models.Pool{Id: poolID}, poolNode[0].NodeType, stepOps, option, opUser)
		if err != nil {
			beego.Error("Fail to create flow instance", err)
			return instances, newNodes, err
		}
		createNodes, err := exec.CreateNodeStates(instance, poolNode)
		if err != nil {
			beego.Error("Fail to create node states for flow: ", instance.Name, err)
			flowService.DeleteBase(instance)
			return instances, newNodes, err
		}
		newNodes[instance.Id] = createNodes
		instances = append(instances, instance)
	}

	return instances, newNodes, nil
}

// Start starts an existing flow instance.
func (exec *FlowExecutor) Start(flow *models.Flow, startNodes []*models.NodeState) error {
	lock.Lock()
	defer lock.Unlock()

	//orign_flow_status := flow.Status

	//if exec.isFlowInState(flow, models.STATUS_RUNNING) {
	//	logInfo := "Flow " + flow.Name + " is in state running: " + strconv.Itoa(flow.Status) + " do not start"
	//	logService.Info(flow.Id, logInfo)
	//
	//	return errors.New(logInfo)
	//}

	exec.SetFlowStatus(flow, models.STATUS_RUNNING)

	// run the flow
	job := func() error {
		logService.Info(flow.Id, "Run flow...")

		err := exec.RunFlow(flow, startNodes)
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

	return flowService.UpdateFlowStatus(flow)
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
func (exec *FlowExecutor) RunFlow(flow *models.Flow, runNodes []*models.NodeState) (err error) {

	var (
		nodeStateList []*models.NodeState
		steps         []*models.ActionImpl
		stepOps       []*models.StepOption
		resultChannel chan *models.NodeState
	)

	logService.Info(flow.Id, fmt.Sprintf("Start running flow[%s,%d]", flow.Name, flow.Id))

	defer func() {
		close(resultChannel)
		logService.Info(flow.Id, fmt.Sprintf("Finish running flow[%s,%d]", flow.Name, flow.Id))
	}()

	if !exec.isFlowInState(flow, models.STATUS_RUNNING) {
		logService.Info(flow.Id, fmt.Sprintf("Flow %s %d state =%d in not running state, ignore", flow.Name, flow.Id, flow.Status))
		return nil
	}

	// get all steps
	if steps, stepOps, err = exec.getSteps(flow); err != nil {
		logService.Error(flow.Id, "Get steps fails: "+err.Error())
		exec.terminateFlow(flow)
		return err
	}

	if len(steps) < 1 {
		logService.Error(flow.Id, "step length bellow 1 is: ", flow.Options)
		exec.terminateFlow(flow)
		return errors.New("step length bellow 1 is: " + flow.Options)

	}

	// load node states
	if nodeStateList, err = exec.loadStartNodeStates(flow, runNodes); err != nil {
		logService.Error(flow.Id, "load nodes error: ", err.Error())
		exec.terminateFlow(flow)
		return err
	}
	//if none node to do task then return and don't change the flow status
	logService.Info(flow.Id, "load nodes length: ", len(nodeStateList))
	if len(nodeStateList) == 0 {
		logService.Info(flow.Id, "none nodes to run")
		exec.terminateFlow(flow)
		return nil
	}
	// set all node to start



	logService.Debug(flow.Id, fmt.Sprintf("Flow[%d] contains %d nodes", flow.Id, len(nodeStateList)))

	resultChannel = make(chan *models.NodeState, len(nodeStateList))
	for _, nodeState := range nodeStateList {
		toRunState := ToRunNodeState{
			resultChannel: resultChannel,
			flow:          flow,
			steps:         steps,
			stepOptions:   stepOps,
			nodeState:     nodeState,
		}
		go workNodeQueue.Submit(toRunState)
	}
	//wait until all nodes have done
	if err = exec.waitNodesResult(flow, resultChannel, nodeStateList); err != nil {
		logService.Warn(flow.Id, err.Error())
	}

	if err = exec.terminateFlow(flow); err != nil {
		return err
	}

	return nil

}

// Run a batch of task.
func (exec *FlowExecutor) RunNodeState(flow *models.Flow, nodeState *models.NodeState,
	steps []*models.ActionImpl, stepOptions []*models.StepOption, resultChannel chan *models.NodeState) (err error) {

	var (
		stepRunTimeArray []*models.StepRunTime
		fid              = flow.Id
		startStepIndex   = nodeState.StepNum
		step             *models.ActionImpl
		runStepIndex     int
		isStopped        bool
	)

	logService.Info(fid, fmt.Sprintf("Run Node, flow:%s flowId:%d nodeId:%d", flow.Name, flow.Id, nodeState.Id))

	defer func() {
		resultChannel <- nodeState // Send nodeState to channel
		logService.Info(fid, fmt.Sprintf("Finish run node, flow:%s flowId:%d", flow.Name, flow.Id))
	}()

	if stepRunTimeArray, err = exec.generateRunStepTime(nodeState, steps); err != nil {
		logService.Error(fid, "Fail to load StepRunTime:", nodeState.StepRunTime, ", err: ", err)
		nodeState.Status = models.STATUS_FAILED
		nodeState.UpdatedTime = time.Now()
		flowService.ChangeNodeStatusById(nodeState)
		return err
	}
	//update nodesState to running
	if isStopped, _ = exec.isStoppedNode(flow, nodeState); isStopped {
		exec.UpdateNodeStatus(steps[startStepIndex].Name, startStepIndex, stepRunTimeArray, nodeState, nodeState.Status)
		return nil
	}

	for runStepIndex = startStepIndex; runStepIndex < len(steps); runStepIndex++ {
		step = steps[runStepIndex]
		//read db to judge node is stopped
		if isStopped, _ = exec.isStoppedNode(flow, nodeState); isStopped {
			logService.Info(fid, "the step: "+step.Name+" is terminate!")
			break
		}
		//check the nodeState has create vm
		if step.Name == "create_vm" && exec.HaveIp(nodeState.Ip) {
			logService.Warn(fid, "the node has already create the node ip: ", nodeState.Ip)
			//nodeState.Status = models.STATUS_FAILED
			continue
		}
		doHandler := handler.GetHandler(step.Type)
		if doHandler == nil {
			logService.Error(fid, fmt.Sprintf("Handler not found for type %s", step.Type))
			nodeState.Status = models.STATUS_FAILED
			break
		}
		// get param values
		stepOption := stepOptions[runStepIndex]
		stepParams := stepOption.Values

		// get retry option
		retryOption := stepOption.Retry
		if retryOption == nil {
			retryOption = dftRetryOpt
		}

		needRunStepNodeState := make([]*models.NodeState, 0)
		needRunStepNodeState = append(needRunStepNodeState, nodeState)

		okNodes, errNodes := exec.RunStep(doHandler, step, runStepIndex, needRunStepNodeState, stepParams, retryOption, stepRunTimeArray)

		logService.Error(fid, fmt.Sprintf("node run result: ok(%d) err(%d)", len(okNodes), len(errNodes)))

		if len(okNodes) == 0 && len(errNodes) != 0{
			nodeState = errNodes[0]
			logService.Error(fid, fmt.Sprintf("node %d status %d run fail at step %s", nodeState.Id, nodeState.Status, step.Name))
			break
		} else if len(errNodes) == 0 && len(okNodes) != 0 {
			nodeState = okNodes[0]
			if nodeState.Status != models.STATUS_RUNNING {
				logService.Error(fid, fmt.Sprintf("node %d status %d stop at step %s", nodeState.Id, nodeState.Status, step.Name))
				break
			} else {
				logService.Info(fid, fmt.Sprintf("node %d status %d run success at step %s", nodeState.Id, nodeState.Status, step.Name))
			}
		}else {
			logService.Error(fid, fmt.Sprintf("Lost node %d status %d run at step %s", nodeState.Id, nodeState.Status, step.Name))
			break
		}
	}

	if runStepIndex == len(steps) {
		logService.Info(fid, fmt.Sprintf("node %d run success all steps", nodeState.Id))
		err = exec.UpdateNodeStatus(steps[runStepIndex-1].Name, runStepIndex, stepRunTimeArray, nodeState, models.STATUS_SUCCESS)
	} else {
		logService.Info(fid, fmt.Sprintf("node %d run at step: %s was terminated", nodeState.Id, steps[runStepIndex].Name))
		if nodeState.Status == models.STATUS_RUNNING {
			nodeState.Status = models.STATUS_FAILED
		}
		err = exec.UpdateNodeStatus(steps[runStepIndex].Name, runStepIndex, stepRunTimeArray, nodeState, nodeState.Status)
	}
	if err != nil {
		logService.Error(fid, fmt.Sprintf("update node state db error: %s", err.Error()))
	}

	return nil
}

// runStep runs one step of a batch
func (exec *FlowExecutor) RunStep(h handler.Handler, step *models.ActionImpl, stepIndex int,
	nstates []*models.NodeState, stepParams map[string]interface{},
	retryOption *models.RetryOption, stepRunTimeArray []*models.StepRunTime) ([]*models.NodeState, []*models.NodeState) {

	var (
		beginRunStepTime = time.Now()
		paramsBytes, _   = json.Marshal(stepParams)
		paramsJson       = string(paramsBytes)
		fid              = nstates[0].Flow.Id
		toRun            = make([]*models.NodeState, 0)
		okNodes          = make([]*models.NodeState, 0)
		errNodes         = make([]*models.NodeState, 0)
	)

	defer func() {
		logService.Info(fid, fmt.Sprintf("Finish running step %s", step.Name))
	}()

	logService.Info(fid, fmt.Sprintf("Start running step %s params: %s", step.Name, paramsJson))

	for _, node := range nstates {
		if isStopped, _ := exec.isStoppedNode(node.Flow, node); isStopped {
			okNodes = append(okNodes, node)
		} else {
			toRun = append(toRun, node)
			stepRunTimeArray[stepIndex].RunTime = time.Since(beginRunStepTime).Seconds()
			err := exec.UpdateNodeStatus(step.Name, stepIndex, stepRunTimeArray, node, models.STATUS_RUNNING)
			if err != nil {
				logService.Error(fid, "update runNode Step err: ", err)
			}
		}
	}

	for i := 0; i < retryOption.RetryTimes+1; i++ {
		// add interval for retry
		if i > 0 {
			time.Sleep(retryInterval * time.Second)
		}

		TrytoRun := make([]*models.NodeState, 0)
		for _, node := range toRun {

			stepRunTimeArray[stepIndex].RunTime = time.Since(beginRunStepTime).Seconds()
			if err := exec.UpdateNodeRunTime(stepRunTimeArray, node); err != nil {
				logService.Error(fid, "update runNode Step err: ", err)
			}

			if isStopped, _ := exec.isStoppedNode(node.Flow, node); isStopped {
				okNodes = append(okNodes, node)
			} else {
				TrytoRun = append(TrytoRun, node)
			}

		}
		toRun = TrytoRun

		if len(toRun) == 0 {
			return okNodes, errNodes
		}

		logService.Info(fid, fmt.Sprintf("Run step %s for %d times", step.Name, i+1))

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
				state.Log += step.Name + ":" + nr.Data + "\n"
				state.Steps = step.Name
				state.StepNum = stepIndex
				if nr.Code == models.STATUS_SUCCESS {
					state.Status = models.STATUS_RUNNING
					okNodes = append(okNodes, state)
				} else {
					state.Status = models.STATUS_FAILED
					errNodes = append(errNodes, state)
				}
			}
			stepRunTimeArray[stepIndex].RunTime = time.Since(beginRunStepTime).Seconds()
			if err := exec.UpdateNodeRunTime(stepRunTimeArray, state); err != nil {
				logService.Error(fid, "update runNode Step err: ", err)
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
			node.Status = models.STATUS_RUNNING
		}
		okNodes = append(okNodes, errNodes...)
		errNodes = make([]*models.NodeState, 0)
	} else {
		for _, node := range errNodes {
			node.Status = models.STATUS_FAILED
		}
	}

	return okNodes, errNodes
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

func (exec *FlowExecutor) UpdateNodeRunTime(stepRunTimeArray []*models.StepRunTime, state *models.NodeState) error {

	nodeStateRunSpendTime := exec.getNodeRunTime(stepRunTimeArray)

	mergedBytes, _ := json.Marshal(stepRunTimeArray)
	merged := string(mergedBytes)
	state.RunTime = nodeStateRunSpendTime
	state.StepRunTime = merged
	state.UpdatedTime = time.Now()

	err := flowService.UpdateNodeRunTime(state)
	if err != nil {
		beego.Error("update node state", state.Ip, "status =", state.Status, "err :", err)
		return err
	}
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

func (exec *FlowExecutor) HaveIp(ip string) bool{

	if parseIp := net.ParseIP(ip); parseIp == nil{
		return false
	}

	return true
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

func (exec *FlowExecutor) terminateFlow(flow *models.Flow) (err error) {

	defer func() {
		if err != nil {
			logService.Error(flow.Id, "Terminate Flow is err: ", err.Error())
		} else {
			logService.Info(flow.Id, "Terminate Flow status: ", flow.Status)
		}
	}()

	var (
		states                 []*models.NodeState
		successNum, stoppedNum int
		failedNum, runningNum  int
		inintNum               int
		flowStatus             int
	)

	if states, err = flowService.GetAllNodeStatesByFlowId(flow.Id); err != nil {
		return err
	}

	for _, ns := range states {
		if ns.Status == models.STATUS_INIT {
			inintNum++
		} else if ns.Status == models.STATUS_RUNNING {
			runningNum++
		} else if ns.Status == models.STATUS_STOPPED {
			stoppedNum++
		} else if ns.Status == models.STATUS_SUCCESS {
			successNum++
		} else if ns.Status == models.STATUS_FAILED {
			failedNum++
		}
	}

	if runningNum != 0 {
		flowStatus = models.STATUS_RUNNING
	} else if successNum != 0 {
		flowStatus = models.STATUS_SUCCESS
	} else if inintNum == len(states) {
		flowStatus = models.STATUS_INIT
	} else if stoppedNum == len(states) {
		flowStatus = models.STATUS_STOPPED
	} else if failedNum == len(states) {
		flowStatus = models.STATUS_FAILED
	} else {
		flowStatus = models.STATUS_FAILED
	}

	if err = exec.SetFlowStatusWithSpenTime(flow, flow.RunTime, flowStatus); err != nil {
		return err
	}

	return nil
}

func (exec *FlowExecutor) getNodeRunTime(stepRunTimeArray []*models.StepRunTime) float64 {

	var nodeStateRunSpendTime = 0.0

	for _, stepTime := range stepRunTimeArray {
		nodeStateRunSpendTime += stepTime.RunTime
	}

	return nodeStateRunSpendTime
}

func (exec *FlowExecutor) loadStartNodeStates(flow *models.Flow, runNodes []*models.NodeState) (states []*models.NodeState, err error) {
	logService.Info(flow.Id, "Load nodes & states for flow", flow.Id)
	defer func() {
		logService.Info(flow.Id, "Finsh load nodes & states for flow", flow.Id)
	}()

	if states, err = flowService.GetNodeStatusByFlowId(flow.Id); err != nil {
		return states, err
	}

	filterNodeList := make([]*models.NodeState, 0)
	for _, state := range states {
		for _, rn := range runNodes {
			if state.Id != rn.Id {
				continue
			}
			if state.Status != models.STATUS_RUNNING && state.Status != models.STATUS_SUCCESS {
				state.Status = models.STATUS_RUNNING
				state.UpdatedTime = time.Now()
				if err := flowService.ChangeNodeStatusById(state); err != nil{
					logService.Info(flow.Id, "update load node status err: ", err.Error())
					continue
				}
				filterNodeList = append(filterNodeList, state)
			}
			break
		}
	}

	return filterNodeList, nil
}

func (exec *FlowExecutor) waitNodesResult(flow *models.Flow, resultChannel chan *models.NodeState, nodes []*models.NodeState) error {

	var maxTime = flow.RunTime

	for i := 0; i < len(nodes); i++ {
		select {
		case ns := <-resultChannel:
			if ns.RunTime > maxTime {
				maxTime = ns.RunTime
			}
		case <-time.After(time.Minute * checkTimeout):
			logService.Error(flow.Id, "Get node run result timeout")
			if err := exec.handleNodeRunTimeOut(flow); err != nil {
				logService.Error(flow.Id, "handle node run timeout err : ", err.Error())
			} else {
				logService.Info(flow.Id, "handle node run timeout done : ", err.Error())
			}
		}
	}

	flow.RunTime = maxTime

	return nil
}

func (exec *FlowExecutor) handleNodeRunTimeOut(flow *models.Flow) (err error) {
	var (
		runningNodes = make([]*models.NodeState, 0)
		runLongNode  *models.NodeState
		runLongTime  = 0.0
	)
	if runningNodes, err = flowService.GetAllNodeStatusByFlowId(flow.Id, models.STATUS_RUNNING); err != nil {
		return err
	}
	if len(runningNodes) == 0 {
		return nil
	}
	for _, runNode := range runningNodes {
		if runNode.RunTime >= runLongTime {
			runLongNode = runNode
			runLongTime = runNode.RunTime
		}
	}
	runLongNode.Status = models.STATUS_FAILED
	runLongNode.UpdatedTime = time.Now()

	if err = flowService.ChangeNodeStatusById(runLongNode); err != nil {
		err = errors.New(fmt.Sprintf("update node status err: %s", err.Error()))
	}
	return err
}
func (exec *FlowExecutor) generateRunStepTime(nodeState *models.NodeState, steps []*models.ActionImpl) (stepRunTimeArray []*models.StepRunTime, err error) {

	if err = json.Unmarshal([]byte(nodeState.StepRunTime), &stepRunTimeArray); err != nil {
		return stepRunTimeArray, err
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
		stepRunTimeArray[nodeState.StepNum].RunTime = 0.0
	}
	return stepRunTimeArray, nil
}

func (exec *FlowExecutor) checkNodeStatus(node *models.NodeState) (*models.NodeState, error) {

	var (
		dbNode *models.NodeState
		err    error
	)

	if dbNode, err = flowService.GetNodeById(node.Id); err != nil {
		return dbNode, err
	}

	return dbNode, nil
}

func (exec *FlowExecutor) isStoppedNode(flow *models.Flow, node *models.NodeState) (bool, error) {

	var (
		dbNode *models.NodeState
		err    error
	)

	if flow, err = flowService.GetFlowWithRel(flow.Id); err != nil {
		return false, err
	}

	if flow.Status == models.STATUS_SUCCESS || flow.Status == models.STATUS_STOPPED {
		node.Status = flow.Status
		return true, nil
	}

	if dbNode, err = exec.checkNodeStatus(node); err != nil {
		return false, err
	}

	if dbNode.Deleted {
		node.Deleted = dbNode.Deleted
		return true, nil
	}

	if dbNode.Status == models.STATUS_STOPPED || dbNode.Status == models.STATUS_SUCCESS {
		node.Status = dbNode.Status
		return true, nil
	}

	return false, nil
}
