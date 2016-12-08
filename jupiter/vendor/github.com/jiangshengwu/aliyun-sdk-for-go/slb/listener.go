package slb

import "github.com/jiangshengwu/aliyun-sdk-for-go/util"

type ListenerService interface {
	CreateLoadBalancerHTTPListener(params map[string]interface{}) (CreateLoadBalancerHTTPListenerResponse, error)
	CreateLoadBalancerHTTPSListener(params map[string]interface{}) (CreateLoadBalancerHTTPSListenerResponse, error)
	CreateLoadBalancerTCPListener(params map[string]interface{}) (CreateLoadBalancerTCPListenerResponse, error)
	CreateLoadBalancerUDPListener(params map[string]interface{}) (CreateLoadBalancerUDPListenerResponse, error)
	DeleteLoadBalancerListener(params map[string]interface{}) (DeleteLoadBalancerListenerResponse, error)
	StartLoadBalancerListener(params map[string]interface{}) (StartLoadBalancerListenerResponse, error)
	StopLoadBalancerListener(params map[string]interface{}) (StopLoadBalancerListenerResponse, error)
	SetListenerAccessControlStatus(params map[string]interface{}) (SetListenerAccessControlStatusResponse, error)
	AddListenerWhiteListItem(params map[string]interface{}) (AddListenerWhiteListItemResponse, error)
	RemoveListenerWhiteListItem(params map[string]interface{}) (RemoveListenerWhiteListItemResponse, error)
	SetLoadBalancerHTTPListenerAttribute(params map[string]interface{}) (SetLoadBalancerHTTPListenerAttributeResponse, error)
	SetLoadBalancerHTTPSListenerAttribute(params map[string]interface{}) (SetLoadBalancerHTTPSListenerAttributeResponse, error)
	SetLoadBalancerTCPListenerAttribute(params map[string]interface{}) (SetLoadBalancerTCPListenerAttributeResponse, error)
	SetLoadBalancerUDPListenerAttribute(params map[string]interface{}) (SetLoadBalancerUDPListenerAttributeResponse, error)
	DescribeLoadBalancerHTTPListenerAttribute(params map[string]interface{}) (DescribeLoadBalancerHTTPListenerAttributeResponse, error)
	DescribeLoadBalancerHTTPSListenerAttribute(params map[string]interface{}) (DescribeLoadBalancerHTTPSListenerAttributeResponse, error)
	DescribeLoadBalancerTCPListenerAttribute(params map[string]interface{}) (DescribeLoadBalancerTCPListenerAttributeResponse, error)
	DescribeLoadBalancerUDPListenerAttribute(params map[string]interface{}) (DescribeLoadBalancerUDPListenerAttributeResponse, error)
	DescribeListenerAccessControlAttribute(params map[string]interface{}) (DescribeListenerAccessControlAttributeResponse, error)
}

type ListenerOperator struct {
	Common *util.CommonParam
}

// Response struct for CreateLoadBalancerHTTPListener
type CreateLoadBalancerHTTPListenerResponse struct {
	util.ErrorResponse
}

// Response struct for CreateLoadBalancerHTTPSListener
type CreateLoadBalancerHTTPSListenerResponse struct {
	util.ErrorResponse
}

// Response struct for CreateLoadBalancerTCPListener
type CreateLoadBalancerTCPListenerResponse struct {
	util.ErrorResponse
}

// Response struct for CreateLoadBalancerUDPListener
type CreateLoadBalancerUDPListenerResponse struct {
	util.ErrorResponse
}

// Response struct for DeleteLoadBalancerListener
type DeleteLoadBalancerListenerResponse struct {
	util.ErrorResponse
}

// Response struct for StartLoadBalancerListener
type StartLoadBalancerListenerResponse struct {
	util.ErrorResponse
}

// Response struct for StopLoadBalancerListener
type StopLoadBalancerListenerResponse struct {
	util.ErrorResponse
}

// Response struct for SetListenerAccessControlStatus
type SetListenerAccessControlStatusResponse struct {
	util.ErrorResponse
}

// Response struct for AddListenerWhiteListItem
type AddListenerWhiteListItemResponse struct {
	util.ErrorResponse
}

// Response struct for RemoveListenerWhiteListItem
type RemoveListenerWhiteListItemResponse struct {
	util.ErrorResponse
}

// Response struct for SetLoadBalancerHTTPListenerAttribute
type SetLoadBalancerHTTPListenerAttributeResponse struct {
	util.ErrorResponse
}

// Response struct for SetLoadBalancerHTTPSListenerAttribute
type SetLoadBalancerHTTPSListenerAttributeResponse struct {
	util.ErrorResponse
}

// Response struct for SetLoadBalancerTCPListenerAttribute
type SetLoadBalancerTCPListenerAttributeResponse struct {
	util.ErrorResponse
}

// Response struct for SetLoadBalancerUDPListenerAttribute
type SetLoadBalancerUDPListenerAttributeResponse struct {
	util.ErrorResponse
}

// Response struct for DescribeLoadBalancerHTTPListenerAttribute
type DescribeLoadBalancerHTTPListenerAttributeResponse struct {
	util.ErrorResponse
	ListenerPort           int    `json:"ListenerPort"`
	BackendServerPort      int    `json:"BackendServerPort"`
	Bandwidth              int    `json:"Bandwidth"`
	Status                 string `json:"Status"`
	XForwardedFor          string `json:"XForwardedFor"`
	Scheduler              string `json:"Scheduler"`
	StickySession          string `json:"StickySession"`
	StickySessionType      string `json:"StickySessionType"`
	CookieTimeout          int    `json:"CookieTimeout"`
	Cookie                 string `json:"Cookie"`
	HealthCheck            string `json:"HealthCheck"`
	HealthCheckDomain      string `json:"HealthCheckDomain"`
	HealthCheckURI         string `json:"HealthCheckURI"`
	HealthyThreshold       int    `json:"HealthyThreshold"`
	UnhealthyThreshold     int    `json:"UnhealthyThreshold"`
	HealthCheckTimeout     int    `json:"HealthCheckTimeout"`
	HealthCheckInterval    int    `json:"HealthCheckInterval"`
	HealthCheckHttpCode    string `json:"HealthCheckHttpCode"`
	HealthCheckConnectPort int    `json:"HealthCheckConnectPort"`
}

// Response struct for DescribeLoadBalancerHTTPSListenerAttribute
type DescribeLoadBalancerHTTPSListenerAttributeResponse struct {
	util.ErrorResponse
	ListenerPort           int    `json:"ListenerPort"`
	BackendServerPort      int    `json:"BackendServerPort"`
	Bandwidth              int    `json:"Bandwidth"`
	Status                 string `json:"Status"`
	XForwardedFor          string `json:"XForwardedFor"`
	Scheduler              string `json:"Scheduler"`
	StickySession          string `json:"StickySession"`
	StickySessionType      string `json:"StickySessionType"`
	CookieTimeout          int    `json:"CookieTimeout"`
	Cookie                 string `json:"Cookie"`
	HealthCheck            string `json:"HealthCheck"`
	HealthCheckDomain      string `json:"HealthCheckDomain"`
	HealthCheckURI         string `json:"HealthCheckURI"`
	HealthyThreshold       int    `json:"HealthyThreshold"`
	UnhealthyThreshold     int    `json:"UnhealthyThreshold"`
	HealthCheckTimeout     int    `json:"HealthCheckTimeout"`
	HealthCheckInterval    int    `json:"HealthCheckInterval"`
	HealthCheckHttpCode    string `json:"HealthCheckHttpCode"`
	HealthCheckConnectPort int    `json:"HealthCheckConnectPort"`
	ServerCertificateId    string `json:"ServerCertificateId"`
}

// Response struct for DescribeLoadBalancerTCPListenerAttribute
type DescribeLoadBalancerTCPListenerAttributeResponse struct {
	util.ErrorResponse
	ListenerPort              int    `json:"ListenerPort"`
	BackendServerPort         int    `json:"BackendServerPort"`
	Bandwidth                 int    `json:"Bandwidth"`
	Status                    string `json:"Status"`
	SynProxy                  string `json:"SynProxy"`
	Scheduler                 string `json:"Scheduler"`
	PersistenceTimeout        int    `json:"PersistenceTimeout"`
	HealthCheckType           string `json:"HealthCheckType"`
	HealthCheck               string `json:"HealthCheck"`
	HealthyThreshold          int    `json:"HealthyThreshold"`
	UnhealthyThreshold        int    `json:"UnhealthyThreshold"`
	HealthCheckConnectTimeout int    `json:"HealthCheckConnectTimeout"`
	HealthCheckConnectPort    int    `json:"HealthCheckConnectPort"`
	HealthCheckInterval       int    `json:"HealthCheckInterval"`
	HealthCheckDomain         string `json:"HealthCheckDomain"`
	HealthCheckURI            string `json:"HealthCheckURI"`
	HealthCheckHttpCode       string `json:"HealthCheckHttpCode"`
}

// Response struct for DescribeLoadBalancerUDPListenerAttribute
type DescribeLoadBalancerUDPListenerAttributeResponse struct {
	util.ErrorResponse
	ListenerPort              int    `json:"ListenerPort"`
	BackendServerPort         int    `json:"BackendServerPort"`
	Bandwidth                 int    `json:"Bandwidth"`
	Status                    string `json:"Status"`
	Scheduler                 string `json:"Scheduler"`
	PersistenceTimeout        int    `json:"PersistenceTimeout"`
	StickySessionType         string `json:"StickySessionType"`
	HealthCheck               string `json:"HealthCheck"`
	HealthyThreshold          int    `json:"HealthyThreshold"`
	UnhealthyThreshold        int    `json:"UnhealthyThreshold"`
	HealthCheckConnectTimeout int    `json:"HealthCheckConnectTimeout"`
	HealthCheckConnectPort    int    `json:"HealthCheckConnectPort"`
	HealthCheckInterval       int    `json:"HealthCheckInterval"`
}

// Response struct for DescribeListenerAccessControlAttribute
type DescribeListenerAccessControlAttributeResponse struct {
	util.ErrorResponse
	AccessControlStatus string `json:"AccessControlStatus"`
	SourceItems         string `json:"SourceItems"`
}

func (op *ListenerOperator) CreateLoadBalancerHTTPListener(params map[string]interface{}) (CreateLoadBalancerHTTPListenerResponse, error) {
	var resp CreateLoadBalancerHTTPListenerResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ListenerOperator) CreateLoadBalancerHTTPSListener(params map[string]interface{}) (CreateLoadBalancerHTTPSListenerResponse, error) {
	var resp CreateLoadBalancerHTTPSListenerResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ListenerOperator) CreateLoadBalancerTCPListener(params map[string]interface{}) (CreateLoadBalancerTCPListenerResponse, error) {
	var resp CreateLoadBalancerTCPListenerResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ListenerOperator) CreateLoadBalancerUDPListener(params map[string]interface{}) (CreateLoadBalancerUDPListenerResponse, error) {
	var resp CreateLoadBalancerUDPListenerResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ListenerOperator) DeleteLoadBalancerListener(params map[string]interface{}) (DeleteLoadBalancerListenerResponse, error) {
	var resp DeleteLoadBalancerListenerResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ListenerOperator) StartLoadBalancerListener(params map[string]interface{}) (StartLoadBalancerListenerResponse, error) {
	var resp StartLoadBalancerListenerResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ListenerOperator) StopLoadBalancerListener(params map[string]interface{}) (StopLoadBalancerListenerResponse, error) {
	var resp StopLoadBalancerListenerResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ListenerOperator) SetListenerAccessControlStatus(params map[string]interface{}) (SetListenerAccessControlStatusResponse, error) {
	var resp SetListenerAccessControlStatusResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ListenerOperator) AddListenerWhiteListItem(params map[string]interface{}) (AddListenerWhiteListItemResponse, error) {
	var resp AddListenerWhiteListItemResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ListenerOperator) RemoveListenerWhiteListItem(params map[string]interface{}) (RemoveListenerWhiteListItemResponse, error) {
	var resp RemoveListenerWhiteListItemResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ListenerOperator) SetLoadBalancerHTTPListenerAttribute(params map[string]interface{}) (SetLoadBalancerHTTPListenerAttributeResponse, error) {
	var resp SetLoadBalancerHTTPListenerAttributeResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ListenerOperator) SetLoadBalancerHTTPSListenerAttribute(params map[string]interface{}) (SetLoadBalancerHTTPSListenerAttributeResponse, error) {
	var resp SetLoadBalancerHTTPSListenerAttributeResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ListenerOperator) SetLoadBalancerTCPListenerAttribute(params map[string]interface{}) (SetLoadBalancerTCPListenerAttributeResponse, error) {
	var resp SetLoadBalancerTCPListenerAttributeResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ListenerOperator) SetLoadBalancerUDPListenerAttribute(params map[string]interface{}) (SetLoadBalancerUDPListenerAttributeResponse, error) {
	var resp SetLoadBalancerUDPListenerAttributeResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ListenerOperator) DescribeLoadBalancerHTTPListenerAttribute(params map[string]interface{}) (DescribeLoadBalancerHTTPListenerAttributeResponse, error) {
	var resp DescribeLoadBalancerHTTPListenerAttributeResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ListenerOperator) DescribeLoadBalancerHTTPSListenerAttribute(params map[string]interface{}) (DescribeLoadBalancerHTTPSListenerAttributeResponse, error) {
	var resp DescribeLoadBalancerHTTPSListenerAttributeResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ListenerOperator) DescribeLoadBalancerTCPListenerAttribute(params map[string]interface{}) (DescribeLoadBalancerTCPListenerAttributeResponse, error) {
	var resp DescribeLoadBalancerTCPListenerAttributeResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ListenerOperator) DescribeLoadBalancerUDPListenerAttribute(params map[string]interface{}) (DescribeLoadBalancerUDPListenerAttributeResponse, error) {
	var resp DescribeLoadBalancerUDPListenerAttributeResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ListenerOperator) DescribeListenerAccessControlAttribute(params map[string]interface{}) (DescribeListenerAccessControlAttributeResponse, error) {
	var resp DescribeListenerAccessControlAttributeResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}
