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
