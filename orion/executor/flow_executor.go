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
	"weibo.com/opendcp/orion/utils"
	"strings"
)

var (
	// Executor is the singleton FlowExecutor instance to run all tasks.
	Executor = &FlowExecutor{}
	flowService = service.Flow
	logService = service.Logs

	// Default retry option
	dftRetryOpt = &models.RetryOption{
		RetryTimes:  0,
		IgnoreError: false,
	}

	// lock for running jobs
	lock = &sync.Mutex{}

	// workers for each pool
	workers = make(map[int]*Worker)
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

// Run creates a flow instance from the given template id, and nodes, and
// starts the flow instance.
func (exec *FlowExecutor) Run(tplID int, name string, option *ExecOption,
	nodes []*models.Node,context map[string]interface{}) error {

	fis, err := exec.Create(tplID, name, option, nodes, context)
	if err != nil {
		return err
	}

	for _, fi  := range fis{
		ins := fi
		err = exec.Start(ins)
		if err  != nil{
			return err
		}
	}

	return nil
}

// Create creates a flow instance with given template id, and nodes.
func (exec *FlowExecutor) Create(tplID int, name string, option *ExecOption,
	nodes []*models.Node, context map[string]interface{}) ([]*models.Flow, error) {

	beego.Info("Create new task from template[", tplID, "]")

	nodesarray := make(map[int][]*models.Node)
	var instances []*models.Flow

	flow := &models.FlowImpl{Id: tplID}
	err := flowService.GetBase(flow)
	if err != nil {
		beego.Error("template not found " + fmt.Sprint(tplID))
		return nil, errors.New("Template not found " + fmt.Sprint(tplID))
	}

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
	err = json.Unmarshal([]byte(flow.Steps), &stepOps)
	if err != nil {
		beego.Error("Bad step options: ", flow.Steps, "[err]: ", err)
		return nil, errors.New("Bad step options: " + flow.Steps + ", err:" + err.Error())
	}

	//merge override params with params
	overrideParams, ok:= context["overrideParams"].(map[string]interface{})
	if !ok {
		beego.Error("bad overrideParams:", context["overrideParams"])
		return nil, errors.New("bad overrideParams!")
	}
	exec.mergeParams(stepOps, overrideParams)

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


	for pid, node := range nodesarray{
		poolID := pid
		poolNode := node
		instance, err := exec.createFlowInstance(name, flow, &models.Pool{Id: poolID}, stepOps, option, opUser)
		if err != nil {
			beego.Error("Fail to create flow instance", err)
			return nil, err
		}
		states, err := exec.createNodeStates(instance, poolNode)
		if err != nil {
			beego.Error("Fail to create node states for flow: ", instance.Name, err)
			flowService.DeleteBase(instance)
			return nil, err
		}
		// create batches
		err = exec.createBatches(instance, states, option.MaxNum)
		if err != nil {
			beego.Error("Fail to create batches for flow: ", instance.Name, err)
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

	correlationId := exec.getCorrelationId(flow.Id,0)

	if exec.isFlowInState(flow, models.STATUS_RUNNING){

		logInfo := "Flow " + flow.Name + " is not in state init/stopped: " + strconv.Itoa(flow.Status)
		logService.Info(flow.Id,0,correlationId,logInfo)

		return errors.New(logInfo)
	}

	exec.setFlowStatus(flow, models.STATUS_RUNNING)

	// run the flow
	job := func() error {
		logService.Info(flow.Id,0,correlationId,"Run flow...")

		err := exec.runFlow(flow)
		if err != nil {
			logService.Error(flow.Id,0,correlationId,"Run flow error",err)
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

	return exec.setFlowStatus(flow, models.STATUS_STOPPED)
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

	return exec.setFlowStatus(flow, models.STATUS_SUCCESS)
}

func (exec *FlowExecutor) setFlowStatus(flow *models.Flow, status int) error {
	beego.Debug("Set flow", flow.Name, "status =", status)
	flow.Status = status
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
func (exec *FlowExecutor) runFlow(flow *models.Flow) error {
	correlationId := exec.getCorrelationId(flow.Id,0)

	logService.Info(flow.Id,0,correlationId,fmt.Sprintf("Start running flow[%s,%d]",flow.Name,flow.Id))
	defer func (){
		logService.Info(flow.Id,0,correlationId,fmt.Sprintf("Finish running flow[%s,%d]",flow.Name,flow.Id))
	}()

	if !exec.isRunning(flow) {
		logService.Info(flow.Id,0,correlationId,fmt.Sprintf("Flow %s %d state =%d not in running state, ignore",flow.Name,flow.Id,flow.Status))
		return nil
	}

	// get all steps
	steps, stepOps, err := exec.getSteps(flow)
	if err != nil {
		logService.Error(flow.Id,0,correlationId,"Get steps fails: " + err.Error())
		return err
	}

	// load batches
	logService.Info(flow.Id,0,correlationId,fmt.Sprintf("Load batches for flow name:%s,id:%d" , flow.Name ,flow.Id))

	var batches []*models.FlowBatch
	flowService.ListByPageWithFilter(0, 1000, models.FlowBatch{}, &batches, "Flow", flow.Id)
	for _, batch := range batches {
		batch.Flow = flow
	}

	logService.Info(flow.Id,0,correlationId,fmt.Sprintf("There are %d batches for flow %s",len(batches),flow.Name))

	// load node states
	idStateMap, err := exec.loadNodeStates(flow)
	if err != nil {
		return err
	}

	//如果没有要执行的节点，即所有执行的节点都成功了
	if len(idStateMap) == 0{
		return nil
	}

	// execute flow by batch
	for i, batch := range batches {
		correlationId = exec.getCorrelationId(flow.Id,batch.Id)

		logService.Info(flow.Id,0,correlationId,fmt.Sprintf("Execute batch %d of flow %s ",i+1,flow.Name))

		if !exec.isRunning(flow) {
			logService.Info(flow.Id,0,correlationId,fmt.Sprintf("Flow %s %d state =%d not in running state, ignore",flow.Name,flow.Id,flow.Status))
			return nil
		}

		switch batch.Status {
		case models.STATUS_SUCCESS:
			logService.Info(flow.Id,0,correlationId,fmt.Sprintf("Batch %d already finished with status %d , skip",i+1,batch.Status))
			continue
		case models.STATUS_INIT:
			logService.Info(flow.Id,0,correlationId,fmt.Sprintf("Batch %d not stated yet, start",i+1))

			err = exec.runBatch(batch, idStateMap, steps, stepOps)
		case models.STATUS_RUNNING:
			logService.Info(flow.Id,0,correlationId,fmt.Sprintf("Batch %d already running, rerun",i+1))

			err = exec.runBatch(batch, idStateMap, steps, stepOps)
		case models.STATUS_FAILED:
			logService.Info(flow.Id,0,correlationId,fmt.Sprintf("Batch %d stated fail, start",i+1))

			err = exec.runBatch(batch, idStateMap, steps, stepOps)
		}

		if err != nil {
			logService.Error(flow.Id,0,correlationId,fmt.Sprintf("Batch %d fails:",i+1),err)

			exec.setFlowStatus(flow, models.STATUS_FAILED)
			return err
		}
	}

	exec.setFlowStatus(flow, models.STATUS_SUCCESS)
	return nil
}

// Run a batch of task.
func (exec *FlowExecutor) runBatch(batch *models.FlowBatch, stateMap map[int]*models.NodeState,
	steps []*models.ActionImpl, stepOptions []*models.StepOption) error {

	fid := batch.Flow.Id
	correlationId := exec.getCorrelationId(fid,batch.Id)

	logService.Info(fid,batch.Id,correlationId,fmt.Sprintf("Run batch, flow:%s batchId:%d", batch.Flow.Name,batch.Id))

	defer func (){
		logService.Info(fid,batch.Id,correlationId,fmt.Sprintf("Finish run batch, flow:%s batchId:%d", batch.Flow.Name,batch.Id))
	}()


	// load nodes & node states
	states, err := exec.getBatchNodeStatesByStates(batch, stateMap)
//	states, err := exec.getBatchNodeStates(batch, stateMap)
	if err != nil {
		return err
	}

	logService.Debug(fid,batch.Id,correlationId,fmt.Sprintf("Batch[%d] contains %d nodes", batch.Id, len(states)))

	//记录每个节点是否开始执行
	isRun := make([]int, len(states))

	all := states
	for i, step := range steps {
		var tempall []*models.NodeState
		//需要对all进行筛选，当步骤已经成功的不需要重新再次执行
		for temi := 0; temi < len(states); temi++ {
			//当前节点已经从错误步骤开始
			if(isRun[temi] != 0){
				tempall = append(tempall, states[temi])
			}else if(strings.Compare(states[temi].Steps, step.Name) == 0){
				//找到当前节点错误步骤
				isRun[temi] = 1;
				tempall = append(tempall, states[temi])
			}
		}
		if len(tempall) == 0 {
			continue;
		}

		logService.Debug(fid,batch.Id,correlationId,fmt.Sprintf("Run step %s of batch[%d]",step.Name,batch.Id))

		// run step using handler
		handler := handler.GetHandler(step.Type)
		if handler == nil {
			logService.Error(fid,batch.Id,correlationId,fmt.Sprintf("Handler not found for type %s", step.Type))

			exec.allFailed(batch, step, tempall)
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

		// use flow-batch id as correlation id
		correlationId := exec.getCorrelationId(batch.Flow.Id, batch.Id)

		tempall, _ = exec.runStep(handler, step, tempall, stepParams, retryOption, correlationId)

		// if all nodes fail in this batch, then we assume this task fails
		if len(tempall) == 0 {
			logService.Error(fid,batch.Id,correlationId,fmt.Sprintf("Flow %s fails at batch[%d] step[%s]", batch.Flow.Name,batch.Id,step.Name))

			exec.allFailed(batch, step, tempall)
			return errors.New("Fail at step " + step.Name)
		}
	}

	exec.allSuccess(batch, all)

	return nil
}



// runStep runs one step of a batch
func (exec *FlowExecutor) runStep(h handler.Handler, step *models.ActionImpl,
	nstates []*models.NodeState, stepParams map[string]interface{},
	retryOption *models.RetryOption, correlationId string) ([]*models.NodeState, []*models.NodeState) {

	paramsBytes , _ := json.Marshal(stepParams)
	paramsJson := string(paramsBytes)

	fid := nstates[0].Flow.Id
	batchId := nstates[0].Batch.Id

	logService.Debug(fid,batchId,correlationId,fmt.Sprintf("Start running step %s params: %s",step.Name, paramsJson))
	defer func (){
		logService.Debug(fid,batchId,correlationId,fmt.Sprintf("Finish running step %s",step.Name))
	}()


	exec.updateStepStatus(nstates, step.Name, models.STATUS_RUNNING)

	toRun := nstates
	var okNodes, errNodes []*models.NodeState
	for i := 0; i < retryOption.RetryTimes+1; i++ {
		// add interval for retry
		if i > 0 {
			time.Sleep(retryInterval * time.Second)
		}

		logService.Debug(fid,batchId,correlationId,fmt.Sprintf("Run step %s for %d times",step.Name, i+1))

		result := h.Handle(step, stepParams, toRun, correlationId)

		// retry if failed
		if result.Code == handler.CODE_ERROR {
			errNodes = toRun
			msg := fmt.Sprintf("Fail to run step [%s]: %s", step.Name, result.Msg)
			logService.Error(fid,batchId,correlationId,msg)

			continue
		}

		// handle result, retry if failed
		results := result.Result
		if results == nil {
			errNodes = toRun
			msg := fmt.Sprintf("Node results is empty for [%s]", step.Name)
			logService.Error(fid,batchId,correlationId,msg)

			continue
		}

		// update result by every node
		errNodes = make([]*models.NodeState, 0)
		for i, state := range toRun {
			node := state.Node
			nr := results[i]
			if nr == nil {
				logService.Warn(fid,batchId,correlationId,fmt.Sprintf("Result for node %d %s missing, set it as failed",node.Id, node.Ip))

				state.Status = models.STATUS_FAILED
				state.Log += step.Name + ":" + "<Missing result>\n"
				errNodes = append(errNodes, state)
			} else {
				logService.Debug(fid,batchId,correlationId,fmt.Sprintf("Result for node [%d %s] is %d %s", node.Id, node.Ip,nr.Code, nr.Data))

				if nr.Code == models.STATUS_SUCCESS {
					okNodes = append(okNodes, state)
					state.Status = models.STATUS_RUNNING
				} else {
					state.Status = models.STATUS_FAILED
					errNodes = append(errNodes, state)
				}
				state.Log += step.Name + ":" + nr.Data + "\n"

				// update progress
				state.Steps = step.Name
			}
			state.UpdatedTime = time.Now()

			err := flowService.UpdateBase(state)
			if err != nil {
				logService.Error(fid,batchId,correlationId,fmt.Sprintf("Fail to update state for node[%d %s]", node.Id, node.Ip))
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
		exec.updateStepStatus(errNodes, "", models.STATUS_SUCCESS)
		okNodes = nstates
		errNodes = []*models.NodeState{}
	} else {
		exec.updateStepStatus(errNodes, "", models.STATUS_FAILED)
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

/*
func (exec *FlowExecutor) loadFlowOption(flow *models.Flow) (*FlowOption) {
    optionsStr := flow.Options
	options := &FlowOption{}
	err := json.Unmarshal([]byte(optionsStr), options)
	if err != nil {
		beego.Error("Cannot unmarshal options from flow", flow.Name, flow.Id)
		options = &FlowOption {
			Retry: map[string]*RetryOption {},
		}
	}
	return options
}
*/

func (exec *FlowExecutor) loadNodeStates(flow *models.Flow) (map[int]*models.NodeState, error) {
	beego.Debug("Load nodes & states for flow", flow.Id)
	states, err := flowService.GetNodeStatusByFlowId(flow.Id)

	//进项筛选，当处于失败状态的节点或者，或者初始化的节点才会执行
	var tempnodeList []*models.NodeState
	for _, state := range states {
		if state.Status == models.STATUS_FAILED || state.Status == models.STATUS_INIT{
			tempnodeList = append(tempnodeList, state)
		}

	}
	states = tempnodeList

	if err != nil {
		return nil, err
	}

	if len(states) == 0 {
		beego.Error("No node states found for flow:", flow.Id)
		return nil, errors.New("No node states found for flow:" + strconv.Itoa(flow.Id))
	}

	nodeIDs := make([]int, len(states))
	for i, state := range states {
		nodeIDs[i] = state.Node.Id
		beego.Debug("id =", state.Node.Id)
	}

	var nodes []*models.Node
	err = flowService.GetByIds(&models.Node{}, &nodes, nodeIDs)
	if len(nodes) != len(states) {
		beego.Error("Flow", flow.Id, " len(nodes) =", len(nodes),
			"len(states) =", len(states))
		return nil, errors.New("node states")
	}

	beego.Debug("There are", len(nodes), "nodes for flow", flow.Name)

	idNodeMap := make(map[int]*models.Node)
	for _, node := range nodes {
		idNodeMap[node.Id] = node
	}

	idStateMap := make(map[int]*models.NodeState)
	for _, state := range states {
		idStateMap[state.Node.Id] = state
		state.Node = idNodeMap[state.Node.Id]
		beego.Debug("Set state.Node =", state.Node.Ip)
	}

	return idStateMap, nil
}

func (exec *FlowExecutor) getBatchNodeStatesByStates(batch *models.FlowBatch,
	stateMap map[int]*models.NodeState) ([]*models.NodeState, error) {
	states := make([]*models.NodeState, len(stateMap))
	var i = 0
	for k, v := range stateMap {
		states[i] = v
		if states[i] == nil {
			msg := fmt.Sprintf("State not found for id=%d in state map", k)
			beego.Error(msg)
			return nil, errors.New(msg)
		}
		i++;
	}

	return states, nil
}

func (exec *FlowExecutor) getBatchNodeStates(batch *models.FlowBatch,
	stateMap map[int]*models.NodeState) ([]*models.NodeState, error) {

	beego.Debug("Load nodes for batch, id =", batch.Id, "nodes =", batch.Nodes)
	var nodeIDs []int
	if err := json.Unmarshal([]byte(batch.Nodes), &nodeIDs); err != nil {
		beego.Error("Fail to load node ids from batch", batch.Id, ", nodes=", batch.Nodes)
		return nil, err
	}

	states := make([]*models.NodeState, len(nodeIDs))
	for i, id := range nodeIDs {
		states[i] = stateMap[id]
		if states[i] == nil {
			msg := fmt.Sprintf("State not found for id=%d in state map", id)
			beego.Error(msg)
			return nil, errors.New(msg)
		}
	}

	return states, nil
}

/*
 * Create a flow instance in DB, set its state to INIT
 */
func (exec *FlowExecutor) createFlowInstance(name string, flow *models.FlowImpl, pool *models.Pool,
	stepOps []*models.StepOption, option *ExecOption,opUser string) (*models.Flow, error) {

	bytes, err := json.Marshal(stepOps)
	if err != nil {
		beego.Error("Fail to json dump :", stepOps, err)
		return nil, err
	}
	optionValues := string(bytes)

	fi := &models.Flow{
		Name: name,
		//Params:      params,
		Options:     optionValues,
		Status:      models.STATUS_INIT,
		Impl:        flow,
		Pool:        pool,
		StepLen:     option.MaxNum,
		OpUser:      opUser,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	err = flowService.InsertBase(fi)
	return fi, err
}

// Create node states for every node in given task
func (exec *FlowExecutor) createNodeStates(flow *models.Flow, nodes []*models.Node) ([]*models.NodeState, error) {
	states := make([]*models.NodeState, len(nodes))
	for i, node := range nodes {
		state := &models.NodeState{
			Ip:     node.Ip,
			VmId:   node.VmId,
			Node:   node,
			Flow:   flow,
			Pool:   node.Pool,
			Status: models.STATUS_INIT,
			Log:    "",
			Steps:  "[]",
		}

		// TODO batch insert all states
		beego.Debug("Create node state for node:", node.Id, ", ip=", node.Ip)
		err := flowService.InsertBase(state)
		if err != nil {
			beego.Error("Fail to create state, error: ", err)
		}

		states[i] = state
	}

	return states, nil
}

// Create batches for the given task.
func (exec *FlowExecutor) createBatches(instance *models.Flow, states []*models.NodeState, max int) error {
	total := len(states)
	if max < 1 {
		max = 1
	}

	batchNum := total / max
	if total%max != 0 {
		batchNum++
	}

	for i := 0; i < batchNum; i++ {
		start, end := i*max, i*max+max
		if end > total {
			end = total
		}

		part := states[start:end]
		ids := make([]int, len(part))
		for i, ns := range part {
			ids[i] = ns.Node.Id
		}

		idsBytes, _ := json.Marshal(ids)
		idsStr := string(idsBytes)

		batch := &models.FlowBatch{
			Flow:        instance,
			Status:      models.STATUS_INIT,
			Step:        -1, // not started
			Nodes:       idsStr,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
		}

		err := flowService.InsertBase(batch)
		if err != nil {
			return err
		}

		// update node states
		corrId := fmt.Sprintf("%d-%d", instance.Id, batch.Id)
		for _, ns := range part {
			ns.CorrId = corrId
			ns.Batch = batch

			flowService.UpdateBase(ns)
		}
	}

	return nil
}

func (exec *FlowExecutor) updateStepStatus(states []*models.NodeState, step string, stateCode int) error {
	beego.Debug("Update node states for step", step, ", num of states is", len(states))

	for _, state := range states {
		state.Status = stateCode
		if step != "" {
			state.Steps = step
		}
		state.UpdatedTime = time.Now()

		beego.Debug("Set node state", state.Ip, " step =", state.Steps, "status =", stateCode)
		err := flowService.UpdateBase(state)
		if err != nil {
			return err
		}
	}

	return nil
}

func (exec *FlowExecutor) allFailed(batch *models.FlowBatch, step *models.ActionImpl,
	states []*models.NodeState) error {

	exec.updateStepStatus(states, step.Name, models.STATUS_FAILED)

	batch.Status = models.STATUS_FAILED
	batch.UpdatedTime = time.Now()
	flowService.UpdateBase(batch)
	return nil
}

func (exec *FlowExecutor) allSuccess(batch *models.FlowBatch, states []*models.NodeState) error {
	exec.updateStepStatus(states, "", models.STATUS_SUCCESS)

	batch.Status = models.STATUS_SUCCESS
	batch.UpdatedTime = time.Now()
	flowService.UpdateBase(batch)

	return nil
}

func (exec *FlowExecutor) mergeParams(options []*models.StepOption,
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

func (exec *FlowExecutor) getCorrelationId(fid int, batchId int) string{
	return utils.GetCorrelationId(fid,batchId)
}
