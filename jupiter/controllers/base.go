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
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

const (
	BAD_REQUEST     = 400
	SERVICE_ERRROR  = 500
	SERVICE_SUCCESS = 200
)

type ApiResponse struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	// 前端分页
	//PageNumber int         `json:"pageNumber"`
	//PageSize   int         `json:"pageSize"`
	//TotalCount int64       `json:"totalCount"`
	Content interface{} `json:"content"`
	Ext     interface{} `json:"ext"`
}

type Select2Object struct {
	Text string `json:"text"`
	Id   int    `json:"id"`
}

type UserPrivResponse struct {
	Res        interface{} `json:"res"`
	Privileges interface{} `json:"privileges"`
}

const (
	SERVICE_ERR_CODE     = 10999
	NO_AUTH_CODE         = 12995
	APPKEY_FAIL_CODE     = 12994
	APPKEY_BAN_CODE      = 12993
	USER_HEADER_ERR_CODE = 12992
)

var (
	SuccessResp = ApiResponse{}

	// Error
	JsonFormatFaildResp = ApiResponse{
		Code: 12999,
		Msg:  "Please input the correct json format!",
	}

	PageErrorResp = ApiResponse{
		Code: 12998,
		Msg:  "Please input the correct page parameter!",
	}

	MissingParamResp = ApiResponse{
		Code: 12997,
		Msg:  "Missing requisite parameter!",
	}

	InputParseFaildResp = ApiResponse{
		Code: 12996,
		Msg:  "Input error!",
	}

	AuthFaildResp = ApiResponse{
		Code: NO_AUTH_CODE,
		Msg:  "You are not authorized!",
	}

	AppkeyFaildResp = ApiResponse{
		Code: APPKEY_FAIL_CODE,
		Msg:  "Incorrect appkey!",
	}

	AppkeyBanndResp = ApiResponse{
		Code: APPKEY_BAN_CODE,
		Msg:  "Appkey has been banned!",
	}

	LoginFaildResp = ApiResponse{
		Code: 11999,
		Msg:  "Invalid username or password!",
	}

	ServiceErrResp = ApiResponse{
		Code: SERVICE_ERR_CODE,
		Msg:  "Server encountered an error!",
	}

	OverRangeResp = ApiResponse{
		Code: 11998,
		Msg:  "Over range to create",
	}

	ResourceInsufficientResp = ApiResponse{
		Code: 11997,
		Msg:  "resource is not enough to apply",
	}

	OverSizeLimitedResp = ApiResponse{
		Code: 11996,
		Msg:  "Over size limited",
	}

	IpAlreadExistedResp = ApiResponse{
		Code: 11995,
		Msg:  "This ip alread existed",
	}

	NoAuthRespStr        = fmt.Sprintf("{\"code\": %d, \"msg\": \"You are not authorized!\"}", NO_AUTH_CODE)
	AppkeyFaildRespStr   = fmt.Sprintf("{\"code\": %d, \"msg\": \"Incorrect appkey!\"}", APPKEY_FAIL_CODE)
	AppkeyBannedRespStr  = fmt.Sprintf("{\"code\": %d, \"msg\": \"Appkey has been banned!\"}", APPKEY_BAN_CODE)
	UserHeaderErrRespStr = fmt.Sprintf("{\"code\": %d, \"msg\": \"Incorrect username in header!\"}", USER_HEADER_ERR_CODE)
)

var (
	MachineIdFailedResp = ApiResponse{
		Code: 13999,
		Msg:  "please input MachineName !",
	}

	MachineIpFailedResp = ApiResponse{
		Code: 13998,
		Msg:  "please input MachineIp !",
	}

	MachineInsertFaildResp = ApiResponse{
		Code: 13997,
		Msg:  "insert machine is Faild !",
	}

	MachineNotFindResp = ApiResponse{
		Code: 13995,
		Msg:  "can't find the machine !",
	}

	MachineNotFindKPResp = ApiResponse{
		Code: 13994,
		Msg:  "can't find the key primary !",
	}

	MachineReadFailedResp = ApiResponse{
		Code: 13993,
		Msg:  "please input machineId!",
	}

	MachineUpdateFaildResp = ApiResponse{
		Code: 13992,
		Msg:  "update the machine failure",
	}
	MachineRemoveFaildResp = ApiResponse{
		Code: 13991,
		Msg:  "remove the machine failure",
	}

	MachineStopFailedResp = ApiResponse{
		Code: 13990,
		Msg:  "Stop the machine failure!",
	}

	MachineDataFailedStoreResp = ApiResponse{
		Code: 13989,
		Msg:  "Store data of the machine failure!",
	}

	MachineDataFailedRemoveResp = ApiResponse{
		Code: 13988,
		Msg:  "Remove data of the machine failure!",
	}

	MachineUpdatePasswordFaildResp = ApiResponse{
		Code: 13987,
		Msg:  "Machine can't be deleted or starting!",
	}

	MachineDataFailedUpdateResp = ApiResponse{
		Code: 13986,
		Msg:  "Machine can't update in the DB!",
	}

	MachineDuplicateFaildResp = ApiResponse{
		Code: 13985,
		Msg:  "Machine instanceId is duplicate in DB!",
	}

	ResourceUpdateFaildResp = ApiResponse{
		Code: 14998,
		Msg:  "Update the resource failure!",
	}

	ResourceInsertFaildResp = ApiResponse{
		Code: 14997,
		Msg:  "Insert the resource failure!",
	}

	ResourceDeleteFaildResp = ApiResponse{
		Code: 14996,
		Msg:  "Delete the resource failure!",
	}

	ResourceNotFindResp = ApiResponse{
		Code: 14995,
		Msg:  "Can't find the resource!",
	}

	ResourceNotFindKPResp = ApiResponse{
		Code: 14994,
		Msg:  "Can't find the key primary!",
	}

	PoolExistedFailedResp = ApiResponse{
		Code: 15999,
		Msg:  "The pool name already exists",
	}

	PoolUpdateFaildResp = ApiResponse{
		Code: 15998,
		Msg:  "Update the pool failure",
	}

	PoolDeleteFaildResp = ApiResponse{
		Code: 15997,
		Msg:  "Delete the pool failure",
	}

	PoolInsertFaildResp = ApiResponse{
		Code: 15996,
		Msg:  "Insert  the pool failure!",
	}

	PoolNotFindResp = ApiResponse{
		Code: 15995,
		Msg:  "Can't find the pool!",
	}

	PoolNotFindKPResp = ApiResponse{
		Code: 15994,
		Msg:  "Can't find the key primary!",
	}

	MachineSuccessResp = ApiResponse{
		Code: 20000,
		Msg:  "Successful operation!",
	}

	ResourceSuccessResp = ApiResponse{
		Code: 20000,
		Msg:  "Successful operation!",
	}

	PoolSuccessResp = ApiResponse{
		Code: 20000,
		Msg:  "Successful operation!",
	}

	MachineFailedResp = 13000

	// Success
	//MachineSuccessResp = 20000

)

var (
	UserInsertSuccessResp = ApiResponse{
		Code: 20000,
		Msg:  "Insert user successfully!",
	}
	UserInsertFailedResp = ApiResponse{
		Code: 16999,
		Msg:  "Failed to insert user!",
	}
	UserExistedResp = ApiResponse{
		Code: 16998,
		Msg:  "User already exists!",
	}
	UserDeleteSuccessResp = ApiResponse{
		Code: 20000,
		Msg:  "Delete user successfully!",
	}
	UserDeleteFailedResp = ApiResponse{
		Code: 16997,
		Msg:  "Failed to delete user!",
	}
	UserNotFindResp = ApiResponse{
		Code: 16996,
		Msg:  "User is not found!",
	}
	UserSelectFailedResp = ApiResponse{
		Code: 16995,
		Msg:  "Failed to select user!",
	}
	UserSelectSuccessResp = ApiResponse{
		Code: 20000,
		Msg:  "Select user successful!",
	}
	UserUpdateFailedResp = ApiResponse{
		Code: 16994,
		Msg:  "Failed to update user!",
	}
	UserUpdateSuccessResp = ApiResponse{
		Code: 20000,
		Msg:  "Update user successful!",
	}
)

var (
	PrivilegeInsertSuccessResp = ApiResponse{
		Code: 20000,
		Msg:  "insert Priviledge successfully !",
	}
	PrivilegeInsertFailedResp = ApiResponse{
		Code: 17999,
		Msg:  "Failed to insert privilege!",
	}
	PrivilegeExistedResp = ApiResponse{
		Code: 17998,
		Msg:  "Privilege already exists!",
	}
	PrivilegeDeleteSuccessResp = ApiResponse{
		Code: 20000,
		Msg:  "Delete privilege successfully!",
	}
	PrivilegeDeleteFailedResp = ApiResponse{
		Code: 17997,
		Msg:  "Failed to delete privilege!",
	}
	PrivilegeNotFindResp = ApiResponse{
		Code: 17996,
		Msg:  "privilege is not found!",
	}
	PrivilegeSelectFailedResp = ApiResponse{
		Code: 17995,
		Msg:  "Failed to select privilege!",
	}
	PrivilegeSelectSuccessResp = ApiResponse{
		Code: 20000,
		Msg:  "Select privilege succeffful!",
	}
	PrivilegeUpdateFailedResp = ApiResponse{
		Code: 17994,
		Msg:  "Failed to update privilege!",
	}
	PrivilegeUpdateSuccessResp = ApiResponse{
		Code: 20000,
		Msg:  "Update privilege successful!",
	}
)

var (
	UserPrivAddFailedRes = ApiResponse{
		Code: 18999,
		Msg:  "Failed to add your Privileges!",
	}
	UserPrivAddSuccessRes = ApiResponse{
		Code: 20000,
		Msg:  "Add your Privileges successful!",
	}
	UserPrivSelectFailedResp = ApiResponse{
		Code: 18998,
		Msg:  "Failed to add user's Privileges!",
	}
	UserPrivSelectSuccessResp = ApiResponse{
		Code: 20000,
		Msg:  "Select user's Privileges successful!",
	}
	UserPrivDeleteFailedRes = ApiResponse{
		Code: 18997,
		Msg:  "Failed to delete user's Privileges!",
	}
	UserPrivDeleteSuccessResp = ApiResponse{
		Code: 20000,
		Msg:  "Delete user's Privileges successful!",
	}
)
var (
	BehaviorInsertSuccessResp = ApiResponse{
		Code: 20000,
		Msg:  "Insert Behavior successfully!",
	}
	BehaviorInsertFailedResp = ApiResponse{
		Code: 14999,
		Msg:  "Failed to Behavior user!",
	}
	BehaviorSelectFailedResp = ApiResponse{
		Code: 14996,
		Msg:  "Failed to select behavior!",
	}
	BehaviorSelectSuccessResp = ApiResponse{
		Code: 20000,
		Msg:  "Select behavior successful!",
	}
	BehaviorNotFindResp = ApiResponse{
		Code: 14997,
		Msg:  "Behavior is not found!",
	}
)

var (
	DayBookNotFindResp = ApiResponse{
		Code: 19995,
		Msg:  "can't find the dayBook !",
	}
	DayBookNotFindKPResp = ApiResponse{
		Code: 19994,
		Msg:  "can't find the key primary !",
	}
	DayBookSuccessResp = ApiResponse{
		Code: 20000,
		Msg:  "Successful operation!",
	}
	DayBooksNotFindResp = ApiResponse{
		Code: 19996,
		Msg:  "can't find some dayBooks !",
	}
	DayBookUpdateFaildResp = ApiResponse{
		Code: 19998,
		Msg:  "update the daybook failure",
	}

	DayBookInsertFaildResp = ApiResponse{
		Code: 19997,
		Msg:  "insert the DayBook failure",
	}

	DayBookDeleteFaildResp = ApiResponse{
		Code: 19999,
		Msg:  "delete the DayBook failure",
	}
)

type BaseController struct {
	beego.Controller
	ApiResponse ApiResponse
	Status      int
}

func (c *BaseController) GetMissingParamResp(param string) {
	r := MissingParamResp
	r.Msg = "Missing requisite parameter(" + param + ")!"
	c.ApiResponse = r
}

func (c *BaseController) GetServiceErrResp(err string) {
	r := ServiceErrResp
	r.Msg = err
	c.ApiResponse = r
}

func (c *BaseController) GetOverSizeLimitedResp(param string, rangeSize string) {
	r := OverSizeLimitedResp
	r.Msg = "The paramter (" + param + ") over size limited and the size range is " + rangeSize + "!"
	c.ApiResponse = r
}

func (c *BaseController) GetPageErrorResp(err string) {
	r := PageErrorResp
	r.Msg = err
	c.ApiResponse = r
}


func (c *BaseController) GetExistedIpResp(msg string) {
	r := IpAlreadExistedResp
	r.Msg = msg
	c.ApiResponse = r
}

func (c *BaseController) RespJsonWithStatus() {
	c.Data["json"] = c.ApiResponse
	c.Ctx.Output.SetStatus(c.Status)
	c.Ctx.Request.Header.Add("Status", strconv.Itoa(c.Status))
	c.ServeJSON()
}

func (c *BaseController) RespMissingParams(param string) {
	c.GetMissingParamResp(param)
	c.Status = BAD_REQUEST
	c.RespJsonWithStatus()
}

func (c *BaseController) RespInputOverLimited(param string, rangeSize string) {
	c.GetOverSizeLimitedResp(param, rangeSize)
	c.Status = BAD_REQUEST
	c.RespJsonWithStatus()
}

func (c *BaseController) RespPageError(err error) {
	c.GetPageErrorResp(err.Error())
	c.Status = BAD_REQUEST
	c.RespJsonWithStatus()
}

func (c *BaseController) RespIpExisted(msg string) {
	c.GetExistedIpResp(msg)
	c.Status = BAD_REQUEST
	c.RespJsonWithStatus()
}

func (c *BaseController) RespInputError() {
	c.ApiResponse = InputParseFaildResp
	c.Status = BAD_REQUEST
	c.RespJsonWithStatus()
}

func (c *BaseController) RespServiceError(err error) {
	c.GetServiceErrResp(err.Error())
	c.Status = SERVICE_ERRROR
	c.RespJsonWithStatus()
}
