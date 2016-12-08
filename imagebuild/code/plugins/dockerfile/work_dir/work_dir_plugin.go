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
)

/**
工作目录
 */
type WorkDirPlugin struct{}

func (p *WorkDirPlugin) Process(params map[string]interface{}, resp *string) error {
	var currentDockerfile string = params["input"].(string)
	config := make(map[string]interface{}, 0)
	json.Unmarshal([]byte(params["config"].(string)), &config)

	workDir := config["input_workDir"].(string)
	newDockerfile, error := service.GetDockerFileOperatorInstance().WorkDir(currentDockerfile, workDir)
	if error != nil {
		return error
	}
	*resp = newDockerfile
	return nil
}

func main() {
	plugin := &WorkDirPlugin{}
	pingo.Register(plugin)
	pingo.Run()
}
