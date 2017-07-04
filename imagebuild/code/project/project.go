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
	log "github.com/Sirupsen/logrus"
	"github.com/go-yaml/yaml"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"
	"weibo.com/opendcp/imagebuild/code/env"
	"weibo.com/opendcp/imagebuild/code/errors"
	"weibo.com/opendcp/imagebuild/code/util"
)

/**
项目信息
 */
type Project interface {
	Init() bool              // 初始化
	View(lang string) string // 返回project的页面数据
	Info() ProjectInfo
	BuildImage() bool                           // 构建镜像
	BuildAndPushImage(tag string, projectId int64) bool          // push镜像
	Save(configs []map[string]interface{}) bool // 保存配置

	ClearLog()
	AppendLog(line string, logLevel string)
	GetLog() string
}

func NewProject(projectName string,
	creator string,
	cluster string,
	defineDockerFileType string,
	dockerfilePlugins *util.ConcurrentMap,
	buildPlugins *util.ConcurrentMap) (project Project, code int) {

	//　项目目录
	if code := createProjectFolder(projectName); code != errors.OK {
		return nil, code
	}

	//　插件目录
	if error := createPluginConfigDirectory(projectName); error != errors.OK {
		return nil, error
	}

	// 临时文件夹
	if error := createTmpDirectory(projectName); error != errors.OK {
		return nil, error
	}

	// 初始化信息文件
	createTime, code := createInfoFile(projectName, creator, cluster, defineDockerFileType)
	log.Info("createInfoFile:", projectName, code)
	if code != errors.OK {
		return nil, code
	}

	// 插件列表
	if code := createPluginListFile(projectName); code != errors.OK {
		return nil, code
	}

	project = BuildPluginProject(projectName,
		creator,
		cluster,
		defineDockerFileType,
		createTime,
		dockerfilePlugins,
		buildPlugins)

	project.Init()

	return project, errors.OK
}

func UpdateInfo(projectName string,
	creator string,
	cluster string,
	defineDockerFileType string) (map[string]string, int) {
	return updateInfoFile(projectName, creator, cluster, defineDockerFileType)
}

func DeleteProject(projectName string, operator string) (code int) {
	//　删除项目
	projectPath := env.PROJECT_CONFIG_BASEDIR + projectName
	if !util.DeleteFile(projectPath) {
		log.Errorf("delete project: %s error", projectName)
		return errors.INTERNAL_ERROR
	}

	return errors.OK
}

func ClearTmp(projectName string)  int {
	projectTmpPath := env.PROJECT_CONFIG_BASEDIR + projectName + "/tmp"
	util.ClearFolder(projectTmpPath)

	return errors.OK
}

func handleCloneError(projectName string) (project Project, code int) {
	DeleteProject(projectName, "system")
	return nil, errors.INTERNAL_ERROR
}

func CloneProject(srcProjectName string,
	dstProjectName string,
	creator string,
	cluster string,
	defineDockerFileType string,
	dockerfilePlugins *util.ConcurrentMap,
	buildPlugins *util.ConcurrentMap) (project Project, code int) {

	// 检查源项目是否存在
	srcProjectPath := env.PROJECT_CONFIG_BASEDIR + srcProjectName
	exists := util.IsDirExists(srcProjectPath)
	if !exists {
		log.Errorf("src project: %s not exist", srcProjectName)
		return nil, errors.CLONE_SRC_PROJECT_NOT_EXIST
	}

	//　创建项目目录
	if code := createProjectFolder(dstProjectName); code != errors.OK {
		return nil, code
	}

	dstProjectPath := env.PROJECT_CONFIG_BASEDIR + dstProjectName

	// 插件目录
	if !util.NewFile(dstProjectPath, "dockerfile", true) {
		log.Errorf("dockerfile plug folder create fail, src project: %s, dst project: %s", srcProjectName, dstProjectName)
		return handleCloneError(dstProjectName)
	}

	if !util.NewFile(dstProjectPath, "build", true) {
		log.Errorf("build plug folder create fail, src project: %s, dst project: %s", srcProjectName, dstProjectName)
		return handleCloneError(dstProjectName)
	}

	if !copyFolder(srcProjectPath+"/"+"dockerfile", dstProjectPath+"/"+"dockerfile") {
		log.Errorf("copy dockerfile plugin config error, src project: %s, dst project: %s", srcProjectName, dstProjectName)
		return handleCloneError(dstProjectName)
	}

	if !copyFolder(srcProjectPath+"/"+"build", dstProjectPath+"/"+"build") {
		log.Errorf("copy build plugin config error, src project: %s, dst project: %s", srcProjectName, dstProjectName)
		return handleCloneError(dstProjectName)
	}

	if !util.NewFile(dstProjectPath, "tmp", true) {
		log.Errorf("create tmp folder error, src project: %s, dst project: %s", srcProjectName, dstProjectName)
		return handleCloneError(dstProjectName)
	}

	createTime, code := createInfoFile(dstProjectName, creator, cluster, defineDockerFileType)
	if code != errors.OK {
		log.Errorf("create tmp folder error, src project: %s, dst project: %s", srcProjectName, dstProjectName)
		return nil, code
	}

	project = BuildPluginProject(dstProjectName,
		creator,
		cluster,
		defineDockerFileType,
		createTime,
		dockerfilePlugins,
		buildPlugins)

	project.Init()

	return project, errors.OK
}

// 深度为１的文件夹拷贝
func copyFolder(srcFolder string, dstFolder string) bool {
	file, error := os.Open(srcFolder)
	if error != nil {
		log.Errorf("%s", error)
		return false
	}

	defer file.Close()

	configNames, error := file.Readdirnames(-1)
	if error != nil {
		log.Errorf("%s", error)
		return false
	}

	for _, configName := range configNames {
		suc := util.NewFile(dstFolder, configName, false)
		if !suc {
			return false
		}

		in, error := os.Open(srcFolder + "/" + configName)
		if error != nil {
			log.Errorf("%s", error)
			return false
		}

		out, error := os.Create(dstFolder + "/" + configName)
		if error != nil {
			log.Errorf("%s", error)
			return false
		}

		_, error = io.Copy(out, in)

		// 直接关闭，而不是调用defer延迟关闭
		in.Close()
		out.Close()

		if error != nil {
			log.Errorf("%s", error)
			return false
		}
	}

	return true
}

func createProjectFolder(projectName string) (code int) {
	projectPath := env.PROJECT_CONFIG_BASEDIR + projectName
	exists := util.IsDirExists(projectPath)
	if exists {
		log.Errorf("project %s already exist", projectName)
		return errors.CREATE_PROJECT_ALREADY_EXIST
	}

	suc := util.NewFile(env.PROJECT_CONFIG_BASEDIR, projectName, true)
	if !suc {
		log.Errorf("project %s folder create fail", projectName)
		return errors.INTERNAL_ERROR
	}

	return errors.OK
}

func createInfoFile(projectName, creator, cluster, defineDockerFileType string) (createTime string, code int) {
	infoMap := make(map[string]string, 0)
	createTime = time.Now().String()
	infoMap["createTime"] = createTime
	infoMap["lastModifyTime"] = createTime
	infoMap["creator"] = creator
	infoMap["lastModifyOperator"] = creator
	infoMap["cluster"] = cluster
	infoMap["defineDockerFileType"] = defineDockerFileType

	infoBytes, error := yaml.Marshal(infoMap)
	if error != nil {
		return "", errors.INTERNAL_ERROR
	}

	code = writeDataToInfo(projectName, infoBytes)
	if code == errors.INTERNAL_ERROR {
		return "", code
	}

	return createTime, code
}

func updateInfoFile(projectName, creator, cluster, defineDockerFileType string) (map[string]string, int) {
	infoMap := make(map[string]string, 0)

	project := &PluggedProject{}
	project.Name = projectName
	project.readInfo()

	updateTime := time.Now().String()
	infoMap["creator"] = project.Creator
	infoMap["createTime"] = project.CreateTime
	infoMap["lastModifyTime"] = updateTime
	infoMap["lastModifyOperator"] = creator
	infoMap["cluster"] = cluster
	infoMap["defineDockerFileType"] = defineDockerFileType

	infoBytes, error := yaml.Marshal(infoMap)
	if error != nil {
		return nil, errors.INTERNAL_ERROR
	}

	code := writeDataToInfo(projectName, infoBytes)
	if code == errors.INTERNAL_ERROR {
		return nil, code
	}

	return infoMap, code
}

func writeDataToInfo(projectName string, infoBytes []byte) (code int) {
	projectPath := env.PROJECT_CONFIG_BASEDIR + projectName

	error := ioutil.WriteFile(projectPath+"/info", infoBytes, 0777)
	if error != nil {
		log.Errorf("create info file error, project: %s, error: %s", projectName, error)
		return errors.INTERNAL_ERROR
	}

	code = errors.OK
	return code
}

func createPluginConfigDirectory(projectName string) (code int) {
	projectPath := env.PROJECT_CONFIG_BASEDIR + projectName

	if !util.NewFile(projectPath, "build", true) {
		return errors.INTERNAL_ERROR
	}

	if !util.NewFile(projectPath, "dockerfile", true) {
		return errors.INTERNAL_ERROR
	}

	return errors.OK
}

func createTmpDirectory(projectName string) (code int) {
	projectPath := env.PROJECT_CONFIG_BASEDIR + projectName

	if !util.NewFile(projectPath, "tmp", true) {
		return errors.INTERNAL_ERROR
	}

	return errors.OK
}

func createProjectTypeFile(projectName string, projectType int) (code int) {
	extensionFile := env.PROJECT_CONFIG_BASEDIR + projectName + "/type"

	projectTypeStr := strconv.Itoa(projectType)

	err := ioutil.WriteFile(extensionFile, []byte(projectTypeStr), 0777)

	if err != nil {
		return errors.INTERNAL_ERROR
	}

	return errors.OK
}

func createPluginListFile(projectName string) (code int) {
	dockerfilePlugListFile := env.PROJECT_CONFIG_BASEDIR + projectName + "/dockerfile/plug_list"
	buildPlugListFile := env.PROJECT_CONFIG_BASEDIR + projectName + "/build/plug_list"

	err := ioutil.WriteFile(dockerfilePlugListFile, []byte(""), 0777)

	if err != nil {
		return errors.INTERNAL_ERROR
	}

	err = ioutil.WriteFile(buildPlugListFile, []byte(""), 0777)

	if err != nil {
		return errors.INTERNAL_ERROR
	}

	return errors.OK
}
