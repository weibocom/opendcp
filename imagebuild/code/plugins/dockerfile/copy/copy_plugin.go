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
