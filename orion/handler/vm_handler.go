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
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"weibo.com/opendcp/orion/models"
	"weibo.com/opendcp/orion/service"
	"weibo.com/opendcp/orion/utils"
	"strconv"
)

const (
	createVM = "create_vm"
	returnVM = "return_vm"

	vmTypeId = "vm_type_id"
)

const (
	vmPending = iota
	vmSuccess
	vmUninit
	vmIniting
	vmInitTimeout
	vmDeleted
	vmUninstalling
	vmUniTimeout
	vmDeleting
	vmError
)

var (
	jupiterAddr = beego.AppConfig.String("vm_mgr_addr")
	timeout     = 90 * 5

	apiCreate = "http://%s" + beego.AppConfig.String("vm_create_url")
	apiReturn = "http://%s" + beego.AppConfig.String("vm_return_url")
	apiCheck  = "http://%s" + beego.AppConfig.String("vm_check_url")
	apiLog    = "http://%s" + beego.AppConfig.String("vm_log_url")
)

// VMHandler handles step involving creating, returning VM machines.
type VMHandler struct {
}

// ListAction implements method of Handler.
func (v *VMHandler) ListAction() []models.ActionImpl {
	return []models.ActionImpl{
		{
			Name: createVM,
			Desc: "Create VMs",
			Type: "vm",
			Params: map[string]interface{}{
				vmTypeId: "Integer",
			},
		},
		{
			Name: returnVM,
			Desc: "Return VMs",
			Type: "vm",
			Params: map[string]interface{}{
				vmTypeId: "Integer",
			},
		},
	}
}

// GetType implements method of Handler.
func (v *VMHandler) GetType() string {
	return "vm"
}

// Handle implements method of Handler.
func (v *VMHandler) Handle(action *models.ActionImpl, actionParams map[string]interface{},
	nodes []*models.NodeState, corrId string) *HandleResult {
	beego.Debug("vm handler recieve new action: [", action.Name, "]")

	switch action.Name {
	case createVM:
		return v.createVMs(actionParams, nodes, corrId)
	case returnVM:
		return v.returnVMs(actionParams, nodes, corrId)
	default:
		beego.Error("Unknown VM action: " + action.Name)
		return Err("Unknown VM action: " + action.Name)
	}
}

// Create vm machines from jupiter.
func (v *VMHandler) createVMs(params map[string]interface{},
	nodes []*models.NodeState, corrId string) *HandleResult {

	num := len(nodes)
	beego.Debug("creating vm, vm_type_id =", params[vmTypeId], reflect.TypeOf(params[vmTypeId]))
	cluStr := params[vmTypeId]
	cluster, err := utils.ToInt(cluStr)
	if err != nil {
		beego.Error("Bad cluster:[", cluStr, "]")
		return Err("Bad cluster")
	}

	// call create vm api
	beego.Info("Creating VM in cluster", cluster, ", num=", num)
	url := fmt.Sprintf(apiCreate, jupiterAddr, cluster, num)
	header := map[string]interface{} {
		"X-CORRELATION-ID": corrId,
	}
	resp, hr := v.callAPI("POST", url, nil, &header)
	if hr != nil {
		// remove all node since it fails here
		for _, nodeState := range nodes {
			beego.Info("Deleting node [", nodeState.Node.Id, "]")
			service.Cluster.DeleteBase(nodeState.Node)

			nodeState.Log = "[jupiter]: " + hr.Msg + "\n"
			service.Cluster.UpdateBase(nodeState)
		}
		return hr
	}

	content := resp["content"]
	fmt.Println(reflect.TypeOf(content))

	tmpList, ok := content.([]interface{})
	if !ok {
		beego.Error("Bad id list", fmt.Sprint(content))
		return Err("Bad id list: " + fmt.Sprint(content))
	}

	list := make([]string, len(tmpList))
	for i, id := range tmpList {
		list[i] = id.(string)
	}
	vmIds := list

	if len(vmIds) != len(nodes) {
		beego.Warn("Number of vm ids (", len(vmIds),
			") doesn't equal that of nodes (", len(nodes), ")")
	}

	// update nodes
	nodeMap := make(map[string]*models.NodeState)
	idxMap := make(map[string]int)
	for i, vmID := range vmIds {
		nodes[i].Node.VmId = vmID
		nodes[i].VmId = vmID
		nodeMap[vmID] = nodes[i]
		idxMap[vmID] = i
	}

	// for missing vm ids, mark then as failed
	for i := 0; i < len(nodes) - len(vmIds) ; i ++ {
		node := nodes[i + len(vmIds)]
		node.Status = CODE_ERROR
		node.UpdatedTime = time.Now()

		service.Cluster.UpdateBase(node)
		service.Cluster.DeleteBase(node.Node)
	}

	// start checking result
	beego.Info("VM creating command sent for cluster:", cluster, ", vm ids = ", vmIds)
	var failed, done []string
	for i := 0; i < timeout/5; i++ {
		time.Sleep(5 * time.Second)
		beego.Info("check result for times ", i+1)

		url := fmt.Sprintf(apiCheck, jupiterAddr, strings.Join(list, ","))
		msg, err := utils.Http.Get(url, nil)
		if err != nil {
			beego.Warn("check result err: \n", err)
			continue
		}

		resp, err := utils.Json.ToMap(msg)
		if err != nil {
			beego.Error("bad response: ", msg, ", err:", err)
			continue
		}

		statuses, ok := resp["content"].([]interface{})
		if !ok {
			beego.Error(`bad response "content": ` + msg)
			continue
		}

		var running []string
		for _, v := range statuses {
			tmp := v.(map[string]interface{})
			id := tmp["instance_id"].(string)
			state := int(tmp["status"].(float64))
			toDel := false
			switch state {
			case vmSuccess:
				beego.Debug("Node[", id, "] OK")
				done = append(done, id)
			case vmInitTimeout:
				beego.Debug("Node[", id, "] init timeout")
				failed = append(failed, id)
				toDel = true
			case vmError, vmUninit:
				beego.Debug("Node[", id, "] error")
				failed = append(failed, id)
				toDel = true
			default:
				beego.Debug("Node[", id, "] in progress, status=", state)
				running = append(running, id)
			}

			if tmp["ip_address"] != nil {
				n := nodeMap[id]
				n.Ip = tmp["ip_address"].(string)
				n.VmId = tmp["instance_id"].(string)
				n.Status = models.STATUS_RUNNING
				n.Node.Ip = n.Ip
				n.Node.VmId = n.VmId

				// save node and this will add node to pool
				service.Cluster.UpdateBase(n.Node)
				service.Cluster.UpdateBase(n)
			}

			// if failed, remove the node from pool
			if toDel {
				beego.Info("Deleting node [", id, "] since it failed to create")
				service.Cluster.DeleteBase(nodeMap[id].Node)
			}
		}

		list = running

		if len(list) == 0 {
			break
		}

	}

	// this nodes are timeout, mark them as failed
	if len(list) != 0 {
		for _, id := range list {
			beego.Debug("Node[", id, "] timeout")
			failed = append(failed, id)
			n := nodeMap[id]
			n.Status = models.STATUS_FAILED
			service.Cluster.UpdateBase(n)

			beego.Info("Deleting node [", id, "] since it failed to create")
			service.Cluster.DeleteBase(n.Node)
		}
	}

	beego.Info("All finished")
	ret := make([]*NodeResult, len(nodes))
	for _, vid := range done {
		idx := idxMap[vid]
		nr := &NodeResult{
			Code: CODE_SUCCESS,
			Data: "OK",
		}
		ret[idx] = nr
	}

	for _, vid := range failed {
		idx := idxMap[vid]
		nr := &NodeResult{
			Code: CODE_ERROR,
			Data: "FAILED",
		}
		ret[idx] = nr
	}

	// handle missing vms
	for i := 0; i < len(nodes) - len(vmIds); i ++ {
		nr := &NodeResult{
			Code: CODE_ERROR,
			Data: "FAILED",
		}
		ret[i + len(vmIds)] = nr
	}

	return &HandleResult{
		Code:   CODE_SUCCESS,
		Msg:    "",
		Result: ret,
	}
}

// Return vm machines to jupiter.
func (v *VMHandler) returnVMs(params map[string]interface{},
	nodes []*models.NodeState, corrId string) *HandleResult {

	ids := make([]string, 0)
	cannotDelete := make(map[int]bool)
	for _, node := range nodes {
		vmId := node.Node.VmId
		if vmId != "" {
			ids = append(ids, vmId)
			cannotDelete[node.Id] = false
		} else {
			// for vmId == "", we cannot delete them
			cannotDelete[node.Id] = true
		}
	}

	url := fmt.Sprintf(apiReturn, jupiterAddr, strings.Join(ids, ","))
	header := map[string]interface{} {
		"X-CORRELATION-ID": corrId,
		"APPKEY": SD_APPKEY,
	}
	_, hr := v.callAPI("DELETE", url, nil, &header)
	if hr != nil {
		return hr
	}

	// delete nodes from pool
	for _, node := range nodes {
		if !cannotDelete[node.Id] {
			service.Cluster.DeleteBase(&models.Node{Id: node.Node.Id})
		}
	}

	success := false
	nRet := make([]*NodeResult, len(nodes))
	for i := 0; i < len(nodes); i++ {
		if cannotDelete[nodes[i].Id] {
			nRet[i] = &NodeResult{
				Code: CODE_ERROR,
				Data: "no vm id for this node:" + strconv.Itoa(nodes[i].Id),
			}
		} else {
			nRet[i] = &NodeResult{
				Code: CODE_SUCCESS,
				Data: "ok",
			}
			success = true
		}
	}

	code, msg := CODE_ERROR, "failed"
	if success {
		code, msg = CODE_SUCCESS, "ok"
	}

	return &HandleResult{
		Code:   code,
		Msg:    msg,
		Result: nRet,
	}
}

// call jupiter api to create/return vms
func (v *VMHandler) callAPI(method string, url string,
	data *map[string]interface{}, header *map[string]interface{}) (map[string]interface{}, *HandleResult) {

	msg, err := utils.Http.Do(method, url, data, header)
	if err != nil {
		beego.Error("Fail to ", method, url, ": ", err)
		return nil, Err("Fail: " + err.Error())
	}

	resp, err := utils.Json.ToMap(msg)
	if err != nil {
		beego.Error("Bad resp:", msg, ", err:", err)
		return nil, Err("Bad resp: " + msg)
	}
	code := int(resp["code"].(float64))

	if code != 0 {
		msg = fmt.Sprint(resp["msg"])
		beego.Error("Fail: " + msg)
		return nil, Err(msg)
	}

	return resp, nil
}

func (v *VMHandler) GetLog(nodeState *models.NodeState) string {
	corrId , instanceId := nodeState.CorrId, nodeState.VmId
	header := make(map[string]interface{})
	header["X-CORRELATION-ID"] = corrId
	header["X-SOURCE"] = "orion"

	beego.Debug("Get log for", instanceId, "...")
	url := fmt.Sprintf(apiLog, jupiterAddr, corrId, instanceId)
	msg, err := v.callAPI("GET", url, nil, &header)
	if err != nil {
		beego.Error("Error get log for", instanceId, ", err:", err)
		return "<NO LOG>"
	}

	log, ok := msg["content"].(string)
	if !ok {
		beego.Debug("Bad content: ", msg["content"])
	}

	return log
}
