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

func (f *FlowService) GetNodeByIp(ip string) (*models.Node, error) {
	o := orm.NewOrm()

	node := &models.Node{}
	err := o.QueryTable(node).Filter("ip", ip).One(node)
	if err != nil {
		return nil, err
	}

	return node, nil
}

/*
func (f *FlowService) GetNodesByFlowId(flowId int) ([]*models.Node, error) {
	o := orm.NewOrm()

	nodeList := make([]*models.Node, 0)

	_, err := o.QueryTable(&models.Node{}).Filter("Flow", flowId).All(&nodeList)

	return nodeList, err
}
*/

func (f *FlowService) GetNodeStatusByFlowId(flowId int) ([]*models.NodeState, error) {
	o := orm.NewOrm()

	nodeList := make([]*models.NodeState, 0)

	_, err := o.QueryTable(&models.NodeState{}).Filter("Flow", flowId).All(&nodeList)

	return nodeList, err
}

func (f *FlowService) DeleteNode(ips []string) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(&models.NodeState{}).Filter("ip__in", ips).Update(orm.Params{
		"deleted": models.DELETED,
	})

	if err != nil {
		beego.Error("Error when update nodestate %v with err:", ips, err)
	}

	return nil
}

func (f *FlowService) ListNodeRegister(obj interface{}, list interface{}, pids []int) (int, error) {
	o := orm.NewOrm()

	num, err := o.QueryTable(obj).Exclude("deleted", models.DELETED).Filter("steps", "register").Filter("pool_id__in", pids).All(list)

	if err != nil {
		return 0, err
	}

	return int(num), nil
}

/*
func (f *FlowService) GetLog(correlationId string) (string, error) {


}
*/
