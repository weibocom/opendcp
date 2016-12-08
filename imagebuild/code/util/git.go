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

package util

import (
	log "github.com/Sirupsen/logrus"
	"github.com/go-errors/errors"
	"os"
	"strings"
	"net/url"
)

/**
git工具
 */
func GitDownload(remoteUrl string, username string, password string, checkoutAs string, project string) *errors.Error {
	var realRemoteUrl string

	username = url.QueryEscape(username)
	password = url.QueryEscape(password)

	if strings.HasPrefix(remoteUrl, "http://") {
		realRemoteUrl = "http://" + username + ":" + password + "@" + remoteUrl[7:]
	} else if strings.HasPrefix(remoteUrl, "https://") {
		realRemoteUrl = "https://" + username + ":" + password + "@" + remoteUrl[8:]
	} else {
		realRemoteUrl = "http://" + username + ":" + password + "@" + remoteUrl
	}

	error := os.RemoveAll(checkoutAs)
	if error != nil {
		PrintErrorStack(error)
		return ErrorWrapper(error)
	}

	LogInit("/tmp/imagebuild.log")
	command := "git clone " + realRemoteUrl + " " + checkoutAs
	log.Infof("git command:%s", command)
	// exec
	_, error = ExecuteFullCommand(command)
	if error != nil {
		log.Errorf("git command error:%s", command)
		PrintErrorStack(error)
		return ErrorWrapper(error)
	}

	return nil
}
