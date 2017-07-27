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

package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	ConfigFile    = "conf/jupiter.json"
	TryTime       = 3
	ERROR_PARAM   = "parameter error"
	ERROR_CONVERT = "Convert error"
	Interval      = 60
	Time4Wait     = 3                   //unit s
	DateForm      = "2006-01-02T15:04Z" //时间格式
)

var (
	Config *Configuration
)

type Configuration struct {
	Password     string
	KeyId        string
	KeySecret    string
	AwsKeyId     string
	AwsKeySecret string
	OpIp         string
	OpPort       string
	OpUserName   string
	OpPassWord   string
	BufferSize   int
	Ansible      *Ansible
	KeyDir       string
}

type Ansible struct {
	Url          string
	GetOctansUrl string
	DefaultRole  string
	ForkNum      int
}

func GetConfig() (*Configuration, error) {
	c, err := getConfigFromFile()
	return c, err
}

func getConfigFromFile() (*Configuration, error) {
	var config Configuration
	if conf, err := ioutil.ReadFile(ConfigFile); err == nil {
		e := json.Unmarshal(conf, &config)
		return &config, e
	} else {
		return &config, err
	}
}

func FileExists(filename string) bool {
	fi, err := os.Stat(filename)
	return (err == nil || os.IsExist(err)) && !fi.IsDir()
}

func InitConf() {
	if config, err := GetConfig(); err == nil {
		Config = config
	}
	return
}
