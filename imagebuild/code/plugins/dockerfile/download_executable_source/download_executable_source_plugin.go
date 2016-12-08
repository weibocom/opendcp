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
	"fmt"
	"github.com/dullgiulio/pingo"
	"github.com/wendal/errors"
	"strconv"
	"strings"
	"time"
	"weibo.com/opendcp/imagebuild/code/service"
	"weibo.com/opendcp/imagebuild/code/util"
)

/**
下载部署文件
 */
type DownloadExecutableSourcePlugin struct{}

var SVN = "svn"

func (p *DownloadExecutableSourcePlugin) Process(params map[string]interface{}, resp *string) error {
	project := params["project"].(string)

	projectPath := params["projectFolder"].(string)

	currentDockerfile := params["input"].(string)

	config := make(map[string]interface{}, 0)
	json.Unmarshal([]byte(params["config"].(string)), &config)

	sourceType := config["source_type"].(string)
	sourceUrl := config["source_url"].(string)
	checkoutAs := config["checkoutAs"].(string)
	containerPath := config["containerPath"].(string)
	username := config["username"].(string)
	password := config["password"].(string)

	// do nothing
	if sourceUrl == "" {
		*resp = currentDockerfile
		errors.New("")
		return nil
	}

	if !strings.HasPrefix(checkoutAs, "/") {
		checkoutAs = "/" + checkoutAs
	}

	var localPath string = "localPath" + strconv.Itoa(time.Now().Nanosecond())
	realCheckoutAs := projectPath + project + "/tmp/" + localPath

	if sourceType == SVN {
		fmt.Printf("svn download")
		util.SvnDownload(sourceUrl, username, password, realCheckoutAs, project)
	} else {
		fmt.Printf("git download")
		util.GitDownload(sourceUrl, username, password, realCheckoutAs, project)
	}

	newDockerfile, error := service.GetDockerFileOperatorInstance().Add(currentDockerfile, containerPath, localPath+checkoutAs)
	if error != nil {
		return error
	}
	*resp = newDockerfile
	return nil
}

func main() {
	plugin := &DownloadExecutableSourcePlugin{}
	pingo.Register(plugin)
	pingo.Run()
}
