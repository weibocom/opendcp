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
	log "github.com/Sirupsen/logrus"
	"weibo.com/opendcp/imagebuild/code/errors"
	"weibo.com/opendcp/imagebuild/code/web/models"
)
/**
build image history
 */
type BuildImageHistoryController struct {
	BasicController
}

func (c *BuildImageHistoryController) Get() {
	project := c.GetString("projectName")
	operator := c.Operator()

	if project == "" || operator == "" {
		log.Error("project,operator should not be empy when get build history!")
		resp := models.BuildResponse(
			errors.PARAMETER_INVALID,
			-1,
			errors.ErrorCodeToMessage(errors.PARAMETER_INVALID))

		c.Data["json"] = resp
		c.ServeJSON(true)
		return
	}

	model := models.AppServer.GetBuildLastHistory(project)

	if model == nil {
		log.Error("get projectName:%s build history is nil!", project)
		resp := models.BuildResponse(
			errors.BUILD_PROJECT_NOT_EXIST,
			-1,
			errors.ErrorCodeToMessage(errors.BUILD_PROJECT_NOT_EXIST))

		c.Data["json"] = resp
		c.ServeJSON(true)
		return
	}

	var resp interface{}
	resp = models.BuildResponse(
		errors.OK,
		model.State(),
		errors.ErrorCodeToMessage(errors.OK))

	c.Data["json"] = resp
	c.ServeJSON(true)

	return
}
