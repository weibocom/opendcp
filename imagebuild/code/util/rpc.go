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



package util

import "weibo.com/opendcp/imagebuild/code/env"

/**
系统变量，插件可能会用到
 */
func PackageSystemEnvIntoParam(params map[string]interface{}) {

	params["projectFolder"] = env.PROJECT_CONFIG_BASEDIR
	params["mysqlHost"] = env.MYSQL_HOST
	params["mysqlPort"] = env.MYSQL_PORT
	params["mysqlUser"] = env.MYSQL_USER
	params["mysqlPassword"] = env.MYSQL_PASSWORD
	params["dockerfilePluginFolder"] = env.DOCKERFILE_PLUGINS_BASEDIR
	params["logPath"] = env.LOG_PATH
	params["defaultHarborAddress"] = env.HARBOR_ADDRESS
	params["defaultHarborUser"] = env.HARBOR_USER
	params["defaultHarborPassword"] = env.HARBOR_PASSWORD
}
