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
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"

	"weibo.com/opendcp/orion/models"
	"weibo.com/opendcp/orion/service"
)

type ClusterApi struct {
	baseAPI
}

type cluster_struct struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
	Biz  string `json:"biz"`
}

type pool_struct struct {
	Id        int                    `json:"id"`
	Name      string                 `json:"name"`
	Desc      string                 `json:"desc"`
	VmType    int                    `json:"vm_type"`
	SdId      int                    `json:"sd_id"`
	Tasks     map[string]interface{} `json:"tasks"`
	ServiceId int                    `json:"service_id"`
	Nodecount int                    `json:"node_count"`
}

type service_struct struct {
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	ServiceType string `json:"service_type"`
	DockerImage string `json:"docker_image"`
	ClusterId   int    `json:"cluster_id"`
}

//因跟 models.Service.ClusterId 结构不一样..所以需要单独定义.
type service_view_struct struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	ServiceType string `json:"service_type"`
	DockerImage string `json:"docker_image"`
	ClusterId   int    `json:"cluster_id"`
}

type delnodes_struct struct {
	nodes []string `json:"nodes"`
}

func (c *ClusterApi) URLMapping() {
	c.Mapping("ClusterInfo", c.ClusterInfo)
	c.Mapping("ClusterList", c.ClusterList)
	c.Mapping("ClusterAppend", c.ClusterAppend)
	c.Mapping("ClusterDelete", c.ClusterDelete)
	c.Mapping("ClusterUpdate", c.ClusterUpdate)

	c.Mapping("ServiceInfo", c.ServiceInfo)
	c.Mapping("ServiceList", c.ServiceList)
	c.Mapping("ServiceAppend", c.ServiceAppend)
	c.Mapping("ServiceDelete", c.ServiceDelete)
	c.Mapping("ServiceUpdate", c.ServiceUpdate)

	c.Mapping("PoolInfo", c.PoolInfo)
	c.Mapping("PoolList", c.PoolList)
	c.Mapping("PoolAppend", c.PoolAppend)
	c.Mapping("PoolDelete", c.PoolDelete)
	c.Mapping("PoolUpdate", c.PoolUpdate)
	c.Mapping("AllPoolList", c.AllPoolList)

	c.Mapping("NodeList", c.NodeList)
	c.Mapping("NodeAppend", c.NodeAppend)
	c.Mapping("NodeDelete", c.NodeDelete)

	c.Mapping("search_by_ip", c.SearchPoolByIP)
}

//集群管理
func (c *ClusterApi) ClusterInfo() {
	idInt := c.clusterCheckId()
	if idInt < 1 {
		c.ReturnFailed("id is error !", 400)
		return
	}

	obj := &models.Cluster{Id: idInt}
	err := service.Cluster.GetBase(obj)
	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}

	c.ReturnSuccess(obj)
}

func (c *ClusterApi) ClusterList() {
	page := c.Query2Int("page", 1)
	pageSize := c.Query2Int("page_size", 10)

	c.CheckPage(&page, &pageSize)

	list := make([]models.Cluster, 0, pageSize)

	count, err := service.Cluster.ListByPageWithSort(page, pageSize, &models.Cluster{}, &list, "-id")
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	c.ReturnPageContent(page, pageSize, count, list)

}

func (c *ClusterApi) ClusterAppend() {
	req := cluster_struct{}
	err := c.clusterCheckParam(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	data := models.Cluster{
		Name: req.Name,
		Desc: req.Desc,
		Biz:  req.Biz,
	}
	err = service.Cluster.InsertBase(&data)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}
	c.ReturnSuccess(data.Id)
}

func (c *ClusterApi) ClusterDelete() {
	idInt := c.clusterCheckId()
	if idInt < 1 {
		c.ReturnFailed("id is error !", 400)
		return
	}

	obj := &models.Cluster{Id: idInt}
	err := service.Cluster.GetBase(obj)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	list := make([]models.Service, 0, 1)

	count, err := service.Cluster.ListByPageWithFilter(0, 1, &models.Service{}, &list, "cluster_id", idInt)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	if count > 0 {
		c.ReturnFailed("service exists in this cluster", 400)
		return
	}

	err = service.Cluster.DeleteBase(obj)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	c.ReturnSuccess(nil)
}

func (c *ClusterApi) ClusterUpdate() {
	idInt := c.clusterCheckId()
	if idInt < 1 {
		c.ReturnFailed("id is error !", 400)
		return
	}

	req := cluster_struct{}
	err := c.clusterCheckParam(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	cluster := &models.Cluster{Id: idInt}
	err = service.Remote.GetBase(cluster)
	if len(cluster.Name) < 1 {
		c.ReturnFailed("old data not found !", 400)
		return
	}

	cluster.Desc = req.Desc
	cluster.Biz = req.Biz

	err = service.Remote.UpdateBase(cluster)

	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}
	c.ReturnSuccess("")

}

func (c *ClusterApi) clusterCheckId() int {
	id := c.Ctx.Input.Param(":id")
	if len(id) < 1 {
		return 0
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}

	return idInt
}

func (c *ClusterApi) clusterCheckParam(req *cluster_struct) error {
	err := c.Body2Json(&req)
	if err != nil {
		return err
	}

	if len(req.Name) < 1 {
		return errors.New("param is error!")
	}

	return nil
}

//服务管理
func (c *ClusterApi) ServiceInfo() {
	idInt := c.serviceCheckId()
	if idInt < 1 {
		c.ReturnFailed("id is error !", 400)
		return
	}

	obj := &models.Service{Id: idInt}
	err := service.Cluster.GetBase(obj)
	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}

	servicestru := service_view_struct{}
	servicestru.Id = obj.Id
	servicestru.Name = obj.Name
	servicestru.Desc = obj.Desc
	servicestru.ServiceType = obj.ServiceType
	servicestru.DockerImage = obj.DockerImage
	servicestru.ClusterId = obj.Cluster.Id

	c.ReturnSuccess(servicestru)
}

func (c *ClusterApi) ServiceList() {
	id := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	page := c.Query2Int("page", 1)
	pageSize := c.Query2Int("page_size", 10)

	c.CheckPage(&page, &pageSize)

	list := make([]models.Service, 0, pageSize)

	count, err := service.Cluster.ListByPageWithFilter(page, pageSize, &models.Service{}, &list, "cluster_id", idInt)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	liststruct := make([]service_view_struct, len(list), pageSize)

	for i, fi := range list {
		liststruct[i].Id = fi.Id
		liststruct[i].Name = fi.Name
		liststruct[i].Desc = fi.Desc
		liststruct[i].ServiceType = fi.ServiceType
		liststruct[i].DockerImage = fi.DockerImage
		liststruct[i].ClusterId = fi.Cluster.Id
	}

	c.ReturnPageContent(page, pageSize, count, liststruct)
}

func (c *ClusterApi) ServiceAppend() {
	req := service_struct{}
	err := c.serviceCheckParam(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	data := models.Service{
		Name:        req.Name,
		Desc:        req.Desc,
		ServiceType: req.ServiceType,
		DockerImage: req.DockerImage,
		Cluster: &models.Cluster{
			Id: req.ClusterId,
		},
	}

	err = service.Cluster.InsertBase(&data)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	c.ReturnSuccess(data.Id)
}

func (c *ClusterApi) ServiceDelete() {
	idInt := c.serviceCheckId()
	if idInt < 1 {
		c.ReturnFailed("id is error !", 400)
		return
	}

	serv := &models.Service{Id: idInt}
	err := service.Cluster.GetBase(serv)
	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}

	poolexists := &models.Pool{Service: serv}
	err = service.Cluster.GetBy(poolexists, "service_id")
	if err == nil {
		c.ReturnFailed("sub pool exists", 400)
		return
	}

	err = service.Cluster.DeleteBase(serv)
	if err != nil {
		c.ReturnFailed("data not found", 400)
		return
	}

	c.ReturnSuccess(nil)
}

func (c *ClusterApi) ServiceUpdate() {
	idInt := c.serviceCheckId()
	if idInt < 1 {
		c.ReturnFailed("id is error !", 400)
		return
	}

	req := service_struct{}
	err := c.serviceCheckParam(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	servicem := &models.Service{Id: idInt}
	err = service.Remote.GetBase(servicem)
	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}

	servicem.Desc = req.Desc
	servicem.ServiceType = req.ServiceType
	servicem.DockerImage = req.DockerImage
	servicem.Cluster = &models.Cluster{
		Id: req.ClusterId,
	}

	err = service.Remote.UpdateBase(servicem)
	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}

	c.ReturnSuccess("")
}

func (c *ClusterApi) serviceCheckId() int {
	id := c.Ctx.Input.Param(":id")
	if len(id) < 1 {
		return 0
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}

	return idInt
}

func (c *ClusterApi) serviceCheckParam(req *service_struct) error {
	err := c.Body2Json(&req)
	if err != nil {
		return err
	}

	if len(req.Name) < 1 || len(req.Name) < 1 || len(req.DockerImage) < 1 || req.ClusterId < 1 {
		return errors.New("param is error!")
	}

	return nil
}

//服务池管理
func (c *ClusterApi) PoolInfo() {
	id := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	obj := &models.Pool{Id: idInt}

	err := service.Cluster.GetBase(obj)
	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}

	poolstr := pool_struct{}
	poolstr.Id = obj.Id
	poolstr.Name = obj.Name
	poolstr.Desc = obj.Desc
	poolstr.VmType = obj.VmType
	poolstr.SdId = obj.SdId
	json.Unmarshal([]byte(obj.Tasks), &poolstr.Tasks)
	poolstr.ServiceId = obj.Service.Id

	c.ReturnSuccess(poolstr)
}

func (c *ClusterApi) PoolList() {
	id := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	page := c.Query2Int("page", 1)
	pageSize := c.Query2Int("page_size", 10)

	c.CheckPage(&page, &pageSize)

	list := make([]models.Pool, 0, pageSize)

	count, err := service.Cluster.ListByPageWithFilter(page, pageSize,
		&models.Pool{}, &list, "service_id", idInt)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	liststruct := make([]pool_struct, len(list), pageSize)

	for i, fi := range list {
		liststruct[i].Id = fi.Id
		liststruct[i].Name = fi.Name
		liststruct[i].Desc = fi.Desc
		liststruct[i].VmType = fi.VmType
		liststruct[i].SdId = fi.SdId

		json.Unmarshal([]byte(fi.Tasks), &liststruct[i].Tasks)
		liststruct[i].ServiceId = fi.Service.Id

		nodeCount, err := service.Cluster.GetCount(&models.Node{}, "Pool", &models.Pool{Id: fi.Id})
		if err != nil {
			c.ReturnFailed(err.Error(), 400)
			return
		}
		liststruct[i].Nodecount = nodeCount

	}

	c.ReturnPageContent(page, pageSize, count, liststruct)
}

func (c *ClusterApi) PoolAppend() {
	req := pool_struct{}

	err := c.Body2Json(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}
	taskbytes, _ := json.Marshal(req.Tasks)
	data := models.Pool{
		Name:   req.Name,
		Desc:   req.Desc,
		VmType: req.VmType,
		SdId:   req.SdId,
		Tasks:  string(taskbytes),
		Service: &models.Service{
			Id: req.ServiceId,
		},
	}
	err = service.Cluster.InsertBase(&data)

	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}
	c.ReturnSuccess(data.Id)
}

func (c *ClusterApi) PoolDelete() {
	id := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	pool := &models.Pool{Id: idInt}
	err := service.Cluster.GetBase(pool)
	if err != nil {
		c.ReturnFailed("data not found", 400)
		return
	}

	list := make([]models.Node, 0, 1)

	count, err := service.Cluster.ListByPageWithFilter(0, 1,
		&models.Node{}, &list, "pool_id", pool.Id)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	if count > 0 {
		c.ReturnFailed("node exists in this pool", 400)
		return
	}

	err = service.Cluster.DeleteBase(&models.Pool{Id: idInt})
	if err != nil {
		c.ReturnFailed("fail to delete pool", 400)
		return
	}

	c.ReturnSuccess(nil)
}

func (c *ClusterApi) PoolUpdate() {

	id := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	req := pool_struct{}

	err := c.Body2Json(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	fmt.Println("bad")

	pool := &models.Pool{Id: idInt}
	err = service.Remote.GetBase(pool)

	pool.Desc = req.Desc
	pool.VmType = req.VmType
	pool.SdId = req.SdId
	taskbytes, _ := json.Marshal(req.Tasks)
	pool.Tasks = string(taskbytes)

	err = service.Remote.UpdateBase(pool)

	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}
	c.ReturnSuccess("")

}

//节点管理
func (c *ClusterApi) NodeList() {
	idStr := c.Ctx.Input.Param(":id")
	poolId, _ := strconv.Atoi(idStr)

	page := c.Query2Int("page", 1)
	pageSize := c.Query2Int("page_size", 10)

	c.CheckPage(&page, &pageSize)

	list := make([]models.Node, 0, pageSize)

	count, err := service.Cluster.ListByPageWithFilter(page, pageSize,
		&models.Node{}, &list, "pool_id", poolId)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	c.ReturnPageContent(page, pageSize, count, list)
}

func (c *ClusterApi) NodeAppend() {
	_id := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(_id)
	req := struct {
		Ips []string `json:"nodes"`
	}{}

	err := c.Body2Json(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	//repeat ip check

	for _, ip := range req.Ips {
		count, err1 := service.Cluster.GetCount(&models.Node{}, "ip", ip)
		if err1 != nil || count > 0 {
			c.ReturnFailed("ip exists already", 404)
			return
		}
	}

	//Pool check
	pool := &models.Pool{
		Id: id,
	}

	err = service.Cluster.GetBase(pool)
	if err != nil {
		c.ReturnFailed("pool_id is not vaild", 404)
		return
	}

	//check IP list
	for _, ip := range req.Ips {
		if strings.TrimSpace(ip) == "" {
			c.ReturnFailed("ip is empty", 400)
			return
		}
	}

	back := service.Cluster.AppendIpList(req.Ips, pool)

	c.ReturnSuccess(back)
}

func (c *ClusterApi) NodeDelete() {

	//req := delnodes_struct{}
	req := struct {
		NodeIds []int `json:"nodes"`
	}{}
	err := c.Body2Json(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	fmt.Println(req)
	for _, id := range req.NodeIds {

		//idInt,_:=strconv.Atoi(id)
		err := service.Cluster.DeleteBase(&models.Node{Id: id})
		if err != nil {
			beego.Error("Error when deleting id:", id, ", error:", err)
			c.ReturnFailed("error when delete id: "+strconv.Itoa(id)+", err:"+err.Error(), 400)
			return
		}
	}

	c.ReturnSuccess(nil)

}

func (c *ClusterApi) SearchPoolByIP() {
	ipList := c.Ctx.Input.Param(":iplist")

	ips := strings.Split(ipList, ",")
	poolIds := service.Cluster.SearchPoolByIP(ips)

	c.ReturnSuccess(poolIds)
}

func (c *ClusterApi) AllPoolList() {
	page := c.Query2Int("page", 1)
	pageSize := c.Query2Int("page_size", 10)

	c.CheckPage(&page, &pageSize)

	list := make([]models.Pool, 0, pageSize)

	count, err := service.Flow.ListByPageWithSort(page, pageSize, &models.Pool{}, &list, "-id")
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	liststruct := make([]pool_struct, len(list), pageSize)

	for i, fi := range list {
		liststruct[i].Id = fi.Id
		liststruct[i].Name = fi.Name
		liststruct[i].Desc = fi.Desc
		liststruct[i].VmType = fi.VmType
		liststruct[i].SdId = fi.SdId

		json.Unmarshal([]byte(fi.Tasks), &liststruct[i].Tasks)
		liststruct[i].ServiceId = fi.Service.Id

		nodeCount, err := service.Cluster.GetCount(&models.Node{}, "Pool", &models.Pool{Id: fi.Id})
		if err != nil {
			c.ReturnFailed(err.Error(), 400)
			return
		}
		liststruct[i].Nodecount = nodeCount

	}
	c.ReturnPageContent(page, pageSize, count, liststruct)
}
