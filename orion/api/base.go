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
