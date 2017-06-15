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



package handler

import (
	log "github.com/Sirupsen/logrus"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"weibo.com/opendcp/imagebuild/code/env"
	p "weibo.com/opendcp/imagebuild/code/plugin"
	"weibo.com/opendcp/imagebuild/code/util"
)

/**
extensible handler
 */
type AbstractExtensibleHandler struct {
	pipeline *p.PluginPipeline

	HandlerInfo

	lock sync.RWMutex
}

func (handler *AbstractExtensibleHandler) SetProjectName(project string) {
	handler.ProjectName = project
}

func (handler *AbstractExtensibleHandler) SetConfigRelativeFolder(folder string) {
	handler.ConfigRelativeFolder = folder
}

// public function
func (handler *AbstractExtensibleHandler) GetProjectName() string {
	return handler.ProjectName
}

func (handler *AbstractExtensibleHandler) View() p.PluginPipelineView {

	// read lock
	handler.lock.RLock()
	defer handler.lock.RUnlock()

	return handler.pipeline.View()
}

func (handler *AbstractExtensibleHandler) Save(configs []map[string]interface{}, plugins *util.ConcurrentMap) bool {

	handler.lock.Lock()
	defer handler.lock.Unlock()

	log.Infof("config to save is: %s, project: %s", configs, handler.ProjectName)

	recordPluginCount := make(map[string]int, 0)

	var pipeline *p.PluginPipeline = p.BuildPluginPipeline(
		handler.pipeline.GetProjectName(),
		handler.pipeline.GetPipelineName(),
		handler.pipeline.GetPipelineDescription())

	handler.pipeline.ClearAllConfig(handler.ConfigRelativeFolder)

	// 重建目录
	for _, config := range configs {
		plugName := config["$$plugin"].(string)
		plugin := p.FindPlugByName(plugName, plugins)
		// not exist
		if plugin == nil {
			continue
		}

		var number int
		if currentNumber, ok := recordPluginCount[plugName]; ok {
			number = currentNumber + 1
		} else {
			number = 1
		}

		recordPluginCount[plugName] = number

		copyPlugin := p.CopyPluginWrapper(plugin)
		copyPlugin.IndexInPipeline = number
		copyPlugin.Save(handler.ProjectName, config)
		pipeline.AddPluginToTail(copyPlugin)
	}

	handler.pipeline = pipeline

	// 替换项目的插件列表
	pluginList := handler.pipeline.PluginList()
	pluginListPath := env.PROJECT_CONFIG_BASEDIR + handler.ProjectName + "/" + handler.ConfigRelativeFolder + "/plug_list"
	ioutil.WriteFile(pluginListPath, []byte(pluginList), 0777)

	return true
}

func (handler *AbstractExtensibleHandler) initPipeline(projectName string, plugins *util.ConcurrentMap) {
	pluginListFile := env.PROJECT_CONFIG_BASEDIR + handler.GetProjectName() + "/" + handler.ConfigRelativeFolder + "/plug_list"
	content, error := ioutil.ReadFile(pluginListFile)
	if error != nil {
		log.Error("Load extension plug list from config file failed, config file path: " + pluginListFile)
		os.Exit(-1)
	}

	if len(content) > 0 {
		recordPluginCount := make(map[string]int, 0)
		pluginList := strings.Split(string(content), ",")
		for _, pluginName := range pluginList {
			var number int
			if currentNumber, ok := recordPluginCount[pluginName]; ok {
				number = currentNumber + 1
			} else {
				number = 1
			}

			recordPluginCount[pluginName] = number

			// find executable plug
			plugin := p.FindPlugByName(pluginName, plugins)

			if plugin == nil {
				continue
			}

			// wrapper
			pluginWrapper := &p.PluginWrapper{
				Plugin_type:      plugin.Plugin_type,
				Plugin_name:      pluginName,
				Plugin_directory: plugin.Plugin_directory,
				Plugin:           plugin.Plugin,
				IndexInPipeline:  number}

			configMap := handler.readConfigByPlug(pluginWrapper)
			pluginWrapper.Config = configMap

			handler.pipeline.AddPluginToTail(pluginWrapper)
		}
	}
	handler.pipeline.SetProjectName(projectName)
}

func (handler *AbstractExtensibleHandler) readConfigByPlug(plugin *p.PluginWrapper) map[string]interface{} {

	// load config
	configPath := env.PROJECT_CONFIG_BASEDIR + handler.GetProjectName() + "/" + handler.ConfigRelativeFolder + "/" + plugin.Plugin_name

	if plugin.IndexInPipeline != 0 {
		configPath += ("_" + strconv.Itoa(plugin.IndexInPipeline))
	}

	exist := util.IsFileExists(configPath)
	configMap := make(map[string]interface{}, 0)
	if !exist {
		return configMap
	}

	config, error := ioutil.ReadFile(configPath)
	if error != nil {
		log.Errorf("%s", error)
		os.Exit(-1)
	}

	yaml.Unmarshal(config, &configMap)

	return configMap
}
