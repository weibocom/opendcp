package ecs

import "github.com/jiangshengwu/aliyun-sdk-for-go/util"

type VRouterService interface {
	/**
	 * VRouterId(optional): 需要查询的 VRouter 的 Id
	 * RegionId(required): 查询指定地域的 VRouter
	 * PageNumber:
	 * PageSize:
	 */
	DescribeVRouters(param map[string]interface{}) (DescribeVRoutersResponse, error)

	/**
	 * VRouterId:
	 * VRouterName(optional): 修改后的 VRouter 名字，不填则为空，默认值为空，[2, 128] 英文或中文字符，必须以大小字母或中文开头，可包含数字，”_”或”-”，这个值会展示在控制台。不能以 http:// 和 https:// 开头
	 * Description(optional): 修改后的 VRouter 描述，不填则为空，默认值为空，[2, 256] 英文或中文字符，不能以 http:// 和 https:// 开头
	 */
	ModifyVRouterAttribute(params map[string]interface{}) (ModifyVRouterAttributeResponse, error)
}

type VRouterOperator struct {
	Common *util.CommonParam
}

type VRouterSetTypes struct {
	VRouter []VRouterSetType `json:"VRouter"`
}

type VRouterSetType struct {
	VRouterId     string        `json:"VRouterId"`
	RegionId      string        `json:"RegionId"`
	VpcId         string        `json:"VpcId"`
	RouteTableIds RouteTableIds `json:"RouteTableIds"`
	VRouterName   string        `json:"VRouterName"`
	Description   string        `json:"Description"`
	CreationTime  string        `json:"CreationTime"`
}

type RouteTableIds struct {
	RouteTableId []string `json:"RouteTableId"`
}

type DescribeVRoutersResponse struct {
	util.ErrorResponse
	util.PageResponse
	VRouters VRouterSetTypes `json:"VRouters"`
}

type ModifyVRouterAttributeResponse struct {
	util.ErrorResponse
}

func (op *VRouterOperator) DescribeVRouters(params map[string]interface{}) (DescribeVRoutersResponse, error) {
	var resp DescribeVRoutersResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *VRouterOperator) ModifyVRouterAttribute(params map[string]interface{}) (ModifyVRouterAttributeResponse, error) {
	var resp ModifyVRouterAttributeResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}
