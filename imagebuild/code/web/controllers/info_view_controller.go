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
	c.Layout = "project.tpl"
	c.TplName = "info.tpl"
	// TODO check code
	_, info := models.AppServer.GetProjectInfo(project)

	c.Data["project"] = project
	c.Data["name"] = info.Name
	c.Data["createTime"] = info.CreateTime
	c.Data["creator"] = info.Creator
	c.Data["lastModifyTime"] = info.LastModifyTime
	c.Data["lastModifyOperator"] = info.LastModifyOperator
	c.Data["page"] = "info"
}
