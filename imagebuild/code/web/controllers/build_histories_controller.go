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
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"weibo.com/opendcp/imagebuild/code/web/models"
)
/**
build历史记录
 */
type BuildHistoriesController struct {
	BasicController
}

var DEFAULT_CURSOR = 0
var DEFAULT_OFFSET = 10

func (c *BuildHistoriesController) Get() {
	projectName := c.GetString("projectName", "")
	if projectName == "" {
		log.Errorf("Project name shoud not be empty when quering build histories..")
		c.Ctx.ResponseWriter.Write([]byte("no"))
		return
	}

	cursor, _ := c.GetInt("cursor", DEFAULT_CURSOR)
	offset, _ := c.GetInt("offset", DEFAULT_OFFSET)

	log.Infof("Query build histories, project is %s, cursor is %s, offset is %s", projectName, cursor, offset)
	histories := models.AppServer.GetBuildHistories(cursor, offset, projectName)
	bytes, _ := json.Marshal(histories)
	log.Infof("Query build histories result is %s", string(bytes))
	c.Ctx.ResponseWriter.Write(bytes)
}
