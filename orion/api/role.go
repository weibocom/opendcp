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
	"fmt"
	"strconv"
	"weibo.com/opendcp/orion/models"
	"weibo.com/opendcp/orion/service"
)

type RoleApi struct {
	baseAPI
}

type role_struct struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	RoleFilePath string `json:"role_file_path"`
	Desc         string `json:"desc"`
	Templates    string `json:"templates"`
	Tasks        string `json:"tasks"`
	Vars         string `json:"vars"`
	UpdateTime   string `json:"update_time"`
}

func (f *RoleApi) URLMapping() {

	f.Mapping("RoleList", f.RoleList)
	f.Mapping("RoleAppend", f.RoleAppend)
	f.Mapping("RoleDelete", f.RoleDelete)
	f.Mapping("GetRole", f.GetRole)
	f.Mapping("RoleUpdate", f.RoleUpdate)

	f.Mapping("RoleResourceInfo", f.RoleResourceInfo)
	f.Mapping("RoleResourceList", f.RoleResourceList)
	f.Mapping("RoleResourceAppend", f.RoleResourceAppend)
	f.Mapping("RoleResourceDelete", f.RoleResourceDelete)
	f.Mapping("RoleResourceUpdate", f.RoleResourceUpdate)

	f.Mapping("TestPack", f.TestPack)
}

func (f *RoleApi) RoleList() {
	page := f.Query2Int("page", 1)
	pageSize := f.Query2Int("page_size", 10)

	f.CheckPage(&page, &pageSize)

	list := make([]models.Role, 0, pageSize)

	count, err := service.Role.ListByPageWithSort(page, pageSize, &models.Role{}, &list, "-id")
	if err != nil {
		f.ReturnFailed(err.Error(), 400)
		return
	}

	liststruct := make([]role_struct, len(list), pageSize)

	for i, fi := range list {
		liststruct[i].Id = fi.Id
		liststruct[i].Name = fi.Name
		liststruct[i].Desc = fi.Desc
		liststruct[i].RoleFilePath = fi.RoleFilePath
		liststruct[i].Templates = fi.Templates
		liststruct[i].Tasks = fi.Tasks
		liststruct[i].Vars = fi.Vars
		liststruct[i].UpdateTime = fi.UpdateTime.String()
	}

	f.ReturnPageContent(page, pageSize, count, liststruct)
}

func (f *RoleApi) RoleAppend() {
	obj := models.Role{}

	err := f.Body2Json(&obj)
	if err != nil {
		f.ReturnFailed("B2J Failed "+err.Error(), 400)
		return
	}

	params := []string{obj.Tasks, obj.Templates, obj.Vars, obj.Files, obj.Handles, obj.Meta}
	res, err := service.Role.CheckRoleParams(params)
	if err != nil {
		f.ReturnFailed(err.Error()+" The illegal param is "+res, 400)
		return
	}

	err = service.Role.InsertBase(&obj)
	if err != nil {
		f.ReturnFailed("Insert Failed "+err.Error(), 400)
		return
	}

	err = service.Role.BuildRoleFile(&obj)
	if err != nil {
		dirtyObj := &models.Role{
			Id:   obj.Id,
			Name: obj.Name,
		}
		_ = service.Role.DeleteBase(dirtyObj)
		_ = service.Role.RemoveRoleFile(dirtyObj)

		f.ReturnFailed("Write Failed "+err.Error(), 400)
		return
	}

	f.ReturnSuccess(obj.Id)
}

func (f *RoleApi) RoleDelete() {
	id := f.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	obj := &models.Role{Id: idInt}
	err := service.Role.GetBase(obj)
	if err != nil {
		f.ReturnFailed("This role dose not exit. "+err.Error(), 404)
		return
	}

	err = service.Role.DeleteBase(obj)
	if err != nil {
		f.ReturnFailed(err.Error(), 400)
		return
	}

	err = service.Role.RemoveRoleFile(obj)

	f.ReturnSuccess(nil)
}

func (f *RoleApi) GetRole() {
	id := f.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	obj := &models.Role{Id: idInt}
	err := service.Role.GetBase(obj)
	if err != nil {
		f.ReturnFailed("This role dose not exit. "+err.Error(), 404)
		return
	}

	f.ReturnSuccess(obj)
}

func (f *RoleApi) RoleUpdate() {
	id := f.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	obj := &models.Role{Id: idInt}

	err := service.Role.GetBase(obj)
	if err != nil {
		f.ReturnFailed("This role dose not exit. "+err.Error(), 404)
		return
	}

	req := &role_struct{}
	err = f.Body2Json(req)
	if err != nil {
		f.ReturnFailed(err.Error(), 400)
		return
	}

	params := []string{req.Tasks, req.Templates, req.Vars}
	res, err := service.Role.CheckRoleParams(params)
	if err != nil {
		f.ReturnFailed(err.Error()+" The illegal param is "+res, 400)
		return
	}

	obj.Tasks = req.Tasks
	obj.Templates = req.Templates
	obj.Vars = req.Vars
	obj.Name = req.Name
	obj.Desc = req.Desc
	obj.RoleFilePath = req.RoleFilePath

	err = service.Role.UpdateBase(obj)
	if err != nil {
		f.ReturnFailed(err.Error(), 400)
		return
	}

	err = service.Role.UpdateRoleFile(obj)
	if err != nil {
		f.ReturnFailed(err.Error(), 400)
		return
	}

	f.ReturnSuccess(nil)
}

func (f *RoleApi) RoleResourceInfo() {
	idInt := f.roleResourceCheckId()
	if idInt < 1 {
		f.ReturnFailed("The id is error !", 400)
		return
	}
	obj := &models.RoleResource{Id: idInt}
	err := service.Role.GetBase(obj)
	if err != nil {
		f.ReturnFailed(err.Error(), 404)
		return
	}
	f.ReturnSuccess(obj)

}

func (f *RoleApi) RoleResourceList() {
	page := f.Query2Int("page", 1)
	pageSize := f.Query2Int("page_size", 10)
	resource_type := f.Ctx.Input.Param(":type")
	f.CheckPage(&page, &pageSize)
	list := make([]models.RoleResource, 0, pageSize)
	if resource_type != "" {

		count, err := service.Role.ListByPageWithFilter(page, pageSize, &models.RoleResource{}, &list, "resource_type", resource_type)
		if err != nil {
			f.ReturnFailed(err.Error(), 400)
			return
		}
		f.ReturnPageContent(page, pageSize, count, list)
	} else {
		count, err := service.Role.ListByPageWithSort(page, pageSize, &models.RoleResource{}, &list, "Id")
		if err != nil {
			f.ReturnFailed(err.Error(), 400)
			return
		}
		f.ReturnPageContent(page, pageSize, count, list)
	}
}

func (f *RoleApi) RoleResourceAppend() {
	req := models.RoleResource{}
	err := f.roleResourceCheckParam(&req)
	if err != nil {
		f.ReturnFailed(err.Error(), 400)
		return
	}

	err = service.Role.InsertBase(&req)
	if err != nil {
		f.ReturnFailed(err.Error(), 400)
		return
	}
	f.ReturnSuccess(req.Id)
}

func (f *RoleApi) RoleResourceDelete() {
	idInt := f.roleResourceCheckId()
	if idInt < 1 {
		f.ReturnFailed("The id is error !", 400)
		return
	}

	obj := &models.RoleResource{Id: idInt}
	err := service.Role.GetBase(obj)
	if err != nil {
		f.ReturnFailed(err.Error(), 400)
		return
	}

	err = service.Role.DeleteBase(obj)
	if err != nil {
		f.ReturnFailed(err.Error(), 400)
		return
	}

	f.ReturnSuccess(nil)
}

func (f *RoleApi) RoleResourceUpdate() {
	idInt := f.roleResourceCheckId()
	if idInt < 1 {
		f.ReturnFailed("The id is error !", 400)
		return
	}

	req := models.RoleResource{}
	err := f.roleResourceCheckParam(&req)
	if err != nil {
		f.ReturnFailed(err.Error(), 400)
		return
	}

	roleResource := &models.RoleResource{Id: idInt}
	err = service.Role.GetBase(roleResource)
	if len(roleResource.Name) < 1 {
		f.ReturnFailed("The old data not found !", 404)
		return
	}

	roleResource.Name = req.Name
	roleResource.Desc = req.Desc
	roleResource.ResourceType = req.ResourceType
	roleResource.ResourceContent = req.ResourceContent
	roleResource.TemplateFilePath = req.TemplateFilePath
	roleResource.TemplateFilePerm = req.TemplateFilePerm
	roleResource.TemplateFileOwner = req.TemplateFileOwner
	roleResource.AssociateRole = req.AssociateRole

	err = service.Role.UpdateBase(roleResource)
	if err != nil {
		f.ReturnFailed(err.Error(), 400)
		return
	}

	f.ReturnSuccess("")
}

func (f *RoleApi) roleResourceCheckParam(req *models.RoleResource) error {
	err := f.Body2Json(&req)
	if err != nil {
		return err
	}

	if len(req.Name) < 1 {
		return fmt.Errorf("param is error!")
	}

	return nil
}

func (f *RoleApi) roleResourceCheckId() int {
	id := f.Ctx.Input.Param(":id")
	if len(id) < 1 {
		return 0
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}
	return idInt
}

func (f *RoleApi) TestPack() {
	step := f.Ctx.Input.Param(":step")
	roleName := f.Ctx.Input.Param(":name")
	roleNames := []string {roleName}

	err := service.Role.PackRoles(step, roleNames)

	if err != nil {
		f.ReturnFailed(err.Error(), 400)
		return
	}

	f.ReturnSuccess("")
}
