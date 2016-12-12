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
)

const (
	CODE_INIT = iota
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
