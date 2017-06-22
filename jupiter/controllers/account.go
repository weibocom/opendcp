package controllers

import (
	"strconv"
	"github.com/astaxie/beego"
	"weibo.com/opendcp/jupiter/service/account"
	"weibo.com/opendcp/jupiter/models"
	"encoding/json"
	"io/ioutil"
)

type AccountController struct {
	BaseController
}

// @Title List accounts.
// @Description list all accounts.
// @router / [get]
func (accountController *AccountController) GetAllAccounts()  {
	bizId := accountController.Ctx.Input.Header("X-Biz-ID")
	bid, err := strconv.Atoi(bizId)
	if bizId=="" || err != nil {
		beego.Error("Get X-Biz-ID err!")
		accountController.RespInputError()
		return
	}
	accounts, err := account.ListAccounts(bid)
	if err != nil {
		beego.Error("Get all accounts err: ", err)
		accountController.RespServiceError(err)
		return
	}

	resp := ApiResponse{}
	resp.Content = accounts
	accountController.ApiResponse = resp
	accountController.Status = SERVICE_SUCCESS
	accountController.RespJsonWithStatus()
}

// @Title Get a account.
// @Description Get a account infomation.
// @Success 200 {object} models.Account
// @Failure 403 body is empty
// @router /:provider [get]
func (accountController *AccountController) GetAccountInfo()  {
	bizId := accountController.Ctx.Input.Header("X-Biz-ID")
	bid, err := strconv.Atoi(bizId)
	if bizId=="" || err != nil {
		beego.Error("Get X-Biz-ID err!")
		accountController.RespInputError()
		return
	}
	provider := accountController.GetString(":provider")
	theAccount, err := account.GetAccount(bid, provider)
	if err != nil {
		beego.Error("Get account info err: ", err)
		accountController.RespServiceError(err)
		return
	}

	resp := ApiResponse{}
	resp.Content = theAccount
	accountController.ApiResponse = resp
	accountController.Status = SERVICE_SUCCESS
	accountController.RespJsonWithStatus()
}

// @Title Create account
// @Description Create account.
// @router / [post]
func (accountController *AccountController) CreateAccount() {
	bizId := accountController.Ctx.Input.Header("X-Biz-ID")
	bid, err := strconv.Atoi(bizId)
	if bizId == "" || err != nil {
		beego.Error("Get X-Biz-ID err!")
		accountController.RespInputError()
		return
	}

	var  theAccount models.Account
	body := accountController.Ctx.Input.RequestBody
	err = json.Unmarshal(body, &theAccount)
	if err != nil {
		beego.Error("Could parse request before the request: ", err)
		accountController.RespInputError()
		return
	}
	theAccount.BizId = bid

	id, err := account.CreateAccount(&theAccount)
	if err != nil {
		beego.Error("Create account err: ", err)
		accountController.RespServiceError(err)
		return
	}

	resp := ApiResponse{}
	resp.Content = id
	accountController.ApiResponse = resp
	accountController.Status = SERVICE_SUCCESS
	accountController.RespJsonWithStatus()
}

// @Title Delete account
// @Description Delete account.
// @router /:provider [delete]
func (accountController *AccountController) DeleteAccount()  {
	bizId := accountController.Ctx.Input.Header("X-Biz-ID")
	bid, err := strconv.Atoi(bizId)
	if bizId=="" || err != nil {
		beego.Error("Get X-Biz-ID err!")
		accountController.RespInputError()
		return
	}
	provider := accountController.GetString(":provider")
	isDeleted, err := account.DeleteAccount(bid, provider)
	if err != nil {
		beego.Error("Get account info err: ", err)
		accountController.RespServiceError(err)
		return
	}

	resp := ApiResponse{}
	resp.Content = isDeleted
	accountController.ApiResponse = resp
	accountController.Status = SERVICE_SUCCESS
	accountController.RespJsonWithStatus()
}

// @Title update account
// @Description update account.
// @router /update
func (accountController *AccountController) UpdateAccount()  {
	bizId := accountController.Ctx.Input.Header("X-Biz-ID")
	bid, err := strconv.Atoi(bizId)
	if bizId=="" || err != nil {
		beego.Error("Get X-Biz-ID err!")
		accountController.RespInputError()
		return
	}
	bytes, err := ioutil.ReadAll(accountController.Ctx.Request.Body)
	if err != nil {
		beego.Error("Get Request Body err: ", err)
		accountController.RespServiceError(err)
		return
	}
	obj := &models.Account{}
	err = json.Unmarshal(bytes, obj)
	if err != nil {
		beego.Error("Unmarshal bytes to account err: ", err)
		accountController.RespServiceError(err)
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
		accountController.RespServiceError(err)
		return
	}

	resp := ApiResponse{}
	resp.Content = true
	accountController.ApiResponse = resp
	accountController.Status = SERVICE_SUCCESS
	accountController.RespJsonWithStatus()
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

	cost, err := account.GetCost(bid, provider)
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
	isExist := account.IsAccountExist(bid, provider)

	resp := ApiResponse{}
	resp.Content = isExist
	ac.ApiResponse = resp
	ac.Status = SERVICE_SUCCESS
	ac.RespJsonWithStatus()
}