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
	Id     int    `json:"id" orm:"pk;auto"`
	Ip     string `json:"ip" orm:"null"`
	VmId   string `json:"vm_id" orm:"null"`
	Status int    `json:"status"`
	Pool   *Pool  `json:"-" orm:"rel(fk);on_delete(cascade)"`
	//Cluster*Cluster`json:"-" orm:"rel(fk);on_delete(cascade)"`
}
