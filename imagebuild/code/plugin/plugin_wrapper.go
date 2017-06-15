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



package plugin

import (
	"bytes"
	"encoding/json"
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/beego/i18n"
	"github.com/dullgiulio/pingo"
	stackError "github.com/go-errors/errors"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"strconv"
	"text/template"
	"weibo.com/opendcp/imagebuild/code/env"
	"weibo.com/opendcp/imagebuild/code/util"
)

var DOCKERFILE_PLUGIN = 1
var BUILD_PLUGIN = 2

/**
* 对现有插件Plugin的扩展
 */
type PluginWrapper struct {
	// 在pipeline中的位置信息
	IndexInPipeline int `json:"indexInPipeline"`
	// 插件类型
	Plugin_type int `json:"plugin_type"`
	// 插件名称
	Plugin_name string `json:"plugin_name"`
	// 插件目录
	Plugin_directory string `json:"plugin_directory"`
	// 插件
	Plugin *pingo.Plugin `json:"plugin"`
	// 该插件的配置信息，所属项目不同配置信息不同。
	Config map[string]interface{} `json:"config"`
}

func CopyPluginWrapper(plugin *PluginWrapper) *PluginWrapper {
	copyPlugin := &PluginWrapper{
		Plugin_type:      plugin.Plugin_type,
		Plugin_name:      plugin.Plugin_name,
		Plugin_directory: plugin.Plugin_directory,
		Plugin:           plugin.Plugin,
		Config:           plugin.Config}

	return copyPlugin
}

func BuildPluginWrapper(pluginName string, pluginDirectory string) *PluginWrapper {
	wrapper := &PluginWrapper{Plugin_name: pluginName, Plugin_directory: pluginDirectory}
	plugin := pingo.NewPlugin("tcp", pluginDirectory+"/"+pluginName+"_plugin")
	plugin.Start()
	wrapper.Plugin = plugin
	return wrapper
}

func (pw *PluginWrapper) View() string {
	t := template.New("")

	t.Funcs(template.FuncMap{"defaultV": util.DefaultValue})
	t.Funcs(template.FuncMap{"defaultA": util.DefaultEmptyArray})
	t.Funcs(template.FuncMap{"startwith": util.StartWith})
	t.Funcs(template.FuncMap{"isArray": util.IsArray})
	t.Funcs(template.FuncMap{"escaped": util.Escaped})
	t.Funcs(template.FuncMap{"i18n": i18n.Tr})

	configPage := pw.Plugin_directory + "/index.html"
	if !util.IsFileExists(configPage) {
		return ""
	}

	configPageContent, error := ioutil.ReadFile(configPage)
	if error != nil {
		return ""
	}

	t, _ = t.Parse(string(configPageContent))
	var htmlContent bytes.Buffer

	var config map[string]interface{}
	if pw.Config == nil {
		config = make(map[string]interface{}, 0)
	} else {
		config = pw.Config
	}
	config["Lang"] = util.Lang

	t.Execute(&htmlContent, config)
	return htmlContent.String()
}

func (pw *PluginWrapper) Process(project string, input interface{}) (error, interface{}) {
	params := make(map[string]interface{}, 0)

	params["project"] = project

	// 系统变量，插件可能会用到
	util.PackageSystemEnvIntoParam(params)

	bytes, _ := json.Marshal(pw.Config)
	// convert map to json string because Pingo doesn't receive `map[string][]string` as parameter
	params["config"] = string(bytes)

	var result string

	pluginName := util.ConvertToHump(pw.Plugin_name+"_plugin")
	if pw.Plugin_type == DOCKERFILE_PLUGIN {
		params["input"] = input
		log.Infof("call %s.Process", pluginName)
		error := pw.Plugin.Call(pluginName+".Process", params, &result)
		log.Infof("%s return is %s", pw.Plugin_name, result)
		// 生成dockerfile的插件返回值是string
		if error != nil {
			log.Errorf("%s", error)
			return error, ""
		}

		if !util.IsString(result) {
			return errors.New(""), ""
		}

		return nil, result
	} else {
		//　用来构建的插件返回值map[string]interface{}
		if util.IsMap(input) {
			for key, value := range input.(map[string]interface{}) {
				params[key] = value
			}
		}


		error := pw.Plugin.Call(pluginName+".Process", params, &result)
		log.Infof("%s return is %s", pw.Plugin_name, result)
		if error != nil {
			log.Errorf("%s", error)
			return error, make(map[string]interface{}, 0)
		}

		resultMap := make(map[string]interface{}, 0)
		error = json.Unmarshal([]byte(result), &resultMap)
		if error != nil {
			return error, ""
		}

		// merge result into input as next input
		if !util.IsMap(input) {
			return nil, resultMap
		}

		for key, value := range input.(map[string]interface{}) {
			resultMap[key] = value
		}

		return nil, resultMap
	}
}

func (pw *PluginWrapper) Save(projectName string, config map[string]interface{}) string {
	pw.Config = config

	// rewrite config file of this plugin
	bytes, error := yaml.Marshal(config)
	if error != nil {
		log.Error(stackError.New(error).ErrorStack())
		return ""
	}
	var configPath string = env.PROJECT_CONFIG_BASEDIR + projectName
	if pw.Plugin_type == BUILD_PLUGIN {
		configPath += "/build/" + pw.Plugin_name
	} else {
		configPath += "/dockerfile/" + pw.Plugin_name
	}
	if pw.IndexInPipeline != 0 {
		configPath += ("_" + strconv.Itoa(pw.IndexInPipeline))
	}

	log.Infof("%s", configPath)
	log.Infof("%s", string(bytes))

	error = ioutil.WriteFile(configPath, bytes, 0777)
	if error != nil {
		log.Error(stackError.New(error).ErrorStack())
	}
	return ""
}

/**
 * find plug by name
 */
func FindPlugByName(dstName string, plugins *util.ConcurrentMap) *PluginWrapper {
	if plugins.ContainsKey(dstName) {
		return plugins.Get(dstName).(*PluginWrapper)
	}
	log.Info(dstName)
	log.Info(plugins.ToPrettyString())
	return nil
}
