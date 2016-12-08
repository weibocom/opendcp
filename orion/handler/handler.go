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
