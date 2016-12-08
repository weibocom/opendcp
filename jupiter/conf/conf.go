// Copyright 2016 Weibo Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	Password   string
	KeyId      string
	KeySecret  string
	BufferSize int
	Ansible    *Ansible
	KeyDir     string
}

type Ansible struct {
	Url         string
	DefaultRole string
	ForkNum     int
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
