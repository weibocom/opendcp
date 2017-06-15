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
	"strconv"
	"weibo.com/opendcp/imagebuild/code/errors"
	"weibo.com/opendcp/imagebuild/code/web/models"
	"github.com/astaxie/beego"
)
/**
build image
 */
type BuildImageController struct {
	BasicController
}

func (c *BuildImageController) Post() {
	log.Infof("BuildImageController: %s", c.Ctx.Request.Form)

	project := c.GetString("projectName")
	cluster := c.BizName()
	tag := c.GetString("tag")

	operator := c.Operator()

	if project == "" || operator == "" || tag == "" || cluster == ""{
		beego.Warn("cluster,project,operator,tag should not be empy when building project")
		beego.Warn(project)
		beego.Warn(cluster)
		beego.Warn(tag)
		beego.Warn(operator)

		log.Error("cluster,project,operator,tag should not be empy when building project")
		resp := models.BuildResponse(
			errors.PARAMETER_INVALID,
			-1,
			errors.ErrorCodeToMessage(errors.PARAMETER_INVALID))

		c.Data["json"] = resp
		c.ServeJSON(true)
		return
	}

	code, id := models.AppServer.BuildImage(cluster, project, tag, operator)
	idStr := strconv.FormatInt(id, 10)

	var resp interface{}

	resp = models.BuildResponse(
		code,
		idStr,
		errors.ErrorCodeToMessage(code))

	c.Data["json"] = resp
	c.ServeJSON(true)

	return
}
