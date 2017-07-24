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

	"weibo.com/opendcp/orion/models"
)

type ClusterService struct {
	BaseService
}

func (c *ClusterService) AppendIpList(ips []string, pool *models.Pool) []int {
	o := orm.NewOrm()

	respDatas := make([]int, 0, len(ips))

	for _, ip := range ips {
		data := models.Node{
			Ip:   ip,
			Pool: pool,
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

// SearchPoolByIP returs the pool id of the ip given, if it exists
func (c *ClusterService) SearchPoolByIP(ips []string) map[string]int {

	result := make(map[string]int)
	for _, ip := range ips {
		node := &models.Node{Ip: ip}
		err := c.GetBy(node, "ip")
		id := -1
		if err == nil {
			id = node.Pool.Id
		}
		result[ip] = id
	}

	return result
}

func (c *ClusterService) ListNodesByType(pool_id int, node_type string) ([]models.Node, error) {
	o := orm.NewOrm()

	list := make([]models.Node, 0)

	_, err := o.QueryTable(&models.Node{}).Filter("Pool", pool_id).Filter("NodeType", node_type).All(&list)

	if err != nil {
		return nil, err
	}
	return list, nil
}
