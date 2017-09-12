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

package service

import (
	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
	"time"
	"weibo.com/opendcp/orion/models"
)

type FlowService struct {
	BaseService
}

func (f *FlowService) GetFlowWithRel(id int) (*models.Flow, error) {
	o := orm.NewOrm()

	flow := &models.Flow{}
	err := o.QueryTable(flow).Filter("Id", id).RelatedSel().One(flow)
	if err != nil {
		return nil, err
	}

	return flow, nil
}

func (f *FlowService) GetFlowImplWithRel(id int) (*models.FlowImpl, error) {
	o := orm.NewOrm()

	flowImpl := &models.FlowImpl{}
	err := o.QueryTable(flowImpl).Filter("Id", id).RelatedSel().One(flowImpl)
	if err != nil {
		return nil, err
	}

	return flowImpl, nil
}

func (f *FlowService) GetActionImplByName(name string) (*models.ActionImpl, error) {
	o := orm.NewOrm()

	action := &models.ActionImpl{}
	err := o.QueryTable(action).Filter("name", name).One(action)
	if err != nil {
		return nil, err
	}

	return action, nil
}

func (f *FlowService) GetNodeByIp(ip string) (*models.NodeState, error) {
	o := orm.NewOrm()

	node := &models.NodeState{}
	err := o.QueryTable(node).Filter("ip", ip).Filter("deleted", false).One(node)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func (f *FlowService) GetNodeById(id int) (*models.NodeState, error) {
	o := orm.NewOrm()

	node := &models.NodeState{Id: id}
	err := o.Read(node)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func (f *FlowService) UpdateNodeMachine(state *models.NodeState) error {
	o := orm.NewOrm()
	if state.Ip == "" || state.Ip == "-" {
		_, err := o.Update(state, "vm_id", "updated_time")
		return err
	}
	_, err := o.Update(state, "ip", "vm_id", "updated_time")
	return err
}

func (f *FlowService) UpdateNode(state *models.NodeState) error {
	o := orm.NewOrm()
	_, err := o.Update(state,
		"status", "steps", "step_num", "log", "last_op",
		"step_run_time", "run_time", "updated_time",
	)
	return err
}

func (f *FlowService) UpdateNodeRunTime(state *models.NodeState) error {
	o := orm.NewOrm()
	_, err := o.Update(state,
		"step_run_time", "run_time", "updated_time",
	)
	return err
}

func (f *FlowService) UpdateNodeWithoutStatus(state *models.NodeState) error {
	o := orm.NewOrm()
	_, err := o.Update(state,
		"steps", "step_num", "log",
		"last_op", "step_run_time", "run_time",
		"updated_time",
	)
	return err
}

func (f *FlowService) DeleteNodeById(state *models.NodeState) error {
	o := orm.NewOrm()
	_, err := o.Update(state, "deleted", "updated_time")
	return err
}

func (f *FlowService) ChangeNodeStatusById(state *models.NodeState) error {
	o := orm.NewOrm()
	_, err := o.Update(state, "status", "updated_time")
	return err
}

func (f *FlowService) ChangeNodeStatusAndLogsById(state *models.NodeState) error {
	o := orm.NewOrm()
	_, err := o.Update(state, "log", "status", "updated_time")
	return err
}
func (f *FlowService) GetNodeStatusByFlowId(flowId int) ([]*models.NodeState, error) {
	o := orm.NewOrm()

	nodeList := make([]*models.NodeState, 0)

	_, err := o.QueryTable(&models.NodeState{}).Filter("Flow", flowId).Filter("deleted", false).All(&nodeList)

	return nodeList, err
}

func (f *FlowService) GetAllNodeStatesByFlowId(flowId int) ([]*models.NodeState, error) {
	o := orm.NewOrm()

	nodeList := make([]*models.NodeState, 0)

	_, err := o.QueryTable(&models.NodeState{}).Filter("Flow", flowId).All(&nodeList)

	return nodeList, err
}

func (f *FlowService) GetAllNodeStatusByFlowId(flowId int, status int) ([]*models.NodeState, error) {
	o := orm.NewOrm()

	nodeList := make([]*models.NodeState, 0)

	_, err := o.QueryTable(&models.NodeState{}).Filter("Flow", flowId).Filter("status", status).All(&nodeList)

	return nodeList, err
}

func (f *FlowService) DeleteNode(ips []string) error {
	o := orm.NewOrm()
	for _, ip := range ips {
		n := &models.NodeState{Ip: ip}
		err := o.Read(n, "ip")
		if err != nil {
			return err
		}
		n.Deleted = true
		n.UpdatedTime = time.Now()
		_, err = o.Update(n, "deleted", "updated_time")
		if err != nil {
			beego.Error("Error when update nodestate ", ip, " with err:", ip, err)
			return err
		}
	}

	return nil
}

func (f *FlowService) ListNodeRegister(obj interface{}, list interface{}, pids []int) (int, error) {
	o := orm.NewOrm()

	num, err := o.QueryTable(obj).Exclude("deleted", true).Filter("steps", "register").Filter("pool_id__in", pids).All(list)

	if err != nil {
		return 0, err
	}

	return int(num), nil
}

func (f *FlowService) UpdateFlowStatus(flow *models.Flow) error {
	o := orm.NewOrm()

	_, err := o.Update(flow, "status", "updated_time")

	return err
}
