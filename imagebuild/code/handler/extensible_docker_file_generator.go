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

package handler

import (
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"strings"
	"weibo.com/opendcp/imagebuild/code/env"
	p "weibo.com/opendcp/imagebuild/code/plugin"
	"weibo.com/opendcp/imagebuild/code/util"
)

/**
可扩展的dockerfile生成器，ExtensibleBasicHandler的子类。
*/
type ExtensibleDockerFileGenerator struct {
	AbstractExtensibleHandler
	plugins *util.ConcurrentMap
}

func BuildExtensibleDockerFileGenerator(projectName string,
	configRelativeFolder string,
	plugins *util.ConcurrentMap) *ExtensibleDockerFileGenerator {

	g := &ExtensibleDockerFileGenerator{
		plugins: plugins}

	g.ProjectName = projectName
	g.ConfigRelativeFolder = configRelativeFolder

	return g
}

func (generator *ExtensibleDockerFileGenerator) Init() bool {

	generator.pipeline = p.BuildPluginPipeline(
		generator.ProjectName,
		"Dockerfile Pipeline",
		"Dockerfile Pipeline")

	// 初始化pipeline
	generator.initPipeline(generator.ProjectName, generator.plugins)

	return true
}

func (builder *ExtensibleDockerFileGenerator) Handle() bool {
	dockerFile := ""

	err, tmp := builder.pipeline.Handle(builder.ProjectName, dockerFile)
	if err != nil {
		log.Errorf("build dockerfile pipeline error")
		return false
	}

	dockerFile = tmp.(string)

	//sort
	dcmdList := strings.Split(dockerFile, "\n")
	fromList := make([]string, 0)
	addList := make([]string, 0)
	otherList := make([]string, 0)
	for _, dcmdItem := range dcmdList {
		dcmd := strings.Split(dcmdItem, " ")[0]
		switch strings.TrimSpace(dcmd) {
		default:
			otherList = append(otherList, dcmdItem)
		case "FROM":
			fromList = append(fromList, dcmdItem)
		case "ADD":
			addList = append(addList, dcmdItem)
		}
	}
	result := ""
	changLine := "\n"
	for _, cmd := range fromList {
		result = result + changLine + cmd
	}
	for _, cmd := range addList {
		result = result + changLine + cmd
	}
	for _, cmd := range otherList {
		result = result + changLine + cmd
	}

	dockerFile = result

	log.Infof("final dockerfile is: \n %s \n", dockerFile)

	if dockerFile == "" {
		return true
	}

	tmpPath := env.PROJECT_CONFIG_BASEDIR + builder.GetProjectName() + "/tmp"
	if !util.IsDirExists(tmpPath) {
		util.NewFile(env.PROJECT_CONFIG_BASEDIR+builder.GetProjectName(), "tmp", true)
	}

	dockerfilePath := env.PROJECT_CONFIG_BASEDIR + builder.GetProjectName() + "/tmp/Dockerfile"
	error := ioutil.WriteFile(dockerfilePath, []byte(dockerFile), 0777)
	if error != nil {
		log.Errorf("write dockerfile to disk error: %s\n", error)
		return false
	}

	return true
}
