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
		return command, ErrorWrapper(err)
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
