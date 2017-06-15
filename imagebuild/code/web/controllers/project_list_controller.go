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
