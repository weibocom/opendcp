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
	. "weibo.com/opendcp/orion/models"
	//l "github.com/astaxie/beego"
)

var (
	remoteHandler = &RemoteHandler{}
	sdHandler     = &ServiceDiscoveryHandler{}
	vmHandler     = &VMHandler{}
	handlers      = map[string]Handler{
		"remote": remoteHandler,
		"sd":     sdHandler,
		"vm":     vmHandler,
	}
)

func GetHandler(typeName string) Handler {
	return handlers[typeName]
}

func RegisterHandler(typeName string, h Handler) error {
	handlers[typeName] = h
	return nil
}

func GetActionImpl(name string) *ActionImpl {
	for _, h := range handlers {
		// TODO cache action list
		acts := h.ListAction()
		for _, act := range acts {
			if act.Name == name {
				return &act
			}
		}
	}

	return nil
}

func GetAllActionImpl() []ActionImpl {
	all := make([]ActionImpl, 0)
	for _, h := range handlers {
		acts := h.ListAction()
		for _, act := range acts {
			all = append(all, act)
		}
	}

	return all
}
