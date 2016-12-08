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

	// 调用插件方法
	paramsNew := util.Transform(params)

	code, ret := models.AppServer.CallExtensionInterface(plugin, method, paramsNew)

	resp := models.BuildResponse(
		code,
		ret,
		errors.ErrorCodeToMessage(code))

	c.Data["json"] = resp

	c.ServeJSON(true)
}
