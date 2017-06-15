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
	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"weibo.com/opendcp/imagebuild/code/util"
	"weibo.com/opendcp/imagebuild/code/web/models"
)

/**
基础controller
 */
type BasicController struct {
	beego.Controller
}

func (c *BasicController) Operator() string {
	operator := c.GetString("operator")
	if operator == "" {
		operator = c.Ctx.Request.Header.Get("Authorization")
	}

	return operator
}

func (c *BasicController) BizName() string{
	bizName := c.GetString("X-Biz-Name")
	if bizName == "" {
		bizName = c.Ctx.Request.Header.Get("X-Biz-Name")
	}

	return bizName
}
func (c *BasicController) Prepare() {
	// 1. Check URL arguments.
	lang := c.Input().Get("lang")

	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = c.Ctx.GetCookie("lang")
	}

	// Check again in case someone modify by purpose.
	if !i18n.IsExist(lang) {
		lang = ""
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := c.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}

	// 4. Default language is English.
	if len(lang) == 0 {
		lang = "zh-CN"
	}

	util.Lang = lang
	// Set language properties.
	c.Data["Lang"] = lang
	models.AppServer.Lang = lang
	log.Infof("BasicController prepare lang:%s", lang)
}
