package ecs

import "github.com/jiangshengwu/aliyun-sdk-for-go/util"

type SecurityGroupService interface {
	CreateSecurityGroup(params map[string]interface{}) (CreateSecurityGroupResponse, error)
	AuthorizeSecurityGroup(params map[string]interface{}) (AuthorizeSecurityGroupResponse, error)
	DescribeSecurityGroupAttribute(params map[string]interface{}) (DescribeSecurityGroupAttributeResponse, error)
	DescribeSecurityGroups(params map[string]interface{}) (DescribeSecurityGroupsResponse, error)
	RevokeSecurityGroup(params map[string]interface{}) (RevokeSecurityGroupResponse, error)
	DeleteSecurityGroup(params map[string]interface{}) (DeleteSecurityGroupResponse, error)
	ModifySecurityGroupAttribute(params map[string]interface{}) (ModifySecurityGroupAttributeResponse, error)
	AuthorizeSecurityGroupEgress(params map[string]interface{}) (AuthorizeSecurityGroupEgressResponse, error)
	RevokeSecurityGroupEgress(params map[string]interface{}) (RevokeSecurityGroupEgressResponse, error)
}

type SecurityGroupOperator struct {
	Common *util.CommonParam
}

// Response struct for CreateSecurityGroup
type CreateSecurityGroupResponse struct {
	util.ErrorResponse
	SecurityGroupId string `json:"SecurityGroupId"`
}

// Response struct for AuthorizeSecurityGroup
type AuthorizeSecurityGroupResponse struct {
	util.ErrorResponse
}

// Response struct for DescribeSecurityGroupAttribute
type DescribeSecurityGroupAttributeResponse struct {
	util.ErrorResponse
	RegionId        string      `json:"RegionId"`
	SecurityGroupId string      `json:"SecurityGroupId"`
	Description     string      `json:"Description"`
	AllPermissions  Permissions `json:"Permissions"`
	VpcId           string      `json:"VpcId"`
}

type Permissions struct {
	AllPermission []PermissionType `json:"Permission"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&permissiontype
type PermissionType struct {
	IpProtocol              string `json:"IpProtocol"`
	PortRange               string `json:"PortRange"`
	SourceCidrIp            string `json:"SourceCidrIp"`
	SourceGroupId           string `json:"SourceGroupId"`
	SourceGroupOwnerAccount string `json:"SourceGroupOwnerAccount"`
	DestCidrIp              string `json:"DestCidrIp"`
	DestGroupId             string `json:"DestGroupId"`
	DestGroupOwnerAccount   string `json:"DestGroupOwnerAccount"`
	Policy                  string `json:"Policy"`
	NicType                 string `json:"NicType"`
	Priority                int    `json:"Priority"`
}

// Response struct for DescribeSecurityGroupsResponse
type DescribeSecurityGroupsResponse struct {
	util.ErrorResponse
	util.PageResponse
	RegionId  string `json:"RegionId"`
	AllGroups Groups `json:"SecurityGroups"`
}

type Groups struct {
	AllGroup []SecurityGroupItemType `json:"SecurityGroup"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&securitygroupitemtype
type SecurityGroupItemType struct {
	SecurityGroupId   string `json:"SecurityGroupId"`
	SecurityGroupName string `json:"SecurityGroupName"`
	Description       string `json:"Description"`
	VpcId             string `json:"VpcId"`
	CreationTime      string `json:"CreationTime"`
}

// Response struct for RevokeSecurity
type RevokeSecurityGroupResponse struct {
	util.ErrorResponse
}

// Response struct for DeleteSecurityGroup
type DeleteSecurityGroupResponse struct {
	util.ErrorResponse
}

// Response struct for ModifySecurityGroupAttribute
type ModifySecurityGroupAttributeResponse struct {
	util.ErrorResponse
}

// Response struct for AuthorizeSecurityGroupEgress
type AuthorizeSecurityGroupEgressResponse struct {
	util.ErrorResponse
}

// Response struct for RevokeSecurityGroupEgress
type RevokeSecurityGroupEgressResponse struct {
	util.ErrorResponse
}

func (op *SecurityGroupOperator) CreateSecurityGroup(params map[string]interface{}) (CreateSecurityGroupResponse, error) {
	var resp CreateSecurityGroupResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *SecurityGroupOperator) AuthorizeSecurityGroup(params map[string]interface{}) (AuthorizeSecurityGroupResponse, error) {
	var resp AuthorizeSecurityGroupResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *SecurityGroupOperator) DescribeSecurityGroupAttribute(params map[string]interface{}) (DescribeSecurityGroupAttributeResponse, error) {
	var resp DescribeSecurityGroupAttributeResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *SecurityGroupOperator) DescribeSecurityGroups(params map[string]interface{}) (DescribeSecurityGroupsResponse, error) {
	var resp DescribeSecurityGroupsResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *SecurityGroupOperator) RevokeSecurityGroup(params map[string]interface{}) (RevokeSecurityGroupResponse, error) {
	var resp RevokeSecurityGroupResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *SecurityGroupOperator) DeleteSecurityGroup(params map[string]interface{}) (DeleteSecurityGroupResponse, error) {
	var resp DeleteSecurityGroupResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *SecurityGroupOperator) ModifySecurityGroupAttribute(params map[string]interface{}) (ModifySecurityGroupAttributeResponse, error) {
	var resp ModifySecurityGroupAttributeResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *SecurityGroupOperator) AuthorizeSecurityGroupEgress(params map[string]interface{}) (AuthorizeSecurityGroupEgressResponse, error) {
	var resp AuthorizeSecurityGroupEgressResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *SecurityGroupOperator) RevokeSecurityGroupEgress(params map[string]interface{}) (RevokeSecurityGroupEgressResponse, error) {
	var resp RevokeSecurityGroupEgressResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}
