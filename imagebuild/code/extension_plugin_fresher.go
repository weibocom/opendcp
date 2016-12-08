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

package code

import (
	"io/ioutil"
	"weibo.com/opendcp/imagebuild/code/env"
	p "weibo.com/opendcp/imagebuild/code/plugin"
	"weibo.com/opendcp/imagebuild/code/util"
)

func FreshDockerfilePlugins(server *Server) {
	defer func() {
		if err := recover(); err != nil {
			go FreshDockerfilePlugins(server)
		}
		// recover
	}()

	plugins, _ := ioutil.ReadDir(env.DOCKERFILE_PLUGINS_BASEDIR)

	for _, plugin := range plugins {
		// contains executable plugin
		if !util.IsFileExists(env.DOCKERFILE_PLUGINS_BASEDIR + "/" + plugin.Name() + "/" + plugin.Name() + "_plugin") {
			continue
		}

		if !server.buildPlugins.ContainsKey(plugin.Name()) {
			continue
		}

		// load plugin
		server.loadNewPlugin(p.DOCKERFILE_PLUGIN, plugin.Name(), env.DOCKERFILE_PLUGINS_BASEDIR+"/"+plugin.Name())
	}
}
