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
	"weibo.com/opendcp/orion/models"
	"weibo.com/opendcp/orion/service"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
)
/**
*  operation related with action and step
 */
type TaskApi struct {
	baseAPI
}

type task_cron_struct struct{
	Id           int       `json:"id"`
	ExecTaskId   int       `json:"exec_task_id"`
	InstanceNum  int       `json:"instance_num"`              //扩容缩容使用作为机器的数量
	ConcurrRatio int       `json:"concurr_ratio"`             //上线使用作为最大并发比例
	ConcurrNum   int       `json:"concurr_num"`               //上线使用作为最大并发数
	WeekDay      int       `json:"week_day"`                  //每周第几天，取值 0 每天,1 周日,2 周一,3 周二,4 周三,5 周四,6 周五,7 周六
	Time         string    `json:"time"`                      //每天时分秒,例如 14:09:08
	Ignore       bool      `json:"ignore"`                    //是否忽略定时任务 0 不忽略，1 忽略
}
type task_depend_struct struct{

}
type exec_task_struct struct{

}


func (f *TaskApi) URLMapping() {
	f.Mapping("GetExpandList", f.GetExpandList)
	f.Mapping("GetCronList", f.GetCronList)
	f.Mapping("AddTask", f.AddTask)
}

/*
Get task depend item list
*/
func (c *TaskApi) GetExpandList() {
	poolId := c.Ctx.Input.Query("pool_id")
	pool_Id, err := strconv.Atoi(poolId)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}
	taskList, err := service.Task.GetAllTaskByPool(pool_Id,"expand")
	c.ReturnSuccess(taskList)
}
/*
Get task cron item list
*/
func (c *TaskApi) GetCronList() {
	poolId := c.Ctx.Input.Query("pool_id")
	pool_Id, err := strconv.Atoi(poolId)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}
	taskList, err := service.Task.GetAllTaskByPool(pool_Id,"upload")
	c.ReturnSuccess(taskList)
}
/*
add exec Task
*/
func (c *TaskApi) AddTask() {
	req := struct {
		poolId              int                      `json:"pool_id"`
		taskType            string                   `json:"task_type"`
		cronList            []map[string]interface{} `json:"cronList"`
		depentList          []map[string]interface{} `json:"dependList"`
	}{}
	err := c.Body2Json(&req)
	if err != nil {
		beego.Error("RUN Add Task, json err:", err)
		c.ReturnFailed(err.Error(), 400)
		return
	}
	taskList, _:= service.Task.GetAllTaskByPool(req.poolId,req.taskType)

	var cronItemList []*models.CronItem
	var depentItemList []*models.DependItem
	if taskList == nil{
		//创建exec_task
		pool := &models.Pool{Id: req.poolId}
		cronItems:=make([]*models.CronItem,0)
		dependItems:=make([]*models.DependItem,0)
		creatTask(pool, cronItems, dependItems, req.taskType, "crontab")
		creatTask(pool, cronItems, dependItems, req.taskType, "depend")
		cronItemList = cronItems
		depentItemList = dependItems
	}else{
		for _, task := range taskList {
			if strings.EqualFold("crontab", task.ExecType){
				cronItemList = task.CronItems
			}
			if strings.EqualFold("depend", task.ExecType){
				depentItemList = task.DependItems
			}
		}
	}

	//get update or add of cron
	cronUpdateItemList := make([]*models.CronItem, 0)
	cronAddItemList := make([]*models.CronItem, 0)
	for _, cron := range req.cronList {
		cronId, ok := cron["id"].(int)
		if !ok {
			beego.Error("cronId :[", cronId, "] has not id")
			continue
		}
		if cronId == 0{
			cron_item,_:= getCronItemForm(cron)
			cronAddItemList = append(cronAddItemList, cron_item)
		}else{
			cron_item,_:= getCronItemForm(cron)
			cronUpdateItemList = append(cronUpdateItemList, cron_item)
		}
	}
	//get need delete cron
	cronDeleteItemList := make([]*models.CronItem, 0)
	for _, cron := range cronItemList{
		isNeedDelete := true
		for _, updateCron := range cronUpdateItemList{
			if cron.Id == updateCron.Id{
				isNeedDelete = false
			}
		}
		if isNeedDelete{
			cronDeleteItemList = append(cronDeleteItemList, cron)
		}
	}
	//get update or add of depend
	var dependUpdateItemList []*models.DependItem
	var dependAddItemList []*models.DependItem
	for _, depend := range req.depentList {
		dependId, ok := depend["id"].(int)
		if !ok {
			beego.Error("cronId :[", depend, "] has not id")
			continue
		}
		if dependId == 0{
			depend_item,_:= getDependItemForm(depend)
			dependAddItemList = append(dependAddItemList, depend_item)
		}else{
			depend_item,_:= getDependItemForm(depend)
			dependUpdateItemList = append(dependUpdateItemList, depend_item)
		}
	}
	//get need delete depend
	dependDeleteItemList := make([]*models.DependItem, 0)
	for _, depend := range depentItemList{
		isNeedDelete := true
		for _, updateDepend := range dependUpdateItemList{
			if depend.Id == updateDepend.Id{
				isNeedDelete = false
			}
		}
		if isNeedDelete{
			dependDeleteItemList = append(dependDeleteItemList, depend)
		}
	}
	//更新cron表
	//创建cron

	//更新depend表




}

func getCronItemForm(cron map[string]interface{})(*models.CronItem, error){
	cron_item := &models.CronItem{}
	cron_item.ExecTask.Id = cron["exec_task_id"].(int)
	cron_item.Ignore = cron["ignore"].(bool)
	cron_item.Time = cron["time"].(string)
	cron_item.WeekDay = cron["weekDay"].(int)
	cron_item.ConcurrRatio = cron["concurrRatio"].(int)
	cron_item.ConcurrNum = cron["concurrNum"].(int)
	cron_item.InstanceNum = cron["instanceNum"].(int)
	return cron_item,nil
}

func getDependItemForm(depend map[string]interface{})(*models.DependItem, error){
	depend_item := &models.DependItem{}
	depend_item.ExecTask.Id = depend["exec_task_id"].(int)
	depend_item.Ignore = depend["ignore"].(bool)
	depend_item.Pool.Id =depend["pool_id"].(int)
	depend_item.ElasticCount = depend["elasticCount"].(int)
	depend_item.Ratio = depend["ratio"].(float64)
	depend_item.StepName = depend["stepName"].(string)
	//该依赖的poolId中所有的任务没有依赖任务
	//需要判断，避免形成环
	return depend_item,nil
}

func creatTask(pool *models.Pool, cronItems []*models.CronItem, dependItems []*models.DependItem, task_type string, task_exec_type string) (*models.ExecTask, error) {
	exec_task:= &models.ExecTask{
		Pool: pool,
		CronItems: cronItems,
		DependItems: dependItems,
		Type:  task_type,
		ExecType: task_exec_type,
	}
	err := service.Task.InsertBase(exec_task)
	return exec_task, err
}

//func creatCron(task_cron_list []*models.CronItem){
//	for _, cron := range task_cron_list{
//		isNeedDelete := true
//		for _, updateCron := range cronUpdateItemList{
//			if cron.Id == updateCron.Id{
//				isNeedDelete = false
//			}
//		}
//		if isNeedDelete{
//			cronDeleteItemList = append(cronDeleteItemList, cron)
//		}
//	}
//}
func creatDepend(task_dependList []*models.DependItem){

}