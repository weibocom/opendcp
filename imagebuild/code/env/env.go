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
