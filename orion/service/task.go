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

package service

import (
	"github.com/astaxie/beego/orm"

	"weibo.com/opendcp/orion/models"
	"fmt"
)

type TaskService struct {
	BaseService
}

func (t *TaskService) GetAllTaskByPool(pool_id int, task_type string) (*models.ExecTask, error) {
	o := orm.NewOrm()
	task := &models.ExecTask{}
	var err error
	if task_type == "" || task_type == " " || task_type == "all" {
		err = o.QueryTable("exec_task").Filter("Pool", pool_id).RelatedSel().One(task)
	} else {
		err = o.QueryTable("exec_task").Filter("Pool", pool_id).Filter("Type", task_type).RelatedSel().One(task)
	}

	if err != nil {
		return nil, err
	}


	o.LoadRelated(task, "CronItems")
	o.LoadRelated(task, "DependItems")

	for _, dep := range task.DependItems {
		if err := o.Read(dep.Pool); err != nil {
			return nil, fmt.Errorf("db load %d DependItems Pool failed: %v", err)
		}
	}

	return task, nil
}
