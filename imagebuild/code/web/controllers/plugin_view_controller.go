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
