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
创建项目
 */
type ProjectNewController struct {
	BasicController
}

func (c *ProjectNewController) Post() {
	log.Info("ProjectNewController: %s", c.Ctx.Request.Form)
	projectName := c.GetString("projectName")
	creator := c.Operator()
	cluster := c.BizName()
	if creator == "" || projectName == "" || cluster == ""{
		log.Error("cluster, creator,projectName should not be empy when building project")
		resp := models.BuildResponse(
			errors.PARAMETER_INVALID,
			-1,
			errors.ErrorCodeToMessage(errors.PARAMETER_INVALID))

		c.Data["json"] = resp
		c.ServeJSON(true)
		return
	}

	_, code := models.AppServer.NewProject(projectName, creator, cluster, "")
	response := models.BuildResponse(code, "", errors.ErrorCodeToMessage(code))
	c.Data["json"] = response
	c.ServeJSON(true)
	return
}
