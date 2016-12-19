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



package main

import (
	"encoding/json"
	"github.com/dullgiulio/pingo"
	"weibo.com/opendcp/imagebuild/code/service"
	"weibo.com/opendcp/imagebuild/code/util"
)

/**
copy
 */
type CopyPlugin struct{}

func (p *CopyPlugin) Process(params map[string]interface{}, resp *string) error {
	var currentDockerfile string = params["input"].(string)
	config := make(map[string]interface{}, 0)
	error := json.Unmarshal([]byte(params["config"].(string)), &config)
	if error != nil {
		return error
	}

	keys := config["key"]
	values := config["value"]

	var keyArray []string
	var valueArray []string

	if util.IsArray(keys) {
		keyArray = make([]string, 0)
		for _, v := range keys.([]interface{}) {
			keyArray = append(keyArray, v.(string))
		}
	} else {
		keyArray = []string{keys.(string)}
	}

	if util.IsArray(values) {
		valueArray = make([]string, 0)
		for _, v := range values.([]interface{}) {
			valueArray = append(valueArray, v.(string))
		}
	} else {
		valueArray = []string{values.(string)}
	}

	keyValuePairs := make([]string, len(keyArray)*2)
	for index, key := range keyArray {
		keyValuePairs[index*2] = key
		keyValuePairs[index*2+1] = valueArray[index]
	}
	newDockerfile, error := service.GetDockerFileOperatorInstance().Env(currentDockerfile, keyValuePairs...)
	if error != nil {
		return error
	}

	*resp = newDockerfile
	return nil
}

func main() {
	plugin := &CopyPlugin{}
	pingo.Register(plugin)
	pingo.Run()
}
