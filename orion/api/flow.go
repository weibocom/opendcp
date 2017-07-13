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

package api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"weibo.com/opendcp/orion/executor"
	"weibo.com/opendcp/orion/handler"
	. "weibo.com/opendcp/orion/models"
	"weibo.com/opendcp/orion/service"
)

type FlowApi struct {
	baseAPI
}

type flow_struct struct {
	Id          int           `json:"id"`
	TplId       int           `json:"template_id"`
	TplName     string        `json:"template_name"`
	Name        string        `json:"task_name"`
	PoolName    string        `json:"pool_name"`
	Status      int           `json:"state"`
	Options     []*StepOption `json:"options"`
	StepLen     int           `json:"step_len"`
	OpUser      string        `json:"opr_user"`
	CreatedTime time.Time     `json:"created"`
	UpdatedTime time.Time     `json:"updated"`
	Stat        []int         `json:"stat`
}

type node_state struct {
	Id       int       `json:"id"`
	Ip       string    `json:"ip"`
	Status   int       `json:"state"`
	Steps    string    `json:"steps"`
	PoolName string    `json:"pool_name"`
	VmId     string    `json:"vm_id"`
	Created  time.Time `josn:"created"`
	Updated  time.Time `josn:""`
}

type detail_struct struct {
	States map[string]node_state `json:"states"`
}

type flowImpl struct {
	Id    int          `json:"id"`
	Name  string       `json:"name"`
	Desc  string       `json:"desc"`
	Steps []StepOption `json:"steps"`
	//Options []StepOption		  `json:"options"`
}

func (f *FlowApi) URLMapping() {
	f.Mapping("AppendFlowImpl", f.AppendFlowImpl)
	f.Mapping("ListFlowImpl", f.ListFlowImpl)
	f.Mapping("DeleteFlowImpl", f.DeleteFlowImpl)
	f.Mapping("GetFlowImpl", f.GetFlowImpl)
	f.Mapping("FlowImplUpdate", f.FlowImplUpdate)

	f.Mapping("ListFlow", f.ListFlow)
	f.Mapping("GetFlow", f.GetFlow)
	f.Mapping("GetTaskDetail", f.GetNodeStates)

	f.Mapping("ListTaskStep", f.ListTaskStep)

	//f.Mapping("AppendFlow", f.AppendFlow)
	f.Mapping("StartFlow", f.StartFlow)
	f.Mapping("RunFlow", f.RunFlow)
	f.Mapping("StopFlow", f.StopFlow)
	f.Mapping("PauseFlow", f.PauseFlow)
}

/*
 * Create a new task template
 */
func (c *FlowApi) AppendFlowImpl() {
	req := flowImpl{}

	err := c.Body2Json(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	//StepName check
	for _, step := range req.Steps {
		name := step.Name
		s := handler.GetActionImpl(name)
		if s == nil {
			c.ReturnFailed("step "+name+" not found, error:"+err.Error(), 404)
			return
		}
	}

	stepsByte, _ := json.Marshal(req.Steps)

	obj := &FlowImpl{
		Name:  req.Name,
		Steps: string(stepsByte),
		Desc:  string(req.Desc),
		//Options:string(optStr),
	}

	err = service.Flow.InsertBase(obj)
	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}

	c.ReturnSuccess(obj.Id)
}

/**
*  load flowimpl by id
 */
func (c *FlowApi) GetFlowImpl() {

	id := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	flowimpl := &FlowImpl{Id: idInt}
	err := service.Remote.GetBase(flowimpl)

	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}
	flowimpls := flowImpl{}
	flowimpls.Id = flowimpl.Id
	flowimpls.Name = flowimpl.Name
	flowimpls.Desc = flowimpl.Desc
	err = json.Unmarshal([]byte(flowimpl.Steps), &flowimpls.Steps)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	c.ReturnSuccess(flowimpls)
}

//列出TaskImpl
func (c *FlowApi) ListFlowImpl() {
	page := c.Query2Int("page", 1)
	pageSize := c.Query2Int("page_size", 10)

	c.CheckPage(&page, &pageSize)

	list := make([]FlowImpl, 0, pageSize)

	count, err := service.Flow.ListByPageWithSort(page, pageSize, &FlowImpl{}, &list, "-id")
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	liststruct := make([]flowImpl, len(list), pageSize)

	for i, fi := range list {
		liststruct[i].Id = fi.Id
		liststruct[i].Name = fi.Name
		liststruct[i].Desc = fi.Desc
		json.Unmarshal([]byte(fi.Steps), &liststruct[i].Steps)
	}

	c.ReturnPageContent(page, pageSize, count, liststruct)
}

//update flow
func (c *FlowApi) FlowImplUpdate() {
	id := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	req := flowImpl{}

	err := c.Body2Json(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	flowimpl := &FlowImpl{Id: idInt}
	err = service.Remote.GetBase(flowimpl)
	flowimpl.Desc = req.Desc
	stepStr, _ := json.Marshal(req.Steps)

	flowimpl.Steps = string(stepStr)
	//paramsStr, _ := json.Marshal(req.Options)
	//flowimpl.Options = string(paramsStr)

	err = service.Remote.UpdateBase(flowimpl)

	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}
	c.ReturnSuccess("")
}

//删除FlowImpl
func (f *FlowApi) DeleteFlowImpl() {
	id := f.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	err := service.Flow.DeleteBase(&FlowImpl{Id: idInt})
	if err != nil {
		f.ReturnFailed("data not found", 404)
		return
	}

	f.ReturnSuccess(nil)
}

//列出可用ActionImpl
func (f *FlowApi) ListTaskStep() {
	page := f.Query2Int("page", 1)
	pageSize := f.Query2Int("page_size", 10)

	f.CheckPage(&page, &pageSize)
	list := handler.GetAllActionImpl()
	f.ReturnPageContent(0, len(list), len(list), list)
}

func (f *FlowApi) RunFlow() {
	req := struct {
		TaskImplId int                      `json:"template_id"`
		TaskName   string                   `json:"task_name"`
		Timeout    int                      `json:"timeout"`
		Auto       int                      `json:"auto"`
		Ratio      int                      `json:"max_ratio"`
		MaxNum     int                      `json:"max_num"`
		RemoteUser string                   `json:"opr_user"`
		Nodes      []map[string]interface{} `json:"nodes"`
		Params     map[string]string        `json:"params"`
	}{}
	err := f.Body2Json(&req)
	if err != nil {
		beego.Error("RUN FLOW, json err:", err)
		f.ReturnFailed(err.Error(), 400)
		return
	}

	//ratio check
	if req.Ratio <= 0 || req.Ratio > 100 {
		f.ReturnFailed("ratio error", 400)
		return
	}

	stepLen := int(float64(len(req.Nodes)) * float64(req.Ratio) / 100.0)
	if stepLen > req.MaxNum {
		stepLen = req.MaxNum
	}

	opUser := f.Ctx.Input.Header("Authorization")

	nodes := make([]string, 0)
	nodeList := make([]*Node, 0, len(nodes))
	errorNodesIp := ""

	for _, n := range req.Nodes {
		nodeIp, ok := n["ip"].(string)
		if !ok {
			beego.Error("node :[", n, "] has not ip")
			continue
		}
		node, err := service.Flow.GetNodeByIp(nodeIp)
		if err != nil {
			beego.Error("node :[", nodeIp, "] not found...")
			errorNodesIp += nodeIp + ","
			continue
		}
		nodeList = append(nodeList, node)
	}

	if len(errorNodesIp) > 0 {
		f.ReturnFailed(fmt.Sprintf("node :[%s] not found...", errorNodesIp), 400)
	}

	context := make(map[string]interface{})
	context["overrideParams"] = map[string]interface{}{}
	context["opUser"] = opUser

	err = executor.Executor.Run(req.TaskImplId, req.TaskName,
		&executor.ExecOption{MaxNum: stepLen}, nodeList, context)

	if err != nil {
		beego.Error("Run", req.TaskName, "[", req.TaskImplId, "] fails:", err)
		f.ReturnFailed("run task fails: "+err.Error(), 400)
	} else {
		f.ReturnSuccess(nil)
	}
}

/*
 * Create a new task instance
 */
/*
func (f *FlowApi) AppendFlow() {
	req := struct {
		TaskImplId int               `json:"task_def_id"`
		Ratio      float32           `json:"ratio"`
		MaxNum     int               `json:"max_num"`
		RemoteUser string            `json:"opr_user"`
		Nodes      []string          `json:"nodes"`
		Params     map[string]string `json:"params"`
	}{}

	err := f.Body2Json(&req)
	if err != nil {
		f.ReturnFailed(err.Error(), 400)
		return
	}

	//ratio check
	if req.Ratio <= 0 || req.Ratio > 1 {
		f.ReturnFailed("ratio error", 400)
		return
	}

	//flowImpl check
	flowImpl := &FlowImpl{Id: req.TaskImplId}
	err = service.Flow.GetBase(flowImpl)
	if err != nil {
		f.ReturnFailed("flowimpl not found", 404)
		return
	}

	//nodes check
	if len(req.Nodes) == 0 {
		f.ReturnFailed("nodes is empty", 400)
		return
	}
	for _, nodeIp := range req.Nodes {
		_, err = service.Flow.GetNodeByIp(nodeIp)
		if err != nil {
			f.ReturnFailed("node not found", 404)
			return
		}
	}

	//计算步长
	stepLen := int(float32(len(req.Nodes)) * req.Ratio)
	if stepLen > req.MaxNum {
		stepLen = req.MaxNum
	}

	//步长check
	if stepLen <= 0 {
		f.ReturnFailed("stepLen is empty", 400)
		return
	}

	//Flow落地
	paramsByte, _ := json.Marshal(req.Params)

	flow := &Flow{
		Options:      string(paramsByte),
		CreatedTime: time.Now(),
		Impl:        flowImpl,
		StepLen:     stepLen,
	}
	err = service.Flow.InsertBase(flow)
	if err != nil {
		f.ReturnFailed(err.Error(), 400)
		return
	}

	f.ReturnSuccess(flow.Id)
}
*/

func (f *FlowApi) StartFlow() {
	_id := f.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		f.ReturnFailed("Bad flow id: "+_id, 400)
		return
	}

	obj, err := service.Flow.GetFlowWithRel(id)
	if err != nil {
		f.ReturnFailed("flow not found id: "+_id, 400)
		return
	}

	//重新读取配置文件
	if obj.Impl != nil {
		obj.Options = obj.Impl.Steps
	} else {
		beego.Error("flowImp is not found " + _id)
		f.ReturnFailed("flowImp is not found "+_id, 400)
		return
	}
	err = executor.Executor.Start(obj)
	if err != nil {
		beego.Error("start flow ", _id, "fails: ", err)
		f.ReturnFailed("start task "+_id+" fails： "+err.Error(), 400)
	} else {
		f.ReturnSuccess(nil)
	}

}

func (f *FlowApi) StopFlow() {
	_id := f.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		f.ReturnFailed("Bad flow id: "+_id, 400)
		return
	}

	obj, err := service.Flow.GetFlowWithRel(id)
	if err != nil {
		f.ReturnFailed("flow not found id: "+_id, 400)
		return
	}

	err = executor.Executor.Stop(obj)
	if err != nil {
		beego.Error("stop flow ", _id, "fails: ", err)
		f.ReturnFailed("stop task "+_id+" fails: "+err.Error(), 400)
	} else {
		f.ReturnSuccess(nil)
	}
}

func (f *FlowApi) PauseFlow() {
	_id := f.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		f.ReturnFailed("Bad flow id: "+_id, 400)
		return
	}

	obj, err := service.Flow.GetFlowWithRel(id)
	if err != nil {
		f.ReturnFailed("flow not found id: "+_id, 400)
		return
	}

	err = executor.Executor.Pause(obj)
	if err != nil {
		beego.Error("pause flow ", _id, "fails: ", err)
		f.ReturnFailed("pause task "+_id+" fails: "+err.Error(), 400)
	} else {
		f.ReturnSuccess(nil)
	}
}

func (f *FlowApi) GetFlow() {
	_id := f.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(_id)
	if err != nil {
		f.ReturnFailed("Bad flow id: "+_id, 400)
		return
	}

	obj, err := service.Flow.GetFlowWithRel(id)
	if err != nil {
		f.ReturnFailed("flow not found id: "+_id, 400)
		return
	}

	flowstru := flow_struct{}
	f.popFlowStruct(obj, &flowstru)

	//def := flow.Impl
	f.ReturnSuccess(flowstru)
}

func (f *FlowApi) ListFlow() {
	page := f.Query2Int("page", 1)
	pageSize := f.Query2Int("page_size", 10)

	f.CheckPage(&page, &pageSize)

	list := make([]Flow, 0, pageSize)

	count, err := service.Flow.ListByPageWithSort(page, pageSize, &Flow{}, &list, "-id")
	if err != nil {
		f.ReturnFailed(err.Error(), 400)
		return
	}

	liststruct := make([]flow_struct, len(list), pageSize)

	for i, fi := range list {
		f.popFlowStruct(&fi, &liststruct[i])
	}

	f.ReturnPageContent(page, pageSize, count, liststruct)
}

func (f *FlowApi) GetNodeStates() {
	_id := f.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(_id)

	flow := &Flow{Id: id}
	err := service.Flow.GetBase(flow)
	if err != nil {
		f.ReturnFailed("Flow not found: "+strconv.Itoa(id), 400)
		return
	}

	list := make([]NodeState, 0)
	_, err = service.Flow.ListByPageWithFilter(0, 1000,
		&NodeState{}, &list, "Flow", id)
	if err != nil {
		f.ReturnFailed(err.Error(), 400)
		return
	}

	ret := make(map[string][]node_state)
	for _, ns := range list {
		state := ns.Status
		key := strconv.Itoa(state)
		st := node_state{}
		f.popNodeStruct(&ns, &st)
		list := ret[key]
		if list == nil {
			list = make([]node_state, 0)
			ret[key] = list
		}
		list = append(list, st)
		ret[key] = list
	}

	out, _ := json.Marshal(ret)
	beego.Debug("node states json:\n", string(out))

	f.ReturnSuccess(ret)
}

func (f *FlowApi) GetFlowLogById() {
	_id := f.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(_id)

	flow := &Flow{Id: id}
	err := service.Flow.GetBase(flow)
	if err != nil {
		f.ReturnFailed("Flow not found: "+strconv.Itoa(id), 400)
		return
	}

	logList := make([]Logs, 0)
	_, err = service.Flow.ListByPageWithFilter(0, 1000, &Logs{}, &logList, "fid", id)
	if err != nil {
		f.ReturnFailed(err.Error(), 400)
		return
	}

	f.ReturnSuccess(logList)
}

// GetLog get log using nodeState Id
func (f *FlowApi) GetLog() {
	idStr := f.Ctx.Input.Param(":nsid")
	nodeStateId, err := strconv.Atoi(idStr)
	if err != nil {
		f.ReturnFailed("Bad node state : "+idStr, 400)
		return
	}

	nodeState := &NodeState{Id: nodeStateId}
	err = service.Flow.GetBase(nodeState)
	if err != nil {
		f.ReturnFailed("Node state not found: "+strconv.Itoa(nodeStateId), 400)
		return
	}

	logs, err := getLog(nodeState)
	if err != nil {
		f.ReturnFailed(err.Error(), 400)
		return
	}

	f.ReturnSuccess(logs)
}

var handlers = make(map[string]*handler.Handler)

func getLog(nodeState *NodeState) ([]map[string]string, error) {
	logs := make([]map[string]string, 0)

	// load flow definition
	ss := strings.Split(nodeState.CorrId, "-")
	flowId, err := strconv.Atoi(ss[0])
	if err != nil {
		return logs, err
	}

	flow := &Flow{Id: flowId}
	err = service.Flow.GetBase(flow)
	if err != nil {
		return logs, err
	}

	// load options
	var stepOptions []*StepOption
	err = json.Unmarshal([]byte(flow.Options), &stepOptions)
	if err != nil {
		beego.Error("Marshal steps: ", flow.Options, ", err:", err)
		return logs, err
	}

	// load logs
	currStep := nodeState.Steps
	for _, option := range stepOptions {
		step := handler.GetActionImpl(option.Name)

		hdl := handlers[step.Type]
		if hdl == nil {
			h := handler.GetHandler(step.Type)
			hdl = &h
			handlers[step.Type] = hdl
		}

		log := (*hdl).GetLog(nodeState)
		stepLog := map[string]string{
			step.Name: log,
		}
		logs = append(logs, stepLog)

		if currStep == step.Name {
			break
		}
	}

	return logs, nil
}

func (f *FlowApi) popFlowStruct(obj *Flow, flowstru *flow_struct) {
	flowstru.Id = obj.Id
	flowstru.Name = obj.Name
	flowstru.Status = obj.Status
	json.Unmarshal([]byte(obj.Options), &flowstru.Options)
	flowstru.StepLen = obj.StepLen
	flowstru.OpUser = obj.OpUser
	flowstru.CreatedTime = obj.CreatedTime
	flowstru.UpdatedTime = obj.UpdatedTime
	flowstru.TplId = obj.Impl.Id
	flowstru.TplName = obj.Impl.Name

	if obj.Pool != nil {
		flowstru.PoolName = obj.Pool.Name
	}

	// get statistics
	stat := make([]int, 5)

	states := make([]*NodeState, 0)
	_, err := service.Flow.ListByPageWithFilter(0, 10000,
		&NodeState{}, &states, "Flow", obj.Id)

	if err != nil {
		beego.Error("Fail to get node states for flow", obj.Name, obj.Id, err.Error())
	}

	for _, ns := range states {
		stat[ns.Status] += 1
	}

	flowstru.Stat = stat
}

func (f *FlowApi) popNodeStruct(obj *NodeState, state *node_state) {
	state.Id = obj.Id
	state.Ip = obj.Ip
	state.VmId = obj.VmId
	state.Status = obj.Status
	state.Steps = obj.Steps
	state.Created = obj.CreatedTime
	state.Updated = obj.UpdatedTime
	if obj.Pool != nil {
		state.PoolName = obj.Pool.Name
	}
}
