package scaler

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"sync"
	"time"

	"weibo.com/opendcp/orion/executor"
	"weibo.com/opendcp/orion/handler"
	"weibo.com/opendcp/orion/helper"
	"weibo.com/opendcp/orion/models"
	"weibo.com/opendcp/orion/service"
	"weibo.com/opendcp/orion/utils"

	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/davecgh/go-spew/spew"
	theGosync "github.com/lrita/gosync"
)

var (
	gmutex = theGosync.NewMutexGroup()
	pmutex = theGosync.NewMutexGroup()
)

func init() {
	theGosync.PanicOnBug = false
}

type configs interface {
	Config(id int) *models.ExecTask
}

func Scale(ctx context.Context, cfgs configs, id, idx int) {
	wg := ctx.Value("wg").(*sync.WaitGroup)
	gmutex.Lock(id)
	defer gmutex.UnLockAndFree(id)
	defer wg.Done()

	cfg := cfgs.Config(id)
	if cfg == nil {
		beego.Error(fmt.Sprintf("task(%d) has not config", id))
		return
	}

	beego.Info(spew.Sprintf("task(%+v) idx(%d) running", cfg, idx))

	if cfg.Type != models.TaskExpend && cfg.Type != models.TaskShrink {
		beego.Error(fmt.Sprintf("id(%d) not support task %q", id, cfg.Type))
		return
	}

	if len(cfg.CronItems)-1 < idx {
		beego.Error(fmt.Sprintf("id(%d) has the %d cronitem", cfg.Id, idx))
		return
	}

	should := cfg.CronItems[idx].InstanceNum

	expand, picked, err := initScalePool(should, cfg.Pool)

	if len(picked) == 0 {
		beego.Info(fmt.Sprintf("task(%d) pool(%d) onlinde nodes is okay!", cfg.Id, cfg.Pool.Id))
		return
	}
	if err != nil {
		beego.Error(err.Error())
		return
	}

	// cancel point
	select {
	case <-ctx.Done():
		beego.Info(fmt.Sprintf("task(%d) pool(%d) canceled", cfg.Id, cfg.Pool.Id))
		return
	default:
	}

	//run node channel
	dependc := make(chan *dependNotice, len(picked))

	scaleDependPool(ctx, cfg, should, len(picked), dependc)

	wg.Add(1)

	ctx = context.WithValue(ctx, "expand", expand)
	ctx = context.WithValue(ctx, "isdep", false)
	ctx = context.WithValue(ctx, "noticeStep", "")
	go scalePool(ctx, cfg.Pool, expand, picked, dependc)
}

func scaleDependPool(ctx context.Context, cfg *models.ExecTask, currentNum, should int, dependc chan *dependNotice) {
	var (
		wg    = ctx.Value("wg").(*sync.WaitGroup)
		ctrls = make(map[int]*dependCtrl)
	)

	for _, dep := range cfg.DependItems {
		if dep.Ignore {
			continue
		}
		// TODO ElasticCount
		num := int(math.Ceil(float64(currentNum) * dep.Ratio))

		dependExpand, dependNodes, err := initScalePool(num, dep.Pool)
		if err != nil {
			beego.Error(fmt.Sprintf("pool(%d) get expand and nodes err: %v", dep.Pool.Id, err.Error()))
			continue
		}
		if len(dependNodes) == 0 {
			beego.Error(fmt.Sprintf("pool(%d) get none nodes to run", dep.Pool.Id))
			continue
		}

		depc := make(chan *dependNotice, len(dependNodes))

		ctrls[dep.Pool.Id] = &dependCtrl{
			base:    currentNum,
			num:     len(dependNodes),
			elastic: dep.ElasticCount,
			ratio:   dep.Ratio,
			dependc: depc,
		}
		wg.Add(1)
		ctx = context.WithValue(ctx, "expand", dependExpand)
		ctx = context.WithValue(ctx, "isdep", true)
		ctx = context.WithValue(ctx, "noticeStep", dep.StepName)
		go scalePool(ctx, dep.Pool, dependExpand, dependNodes, depc)
	}

	wg.Add(1)
	go dependNoticeGoroutine(ctx, should, dependc, ctrls)
}

func initScalePool(should int, pool *models.Pool) (bool, []*models.NodeState, error) {
	var (
		expand = false
		picked = make([]*models.NodeState, 0)
	)
	onlineCount, init, ok, failed, running, stopped, err := onlineNodesList(pool.Id)
	if err != nil {
		return expand, picked, fmt.Errorf(fmt.Sprintf("get shit online pool(%d) nodes failed: %v", pool.Id, err))
	}
	beego.Info(fmt.Sprintf("pool(%d) should(%d) online(%d) ok(%d) failed(%d) stopped(%d) running(%d)",
		pool.Id, should, onlineCount, len(ok), len(failed), len(stopped), len(running)))

	num := should - onlineCount

	if num >= 0 {
		expand = true
	}

	// pick up which nodes should to expand/shrink
	switch {
	case num == 0:
		for _, nn := range [][]*models.NodeState{init, failed, stopped} {
			picked = append(picked, nn...)
		}
	case num > 0:
		for _, nn := range [][]*models.NodeState{init, failed, stopped} {
			picked = append(picked, nn...)
		}
		nn := make([]*models.NodeState, num)
		for i := range nn {
			nn[i] = new(models.NodeState)
		}
		picked = append(picked, nn...)
	case num < 0:
		num = -num
		for _, nn := range [][]*models.NodeState{init, failed, stopped, running, ok} {
			if n := num - len(picked); n > 0 {
				picked = append(picked, nn[:min(n, len(nn))]...)
			}
		}
	default:
		return expand, picked, fmt.Errorf(fmt.Sprintf("pool(%d) has a fuck bug!!! should(%d) online(%d) expand(%v)",
			pool.Id, should, onlineCount, expand))
	}

	return expand, picked, nil
}

// why so long ?
func scalePool(ctx context.Context, pool *models.Pool, expand bool, picked []*models.NodeState, dependc chan *dependNotice) {
	var (
		operation  string
		nodeStates []*models.NodeState
		actions    []*models.ActionImpl
		steps      []*models.StepOption
		flow       *models.FlowImpl
		ff         *models.Flow
		wg         = ctx.Value("wg").(*sync.WaitGroup)
	)

	pmutex.Lock(pool.Id)
	defer wg.Done()
	defer pmutex.UnLockAndFree(pool.Id)

	tname := models.TaskExpend
	if !expand {
		tname = models.TaskShrink
	}

	flow, steps, err := flowAndSteps(pool, tname)

	if err != nil {
		beego.Error(fmt.Sprintf("pool(%d) get flow and steps failed %v", pool.Id, err))
		return
	}

	if len(steps) == 0 {
		beego.Error(fmt.Sprintf("pool(%d) has no steps", pool.Id))
		return
	}

	if expand {
		if steps[0].Name != "create_vm" {
			beego.Error(fmt.Sprintf("pool(%d) first step of expand template is not create_vm",
				pool.Id))
			return
		}
		executor.Executor.MergeParams(steps, map[string]interface{}{
			helper.CREATE_VM: map[string]interface{}{helper.KEY_VM_TYPE: pool.VmType},
			helper.REGISTER:  map[string]interface{}{helper.KEY_SD_ID: pool.SdId},
		})
		operation = models.TaskExpend
	} else {
		if steps[len(steps)-1].Name != "return_vm" {
			beego.Error(fmt.Sprintf("pool(%d) last step of expand template is not return_vm",
				pool.Id))
			return
		}
		executor.Executor.MergeParams(steps, map[string]interface{}{
			helper.RETURN_VM:  map[string]interface{}{helper.KEY_VM_TYPE: pool.VmType},
			helper.UNREGISTER: map[string]interface{}{helper.KEY_SD_ID: pool.SdId},
		})
		operation = models.TaskShrink
	}

	if ff, err = executor.Executor.CreateFlowInstance(pool.Name+"_"+tname, flow, pool,
		models.Crontab, steps, &executor.ExecOption{MaxNum: len(picked)}, models.Crontab); err != nil {
		beego.Error(fmt.Sprintf("create the fucking real flow failed: %v", err))
		return
	}

	lg := &logger{fid: ff.Id}
	ctx = context.WithValue(ctx, "lg", lg)

	defer func() {
		if err != nil {
			if e := executor.Executor.SetFlowStatus(ff, models.STATUS_FAILED); e != nil {
				lg.Errorf("pool(%d) set real flow status failed err: %v", pool.Id, e)
			}
			lg.Errorf(err.Error())
		}
	}()

	if nodeStates, err = createNodeState(pool, ff, picked, operation); err != nil {
		return
	}

	if actions, err = actionImpls(steps); err != nil {
		return
	}

	if err = executor.Executor.SetFlowStatus(ff, models.STATUS_RUNNING); err != nil {
		return
	}
	// cancel point
	select {
	case <-ctx.Done():
		lg.Infof("pool(%d) task canceled", pool.Id)
		executor.Executor.SetFlowStatus(ff, models.STATUS_STOPPED)
		return
	default:
	}

	oknum, failednum, _ := doEachStep(ctx, ff, nodeStates, steps, actions, dependc)

	lg.Infof("pool(%d) flow(%d) finished ok(%d) failed(%d) status(%d)",
		pool.Id, ff.Id, oknum, failednum, ff.Status)
}

func flowAndSteps(pool *models.Pool, taskname string) (*models.FlowImpl, []*models.StepOption, error) {
	task, err := utils.Json.ToMap(pool.Tasks)
	if err != nil {
		return nil, nil, fmt.Errorf("pool(%d), invalid pool task %q",
			pool.Id, pool.Tasks)
	}

	tidi := task[taskname]
	if tidi == nil {
		return nil, nil, fmt.Errorf("pool(%d) task has no taskname(%s)",
			pool.Id, pool.Tasks)
	}
	tid, err := utils.ToInt(tidi)
	if err != nil {
		return nil, nil, fmt.Errorf("pool(%d) pool task has no taskname(%s) to id",
			pool.Id, pool.Tasks)
	}

	flow := &models.FlowImpl{Id: tid}
	if err := service.Flow.GetBase(flow); err != nil {
		return nil, nil, fmt.Errorf("pool(%d) tid(%d) pool not found template",
			pool.Id, tid)
	}

	var steps []*models.StepOption
	if err := json.Unmarshal([]byte(flow.Steps), &steps); err != nil {
		return nil, nil, fmt.Errorf("pool(%d) tid(%d) step decode failed",
			pool.Id, tid)
	}
	return flow, steps, nil
}

func createNodeState(pool *models.Pool, ff *models.Flow, nodes []*models.NodeState, operation string) (ns []*models.NodeState, err error) {
	o := orm.NewOrm()

	ns = make([]*models.NodeState, 0)

	if err = o.Begin(); err != nil {
		err = fmt.Errorf("tx begin failed: %v", err)
		return ns, err
	}

	defer func() {
		if err != nil {
			if e := o.Rollback(); e != nil {
				err = fmt.Errorf("%v, rollback failed %v", err, e)
			}
		}
	}()

	//delete old nodes
	deleteNodes := make([]*models.NodeState, 0)
	for _, n := range nodes {
		//if nodeStatus is not running
		if n.Id != 0 && n.Status != models.STATUS_RUNNING {
			n.Deleted = true
			n.UpdatedTime = time.Now()
			if _, err := o.Update(n, "deleted", "updated_time"); err != nil {
				err = fmt.Errorf("delete node %d failed %v", n.Id, err)
				beego.Error(err)
			} else {
				deleteNodes = append(deleteNodes, n)
			}
		} else if n.Id == 0 {
			deleteNodes = append(deleteNodes, n)
		}
	}
	//insert nodes
	for _, n := range deleteNodes {
		state := &models.NodeState{
			Ip:          "-",
			VmId:        "",
			Flow:        ff,
			Pool:        pool,
			Status:      models.STATUS_INIT,
			Log:         "",
			Steps:       "[]",
			LastOp:      n.LastOp,
			StepRunTime: "[]",
			StepNum:     0,
			RunTime:     0.0,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
			Deleted:     false,
			NodeType:    models.Crontab,
		}
		if n.Id != 0 {
			state.Ip = n.Ip
			state.VmId = n.VmId
		}
		if n.LastOp != operation {
			// Reset step num if current operation is different from last operation.
			// If current operation is equal the last operation, we run this task as
			// a repeat task, it will retry those nodes which are failed in last task.
			// If current operation is not equal the last operation, we set the step
			// num to 0, and run this task from the beginning.
			state.StepNum = 0
			state.LastOp = operation
		}

		if _, err := o.Insert(state); err != nil {
			err = fmt.Errorf("node insert %d failed %v", n.Id, err)
			return ns, err
		}
		ns = append(ns, state)
	}
	err = o.Commit()
	return ns, err
}

func min(a, b int) int {

	if a > b {
		return b
	}

	return a
}

func actionImpls(steps []*models.StepOption) ([]*models.ActionImpl, error) {
	var actions []*models.ActionImpl
	for _, s := range steps {
		action := handler.GetActionImpl(s.Name)
		if action == nil {
			return nil, fmt.Errorf("get action %s failed", s.Name)
		}
		actions = append(actions, action)
	}
	return actions, nil
}

func runNodeToChannel(ctx context.Context, flow *models.Flow, actions []*models.ActionImpl,
	steps []*models.StepOption, runNodes []*models.NodeState,
	resultChannel chan *models.NodeState, dependc chan *dependNotice) error {

	var (
		runSleep = time.Microsecond // sleep 1us to run next
	)
	//put nodes to channel
	runNodeChannel := make(chan *models.NodeState, len(runNodes))
	defer close(runNodeChannel)

	for _, n := range runNodes {
		runNodeChannel <- n
		beego.Info("runNode channel add runNnode")
	}
	//run nodes from channel
	for i := 0; i < len(runNodes); i++ {
		ns, ok := <-runNodeChannel
		if !ok {
			return errors.New("runNodeChannel was closed!")
		}
		go doNodeEachStep(ctx, flow, ns, actions, steps, resultChannel, dependc)
		time.Sleep(runSleep) // sleep 1us to run next
	}
	return nil
}
func doEachStep(ctx context.Context, flow *models.Flow, nodes []*models.NodeState, steps []*models.StepOption,
	actions []*models.ActionImpl, dependc chan *dependNotice) (oknum, failednum, stoppednum int) {

	var (
		lg                    = ctx.Value("lg").(*logger)
		timeout               = time.After(15 * time.Minute)
		resultFlowStatus      = models.STATUS_FAILED
		resultChannel         = make(chan *models.NodeState, len(nodes))
		maxNodeStatesCostTime = 0.0
	)

	defer close(resultChannel)

	if err := runNodeToChannel(ctx, flow, actions, steps, nodes, resultChannel, dependc); err != nil {
		lg.Errorf("run node to channel err: ", err.Error())
		return oknum, failednum, stoppednum
	}

	for i := 0; i < len(nodes); i++ {
		select {
		case nodeStatesResult := <-resultChannel:
			if nodeStatesResult.Status == models.STATUS_SUCCESS {
				oknum++
			} else if nodeStatesResult.Status == models.STATUS_FAILED {
				failednum++
			} else if nodeStatesResult.Status == models.STATUS_STOPPED {
				stoppednum++
			}
			if nodeStatesResult.RunTime > maxNodeStatesCostTime {
				maxNodeStatesCostTime = nodeStatesResult.RunTime
			}
		case <-timeout:
			failednum++
			lg.Errorf("RunAndCheck run node timeout!")
		}
	}
	//if all nodeNodestate failed then the flow is failed
	if oknum != 0 {
		resultFlowStatus = models.STATUS_SUCCESS
	} else if failednum == len(nodes) {
		resultFlowStatus = models.STATUS_FAILED
	} else if stoppednum == len(nodes) {
		resultFlowStatus = models.STATUS_STOPPED
	}
	executor.Executor.SetFlowStatusWithSpenTime(flow, maxNodeStatesCostTime, resultFlowStatus)

	return oknum, failednum, stoppednum
}

func doNodeEachStep(ctx context.Context, flow *models.Flow, nodeState *models.NodeState,
	steps []*models.ActionImpl, stepOptions []*models.StepOption,
	resultChannel chan *models.NodeState, dependc chan *dependNotice) {

	var (
		stopped    = false
		depidx     = -1
		expand     = ctx.Value("expand").(bool)
		isdep      = ctx.Value("isdep").(bool)
		noticeStep = ctx.Value("noticeStep")
		lg         = ctx.Value("lg").(*logger)
		runSuccess = true
		hasNotice  = false
	)

	defer func() {
		resultChannel <- nodeState // Send nodeState to channel
		//send notice
		if isdep && !hasNotice {
			dependc <- &dependNotice{num: 0}
			hasNotice = true
		}
	}()

	if expand && !isdep {
		depidx = len(steps) - 1
	} else if !expand && !isdep {
		depidx = 0
	} else if !expand && isdep {
		noticeStep = helper.RETURN_VM
		depidx = -1
	} else if expand && isdep {
		depidx = -1
	}

	lg.Infof("depend action index is %d", depidx)

	stepRunTimeArray, startStepIndex, err := generateRunTimeStep(nodeState, steps)
	if err != nil {
		lg.Errorf("Fail to load StepRunTime:", nodeState.StepRunTime, ", err:", err)
	}

	if isdep {
		for initStart := 0; initStart < startStepIndex; initStart++ {
			if steps[initStart].Name == noticeStep && !hasNotice {
				dependc <- &dependNotice{num: 1}
				hasNotice = true
				return
			}
		}
	}
	//update nodesState to running
	executor.Executor.UpdateNodeStatus(steps[startStepIndex].Name, startStepIndex, stepRunTimeArray, nodeState, models.STATUS_INIT)
	i := startStepIndex
	step := steps[0]
	for ; i < len(steps); i++ {
		step = steps[i]
		lg.Infof("run step %s(%d)", step.Name, i)

		if stopped = checFlowAndNodeStop(ctx, flow, nodeState); stopped {
			lg.Infof("flow(%d) stopped at step %s(%d)", flow.Id, step.Name, i)
			if nodeState.Status != models.STATUS_STOPPED && nodeState.Status != models.STATUS_SUCCESS{
				nodeState.Status = models.STATUS_STOPPED
			}
			runSuccess = false
			break
		}
		if i == depidx {
			if !isdep {
				should := waitDependNotice(ctx, dependc, lg)
				if should == 0 {
					lg.Errorf("depend node is error!")
					nodeState.Status = models.STATUS_FAILED
					runSuccess = false
					break
				}
				// check flow status again
				if stopped = checFlowAndNodeStop(ctx, flow, nodeState); stopped {
					lg.Infof("flow(%d) stopped at step %s(%d)", flow.Id, step.Name, i)
					if nodeState.Status != models.STATUS_STOPPED && nodeState.Status != models.STATUS_SUCCESS{
						nodeState.Status = models.STATUS_STOPPED
					}
					runSuccess = false
					break
				}
			}
		}
		theHandler := handler.GetHandler(step.Type)
		if theHandler == nil {
			lg.Errorf(fmt.Sprintf("Handler not found for type %s", step.Type))
			nodeState.Status = models.STATUS_FAILED
			runSuccess = false
			break
		}
		// get param values
		stepOption := stepOptions[i]
		stepParams := stepOption.Values
		// get retry option
		retryOption := stepOption.Retry
		if retryOption == nil {
			retryOption = &models.RetryOption{
				RetryTimes:  0,
				IgnoreError: false,
			}
		}

		if nodeState.Ip != "-" && step.Name == "create_vm" {
			lg.Infof(fmt.Sprintf("node %d already create_vm no neeed to run: %s skip", nodeState.Id, step.Name))
			nodeState.Status = models.STATUS_RUNNING
			//send notice
			if isdep && step.Name == noticeStep && runSuccess && !hasNotice {
				dependc <- &dependNotice{num: 1}
				hasNotice = true
			}
			continue
		}
		needRunStepNodeState := make([]*models.NodeState, 0)
		needRunStepNodeState = append(needRunStepNodeState, nodeState)

		okNodes, _ := executor.Executor.RunStep(theHandler, step, i, needRunStepNodeState, stepParams, retryOption, stepRunTimeArray)

		if len(okNodes) == 0 {
			lg.Warnf(fmt.Sprintf("node %d run fail at step %s", nodeState.Id, step.Name))
			nodeState.Status = models.STATUS_FAILED
			runSuccess = false
			break

		} else {
			if okNodes[0].Status != models.STATUS_RUNNING{
				nodeState.Status = okNodes[0].Status
				runSuccess = false
				lg.Infof(fmt.Sprintf("node %d status %d run stop at step %s", nodeState.Id, nodeState.Status, step.Name))
				break
			}else {
				lg.Infof(fmt.Sprintf("node %d run success at step %s", nodeState.Id, step.Name))
				nodeState.Status = models.STATUS_RUNNING
			}
		}
		//send notice
		if isdep && step.Name == noticeStep && runSuccess && !hasNotice {
			dependc <- &dependNotice{num: 1}
			hasNotice = true
		}
	}

	if runSuccess {
		nodeState.Status = models.STATUS_SUCCESS
		lg.Infof(fmt.Sprintf("node %d run success all steps", nodeState.Id))
	}
	if err := executor.Executor.UpdateNodeStatus(step.Name, i, stepRunTimeArray, nodeState, nodeState.Status); err != nil {
		lg.Errorf(fmt.Sprintf("update node state db error: %s", err.Error()))
	}
}

func waitDependNotice(ctx context.Context, dependc chan *dependNotice, lg *logger) int {
	select {
	case n, ok := <-dependc:
		if !ok {
			return int(math.MaxInt32)
		}
		return n.num
	case <-time.After(30 * time.Minute):
		lg.Errorf("wait depend signal timeout")
	case <-ctx.Done():
		lg.Warnf("wait depend canceled")
	}

	return 0
}

//generate nodes runTime steps
func generateRunTimeStep(nodeState *models.NodeState, steps []*models.ActionImpl) ([]*models.StepRunTime, int, error) {
	var (
		stepRunTimeArray = make([]*models.StepRunTime, 0)
		startIndex       = 0
	)

	if err := json.Unmarshal([]byte(nodeState.StepRunTime), &stepRunTimeArray); err != nil {
		return stepRunTimeArray, startIndex, err
	}
	//generate nodes runTime steps
	for _, step := range steps {
		stepRunTimeElement := &models.StepRunTime{
			Name:    step.Name,
			RunTime: 0.0,
		}
		stepRunTimeArray = append(stepRunTimeArray, stepRunTimeElement)

	}

	if steps[0].Name == helper.CREATE_VM && nodeState.Ip != "-" {
		startIndex = 1
	}

	return stepRunTimeArray, startIndex, nil
}

func checFlowAndNodeStop(ctx context.Context, flow *models.Flow, node *models.NodeState) bool {
	var stopped = false
	freshFlow, _ := service.Flow.GetFlowWithRel(flow.Id)
	if freshFlow.Status == models.STATUS_STOPPED || freshFlow.Status  == models.STATUS_SUCCESS{
		node.Status = freshFlow.Status
		stopped = true
	} else {
		freshNode, _ := service.Flow.GetNodeById(node.Id)
		if freshNode.Status == models.STATUS_STOPPED || freshNode.Status == models.STATUS_SUCCESS{
			stopped = true
			node.Status = freshNode.Status
		}
	}
	// cancel point
	select {
	case <-ctx.Done():
		stopped = true
	default:
	}

	return stopped
}
