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
)

const (
	CODE_INIT    = iota
	CODE_RUNNING
	CODE_SUCCESS
	CODE_ERROR
	CODE_PARTIAL
)

// Handler defines how a certain type of Step is processed.
type Handler interface {
	ListAction() []ActionImpl
	Handle(*ActionImpl, map[string]interface{}, []*NodeState, string) *HandleResult
	GetType() string
	GetLog(*NodeState) string
}

// NodeResult represents the result of a node handled by Handler.
type NodeResult struct {
	Code int
	Data string
}

// HandleResult represents all result of nodes handled by Handler.
type HandleResult struct {
	Code   int
	Msg    string
	Result []*NodeResult
}

func Err(msg string) *HandleResult {
	return &HandleResult{
		Code: CODE_ERROR,
		Msg:  msg,
	}
}
