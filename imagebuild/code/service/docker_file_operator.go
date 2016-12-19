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



package service

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"weibo.com/opendcp/imagebuild/code/util"
)

/**
组装Dockerfile
 */
type DockerFileOperator struct {
}

var instance_1 *DockerFileOperator
var once_1 sync.Once

// single instance
func GetDockerFileOperatorInstance() *DockerFileOperator {
	once_1.Do(func() {
		instance_1 = &DockerFileOperator{}
	})

	return instance_1
}

func (dfo *DockerFileOperator) DockerfileContent(currentDockerfile, dockerfileContent string) (string, error) {
	if util.IsEmpty(dockerfileContent) {
		return currentDockerfile, errors.New("Add From Instrument Error, CurrentDockerfile is: " + currentDockerfile)
	}
	currentDockerfile = currentDockerfile + "\n" + dockerfileContent
	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) From(currentDockerfile string, baseImage string) (string, error) {
	if util.IsEmpty(baseImage) {
		return currentDockerfile, errors.New("Add From Instrument Error, CurrentDockerfile is: " + currentDockerfile)
	}
	currentDockerfile = currentDockerfile + "\n" + "FROM " + baseImage
	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) Maintainer(currentDockerfile string, author string) (string, error) {
	if util.IsEmpty(author) {
		return currentDockerfile, errors.New("Add Maintainer Instrument Error")
	}
	currentDockerfile = currentDockerfile + "\n" + "MAINTAINER " + author
	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) WorkDir(currentDockerfile string, workDir string) (string, error) {
	if util.IsEmpty(workDir) {
		return currentDockerfile, errors.New("Add workDir Instrument Error")
	}
	currentDockerfile = currentDockerfile + "\n" + "WORKDIR " + workDir
	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) RunShell(currentDockerfile string, command string) (string, error) {
	if util.IsEmpty(command) {
		return currentDockerfile, errors.New("Add Run Instrument Error")
	}
	currentDockerfile = currentDockerfile + "\n" + "RUN " + command
	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) RunExec(currentDockerfile string, params string) (string, error) {
	if util.IsEmpty(params) {
		return currentDockerfile, errors.New("Add Run Instrument Error")
	}
	currentDockerfile = currentDockerfile + "\n" + "RUN " + "["

	runParams := strings.Split(params, " ")
	for inx, param := range runParams {
		if param != ""{
			currentDockerfile = currentDockerfile + " \"" + param + "\""
			if inx != len(runParams)-1 {
				currentDockerfile = currentDockerfile + ","
			}
		}
	}

	currentDockerfile = currentDockerfile + "]"
	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) CmdShell(currentDockerfile string, command string, params ...string) (string, error) {
	if util.IsEmpty(command) {
		return currentDockerfile, errors.New("Add Cmd Instrument Error")
	}
	currentDockerfile = currentDockerfile + "\n" + "CMD " + command
	for _, param := range params {
		currentDockerfile = currentDockerfile + " " + param
	}
	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) CmdExec(currentDockerfile string, command string, params ...string) (string, error) {
	var tmp []string

	if util.IsEmpty(command) {
		if len(params) == 0 {
			return currentDockerfile, errors.New("Add Cmd Instrument Error")
		}
		currentDockerfile = currentDockerfile + "\n" + "CMD" + "[" + "\"" + params[0] + "\""
		if len(params) == 1 {
			tmp = make([]string, 0)
		} else {
			tmp = params[1:]
		}
	} else {
		currentDockerfile = currentDockerfile + "\n" + "CMD " + "[" + "\"" + command + "\""
		tmp = params
	}

	for _, param := range tmp {
		currentDockerfile = currentDockerfile + ", \"" + param + "\""
	}

	currentDockerfile = currentDockerfile + "]"
	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) Label(currentDockerfile string, labels ...string) (string, error) {
	if len(labels) == 0 || len(labels)%2 != 0 {
		return currentDockerfile, errors.New("Add Label Instrument Error")
	}

	currentDockerfile = currentDockerfile + "\n" + "LABEL"
	for i := 0; i < len(labels); i = i + 2 {
		currentDockerfile = currentDockerfile + " " + labels[i] + "=" + "\"" + labels[i+1] + "\""
	}

	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) Expose(currentDockerfile string, ports ...string) (string, error) {
	if len(ports) == 0 {
		errorMsg := fmt.Sprintf("Add Expose Instrument Error! \ncurrentDockerfile is:\n%s\n expose is:%s\n", currentDockerfile, ports)
		return currentDockerfile, errors.New(errorMsg)
	}

	currentDockerfile = currentDockerfile + "\n" + "EXPOSE"
	for _, port := range ports {
		currentDockerfile += (" " + port)
	}

	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) Env(currentDockerfile string, envs ...string) (string, error) {
	if len(envs) == 0 || len(envs)%2 != 0 {
		errorMsg := fmt.Sprintf("Add Env Instrument Error! \ncurrentDockerfile is:\n%s\n env is:%s\n", currentDockerfile, envs)
		return currentDockerfile, errors.New(errorMsg)
	}

	for i := 0; i < len(envs); i = i + 2 {
		currentDockerfile += ("\n" + "ENV" + " " + envs[i] + " " + envs[i+1])
	}

	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) Add(currentDockerfile string, dstPath string, srcPaths ...string) (string, error) {
	if len(srcPaths) == 0 {
		return currentDockerfile, errors.New("Add Add Instrument Error")
	}

	containsWhitespace := false
	for _, srcPath := range srcPaths {
		if strings.Contains(srcPath, " ") {
			containsWhitespace = true
		}
	}

	if strings.Contains(dstPath, " ") {
		containsWhitespace = true
	}

	currentDockerfile += ("\n" + "ADD ")

	if containsWhitespace {
		currentDockerfile += ("[\"" + strings.Join(srcPaths, "\",\"") + "\" \"" + dstPath + "\"]")
	} else {
		currentDockerfile += (strings.Join(srcPaths, ",") + " " + dstPath)
	}

	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) Copy(currentDockerfile string, dstPath string, srcPaths ...string) (string, error) {
	if len(srcPaths) == 0 {
		return currentDockerfile, errors.New("Add Copy Instrument Error")
	}

	containsWhitespace := false
	for _, srcPath := range srcPaths {
		if strings.Contains(srcPath, " ") {
			containsWhitespace = true
		}
	}

	if strings.Contains(dstPath, " ") {
		containsWhitespace = true
	}

	currentDockerfile = currentDockerfile + "\n" + "ADD "

	if containsWhitespace {
		currentDockerfile += ("[\"" + strings.Join(srcPaths, "\",\"") + "\" \"" + dstPath + "\"]")
	} else {
		currentDockerfile += (strings.Join(srcPaths, ",") + " " + dstPath)
	}

	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) EntrypointExec(currentDockerfile string, command string, params ...string) (string, error) {
	if util.IsEmpty(command) {
		return currentDockerfile, errors.New("Add Entrypoint Instrument Error")
	}

	currentDockerfile += ("\n" + "ENTRYPOINT [\"" + command)

	paramsString := strings.Join(params, "\", \"")
	if paramsString == "" {
		currentDockerfile += "\"]"
	} else {
		currentDockerfile += ("\", \"" + paramsString + "\"]")
	}

	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) EntrypointShell(currentDockerfile string, command string, params ...string) (string, error) {
	if util.IsEmpty(command) {
		return currentDockerfile, errors.New("Add Entrypoint Instrument Error")
	}

	currentDockerfile += "\n" + "ENTRYPOINT " + command

	paramsString := strings.Join(params, " ")
	if paramsString != "" {
		currentDockerfile += (" " + paramsString)
	}

	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) Volume(currentDockerfile string, volumes ...string) (string, error) {
	if len(volumes) == 0 {
		return currentDockerfile, errors.New("Add Volume Instrument Error")
	}

	containsWhitespace := false
	for _, volume := range volumes {
		if strings.Contains(volume, " ") {
			containsWhitespace = true
		}
	}

	currentDockerfile += ("\n" + "VOLUME ")

	if containsWhitespace {
		currentDockerfile += ("[\"" + strings.Join(volumes, "\", \"") + "\"]")
	} else {
		currentDockerfile += strings.Join(volumes, " ")
	}

	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) User(currentDockerfile string, user string) (string, error) {
	if util.IsEmpty(user) {
		return currentDockerfile, errors.New("Add User Instrument Error")
	}

	currentDockerfile += ("\n" + "USER " + user)
	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) Workdir(currentDockerfile string, workdirs ...string) (string, error) {
	if len(workdirs) == 0 {
		return "", errors.New("Add Workdir Instrument Error")
	}

	for _, workdir := range workdirs {
		currentDockerfile += ("\n" + "WORKDIR " + workdir)
	}

	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) Arg(currentDockerfile string, args ...string) (string, error) {
	if len(args) == 0 {
		return currentDockerfile, errors.New("Add Arg Instrument Error")
	} else if len(args) == 1 {
		currentDockerfile += ("\n" + "ARG " + args[0])
	} else {
		currentDockerfile += ("\n" + "ARG " + args[0] + "=" + args[1])
	}

	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) Onbuild(currentDockerfile string, instruction string) (string, error) {
	if util.IsEmpty(instruction) {
		return currentDockerfile, errors.New("Add Onbuild Instrument Error")
	}

	currentDockerfile += ("\n" + "ONBUILD " + instruction)

	return currentDockerfile, nil
}

func (dfo *DockerFileOperator) Stopsignal(currentDockerfile string, signal string) (string, error) {
	if util.IsEmpty(signal) {
		return currentDockerfile, errors.New("Add Stopsignal Instrument Error")
	}

	currentDockerfile += ("\n" + "STOPSIGNAL " + signal)
	return currentDockerfile, nil
}

// TODO ADD MORE
