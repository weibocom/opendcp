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

package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"

	"weibo.com/opendcp/orion/executor"
	"weibo.com/opendcp/orion/models"
	"weibo.com/opendcp/orion/service"
	"weibo.com/opendcp/orion/utils"
)

const (
	EXPAND = "expand"
	SHRINK = "shrink"
	DEPLOY = "deploy"

	START_SERVICE = "start_service"

	CREATE_VM = "create_vm"
	RETURN_VM = "return_vm"

	REGISTER   = "register"
	UNREGISTER = "unregister"

	KEY_SD_ID   = "service_discovery_id"
	KEY_VM_TYPE = "vm_type_id"
	KEY_TAG     = "tag"
)

// Expand will expand a service pool by add vms and start service on them.
func Expand(poolId int, num int, opUser string) error {
	pool, flow, steps, error := getModels(poolId, EXPAND)
	if error != nil {
		return error
	}

	if len(steps) < 1 || steps[0].Name != "create_vm" {
		return errors.New("first step of expand template is not create_vm: " + flow.Steps)
	}

	if num < 0 || num > 200 {
		return errors.New("Bad num: " + strconv.Itoa(num))
	}

	// create nodes & states
	beego.Debug("creating nodes...")
	nodes := make([]*models.Node, num)
	for i := 0; i < num; i++ {
		n := &models.Node{
			Ip:       "-",
			Pool:     pool,
			Status:   models.STATUS_INIT,
			NodeType: models.Manual,
		}
		nodes[i] = n

		// save the node to DB
		service.Cluster.InsertBase(n)
	}

	// Use vm_type & service discovery info from Pool
	override := map[string]interface{}{
		CREATE_VM: map[string]interface{}{KEY_VM_TYPE: pool.VmType},
		REGISTER:  map[string]interface{}{KEY_SD_ID: pool.SdId},
		//name:      map[string]interface{}{KEY_TAG: pool.Service.DockerImage},
	}

	name := searchStartServiceStep(steps)
	if name == "" {
		beego.Warn("No step found starting with 'start_service' in flow: ", flow.Id, flow.Name)
	} else {
		// override tag with the new tag given
		override[name] = map[string]interface{}{KEY_TAG: pool.Service.DockerImage}
	}

	context := make(map[string]interface{})
	context["overrideParams"] = override
	context["opUser"] = opUser

	beego.Info("Expand pool[", pool.Name, "] vm_type =", pool.VmType,
		",sd_id =", pool.SdId)

	beego.Debug("exec flow ...")
	err := executor.Executor.Run(flow.Id, "expand_"+pool.Name,
		&executor.ExecOption{MaxNum: num}, nodes, context)
	beego.Debug("exec flow ... [DONE]")

	return err
}

// Shrink will shrink a service pool by stopping service on vms and return them.
func Shrink(poolId int, nodeIps []string, opUser string) error {

	pool, flow, steps, error := getModels(poolId, SHRINK)
	if error != nil {
		return error
	}

	if len(steps) < 1 || steps[len(steps)-1].Name != "return_vm" {
		return errors.New("last step of shrink template is not return_vm: " + flow.Steps)
	}

	nodes := make([]*models.Node, 0)
	for _, ip := range nodeIps {
		n := &models.Node{Ip: ip}
		err := service.Cluster.GetBy(n, "Ip")
		if err != nil {
			beego.Error("Node with IP ", ip, "not foud, ignore")
			continue
		}
		nodes = append(nodes, n)
	}

	// Use vm_type & service discovery info from Pool
	override := map[string]interface{}{
		RETURN_VM:  map[string]interface{}{KEY_VM_TYPE: pool.VmType},
		UNREGISTER: map[string]interface{}{KEY_SD_ID: pool.SdId},
	}

	context := make(map[string]interface{})
	context["overrideParams"] = override
	context["opUser"] = opUser

	beego.Debug("exec shrink flow...")
	err := executor.Executor.Run(flow.Id, "shrink_"+pool.Name,
		&executor.ExecOption{MaxNum: len(nodes)}, nodes, context)
	beego.Debug("exec flow ... [DONE]")

	return err
}

func Deploy(poolId int, tag string, maxNum int, opUser string) error {

	pool, flow, steps, error := getModels(poolId, DEPLOY)
	if error != nil {
		return error
	}

	nodes := make([]*models.Node, 0)
	count, err := service.Cluster.ListByPageWithFilter(0, 10000,
		&models.Node{}, &nodes, "Pool", poolId)
	if err != nil {
		return err
	}

	override := make(map[string]interface{})
	name := searchStartServiceStep(steps)
	if name == "" {
		beego.Warn("No step found starting with 'start_service' in flow: ", flow.Id, flow.Name)
	} else {
		// override tag with the new tag given
		override[name] = map[string]interface{}{KEY_TAG: tag}
	}

	context := make(map[string]interface{})
	context["overrideParams"] = override
	context["opUser"] = opUser

	beego.Debug("exec flow on Pool[", pool.Name, "] node_cound=", count, "...")
	err = executor.Executor.Run(flow.Id, "deploy_"+pool.Name,
		&executor.ExecOption{MaxNum: maxNum}, nodes, context)
	beego.Debug("exec flow ... [DONE]")

	return err
}

func getModels(poolId int, tplType string) (*models.Pool, *models.FlowImpl, []*models.StepOption, error) {
	pool := &models.Pool{Id: poolId}
	err := service.Cluster.GetBase(pool)
	if err != nil {
		return nil, nil, nil, errors.New("Pool not found : " + strconv.Itoa(poolId))
	}

	serv := &models.Service{Id: pool.Service.Id}
	err = service.Cluster.GetBase(serv)
	if err != nil {
		return nil, nil, nil, errors.New("Service not found : " + strconv.Itoa(pool.Service.Id))
	}

	pool.Service = serv

	tasks, err := utils.Json.ToMap(pool.Tasks)
	if err != nil {
		return nil, nil, nil, errors.New("No task def  found")
	}

	_tid := tasks[tplType]
	if _tid == nil {
		return nil, nil, nil, errors.New("No task id found for:" + tplType)
	}

	tid, err := utils.ToInt(_tid)
	if err != nil {
		return nil, nil, nil, errors.New("Bad " + tplType + " task id :" +
			fmt.Sprint(_tid) + ", " + err.Error())
	}

	beego.Info("Get", tplType, "tpl id", tid)

	flow := &models.FlowImpl{Id: tid}
	err = service.Flow.GetBase(flow)
	if err != nil {
		return nil, nil, nil, errors.New("template not found: " + strconv.Itoa(tid))
	}

	var steps []*models.StepOption
	err = json.Unmarshal([]byte(flow.Steps), &steps)
	if err != nil {
		beego.Error("Fail to load template steps:", flow.Steps, ", err:", err)
		return nil, nil, nil, errors.New("fail to load template steps: " + flow.Steps +
			", err: " + err.Error())
	}

	// get flow
	return pool, flow, steps, nil
}

func searchStartServiceStep(steps []*models.StepOption) string {
	name := ""
	for _, step := range steps {
		if strings.HasPrefix(step.Name, START_SERVICE) {
			name = step.Name
			break
		}
	}
	return name
}
