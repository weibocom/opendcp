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
	"weibo.com/opendcp/imagebuild/code/web/models"
)
/**
info视图
 */
type InfoViewController struct {
	BasicController
}

func (c *InfoViewController) Get() {
	project := c.GetString("projectName")
	cluster := c.BizName()
	c.Layout = "project.tpl"
	c.TplName = "info.tpl"
	// TODO check code
	_, info := models.AppServer.GetProjectInfo(cluster, project)

	c.Data["cluster"] = cluster
	c.Data["project"] = project
	c.Data["name"] = info.Name
	c.Data["createTime"] = info.CreateTime
	c.Data["creator"] = info.Creator
	c.Data["lastModifyTime"] = info.LastModifyTime
	c.Data["lastModifyOperator"] = info.LastModifyOperator
	c.Data["page"] = "info"
}
