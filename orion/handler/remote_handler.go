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
package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"

	"weibo.com/opendcp/orion/models"
	"weibo.com/opendcp/orion/service"
	"weibo.com/opendcp/orion/utils"
	"strings"
)

// RemoteHandler handle all remote steps. It will create playbook from remote
// steps and send the command to Ansible or other channels.
type RemoteHandler struct {
}

// struct for get the check response.
type remoteCheckResponse struct {
	Task struct {
		ID     int    `json:"id"`
		Status int    `json:"status"`
		Err    string `json:"err"`
	} `json:"task"`
	Nodes []struct {
		IP     string `json:"ip"`
		Status int    `json:"status"`
		Log    string
	} `json:"nodes"`
}

var (
	remoteAddr = beego.AppConfig.String("remote_mgr_addr")
	runURL   = "http://%s" + beego.AppConfig.String("remote_command_url")
	checkURL = "http://%s" + beego.AppConfig.String("remote_check_url")
	logURL = "http://%s" + beego.AppConfig.String("remote_log_url")
	forkNum  = 5
)

const (
	checkTimeout = 120
	maxCount     = 1000
)

// GetType implements method of interface Handler, returns "remote".
func (h *RemoteHandler) GetType() string {
	return "remote"
}

// Handle implements method of interface Handler, handles remote step.
func (h *RemoteHandler) Handle(action *models.ActionImpl,
	stepParams map[string]interface{}, nodes []*models.NodeState, corrId string) *HandleResult {

	step := action.Name
	beego.Debug("Handle remote step", step)
	rstep := models.RemoteStep{Name: step}
	err := service.Remote.GetBy(&rstep, "Name")

	if err != nil {
		beego.Error("remote step not found", step)
		return Err("remote step not found: " + step)
	}

	// get actions
	actions := rstep.Actions
	beego.Debug("remote step has actions", actions)
	var actNames []string
	json.Unmarshal([]byte(actions), &actNames)
	beego.Debug("remote step has actions", actNames, len(actNames))

	var actionList []models.RemoteAction
	service.Remote.GetByStringValues(&models.RemoteAction{}, &actionList,
		"name", actNames)

	// generate playbook
	tpls := make([]interface{}, len(actionList))
	for _, action := range actionList {
		actID := action.Id
		beego.Debug("getting act impl for ", actID)
		act := models.RemoteActionImpl{ActionId: actID}
		err = service.Remote.GetBy(&act, "ActionId")
		if err != nil {
			beego.Error("Cannot find act impl", actID)
			return Err("act impl not found: " + strconv.Itoa(actID))
		}

		tpl := make(map[string]interface{})
		err = json.Unmarshal([]byte(act.Template), &tpl)
		if err != nil {
			beego.Error("Bad act impl template: ", actID, act.Template, err)
			return Err("bad impl template: " + strconv.Itoa(actID))
		}

		idx := h.indexOf(actNames, action.Name)
		if idx == -1 {
			beego.Error("Action [", action.Name, "] not in action list:", actNames)
			return Err("Action: " + action.Name)
		}

		tpls[idx] = tpl
		beego.Debug("template of", actID, "is", act.Template)
	}

	beego.Debug("remote step template:", tpls)

	// call ansible executor
	ips := make([]string, len(nodes))
	ipIDMap := make(map[string]int)
	ipIdxMap := make(map[string]int)
	for i, node := range nodes {
		ips[i] = node.Ip
		ipIDMap[node.Ip] = node.Id
		ipIdxMap[node.Ip] = i
	}

	execID, user := rstep.Name+"_"+fmt.Sprint(time.Now().UnixNano()), "root"
	_, err = h.callExecutor(&ips, user, execID, &tpls, &stepParams, corrId)
	if err != nil {
		return Err("fail to execute command: " + err.Error())
	}

	// check until got result
	for i := 0; i < checkTimeout; i++ {
		time.Sleep(5 * time.Second)

		beego.Debug("Checking result for task", execID, "for times", i+1)
		resp, err := h.checkResult(execID, corrId)

		if err == nil {
			beego.Debug("Checking result for task", execID, "for times", i+1,
				"status:", resp.Task.Status)

			switch resp.Task.Status {
			case CODE_INIT, CODE_RUNNING:
				continue
			case CODE_ERROR:
				return Err(resp.Task.Err)
			default:
				ret := make([]*NodeResult, len(nodes))
				for _, nodeResp := range resp.Nodes {
					ip := nodeResp.IP
					ret[ipIdxMap[ip]] = &NodeResult{
						Code: nodeResp.Status,
						Data: nodeResp.Log,
					}
				}

				return &HandleResult{
					Code:   CODE_SUCCESS,
					Msg:    "",
					Result: ret,
				}
			}

		} else {
			beego.Error("Checking result for task", execID, "for times", i+1, "FAIL:\n", err.Error())
		}
	}

	return Err("Timeout to check result")
}

// ListAction implements method of interface Handler, and will return all
// remote steps defined bu users.
func (h *RemoteHandler) ListAction() []models.ActionImpl {
	var list []models.RemoteStep
	count, err := service.Remote.ListByPage(0, maxCount, &models.RemoteStep{}, &list)
	if err != nil {
		beego.Error("Fail to get remote steps ", err)
		return []models.ActionImpl{}
	}

	acts := make([]models.ActionImpl, count)
	for i, step := range list {
		params, err := h.getStepParams(&step)
		if err != nil {
			beego.Error("Fail to get params for step", step.Name, err)
			continue
		}

		act := models.ActionImpl{
			Name:   step.Name,
			Desc:   step.Desc,
			Type:   "remote",
			Params: params,
		}

		acts[i] = act
	}

	return acts
}

func (h *RemoteHandler) callExecutor(ips *[]string, user string, execID string,
	content *[]interface{}, params *map[string]interface{}, corrId string) (string, error) {

	header := make(map[string]interface{})
	header["X-CORRELATION-ID"] = corrId
	header["X-SOURCE"] = "orion"

	data := make(map[string]interface{})
	data["nodes"] = ips
	data["user"] = user
	data["name"] = execID
	data["tasks"] = content
	data["params"] = params
	data["fork_num"] = forkNum

	url := fmt.Sprintf(runURL, remoteAddr)
	msg, err := utils.Http.Post(url, &data, &header)
	return msg, err
}

func (h *RemoteHandler) checkResult(id string, corrId string) (*remoteCheckResponse, error) {
	data := make(map[string]interface{})
	data["name"] = id

	header := make(map[string]interface{})
	header["X-CORRELATION-ID"] = corrId
	header["X-SOURCE"] = "orion"

	url := fmt.Sprintf(checkURL, remoteAddr)
	msg, err := utils.Http.Post(url, &data, &header)

	if err != nil {
		return nil, err
	}

	resp := &struct {
		Content remoteCheckResponse `json:"content"`
	}{}
	err = json.Unmarshal([]byte(msg), resp)
	if err != nil {
		return &resp.Content, err
	}

	return &resp.Content, nil
}

func (h *RemoteHandler) getStepParams(step *models.RemoteStep) (map[string]interface{}, error) {
	tmp := step.Actions
	var acts []string
	err := json.Unmarshal([]byte(tmp), &acts)
	if err != nil {
		return nil, err
	}

	params := make(map[string]interface{})
	for _, act := range acts {
		ra := &models.RemoteAction{Name: act}
		err = service.Remote.GetBy(ra, "Name")
		if err != nil {
			msg := "cannot find remote action " + act
			beego.Error("Getting params for", step.Name, ", err: ", msg)
			return nil, errors.New(msg)
		}

		p := make(map[string]interface{})
		err = json.Unmarshal([]byte(ra.Params), &p)
		if err != nil {
			msg := "bad remote action [" + act + "] params:" + ra.Params
			beego.Error(msg)
			return nil, errors.New(msg)
		}

		for k, v := range p {
			params[k] = v
		}
	}

	return params, nil
}

type logResp struct {
	Content struct {
		Log []string
	}
}

func (h *RemoteHandler) GetLog(nodeState *models.NodeState) string {
	corrId , instanceId := nodeState.CorrId, nodeState.VmId

	header := make(map[string]interface{})
	header["X-CORRELATION-ID"] = corrId
	header["X-SOURCE"] = "orion"

	data := make(map[string]interface{})
	data["host"] = nodeState.Ip
	data["source"] = "orion"

	beego.Debug("Get log for", instanceId, nodeState.Ip, "....")

	url := fmt.Sprintf(logURL, remoteAddr)
	raw, err := utils.Http.Post(url, &data, &header)
	if err != nil {
		beego.Error("Error when getting log for", instanceId, "err:", err)
		return "<NO LOG>"
	}

	resp := &logResp {}
	err = json.Unmarshal([]byte(raw), &resp)
	if err != nil {
		beego.Error("Error when parsing log for", instanceId, "err:", err)
		return "<NO LOG>"
	}

	return strings.Join(resp.Content.Log, "\n")
}

func (h *RemoteHandler) indexOf(array []string, v string) int {
	for i, x := range array {
		if x == v {
			return i
		}
	}

	return -1
}
