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

	cluster := params["cluster"].(string)

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
	realCheckoutAs := projectPath + cluster + "/" + project + "/tmp/" + localPath

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
