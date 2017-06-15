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
	"errors"
	"github.com/dullgiulio/pingo"
	stackError "github.com/go-errors/errors"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
	"weibo.com/opendcp/imagebuild/code/service"
	"weibo.com/opendcp/imagebuild/code/util"
)

/**
下载dockerfile
 */
type DownloadDockerfilePlugin struct {
}

func (p *DownloadDockerfilePlugin) Process(params map[string]interface{}, resp *string) error {
	var currentDockerfile string = params["input"].(string)
	var project string = params["project"].(string)
	var projectPath string = params["projectFolder"].(string)
	config := make(map[string]interface{}, 0)
	error := json.Unmarshal([]byte(params["config"].(string)), &config)

	if error != nil {
		return util.ErrorWrapper(error)
	}

	sourceType := config["sourceType"].(string)
	sourceUrl := config["sourceUrl"].(string)
	checkoutAs := config["checkoutAs"].(string)
	username := config["username"].(string)
	password := config["password"].(string)

	var localPath string = "localPath" + strconv.Itoa(time.Now().Nanosecond())
	realCheckoutAs := projectPath + project + "/tmp/" + localPath

	if sourceType == "git" {
		error := util.GitDownload(sourceUrl, username, password, realCheckoutAs, project)
		if error != nil {
			return errors.New(error.ErrorStack())
		}
	} else {
		error := util.SvnDownload(sourceUrl, username, password, realCheckoutAs, project)
		if error != nil {
			return errors.New(error.ErrorStack())
		}
	}

	if !strings.HasPrefix(checkoutAs, "/") {
		checkoutAs = "/" + checkoutAs
	}
	content, error := ioutil.ReadFile(realCheckoutAs + checkoutAs)
	if error != nil {
		return error
	}

	newDockerfile, error := service.GetDockerFileOperatorInstance().DockerfileContent(currentDockerfile, string(content))
	if error != nil {
		return error
	}

	*resp = newDockerfile
	return nil
}

func (p *DownloadDockerfilePlugin) generateRandomTag(projectName string) (*stackError.Error, string) {
	re, error := regexp.Compile("[ \\-\\:]")
	if error != nil {
		return util.ErrorWrapper(error), ""
	}

	time := re.ReplaceAllString(time.Now().Format("2006-01-02 15:04:05"), "_")
	return nil, time

}

func main() {
	// &NYPizzaStore{pizzaStore: pizzaStore{new(NYPizzaStore)}}
	plugin := &DownloadDockerfilePlugin{}
	pingo.Register(plugin)
	pingo.Run()
}
