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
