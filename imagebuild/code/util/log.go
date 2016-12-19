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
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os"
)

/**
初始化log
 */
func LogInit(logPath string) {
	var logHandler *os.File

	logHandler, err := os.Open(logPath)
	if !IsFileExists(logPath) {
		logHandler, err = os.Create(logPath)
		if err != nil {
			fmt.Printf("create logfile fail, logpath: %s\n", logPath)
			os.Exit(-1)
		}
	} else {
		logHandler, err = os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Printf("open logfile fail, logpath: %s\n", logPath)
			os.Exit(-1)
		}
	}

	log.SetOutput(logHandler)
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
	log.Infof("Logpath: %s", logPath)
}
