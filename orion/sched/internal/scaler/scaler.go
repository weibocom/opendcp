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

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/davecgh/go-spew/spew"
	gosync "github.com/lrita/gosync"
)

var (
	gmutex = gosync.NewMutexGroup()
	pmutex = gosync.NewMutexGroup()
)

func init() {
	gosync.PanicOnBug = false
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
	count,_, _, _, _, _, err := onlineNodesList(cfg.Pool.Id)
	if err != nil {
		beego.Error(fmt.Sprintf("get online pool(%d) nodes failed: %v", cfg.Pool.Id, err))
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
	dependc := make(chan *dependNotice, should)
	expand := false

	if should-count>= 0 {
		expand = true
	}

	ctx = context.WithValue(ctx, "expand", expand)
	ctx = context.WithValue(ctx, "isdep", false)

	scaleDependPool(ctx, cfg, expand, should, dependc)

	wg.Add(1)
	go scalePool(ctx, cfg.Pool, expand, false, should, dependc)
}

func scaleDependPool(ctx context.Context, cfg *models.ExecTask, expand bool, should int, dependc chan *dependNotice) {
	var (
		wg    = ctx.Value("wg").(*sync.WaitGroup)
		ctrls = make(map[int]*dependCtrl)
	)

	for _, dep := range cfg.DependItems {
		if dep.Ignore {
			continue
		}
		// TODO ElasticCount
		num := int(math.Ceil(float64(should) * dep.Ratio))
		depc := make(chan *dependNotice, num)
		ctrls[dep.Pool.Id] = &dependCtrl{
			base:    should,
			num:     num,
			elastic: dep.ElasticCount,
			ratio:   dep.Ratio,
			dependc: depc,
		}
		wg.Add(1)
		ctx = context.WithValue(ctx, "isdep", true)
		go scalePool(ctx, dep.Pool, expand, true, num, depc)
	}

	wg.Add(1)
	go dependNoticeGoroutine(ctx, expand, should, dependc, ctrls)
}

// why so long ?
func scalePool(ctx context.Context, pool *models.Pool, expand, isdep bool, should int, dependc chan *dependNotice) {
	var (
		operation                            string
		picked 				     []*models.NodeState
		actions                              []*models.ActionImpl
		steps                                []*models.StepOption
		flow                                 *models.FlowImpl
		ff                                   *models.Flow
		wg                                   = ctx.Value("wg").(*sync.WaitGroup)
	)

	pmutex.Lock(pool.Id)
	defer wg.Done()
	defer pmutex.UnLockAndFree(pool.Id)

	onlineCount, init, ok, failed, running, stopped, err := onlineNodesList(pool.Id)
	if err != nil {
		beego.Error(fmt.Sprintf("get online pool(%d) nodes failed: %v", pool.Id, err))
		return
	}

	defer func() {
		if err != nil {
			beego.Error(fmt.Sprintf("scalePool pool(%d) exit on error: %v", err, pool.Id))
			if expand == isdep {
				dependc <- &dependNotice{pid: pool.Id, num: len(ok)}
			}
		}
	}()

	beego.Info(fmt.Sprintf("pool(%d) should(%d) online(%d) ok(%d) failed(%d) stopped(%d) running(%d)",
		pool.Id, should, onlineCount, len(ok), len(failed), len(stopped), len(running)))

	num := should - onlineCount
	if num >= 0 isdep{
		noticeDepentSuccess(ok,dependc)
		noticeDepentSuccess(running,dependc)
	}
	// pick up which nodes should to expand/shrink
	switch {
	case num == 0 && expand:
		for _, nn := range [][]*models.NodeState{init, failed, stopped} {
			picked = append(picked, nn...)
		}
		if len(picked) == 0 {
			beego.Info(fmt.Sprintf("pool(%d) has enough nodes online", pool.Id))
			return
		}
	case num > 0 && expand:
		for _, nn := range [][]*models.NodeState{init, failed, stopped} {
			picked = append(picked, nn...)
		}
		nn := make([]*models.NodeState, num)
		for i := range nn {
			nn[i] = new(models.NodeState)
		}
		picked = append(picked, nn...)
	case num < 0 && !expand:
		num = -num
		for _, nn := range [][]*models.NodeState{init, failed, stopped, running, ok} {
			if n := num - len(picked); n > 0 {
				picked = append(picked, nn[:min(n, len(nn))]...)
			}
		}
		if isdep && should > len(picked) {
			for i := 0; i < should - len(picked); i++ {
				dependc <- &dependNotice{num: 1}
			}
		}
	default:
		beego.Error(fmt.Sprintf("pool(%d) has a bug!!! should(%d) online(%d) expand(%v)",
			pool.Id, should, onlineCount, expand))
		return
	}

	tname := models.TaskExpend
	if !expand {
		tname = models.TaskShrink
	}

	flow, steps, err = flowAndSteps(pool, tname)
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

	ff, err = executor.Executor.CreateFlowInstance(pool.Name+"_"+tname, flow, pool,
		models.Crontab, steps, &executor.ExecOption{MaxNum: num}, models.Crontab)
	if err != nil {
		beego.Error(fmt.Sprintf("create the fucking real flow failed: %v", err))
		return
	}

	lg := &logger{fid: ff.Id}
	ctx = context.WithValue(ctx, "lg", lg)

	defer func() {
		if err != nil {
			if _, e := orm.NewOrm().Delete(ff); e != nil {
				lg.Errorf("pool(%d) delete real flow failed %v", pool.Id, e)
			}
		}
	}()

	nodeStates, resultErr := createNodeState(pool, ff, picked, operation)
	if resultErr != nil {
		err = resultErr
		lg.Errorf(err)
		return
	}

	if actions, err = actionImpls(steps); err != nil {
		lg.Errorf(err)
		return
	}

	executor.Executor.SetFlowStatus(ff, models.STATUS_RUNNING)
	// cancel point
	select {
	case <-ctx.Done():
		lg.Infof("pool(%d) task canceled", pool.Id)
		executor.Executor.SetFlowStatus(ff, models.STATUS_STOPPED)
		return
	default:
	}

	oknum, failednum, stopped:= doEachStep(ctx, ff, nodeStates, steps, actions, expand, isdep, dependc)

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

func noticeDepentSuccess(nodes []*models.NodeState, dependc chan *dependNotice){
	for i := 0; i < len(nodes); i++{
		dependc <- &dependNotice{num: 1}
	}
}
func createNodeState(pool *models.Pool, ff *models.Flow, nodes []*models.NodeState, operation string) (ns []*models.NodeState, err error) {
	o := orm.NewOrm()

	ns = make([]*models.NodeState, 0)

	if err = o.Begin(); err != nil {
		err = fmt.Errorf("tx begin failed: %v", err)
		return
	}

	defer func() {
		if err != nil {
			if e := o.Rollback(); e != nil {
				err = fmt.Errorf("%v, rollback failed %v", err, e)
			}
		}
	}()

	//delete old nodes
	for _, n := range nodes {
		//if nodeStatus is not running
		if n.Id != 0 && n.Status != models.STATUS_RUNNING{
			service.Flow.DeleteNodeById(n)
			ns = append(ns, n)
		}else if n.Id == 0{
			ns = append(ns, n)
		}
	}
	//insert nodes
	for _, n := range ns {
		state := &models.NodeState{
			Ip:          n.Ip,
			VmId:        n.VmId,
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
			return
		}
		ns = append(ns, state)
	}
	err = o.Commit()
	return
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

func doEachStep(ctx context.Context, flow *models.Flow, nodes []*models.NodeState, steps []*models.StepOption,
		actions []*models.ActionImpl, expand, isdep bool, dependc chan *dependNotice) (oknum, failednum, stoppednum int) {
	lg := ctx.Value("lg").(*logger)
	resultChannel := make(chan *models.NodeState, len(nodes))

	for _, ns := range nodes {
		toRunNodes := make([]*models.NodeState, 0)
		toRunNodes = append(toRunNodes, ns)
		go doNodeEachStep(ctx, flow, toRunNodes, actions, steps, resultChannel, dependc)
	}

	resultFlowStatus := models.STATUS_FAILED
	maxNodeStatesCostTime := 0.0
	for i := 0; i < len(nodes); i++ {
		select {
		case nodeStatesResult := <-resultChannel:
			if nodeStatesResult.Status == models.STATUS_SUCCESS {
				resultFlowStatus = models.STATUS_SUCCESS
				oknum++
			} else if nodeStatesResult.Status == models.STATUS_FAILED {
				failednum++
			} else if nodeStatesResult.Status == models.STATUS_STOPPED {
				stoppednum++
			}
			if nodeStatesResult.RunTime > maxNodeStatesCostTime {
				maxNodeStatesCostTime = nodeStatesResult.RunTime
			}
			if isdep && expand{

			}
		case <-time.After(time.Minute*15):
			failednum++
			lg.Errorf(flow.Id, "RunAndCheck node timeout!")
		}
	}
	close(resultChannel)
	//if all nodeNodestate failed then the flow is failed
	if failednum == len(nodes) && oknum == 0 {
		resultFlowStatus = models.STATUS_FAILED
	}
	if stoppednum == len(nodes) && oknum == 0{
		resultFlowStatus = models.STATUS_STOPPED
	}
	executor.Executor.SetFlowStatusWithSpenTime(flow, maxNodeStatesCostTime, resultFlowStatus)
}


func doNodeEachStep(ctx context.Context, flow *models.Flow, nodeState *models.NodeState,
	steps []*models.ActionImpl, stepOptions []*models.StepOption,
	resultChannel chan *models.NodeState, dependc chan *dependNotice) {

	var (
		stopped bool
		depidx  = -1
		expand  = ctx.Value("expand").(bool)
		isdep   = ctx.Value("isdep").(bool)
		lg      = ctx.Value("lg").(*logger)
	)

	if expand && !isdep {
		depidx = len(steps) - 1
	} else if !expand && isdep {
		depidx = 0
	}

	lg.Infof("depend action index is %d", depidx)

	//nodeRunnedTime := nodeState.RunTime
	startStepIndex := nodeState.StepNum

	var stepRunTimeArray []*models.StepRunTime
	err := json.Unmarshal([]byte(nodeState.StepRunTime), &stepRunTimeArray)
	if err != nil {
		lg.Errorf("Fail to load StepRunTime:", nodeState.StepRunTime, ", err:", err)
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
	executor.Executor.UpdateNodeStatus(steps[startStepIndex].Name, startStepIndex, stepRunTimeArray, nodeState, models.STATUS_INIT)


	success := true
	i := startStepIndex
	step := steps[startStepIndex]
	for i = startStepIndex; i < len(steps); i++ {
		step = steps[i]
		lg.Infof("run step %s(%d)", step.Name, i)
		flow, _ := service.Flow.GetFlowWithRel(flow.Id)
		if flow.Status == models.STATUS_STOPPED {
			stopped = true
		}
		// cancel point
		select {
		case <-ctx.Done():
			stopped = true
		default:
		}
		if stopped {
			lg.Infof("flow(%d) stopped at step %s(%d)", flow.Id, step.Name, i)
			nodeState.Status = models.STATUS_STOPPED
			success = false
			break
		}
		if i == depidx {
			if !isdep {
				should := waitDependNotice(ctx, dependc, lg)
				if should == 0{
					lg.Errorf("depend node is error!")
					nodeState.Status = models.STATUS_FAILED
					success = false
					break
				}
			}
		}
		// check flow status again
		flow, _ = service.Flow.GetFlowWithRel(flow.Id)
		if flow.Status == models.STATUS_STOPPED {
			lg.Infof("flow(%d) stopped at step %s(%d)", flow.Id, step.Name, i)
			nodeState.Status = models.STATUS_STOPPED
			success = false
			break
		}

		theHandler := handler.GetHandler(step.Type)
		if theHandler == nil {
			lg.Errorf(fmt.Sprintf("Handler not found for type %s", step.Type))
			nodeState.Status = models.STATUS_FAILED
			success = false
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

		needRunStepNodeState := make([]*models.NodeState, 0)
		needRunStepNodeState = append(needRunStepNodeState, nodeState)

		okNodes, _ := executor.Executor.RunStep(theHandler, step, i, needRunStepNodeState, stepParams, retryOption, stepRunTimeArray)

		if len(okNodes) == 0 {
			lg.Warnf(fmt.Sprintf("node %d run fail at step %s", nodeState.Id, step.Name))
			nodeState.Status = models.STATUS_FAILED
			success = false
			break

		} else {
			lg.Infof(fmt.Sprintf("node %d run success at step %s", nodeState.Id, step.Name))
			nodeState.Status = models.STATUS_RUNNING
		}
	}

	if success {
		nodeState.Status = models.STATUS_SUCCESS
		lg.Infof(fmt.Sprintf("node %d run success all steps", nodeState.Id))
	}
	err = executor.Executor.UpdateNodeStatus(step.Name, i, stepRunTimeArray, nodeState, nodeState.Status)
	if err != nil {
		lg.Errorf(fmt.Sprintf("update node state db error: %s", err.Error()))
	}
	//put nodeState to chan
	resultChannel <- nodeState // Send nodeState to channel

	if isdep && success{
		dependc <- &dependNotice{num: 1}
	}else if isdep && !success{
		dependc <- &dependNotice{num: 0}
	}

	return
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

