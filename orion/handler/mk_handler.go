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

func (v *MockH) ListAction(biz_id int) []ActionImpl {
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

func (v *MockH) GetLog(nodeState *NodeState) string {
	return ""
}

func (v *MockH) HandleInit(*ActionImpl, map[string]interface{}) *HandleResult{
	return nil
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
