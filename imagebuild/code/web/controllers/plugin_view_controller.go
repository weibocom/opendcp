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
	"github.com/beego/i18n"
	"weibo.com/opendcp/imagebuild/code/web/models"
)

/**
插件视图
 */
type PluginViewController struct {
	BasicController
}

func (c *PluginViewController) Get() {
	pluginType, _ := c.GetInt("type")
	pluginName := c.GetString("plugin")
	originPluginView := c.GetString("plugin")

	pluginView := models.AppServer.GetPluginView(pluginType, pluginName)

	pluginName = i18n.Tr(c.Data["Lang"].(string), pluginName)
	var view string = "<div class='panel panel-default col-sm-10'>" +
		"<input type='hidden' name='$$plugin' id='$$plugin' value='" + originPluginView + "'/>"

	if originPluginView != "download_executable_source" && originPluginView != "static_dockerfile" && originPluginView != "download_dockerfile" {
		view = view + "<div class='panel-heading'>" + pluginName +
			"<div style='position:relative; float:right; right:0px'>" +
			"<span class='glyphicon glyphicon-chevron-up' onclick='movePluginUp(event)'></span>" +
			"<span class='glyphicon glyphicon-chevron-down' onclick='movePluginDown(event)'></span>" +
			"<span class='glyphicon glyphicon-remove' onclick='deletePlugin(event)'></span>" +
			"</div>" +
			"</div>"
	}
	view = view +
		"<div class='panel-body'>" + pluginView + "</div>" + "</div>"

	c.Ctx.ResponseWriter.Write([]byte(view))
}
