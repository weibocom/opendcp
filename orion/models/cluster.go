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

package models

type Cluster struct {
	Id   int    `json:"id" orm:"pk;auto"`
	Name string `json:"name" orm:"size(50)"`
	Desc string `json:"desc" orm:"size(255);null"`
	Biz  string `json:"biz"` //产品线
}

type Service struct {
	Id          int      `json:"id" orm:"pk;auto"`
	Name        string   `json:"name" orm:"size(50)"`
	Desc        string   `json:"desc" orm:"size(255);null"`
	ServiceType string   `json:"service_type"` //服务类型
	DockerImage string   `json:"docker_image"` //Docker镜像
	Cluster     *Cluster `json:"-" orm:"rel(fk);on_delete(cascade)"`
}

type Pool struct {
	Id      int      `json:"id" orm:"pk;auto"`
	Name    string   `json:"name" orm:"size(50)"`
	Desc    string   `json:"desc" orm:"size(255);null"`
	VmType  int      `json:"vm_type"` //VM类型
	SdId    int      `json:"sd_id"`   //服务发现ID
	Tasks   string   `json:"tasks"`   //对应任务(task_name arr)
	Service *Service `json:"-" orm:"rel(fk);on_delete(cascade)"`
}

type Node struct {
	Id       int    `json:"id" orm:"pk;auto"`
	Ip       string `json:"ip" orm:"null"`
	VmId     string `json:"vm_id" orm:"null"`
	Status   int    `json:"status"`
	Pool     *Pool  `json:"-" orm:"rel(fk);on_delete(cascade)"`
	NodeType string `json:"node_type" orm:"default(manual)"`
	//Cluster*Cluster`json:"-" orm:"rel(fk);on_delete(cascade)"`
}
