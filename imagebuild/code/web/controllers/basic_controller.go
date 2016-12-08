/*
 *  Copyright 2009-2016 Weibo, Inc.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
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
