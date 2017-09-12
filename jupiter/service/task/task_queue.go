package task

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"time"
	"weibo.com/opendcp/jupiter/dao"
	"weibo.com/opendcp/jupiter/future"
	"weibo.com/opendcp/jupiter/logstore"
	"weibo.com/opendcp/jupiter/models"
	"weibo.com/opendcp/jupiter/provider"
)

const (
	MAX_RUNNING_TASK = 1
	TASK_CACHE       = 5000
	INTERVAL_TASK    = 1  //second to run next expand machine task
)

var (
	taskService  = &InstanceTaskService{}
	instanceTask = &InstanceQueue{}
)

type InstanceQueue struct {
	isStarted bool
}

func (iq *InstanceQueue) Stop() {
	iq.isStarted = false
}

func (iq *InstanceQueue) Start() {
	beego.Info("The task of creating instance starting ...")
	iq.loop()
}

func (iq *InstanceQueue) loop() {
	iq.isStarted = true
	count := 1

	for {
		if !iq.isStarted {
			break
		}

		iq.run()

		if count%TASK_CACHE == 0 {
			iq.clearOldTasks()
			count = 0
		}
		count++
		time.Sleep(INTERVAL_TASK * time.Second)
	}
}

func (iq *InstanceQueue) run() {
	defer func() {
		if r := recover(); r != nil {
			beego.Error("Recovered from err when task run:", r)
		}
	}()

	initTask, err := taskService.GetTasks(models.StateReady)
	if err != nil {
		beego.Error("Failed to get tasks, err:", err)
		return
	}

	if len(initTask) == 0 {
		//beego.Info("There're no new tasks to run!")
		return
	}

	var tasks []models.InstanceItem
	if len(initTask) <= MAX_RUNNING_TASK {
		tasks = initTask
	} else {
		tasks = initTask[len(initTask)-MAX_RUNNING_TASK:]
	}

	iq.setTaskState(tasks, models.StateRunning)

	startTime := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf("--- Begin %d instance tasks at %s, %d left ---", len(tasks), startTime, len(initTask)-len(tasks))
	beego.Info(msg)
	go iq.createInstances(tasks)
}

func (iq *InstanceQueue) setTaskState(tasks []models.InstanceItem, state models.TaskState) {
	for _, task := range tasks {
		task.Status = state
		taskService.UpdateTask(task)
	}
}

func (iq *InstanceQueue) createInstances(tasks []models.InstanceItem) {
	//num := len(tasks)
	//success := make(chan *models.InstanceItem, num)
	//failed := make(chan *models.InstanceItem, num)
	//errs := make(chan string, num)

	for _, task := range tasks {
		go func(task models.InstanceItem) {
			ins, err := iq.createOneInstance(task)
			if err != nil {
				beego.Error("Task", task.TaskId, "executes failed,", "id:", task.Id)
				task.ErrLog = err.Error()
				task.Status = models.StateFailed
				taskService.UpdateTask(task)
				//failed <- &task
				//errs <- err.Error()
				//return
			}else {
				task.InstanceId = ins.InstanceId
				beego.Info("Task", task.TaskId, "executes successfully,", "id:", task.Id)
				task.Status = models.StateSuccess
				taskService.UpdateTask(task)
			}
			//success <- &task
		}(task)
	}

	//successCount := 0
	//failedCount := 0
	//for i := 0; i < num; i++ {
	//	select {
	//	case task := <-success:
	//		task.Status = models.StateSuccess
	//		taskService.UpdateTask(*task)
	//		successCount++
	//	case task := <-failed:
	//		task.Status = models.StateFailed
	//		err := <-errs
	//		task.ErrLog = err
	//		taskService.UpdateTask(*task)
	//		failedCount++
	//	}
	//}
	//
	//close(success)
	//close(failed)
	//close(errs)

	//beego.Info("Task completed, success:", successCount, "failed:", failedCount)
}

func (iq *InstanceQueue) createOneInstance(task models.InstanceItem) (*models.Instance, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			beego.Error("Recovered from err when create instance:", r)
			switch r.(type) {
			case error:
				err = r.(error)
			default:
				err = errors.New(fmt.Sprint("call create service failed:", r))
			}
		}
	}()

	ins, err := createCloudInstance(task.Cluster, task.CorrelationId)
	if err != nil {
		return nil, err
	}

	return ins, err
}

func (iq *InstanceQueue) clearOldTasks() {
	err := taskService.DeleteOldTasks(time.Now().UTC().AddDate(0, 0, -7))
	if err != nil {
		beego.Error("Delete old tasks err:", err)
	}
}

func createCloudInstance(cluster *models.Cluster, correlationId string) (*models.Instance, error) {
	providerDriver, err := provider.New(cluster.Provider)
	if err != nil {
		return nil, err
	}
	beego.Info("Begin to create instance from cloud")
	insIds, errs := providerDriver.Create(cluster, 1)
	if len(insIds) == 0 {
		return nil, errs[0]
	}
	logstore.Info(correlationId, insIds[0], "Begin to insert instance into db and start instance", insIds[0])
	logstore.Info(correlationId, insIds[0], "1. Begin to insert instance into db, instance id:", insIds[0])
	ins, err := providerDriver.GetInstance(insIds[0])
	if err != nil {
		logstore.Error(correlationId, insIds[0], "get instance info error:", err)
		return nil, err
	}
	logstore.Info(correlationId, insIds[0], "get instance info successfully")
	ins.Cluster = cluster
	ins.Cpu = cluster.Cpu
	ins.Ram = cluster.Ram
	ins.DataDiskCategory = cluster.DataDiskCategory
	ins.DataDiskSize = cluster.DataDiskSize
	ins.DataDiskNum = cluster.DataDiskNum
	ins.SystemDiskCategory = cluster.SystemDiskCategory
	ins.InstanceType = cluster.InstanceType
	ins.Status = models.Pending
	if err := dao.InsertInstance(ins); err != nil {
		logstore.Error(correlationId, insIds[0], "insert instance to db error:", err)
		return nil, err
	}
	logstore.Info(correlationId, insIds[0], "insert instance into db successfully")
	logstore.Info(correlationId, insIds[0], "2. Begin start instance in future")
	startFuture := future.NewStartFuture(insIds[0], cluster.Provider, true, ins.PrivateIpAddress, correlationId)
	go future.Exec.Submit(startFuture)

	return ins, nil
}

func InitInstanceTask() {
	go func() {
		instanceTask.Start()
	}()
}
