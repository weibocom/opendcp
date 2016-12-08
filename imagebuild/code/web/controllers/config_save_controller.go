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
	"reflect"
	"strings"
	"weibo.com/opendcp/imagebuild/code/errors"
	"weibo.com/opendcp/imagebuild/code/web/models"
	"net/url"
)

var specialStrings = []string{"!","@","#","$","%","^","&","*","(",")","=","'","\"","/","\\","|","<",">","{","}","[","]"}

/**
保存Dockerfile配置
 */
type ConfigSaveController struct {
	BasicController
}

func (c *ConfigSaveController) Post() {
	log.Infof("BuildProgressController: ", c.Ctx.Request.Form)

	// get origin request body
	body := string(c.Ctx.Input.RequestBody)

	keyValues := strings.Split(body, "&")

	// config of each plugin is mapped to a map
	configs := make([]map[string]interface{}, 0)
	pluginMap := make(map[string]interface{}, 0)

	var project string
	var addOrUpdate string
	var cluster string
	var defineDockerFileType string
	var pluginConfig map[string]interface{}

	for _, keyValueStr := range keyValues {
		keyValue := strings.Split(keyValueStr, "=")
		attributeName := keyValue[0]
		attributeValue,_ := url.QueryUnescape(keyValue[1])
		attributeValue = strings.Trim(attributeValue, " ")
		if attributeName == "$$plugin" {
			if _, ok := pluginMap["$$plugin"]; !ok {
				pluginConfig = make(map[string]interface{}, 0)
				pluginConfig["$$plugin"] = attributeValue
				pluginMap[attributeValue] = pluginConfig
				configs = append(configs, pluginConfig)
			}
		} else if attributeName == "project" {
			project = attributeValue
		} else if attributeName == "Cluster" {
			cluster = attributeValue
		} else if attributeName == "DefineDockerFileType" {
			defineDockerFileType = attributeValue
		} else if attributeName == "addOrUpdate" {
			addOrUpdate = attributeValue
		}else {
			tmp := strings.Split(attributeName, ".")
			if len(tmp) != 2 {
				continue
			}
			pluginName := tmp[0]
			attributeName := tmp[1]
			if _, ok := pluginMap[pluginName]; !ok {
				pluginConfig = make(map[string]interface{}, 0)
				pluginConfig["$$plugin"] = pluginName
				pluginMap[pluginName] = pluginConfig
				configs = append(configs, pluginConfig)
			} else {
				pluginConfig = pluginMap[pluginName].(map[string]interface {})
			}
			if _, ok := pluginConfig[attributeName]; !ok {
				pluginConfig[attributeName] = attributeValue
			} else {
				currentValue := pluginConfig[attributeName]
				if reflect.TypeOf(currentValue).String() != "string" {
					// add to tail
					currentValue = append(currentValue.([]string), attributeValue)
					pluginConfig[attributeName] = currentValue
				} else {
					// if a attribute appear more than once, transform the attribute value to array
					pluginConfig[attributeName] = []string{currentValue.(string), attributeValue}
				}
			}
		}
	}

	log.Infof("configs---------->%s", configs)

	projectName := project

	projectName = strings.ToLower(projectName)
	for _,spec := range specialStrings {
		if strings.Contains(projectName, spec) {
			var resp = models.BuildResponse(
				errors.PARAMETER_INVALID,
				"projectName contains special char:" + spec,
				errors.ErrorCodeToMessage(errors.PARAMETER_INVALID))
			c.Data["json"] = resp
			c.ServeJSON(true)
			return
		}
	}

	creator := c.Operator()

	exist := models.AppServer.IsProjectExist(projectName)
	if addOrUpdate == "add" && exist{
		var resp = models.BuildResponse(
			errors.CREATE_PROJECT_ALREADY_EXIST,
			"project:" + projectName + " already exist",
			errors.ErrorCodeToMessage(errors.CREATE_PROJECT_ALREADY_EXIST))
		c.Data["json"] = resp
		c.ServeJSON(true)
		return
	}

	var ok bool
	var code int
	if exist {
		ok, code = models.AppServer.UpdateProject(projectName, creator, cluster, defineDockerFileType)
	} else {
		ok, code = models.AppServer.NewProject(projectName, creator, cluster, defineDockerFileType)
	}

	if !ok {
		response := models.BuildResponse(code, "", errors.ErrorCodeToMessage(code))
		c.Data["json"] = response
		c.ServeJSON(true)
		return
	}

	succ := models.AppServer.SaveProjectConfig(project, configs)

	var resp interface{}

	if succ {
		resp = models.BuildResponse(
			errors.OK,
			"",
			errors.ErrorCodeToMessage(errors.OK))
	} else {
		resp = models.BuildResponse(
			errors.INTERNAL_ERROR,
			"",
			errors.ErrorCodeToMessage(errors.INTERNAL_ERROR))
	}

	c.Data["json"] = resp
	c.ServeJSON(true)
	return
}
