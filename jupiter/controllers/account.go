package controllers

import (
	"strconv"
	"github.com/astaxie/beego"
	"weibo.com/opendcp/jupiter/service/account"
	"weibo.com/opendcp/jupiter/models"
	"encoding/json"
	"weibo.com/opendcp/jupiter/service/instance"
	"io/ioutil"
)


type AccountController struct {
	BaseController
}

// @Title List accounts.
// @Description list all accounts.
// @router / [get]
func (ac *AccountController) GetAllAccounts()  {
	bizId := ac.Ctx.Input.Header("X-Biz-ID")
	bid, err := strconv.Atoi(bizId)
	if bizId=="" || err != nil {
		beego.Error("Get X-Biz-ID err!")
		ac.RespInputError()
		return
	}
	accounts, err := account.ListAccounts(bid)
	if err != nil {
		beego.Error("Get all accounts err: ", err)
		ac.RespServiceError(err)
		return
	}

	resp := ApiResponse{}
	resp.Content = accounts
	ac.ApiResponse = resp
	ac.Status = SERVICE_SUCCESS
	ac.RespJsonWithStatus()
}

// @Title Get a account.
// @Description Get a account infomation.
// @Success 200 {object} models.Account
// @Failure 403 body is empty
// @router /:provider [get]
func (ac *AccountController) GetAccountInfo()  {
	bizId := ac.Ctx.Input.Header("X-Biz-ID")
	bid, err := strconv.Atoi(bizId)
	if bizId=="" || err != nil {
		beego.Error("Get X-Biz-ID err!")
		ac.RespInputError()
		return
	}
	provider := ac.GetString(":provider")
	theAccount, err := account.GetAccount(bid, provider)
	if err != nil {
		beego.Error("Get account info err: ", err)
		ac.RespServiceError(err)
		return
	}

	resp := ApiResponse{}
	resp.Content = theAccount
	ac.ApiResponse = resp
	ac.Status = SERVICE_SUCCESS
	ac.RespJsonWithStatus()
}

// @Title Create account
// @Description Create account.
// @router / [post]
func (ac *AccountController) CreateAccount() {
	bizId := ac.Ctx.Input.Header("X-Biz-ID")
	bid, err := strconv.Atoi(bizId)
	if bizId == "" || err != nil {
		beego.Error("Get X-Biz-ID err!")
		ac.RespInputError()
		return
	}

	var  theAccount models.Account
	body := ac.Ctx.Input.RequestBody
	err = json.Unmarshal(body, &theAccount)
	if err != nil {
		beego.Error("Could parse request before the request: ", err)
		ac.RespInputError()
		return
	}
	theAccount.BizId = bid

	instances, err := instance.ListTestingInstances(bid, theAccount.Provider)
	if err != nil {
		beego.Error("Get testing instances err: ", err)
		ac.RespServiceError(err)
		return
	}
	beego.Warn("Create account: ", len(instances))

	err = instance.DeleteInstances(instances, bid)
	if err != nil {
		beego.Error("Delete testing instances err: ", err)
		ac.RespServiceError(err)
		return
	}

	id, err := account.CreateAccount(&theAccount)
	if err != nil {
		beego.Error("Create account err: ", err)
		ac.RespServiceError(err)
		return
	}

	resp := ApiResponse{}
	resp.Content = id
	ac.ApiResponse = resp
	ac.Status = SERVICE_SUCCESS
	ac.RespJsonWithStatus()
}

// @Title Delete account
// @Description Delete account.
// @router /:provider [delete]
func (ac *AccountController) DeleteAccount()  {
	bizId := ac.Ctx.Input.Header("X-Biz-ID")
	bid, err := strconv.Atoi(bizId)
	if bizId=="" || err != nil {
		beego.Error("Get X-Biz-ID err!")
		ac.RespInputError()
		return
	}
	provider := ac.GetString(":provider")
	isDeleted, err := account.DeleteAccount(bid, provider)
	if err != nil {
		beego.Error("Get account info err: ", err)
		ac.RespServiceError(err)
		return
	}

	resp := ApiResponse{}
	resp.Content = isDeleted
	ac.ApiResponse = resp
	ac.Status = SERVICE_SUCCESS
	ac.RespJsonWithStatus()
}

// @Title update account
// @Description update account.
// @router /update
func (ac *AccountController) UpdateAccount()  {
	bizId := ac.Ctx.Input.Header("X-Biz-ID")
	bid, err := strconv.Atoi(bizId)
	if bizId=="" || err != nil {
		beego.Error("Get X-Biz-ID err!")
		ac.RespInputError()
		return
	}
	bytes, err := ioutil.ReadAll(ac.Ctx.Request.Body)
	if err != nil {
		beego.Error("Get Request Body err: ", err)
		ac.RespServiceError(err)
		return
	}
	obj := &models.Account{}
	err = json.Unmarshal(bytes, obj)
	if err != nil {
		beego.Error("Unmarshal bytes to account err: ", err)
		ac.RespServiceError(err)
		return
	}
	obj.BizId = bid
	obj.KeySecret = account.Encode(obj.KeySecret)
	fields  := []string{
		"KeyId",
		"KeySecret",
	}

	err = account.UpdateAccountInfo(obj,fields)
	if err != nil {
		beego.Error("Get account info err: ", err)
		ac.RespServiceError(err)
		return
	}

	resp := ApiResponse{}
	resp.Content = true
	ac.ApiResponse = resp
	ac.Status = SERVICE_SUCCESS
	ac.RespJsonWithStatus()
}


// @Title get cost
// @Description Get cost info
// @router /cost [get]
func (ac *AccountController) GetCost() {
	bizId := ac.Ctx.Input.Header("X-Biz-ID")
	bid, err := strconv.Atoi(bizId)
	if bizId=="" || err != nil {
		beego.Error("Get X-Biz-ID err!")
		ac.RespInputError()
		return
	}

	provider := ac.GetString("provider")

	cost, err := instance.GetCost(bid, provider)
	if err != nil {
		beego.Error("Get cost err: ", err)
		ac.RespServiceError(err)
		return
	}

	resp := ApiResponse{}
	resp.Content = cost
	ac.ApiResponse = resp
	ac.Status = SERVICE_SUCCESS
	ac.RespJsonWithStatus()
}

// @Title get cost
// @Description Get cost info
// @router /exist [get]
func (ac *AccountController) IsExist() {
	bizId := ac.Ctx.Input.Header("X-Biz-ID")
	bid, err := strconv.Atoi(bizId)
	if bizId=="" || err != nil {
		beego.Error("Get X-Biz-ID err!")
		ac.RespInputError()
		return
	}

	provider := ac.GetString("provider")
	isExist := instance.IsAccountExist(bid, provider)

	resp := ApiResponse{}
	resp.Content = isExist
	ac.ApiResponse = resp
	ac.Status = SERVICE_SUCCESS
	ac.RespJsonWithStatus()
}