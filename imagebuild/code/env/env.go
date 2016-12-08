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

package env

var (
	SERVER_BASEDIR             string
	GLOBLE_CONFIG_BASEDIR      string
	PROJECT_CONFIG_BASEDIR     string
	DOCKERFILE_PLUGINS_BASEDIR string

	SERVER_HOST string

	MYSQL_HOST     string
	MYSQL_PORT     string
	MYSQL_USER     string
	MYSQL_PASSWORD string

	LOG_PATH string

	PLUGIN_VIEW_RUL         string
	EXTENSION_INTERFACE_RUL string

	HARBOR_ADDRESS  string
	HARBOR_USER     string
	HARBOR_PASSWORD string

	CLUSTERLIST_ADDRESS string
)

func InitEnv(
	harborAddress string,
	harborUser string,
	harborPassword string,
	pluginViewUrl string,
	extensionInterfaceUrl string,
	logPath string,
	serverIp string,
	serverPort string,
	serverBaseDir string,
	mysqlHost string,
	mysqlPort string,
	mysqlUser string,
	mysqlPassword string,
	clusterlistAddress string) {

	SERVER_BASEDIR = serverBaseDir + "/code/"
	PROJECT_CONFIG_BASEDIR = serverBaseDir + "/project/"
	DOCKERFILE_PLUGINS_BASEDIR = SERVER_BASEDIR + "plugins/dockerfile/"
	GLOBLE_CONFIG_BASEDIR = serverBaseDir + "/globle_config/"

	SERVER_HOST = serverIp + ":" + serverPort

	MYSQL_HOST = mysqlHost
	MYSQL_PORT = mysqlPort
	MYSQL_USER = mysqlUser
	MYSQL_PASSWORD = mysqlPassword

	LOG_PATH = logPath

	PLUGIN_VIEW_RUL = pluginViewUrl
	EXTENSION_INTERFACE_RUL = extensionInterfaceUrl

	HARBOR_ADDRESS = harborAddress
	HARBOR_USER = harborUser
	HARBOR_PASSWORD = harborPassword

	CLUSTERLIST_ADDRESS = clusterlistAddress
}
