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
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"fmt"
	"time"
	"weibo.com/opendcp/orion/models"
)

type ClusterService struct {
	BaseService
}

const NODE_STATE_TABLE = "node_state"

func (c *ClusterService) GetAllExecTask() ([]*models.ExecTask, error) {
	var (
		taskList []*models.ExecTask
		o        = orm.NewOrm()
	)
	if _, err := o.QueryTable(&models.ExecTask{}).
		RelatedSel().All(&taskList); err != nil {
		return nil, fmt.Errorf("db load ExecTask failed: %v", err)
	}
	for _, task := range taskList {
		if _, err := o.LoadRelated(task, "CronItems"); err != nil {
			return nil, fmt.Errorf("db load %d CronItems failed: %v", task.Id, err)
		}
		if _, err := o.LoadRelated(task, "DependItems"); err != nil {
			return nil, fmt.Errorf("db load %d DependItems failed: %v", task.Id, err)
		}
	}

	return taskList, nil
}
func (c *ClusterService) AppendIpList(ips []string, pool *models.Pool, label string) []int {
	o := orm.NewOrm()
	beego.Info(label)
	respDatas := make([]int, 0, len(ips))

	for _, ip := range ips {
		data := models.NodeState{
			Ip:          ip,
			Pool:        pool,
			Flow:        &models.Flow{Id: 0},
			Status:      models.STATUS_INIT,
			CreatedTime: time.Now(),
			UpdatedTime: time.Now(),
			NodeType:    models.Manual,
			Deleted:     false,
			Label:       label,
		}
		_, err := o.Insert(&data)
		if err != nil {
			beego.Error("insert node failed, error: ", err)
			return nil
		}
		respDatas = append(respDatas, data.Id)
	}

	return respDatas
}

func (c *ClusterService) AppendIp(ip string, instanceId string, pool *models.Pool, label string) int {
	o := orm.NewOrm()
	data := models.NodeState{
		Ip:          ip,
		VmId:	     instanceId,
		Pool:        pool,
		Flow:        &models.Flow{Id: 0},
		Status:      models.STATUS_INIT,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
		NodeType:    models.Manual,
		Deleted:     false,
		Label:       label,
	}
	_, err := o.Insert(&data)
	if err != nil {
		beego.Error("insertinode failed,error:", err)
		return 0
	}
	return data.Id

}

// SearchPoolByIP returs the pool id of the ip given, if it exists
func (c *ClusterService) SearchPoolByIP(ips []string) map[string]int {
	o := orm.NewOrm()

	result := make(map[string]int)
	for _, ip := range ips {
		node := &models.NodeState{Ip: ip}
		err := o.QueryTable(node).Filter("deleted", false).Filter("ip", ip).One(node)
		id := -1
		if err == nil && !node.Deleted {
			id = node.Pool.Id
		}
		result[ip] = id
	}

	return result
}

func (c *ClusterService) ListNodesByType(pool_id int, node_type string) ([]models.NodeState, error) {
	o := orm.NewOrm()

	list := make([]models.NodeState, 0)

	_, err := o.QueryTable(&models.NodeState{}).Filter("Pool", pool_id).Filter("NodeType", node_type).All(&list)

	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *ClusterService) GetAllLabels() ([]string, error) {
	o := orm.NewOrm()
	var providers []string
	sql := fmt.Sprintf("SELECT DISTINCT LABEL FROM %s ", NODE_STATE_TABLE)
	_, err := o.Raw(sql).QueryRows(&providers)
	if err != nil {
		return nil, err
	}
	return providers, nil
}
