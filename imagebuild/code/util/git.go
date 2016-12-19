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
