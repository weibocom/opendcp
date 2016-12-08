// Copyright 2016 Weibo Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"weibo.com/opendcp/jupiter/models"
	"weibo.com/opendcp/jupiter/service/bill"
	"weibo.com/opendcp/jupiter/service/organization"
)

// Operations about organization
type OrganizationController struct {
	BaseController
}

// @Title List organizations.
// @Description List all organizations.
// @router / [get]
func (oc *OrganizationController) ListAllOrganizations() {
	var organizations []models.Organization
	var err error
	organizations, err = organization.ListAll()
	if err != nil {
		beego.Error("get organization err: ", err)
		oc.RespServiceError(err)
		return
	}
	resp := ApiResponse{}
	resp.Content = organizations
	oc.ApiResponse = resp
	oc.Status = SERVICE_SUCCESS
	oc.RespJsonWithStatus()
}

// @Title Get a organization.
// @Description List a organization.
// @router /:organizationId [get]
func (oc *OrganizationController) GetOrganizationById() {
	organizationId, err := oc.GetInt64(":organizationId")
	if err != nil {
		beego.Error("parse organization id err: ", err)
		oc.RespInputError()
		return
	}
	og, err := organization.GetOrganization(organizationId)
	if err != nil {
		beego.Error("get one organization err: ", err)
		oc.RespServiceError(err)
		return
	}
	resp := ApiResponse{}
	resp.Content = og
	oc.ApiResponse = resp
	oc.Status = SERVICE_SUCCESS
	oc.RespJsonWithStatus()
}

// @Title List instances in organization
// @Description List all instances in organization.
// @router /instances/:organizationId [get]
func (oc *OrganizationController) GetInstancesByOrganizationId() {
	organizationId, err := oc.GetInt64(":organizationId")
	if err != nil {
		beego.Error("parse organization id err: ", err)
		oc.RespInputError()
	}
	resp := ApiResponse{}
	instances, err := organization.GetInstancesByOrganizationId(organizationId)
	if err != nil {
		beego.Error("get instances in organization, error: ", err)
		oc.RespServiceError(err)
	}
	resp.Content = instances
	oc.ApiResponse = resp
	oc.Status = SERVICE_SUCCESS
	oc.RespJsonWithStatus()
}

// @Title New organization
// @Description New a organization.
// @router / [post]
func (oc *OrganizationController) CreateOrganization() {
	body := oc.Ctx.Input.RequestBody
	var og models.Organization
	err := json.Unmarshal(body, &og)
	if err != nil {
		beego.Error("parse organization body err: ", err)
		oc.RespInputError()
	}
	if og.Name == "" {
		beego.Error("Organization name is null")
		oc.RespMissingParams("organization.Name")
	}
	err = organization.New(&og)
	if err != nil {
		beego.Error("New organization err: ", err)
		oc.RespServiceError(err)
	}
	resp := ApiResponse{}
	resp.Content = true
	oc.ApiResponse = resp
	oc.Status = SERVICE_SUCCESS
	oc.RespJsonWithStatus()
}

// @Title Delete organization
// @Description Delete a organization.
// @router /:organizationId [delete]
func (oc *OrganizationController) DeleteOrganization() {
	organizationId, err := oc.GetInt64(":organizationId")
	if err != nil {
		beego.Error("parse organization id err: ", err)
		oc.RespInputError()
	}
	err = organization.Delete(organizationId)
	if err != nil {
		beego.Error("del organization err: ", err)
		oc.RespServiceError(err)
		return
	}
	resp := ApiResponse{}
	resp.Content = true
	oc.ApiResponse = resp
	oc.Status = SERVICE_SUCCESS
	oc.RespJsonWithStatus()
}

// @Title List clusters on a organization
// @Description List clusters on a organization.
// @Success 200 {object} models.Cluster
// @Failure 403 body is empty
// @router /cluster/:organizationId [get]
func (oc *OrganizationController) GetClustersByOrganizationId() {
	organizationId, err := oc.GetInt64(":organizationId")
	if err != nil {
		beego.Error("parse organization id err: ", err)
		oc.RespInputError()
		return
	}
	resp := ApiResponse{}
	clusters, err := organization.GetClusters(organizationId)
	if err != nil {
		beego.Error("get organization clusters err: ", err)
		oc.RespServiceError(err)
		return
	}
	resp.Content = clusters
	oc.ApiResponse = resp
	oc.Status = SERVICE_SUCCESS
	oc.RespJsonWithStatus()
}

// @Title Get one cluster usage
// @Description GetUsage handles request to /organiazation/usage
// @router /usage/:clusterId [get]
func (oc *OrganizationController) GetUsage() {
	resp := ApiResponse{}
	clusterId, err := oc.GetInt64(":clusterId")
	if err != nil {
		beego.Error("parse cluster id err: ", err)
		oc.RespInputError()
		return
	}
	hours, err := bill.GetCosts(clusterId)
	if err != nil {
		beego.Error("get someone cluster usage in organization error:", err)
		oc.RespServiceError(err)
		return
	}
	resp.Content = hours
	oc.ApiResponse = resp
	oc.Status = SERVICE_SUCCESS
	oc.RespJsonWithStatus()
}

// @Title Increase one cluster credit
// @Description IncreaseCredit handles request to /organiazation/credit
// @router /credit/:clusterId/:hours [post]
func (oc *OrganizationController) IncreaseCredit() {
	resp := ApiResponse{}
	clusterId, err := oc.GetInt64(":clusterId")
	if err != nil {
		beego.Error("parse cluster id err: ", err)
		oc.RespInputError()
		return
	}
	hours, err := oc.GetInt(":hours")
	if err != nil {
		beego.Error("parse hours err: ", err)
		oc.RespInputError()
		return
	}
	if hours < 0 {
		oc.RespInputOverLimited("hours", "need to large 0.")
	}
	ret, err := bill.IncreaseCredit(clusterId, hours)
	if err != nil {
		beego.Error("Increase credit error:", err)
		oc.RespServiceError(err)
		return
	}
	resp.Content = ret
	oc.ApiResponse = resp
	oc.Status = SERVICE_SUCCESS
	oc.RespJsonWithStatus()
}

// @Title Get one cluster credit
// @Description GetCredit handles request to /organiazation/credit
// @router /credit/:clusterId [get]
func (oc *OrganizationController) GetCredit() {
	resp := ApiResponse{}
	clusterId, err := oc.GetInt64(":clusterId")
	if err != nil {
		beego.Error("parse cluster id err: ", err)
		oc.RespInputError()
		return
	}
	hours, err := bill.GetCredit(clusterId)
	if err != nil {
		beego.Error("get someone cluster usage in organization error:", err)
		oc.RespServiceError(err)
		return
	}
	resp.Content = hours
	oc.ApiResponse = resp
	oc.Status = SERVICE_SUCCESS
	oc.RespJsonWithStatus()
}

// @Title Get one cluster credit
// @Description GetCredit handles request to /organiazation/credit
// @router /bill/:clusterId [get]
func (oc *OrganizationController) GetBill() {
	resp := ApiResponse{}
	clusterId, err := oc.GetInt64(":clusterId")
	if err != nil {
		beego.Error("parse cluster id err: ", err)
		oc.RespInputError()
		return
	}
	theBill, err := bill.GetBill(clusterId)
	if err != nil {
		beego.Error("get someone cluster bill in organization error:", err)
		oc.RespServiceError(err)
		return
	}
	costsCreditMap := make(map[string]int)
	costsCreditMap["costs"] = theBill.Costs
	costsCreditMap["credit"] = theBill.Credit
	beego.Info(costsCreditMap)
	resp.Content = costsCreditMap
	oc.ApiResponse = resp
	oc.Status = SERVICE_SUCCESS
	oc.RespJsonWithStatus()
}
