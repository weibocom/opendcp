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
)

/**
svn工具
 */
func SvnDownload(remoteUrl string, username string, password string, checkoutAs string, project string) *errors.Error {

	command := "svn co --username=" + username + " --password=" + password + " --trust-server-cert --non-interactive --no-auth-cache " + remoteUrl + " " + checkoutAs

	LogInit("/tmp/imagebuild.log")

	log.Infof("svn command:%s", command)
	// exec
	_, error := ExecuteFullCommand(command)
	if error != nil {
		log.Errorf("svn command error:%s", error)
		PrintErrorStack(error)
		return ErrorWrapper(error)
	}

	return nil
}
