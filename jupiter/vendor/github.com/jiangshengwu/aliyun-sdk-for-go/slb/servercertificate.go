package slb

import "github.com/jiangshengwu/aliyun-sdk-for-go/util"

type ServerCertificateService interface {
	UploadServerCertificate(params map[string]interface{}) (UploadServerCertificateResponse, error)
	DeleteServerCertificate(params map[string]interface{}) (DeleteServerCertificateResponse, error)
	DescribeServerCertificates(params map[string]interface{}) (DescribeServerCertificatesResponse, error)
	SetServerCertificateName(params map[string]interface{}) (SetServerCertificateNameResponse, error)
}

type ServerCertificateOperator struct {
	Common *util.CommonParam
}

// Response struct for UploadServerCertificate
type UploadServerCertificateResponse struct {
	util.ErrorResponse
	ServerCertificateId   string `json:"ServerCertificateId"`
	ServerCertificateName string `json:"ServerCertificateName"`
	Fingerprint           string `json:"Fingerprint"`
}

// Response struct for DeleteServerCertificate
type DeleteServerCertificateResponse struct {
	util.ErrorResponse
}

// Response struct for DescribeServerCertificates
type DescribeServerCertificatesResponse struct {
	util.ErrorResponse
	ServerCertificates ServerCertificateList `json:"ServerCertificates"`
}

type ServerCertificateList struct {
	ServerCertificate []ServerCertificateType `json:"ServerCertificate"`
}

type ServerCertificateType struct {
	ServerCertificateId   string `json:"ServerCertificateId"`
	ServerCertificateName string `json:"ServerCertificateName"`
	RegionId              string `json:"RegionId"`
	Fingerprint           string `json:"Fingerprint"`
}

// Response struct for SetServerCertificateName
type SetServerCertificateNameResponse struct {
	util.ErrorResponse
}

func (op *ServerCertificateOperator) UploadServerCertificate(params map[string]interface{}) (UploadServerCertificateResponse, error) {
	var resp UploadServerCertificateResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ServerCertificateOperator) DeleteServerCertificate(params map[string]interface{}) (DeleteServerCertificateResponse, error) {
	var resp DeleteServerCertificateResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ServerCertificateOperator) DescribeServerCertificates(params map[string]interface{}) (DescribeServerCertificatesResponse, error) {
	var resp DescribeServerCertificatesResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *ServerCertificateOperator) SetServerCertificateName(params map[string]interface{}) (SetServerCertificateNameResponse, error) {
	var resp SetServerCertificateNameResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}
