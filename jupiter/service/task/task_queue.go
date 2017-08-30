package task

import (
	"github.com/astaxie/beego"
	"time"
	"weibo.com/opendcp/jupiter/models"
)

const MAX_RUNNING_NUM = 1

var taskService = &InstanceTaskService{}

type InstanceQueue struct {
	isStarted bool
}

func (iq *InstanceQueue) Stop() {
	iq.isStarted = false
}

func (iq *InstanceQueue) Start() {
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

		if count%10000 == 0 {
			iq.clearOldTasks()
			count = 0
		}
		count++
		time.Sleep(1 * time.Second)
	}
}

func (iq *InstanceQueue) run() {
	defer func() {
		if r := recover(); r != nil {
			beego.Error("Recovered from err:", r)
		}
	}()

	initTask, err := taskService.GetTasks(models.StateReady)
	if err != nil {
		beego.Error("Failed to get tasks, err:", err)
		return
	}

	if len(initTask) == 0 {
		beego.Info("There're no tasks to run")
		return
	}

	var tasks []models.InstanceItem
	if len(initTask) <= MAX_RUNNING_NUM {
		tasks = initTask
	} else {
		tasks = initTask[len(initTask)-MAX_RUNNING_NUM:]
	}

	iq.setTaskState(tasks, models.StateRunning)
	iq.createInstances(tasks)
}

func (iq *InstanceQueue) setTaskState(tasks []models.InstanceItem, state models.TaskState) {
	for _, task := range tasks {
		task.Status = state
		taskService.UpdateTask(task)
	}
}

func (iq *InstanceQueue) createInstances(tasks []models.InstanceItem) {

}

func (iq *InstanceQueue) createOneInstance(task models.InstanceItem) (models.Instance, error){

}

func (iq *InstanceQueue) clearOldTasks() {
	err := taskService.DeleteOldTasks(time.Now().UTC().AddDate(0, 0, -7))
	if err != nil {
		beego.Error("Delete old tasks err:", err)
	}
}
