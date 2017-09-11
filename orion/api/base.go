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

package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/astaxie/beego"
)

/*
*	base api interface as a controller
 */
type baseAPI struct {
	beego.Controller
}

type apiResponse struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Content interface{} `json:"data"`
}

type pageAPIResponse struct {
	apiResponse
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Count    int `json:"query_count"`
}

const (
	HTTP_SUCCESS = 200
)

/*
*	read all request parameter to a struct
 */
func (b *baseAPI) Body2Json(obj interface{}) error {
	bytes, err := ioutil.ReadAll(b.Ctx.Request.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(bytes[:]))
	return json.Unmarshal(bytes, obj)
}

/*
*	get int request parameter
 */
func (b *baseAPI) Query2Int(key string, def int) int {
	val := b.Ctx.Input.Query(key)
	valInt, err := strconv.Atoi(val)
	if err != nil {
		valInt = def
	}
	return valInt
}

/*
*	check and modify wrong paging data
 */
func (b *baseAPI) CheckPage(page, pageSize *int) {
	if *page <= 0 {
		*page = 1
	}
	if *pageSize <= 0 {
		*pageSize = 10
	}
}

/*
*	return request with success and data
 */
func (b *baseAPI) ReturnSuccess(content interface{}) {
	if content == nil {
		content = struct{}{}
	}
	b.Data["json"] = apiResponse{
		Code:    0,
		Content: content,
	}
	b.Ctx.Output.SetStatus(HTTP_SUCCESS)
	b.ServeJSON()
}

/*
*	return request with paged data
 */
func (b *baseAPI) ReturnPageContent(page, page_size, count int, content interface{}) {
	b.Data["json"] = pageAPIResponse{
		apiResponse: apiResponse{
			Code:    0,
			Content: content,
		},
		Page:     page,
		PageSize: page_size,
		Count:    count,
	}
	b.Ctx.Output.SetStatus(HTTP_SUCCESS)
	b.ServeJSON()
}

/*
*	return request with fail info
 */
func (b *baseAPI) ReturnFailed(msg string, code int) {
	b.Data["json"] = apiResponse{
		Code:    code,
		Msg:     msg,
		Content: struct{}{},
	}
	b.Ctx.Output.SetStatus(code)
	b.ServeJSON()
}

func (b *baseAPI) ReturnFailedWithContent(content interface{}, code int) {
	if content == nil {
		content = struct{}{}
	}
	b.Data["json"] = apiResponse{
		Code:    code,
		Content: content,
	}
	b.Ctx.Output.SetStatus(code)
	b.ServeJSON()
}
