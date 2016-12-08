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

package models

import (
	"encoding/json"
)

/**
项目列表返回信息
 */
type ProjectListResponse struct {
	Code       int         `json:"code"`
	Content    interface{} `json:"content"`
	ErrMsg     string      `json:"errMsg"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalCount int         `json:"total_count"`
}

func BuildProjectListResponse(code int, content interface{}, errMsg string, page, page_size, total_count int) ProjectListResponse {
	return ProjectListResponse{
		Code:       code,
		Content:    content,
		ErrMsg:     errMsg,
		Page:       page,
		PageSize:   page_size,
		TotalCount: total_count}
}

func (r ProjectListResponse) ToString() string {
	bytes, _ := json.Marshal(r)
	return string(bytes)
}
