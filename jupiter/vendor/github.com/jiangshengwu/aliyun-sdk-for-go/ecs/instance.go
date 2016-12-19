package ecs

import "github.com/jiangshengwu/aliyun-sdk-for-go/util"

//import "github.com/LiuJunlang/aliyun-sdk-for-go/util"

type InstanceService interface {
	/**
	 * RegionId(required):
	 * ZoneId(optional):
	 * ImageId(required):
	 * InstanceType(required): OtherService.DescribeInstanceTypes
	 * SecurityGroupId(required):
	 * InstanceName(optional):
	 * Description(optional):
	 * InternetChargeType(optional): 网络计费类型，按流量计费还是按固定带宽计费: PayByBandwidth|PayByTraffic，默认PayByBandwidth
	 * InternetMaxBandwidthIn(optional):
	 * InternetMaxBandwidthOut(optional):
	 * HostName(optional):
	 * Password(optional):
	 * IoOptimized(optional): IO优化, none|optimized, 默认none
	 * SystemDisk.Category(optional):
	 * SystemDisk.DiskName(optional):
	 * SystemDisk.Description(optional):
	 * DataDisk.n.Size(optional):
	 * DataDisk.n.Category(optional):
	 * DataDisk.n.SnapshotId(optional):
	 * DataDisk.n.DiskName(optional):
	 * DataDisk.n.Description(optional):
	 * DataDisk.n.Device(optional):
	 * DataDisk.n.DeleteWithInstance(optional):
	 * VSwitchId(optional): 如果是创建 VPC 类型的实例，需要指定交换机的 ID
	 * PrivateIpAddress(optional):
	 */
	CreateInstance(params map[string]interface{}) (CreateInstanceResponse, error)
	StartInstance(params map[string]interface{}) (StartInstanceResponse, error)
	StopInstance(params map[string]interface{}) (StopInstanceResponse, error)
	RebootInstance(params map[string]interface{}) (RebootInstanceResponse, error)
	ModifyInstanceAttribute(params map[string]interface{}) (ModifyInstanceAttributeResponse, error)
	ModifyInstanceVpcAttribute(params map[string]interface{}) (ModifyInstanceVpcAttributeResponse, error)
	DescribeInstanceStatus(params map[string]interface{}) (DescribeInstanceStatusResponse, error)
	DescribeInstanceAttribute(params map[string]interface{}) (DescribeInstanceAttributeResponse, error)
	DescribeInstances(params map[string]interface{}) (DescribeInstancesResponse, error)
	DeleteInstance(params map[string]interface{}) (DeleteInstanceResponse, error)
	JoinSecurityGroup(params map[string]interface{}) (JoinSecurityGroupResponse, error)
	LeaveSecurityGroup(params map[string]interface{}) (LeaveSecurityGroupResponse, error)
	DescribeInstanceVncUrl(params map[string]interface{}) (DescribeInstanceVncUrlResponse, error)
	ModifyInstanceVncPasswd(params map[string]interface{}) (ModifyInstanceVncPasswdResponse, error)
}

type InstanceOperator struct {
	Common *util.CommonParam
}

// Response struct for CreateInstance
type CreateInstanceResponse struct {
	util.ErrorResponse
	InstanceId string `json:"InstanceId"`
}

// Response struct for StartInstance
type StartInstanceResponse struct {
	util.ErrorResponse
}

// Response struct for StopInstance
type StopInstanceResponse struct {
	util.ErrorResponse
}

// Response struct for RebootInstance
type RebootInstanceResponse struct {
	util.ErrorResponse
}

// Response struct for ModifyInstanceAttribute
type ModifyInstanceAttributeResponse struct {
	util.ErrorResponse
}

// Response struct for ModifyInstanceVpcAttribute
type ModifyInstanceVpcAttributeResponse struct {
	util.ErrorResponse
}

// Response struct for DescribeInstanceStatus
type DescribeInstanceStatusResponse struct {
	util.ErrorResponse
	util.PageResponse
	AllInstanceStatuses InstanceStatusSetType `json:"InstanceStatuses"`
}

type InstanceStatusSetType struct {
	AllInstanceStatus []InstanceStatusItemType `json:"InstanceStatus"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&instancestatusitemtype
type InstanceStatusItemType struct {
	InstanceId string `json:"InstanceId"`
	Status     string `json:"Status"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&instanceattributestype
type InstanceAttributesType struct {
	InstanceId              string                  `json:"InstanceId"`
	InstanceName            string                  `json:"InstanceName"`
	Description             string                  `json:"Description"`
	ImageId                 string                  `json:"ImageId"`
	RegionId                string                  `json:"RegionId"`
	ZoneId                  string                  `json:"ZoneId"`
	ClusterId               string                  `json:"ClusterId"`
	InstanceType            string                  `json:"InstanceType"`
	HostName                string                  `json:"HostName"`
	Status                  string                  `json:"Status"`
	AllOperationLocks       OperationLocks          `json:"OperationLocks"`
	AllSecurityGroupIds     SecurityGroupIdSetType  `json:"SecurityGroupIds"`
	PublicIpAddress         IpAddressSetType        `json:"PublicIpAddress"`
	InnerIpAddress          IpAddressSetType        `json:"InnerIpAddress"`
	InternetMaxBandwidthIn  int                     `json:"InternetMaxBandwidthIn"`
	InternetMaxBandwidthOut int                     `json:"InternetMaxBandwidthOut"`
	InternetChargeType      string                  `json:"InternetChargeType"`
	CreationTime            string                  `json:"CreationTime"`
	InstanceNetworkType     string                  `json:"InstanceNetworkType"`
	VpcAttributes           VpcAttributesType       `json:"VpcAttributes"`
	EipAddress              EipAddressAssociateType `json:"EipAddress"`
	InstanceChargeType      string                  `json:"InstanceChargeType"` // PrePaid：预付费，即包年包月; PostPaid：后付费，即按量付费
	DeviceAvailable         bool                    `json:"DeviceAvailable"`    // 实例是否还可以挂载磁盘
	IoOptimized             bool                    `json:"IoOptimized"`        // 是否是 IO 优化型实例
	ExpiredTime             string                  `json:"ExpiredTime"`        // 过期时间，按照ISO8601标准表示，并需要使用UTC时间。格式为：YYYY-MM-DDThh:mmZ
}

// hack for the different returned type of IoOptimized
type InstanceAttributesTypeExt struct {
	InstanceAttributesType
	IoOptimized string `json:"IoOptimized"` // 是否是 IO 优化型实例
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&vpcattributestype
type VpcAttributesType struct {
	VpcId            string           `json:"VpcId"`
	VSwitchId        string           `json:"VSwitchId"`
	PrivateIpAddress IpAddressSetType `json:"PrivateIpAddress"`
	NatIpAddress     string           `json:"NatIpAddress"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&eipaddressassociatetype
type EipAddressAssociateType struct {
	AllocationId       string `json:"AllocationId"`
	IpAddress          string `json:"IpAddress"`
	Bandwidth          int    `json:"Bandwidth"`
	InternetChargeType string `json:"InternetChargeType"`
}

// Response struct for DescribeInstanceAttribute
type DescribeInstanceAttributeResponse struct {
	util.ErrorResponse
	InstanceAttributesTypeExt
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&securitygroupidsettype
type SecurityGroupIdSetType struct {
	AllSecurityGroupId []string `json:"SecurityGroupId"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&ipaddresssettype
type IpAddressSetType struct {
	AllIpAddress []string `json:"IpAddress"`
}

// Response struct for DescribeInstances
type DescribeInstancesResponse struct {
	util.ErrorResponse
	util.PageResponse
	AllInstances Instances `json:"Instances"`
}

type Instances struct {
	AllInstance []InstanceAttributesType `json:"Instance"`
}

// Response struct for DeleteInstance
type DeleteInstanceResponse struct {
	util.ErrorResponse
}

// Response struct for JoinSecurityGroup
type JoinSecurityGroupResponse struct {
	util.ErrorResponse
}

// Response struct for LeaveSecurityGroup
type LeaveSecurityGroupResponse struct {
	util.ErrorResponse
}

// Response struct for DescribeInstanceVncUrl
type DescribeInstanceVncUrlResponse struct {
	util.ErrorResponse
	VncUrl string `json:"VncUrl"`
}

// Response struct for ModifyInstanceVncPasswd
type ModifyInstanceVncPasswdResponse struct {
	util.ErrorResponse
}

func (op *InstanceOperator) CreateInstance(params map[string]interface{}) (CreateInstanceResponse, error) {
	var resp CreateInstanceResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *InstanceOperator) StartInstance(params map[string]interface{}) (StartInstanceResponse, error) {
	var resp StartInstanceResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *InstanceOperator) StopInstance(params map[string]interface{}) (StopInstanceResponse, error) {
	var resp StopInstanceResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *InstanceOperator) RebootInstance(params map[string]interface{}) (RebootInstanceResponse, error) {
	var resp RebootInstanceResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *InstanceOperator) ModifyInstanceAttribute(params map[string]interface{}) (ModifyInstanceAttributeResponse, error) {
	var resp ModifyInstanceAttributeResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *InstanceOperator) ModifyInstanceVpcAttribute(params map[string]interface{}) (ModifyInstanceVpcAttributeResponse, error) {
	var resp ModifyInstanceVpcAttributeResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *InstanceOperator) DescribeInstanceStatus(params map[string]interface{}) (DescribeInstanceStatusResponse, error) {
	var resp DescribeInstanceStatusResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *InstanceOperator) DescribeInstanceAttribute(params map[string]interface{}) (DescribeInstanceAttributeResponse, error) {
	var resp DescribeInstanceAttributeResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *InstanceOperator) DescribeInstances(params map[string]interface{}) (DescribeInstancesResponse, error) {
	var resp DescribeInstancesResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *InstanceOperator) DeleteInstance(params map[string]interface{}) (DeleteInstanceResponse, error) {
	var resp DeleteInstanceResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *InstanceOperator) JoinSecurityGroup(params map[string]interface{}) (JoinSecurityGroupResponse, error) {
	var resp JoinSecurityGroupResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *InstanceOperator) LeaveSecurityGroup(params map[string]interface{}) (LeaveSecurityGroupResponse, error) {
	var resp LeaveSecurityGroupResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *InstanceOperator) DescribeInstanceVncUrl(params map[string]interface{}) (DescribeInstanceVncUrlResponse, error) {
	var resp DescribeInstanceVncUrlResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *InstanceOperator) ModifyInstanceVncPasswd(params map[string]interface{}) (ModifyInstanceVncPasswdResponse, error) {
	var resp ModifyInstanceVncPasswdResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}
