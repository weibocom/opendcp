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
		beego.Error("task(%d) has not config", id)
		return
	}

	beego.Info(spew.Sprintf("task(%+v) idx(%d) running", cfg, idx))

	if cfg.Type != models.TaskExpend && cfg.Type != models.TaskShrink {
		beego.Error(fmt.Sprintf("id(%d) not support task %q", id, cfg.Type))
		return
	}

	if len(cfg.CronItems)-1 < idx {
		beego.Error("id(%d) has the %d cronitem", cfg.Id, idx)
		return
	}

	should := cfg.CronItems[idx].InstanceNum
	online, err := onlineNodesList(cfg.Pool.Id)
	if err != nil {
		beego.Error("get online pool(%d) nodes failed: %v", cfg.Pool.Id, err)
		return
	}

	// cancel point
	select {
	case <-ctx.Done():
		beego.Info("task(%d) pool(%d) canceled", cfg.Id, cfg.Pool.Id)
		return
	default:
	}

	dependc := make(chan *dependNotice, 1)
	expand := false

	if should-len(online) >= 0 {
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
		depc := make(chan *dependNotice, 1)
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
	go dependNoticeGoroutine(ctx, expand, dependc, ctrls)
}

// why so long ?
func scalePool(ctx context.Context, pool *models.Pool, expand, isdep bool, should int, dependc chan *dependNotice) {
	var (
		operation                            string
		ok, failed, running, stopped, picked []*Node
		actions                              []*models.ActionImpl
		steps                                []*models.StepOption
		flow                                 *models.FlowImpl
		ff                                   *models.Flow
		batch                                *models.FlowBatch
		wg                                   = ctx.Value("wg").(*sync.WaitGroup)
	)

	pmutex.Lock(pool.Id)
	defer wg.Done()
	defer pmutex.UnLockAndFree(pool.Id)

	online, err := onlineNodesList(pool.Id)
	if err != nil {
		beego.Error("get online pool(%d) nodes failed: %v", pool.Id, err)
		return
	}

	for _, n := range online {
		switch n.n.Status {
		case models.STATUS_INIT, models.STATUS_RUNNING:
			running = append(running, n)
		case models.STATUS_SUCCESS:
			ok = append(ok, n)
		case models.STATUS_FAILED:
			failed = append(failed, n)
		case models.STATUS_STOPPED:
			stopped = append(stopped, n)
		}
	}

	defer func() {
		if err != nil {
			beego.Error("scalePool pool(%d) exit on error:", err, pool.Id)
			if expand == isdep {
				dependc <- &dependNotice{pid: pool.Id, num: len(ok)}
			}
		}
	}()

	beego.Info(fmt.Sprintf("pool(%d) should(%d) online(%d) ok(%d) failed(%d) stopped(%d) running(%d)",
		pool.Id, should, len(online), len(ok), len(failed), len(stopped), len(running)))

	num := should - len(online)
	// pick up which nodes should to expand/shrink
	switch {
	case num == 0 && expand:
		for _, nn := range [][]*Node{failed, running, stopped} {
			picked = append(picked, nn...)
		}
		if len(picked) == 0 {
			beego.Info(fmt.Sprintf("pool(%d) has enough nodes online", pool.Id))
			if isdep {
				dependc <- &dependNotice{pid: pool.Id, num: should}
			}
			return
		}
	case num > 0 && expand:
		for _, nn := range [][]*Node{failed, running, stopped} {
			picked = append(picked, nn...)
		}
		nn := make([]*Node, num)
		for i := range nn {
			nn[i] = new(Node)
		}
		picked = append(picked, nn...)
	case num < 0 && !expand:
		num = -num
		for _, nn := range [][]*Node{failed, stopped, running, ok} {
			if n := num - len(picked); n > 0 {
				picked = append(picked, nn[:min(n, len(nn))]...)
			}
		}
	default:
		err = fmt.Errorf("pool(%d) has a bug!!! should(%d) online(%d) expand(%v)",
			pool.Id, should, len(online), expand)
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
		err = fmt.Errorf("pool(%d) has no steps", pool.Id)
		return
	}

	if expand {
		if steps[0].Name != "create_vm" {
			err = fmt.Errorf("pool(%d) first step of expand template is not create_vm",
				pool.Id)
			return
		}
		executor.Executor.MergeParams(steps, map[string]interface{}{
			helper.CREATE_VM: map[string]interface{}{helper.KEY_VM_TYPE: pool.VmType},
			helper.REGISTER:  map[string]interface{}{helper.KEY_SD_ID: pool.SdId},
		})
		operation = models.TaskExpend
	} else {
		if steps[len(steps)-1].Name != "return_vm" {
			err = fmt.Errorf("pool(%d) last step of expand template is not return_vm",
				pool.Id)
			return
		}
		executor.Executor.MergeParams(steps, map[string]interface{}{
			helper.RETURN_VM:  map[string]interface{}{helper.KEY_VM_TYPE: pool.VmType},
			helper.UNREGISTER: map[string]interface{}{helper.KEY_SD_ID: pool.SdId},
		})
		operation = models.TaskShrink
	}

	ff, err = createRealFlow(pool.Name+"_"+tname, len(picked), steps, pool, flow)
	if err != nil {
		err = fmt.Errorf("create the fucking real flow failed: %v", err)
		return
	}

	lg := &logger{fid: ff.Id}
	ctx = context.WithValue(ctx, "lg", lg)

	defer func() {
		if err != nil {
			if _, e := orm.NewOrm().Delete(ff); e != nil {
				lg.Infof("pool(%d) delete real flow failed %v", pool.Id, e)
			}
		}
	}()

	if err = createNodeState(pool, ff, picked, operation); err != nil {
		return
	}

	if batch, err = createFlowBatch(picked, ff); err != nil {
		err = fmt.Errorf("pool(%d) create batch failed %v", pool.Id, err)
		lg.Infof("%v", err)
		return
	}

	if actions, err = actionImpls(steps); err != nil {
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

	oknum, failednum, status := doEachStep(ctx, picked, batch, steps, actions, dependc)

	executor.Executor.SetFlowStatus(ff, status)

	if expand == isdep {
		dependc <- &dependNotice{pid: pool.Id, num: oknum}
	}

	lg.Infof("pool(%d) flow(%d) batch(%d) finished ok(%d) failed(%d) status(%d)",
		pool.Id, ff.Id, batch.Id, oknum, failednum, status)
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

func createRealFlow(name string, num int, steps []*models.StepOption,
	pool *models.Pool, flow *models.FlowImpl) (*models.Flow, error) {
	data, _ := json.Marshal(steps)
	ff := &models.Flow{
		Name:        name,
		Options:     string(data),
		Status:      models.STATUS_INIT,
		Impl:        flow,
		Pool:        pool,
		StepLen:     num,
		OpUser:      models.Crontab, // TODO just use this string as operation user ???
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	if _, err := orm.NewOrm().Insert(ff); err != nil {
		return nil, err
	}
	return ff, nil
}

func createNodeState(pool *models.Pool, ff *models.Flow, nodes []*Node, operation string) error {
	o := orm.NewOrm()
	for i := 0; i < len(nodes); i++ {
		if nodes[i].n.Ip == "" {
			nodes[i].n.Ip = "-"
		}
		nodes[i].n.Pool = pool
		nodes[i].n.NodeType = models.Crontab
		nodes[i].s.Ip = nodes[i].n.Ip
		nodes[i].s.VmId = nodes[i].n.VmId
		nodes[i].s.Node = &nodes[i].n
		nodes[i].s.Flow = ff
		nodes[i].s.Pool = pool
		nodes[i].s.Log = ""
		nodes[i].s.Steps = "[]"
		nodes[i].s.CreatedTime = time.Now()

		if nodes[i].s.LastOp != operation {
			// Reset step num if current operation is different from last operation.
			// If current operation is equal the last operation, we run this task as
			// a repeat task, it will retry those nodes which are failed in last task.
			// If current operation is not equal the last operation, we set the step
			// num to 0, and run this task from the beginning.
			nodes[i].s.StepNum = 0
			nodes[i].s.LastOp = operation
		}

		if nodes[i].n.Id == 0 {
			if _, err := o.Insert(&nodes[i].n); err != nil {
				return fmt.Errorf("node insert %d failed %v", i, err)
			}
		} else {
			if _, err := o.Update(&nodes[i].n); err != nil {
				return fmt.Errorf("node insert %d failed %v", i, err)
			}
		}

		nodes[i].s.Id = 0
		if _, err := o.Insert(&nodes[i].s); err != nil {
			return fmt.Errorf("nodestate insert %d failed %v", i, err)
		}
	}
	return nil
}

func createFlowBatch(nodes []*Node, ff *models.Flow) (*models.FlowBatch, error) {
	ids := make([]int, len(nodes))
	for i := 0; i < len(nodes); i++ {
		ids[i] = nodes[i].n.Id
	}
	data, _ := json.Marshal(ids)
	batch := &models.FlowBatch{
		Flow:        ff,
		Status:      models.STATUS_INIT,
		Step:        -1, // not started
		Nodes:       string(data),
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}
	o := orm.NewOrm()
	if _, err := o.Insert(batch); err != nil {
		return nil, err
	}

	for i := 0; i < len(nodes); i++ {
		nodes[i].s.CorrId = fmt.Sprintf("%d-%d", ff.Id, batch.Id)
		nodes[i].s.Batch = batch
		if _, err := o.Update(&nodes[i].s); err != nil {
			return nil, err
		}
	}
	return batch, nil
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

func doEachStep(ctx context.Context, nodes []*Node, batch *models.FlowBatch,
	steps []*models.StepOption, actions []*models.ActionImpl, dependc chan *dependNotice) (oknum, failednum, status int) {
	var (
		stopped bool
		depidx  = -1
		finidx  = len(actions) - 1
		expand  = ctx.Value("expand").(bool)
		isdep   = ctx.Value("isdep").(bool)
		lg      = ctx.Value("lg").(*logger)
	)

	if expand && !isdep {
		depidx = len(actions) - 1
	} else if !expand && isdep {
		depidx = 0
	}

	lg.Infof("batch(%d) depend action index is %d", batch.Id, depidx)

	for idx, action := range actions {
		var picked []*models.NodeState

		lg.Infof("batch(%d) run step %s(%d)", batch.Id, action.Name, idx)
		for i := 0; i < len(nodes); i++ {
			if nodes[i].s.StepNum == idx {
				picked = append(picked, &nodes[i].s)
			}
		}

		// check stopped
		flow, _ := service.Flow.GetFlowWithRel(batch.Flow.Id)
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
			lg.Infof("batch(%d) stopped at step %s(%d)", batch.Id, action.Name, idx)
			status = models.STATUS_STOPPED
			updateNode(picked, idx, status, action.Name)
			updateBatch(batch, idx, status)
			return
		}

		if idx == depidx {
			should := waitDependNotice(ctx, dependc, lg)
			lg.Infof("batch(%d) %d step(%s) got depend should(%d), picked(%d)",
				batch.Id, idx, action.Name, should, len(picked))
			if should < len(picked) {
				var dependFailed []*models.NodeState
				picked, dependFailed = picked[:should], picked[should:]
				for _, n := range dependFailed {
					n.Node.Status = models.STATUS_FAILED
					service.Cluster.UpdateBase(n.Node)
					failednum++
					lg.Infof("batch(%d) %d step(%s) node(%d) depend failed",
						batch.Id, idx, action.Name, n.Node.Ip)
				}
				updateNode(dependFailed, idx, models.STATUS_FAILED, action.Name)
			}
		}

		// check again
		flow, _ = service.Flow.GetFlowWithRel(batch.Flow.Id)
		if flow.Status == models.STATUS_STOPPED {
			lg.Infof("batch(%d) stopped at step %s(%d)", batch.Id, action.Name, idx)
			status = models.STATUS_STOPPED
			updateNode(picked, idx, status, action.Name)
			updateBatch(batch, idx, status)
			return
		}

		if len(picked) == 0 {
			lg.Infof("batch(%d) %d step(%s) has no node", batch.Id, idx, action.Name)
			if idx != finidx {
				continue
			} else {
				status = models.STATUS_FAILED
				updateBatch(batch, idx, status)
				return
			}
		}

		lg.Infof("batch(%d) %d step(%s) picked(%d)", batch.Id, idx, action.Name, len(picked))

		h := handler.GetHandler(action.Type)
		if h == nil {
			lg.Infof("batch(%d) %d step(%s) get nil handle", batch.Id, idx, action.Name)
			status = models.STATUS_FAILED
			updateNode(picked, idx, status, action.Name)
			updateBatch(batch, idx, status)
			return
		}

		cid := utils.GetCorrelationId(flow.Id, batch.Id)
		param := steps[idx].Values
		retry := steps[idx].Retry
		if retry == nil {
			retry = &models.RetryOption{
				RetryTimes:  0,
				IgnoreError: false,
			}
		}

		ok, failed := executor.Executor.RunStep(h, action, idx, picked, param, retry, cid)

		for _, n := range failed {
			n.Node.Status = models.STATUS_FAILED
			service.Cluster.UpdateBase(n.Node)
			failednum++
		}

		if idx == finidx {
			oknum = len(ok)
			for _, n := range ok {
				n.Node.Status = models.STATUS_SUCCESS
				service.Cluster.UpdateBase(n.Node)
			}
			updateNode(ok, idx, models.STATUS_SUCCESS, actions[idx].Name)
		} else {
			updateNode(ok, idx+1, models.STATUS_RUNNING, actions[idx+1].Name)
		}
	}

	if oknum > 0 {
		status = models.STATUS_SUCCESS
	} else {
		status = models.STATUS_FAILED
	}
	updateBatch(batch, finidx, status)
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
		lg.Infof("wait depend signal timeout")
	case <-ctx.Done():
		lg.Infof("wait depend canceled")
	}
	return 0
}

func updateNode(nn []*models.NodeState, step, status int, name string) {
	o := orm.NewOrm()
	for i := 0; i < len(nn); i++ {
		nn[i].Steps = name
		nn[i].StepNum = step
		nn[i].Status = status
		nn[i].UpdatedTime = time.Now()
		o.Update(nn[i])
	}
}

func updateBatch(batch *models.FlowBatch, step, status int) {
	batch.Status = status
	batch.Step = step
	batch.UpdatedTime = time.Now()
	orm.NewOrm().Update(batch)
}
