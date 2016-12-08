package slb

import "github.com/jiangshengwu/aliyun-sdk-for-go/util"

type LoadBalancerService interface {
	CreateLoadBalancer(params map[string]interface{}) (CreateLoadBalancerResponse, error)
	ModifyLoadBalancerInternetSpec(params map[string]interface{}) (ModifyLoadBalancerInternetSpecResponse, error)
	DeleteLoadBalancer(params map[string]interface{}) (DeleteLoadBalancerResponse, error)
	SetLoadBalancerStatus(params map[string]interface{}) (SetLoadBalancerStatusResponse, error)
	SetLoadBalancerName(params map[string]interface{}) (SetLoadBalancerNameResponse, error)
	DescribeLoadBalancers(params map[string]interface{}) (DescribeLoadBalancersResponse, error)
	DescribeLoadBalancerAttribute(params map[string]interface{}) (DescribeLoadBalancerAttributeResponse, error)
	DescribeRegions(params map[string]interface{}) (DescribeRegionsResponse, error)
}

type LoadBalancerOperator struct {
	Common *util.CommonParam
}

// Response struct for CreateLoadBalancer
type CreateLoadBalancerResponse struct {
	util.ErrorResponse
	LoadBalancerId   string `json:"LoadBalancerId"`
	Address          string `json:"Address"`
	NetworkType      string `json:"NetworkType"`
	VpcId            string `json:"VpcId"`
	VSwitchId        string `json:"VSwitchId"`
	LoadBalancerName string `json:"LoadBalancerName"`
}

// Response struct for ModifyLoadBalancerInternetSpec
type ModifyLoadBalancerInternetSpecResponse struct {
	util.ErrorResponse
}

// Response struct for DeleteLoadBalancer
type DeleteLoadBalancerResponse struct {
	util.ErrorResponse
}

// Response struct for SetLoadBalancerStatus
type SetLoadBalancerStatusResponse struct {
	util.ErrorResponse
}

// Response struct for SetLoadBalancerName
type SetLoadBalancerNameResponse struct {
	util.ErrorResponse
}

// Response struct for DescribeLoadBalancers
type DescribeLoadBalancersResponse struct {
	util.ErrorResponse
	LoadBalancers LoadBalancerList `json:"LoadBalancers"`
}

type LoadBalancerList struct {
	LoadBalancer []LoadBalancerItemType `json:"LoadBalancer"`
}

type LoadBalancerItemType struct {
	LoadBalancerId     string `json:"LoadBalancerId"`
	LoadBalancerName   string `json:"LoadBalancerName"`
	LoadBalancerStatus string `json:"LoadBalancerStatus"`
	Address            string `json:"Address"`
	RegionId           string `json:"RegionId"`
	RegionIdAlias      string `json:"RegionIdAlias"`
	AddressType        string `json:"AddressType"`
	VSwitchId          string `json:"VSwitchId"`
	VpcId              string `json:"VpcId"`
	NetworkType        string `json:"NetworkType"`
	Bandwidth          int    `json:"Bandwidth"`
	InternetChargeType int    `json:"InternetChargeType"`
	CreateTime         string `json:"CreateTime"`
}

// Response struct for DescribeLoadBalancerAttribute
type DescribeLoadBalancerAttributeResponse struct {
	util.ErrorResponse
	LoadBalancerId           string                      `json:"LoadBalancerId"`
	RegionId                 string                      `json:"RegionId"`
	RegionIdAlias            string                      `json:"RegionIdAlias"`
	LoadBalancerName         string                      `json:"LoadBalancerName"`
	LoadBalancerStatus       string                      `json:"LoadBalancerStatus"`
	Address                  string                      `json:"Address"`
	AddressType              string                      `json:"AddressType"`
	NetworkType              string                      `json:"NetworkType"`
	VpcId                    string                      `json:"VpcId"`
	VswitchId                string                      `json:"VswitchId"`
	InternetChargeType       string                      `json:"InternetChargeType"`
	Bandwidth                int                         `json:"Bandwidth"`
	CreateTime               string                      `json:"CreateTime"`
	ListenerPorts            ListenerPortList            `json:"ListenerPorts"`
	ListenerPortsAndProtocal ListenerPortAndProtocalList `json:"ListenerPortsAndProtocal"`
	ListenerPortsAndProtocol ListenerPortAndProtocolList `json:"ListenerPortsAndProtocol"`
	BackendServers           BackendServerList           `json:"BackendServers"`
}
type ListenerPortList struct {
	ListenerPort []int `json:"ListenerPort"`
}

type ListenerPortAndProtocalList struct {
	ListenerPortAndProtocal []ListenerPortAndProtocalType `json:"ListenerPortAndProtocal"`
}

type ListenerPortAndProtocalType struct {
	ListenerPort     int    `json:"ListenerPort"`
	ListenerProtocal string `json:"ListenerProtocal"`
}

type ListenerPortAndProtocolList struct {
	ListenerPortAndProtocol []ListenerPortAndProtocolType `json:"ListenerPortAndProtocol"`
}

type ListenerPortAndProtocolType struct {
	ListenerPort     int    `json:"ListenerPort"`
	ListenerProtocol string `json:"ListenerProtocol"`
}

// Response struct for DescribeRegions
type DescribeRegionsResponse struct {
	util.ErrorResponse
	Regions RegionList `json:"Regions"`
}

type RegionList struct {
	Region []RegionType `json:"Region"`
}

type RegionType struct {
	RegionId  string `json:"RegionId"`
	LocalName string `json:"LocalName"`
}

func (op *LoadBalancerOperator) CreateLoadBalancer(params map[string]interface{}) (CreateLoadBalancerResponse, error) {
	var resp CreateLoadBalancerResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *LoadBalancerOperator) ModifyLoadBalancerInternetSpec(params map[string]interface{}) (ModifyLoadBalancerInternetSpecResponse, error) {
	var resp ModifyLoadBalancerInternetSpecResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *LoadBalancerOperator) DeleteLoadBalancer(params map[string]interface{}) (DeleteLoadBalancerResponse, error) {
	var resp DeleteLoadBalancerResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *LoadBalancerOperator) SetLoadBalancerStatus(params map[string]interface{}) (SetLoadBalancerStatusResponse, error) {
	var resp SetLoadBalancerStatusResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *LoadBalancerOperator) SetLoadBalancerName(params map[string]interface{}) (SetLoadBalancerNameResponse, error) {
	var resp SetLoadBalancerNameResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *LoadBalancerOperator) DescribeLoadBalancers(params map[string]interface{}) (DescribeLoadBalancersResponse, error) {
	var resp DescribeLoadBalancersResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *LoadBalancerOperator) DescribeLoadBalancerAttribute(params map[string]interface{}) (DescribeLoadBalancerAttributeResponse, error) {
	var resp DescribeLoadBalancerAttributeResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *LoadBalancerOperator) DescribeRegions(params map[string]interface{}) (DescribeRegionsResponse, error) {
	var resp DescribeRegionsResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}
