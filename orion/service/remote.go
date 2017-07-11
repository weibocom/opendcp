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
	"fmt"

	"github.com/astaxie/beego/orm"
	"weibo.com/opendcp/orion/models"
)

type RemoteStepService struct {
	BaseService
}

/*
type RemoteActionService struct {
	BaseService
}

type RemoteActionImplService struct {
	BaseService
}
*/

var (
	remoteStepService = &RemoteStepService{}
	//remoteActionService = &RemoteActionService{}
	//remoteActionImplService = &RemoteActionImplService{}
)

/**
*  load all actions from db
 */
func (f *RemoteStepService) ActionDelete(id int) error {
	o := orm.NewOrm()

	action := &models.RemoteAction{}

	_, err := o.QueryTable(action).Filter("id", id).Delete()

	return err
}

/**
*  check step then delete actions from db
*
*  if step is using.. return error
 */
func (f *RemoteStepService) CheckActionDelete(id int) error {
	objItem := &models.RemoteAction{Id: id}
	err := f.GetBase(objItem)
	if err != nil {
		return err
	}

	//check Step
	o := orm.NewOrm()
	stepItem := &models.RemoteStep{}
	err = o.QueryTable(stepItem).Filter("actions__icontains", "\""+objItem.Name+"\"").One(stepItem)

	if len(stepItem.Name) > 0 {
		return fmt.Errorf("action is using ! step id:%v,step name:%v", stepItem.Id, stepItem.Name)
	}

	return f.ActionDelete(id)
}

/**
*  check template then delete step from db
*
*  if step is using.. return error
 */
func (f *RemoteStepService) CheckStepDelete(id int) error {
	stepItem := &models.RemoteStep{Id: id}
	err := f.GetBase(stepItem)
	if err != nil {
		return err
	}

	//check task
	orm.Debug = true
	o := orm.NewOrm()
	flowItem := &models.FlowImpl{}
	err = o.QueryTable(flowItem).Filter("steps__icontains", "\"name\":\""+stepItem.Name+"\"").One(flowItem)

	if len(flowItem.Name) > 0 {
		return fmt.Errorf("step is using ! template id:%v,template name:%v", flowItem.Id, flowItem.Name)
	}

	return f.DeleteBase(stepItem)
}
