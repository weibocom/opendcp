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
	"weibo.com/opendcp/imagebuild/code/web/models"
)
/**
build image history
 */
type BuildImageHistoryController struct {
	BasicController
}

func (c *BuildImageHistoryController) Get() {
	project := c.GetString("projectName")
	operator := c.Operator()

	if project == "" || operator == "" {
		log.Error("project,operator should not be empy when get build history!")
		resp := models.BuildResponse(
			errors.PARAMETER_INVALID,
			-1,
			errors.ErrorCodeToMessage(errors.PARAMETER_INVALID))

		c.Data["json"] = resp
		c.ServeJSON(true)
		return
	}

	model := models.AppServer.GetBuildLastHistory(project)

	if model == nil {
		log.Error("get projectName:%s build history is nil!", project)
		resp := models.BuildResponse(
			errors.BUILD_PROJECT_NOT_EXIST,
			-1,
			errors.ErrorCodeToMessage(errors.BUILD_PROJECT_NOT_EXIST))

		c.Data["json"] = resp
		c.ServeJSON(true)
		return
	}

	var resp interface{}
	resp = models.BuildResponse(
		errors.OK,
		map[string]interface{}{
			"state":model.State(),
			"logs": model.Logs(),
		},
		errors.ErrorCodeToMessage(errors.OK))

	c.Data["json"] = resp
	c.ServeJSON(true)

	return
}
