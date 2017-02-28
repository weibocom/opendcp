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
	log "github.com/Sirupsen/logrus"
	stackError "github.com/go-errors/errors"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
	"weibo.com/opendcp/imagebuild/code/util"
	"weibo.com/opendcp/imagebuild/code/env"
)

/**
docker build push
 */
type DockerOperator struct {
}

var instance *DockerOperator
var once sync.Once

// single instance
func GetDockerOperatorInstance() *DockerOperator {
	once.Do(func() {
		instance = &DockerOperator{}
	})

	return instance
}

func (do *DockerOperator) BuildImage(dockerFilePath string, tag string) (string, *stackError.Error) {
	logBack, err := util.ExecuteFullCommand("docker build -t " + tag + " " + dockerFilePath)

	if err != nil {
		util.PrintErrorStack(err)
		return logBack, util.ErrorWrapper(err)
	}

	return logBack, nil
}

func (do *DockerOperator) timeIsOk(imageTime string, startTime time.Time) bool {
	if strings.Contains(imageTime, "hour") || strings.Contains(imageTime, "day") || strings.Contains(imageTime, "week") || strings.Contains(imageTime, "month") || strings.Contains(imageTime, "year") {
		return false
	}

	now := time.Now()
	tmp := strings.Split(imageTime, " ")
	var timeToCompare time.Time

	if tmp[0] == "less" {
		return true
	}

	if tmp[1] == "a" {
		timeToCompare = startTime.Add(time.Duration(1 * time.Minute))
	} else if tmp[1] == "minutes" {
		i, _ := strconv.Atoi(tmp[0])
		timeToCompare = startTime.Add(time.Duration(i * int(time.Minute)))
	} else {
		i, _ := strconv.Atoi(tmp[0])
		timeToCompare = startTime.Add(time.Duration(i * int(time.Second)))
	}

	if timeToCompare.Before(now) {
		return true
	}

	return false
}

func (do *DockerOperator) CheckImageExist(tag string, startTime time.Time) bool {
	out, error := util.ExecuteFullCommand("docker images")
	if error != nil {
		log.Errorf("%s", error)
		return false
	}

	images := strings.Split(out, "\n")
	for _, image := range images {
		re := regexp.MustCompile("[ ]+")
		imageNew := re.ReplaceAllString(image, " ")
		attributes := strings.Split(imageNew, " ")

		imageTag := attributes[0]
		createTime := attributes[4]

		if imageTag == tag && do.timeIsOk(createTime, startTime) {
			return true
		}
	}

	return false
}

// login harbor
func (do *DockerOperator) LoginHarbor() *stackError.Error {
	_, error := util.ExecuteFullCommand("docker login -u " + env.HARBOR_USER + " -p " + env.HARBOR_PASSWORD + " " + env.HARBOR_ADDRESS)
	if error != nil {
		util.PrintErrorStack(error)
		return util.ErrorWrapper(error)
	}

	return nil
}

// push image
func (do *DockerOperator) PushImage(dockerfilePath, tag string) (string, *stackError.Error) {
	logStr, err := util.ExecuteFullCommand("docker push " + tag)
	if err != nil {
		util.PrintErrorStack(err)
		return logStr, util.ErrorWrapper(err)
	}

	return logStr, nil
}

// delete image
func (do *DockerOperator) DeleteImage(tag string) bool {
	_, error := util.ExecuteFullCommand("docker rmi " + tag)
	if error != nil {
		log.Errorf("%s", error)
		return false
	}

	return true
}

func (do *DockerOperator) GenerateRandomTag() (*stackError.Error, string) {
	re, error := regexp.Compile("[ \\-\\:]")
	if error != nil {
		return util.ErrorWrapper(error), ""
	}

	time := re.ReplaceAllString(time.Now().Format("2006-01-02 15:04:05"), "_")
	return nil, time

}
