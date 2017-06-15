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

func BuildExtensibleDockerFileGenerator(cluster string, projectName string,
	configRelativeFolder string,
	plugins *util.ConcurrentMap) *ExtensibleDockerFileGenerator {

	g := &ExtensibleDockerFileGenerator{
		plugins: plugins}

	g.Cluster = cluster
	g.ProjectName = projectName
	g.ConfigRelativeFolder = configRelativeFolder

	return g
}

func (generator *ExtensibleDockerFileGenerator) Init() bool {

	generator.pipeline = p.BuildPluginPipeline(
		generator.Cluster,
		generator.ProjectName,
		"Dockerfile Pipeline",
		"Dockerfile Pipeline")

	// 初始化pipeline
	generator.initPipeline(generator.Cluster, generator.ProjectName, generator.plugins)

	return true
}

func (builder *ExtensibleDockerFileGenerator) Handle() bool {
	dockerFile := ""

	err, tmp := builder.pipeline.Handle(builder.Cluster, builder.ProjectName, dockerFile)
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

	tmpPath := env.PROJECT_CONFIG_BASEDIR + builder.Cluster + "/" + builder.GetProjectName() + "/tmp"
	if !util.IsDirExists(tmpPath) {
		util.NewFile(env.PROJECT_CONFIG_BASEDIR + builder.Cluster + "/" + builder.GetProjectName(), "tmp", true)
	}

	dockerfilePath := env.PROJECT_CONFIG_BASEDIR + builder.Cluster + "/" + builder.GetProjectName() + "/tmp/Dockerfile"
	error := ioutil.WriteFile(dockerfilePath, []byte(dockerFile), 0777)
	if error != nil {
		log.Errorf("write dockerfile to disk error: %s\n", error)
		return false
	}

	return true
}
