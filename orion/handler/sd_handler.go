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
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"weibo.com/opendcp/orion/models"
	"weibo.com/opendcp/orion/utils"
)

const (
	REG   = "register"
	UNREG = "unregister"

	SV_ID = "service_discovery_id"
)

var (
	SD_ADDR   = beego.AppConfig.String("sd_mgr_addr")
	SD_APPKEY = beego.AppConfig.String("sd_appkey")

	REG_URL      = "http://%s" + beego.AppConfig.String("sd_register_url")
	UNREG_URL    = "http://%s" + beego.AppConfig.String("sd_unregister_url")
	SD_CHECK_URL = "http://%s" + beego.AppConfig.String("sd_check_url") // + "?%s=%d&%s=%s"
	SD_LOG_URL   = "http://%s" + beego.AppConfig.String("sd_log_url")
)

type ServiceDiscoveryHandler struct {
}

type sdCmdResp struct {
	Code    int
	Message string `json:"msg"`
	Content struct {
		Type   string
		TaskId string `json:"task_id"`
	} `json:"data"`
}

type sdLogResp struct {
	Code    int
	Message string `json:"msg"`
	Content string `json:"data"`
}

type sdChkResp struct {
	Code    int
	Message string `json:"msg"`
	Content struct {
		State  int
		Detail []struct {
			Ip    string
			State int
		}
	} `json:"data"`
}

func (v *ServiceDiscoveryHandler) ListAction() []models.ActionImpl {
	return []models.ActionImpl{
		{
			Name: REG,
			Desc: "Register to service",
			Type: "sd",
			Params: map[string]interface{}{
				SV_ID: "Integer",
			},
		},
		{
			Name: UNREG,
			Desc: "Unregister from service",
			Type: "sd",
			Params: map[string]interface{}{
				SV_ID: "Integer",
			},
		},
	}
}

func (h *ServiceDiscoveryHandler) GetType() string {
	return "sd"
}

func (h *ServiceDiscoveryHandler) Handle(action *models.ActionImpl,
	actionParams map[string]interface{}, nodes []*models.NodeState, corrId string) *HandleResult {

	fid := nodes[0].Flow.Id
	batchId := nodes[0].Batch.Id

	logService.Debug(fid, batchId, corrId, fmt.Sprintf("sd handler recieve new action: [%s]", action.Name))

	switch action.Name {
	case REG:
		return h.register(actionParams, nodes, corrId)
	case UNREG:
		return h.unregister(actionParams, nodes, corrId)
	default:
		logService.Error(fid, batchId, corrId, fmt.Sprintf("Unknown SD action: [%s]", action.Name))

		return Err("Unknown action: " + action.Name)
	}
}

func (h *ServiceDiscoveryHandler) register(params map[string]interface{},
	nodes []*models.NodeState, corrId string) *HandleResult {

	return h.do(REG_URL, params, nodes, corrId)
}

func (h *ServiceDiscoveryHandler) unregister(params map[string]interface{},
	nodes []*models.NodeState, corrId string) *HandleResult {

	return h.do(UNREG_URL, params, nodes, corrId)
}

func (h *ServiceDiscoveryHandler) do(action string, params map[string]interface{},
	nodes []*models.NodeState, corrId string) *HandleResult {

	fid := nodes[0].Flow.Id
	batchId := nodes[0].Batch.Id

	logService.Debug(fid, batchId, corrId, fmt.Sprintf("sd , service_discovery_id =%v,corrId =%s", params[SV_ID], corrId))

	svVal := params[SV_ID]
	sv, err := utils.ToInt(svVal)

	if err != nil {
		logService.Error(fid, batchId, corrId, fmt.Sprintf("Bad service_discovery_id :[%v]", svVal))

		return Err("Bad servicd_id")
	}

	// call api
	logService.Debug(fid, batchId, corrId, fmt.Sprintf("SD:%d , nodes = %v", sv, nodes))

	ips := make([]string, len(nodes))
	for i, node := range nodes {
		ips[i] = node.Ip
	}

	data := make(map[string]interface{})
	data["type_id"] = sv
	data["ips"] = strings.Join(ips, ",")
	data["user"] = "root"

	header := make(map[string]interface{})
	header["X-CORRELATION-ID"] = corrId
	header["APPKEY"] = SD_APPKEY

	resp := &sdCmdResp{}
	url := fmt.Sprintf(action, SD_ADDR)
	hr := h.callAPI("POST", url, &data, &header, resp)
	if hr != nil {
		return hr
	}

	if resp.Code != 0 {
		return Err(resp.Message)
	}

	// return directly if sync
	if resp.Content.Type == "sync" {
		return h.success(nodes)
	}

	// check result if async
	taskId := resp.Content.TaskId
	logService.Debug(fid, batchId, corrId, fmt.Sprintf("task id = %s", taskId))

	// start checking result
	for i := 0; i < timeout/5; i++ {
		time.Sleep(5 * time.Second)
		logService.Info(fid, batchId, corrId, fmt.Sprintf("check result for times %d", i+1))

		//data := make(map[string]interface{})
		//data["task_id"] = taskId
		//data["appkey"] = SD_APPKEY

		//header := map[string]interface{} {
		//	"APPKEY": SD_APPKEY,
		//}

		url := fmt.Sprintf(SD_CHECK_URL, SD_ADDR) //, "task_id", taskId, "appkey", SD_APPKEY)
		msg, err := utils.Http.Get(url, &header)
		if err != nil {
			logService.Warn(fid, batchId, corrId, fmt.Sprintf("check result err: \n%v", err))

			continue
		}

		resp := &sdChkResp{}
		err = json.Unmarshal([]byte(msg), resp)
		if err != nil {
			logService.Error(fid, batchId, corrId, fmt.Sprintf("bad response: %s", msg))

			continue
		}

		if resp.Code != 0 {
			logService.Error(fid, batchId, corrId, fmt.Sprintf("check result return fail"))

			continue
		}

		if resp.Content.State == CODE_ERROR { // fail
			return Err("FAIL")
		}

		if resp.Content.State == CODE_SUCCESS { // success
			return h.success(nodes)
		}
	}

	return Err("Timeout checking result")
}

func (v *ServiceDiscoveryHandler) callAPI(method string, url string,
	data *map[string]interface{}, header *map[string]interface{}, obj interface{}) *HandleResult {

	msg, err := utils.Http.Do(method, url, data, header)
	if err != nil {
		beego.Error("Fail to ", method, url, ": ", err)
		return Err("Fail: " + err.Error())
	}

	err = json.Unmarshal([]byte(msg), obj)
	if err != nil {
		beego.Error("Fail to unmarshal", msg, "err:", err)
		beego.Error("Bad resp:", msg)
		return Err("Bad resp: " + msg)
	}

	return nil
}

func (h *ServiceDiscoveryHandler) GetLog(nodeState *models.NodeState) string {
	corrId, instanceId := nodeState.CorrId, nodeState.VmId

	header := make(map[string]interface{})
	header["X-CORRELATION-ID"] = corrId
	header["APPKEY"] = SD_APPKEY

	resp := &sdLogResp{}
	url := fmt.Sprintf(SD_LOG_URL, SD_ADDR, corrId)
	err := h.callAPI("GET", url, nil, &header, resp)
	if err != nil {
		beego.Error("Get log for", instanceId, "fails:", err)
		return "<NO LOG>"
	}

	return resp.Content
}

func (h *ServiceDiscoveryHandler) success(nodes []*models.NodeState) *HandleResult {
	nRet := make([]*NodeResult, len(nodes))
	for i := 0; i < len(nodes); i++ {
		nRet[i] = &NodeResult{
			Code: CODE_SUCCESS,
			Data: "OK",
		}
	}

	return &HandleResult{
		Code:   CODE_SUCCESS,
		Msg:    "OK",
		Result: nRet,
	}
}
