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
	"net/http"
	"weibo.com/opendcp/imagebuild/code/web/models"
)
/**
Independent配置视图
 */
type IndependentConfigViewController struct {
	BasicController
}

func (c *IndependentConfigViewController) Get() {
	project := c.GetString("projectName")
	_, configView := models.AppServer.GetProjectConfigView(project)
	if configView == "" {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		return
	}

	c.TplName = "config.tpl"

	c.Data["project"] = project
	c.Data["view"] = configView
	c.Data["showSubmit"] = ""
}
