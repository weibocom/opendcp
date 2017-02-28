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



package project

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"github.com/beego/i18n"
	"github.com/go-yaml/yaml"
	"html/template"
	"io/ioutil"
	"sort"
	"weibo.com/opendcp/imagebuild/code/env"
	h "weibo.com/opendcp/imagebuild/code/handler"
	"weibo.com/opendcp/imagebuild/code/handler/interfaces"
	plu "weibo.com/opendcp/imagebuild/code/plugin"
	"weibo.com/opendcp/imagebuild/code/service"
	"weibo.com/opendcp/imagebuild/code/util"
	"encoding/json"
)

/**
项目封装
 */
type PluggedProject struct {
	ProjectInfo
	DockerFileGenerator interfaces.Handler

	DockerfilePlugins *util.ConcurrentMap
	BuildPlugins      *util.ConcurrentMap

	logs    []string
}

func (p *PluggedProject) Init() bool {
	p.readInfo()
	p.DockerFileGenerator.Init()
	return true
}

func (p *PluggedProject) View(lang string) string {
	view := ProjectView{}
	view.DockerfileView = p.DockerFileGenerator.View()

	dockerfilePlugins := make([]string, 0)
	for plugin := range p.DockerfilePlugins.Iterator() {
		dockerfilePlugins = append(dockerfilePlugins, plugin.Value.(*plu.PluginWrapper).Plugin_name)
	}
	sort.Strings(dockerfilePlugins)

	templatePath := env.GLOBLE_CONFIG_BASEDIR + "/" + "plugins.html"
	t := template.New("")
	t.Funcs(template.FuncMap{"defaultV": util.DefaultValue})
	t.Funcs(template.FuncMap{"defaultA": util.DefaultEmptyArray})
	t.Funcs(template.FuncMap{"endwith": util.EndWith})
	t.Funcs(template.FuncMap{"startwith": util.StartWith})
	t.Funcs(template.FuncMap{"isArray": util.IsArray})
	t.Funcs(template.FuncMap{"unescaped": util.Unescaped})
	t.Funcs(template.FuncMap{"i18n": i18n.Tr})

	configPageContent, error := ioutil.ReadFile(templatePath)
	if error != nil {
		return ""
	}

	t, _ = t.Parse(string(configPageContent))
	var htmlContent bytes.Buffer

	config := make(map[string]interface{}, 0)
	config["view"] = view
	config["pluginViewUrl"] = env.PLUGIN_VIEW_RUL
	config["extensionInterfaceUrl"] = env.EXTENSION_INTERFACE_RUL
	config["dockerfilePlugins"] = dockerfilePlugins
	config["Lang"] = lang
	config["project"] = p.Name

	config["cluster"] = p.Cluster
	config["defineDockerFileType"] = p.DefineDockerFileType
	config["server"] = env.SERVER_HOST
	config["DefaultProjectName"] = "DefaultProjectName"

	t.Execute(&htmlContent, config)
	return htmlContent.String()
}

func (p *PluggedProject) BuildImage() bool {

	projectPath := env.PROJECT_CONFIG_BASEDIR
	dockerFilePath := projectPath + p.Name + "/tmp/"
	util.ClearFolder(dockerFilePath)

	// create docker file
	if !p.DockerFileGenerator.Handle() {
		return false
	}

	return true
}

func (p *PluggedProject) BuildAndPushImage(tag string) bool {
	registry := env.HARBOR_ADDRESS
	fullImageName := registry + "/" + p.Cluster + "/" + p.Name + ":" + tag

	projectPath := env.PROJECT_CONFIG_BASEDIR
	dockerFilePath := projectPath + p.Name + "/tmp/"

	log.Info("BuildImage dockerFilePath:" + dockerFilePath + " fullImageName:" + fullImageName)

	p.appendLog("start build image ....")

	logStr, err := service.GetDockerOperatorInstance().BuildImage(dockerFilePath, fullImageName)

	p.appendLog(logStr)

	if  err != nil {
		log.Error("Build Image with error:", err)
		p.appendLog(err.Error())
		return false
	}

	log.Info("Login Harbor")
	if err := service.GetDockerOperatorInstance().LoginHarbor(); err != nil {
		log.Error("Login Harbor with error:", err)
		p.appendLog(err.Error())
		return false
	}

	p.appendLog("login haror success ...")

	log.Info("PushImage fullImageName:" + fullImageName)

	logStr, err = service.GetDockerOperatorInstance().PushImage(dockerFilePath, fullImageName)

	p.appendLog(logStr)

	if err != nil {
		log.Error("Push Image with error:", err)
		return false
	}

	return true
}

func (p *PluggedProject) Save(configs []map[string]interface{}) bool {
	p.DockerFileGenerator.Save(configs, p.DockerfilePlugins)
	return true
}

func (p *PluggedProject) readInfo() {
	// load project info
	content, error := ioutil.ReadFile(env.PROJECT_CONFIG_BASEDIR + p.Name + "/" + "info")
	if error != nil {
		log.Error("readfile with error:", error)
		panic("Init Failed!")
	}
	infoMap := make(map[string]string, 0)
	yaml.Unmarshal(content, &infoMap)
	p.CreateTime = infoMap["createTime"]
	p.LastModifyTime = infoMap["lastModifyTime"]
	p.Creator = infoMap["creator"]
	p.LastModifyOperator = infoMap["lastModifyOperator"]
	p.Cluster = infoMap["cluster"]
	p.DefineDockerFileType = infoMap["defineDockerFileType"]
}

func (p *PluggedProject) appendLog(line string) {
	p.logs = append(p.logs, line+"\n")
}

func (p *PluggedProject) GetLogInJson() string {
	logsBytes, err := json.Marshal(p.logs)

	if err != nil {
		return ""
	}
	return string(logsBytes)
}

func (p *PluggedProject) ClearLog() {
	p.logs = make([]string, 0)
}

// 构建project对象
func BuildPluginProject(projectName string,
	creator string,
	cluster string,
	defineDockerFileType string,
	createTime string,
	dockerfilePlugins *util.ConcurrentMap,
	buildPlugins *util.ConcurrentMap) *PluggedProject {

	// build Project object
	project := &PluggedProject{}
	project.Creator = creator
	project.Name = projectName
	project.CreateTime = createTime
	project.LastModifyTime = createTime
	project.LastModifyOperator = creator
	project.Cluster = cluster
	project.DefineDockerFileType = defineDockerFileType

	project.logs = make([]string, 0)

	var dockerfileBuilder interfaces.Handler

	dockerfileBuilder = h.BuildExtensibleDockerFileGenerator(projectName,
		"dockerfile",
		dockerfilePlugins)

	project.DockerFileGenerator = dockerfileBuilder
	project.DockerfilePlugins = dockerfilePlugins
	project.BuildPlugins = buildPlugins

	return project
}
