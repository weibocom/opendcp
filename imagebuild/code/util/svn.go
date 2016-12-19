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
