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
	"strings"
	"weibo.com/opendcp/imagebuild/code/service"
	"weibo.com/opendcp/imagebuild/code/util"
)

/**
入口
 */
type EntrypointPlugin struct{}

func (p *EntrypointPlugin) Process(params map[string]interface{}, resp *string) error {
	var currentDockerfile string = params["input"].(string)
	config := make(map[string]interface{}, 0)
	json.Unmarshal([]byte(params["config"].(string)), &config)

	entrypoint := config["input_entrypoint"].(string)
	if util.IsEmpty(entrypoint) {
		*resp = currentDockerfile
		return nil
	}

	splits := strings.Split(entrypoint, " ")
	var entryCommand string
	var entryParams []string
	if len(splits) > 1 {
		entryCommand = splits[0]
		entryParams = splits[1:]
	} else {
		entryCommand = splits[0]
		entryParams = []string{}
	}

	newDockerfile, error := service.GetDockerFileOperatorInstance().EntrypointExec(currentDockerfile, entryCommand, entryParams...)
	if error != nil {
		return error
	}
	*resp = newDockerfile
	return nil
}

func main() {
	plugin := &EntrypointPlugin{}
	pingo.Register(plugin)
	pingo.Run()
}
