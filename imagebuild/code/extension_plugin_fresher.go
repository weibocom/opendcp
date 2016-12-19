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
