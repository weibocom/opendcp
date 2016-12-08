/*
 *  Copyright 2009-2016 Weibo, Inc.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package controllers

import (
	log "github.com/Sirupsen/logrus"
	"weibo.com/opendcp/imagebuild/code/errors"
	pro "weibo.com/opendcp/imagebuild/code/project"
	"weibo.com/opendcp/imagebuild/code/web/models"
)
/**
项目列表
 */
type ProjectListController struct {
	BasicController
}

func (c *ProjectListController) Get() {
	projectName := c.GetString("projectName")
	page, _ := c.GetInt("page", 1)
	pageSize, _ := c.GetInt("page_size", 20)
	if pageSize == 0 {
		pageSize = 20
	}

	if pageSize > 20 {
		pageSize = 20
	}

	creator := c.Operator()
	if creator == "" {
		log.Error("creator should not be empy when building project")
		resp := models.BuildResponse(
			errors.PARAMETER_INVALID,
			-1,
			errors.ErrorCodeToMessage(errors.PARAMETER_INVALID))

		c.Data["json"] = resp
		c.ServeJSON(true)
		return
	}

	projects := models.AppServer.GetProjects(projectName)

	totalCount := projects.Len()
	totalPages := totalCount / pageSize
	if totalCount%pageSize != 0 {
		totalPages++
	}

	if page <= 0 {
		page = 1
	}

	if page > totalPages {
		page = totalPages
	}

	var datas []pro.ProjectInfo
	if totalCount > 0 {
		datas = projects[(page-1)*pageSize:]
	}
	if len(datas) > pageSize {
		datas = datas[:pageSize]
	}

	resp := models.BuildProjectListResponse(
		errors.OK,
		datas,
		errors.ErrorCodeToMessage(errors.OK),
		page,
		pageSize,
		totalCount)

	c.Data["json"] = resp
	c.ServeJSON(true)
	return
}
