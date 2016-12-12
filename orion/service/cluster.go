/*
 *  Copyright 2009-2016 Weibo, Inc.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
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
func (c* ClusterService) SearchPoolByIP(ips []string) map[string]int {

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
