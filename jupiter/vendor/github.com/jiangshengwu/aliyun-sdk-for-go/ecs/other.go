package ecs

import "github.com/jiangshengwu/aliyun-sdk-for-go/util"

type OtherService interface {
	DescribeInstanceTypes(params map[string]interface{}) (DescribeInstanceTypesResponse, error)
}

type OtherOperator struct {
	Common *util.CommonParam
}

// Response struct for DescribeInstanceTypes
type DescribeInstanceTypesResponse struct {
	util.ErrorResponse
	AllInstanceTypes InstanceTypes `json:"InstanceTypes"`
}

type InstanceTypes struct {
	AllInstanceType []InstanceTypeItemType `json:"InstanceType"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&instancetypeitemtype
type InstanceTypeItemType struct {
	InstanceTypeFamily string  `json:"InstanceTypeFamily"`
	InstanceTypeId     string  `json:"InstanceTypeId"`
	CpuCoreCount       int     `json:"CpuCoreCount"`
	MemorySize         float64 `json:"MemorySize"`
}

func (op *OtherOperator) DescribeInstanceTypes(params map[string]interface{}) (DescribeInstanceTypesResponse, error) {
	var resp DescribeInstanceTypesResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}
