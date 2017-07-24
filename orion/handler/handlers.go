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
