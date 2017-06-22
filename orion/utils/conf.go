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

package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	ConfigFile = "conf/orion_dbinit.json"
)

var (
	InitConfig *InitConfiguration
)

type Record map[string]interface{}

type InitConfiguration []struct {
	Table   string
	Records []string
}

func GetConfig() (*InitConfiguration, error) {
	c, err := getConfigFromFile()
	return c, err
}

func getConfigFromFile() (*InitConfiguration, error) {
	var config InitConfiguration
	if conf, err := ioutil.ReadFile(ConfigFile); err == nil {
		//fmt.Println(string(conf))
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
	config, err := GetConfig()
	if err == nil {
		InitConfig = config
	}

	fmt.Println(err)
	return
}

//func main() {
//	InitConf()
//	fmt.Println(InitConfig)
//	fmt.Println("aaa")
//
//
//
//}
