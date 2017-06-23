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
	"weibo.com/opendcp/imagebuild/code/errors"
	"weibo.com/opendcp/imagebuild/code/util"
	"weibo.com/opendcp/imagebuild/code/web/models"
)

/**
插件扩展接口
 */
type PluginExtensionInterfaceController struct {
	BasicController
}

func (c *PluginExtensionInterfaceController) Post() {
	plugin := c.GetString("plugin")
	method := c.GetString("method")
	params := c.Ctx.Request.Form
	if plugin == "" || method == "" {
		response := models.Response{
			Code:   errors.EXTENSION_INTERFACE_PARAM_ERROR,
			ErrMsg: errors.ErrorCodeToMessage(errors.EXTENSION_INTERFACE_PARAM_ERROR)}

		c.Data["json"] = response
		c.ServeJSON(true)
		return
	}
	cluster := c.HarborProjectName();
	// 调用插件方法
	paramsNew := util.Transform(params)
	//设置cluster
	paramsNew["cluster"] = cluster

	code, ret := models.AppServer.CallExtensionInterface(plugin, method, paramsNew)

	resp := models.BuildResponse(
		code,
		ret,
		errors.ErrorCodeToMessage(code))

	c.Data["json"] = resp

	c.ServeJSON(true)
}
