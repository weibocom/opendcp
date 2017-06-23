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



package code

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
	"weibo.com/opendcp/imagebuild/code/env"
	"weibo.com/opendcp/imagebuild/code/errors"
	"weibo.com/opendcp/imagebuild/code/model"
	plu "weibo.com/opendcp/imagebuild/code/plugin"
	pro "weibo.com/opendcp/imagebuild/code/project"
	"weibo.com/opendcp/imagebuild/code/service"
	"weibo.com/opendcp/imagebuild/code/util"
	log "github.com/Sirupsen/logrus"
)

var DefaultProjectName = "DefaultProjectName"
var specialStrings = []string{"!","@","#","$","%","^","&","*","(",")","=","'","\"","/","\\","|","<",">","{","}","[","]"}

type Server struct {
	// app version
	version string
	// app start time
	startTime string
	// ip
	ip string
	// port
	port string
	// lang
	Lang string

	// all internal plugins of procedure "dockerfile"
	dockerfilePlugins *util.ConcurrentMap

	// all internal plugins of procedure "build"
	buildPlugins *util.ConcurrentMap

	projects map[string]pro.Project

	projectLock sync.RWMutex
}

func (app *Server) Init(ip string, port string) {
	log.Info("------start init server")
	app.version = "v1.0"
	log.Info("app version v1.0")
	app.startTime = time.Now().String()
	log.Infof("start time is %s", app.startTime)

	// load all plugins: dockerfile plugins and build plugins
	log.Info("start load dockerfile plugins")
	app.loadDockerfilePlugins()
	log.Info("finish load dockerfile plugins")
	log.Info(app.dockerfilePlugins.ToPrettyString())

	// load all project
	log.Info("start load all projects")
	app.loadProjects()
	log.Info("finish load all projects")
	bytes, _ := json.MarshalIndent(app.projects, "", "  ")
	log.Info(string(bytes))

	log.Info("start create defaultProject")
	app.createDefaultProject()
	log.Info("finish create defaultProject")

	app.ip = ip
	app.port = port

	log.Info("------finish init server")
}

func (app *Server) CloneProject(srcCluster, srcProjectName, dstProjectName, creator, cluster, defineDockerFileType string) (bool, int) {
	project, code := pro.CloneProject(srcCluster, srcProjectName,
		dstProjectName,
		creator,
		cluster,
		defineDockerFileType,
		app.dockerfilePlugins,
		app.buildPlugins)

	if code != errors.OK {
		return false, code
	}

	// write lock
	app.projectLock.Lock()
	defer app.projectLock.Unlock()

	var dstWholeProjectName = app.getWholeProjectName(cluster, dstProjectName)
	app.projects[dstWholeProjectName] = project
	log.Infof("clone project: %s from project: %s success", dstProjectName, srcProjectName)
	return true, code
}

func (app *Server) IsProjectExist(cluster, projectName string) bool {

	projectWholeName := app.getWholeProjectName(cluster, projectName)
	var project pro.Project = app.getProject(projectWholeName)
	if project != nil {
		log.Infof("Project: %s is exist", projectName)
		return true
	}
	return false
}

func (app *Server) IsDefaultProjectExist(projectName string) bool {

	var project pro.Project = app.getProject(projectName)
	if project != nil {
		log.Infof("Project: %s is exist", projectName)
		return true
	}
	return false
}

func (app *Server) UpdateProject(projectName, creator, cluster, defineDockerFileType string) (bool, int) {
	// write lock
	app.projectLock.Lock()
	defer app.projectLock.Unlock()
	infoMap, code := pro.UpdateInfo(projectName, creator, cluster, defineDockerFileType)

	if code != errors.OK {
		return false, code
	}

	projectWholeName := app.getWholeProjectName(cluster, projectName)
	project := app.projects[projectWholeName]
	project.(*pro.PluggedProject).LastModifyOperator = infoMap["lastModifyOperator"]
	project.(*pro.PluggedProject).LastModifyTime = infoMap["lastModifyTime"]
	project.(*pro.PluggedProject).Creator = infoMap["creator"]
	project.(*pro.PluggedProject).CreateTime = infoMap["createTime"]
	project.(*pro.PluggedProject).DefineDockerFileType = infoMap["defineDockerFileType"]
	project.(*pro.PluggedProject).Cluster = infoMap["cluster"]

	app.projects[projectWholeName] = project
	log.Infof("update project: %s success", projectName)
	return true, code
}

func (app *Server) NewProject(projectName, creator, cluster, defineDockerFileType string) (bool, int) {
	project, code := pro.NewProject(projectName,
		creator,
		cluster,
		defineDockerFileType,
		app.dockerfilePlugins,
		app.buildPlugins)

	if code != errors.OK {
		return false, code
	}

	// write lock
	app.projectLock.Lock()
	defer app.projectLock.Unlock()

	var projectWholeName = app.getWholeProjectName(cluster, projectName)
	app.projects[projectWholeName] = project
	log.Infof("new project: %s success", projectWholeName)
	return true, code
}

func (app *Server) DeleteProject(cluster, projectName string, operator string) (bool, int) {
	// write lock
	app.projectLock.Lock()
	defer app.projectLock.Unlock()
	var projectWholeName = app.getWholeProjectName(cluster, projectName)
	if _, ok := app.projects[projectWholeName]; !ok {
		log.Errorf("project: %s to delete no exist", projectWholeName)
		return false, errors.DELETE_PROJECT_NOT_EXIST
	}

	code := pro.DeleteProject(cluster, projectName, operator)
	if code != errors.OK {
		return false, code
	}

	delete(app.projects, projectWholeName)

	log.Infof("delete project: %s success", projectName)
	return true, errors.OK
}

func (app *Server) SaveProjectConfig(cluster, projectName string, configs []map[string]interface{}) bool {
	projectWholeName := app.getWholeProjectName(cluster, projectName)
	var project pro.Project = app.getProject(projectWholeName)
	if project == nil {
		log.Errorf("Project: %s %s to save config not exist", cluster, projectName)
		return false
	}
	return project.Save(configs)
}

func (app *Server) GetProjectConfigView(cluster string, projectName string) (int, string) {
	projectWholeName := app.getWholeProjectName(cluster, projectName)
	var project pro.Project = app.getProject(projectWholeName)

	if project == nil {
		log.Infof("Project: %s %s config not exist", cluster, projectName)
		project = app.getProject(DefaultProjectName)
	}

	projectView := project.View(app.Lang)

	return errors.OK, projectView
}

func (app *Server) GetBuildExtensionPlugins() []string {
	buildExtensionPlugins := make([]string, 0)
	for plugin := range app.buildPlugins.Iterator() {
		buildExtensionPlugins = append(buildExtensionPlugins, plugin.Value.(*plu.PluginWrapper).Plugin_name)
	}

	return buildExtensionPlugins
}

func (app *Server) GetDockerfileExtensionPlugins() []string {
	dockerfileExtensionPlugins := make([]string, 0)
	for plugin := range app.dockerfilePlugins.Iterator() {
		dockerfileExtensionPlugins = append(dockerfileExtensionPlugins, plugin.Value.(*plu.PluginWrapper).Plugin_name)
	}

	return dockerfileExtensionPlugins
}

func (app *Server) GetProjectInfo(cluster string, projectName string) (int, pro.ProjectInfo) {
	projectWholeName := app.getWholeProjectName(cluster, projectName)
	var project pro.Project = app.getProject(projectWholeName)
	if project == nil {
		log.Errorf("Project: %s to query info not exist", projectName)
		return errors.PROJECT_NOT_EXIST, pro.BuildEmptyProjectInfo()
	}

	return errors.OK, project.Info()
}

func (app *Server) GetProjects(cluster, projectName string) pro.ProjectInfoList {
	app.projectLock.RLock()
	defer app.projectLock.RUnlock()

	projectInfos := make([]pro.ProjectInfo, 0)
	for _, project := range app.projects {
		projectInfo := pro.Project(project).Info()
		if projectInfo.Name != DefaultProjectName && strings.Contains(projectInfo.Name, projectName) &&
			strings.Compare(projectInfo.Cluster, cluster) == 0{
			projectInfos = append(projectInfos, projectInfo)
		}
	}

	// sort
	sort.Sort(pro.ProjectInfoList(projectInfos))

	return projectInfos
}

func (app *Server) BuildImage(cluster string, projectName, tag, operator string) (int, int64) {
	projectWholeName := app.getWholeProjectName(cluster, projectName)
	var project pro.Project = app.getProject(projectWholeName)
	if project == nil {
		log.Errorf("Project: %s to build not exist, operator", projectName, operator)
		return errors.BUILD_PROJECT_NOT_EXIST, -1
	}

	buildHistoryService := service.GetBuildHistoryServiceInstance()
	if buildHistoryService == nil {
		return errors.INTERNAL_ERROR, -1
	}

	id := buildHistoryService.InsertRecord(cluster, operator, projectName)

	// 异步线程处理构建并且进行更新任务状态
	go func() {
		//清空日志
		project.ClearLog()
		project.AppendLog(fmt.Sprintf("%s\t[%s]\t%s begin generate dockerfile",app.timeNow(),"Info",projectName))
		buildHistoryService.UpdateRecord(id, project.GetLog(), service.BUILDING)
		success := project.BuildImage()
		if success {
			log.Infof("%s build dockerfile success id:%d", projectName, id)
			project.AppendLog(fmt.Sprintf("%s\t[%s]\t%s build dockerfile success id:%d",app.timeNow(),"Info",projectName, id))
			buildHistoryService.UpdateRecord(id, project.GetLog(), service.BUILDING)
			log.Infof("start build and push image with project:%s state for build id:%d", projectName, id)
			pushSuccess := project.BuildAndPushImage(id, tag)
			if pushSuccess {
				log.Infof("%s push success id:%d tag:%s", projectName, id, tag)
				project.AppendLog(fmt.Sprintf("%s\t[%s]\t%s push success id:%d tag:%s",app.timeNow(),"Info", projectName, id, tag))
				pro.ClearTmp(cluster, projectName)
				if id != -1 {
					log.Infof("start update project %s state for id:%d", projectName, id)
					buildHistoryService.UpdateRecord(id, project.GetLog(), service.SUCCESS)
					log.Infof("finish update project %s state for id:%d", projectName, id)
				}
			} else {
				log.Errorf("%s push fail id:%d tag:%s", projectName, id, tag)
				if id != -1 {
					log.Infof("start update project %s state for build id:%d", projectName, id)
					buildHistoryService.UpdateRecord(id, project.GetLog(), service.FAIL)
					log.Infof("finish update project %s state for build id:%d", projectName, id)
				}
			}
		} else {
			log.Errorf("%s build fail id:%d", projectName, id)
			if id != -1 {
				project.AppendLog(fmt.Sprintf("%s\t[%s]\t%s build dockerfile failure id:%d",app.timeNow(),"Error",projectName, id))
				log.Infof("start update project %s state for build id:%d", projectName, id)
				buildHistoryService.UpdateRecord(id, project.GetLog(), service.FAIL)
				log.Infof("finish update project %s state for build id:%d", projectName, id)
			}
		}

		defer func() {
			if r := recover(); r != nil {
				log.Infof("start update project %s state for build id:%d", projectName, id)
				buildHistoryService.UpdateRecord(id, project.GetLog(), service.FAIL)
				log.Infof("finish update project %s state for build id:%d", projectName, id)
			}
		}()
	}()

	if id < 0 {
		return errors.INTERNAL_ERROR, id
	} else {
		return errors.OK, id
	}
}

func (app *Server) GetBuildLastHistory(cluster, projectName string) *model.BuildHistory {
	return service.GetBuildHistoryServiceInstance().QueryLastBuildRecord(cluster, projectName)
}

func (app *Server) GetBuildHistories(cursor int, offset int, cluster string, projectName string) []*model.BuildHistory {
	return service.GetBuildHistoryServiceInstance().QueryRecordList(cursor, offset, cluster, projectName)
}

func (app *Server) GetBuildHistory(id int) *model.BuildHistory {
	return service.GetBuildHistoryServiceInstance().QueryRecord(id)
}

func (app *Server) GetPluginView(pluginType int, pluginName string) string {
	var plugin interface{}
	if pluginType == plu.DOCKERFILE_PLUGIN {
		plugin = app.dockerfilePlugins.Get(pluginName)
	} else {
		plugin = app.buildPlugins.Get(pluginName)
	}

	if plugin == nil {
		return ""
	}
	return plugin.(*plu.PluginWrapper).View()
}

func (app *Server) GetServerAddress() string {
	if app.port == "" {
		return app.ip
	}

	return app.ip + ":" + app.port
}

func (app *Server) CallExtensionInterface(pluginName string, method string, params map[string]interface{}) (int, interface{}) {
	p := app.dockerfilePlugins.Get(pluginName)

	if p == nil {
		p = app.buildPlugins.Get(pluginName)
	}

	if p == nil {
		log.Errorf("plugin: %s not exist when calling extension interface: %s", pluginName, method)
		return errors.PLUGIN_NOT_EXIST, ""
	}

	plugin := p.(*plu.PluginWrapper)

	// 系统变量，插件可能会用到
	util.PackageSystemEnvIntoParam(params)
	realPluginName := util.ConvertToHump(pluginName+"_plugin")
	log.Infof("call %s.%s", realPluginName, method)
	var result interface{}
	error := plugin.Plugin.Call(realPluginName+"."+method, params, &result)
	if error != nil {
		log.Errorf("call %s.%s with error:%s", realPluginName, method, error)
		return errors.INTERNAL_ERROR, fmt.Sprintf("%s", error)
	}

	return errors.OK, result
}
//验证项目名称是否合法
func (app *Server) ValidateProjectName(projectName string) (bool, string){
	for _,spec := range specialStrings {
		if strings.Contains(projectName, spec) {
			log.Errorf("projectName: %s contains special char: %s", projectName, spec)
			return false, spec
		}
	}
	return true, ""
}
// ===================== private function ======================
func (app *Server) getProject(projectName string) pro.Project {
	// read lock
	app.projectLock.RLock()
	defer app.projectLock.RUnlock()

	if project, ok := app.projects[projectName]; ok {
		return project
	}

	return nil
}

func (app *Server) createDefaultProject() {
	app.projectLock.RLock()
	defer app.projectLock.RUnlock()

	exist := app.IsDefaultProjectExist(DefaultProjectName)
	if !exist {
		project, code := pro.NewProject(DefaultProjectName, "", "", "", app.dockerfilePlugins, app.buildPlugins)
		if code != errors.OK {
			log.Errorf("init failed, %s", code)
			return
		}
		app.projects[DefaultProjectName] = project
	}
}

func (app *Server) loadProjects() {
	app.projects = make(map[string]pro.Project, 0)

	fileInfos, error := ioutil.ReadDir(env.PROJECT_CONFIG_BASEDIR)
	if error != nil {
		error := os.Mkdir(env.PROJECT_CONFIG_BASEDIR, 0700)
		if error != nil {
			log.Errorf("create project dir with error:", error)
			os.Exit(-1)
		}
	}

	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			log.Errorf("%s is not a folder..", fileInfo.Name())
			continue
		}

		projectName := fileInfo.Name()

		project := pro.BuildPluginProject(projectName,
			"",
			"",
			"",
			"",
			app.dockerfilePlugins,
			app.buildPlugins)

		project.Init()

		log.Infof("project %s load success", projectName)
		app.projects[projectName] = project
	}
}

func (app *Server) loadDockerfilePlugins() {
	app.dockerfilePlugins = util.MakeConcurrentMap()
	app.loadPlugins(plu.DOCKERFILE_PLUGIN, env.DOCKERFILE_PLUGINS_BASEDIR, app.dockerfilePlugins)
}

func (app *Server) loadPlugins(pluginType int, basedir string, pluginWrappers *util.ConcurrentMap) {
	fileInfos, error := ioutil.ReadDir(basedir)
	if error != nil {
		log.Errorf("init failed, %s", error)
		os.Exit(-1)
	}

	for _, fileInfo := range fileInfos {
		// 如果不是文件夹则直接跳过
		if !fileInfo.IsDir() {
			log.Warning("%s is not a folder..", fileInfo.Name())
			continue
		}

		// 检查是否有可执行文件，没有则直接跳过
		executableFileExist := util.IsFileExists(basedir + "/" + fileInfo.Name() + "/" + fileInfo.Name() + "_plugin")
		if !executableFileExist {
			log.Warning("%s doesn't contain an executable file..", fileInfo.Name())
			continue
		}

		log.Infof("start load plugin %s", fileInfo.Name())

		pluginWrapper := plu.BuildPluginWrapper(fileInfo.Name(), basedir+"/"+fileInfo.Name())
		pluginWrapper.Plugin_type = pluginType
		pluginWrappers.Put(pluginWrapper.Plugin_name, pluginWrapper)

		log.Infof("plugin %s load success", fileInfo.Name())
	}

}

func (app *Server) loadNewPlugin(pluginType int, name string, path string) {
	pluginWrapper := plu.BuildPluginWrapper(name, path)
	pluginWrapper.Plugin_type = pluginType

	if pluginType == plu.BUILD_PLUGIN {
		app.buildPlugins.Put(name, pluginWrapper)
	} else {
		app.dockerfilePlugins.Put(name, pluginWrapper)
	}
}

//返回完整名称
func (app *Server) getWholeProjectName(cluster string, projectName string)(string){
	projectName = strings.ToLower(projectName)
	return cluster + "^" + projectName
}

//获取当前时间
func (app *Server) timeNow() string {
	return time.Now().Format("2006-01-02 15:04:05") + "\t"
}

