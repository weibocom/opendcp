// Copyright 2016 Weibo Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use sf file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package future

import (
	"fmt"
	"strings"
	"time"

	"weibo.com/opendcp/jupiter/conf"
	"weibo.com/opendcp/jupiter/dao"
	"weibo.com/opendcp/jupiter/models"
	"weibo.com/opendcp/jupiter/response"
	"weibo.com/opendcp/jupiter/logstore"
	"errors"
)

const (
	STATUS_INIT = iota
	STATUS_RUNNING
	STATUS_SUCCESS
	STATUS_FAILED
	STATUS_STOPPED
)

type AnsibleTaskFuture struct {
	InstanceId string
	Ip    string
	Roles []string
	CorrelationId string
}

func NewAnsibleTaskFuture(instanceId, ip string, roles []string, correlationId string) *AnsibleTaskFuture {
	return &AnsibleTaskFuture{
		InstanceId: instanceId,
		Ip:    ip,
		Roles: roles,
		CorrelationId: correlationId,
	}
}

func (atf *AnsibleTaskFuture) Run() error {
	initEnv := strings.Join(atf.Roles, ",")
	rolesStr := strings.Join(atf.Roles, "\",\"")
	logstore.Info(atf.CorrelationId, atf.InstanceId, "initEnv [", initEnv, "], rolesStr [", rolesStr, "]")
	taskName := atf.Ip + time.Now().Format("2006-01-02 15:04:05")
	body := "{\"nodes\": [\"%s\"],\"tasks\": [\"%s\"],\"tasktype\": \"ansible_role\", \"user\": \"root\", \"name\": \"%s\", \"fork_num\": %d}"
	body = fmt.Sprintf(body, atf.Ip, rolesStr, atf.Ip + taskName, conf.Config.Ansible.ForkNum)
	logstore.Info(atf.CorrelationId, atf.InstanceId, "init req body: ", body)
	url := conf.Config.Ansible.Url + "/api/run"
	msg, err := response.CallApi(body, "POST", url, atf.CorrelationId)
	logstore.Info(atf.CorrelationId, atf.InstanceId, "call octaons api response:", msg)
	if err != nil {
		logstore.Error(atf.CorrelationId, atf.InstanceId, "call octans api error:", err)
		for i:=0; i < 3; i++ {
			msg, err = response.CallApi(body, "POST", url, atf.CorrelationId)
			if err == nil {
				break
			}
			logstore.Error(atf.CorrelationId, atf.InstanceId, "call octans api error:", err)
			return err
		}
	}
	respMap := response.RespToMap(msg)
	if respMap == nil {
		return errors.New("Octans response can't resolve.")
	}
	code := int(respMap["code"].(float64))
	if code != 0 {
		msg = fmt.Sprint(respMap["msg"])
		return errors.New("Octans init response code != 0!")
	}
	content := respMap["content"].(map[string]interface{})
	taskId := int(content["id"].(float64))
	// 添加重试
	err = checkTaskState(taskId, atf.Ip, atf.InstanceId, atf.CorrelationId)
	if err != nil {
		for i := 0; i < 3; i++ {
			err = checkTaskState(taskId, atf.Ip, atf.InstanceId, atf.CorrelationId)
			if err == nil {
				return nil
			}
			logstore.Error(atf.CorrelationId, atf.InstanceId, "octans check api error:", err)
		}
		logstore.Error(atf.CorrelationId, atf.InstanceId, "octans check api error:", err)
		return err
	}
	return nil
}

func (atf *AnsibleTaskFuture) Success() {
	if err := dao.UpdateInstanceStatus(atf.Ip, models.Success); err != nil {
		logstore.Error(atf.CorrelationId, atf.InstanceId, "update db err: ", err)
	}
	logstore.Info(atf.CorrelationId, atf.InstanceId, "init finished: ", atf.Ip)
}

func (atf *AnsibleTaskFuture) Failure(err error) {
	if err := dao.UpdateInstanceStatus(atf.Ip, models.Uninit); err != nil {
		logstore.Error(atf.CorrelationId, atf.InstanceId, "update db err: ", err)
	}
	logstore.Error(atf.CorrelationId, atf.InstanceId, "init err: ", err)
}

func (atf *AnsibleTaskFuture) ShutDown() {
}

func checkTaskState(taskId int, ip string, instanceId, correlationId string) error {
	body := "{\"id\": %d}"
	body = fmt.Sprintf(body, taskId)
	logstore.Info(correlationId, instanceId, "check cotans http api request body")
	for i := 0; i < 60; i++ {
		time.Sleep(3 * time.Second)
		url := conf.Config.Ansible.Url + "/api/check"
		msg, err := response.CallApi(body, "POST", url, correlationId)
		if err != nil {
			logstore.Error(correlationId, instanceId, "ocatan check running result error", "the response is", msg, "error is", err)
			return err
		}
		logstore.Info(correlationId, instanceId, "wait for instance", ip, "to init:", i)
		respMap := response.RespToMap(msg)
		if respMap == nil {
			logstore.Error(correlationId, instanceId, "Can't resolve ocatans check response")
			return errors.New("Can't resolve octans check response")
		}
		code := int(respMap["code"].(float64))
		if code != 0 {
			msg = fmt.Sprint(respMap["msg"])
			logstore.Error(correlationId, instanceId, "Check Ocatans result error: " + msg)
			return errors.New("init response code !=  0!")
		}
		content := respMap["content"].(map[string]interface{})
		taskStateMap := content["task"].(map[string]interface{})
		state := int(taskStateMap["status"].(float64))
		switch state {
		case STATUS_SUCCESS:
			logstore.Info(correlationId, instanceId, content)
			logstore.Info(correlationId, instanceId, "Init success, task Id is", taskId)
			return nil
		case STATUS_FAILED:
			errMsg := taskStateMap["err"].(string)
			logstore.Info(correlationId, instanceId, content)
			logstore.Error(correlationId, instanceId, taskId, "The task id", taskId, "falied:", errMsg)
			return errors.New(errMsg)
		}
	}
	dao.UpdateInstanceStatus(ip, models.InitTimeout)
	return errors.New("Wait octans return success timeout")
}
