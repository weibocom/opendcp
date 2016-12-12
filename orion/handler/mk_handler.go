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
package handler

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/astaxie/beego"

	. "weibo.com/opendcp/orion/models"
	u "weibo.com/opendcp/orion/utils"
)

type MockH struct {
}

func (v *MockH) ListAction() []ActionImpl {
	return []ActionImpl{
		{
			Name:   "ok",
			Desc:   "ok",
			Type:   "mk",
			Params: map[string]interface{}{"name": "String"},
		},
		{
			Name:   "fail",
			Desc:   "fail",
			Type:   "mk",
			Params: map[string]interface{}{"name": "String"},
		},
		{
			Name:   "half_fail",
			Desc:   "half_fail",
			Type:   "mk",
			Params: map[string]interface{}{"name": "String"},
		},
		{
			Name:   "one_fail",
			Desc:   "one_fail",
			Type:   "mk",
			Params: map[string]interface{}{"name": "String"},
		},
		{
			Name:   "long_ok",
			Desc:   "long_ok",
			Type:   "mk",
			Params: map[string]interface{}{"name": "String"},
		},
		{
			Name:   "may_fail",
			Desc:   "may_fail",
			Type:   "mk",
			Params: map[string]interface{}{"ratio": "Integer"},
		},
	}
}

func (v *MockH) GetType() string {
	return "vm"
}

func (v *MockH) Handle(action *ActionImpl, actionParams map[string]interface{},
	nodes []*NodeState, corrId string) *HandleResult {

	result := make([]*NodeResult, len(nodes))
	switch action.Name {
	case "ok", "long_ok":
		for i, _ := range nodes {
			result[i] = &NodeResult{
				Code: CODE_SUCCESS,
				Data: "DATA" + fmt.Sprint(i),
			}
		}
	case "fail":
		for i, _ := range nodes {
			result[i] = &NodeResult{
				Code: CODE_ERROR,
				Data: "DATA" + fmt.Sprint(i),
			}
		}
	case "half_fail":
		for i, _ := range nodes {
			code := CODE_SUCCESS
			if i%2 == 1 {
				code = CODE_ERROR
			}
			result[i] = &NodeResult{
				Code: code,
				Data: "DATA" + fmt.Sprint(i),
			}
		}
	case "one_fail":
		for i, _ := range nodes {
			code := CODE_SUCCESS
			if i == 0 {
				code = CODE_ERROR
			}
			result[i] = &NodeResult{
				Code: code,
				Data: "DATA" + fmt.Sprint(i),
			}
		}
	case "may_fail":
		r, _ := u.ToInt(actionParams["ratio"])
		for i, _ := range nodes {
			v := rand.Intn(r)
			code := CODE_SUCCESS
			if v == 0 {
				code = CODE_ERROR
			}
			result[i] = &NodeResult{
				Code: code,
				Data: "DATA" + fmt.Sprint(i),
			}
		}
	}

	if action.Name == "long_ok" {
		beego.Debug("Sleeping for 10s")
		time.Sleep(10 * time.Second)
	}

	return &HandleResult{
		Code:   CODE_SUCCESS,
		Msg:    "OK",
		Result: result,
	}
}
