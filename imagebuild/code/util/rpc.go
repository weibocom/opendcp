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
