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
