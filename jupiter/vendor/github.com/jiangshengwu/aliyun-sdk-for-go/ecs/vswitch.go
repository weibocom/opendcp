package ecs

import "github.com/jiangshengwu/aliyun-sdk-for-go/util"

type VSwitchService interface {
	/**
	 * ZoneId(required): 可用区 Id
	 * CidrBlock(required): 指定VSwitch的网段
	 * VpcId(required): 指定VSwitch所在的 VPC
	 * VSwitchName(optional): VSwitch名称，不填则为空，默认值为空，[2, 128] 英文或中文字符，必须以大小字母或中文开头，可包含数字，”_”或”-”，这个值会展示在控制台。不能以 http:// 和 https:// 开头
	 * Description(optional): VSwitch 描述，不填则为空，默认值为空，[2, 256] 英文或中文字符，不能以 http:// 和 https:// 开头
	 * ClientToken(optional): 用于保证请求的幂等性。由客户端生成该参数值，要保证在不同请求间唯一，最大不值过 64 个 ASCII 字符
	 */
	CreateVSwitch(param map[string]interface{}) (CreateVSwitchResponse, error)

	/**
	 * VSwitchId(required): 需要删除的 VSwitch 的 Id
	 */
	DeleteVSwitch(param map[string]interface{}) (DeleteVSwitchResponse, error)

	/**
	 * VpcId(required): VpcId
	 * VSwitchId(optional): 需要查询的 VSwitch 的 Id
	 * ZoneId(optional): 可用区 Id
	 * PageNumber(optional):
	 * PageSize(optional):
	 */
	DescribeVSwitches(param map[string]interface{}) (DescribeVSwitchesResponse, error)

	/**
	 * VSwitchId(required)
	 * VSwitchName(optional)
	 * Description(optional)
	 */
	ModifyVSwitchAttribute(param map[string]interface{}) (ModifyVSwitchAttributeResponse, error)
}

type VSwitchOperator struct {
	Common *util.CommonParam
}

type CreateVSwitchResponse struct {
	util.ErrorResponse
	VSwitchId string `json:"VSwitchId"`
}

type DeleteVSwitchResponse struct {
	util.ErrorResponse
}

type DescribeVSwitchesResponse struct {
	util.ErrorResponse
	util.PageResponse
	AllVSwitches VSwitchTypes `json:"VSwitches"`
}

type ModifyVSwitchAttributeResponse struct {
	util.ErrorResponse
}

type VSwitchTypes struct {
	AllVSwitch []VSwitchType `json:"VSwitch"`
}

type VSwitchType struct {
	VSwitchId               string `json:"VSwitchId"`
	VpcId                   string `json:"VpcId"`
	Status                  string `json:"Status"`
	CidrBlock               string `json:"CidrBlock"`
	ZoneId                  string `json:"ZoneId"`
	AvailableIpAddressCount int    `json:"AvailableIpAddressCount"`
	Description             string `json:"Description"`
	VSwitchName             string `json:"VSwitchName"`
	CreationTime            string `json:"CreationTime"`
}

func (op *VSwitchOperator) CreateVSwitch(params map[string]interface{}) (CreateVSwitchResponse, error) {
	var resp CreateVSwitchResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *VSwitchOperator) DeleteVSwitch(params map[string]interface{}) (DeleteVSwitchResponse, error) {
	var resp DeleteVSwitchResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *VSwitchOperator) DescribeVSwitches(params map[string]interface{}) (DescribeVSwitchesResponse, error) {
	var resp DescribeVSwitchesResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *VSwitchOperator) ModifyVSwitchAttribute(params map[string]interface{}) (ModifyVSwitchAttributeResponse, error) {
	var resp ModifyVSwitchAttributeResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}
