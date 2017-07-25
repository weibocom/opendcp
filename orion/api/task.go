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
	"github.com/astaxie/beego"
	"strconv"
	"weibo.com/opendcp/orion/models"
	"weibo.com/opendcp/orion/sched"
	"weibo.com/opendcp/orion/service"
)

/**
*  operation exec_task cronItem and dependItem
 */
type TaskApi struct {
	baseAPI
}

type task_cron_struct struct {
	Id           int    `json:"id"`
	ExecTaskId   int    `json:"exec_task_id"`
	InstanceNum  int    `json:"instance_num"`  //扩容缩容使用作为机器的数量
	ConcurrRatio int    `json:"concurr_ratio"` //上线使用作为最大并发比例
	ConcurrNum   int    `json:"concurr_num"`   //上线使用作为最大并发数
	WeekDay      int    `json:"week_day"`      //每周第几天，取值 0 每天,1 周日,2 周一,3 周二,4 周三,5 周四,6 周五,7 周六
	Time         string `json:"time"`          //每天时分秒,例如 14:09:08
	Ignore       int    `json:"ignore"`        //是否忽略定时任务 0 不忽略，1 忽略
}
type task_depend_struct struct {
	Id           int     `json:"id"`
	ExecTaskId   int     `json:"exec_task_id"` //依赖任务
	PoolId       int     `json:"pool_id"`
	Ratio        float64 `json:"ratio"`         //依赖比例
	ElasticCount int     `json:"elastic_count"` //冗余机器数量
	StepName     string  `json:"step_name"`     //依赖步骤名称
	Ignore       int     `json:"ignore"`        //是否忽略依赖 0 不忽略，1 忽略
}
type exec_task_struct struct {
	Id          int                  `json:"id"`
	PoolId      int                  `json:"pool_id"`      //服务池id
	CronItems   []task_cron_struct   `json:"cron_itmes"`   //定时任务列表
	DependItems []task_depend_struct `json:"depend_itmes"` //依赖任务列表
	Type        string               `json:"type"`         //模版任务类型 expand/upload
	ExecType    string               `json:"exec_type"`    //执行任务类型 crontab/depend
}

func (c *TaskApi) URLMapping() {
	c.Mapping("GetExpandList", c.GetExpandList)
	c.Mapping("GetUploadList", c.GetUploadList)
	c.Mapping("SaveTask", c.SaveTask)
}

/*
Get task depend item list
*/
func (c *TaskApi) GetExpandList() {
	beego.Debug("Begin enter GetExpandList")
	poolId := c.Ctx.Input.Param(":poolId")
	pool_id, err := strconv.Atoi(poolId)
	if err != nil {
		c.ReturnFailed("Bad pool id: "+poolId, 400)
		return
	}
	taskList, err := service.Task.GetAllTaskByPool(pool_id, models.TaskExpend)
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
		c.ReturnFailed("Bad pool id: "+poolId, 400)
		return
		return
	}
	taskList, err := service.Task.GetAllTaskByPool(pool_id, models.TaskDdeploy)
	c.ReturnSuccess(taskList)
}

/*
add exec Task
CronItem add check for time unique
DependItem add chekc for depend pool unique
*/

func (c *TaskApi) SaveTask() {
	beego.Info(".....SaveTask.....")
	save_exe_task := exec_task_struct{}
	err := c.Body2Json(&save_exe_task)
	if err != nil {
		beego.Error("RUN Add Task, parse json err:", err)
		c.ReturnFailed(err.Error(), 400)
		return
	}

	exec_task := &models.ExecTask{}
	exec_task.Id = save_exe_task.Id

	pool := &models.Pool{Id: save_exe_task.PoolId}
	err = service.Cluster.GetBase(pool)
	if err != nil {
		c.ReturnFailed(err.Error(), 500)
	}
	exec_task.Pool = pool

	exec_task.Type = save_exe_task.Type
	exec_task.ExecType = save_exe_task.ExecType

	cronItems := make([]*models.CronItem, 0)
	dependItems := make([]*models.DependItem, 0)

	timeMap := make(map[string]bool)
	dependPoolMap := make(map[string]bool)

	for _, cron := range save_exe_task.CronItems {
		cronItem := &models.CronItem{}
		cronItem.Id = cron.Id
		cronItem.ExecTask = exec_task
		if _, ok := timeMap[cron.Time]; ok {
			c.ReturnFailed("Cron time "+cron.Time+" is duplicate!", 500)
		} else {
			timeMap[cron.Time] = true
		}
		cronItem.Time = cron.Time
		cronItem.ConcurrNum = cron.ConcurrNum
		cronItem.ConcurrRatio = cron.ConcurrRatio
		if cron.Ignore == 0 {
			cronItem.Ignore = false
		} else {
			cronItem.Ignore = true
		}

		cronItem.InstanceNum = cron.InstanceNum
		cronItem.WeekDay = cron.WeekDay

		cronItems = append(cronItems, cronItem)
	}

	for _, depend := range save_exe_task.DependItems {
		dependItem := &models.DependItem{}
		dependItem.Id = depend.Id
		if depend.Ignore == 0 {
			dependItem.Ignore = false
		} else {
			dependItem.Ignore = true
		}
		pool := &models.Pool{Id: depend.PoolId}
		err = service.Cluster.GetBase(pool)
		if err != nil {
			c.ReturnFailed(err.Error(), 500)
		}
		if _, ok := dependPoolMap[pool.Name]; ok {
			c.ReturnFailed("Depend pool "+pool.Name+" is duplicate!", 500)
		} else {
			dependPoolMap[pool.Name] = true
		}
		dependItem.Pool = pool
		dependItem.Ratio = depend.Ratio
		dependItem.ElasticCount = depend.ElasticCount
		dependItem.StepName = depend.StepName
		dependItem.ExecTask = exec_task
		dependItems = append(dependItems, dependItem)
	}
	exec_task.CronItems = cronItems
	exec_task.DependItems = dependItems

	if exec_task.Id != 0 {
		err := sched.Scheduler.Update(exec_task)
		if err != nil {
			c.ReturnFailed(err.Error(), 500)
		}
	} else {
		err := sched.Scheduler.Create(exec_task)
		if err != nil {
			c.ReturnFailed(err.Error(), 500)
		}
	}
	c.ReturnSuccess(true)
}

/*func (c *TaskApi) SaveTask() {
	beego.Info(".....SaveTask.....")
	save_exe_task := exec_ask_struct{}
	err := c.Body2Json(&save_exe_task)
	if err != nil {
		beego.Error("RUN Add Task, json err:", err)
		c.ReturnFailed(err.Error(), 400)
		return
	}
	pool_id := save_exe_task.PoolId
	Type := save_exe_task.Type
	exe_Type := save_exe_task.ExecType
	beego.Info("pool_id: ", pool_id, " Type: ", Type, " exe_Type:", exe_Type)
	cronList := save_exe_task.CronItems
	depentList := save_exe_task.DependItems

	beego.Info("cronList length: ", len(cronList), " depentList length: ", len(depentList))

	taskList, _:= service.Task.GetAllTaskByPool(pool_id, Type)

	var cronItemList []*models.CronItem
	var depentItemList []*models.DependItem
	if taskList == nil{
		//创建exec_task
		beego.Info("create exec_task!")
		pool := &models.Pool{Id: pool_id}
		cronItems:=make([]*models.CronItem,0)
		dependItems:=make([]*models.DependItem,0)
		taskList, err = creatTask(pool, cronItems, dependItems, Type)
		if err != nil{
			beego.Error("create Task is failed: ", err)
			c.ReturnFailed(err.Error(), 400)
			return
		}
		err :=sched.Scheduler.Create(taskList)
		if err != nil{
			c.ReturnFailed(err.Error(), 500)
		}
		cronItemList = cronItems
		depentItemList = dependItems
	}else{
		updateTask(taskList, save_exe_task)
		beego.Info("exec_task id :", taskList.Id)
		cronItemList = taskList.CronItems
		depentItemList = taskList.DependItems
	}
	beego.Info("get cronItemList length:",  len(cronItemList))
	beego.Info("get depentItemList length:",  len(depentItemList))
	//get update or add of cron
	beego.Info("get update or add of cron")
	cronUpdateItemList := make([]*models.CronItem, 0)
	cronAddItemList := make([]*models.CronItem, 0)
	for _, cron := range cronList {
		cronId := cron.Id
		beego.Debug("get cronList cron id: ", cronId)
		if cronId == 0{
			cron_item,_:= getCronItemForm(cron)
			cronAddItemList = append(cronAddItemList, cron_item)
		}else{
			cron_item,_:= getCronItemForm(cron)
			cronUpdateItemList = append(cronUpdateItemList, cron_item)
		}
	}
	beego.Info("get update cron length:",  len(cronUpdateItemList), " get add cron length:", len(cronAddItemList))
	//get need delete cron
	cronDeleteItemList := make([]*models.CronItem, 0)
	for _, cron := range cronItemList{
		isNeedDelete := true
		for _, updateCron := range cronUpdateItemList{
			if cron.Id == updateCron.Id{
				beego.Debug("delete cronList cron id: ", cron.Id , updateCron.Id)
				isNeedDelete = false
				break
			}
		}
		beego.Debug("isNeedDelete: ", isNeedDelete)
		if isNeedDelete{
			cronDeleteItemList = append(cronDeleteItemList, cron)
		}
	}
	beego.Info("get delete cron length:",  len(cronDeleteItemList))
	//更新cron表
	creatCron(cronAddItemList, taskList)
	updateCron(cronUpdateItemList, taskList)
	deleteCron(cronDeleteItemList, taskList)

	//get update or add of depend
	beego.Info("get update or add of depend")
	var dependUpdateItemList []*models.DependItem
	var dependAddItemList []*models.DependItem
	for _, depend := range depentList {
		dependId := depend.Id
		beego.Info("depentList dependId: ", dependId)
		if dependId == 0{
			depend_item,_:= getDependItemForm(depend)
			dependAddItemList = append(dependAddItemList, depend_item)
		}else{
			depend_item,_:= getDependItemForm(depend)
			dependUpdateItemList = append(dependUpdateItemList, depend_item)
		}
	}
	beego.Info("get update depend length:",  len(dependUpdateItemList), " get add depend length:", len(dependAddItemList))
	//get need delete depend
	dependDeleteItemList := make([]*models.DependItem, 0)
	for _, depend := range depentItemList{
		isNeedDelete := true
		beego.Debug("depentItemList depend id: ", depend.Id)
		for _, updateDepend := range dependUpdateItemList{
			if depend.Id == updateDepend.Id{
				beego.Debug("delete dependList cron id: ", depend.Id , updateDepend.Id)
				isNeedDelete = false
				break
			}
		}
		if isNeedDelete{
			dependDeleteItemList = append(dependDeleteItemList, depend)
		}
	}
	beego.Info("get delete depend length:",  len(dependDeleteItemList))
	//更新depend表
	creatDepend(dependAddItemList, taskList)
	updateDepend(dependUpdateItemList, taskList)
	deleteDepend(dependDeleteItemList, taskList)

	//更新内存
	taskList, _= service.Task.GetAllTaskByPool(pool_id, Type)
	err = sched.Scheduler.Update(taskList)
	if err != nil{
		c.ReturnFailed(err.Error(), 500)
	}
	c.ReturnSuccess(nil)
}*/

func getCronItemForm(cron task_cron_struct) (*models.CronItem, error) {
	beego.Info("getCronItemForm")
	beego.Info("cron.ExecTaskId ", cron.ExecTaskId)
	cron_item := &models.CronItem{}
	cron_item.Id = cron.Id
	exe_Task := &models.ExecTask{Id: cron.ExecTaskId}
	cron_item.ExecTask = exe_Task
	beego.Info("cron.ExecTaskId ", cron.ExecTaskId)
	if cron.Ignore == 0 {
		cron_item.Ignore = false
	} else {
		cron_item.Ignore = true
	}
	beego.Info("cron.Ignore ", cron_item.Ignore)
	cron_item.Time = cron.Time
	beego.Info("cron_item.Time ", cron_item.Time)
	cron_item.WeekDay = cron.WeekDay
	beego.Info("cron_item.WeekDay ", cron_item.WeekDay)
	cron_item.ConcurrRatio = cron.ConcurrRatio
	beego.Info("cron_item.ConcurrRatio  ", cron_item.ConcurrRatio)
	cron_item.ConcurrNum = cron.ConcurrNum
	beego.Info("cron_item.ConcurrNum ", cron_item.ConcurrNum)
	cron_item.InstanceNum = cron.InstanceNum
	beego.Info("cron_item.InstanceNum  ", cron_item.InstanceNum)
	return cron_item, nil
}

func getDependItemForm(depend task_depend_struct) (*models.DependItem, error) {
	beego.Info("getDependItemForm")
	depend_item := &models.DependItem{}
	beego.Info("depend.Id: ", depend.Id)
	depend_item.Id = depend.Id
	beego.Info("depend.ExecTaskId: ", depend.ExecTaskId)
	exe_Task := &models.ExecTask{Id: depend.ExecTaskId}
	depend_item.ExecTask = exe_Task
	if depend.Ignore == 0 {
		depend_item.Ignore = false
	} else {
		depend_item.Ignore = true
	}
	beego.Info("depend_item.Ignore: ", depend_item.Ignore)
	depend_pool := &models.Pool{Id: depend.PoolId}
	beego.Info("depend_item.Ignore: ", depend.PoolId)
	depend_item.Pool = depend_pool
	depend_item.ElasticCount = depend.ElasticCount
	beego.Info("depend.ElasticCount: ", depend.ElasticCount)
	depend_item.Ratio = depend.Ratio
	beego.Info("depend.Ratio: ", depend.Ratio)
	depend_item.StepName = depend.StepName
	beego.Info("depend.StepName: ", depend.StepName)
	//该依赖的poolId中所有的任务没有依赖任务
	//需要判断，避免形成环
	return depend_item, nil
}

func updateTask(dbtask *models.ExecTask, save_exe_task exec_task_struct) error {
	dbtask.Type = save_exe_task.Type
	dbtask.ExecType = save_exe_task.ExecType
	err := service.Task.UpdateBase(dbtask)
	return err
}
func creatTask(pool *models.Pool, cronItems []*models.CronItem, dependItems []*models.DependItem, task_type string) (*models.ExecTask, error) {
	exec_task := &models.ExecTask{
		Pool:        pool,
		CronItems:   cronItems,
		DependItems: dependItems,
		Type:        task_type,
		ExecType:    models.Crontab,
	}
	err := service.Task.InsertBase(exec_task)
	if err != nil {
		beego.Error("creat task_cron is failed")
		return nil, err
	}
	return exec_task, nil
}

func creatCron(task_cron_list []*models.CronItem, exec_task *models.ExecTask) {
	for _, cron := range task_cron_list {
		cron.ExecTask = exec_task
		err := service.Task.InsertBase(cron)
		if err != nil {
			beego.Warning("insert cron error already skip! " + err.Error())
			continue
		}
	}
}
func updateCron(task_cron_list []*models.CronItem, exec_task *models.ExecTask) {
	for _, cron := range task_cron_list {
		cron.ExecTask = exec_task
		err := service.Task.UpdateBase(cron)
		if err != nil {
			beego.Warning("insert cron error already skip! " + err.Error())
			continue
		}

	}
}
func deleteCron(task_cron_list []*models.CronItem, exec_task *models.ExecTask) {
	for _, cron := range task_cron_list {
		cron.ExecTask = exec_task
		err := service.Task.DeleteBase(cron)
		if err != nil {
			beego.Warning("insert cron error already skip! " + err.Error())
			continue
		}
	}
}
func creatDepend(task_dependList []*models.DependItem, exec_task *models.ExecTask) {
	for _, depend := range task_dependList {
		depend.ExecTask = exec_task
		err := service.Task.InsertBase(depend)
		if err != nil {
			beego.Warning("insert depend error already skip! " + err.Error())
			continue
		}
	}
}
func updateDepend(task_dependList []*models.DependItem, exec_task *models.ExecTask) {
	for _, depend := range task_dependList {
		depend.ExecTask = exec_task
		err := service.Task.UpdateBase(depend)
		if err != nil {
			beego.Warning("insert depend error already skip! " + err.Error())
			continue
		}
	}
}
func deleteDepend(task_dependList []*models.DependItem, exec_task *models.ExecTask) {
	for _, depend := range task_dependList {
		depend.ExecTask = exec_task
		err := service.Task.DeleteBase(depend)
		if err != nil {
			beego.Warning("insert depend error already skip! " + err.Error())
			continue
		}
	}
}
