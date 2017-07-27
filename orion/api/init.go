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
	"fmt"
	"weibo.com/opendcp/orion/service"
	"weibo.com/opendcp/orion/utils"
	"weibo.com/opendcp/orion/handler"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
)

type InitApi struct {
	baseAPI
}

type ArrayStr []string


func (c *InitApi) URLMapping() {
	c.Mapping("Init", c.InitDB)
}

func (c *InitApi) InitDB() {
	biz := c.Ctx.Input.Header("X-Biz-ID")
	biz_id,err := strconv.Atoi(biz)
	if err !=nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}


	//1、调用jupiter接口，获取vm_id

	actionImpl := handler.GetActionImpl(biz_id,handler.RequestVMTypeId)

	h := handler.GetHandler(actionImpl.Type)

	result := h.HandleInit(actionImpl,nil)

	if result.Code == handler.CODE_ERROR {
		ret := fmt.Sprintf("request vm_type_id faild for biz %s",biz)
		c.ReturnFailed(ret, 400)
		return
	}

	vm_type_id,err := strconv.Atoi(result.Msg)
	if err !=nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}


	//2、调用hubble接口， 获取sid
	service_discovery_id := 9

	actionImpl = handler.GetActionImpl(biz_id,handler.REQUESTSID)

	h = handler.GetHandler(actionImpl.Type)
	result = h.HandleInit(actionImpl,nil)

	if result.Code == handler.CODE_ERROR {
		ret := fmt.Sprintf("request service_discovery_id faild for biz %s",biz)
		c.ReturnFailed(ret, 400)
		return
	}
	service_discovery_id,err = strconv.Atoi(result.Msg)
	if err !=nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}



	//3、初始化数据库
	utils.InitConf()

	content := make(map[string][]string,len(*utils.InitConfig))

	for i := 0 ; i<len(*utils.InitConfig);i++{
		tbName := (*utils.InitConfig)[i].Table

		records := (*utils.InitConfig)[i].Records

		content[tbName] = records
	}

	c.DeleteALl(biz_id,content)

	err = c.CreateRemoteAction(biz_id, content)

	if err !=nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	err = c.CreateRemoteStep(biz_id, content)

	if err !=nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	//err = c.CreateFlowImpl(vm_type_id,service_discovery_id,biz_id, content)
	//
	//if err !=nil {
	//	c.ReturnFailed(err.Error(), 400)
	//	return
	//}

	err = c.CreateCSP(vm_type_id,service_discovery_id,biz_id, content)



	if err !=nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}


	c.ReturnSuccess(true)

}

func (c *InitApi) DeleteALl (biz_id int, content map[string][]string) (err error) {
	deleteSql := "delete from %s where biz_id=%d"

	for table,_:= range content {
		service.Init.DeleteBysql(fmt.Sprintf(deleteSql,table,biz_id))
	}
	service.Init.DeleteBysql(fmt.Sprintf("delete from flow where biz_id=%d",biz_id))
	return nil

}

func (c *InitApi) CreateCSP(vm_type_id int, service_discovery_id int, biz_id int, content map[string][]string) (err error) {
	//0、创建模板
	flowImpls := content["flow_impl"]
	flowIds := make([]int,len(flowImpls))
	for index,sql := range flowImpls {
		if index >=0 && index <=1{
			sql = fmt.Sprintf(sql,vm_type_id,biz_id)
			sql = strings.Replace(sql,"host_ip",beego.AppConfig.String("octans_host"),-1)
		}else if index ==3 {
			sql = fmt.Sprintf(sql,vm_type_id,biz_id)
		}else if index ==4 {
			sql = fmt.Sprintf(sql,vm_type_id,biz_id)
		}else{
			sql = fmt.Sprintf(sql,biz_id)
		}

		id64, err := service.Init.InsertBysql(sql)
		if err != nil {
			return err
		}
		id := int(id64)
		flowIds[index] = id

	}
	//1、创建cluster
	clusters := content["cluster"]
	clusterId := -1
	serviceIds := make([]int,2)
	for _,sql := range clusters {
		sql = fmt.Sprintf(sql,"default",biz_id)
		id64, err := service.Init.InsertBysql(sql)
		if err != nil {
			return err
		}
		clusterId = int(id64)
	}
	//2、创建service
	services := content["service"]
	for index,sql := range services {
		sql = fmt.Sprintf(sql,clusterId)
		id64, err := service.Init.InsertBysql(sql)
		if err != nil {
			return err
		}
		serviceIds[index] = int(id64)
	}
	//3、创建pool
	pools := content["pool"]
	for index,sql := range pools {
		if index == 0 {
			sql = fmt.Sprintf(sql,vm_type_id,service_discovery_id,flowIds[3-1],flowIds[1-1],flowIds[2-1],serviceIds[index])
		}else if index ==1 {
			sql = fmt.Sprintf(sql,vm_type_id,service_discovery_id,flowIds[6-1],flowIds[4-1],flowIds[5-1],serviceIds[index])
		}
		_, err := service.Init.InsertBysql(sql)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *InitApi) CreateFlowImpl(vm_type_id int, service_discovery_id int, biz_id int, content map[string][]string) (err error) {
	flowImpls := content["flow_impl"]
	flowIds := make([]int,len(flowImpls))
	for index,sql := range flowImpls {
		if index >=0 && index <=3{
			sql = fmt.Sprintf(sql,vm_type_id,biz_id)
		}else if index ==4 {
			sql = fmt.Sprintf(sql,service_discovery_id,vm_type_id,biz_id)
		}else{
			sql = fmt.Sprintf(sql,biz_id)
		}

		id64, err := service.Init.InsertBysql(sql)
		if err != nil {
			return err
		}
		id := int(id64)
		flowIds[index] = id

	}

	return nil

}
func (c *InitApi) CreateRemoteStep(biz_id int, content map[string][]string) (err error) {

	steps := content["remote_step"]
	for _,sql := range steps {
		sql = fmt.Sprintf(sql,biz_id)
		_, err := service.Init.InsertBysql(sql)
		if err != nil {
			return err
		}
	}

	return nil

}


func (c *InitApi) CreateRemoteAction(biz_id int, content map[string][]string) (err error) {

	actions := content["remote_action"]

	actionImpls := content["remote_action_impl"]

	for index,sql := range actions {

		sql = fmt.Sprintf(sql,biz_id)

		id64, err := service.Init.InsertBysql(sql)
		if err != nil {
			return err
		}
		id := int(id64)


		//插入RemoteActionImpl
		sql = actionImpls[index]
		if index == 5{
			sql = strings.Replace(sql,"action_id_value",strconv.Itoa(id),-1)
			sql = strings.Replace(sql,"biz_id_value",strconv.Itoa(biz_id),-1)
		}else{
			sql = fmt.Sprintf(sql,id,biz_id)
		}


		fmt.Println(sql)

		id64, err = service.Init.InsertBysql(sql)
		if err != nil {
			return err
		}

	}

	return  nil


}


