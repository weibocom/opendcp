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
