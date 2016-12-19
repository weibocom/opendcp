package slb

import "github.com/jiangshengwu/aliyun-sdk-for-go/util"

type BackendServerService interface {
	AddBackendServers(params map[string]interface{}) (AddBackendServersResponse, error)
	RemoveBackendServers(params map[string]interface{}) (RemoveBackendServersResponse, error)
	SetBackendServers(params map[string]interface{}) (SetBackendServersResponse, error)
	DescribeHealthStatus(params map[string]interface{}) (DescribeHealthStatusResponse, error)
}

type BackendServerOperator struct {
	Common *util.CommonParam
}

type BackendServerList struct {
	BackendServer []BackendServerType `json:"BackendServer"`
}

type BackendServerType struct {
	ServerId           string `json:"ServerId"`
	Weight             int    `json:"Weight"`
	ServerHealthStatus string `json:"ServerHealthStatus"`
}

// Response struct for AddBackendServers
type AddBackendServersResponse struct {
	util.ErrorResponse
	LoadBalancerId string            `json:"LoadBalancerId"`
	BackendServers BackendServerList `json:"BackendServers"`
}

// Response struct for RemoveBackendServers
type RemoveBackendServersResponse struct {
	util.ErrorResponse
	LoadBalancerId string            `json:"LoadBalancerId"`
	BackendServers BackendServerList `json:"BackendServers"`
}

// Response struct for SetBackendServers
type SetBackendServersResponse struct {
	util.ErrorResponse
	LoadBalancerId string            `json:"LoadBalancerId"`
	BackendServers BackendServerList `json:"BackendServers"`
}

// Response struct for DescribeHealthStatus
type DescribeHealthStatusResponse struct {
	util.ErrorResponse
	BackendServers BackendServerList `json:"BackendServers"`
}

func (op *BackendServerOperator) AddBackendServers(params map[string]interface{}) (AddBackendServersResponse, error) {
	var resp AddBackendServersResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *BackendServerOperator) RemoveBackendServers(params map[string]interface{}) (RemoveBackendServersResponse, error) {
	var resp RemoveBackendServersResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *BackendServerOperator) SetBackendServers(params map[string]interface{}) (SetBackendServersResponse, error) {
	var resp SetBackendServersResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *BackendServerOperator) DescribeHealthStatus(params map[string]interface{}) (DescribeHealthStatusResponse, error) {
	var resp DescribeHealthStatusResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}
