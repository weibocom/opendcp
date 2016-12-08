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

func (do *DockerOperator) BuildImage(dockerFilePath string, tag string) *stackError.Error {
	_, error := util.ExecuteFullCommand("docker build -t " + tag + " " + dockerFilePath)

	if error != nil {
		util.PrintErrorStack(error)
		return util.ErrorWrapper(error)
	}

	return nil
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
func (do *DockerOperator) PushImage(dockerfilePath, tag string) *stackError.Error {
	_, error := util.ExecuteFullCommand("docker push " + tag)
	if error != nil {
		util.PrintErrorStack(error)
		return util.ErrorWrapper(error)
	}

	return nil
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
