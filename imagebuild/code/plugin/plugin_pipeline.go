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
	log "github.com/Sirupsen/logrus"
	"weibo.com/opendcp/imagebuild/code/env"
	"weibo.com/opendcp/imagebuild/code/util"
	"github.com/astaxie/beego"
)

/**
由多个插件按照顺序构建而成的pipeline。前一个插件处理完任务后，将结果交给下一个
插件继续执行。。
*/
type PluginPipeline struct {
	//所属集群
	cluster string
	// 所属项目平成
	projectName string
	// pipeline名称
	pipelineName string
	// pipeline描述
	pipelineDescription string
	// 插件集合
	plugs []*PluginWrapper
}

func BuildPluginPipeline(cluster string, projectName string, pipelineName string, pipelineDescription string) *PluginPipeline {
	pipeline := &PluginPipeline{
		cluster: 	     cluster,
		projectName:         projectName,
		pipelineName:        pipelineName,
		pipelineDescription: pipelineDescription,
		plugs:               make([]*PluginWrapper, 0)}

	return pipeline
}

func BuildPluginPipelineWithPlugins(cluster string, projectName string, pipelineName string, pipelineDescription string, plugs []*PluginWrapper) *PluginPipeline {
	pipeline := &PluginPipeline{
		cluster:              cluster,
		projectName:         projectName,
		pipelineName:        pipelineName,
		pipelineDescription: pipelineDescription,
		plugs:               plugs}

	return pipeline
}

func (pp *PluginPipeline) GetCluster() string {
	return pp.cluster
}
func (pp *PluginPipeline) GetProjectName() string {
	return pp.projectName
}

func (pp *PluginPipeline) GetPipelineName() string {
	return pp.pipelineName
}

func (pp *PluginPipeline) GetPipelineDescription() string {
	return pp.pipelineDescription
}

func (pp *PluginPipeline) SetCluster(cluster string) {
	pp.cluster = cluster
}

func (pp *PluginPipeline) SetProjectName(projectName string) {
	pp.projectName = projectName
}

func (pp *PluginPipeline) SetPipelineName(pipelineName string) {
	pp.pipelineName = pipelineName
}

func (pp *PluginPipeline) AddPluginToTail(plugin *PluginWrapper) {
	pp.plugs = append(pp.plugs, plugin)
}

func (pp *PluginPipeline) AddPluginToHead(plugin *PluginWrapper) {
	pp.plugs = insert(pp.plugs, 0, plugin)
}

func (pp *PluginPipeline) InsertPlugin(plugin *PluginWrapper, i int) {
	pp.plugs = insert(pp.plugs, i, plugin)
}

func (pp *PluginPipeline) GetPlugin(pluginName string) *PluginWrapper {
	for _, plug := range pp.plugs {
		if plug.Plugin_name == pluginName {
			return plug
		}
	}

	return nil
}

func (pp *PluginPipeline) View() PluginPipelineView {
	pluginPipelineView := PluginPipelineView{}
	pluginPipelineView.ViewName = pp.pipelineName
	pluginPipelineView.PluginViews = make([]PluginView, 0)
	for _, plug := range pp.plugs {
		pluginPipelineView.PluginViews = append(pluginPipelineView.PluginViews, PluginView{PluginName: plug.Plugin_name,
			View: plug.View()})
	}

	return pluginPipelineView
}

func (pp *PluginPipeline) ClearAllConfig(relativeConfigFolder string) {
	configFolder := env.PROJECT_CONFIG_BASEDIR + pp.cluster + "/" + pp.projectName + "/" + relativeConfigFolder + "/"

	util.ClearFolder(configFolder)
}

func (pp *PluginPipeline) PluginList() string {
	var pluginList string = ""
	for _, plugin := range pp.plugs {
		pluginList += "," + plugin.Plugin_name
	}

	if len(pluginList) > 1 {
		return pluginList[1:]
	}

	return pluginList
}

func (pp *PluginPipeline) Handle(cluster string, project string, input interface{}) (error, interface{}) {

	beego.Warn("pp *PluginPipeline) Handle")

	var response interface{}
	in := input

	for inx, plug := range pp.plugs {
		beego.Warn("inx:%d  plug:%s", inx, plug.Plugin_name)
		log.Infof("inx:%d  plug:%s", inx, plug.Plugin_name)
	}
	for _, plug := range pp.plugs {
		beego.Warn("%s start process..", plug.Plugin_name)
		log.Infof("%s start process..", plug.Plugin_name)
		err, ret := plug.Process(cluster, project, in)
		if err != nil {
			return err, ""
		}

		in = ret
		response = ret
		log.Infof("%s finish process..", plug.Plugin_name)
		log.Infof("%s finish process..", plug.Plugin_name)
	}

	if response == nil {
		response = ""
	}

	return nil, response
}

func insert(slice []*PluginWrapper, index int, value *PluginWrapper) []*PluginWrapper {
	slice = slice[0 : len(slice)+1]
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}
