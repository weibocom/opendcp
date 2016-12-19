package ecs

import "github.com/jiangshengwu/aliyun-sdk-for-go/util"

type MonitorService interface {
	DescribeInstanceMonitorData(params map[string]interface{}) (DescribeInstanceMonitorDataResponse, error)
	DescribeEipMonitorData(params map[string]interface{}) (DescribeEipMonitorDataResponse, error)
	DescribeDiskMonitorData(params map[string]interface{}) (DescribeDiskMonitorDataResponse, error)
}

type MonitorOperator struct {
	Common *util.CommonParam
}

// Response struct for DescribeInstanceTypes
type DescribeInstanceMonitorDataResponse struct {
	util.ErrorResponse
	AllMonitorData InstanceMonitorData `json:"MonitorData"`
}

type InstanceMonitorData struct {
	InstanceMonitor []InstanceMonitorDataType `json:"InstanceMonitorData"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&instancemonitordatatype
type InstanceMonitorDataType struct {
	InstanceId        string `json:"InstanceId"`
	CPU               int    `json:"CPU"`
	IntranetRX        int    `json:"IntranetRX"`
	IntranetTX        int    `json:"IntranetTX"`
	IntranetBandwidth int    `json:"IntranetBandwidth"`
	InternetRX        int    `json:"InternetRX"`
	InternetTX        int    `json:"InternetTX"`
	InternetBandwidth int    `json:"InternetBandwidth"`
	IOPSRead          int    `json:"IOPSRead"`
	IOPSWrite         int    `json:"IOPSWrite"`
	BPSRead           int    `json:"BPSRead"`
	BPSWrite          int    `json:"BPSWrite"`
	TimeStamp         string `json:"TimeStamp"`
}

// Response struct for DescribeEipMonitorData
type DescribeEipMonitorDataResponse struct {
	util.ErrorResponse
	AllEipMonitorData EipMonitorData `json:"EipMonitorDatas"`
}

type EipMonitorData struct {
	EipMonitor []EipMonitorDataType `json:"EipMonitorData"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&eipmonitordatatype
type EipMonitorDataType struct {
	EipRX        int    `json:"EipRX"`
	EipTX        int    `json:"EipTX"`
	EipFlow      int    `json:"EipFlow"`
	EipBandwidth int    `json:"EipBandwidth"`
	EipPackets   int    `json:"EipPackets"`
	TimeStamp    string `json:"TimeStamp"`
}

// Response struct for DescribeDiskMonitorData
type DescribeDiskMonitorDataResponse struct {
	util.ErrorResponse
	TotalCount     int             `json:"TotalCount"`
	AllMonitorData DiskMonitorData `json:"MonitorData"`
}

type DiskMonitorData struct {
	DiskMonitor []DiskMonitorDataType `json:"DiskMonitorData"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&diskmonitordatatype
type DiskMonitorDataType struct {
	DiskId    string `json:"DiskId"`
	IOPSRead  int    `json:"IOPSRead"`
	IOPSWrite int    `json:"IOPSWrite"`
	IOPSTotal int    `json:"IOPSTotal"`
	BPSRead   int    `json:"BPSRead"`
	BPSWrite  int    `json:"BPSWrite"`
	BPSTotal  int    `json:"BPSTotal"`
	TimeStamp string `json:"TimeStamp"`
}

func (op *MonitorOperator) DescribeInstanceMonitorData(params map[string]interface{}) (DescribeInstanceMonitorDataResponse, error) {
	var resp DescribeInstanceMonitorDataResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *MonitorOperator) DescribeEipMonitorData(params map[string]interface{}) (DescribeEipMonitorDataResponse, error) {
	var resp DescribeEipMonitorDataResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *MonitorOperator) DescribeDiskMonitorData(params map[string]interface{}) (DescribeDiskMonitorDataResponse, error) {
	var resp DescribeDiskMonitorDataResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}
