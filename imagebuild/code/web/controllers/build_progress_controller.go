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
根据id获取build历史记录
 */
type BuildProgressController struct {
	BasicController
}

func (c *BuildProgressController) Get() {

	log.Info("BuildProgressController: %s", c.Ctx.Request.Form)

	id, _ := c.GetInt("id", -1)
	if id == -1 {
		log.Error("Id should not be empty when quering build progress")
		resp := models.BuildResponse(
			errors.PARAMETER_INVALID,
			"",
			errors.ErrorCodeToMessage(errors.PARAMETER_INVALID))
		c.Data["json"] = resp
		c.ServeJSON(true)
		return
	}

	log.Infof("Query build progress, id is %s", id)
	history := models.AppServer.GetBuildHistory(id)

	var resp interface{}
	if history == nil {
		resp = models.BuildResponse(
			errors.INTERNAL_ERROR,
			"",
			errors.ErrorCodeToMessage(errors.INTERNAL_ERROR))
	} else {
		resp = models.BuildResponse(
			errors.OK,
			history.State(),
			errors.ErrorCodeToMessage(errors.OK))
	}

	c.Data["json"] = resp
	c.ServeJSON(true)
	return
}
