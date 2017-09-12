package task

import (
	"github.com/astaxie/beego"
	"time"
	"weibo.com/opendcp/jupiter/dao"
	"weibo.com/opendcp/jupiter/models"
)

const (
	WAIT_AGAIN_TIMES = 70 //AWS创建机器时间比较长，此处设置较大的值
	TIME_INTERVAL    = 3
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
	num := len(tasks)
	for i := 0; i < num+WAIT_AGAIN_TIMES; i++ { //等待所有instance获取到instanceId
		//allDone := true
		beego.Debug("Wait task complete, times:", i+1)
		instanceCount := 0;
		for index, task := range tasks {
			//if task.Status == models.StateSuccess || task.Status == models.StateFailed {
			//	continue
			//}

			taskItem, err := dao.GetItemById(task.Id)
			if err != nil {
				beego.Error("Get task ", task.TaskId, "failed,id:", task.Id, ", err:", err)
				continue
			}

			tasks[index].Status = taskItem.Status
			tasks[index].InstanceId = taskItem.InstanceId

			if len(taskItem.InstanceId) != 0{
				instanceCount++;
				//allDone = false
			}
			//if task.Status == models.StateSuccess || task.Status == models.StateFailed {
			//	beego.Debug("Task", task.TaskId, "finished id:", task.Id, "status:", taskItem.Status)
			//	continue
			//}

			//if task.Status != models.StateSuccess || task.Status != models.StateFailed {
			//	allDone = false
			//}

		}

		if instanceCount == len(tasks) {
			beego.Debug("Tasks have completed")
			break
		} else {
			time.Sleep(TIME_INTERVAL * time.Second)
		}
	}

	return nil
}
