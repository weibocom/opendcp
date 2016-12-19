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
