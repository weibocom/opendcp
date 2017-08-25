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

package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"

	"strings"
	"weibo.com/opendcp/orion/models"
	"weibo.com/opendcp/orion/service"
	"weibo.com/opendcp/orion/utils"
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
	runURL     = "http://%s" + beego.AppConfig.String("remote_command_url")
	checkURL   = "http://%s" + beego.AppConfig.String("remote_check_url")
	logURL     = "http://%s" + beego.AppConfig.String("remote_log_url")
	forkNum    = 5

	logService = service.Logs
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

	fid := nodes[0].Flow.Id

	step := action.Name

	logService.Debug(fid, "Handle remote step:", step)

	rstep := models.RemoteStep{Name: step}
	err := service.Remote.GetBy(&rstep, "Name")

	if err != nil {
		logService.Error(fid, "remote step not found step:", step)

		return Err("remote step not found: " + step)
	}

	// get actions
	actions := rstep.Actions
	logService.Debug(fid, "remote step has actions", actions)

	var actNames []string
	json.Unmarshal([]byte(actions), &actNames)
	logService.Debug(fid, fmt.Sprintf("remote step has actions name:%s len:%d", actNames, len(actNames)))

	var actionList []models.RemoteAction
	service.Remote.GetByStringValues(&models.RemoteAction{}, &actionList,
		"name", actNames)

	// generate playbook
	tpls := make([]interface{}, len(actionList))
	for _, action := range actionList {
		actID := action.Id
		logService.Debug(fid, fmt.Sprintf("getting act impl for %d", actID))

		act := models.RemoteActionImpl{ActionId: actID}
		err = service.Remote.GetBy(&act, "ActionId")
		if err != nil {
			logService.Error(fid, fmt.Sprintf("Cannot find act impl %d", actID))
			return Err("act impl not found: " + strconv.Itoa(actID))
		}

		tpl := make(map[string]interface{})
		err = json.Unmarshal([]byte(act.Template), &tpl)
		if err != nil {
			logService.Error(fid, fmt.Sprintf("Bad act impl template, actid:%d Template:%s", actID, act.Template), err)

			return Err("bad impl template: " + strconv.Itoa(actID))
		}

		idx := h.indexOf(actNames, action.Name)
		if idx == -1 {
			logService.Error(fid, fmt.Sprintf("Action [%s] not in action list:%v", action.Name, actNames))

			return Err("Action: " + action.Name)
		}

		tpls[idx] = tpl
		logService.Debug(fid, fmt.Sprintf("template of %d is %s", actID, act.Template))
	}

	logService.Debug(fid, fmt.Sprintf("remote step template:%v", tpls))

	// call ansible executor

	ret := make([]*NodeResult, len(nodes))
	ipsChan := make(chan map[string]*NodeResult, len(nodes))
	ipRet := make(map[string]*NodeResult, len(nodes))

	for _, node := range nodes {
		go h.callAndCheck(fid, corrId, node.Ip, rstep.Name, &tpls, &stepParams, ipsChan)
	}

	for i := 0; i < len(nodes); i++ {
		select {
		case nodeRespMapList := <-ipsChan:
			for ipString, nodeResp := range nodeRespMapList {
				logService.Debug(fid, fmt.Sprintf("%s runAndCheck is end !", ipString))
				ipRet[ipString] = nodeResp
			}

		case <-time.After(time.Second*checkTimeout*5 + 1):
			logService.Debug(fid, "runAndCheck timeout !")
		}
	}

	//如果全部成功则为成功 ,如果全部失败则为失败.如果有成功有失败..则返回部分成功.
	haveFail := false
	haveSucc := false
	for i, node := range nodes {
		ip := node.Ip

		if ipRet[ip] == nil {
			ipRs := &NodeResult{
				Code: CODE_ERROR,
				Data: fmt.Sprintf(" %s runAndCheck timeout !", ip),
			}

			ret[i] = ipRs
			haveFail = true
			continue
		}

		if ipRet[ip].Code == CODE_SUCCESS && haveSucc == false {
			haveSucc = true
		} else if ipRet[ip].Code != CODE_SUCCESS && haveFail == false {
			haveFail = true
		}

		ret[i] = ipRet[ip]
	}

	taskRsCode := 0
	if haveFail && haveSucc {
		taskRsCode = CODE_PARTIAL
	} else if haveFail {
		taskRsCode = CODE_ERROR
	} else if haveSucc {
		taskRsCode = CODE_SUCCESS
	}

	return &HandleResult{
		Code:   taskRsCode,
		Msg:    "",
		Result: ret,
	}
}

func (h *RemoteHandler) callAndCheck(fid int, corrId string, ip string, setupName string, tpls *[]interface{}, stepParams *map[string]interface{}, ipsChan chan map[string]*NodeResult) {
	execID, user := ip+"_"+setupName+"_"+fmt.Sprint(time.Now().UnixNano()), "root"

	_, err := h.callExecutor(&[]string{ip}, user, execID, tpls, stepParams, corrId)
	if err != nil {
		logService.Error(fid, fmt.Sprintf("%s fail to execute command %v", ip, err.Error()))

		rs := make(map[string]*NodeResult)
		rs[ip] = &NodeResult{
			Code: CODE_ERROR,
			Data: fmt.Sprintf("%s fail to execute command %v", ip, err.Error()),
		}

		ipsChan <- rs
		return
	}

	// check until got result
	for i := 0; i < checkTimeout; i++ {
		time.Sleep(5 * time.Second)
		logService.Debug(fid, fmt.Sprintf("Checking result for task %s for times %d", execID, i+1))

		resp, err := h.checkResult(execID, corrId)
		if err != nil {
			logService.Error(fid, fmt.Sprintf("Checking result for task %s for times %d FAIL:\n %s", execID, i+1, err.Error()))
			continue
		}

		logService.Debug(fid, fmt.Sprintf("Checking result for task %s for times %d status:%d", execID, i+1, resp.Task.Status))
		switch resp.Task.Status {
		case CODE_INIT, CODE_RUNNING:
			continue
		case CODE_ERROR:
			rs := make(map[string]*NodeResult)
			rs[ip] = &NodeResult{
				Code: CODE_ERROR,
				Data: resp.Task.Err,
			}

			ipsChan <- rs
			return
		default:
			for _, nodeResp := range resp.Nodes {
				rs := make(map[string]*NodeResult)
				rs[nodeResp.IP] = &NodeResult{
					Code: nodeResp.Status,
					Data: nodeResp.Log,
				}

				ipsChan <- rs
				return
			}
		}
	}

	rs := make(map[string]*NodeResult)
	rs[ip] = &NodeResult{
		Code: CODE_ERROR,
		Data: fmt.Sprintf("%s checkResult timeout !", ip),
	}

	ipsChan <- rs
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
	corrId, instanceId := strconv.Itoa(nodeState.Flow.Id), nodeState.VmId

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

	resp := &logResp{}
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
