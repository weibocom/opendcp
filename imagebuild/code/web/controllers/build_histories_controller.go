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
	cluster := c.HarborProjectName()
	if projectName == "" || cluster ==""{
		log.Errorf("Project name shoud not be empty when quering build histories..")
		c.Ctx.ResponseWriter.Write([]byte("no"))
		return
	}

	cursor, _ := c.GetInt("cursor", DEFAULT_CURSOR)
	offset, _ := c.GetInt("offset", DEFAULT_OFFSET)

	log.Infof("Query build histories, project is %s, cursor is %s, offset is %s", projectName, cursor, offset)
	histories := models.AppServer.GetBuildHistories(cursor, offset, cluster, projectName)
	bytes, _ := json.Marshal(histories)
	log.Infof("Query build histories result is %s", string(bytes))
	c.Ctx.ResponseWriter.Write(bytes)
}
