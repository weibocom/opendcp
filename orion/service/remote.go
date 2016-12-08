package service

import (
	"github.com/astaxie/beego/orm"
	"weibo.com/opendcp/orion/models"
	"errors"
	"fmt"
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
	if(err != nil) {
		return err
	}

	//check Step
	o := orm.NewOrm()
	stepItem := &models.RemoteStep{}
	err = o.QueryTable(stepItem).Filter("actions__icontains","\""+objItem.Name+"\"").One(stepItem)

	if(len(stepItem.Name) > 0) {
		return errors.New(fmt.Sprintf("action is using ! step id:%v,step name:%v",stepItem.Id,stepItem.Name))
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
	if(err != nil) {
		return err
	}

	//check task
	orm.Debug = true
	o := orm.NewOrm()
	flowItem := &models.FlowImpl{}
	err = o.QueryTable(flowItem).Filter("steps__icontains","\"name\":\""+stepItem.Name+"\"").One(flowItem)

	if(len(flowItem.Name) > 0) {
		return errors.New(fmt.Sprintf("step is using ! template id:%v,template name:%v",flowItem.Id,flowItem.Name))
	}

	return f.DeleteBase(stepItem)
}
