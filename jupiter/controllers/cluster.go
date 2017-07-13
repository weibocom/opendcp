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

package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"weibo.com/opendcp/jupiter/models"
	"weibo.com/opendcp/jupiter/service/bill"
	"weibo.com/opendcp/jupiter/service/cluster"
	"weibo.com/opendcp/jupiter/service/instance"
	"strconv"
	"strings"
)

// Operations about cluster
type ClusterController struct {
	BaseController
}

// @Title List clusters.
// @Description list all cluster.
// @router / [get]
func (clusterController *ClusterController) GetClusters() {
	resp := ApiResponse{}
	theCluster, err := cluster.ListClusters()
	if err != nil {
		beego.Error("get one cluster err: ", err)
		clusterController.RespServiceError(err)
		return
	}
	resp.Content = theCluster
	clusterController.ApiResponse = resp
	clusterController.Status = SERVICE_SUCCESS
	clusterController.RespJsonWithStatus()
}

// @Title Get a cluster.
// @Description Get a cluster infomation.
// @Success 200 {object} models.Cluster
// @Failure 403 body is empty
// @router /:clusterId [get]
func (clusterController *ClusterController) GetClusterInfo() {
	clusterId, err := clusterController.GetInt64(":clusterId")
	if err != nil {
		beego.Error("parse cluster id err: ", err)
		clusterController.RespInputError()
		return
	}
	resp := ApiResponse{}
	theCluster, err := cluster.GetCluster(clusterId)
	if err != nil {
		beego.Error("get one cluster err: ", err)
		clusterController.RespServiceError(err)
		return
	}
	resp.Content = theCluster
	clusterController.ApiResponse = resp
	clusterController.Status = SERVICE_SUCCESS
	clusterController.RespJsonWithStatus()
}

// @Title Create cluster
// @Description Create cluster.
// @router / [post]
func (clusterController *ClusterController) CreateCluster() {
	var theCluster models.Cluster
	body := clusterController.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &theCluster)
	if err != nil {
		beego.Error("Could parase request before crate cluster: ", err)
		clusterController.RespInputError()
		return
	}
	if theCluster.Provider == "aliyun" {
		if theCluster.DataDiskNum > 4 || theCluster.DataDiskNum < 0 {
			clusterController.RespInputOverLimited("DataDiskNum", "larger than 0 and less equal 4.")
			return
		}
		switch theCluster.DataDiskCategory {
		case "cloud":
			if theCluster.DataDiskSize > 2000 || theCluster.DataDiskSize < 5 {
				clusterController.RespInputOverLimited("DdatCategory", "larger or equal 5 and less or equal 2000G")
				return
			}
		case "cloud_efficiency":
			if theCluster.DataDiskSize > 32768 || theCluster.DataDiskSize < 20 {
				clusterController.RespInputOverLimited("DataCategory", "larger or equal 20 and less or equal 32768G")
				return
			}
		case "cloud_ssd":
			if theCluster.DataDiskSize > 32768 || theCluster.DataDiskSize < 20 {
				clusterController.RespInputOverLimited("DataCategory", "larger or equal 20 and less or equal 32768G")
				return
			}
		case "ephemeral_ssd":
			if theCluster.DataDiskSize > 800 || theCluster.DataDiskSize < 5 {
				clusterController.RespInputOverLimited("DataCategory", "larger or equal 5 and less or equal 800G")
				return
			}
		}
		if theCluster.Network.InternetChargeType == "PayByBandwidth" {
			if theCluster.Network.InternetMaxBandwidthOut > 100 || theCluster.Network.InternetMaxBandwidthOut < 0 {
				clusterController.RespInputOverLimited("InternetMaxBandwidthOut", "larger or equal 0 and less or equal 100")
				return
			}
		}
		if theCluster.Network.InternetChargeType == "PayByTraffic" {
			if theCluster.Network.InternetMaxBandwidthOut > 100 || theCluster.Network.InternetMaxBandwidthOut < 1 {
				clusterController.RespInputOverLimited("InternetMaxBandwidthOut", "larger or equal 0 and less or equal 100")
				return
			}
		}
	}
	id, err := cluster.CreateCluster(&theCluster)
	if err != nil {
		beego.Error("Ceate cluster err: ", err)
		clusterController.RespServiceError(err)
		return
	}
	resp := ApiResponse{}
	resp.Content = id
	clusterController.ApiResponse = resp
	clusterController.Status = SERVICE_SUCCESS
	clusterController.RespJsonWithStatus()
}

// @Title Delete cluster
// @Description Delete cluster.
// @router /:clusterId [delete]
func (clusterController *ClusterController) DeleteCluster() {
	clusterId, err := clusterController.GetInt64(":clusterId")
	if err != nil {
		beego.Error("parse cluster id err: ", err)
		clusterController.RespInputError()
		return
	}
	isDeleted, err := cluster.DeleteCluster(clusterId)
	if err != nil {
		beego.Error("Delete cluster err: ", err)
		clusterController.RespServiceError(err)
		return
	}
	resp := ApiResponse{}
	resp.Content = isDeleted
	clusterController.ApiResponse = resp
	clusterController.Status = SERVICE_SUCCESS
	clusterController.RespJsonWithStatus()
}

// @Title Expand cluster
// @Description Expand cluster.
// @Param	body		int 	count 	true "the number of instance"
// @Success 200 body models.Cluster
// @Failure 403 body is empty
// @router /:clusterId/expand/:number [post]
func (clusterController *ClusterController) ExpandInstances() {
	correlationId := clusterController.Ctx.Input.Header("X-CORRELATION-ID")
	if len(correlationId) <= 0 {
		clusterController.RespMissingParams("X-CORRELATION-ID")
		return
	}
	resp := ApiResponse{}
	clusterId, err := clusterController.GetInt64(":clusterId")
	if err != nil {
		beego.Error("Need to pass vaild cluster id: ", err)
		clusterController.RespInputError()
		return
	}
	expandNumber, err := clusterController.GetInt(":number")
	if err != nil {
		beego.Error("Need to pass vaild expand number: ", err)
		clusterController.RespInputError()
		return
	}
	if clusterId < 1 || expandNumber < 1 {
		beego.Error("the cluster id and expand number need to large than 0, now cluster id is", clusterId, "and expand number is", expandNumber)
		clusterController.ApiResponse = InputParseFaildResp
		return
	}
	theCluster, err := cluster.GetCluster(clusterId)
	if err != nil {
		beego.Error("Get cluster error:", err)
		clusterController.RespServiceError(err)
		return
	}
	if theCluster.Provider == instance.PhyDev {
		msg := "physical device can't expand."
		err = errors.New(msg)
		beego.Error(err)
		clusterController.RespServiceError(err)
		return
	}
	if !bill.CanCreate(theCluster) {
		err = fmt.Errorf("Sorry, over credit limit.")
		resp.Msg = err.Error()
		clusterController.ApiResponse = resp
		clusterController.Status = BAD_REQUEST
		clusterController.RespJsonWithStatus()
		return
	}
	instanceIds, err := cluster.Expand(theCluster, expandNumber, correlationId)
	if len(instanceIds) == 0 {
		beego.Error("expand instances failed:", err)
		clusterController.RespServiceError(err)
		return
	}
	resp.Content = instanceIds
	resp.Ext = clusterId
	clusterController.ApiResponse = resp
	clusterController.Status = SERVICE_SUCCESS
	clusterController.RespJsonWithStatus()
}


// @Title get total instances number
// @Description get total instances number
// @router /number/:hour [get]
func (cc *ClusterController) GetInstancesNumber() {
	hour, err := cc.GetInt(":hour")
	if err != nil {
		beego.Error("Can't parse the hour err:", err)
		cc.RespInputError()
		return
	}

	var details []models.InstanceDetail
	if hour == 0 {
		details, err = cluster.GetLatestInstanceDetail()
		if err != nil {
			beego.Error("Get instance detail err:", err)
			cc.RespServiceError(err)
			return
		}
	} else {
		details, err = cluster.GetRecentInstanceDetail(hour)
		if err != nil {
			beego.Error("Get instance detail err:", err)
			cc.RespServiceError(err)
			return
		}
	}

	result := make([] map[string]string,0)
	for _, detail := range details {
		info := make(map[string] string)
		info["time"] = detail.RunningTime
		for k, v := range detail.InstanceNumber {
			info[k] = strconv.Itoa(v)
		}
		result = append(result, info)
	}

	resp := ApiResponse{}
	resp.Content = result
	cc.ApiResponse = resp
	cc.Status = SERVICE_SUCCESS
	cc.RespJsonWithStatus()
}

// @Title get past instances number
// @Description get past instances number
// @router /oldnumber/:time [get]
func (cc *ClusterController) GetPastInstancesNumber() {
	timeStr := cc.GetString(":time")
	specificTime := strings.Replace(timeStr,"%20"," ", -1)
 	detail, err := cluster.GetPastInstanceDetail(specificTime)
	if err != nil {
		beego.Error("Get instance detail at the time err:", err)
		cc.RespServiceError(err)
		return
	}
	result := make(map[string]string)
	result["time"] = detail.RunningTime
	for k, v := range detail.InstanceNumber {
		result[k] = strconv.Itoa(v)
	}
	resp := ApiResponse{}
	resp.Content = result
	cc.ApiResponse = resp
	cc.Status = SERVICE_SUCCESS
	cc.RespJsonWithStatus()
}

// @Title update instances number
// @Description update instances number
// @router /update [get]
func (cc *ClusterController) UpdateInstanceInfo() {
	err := cluster.UpdateInstanceDetail()
	if err != nil {
		beego.Error("Update instances detail err", err)
		cc.RespServiceError(err)
		return
	}
	resp := ApiResponse{}
	resp.Content = true
	cc.ApiResponse = resp
	cc.Status = SERVICE_SUCCESS
	cc.RespJsonWithStatus()
}
