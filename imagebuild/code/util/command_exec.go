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
	"bytes"
	log "github.com/Sirupsen/logrus"
	"os/exec"
	"strings"
)

/**
 * eg: ExecuteCommand("ls", "-a", "-l")
 */
func ExecuteCommand(command string, params ...string) (string, error) {
	cmd := exec.Command(command, params...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Errorf("Command: %s fail, Output: %s", command, out.String())
		return out.String(), ErrorWrapper(err)
	}

	log.Infof("Command: %s success, Output: %s", command, out.String())

	return out.String(), nil
}

/*
 * eg: ExecuteCommand("ls -a -l")
 * you shouldn't use this function while any param of the command has blank.
 * eg: ExecuteCommand("echo 'abc def'")
 */
func ExecuteFullCommand(fullCommand string) (string, error) {
	commandSplits := strings.Split(fullCommand, " ")
	return ExecuteCommand(commandSplits[0], commandSplits[1:]...)
}
