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

type RetryOption struct {
	RetryTimes  int  `json:"retry_times"`
	IgnoreError bool `json:"ignore_error"`
}

type ParamValues map[string]interface{}

type StepOption struct {
	Name   string       `json:"name"`
	Values ParamValues  `json:"param_values"`
	Retry  *RetryOption `json:"retry"`
}

//任务流定义
type FlowImpl struct {
	Id    int    `json:"id" orm:"pk;auto"`
	Name  string `json:"name" orm:"size(50);unique"`
	Desc  string `json:"name" orm:"size(255)"`
	Steps string `json:"steps" orm:"type(text)"` //action_name list
}

//任务流对应单步定义
type ActionImpl struct {
	Id     int                    `json:"id" orm:"pk;auto"`
	Name   string                 `json:"name"`
	Desc   string                 `json:"desc"`
	Type   string                 `json:"type"`
	Params map[string]interface{} `json:"params"`
}
