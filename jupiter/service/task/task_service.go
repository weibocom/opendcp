package task

import (
	"time"
	"weibo.com/opendcp/jupiter/dao"
	"weibo.com/opendcp/jupiter/models"
	"github.com/astaxie/beego"
)

const (
	waitTimes    = 10
	timeInterVal = 5
)

type InstanceTaskService struct {
}

func (its *InstanceTaskService) CreateTask(task models.InstanceItem) error {
	err := dao.InsertItem(task)
	if err != nil {
		return err
	}
	return nil
}

func (its *InstanceTaskService) CreateTasks(tasks []models.InstanceItem) error {
	err := dao.InsertItems(tasks)
	if err != nil {
		return err
	}
	return nil
}

func (its *InstanceTaskService) DeleteTask(task models.InstanceItem) error {
	err := dao.DeleteItem(task)
	if err != nil {
		return err
	}
	return nil
}

func (its *InstanceTaskService) DeleteOldTasks(before time.Time) error {
	err := dao.DeleteOldItems(before)
	if err != nil {
		return err
	}
	return nil
}

func (its *InstanceTaskService) UpdateTask(task models.InstanceItem) error {
	err := dao.UpdateItem(task)
	if err != nil {
		return err
	}
	return err
}

func (its *InstanceTaskService) GetTasks(status models.TaskState) ([]models.InstanceItem, error) {
	tasks, err := dao.GetItemsByStatus(int(status))
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (its *InstanceTaskService) WaitTasksComplete(tasks []models.InstanceItem) error {
	for i := 0; i < waitTimes; i++ {
		allDone := true
		beego.Warn(">>>>>>>>>>>>>>>>>>Wait task times:", i+1)
		for index, task := range tasks {
			if task.Status == models.StateSuccess || task.Status == models.StateFailed {
				continue
			}

			taskItem, err := dao.GetItemById(task.Id)
			if err != nil {
				beego.Error("Get task ", task.TaskId, "failed,id:",task.Id,", err:",err)
				continue
			}

			tasks[index].Status = taskItem.Status

			if task.Status == models.StateSuccess || task.Status == models.StateFailed {
				beego.Debug("Task", task.TaskId, "finished id:", task.Id, "status:", taskItem.Status)
				continue
			}

			allDone = false
		}

		if allDone {
			break
		} else {
			time.Sleep(timeInterVal * time.Second)
		}
	}

	return nil
}
