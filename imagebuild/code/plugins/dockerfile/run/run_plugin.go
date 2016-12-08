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
	log "github.com/Sirupsen/logrus"
	"github.com/dullgiulio/pingo"
	"weibo.com/opendcp/imagebuild/code/service"
	"weibo.com/opendcp/imagebuild/code/util"
)

/**
run
 */
type RunPlugin struct{}

func (p *RunPlugin) Process(params map[string]interface{}, resp *string) error {
	var currentDockerfile string = params["input"].(string)
	config := make(map[string]interface{}, 0)
	error := json.Unmarshal([]byte(params["config"].(string)), &config)
	if error != nil {
		return error
	}

	logPath := "/tmp/imagebuild.log"
	util.LogInit(logPath)

	values := config["value"]
	log.Infof("....values:", values)

	var valueArray []string

	if util.IsArray(values) {
		valueArray = make([]string, 0)
		for _, v := range values.([]interface{}) {
			valueArray = append(valueArray, v.(string))
		}
	} else {
		valueArray = []string{values.(string)}
	}
	for _, run := range valueArray {
		log.Infof("...run:", run)
		currentDockerfile, error = service.GetDockerFileOperatorInstance().RunExec(currentDockerfile, run)
		if error != nil {
			return error
		}
	}

	*resp = currentDockerfile
	return nil
}

func main() {
	plugin := &RunPlugin{}
	pingo.Register(plugin)
	pingo.Run()
}
