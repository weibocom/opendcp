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
*  operation exec_task cronItem and dependItem
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
	Id           int       `json:"id"`
	ExecTaskId   int       `json:"exec_task_id"`              //依赖任务
	PoolId       int       `json:"pool_id"`
	Ratio        float64   `json:"ratio"`                     //依赖比例
	ElasticCount int       `json:"elastic_count"`             //冗余机器数量
	StepName     string    `json:"step_name"`                 //依赖步骤名称
	Ignore       bool      `json:"ignore"` 			  //是否忽略依赖 0 不忽略，1 忽略
}
type exec_task_struct struct{
	Id          int                   `json:"id"`
	PoolId       int                  `json:"pool_id"`         //服务池id
	CronItems   []task_cron_struct    `json:"cron_itmes"`  	   //定时任务列表
	DependItems []task_depend_struct  `json:"depend_itmes"`    //依赖任务列表
	Type        string                `json:"type"`            //模版任务类型 expand/upload
	ExecType    string                `json:"exec_type"`       //执行任务类型 crontab/depend
}

func (c *TaskApi) URLMapping() {
	c.Mapping("GetExpandList", c.GetExpandList)
	c.Mapping("GetUploadList", c.GetUploadList)
	c.Mapping("AddTask", c.AddTask)
}

/*
Get task depend item list
*/
func (c *TaskApi) GetExpandList() {
	beego.Debug("Begin enter GetExpandList")
	poolId := c.Ctx.Input.Param(":poolId")
	pool_id, err := strconv.Atoi(poolId)
	if err != nil {
		c.ReturnFailed("Bad pool id: " + poolId, 400)
		return
	}
	taskList, err := service.Task.GetAllTaskByPool(pool_id,"expand")
	c.ReturnSuccess(taskList)
}
/*
Get task cron item list
*/
func (c *TaskApi) GetUploadList() {
	beego.Debug("Begin enter GetUploadList")
	poolId := c.Ctx.Input.Param(":poolId")
	pool_id, err := strconv.Atoi(poolId)
	if err != nil {
		c.ReturnFailed("Bad pool id: " + poolId, 400)
		return
		return
	}
	taskList, err := service.Task.GetAllTaskByPool(pool_id,"upload")
	c.ReturnSuccess(taskList)
}
/*
add exec Task
*/
func (c *TaskApi) AddTask() {
	//req := struct {
	//	poolId              int                      `json:"pool_id"`
	//	taskType            string                   `json:"task_type"`
	//	cronList            []map[string]interface{} `json:"cronList"`
	//	depentList          []map[string]interface{} `json:"dependList"`
	//}{}
	var exe_taskList []exec_task_struct
	err := c.Body2Json(&exe_taskList)
	if err != nil {
		beego.Error("RUN Add Task, json err:", err)
		c.ReturnFailed(err.Error(), 400)
		return
	}
	if len(exe_taskList) != 2{
		beego.Error("exe_taskList matrix length have not two elements!")
		c.ReturnFailed(err.Error(), 400)
		return
	}
	pool_id := exe_taskList[0].PoolId
	Type := exe_taskList[0].Type

	cronList := exe_taskList[1].CronItems
	depentList := exe_taskList[0].DependItems
	if strings.EqualFold("crontab", exe_taskList[0].ExecType){
		cronList = exe_taskList[0].CronItems
		depentList = exe_taskList[1].DependItems
	}

	taskList, _:= service.Task.GetAllTaskByPool(pool_id, Type)

	var cronItemList []*models.CronItem
	var depentItemList []*models.DependItem
	if taskList == nil{
		//创建exec_task
		pool := &models.Pool{Id: pool_id}
		cronItems:=make([]*models.CronItem,0)
		dependItems:=make([]*models.DependItem,0)
		taskList, err = creatTask(pool, cronItems, dependItems, Type)
		cronItemList = cronItems
		depentItemList = dependItems
		if err != nil{
			beego.Error("create Task is failed: ", err)
			c.ReturnFailed(err.Error(), 400)
			return
		}
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
	for _, cron := range cronList {
		cronId := cron.Id
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
	//更新cron表
	cron_exec_task := taskList[1];
	if strings.EqualFold(taskList[0].ExecType,"crontab"){
		cron_exec_task = taskList[0];
	}
	creatCron(cronAddItemList, cron_exec_task)
	updateCron(cronUpdateItemList, cron_exec_task)
	deleteCron(cronDeleteItemList, cron_exec_task)

	//get update or add of depend
	var dependUpdateItemList []*models.DependItem
	var dependAddItemList []*models.DependItem
	for _, depend := range depentList {
		dependId := depend.Id
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
	//更新depend表
	depend_exec_task := taskList[1];
	if strings.EqualFold(taskList[0].ExecType,"depend"){
		depend_exec_task = taskList[0];
	}
	creatDepend(dependAddItemList, depend_exec_task)
	updateDepend(dependUpdateItemList, depend_exec_task)
	deleteDepend(dependDeleteItemList, depend_exec_task)



}

func getCronItemForm(cron task_cron_struct)(*models.CronItem, error){
	cron_item := &models.CronItem{}
	cron_item.ExecTask.Id = cron.ExecTaskId
	cron_item.Ignore = cron.Ignore
	cron_item.Time = cron.Time
	cron_item.WeekDay = cron.WeekDay
	cron_item.ConcurrRatio = cron.ConcurrRatio
	cron_item.ConcurrNum = cron.ConcurrNum
	cron_item.InstanceNum = cron.InstanceNum
	return cron_item,nil
}

func getDependItemForm(depend task_depend_struct)(*models.DependItem, error){
	depend_item := &models.DependItem{}
	depend_item.ExecTask.Id = depend.ExecTaskId
	depend_item.Ignore = depend_item.Ignore
	depend_item.Pool.Id =depend.PoolId
	depend_item.ElasticCount = depend.ElasticCount
	depend_item.Ratio = depend.Ratio
	depend_item.StepName = depend.StepName
	//该依赖的poolId中所有的任务没有依赖任务
	//需要判断，避免形成环
	return depend_item,nil
}

func creatTask(pool *models.Pool, cronItems []*models.CronItem, dependItems []*models.DependItem, task_type string,) ([]*models.ExecTask, error) {
	tasks := make([]*models.ExecTask,0)
	exec_task_cron:= &models.ExecTask{
		Pool: pool,
		CronItems: cronItems,
		DependItems: dependItems,
		Type:  task_type,
		ExecType: "crontab",
	}
	exec_task_depend:= &models.ExecTask{
		Pool: pool,
		CronItems: cronItems,
		DependItems: dependItems,
		Type:  task_type,
		ExecType: "depend",
	}
	err := service.Task.InsertBase(exec_task_cron)
	if err != nil{
		beego.Error("creat task_cron is failed")
		return nil, err
	}
	tasks = append(tasks, exec_task_cron)
	err = service.Task.InsertBase(exec_task_depend)
	if err != nil{
		beego.Error("creat task_depend is failed")
		return nil, err
	}
	tasks = append(tasks, exec_task_depend)
	return tasks, nil
}

func creatCron(task_cron_list []*models.CronItem, exec_task *models.ExecTask){
	for _, cron := range task_cron_list{
		cron.ExecTask = exec_task;
		err := service.Task.InsertBase(cron)
		if err != nil{
			beego.Warning("insert cron error already skip! " + err.Error())
			continue
		}
	}
}
func updateCron(task_cron_list []*models.CronItem, exec_task *models.ExecTask){
	for _, cron := range task_cron_list{
		cron.ExecTask = exec_task;
		err := service.Task.UpdateBase(cron)
		if err != nil{
			beego.Warning("insert cron error already skip! " + err.Error())
			continue
		}
	}
}
func deleteCron(task_cron_list []*models.CronItem, exec_task *models.ExecTask){
	for _, cron := range task_cron_list{
		cron.ExecTask = exec_task;
		err := service.Task.DeleteBase(cron)
		if err != nil{
			beego.Warning("insert cron error already skip! " + err.Error())
			continue
		}
	}
}
func creatDepend(task_dependList []*models.DependItem, exec_task *models.ExecTask){
	for _, depend := range task_dependList{
		depend.ExecTask = exec_task;
		err := service.Task.InsertBase(depend)
		if err != nil{
			beego.Warning("insert depend error already skip! " + err.Error())
			continue
		}
	}
}
func updateDepend(task_dependList []*models.DependItem, exec_task *models.ExecTask){
	for _, depend := range task_dependList{
		depend.ExecTask = exec_task;
		err := service.Task.UpdateBase(depend)
		if err != nil{
			beego.Warning("insert depend error already skip! " + err.Error())
			continue
		}
	}
}
func deleteDepend(task_dependList []*models.DependItem, exec_task *models.ExecTask){
	for _, depend := range task_dependList{
		depend.ExecTask = exec_task;
		err := service.Task.DeleteBase(depend)
		if err != nil{
			beego.Warning("insert depend error already skip! " + err.Error())
			continue
		}
	}
}
