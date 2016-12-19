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
