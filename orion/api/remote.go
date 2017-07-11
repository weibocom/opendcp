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
	"encoding/json"
	"strconv"
	"unicode/utf8"

	"weibo.com/opendcp/orion/models"
	"weibo.com/opendcp/orion/service"
)

/**
*  operation related with action and step
 */
type RemoteApi struct {
	baseAPI
}

const (
	ANSIBLE = "ansible"
)

type action_struct struct {
	Id     int                    `json:"id"`
	Name   string                 `json:"name"`
	Desc   string                 `json:"desc"`
	Params map[string]interface{} `json:"params"`
}

type remotestep_struct struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Desc    string   `json:"desc"`
	Actions []string `json:"actions"`
}

type remoteactionimpl_struct struct {
	Id       int                    `json:"id"`
	Type     string                 `json:"type"`
	ActionId int                    `json:"action_id"`
	Template map[string]interface{} `json:"template"`
}

func (f *RemoteApi) URLMapping() {

	f.Mapping("ActionList", f.ActionList)
	f.Mapping("ActionDelete", f.ActionDelete)
	f.Mapping("GetAction", f.GetAction)
	f.Mapping("ActionUpdate", f.ActionUpdate)
	f.Mapping("ActionAppend", f.ActionAppend)

	f.Mapping("RemoteStepList", f.RemoteStepList)
	f.Mapping("RemoteStepDelete", f.RemoteStepDelete)
	f.Mapping("GetRemoteStep", f.GetRemoteStep)
	f.Mapping("RemoteStepUpdate", f.RemoteStepUpdate)
	f.Mapping("RemoteStepAppend", f.RemoteStepAppend)

	f.Mapping("RemoteActionImplAppend", f.RemoteActionImplAppend)
	f.Mapping("GetRemoteActionImpl", f.GetRemoteActionImpl)
	f.Mapping("RemoteActionImplDelete", f.RemoteActionImplDelete)
	f.Mapping("RemoteActionImplList", f.RemoteActionImplList)
	f.Mapping("RemoteActionImplUpdate", f.RemoteActionImplUpdate)

}

/**
*  create new action
 */
func (c *RemoteApi) ActionAppend() {

	req := action_struct{}

	err := c.Body2Json(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	paramsStr, err := json.Marshal(req.Params)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	// check unicode
	if len(string(paramsStr)) != utf8.RuneCountInString(string(paramsStr)) {
		c.ReturnFailed("there is unicode character in param", 400)
		return
	}

	data := models.RemoteAction{
		Name:   req.Name,
		Desc:   req.Desc,
		Params: string(paramsStr),
	}
	err = service.Cluster.InsertBase(&data)

	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}
	c.ReturnSuccess(data.Id)
}

/**
*  load actions by page
 */
func (c *RemoteApi) ActionList() {

	page := c.Query2Int("page", 1)
	pageSize := c.Query2Int("page_size", 10)

	c.CheckPage(&page, &pageSize)

	list := make([]models.RemoteAction, 0, pageSize)

	count, err := service.Remote.ListByPageWithSort(page, pageSize, &models.RemoteAction{}, &list, "-id")
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	liststruct := make([]action_struct, len(list), pageSize)

	for i, fi := range list {
		liststruct[i].Id = fi.Id
		liststruct[i].Name = fi.Name
		liststruct[i].Desc = fi.Desc
		json.Unmarshal([]byte(fi.Params), &liststruct[i].Params)

	}

	c.ReturnPageContent(page, pageSize, count, liststruct)
}

/**
*  load action by id
 */
func (c *RemoteApi) GetAction() {

	id := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	action := &models.RemoteAction{Id: idInt}
	err := service.Remote.GetBase(action)

	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}

	actionstr := action_struct{}
	actionstr.Id = action.Id
	actionstr.Name = action.Name
	actionstr.Desc = action.Desc
	json.Unmarshal([]byte(action.Params), &actionstr.Params)

	c.ReturnSuccess(actionstr)
}

/**
* update single action by id
 */
func (c *RemoteApi) ActionUpdate() {

	id := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	req := struct {
		Name   string                 `json:"name"`
		Desc   string                 `json:"desc"`
		Params map[string]interface{} `json:"params"`
	}{}

	err := c.Body2Json(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	action := &models.RemoteAction{Id: idInt}
	err = service.Remote.GetBase(action)
	action.Desc = req.Desc

	paramsStr, _ := json.Marshal(req.Params)

	action.Params = string(paramsStr)

	err = service.Remote.UpdateBase(action)

	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}
	c.ReturnSuccess("")
}

/**
* remove single action by id
 */
func (c *RemoteApi) ActionDelete() {

	id := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	err := service.Remote.CheckActionDelete(idInt)
	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}
	c.ReturnSuccess("")
}

/**
*  create new RemoteStep
 */
func (c *RemoteApi) RemoteStepAppend() {
	req := remotestep_struct{}

	err := c.Body2Json(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	actionsStr, _ := json.Marshal(req.Actions)

	data := models.RemoteStep{
		Name:    req.Name,
		Desc:    req.Desc,
		Actions: string(actionsStr),
	}
	err = service.Cluster.InsertBase(&data)

	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}
	c.ReturnSuccess(data.Id)
}

/**
* update single RemoteStep update by id
 */
func (c *RemoteApi) RemoteStepUpdate() {

	id := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	req := struct {
		Name    string   `json:"name"`
		Desc    string   `json:"desc"`
		Actions []string `json:"actions"`
	}{}

	err := c.Body2Json(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	remotestep := &models.RemoteStep{Id: idInt}
	err = service.Remote.GetBase(remotestep)
	remotestep.Desc = req.Desc

	paramsStr, _ := json.Marshal(req.Actions)

	remotestep.Actions = string(paramsStr)

	err = service.Remote.UpdateBase(remotestep)

	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}
	c.ReturnSuccess("")
}

/**
*  load RemoteStep by id
 */
func (c *RemoteApi) GetRemoteStep() {

	id := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	remotestep := &models.RemoteStep{Id: idInt}
	err := service.Remote.GetBase(remotestep)

	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}

	remotestepstruct := remotestep_struct{}
	remotestepstruct.Id = remotestep.Id
	remotestepstruct.Name = remotestep.Name
	remotestepstruct.Desc = remotestep.Desc
	json.Unmarshal([]byte(remotestep.Actions), &remotestepstruct.Actions)
	c.ReturnSuccess(remotestepstruct)
}

/**
*  load RemoteStep by page
 */
func (c *RemoteApi) RemoteStepList() {

	page := c.Query2Int("page", 1)
	pageSize := c.Query2Int("page_size", 10)

	c.CheckPage(&page, &pageSize)

	list := make([]models.RemoteStep, 0, pageSize)

	count, err := service.Remote.ListByPageWithSort(page, pageSize, &models.RemoteStep{}, &list, "-id")
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	liststruct := make([]remotestep_struct, len(list), pageSize)

	for i, fi := range list {
		liststruct[i].Id = fi.Id
		liststruct[i].Name = fi.Name
		liststruct[i].Desc = fi.Desc
		json.Unmarshal([]byte(fi.Actions), &liststruct[i].Actions)

	}

	c.ReturnPageContent(page, pageSize, count, liststruct)
}

/**
* remove single RemoteStep by id
 */
func (c *RemoteApi) RemoteStepDelete() {

	id := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	err := service.Remote.CheckStepDelete(idInt)
	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}
	c.ReturnSuccess("")
}

/**
*  create new RemoteStepImpl
 */
func (c *RemoteApi) RemoteActionImplAppend() {
	req := remoteactionimpl_struct{}

	err := c.Body2Json(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	templatestr, err := json.Marshal(req.Template)
	if err != nil {
		c.ReturnFailed("Bad parameter: "+err.Error(), 400)
		return
	}

	// check if action with id = action_id exists
	act := &models.RemoteAction{Id: req.ActionId}
	err = service.Remote.GetBase(act)
	if err != nil {
		c.ReturnFailed("action "+
			strconv.Itoa(req.ActionId)+" not found", 400)
	}

	// check if an action impl already exists
	old := &models.RemoteActionImpl{}
	conditions := make(map[string]interface{})
	conditions["Type"] = req.Type
	conditions["ActionId"] = req.ActionId

	err = service.Remote.GetByMultiFieldValue(old, conditions)
	if err == nil {
		c.ReturnFailed("action impl already exists for action "+
			strconv.Itoa(req.ActionId)+"["+req.Type+"]", 400)
	}

	//if !utils.GetValidateUtil().ValidateString(string(templatestr)) {
	//	c.ReturnFailed("the template format is not correct "+
	//		strconv.Itoa(req.ActionId)+"["+req.Type+"]", 400)
	//}

	data := models.RemoteActionImpl{
		//Id:       req.Id,
		Type:     req.Type,
		Template: string(templatestr),
		ActionId: req.ActionId,
	}
	err = service.Cluster.InsertBase(&data)

	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}
	c.ReturnSuccess(data.Id)
}

/**
*  load RemoteActionImpl by id
 */
func (c *RemoteApi) GetRemoteActionImpl() {

	actionId := c.Ctx.Input.Param(":actionId")
	idInt, _ := strconv.Atoi(actionId)
	//actionType := c.Ctx.Input.Param(":type")
	actionType := ANSIBLE

	remoteactionimpl := &models.RemoteActionImpl{}

	conditions := make(map[string]interface{})
	conditions["Type"] = actionType
	conditions["ActionId"] = idInt

	err := service.Remote.GetByMultiFieldValue(remoteactionimpl, conditions)

	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}

	remoteactionimpltruct := remoteactionimpl_struct{}
	remoteactionimpltruct.Id = remoteactionimpl.Id
	remoteactionimpltruct.Type = remoteactionimpl.Type
	json.Unmarshal([]byte(remoteactionimpl.Template), &remoteactionimpltruct.Template)
	remoteactionimpltruct.ActionId = remoteactionimpl.ActionId
	c.ReturnSuccess(remoteactionimpltruct)
}

/**
* remove single RemoteActionImplDelete by id
 */
func (c *RemoteApi) RemoteActionImplDelete() {

	id := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)
	//actionType := c.Ctx.Input.Param(":type")
	actionType := ANSIBLE

	remoteactionimpl := &models.RemoteActionImpl{}

	conditions := make(map[string]interface{})
	conditions["Type"] = actionType
	conditions["ActionId"] = idInt

	err := service.Remote.DeleteByMultiFieldValue(remoteactionimpl, conditions)

	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}
	c.ReturnSuccess("")
}

/**
* update single RemoteActionImpl update by id
 */
func (c *RemoteApi) RemoteActionImplUpdate() {

	id := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)
	//actionType := c.Ctx.Input.Param(":type")
	actionType := ANSIBLE

	conditions := make(map[string]interface{})
	conditions["Type"] = actionType
	conditions["ActionId"] = idInt

	req := remoteactionimpl_struct{}

	err := c.Body2Json(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	remoteactionimpl := &models.RemoteActionImpl{}
	err = service.Remote.GetByMultiFieldValue(remoteactionimpl, conditions)

	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}

	remoteactionimpl.Type = req.Type

	templateStr, _ := json.Marshal(req.Template)
	remoteactionimpl.Template = string(templateStr)
	//remoteactionimpl.ActionId = req.ActionId

	err = service.Remote.UpdateBase(remoteactionimpl)

	if err != nil {
		c.ReturnFailed(err.Error(), 404)
		return
	}
	c.ReturnSuccess("")
}

/**
*  load RemoteActionImpl by page
 */
func (c *RemoteApi) RemoteActionImplList() {

	page := c.Query2Int("page", 1)
	pageSize := c.Query2Int("page_size", 10)

	c.CheckPage(&page, &pageSize)

	list := make([]models.RemoteActionImpl, 0, pageSize)

	count, err := service.Remote.ListByPage(page, pageSize, &models.RemoteActionImpl{}, &list)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	liststruct := make([]remoteactionimpl_struct, len(list), pageSize)

	for i, fi := range list {
		liststruct[i].Id = fi.Id
		liststruct[i].Type = fi.Type
		json.Unmarshal([]byte(fi.Template), &liststruct[i].Template)
		liststruct[i].ActionId = fi.ActionId

	}

	c.ReturnPageContent(page, pageSize, count, liststruct)
}
