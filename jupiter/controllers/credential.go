package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about credential
type CredentialController struct {
	beego.Controller
}

// @Title Add credential
// @Description Add providers credentials.
// @Param	body		body 	models.Credential	true		"body for credential content"
// @Success 200 {object} models.Credential
// @Failure 403 body is empty
// @router / [post]
func (c *CredentialController) Post() {
	c.ServeJSON()
}

// @Title Authorization check
// @Description Check providers uthorization.
// @Success 200 True
// @Failure 402 False
// @router /authorization [post]
func (c *CredentialController) Authorize() {
	c.ServeJSON()
}
