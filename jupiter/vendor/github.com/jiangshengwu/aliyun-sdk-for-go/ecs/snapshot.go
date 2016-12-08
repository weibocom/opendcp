package ecs

import "github.com/jiangshengwu/aliyun-sdk-for-go/util"

type SnapshotService interface {
	CreateSnapshot(params map[string]interface{}) (CreateSnapshotResponse, error)
	DeleteSnapshot(params map[string]interface{}) (DeleteSnapshotResponse, error)
	DescribeSnapshots(params map[string]interface{}) (DescribeSnapshotsResponse, error)
	ModifyAutoSnapshotPolicy(params map[string]interface{}) (ModifyAutoSnapshotPolicyResponse, error)
	DescribeAutoSnapshotPolicy(params map[string]interface{}) (DescribeAutoSnapshotPolicyResponse, error)
}

type SnapshotOperator struct {
	Common *util.CommonParam
}

// Response struct for CreateSnapshot
type CreateSnapshotResponse struct {
	util.ErrorResponse
	SnapshotId string `json:"SnapshotId"`
}

// Response struct for DeleteSnapshot
type DeleteSnapshotResponse struct {
	util.ErrorResponse
}

// Response struct for DescribeSnapshots
type DescribeSnapshotsResponse struct {
	util.ErrorResponse
	util.PageResponse
	AllSnapshots Snapshots `json:"Snapshots"`
}

type Snapshots struct {
	AllSnapshot []SnapshotType `json:"Snapshot"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&snapshottype
type SnapshotType struct {
	SnapshotId     string `json:"SnapshotId"`
	SnapshotName   string `json:"SnapshotName"`
	Description    string `json:"Description"`
	Progress       string `json:"Progress"`
	SourceDiskId   string `json:"SourceDiskId"`
	SourceDiskSize int    `json:"SourceDiskSize"`
	SourceDiskType string `json:"SourceDiskType"`
	ProductCode    string `json:"ProductCode"`
	CreationTime   string `json:"CreationTime"`
}

// Response struct for ModifyAutoSnapshotPolicy
type ModifyAutoSnapshotPolicyResponse struct {
	util.ErrorResponse
}

// Response struct for DescribeAutoSnapshotPolicy
type DescribeAutoSnapshotPolicyResponse struct {
	util.ErrorResponse
	AutoSnapshotExecutionStatus AutoSnapshotExecutionStatusType `json:"AutoSnapshotExecutionStatus"`
	AutoSnapshotPolicy          AutoSnapshotPolicyType          `json:"AutoSnapshotPolicy"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&autosnapshotexecutionstatustype
type AutoSnapshotExecutionStatusType struct {
	SystemDiskExecutionStatus string `json:"SystemDiskExecutionStatus"`
	DataDiskExecutionStatus   string `json:"DataDiskExecutionStatus"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&autosnapshotpolicytype
type AutoSnapshotPolicyType struct {
	SystemDiskPolicyEnabled           string `json:"SystemDiskPolicyEnabled"`
	SystemDiskPolicyTimePeriod        int    `json:"SystemDiskPolicyTimePeriod"`
	SystemDiskPolicyRetentionDays     int    `json:"SystemDiskPolicyRetentionDays"`
	SystemDiskPolicyRetentionLastWeek string `json:"SystemDiskPolicyRetentionLastWeek"`
	DataDiskPolicyEnabled             string `json:"DataDiskPolicyEnabled"`
	DataDiskPolicyTimePeriod          int    `json:"DataDiskPolicyTimePeriod"`
	DataDiskPolicyRetentionDays       int    `json:"DataDiskPolicyRetentionDays"`
	DataDiskPolicyRetentionLastWeek   string `json:"DataDiskPolicyRetentionLastWeek"`
}

func (op *SnapshotOperator) CreateSnapshot(params map[string]interface{}) (CreateSnapshotResponse, error) {
	var resp CreateSnapshotResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *SnapshotOperator) DeleteSnapshot(params map[string]interface{}) (DeleteSnapshotResponse, error) {
	var resp DeleteSnapshotResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *SnapshotOperator) DescribeSnapshots(params map[string]interface{}) (DescribeSnapshotsResponse, error) {
	var resp DescribeSnapshotsResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *SnapshotOperator) ModifyAutoSnapshotPolicy(params map[string]interface{}) (ModifyAutoSnapshotPolicyResponse, error) {
	var resp ModifyAutoSnapshotPolicyResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *SnapshotOperator) DescribeAutoSnapshotPolicy(params map[string]interface{}) (DescribeAutoSnapshotPolicyResponse, error) {
	var resp DescribeAutoSnapshotPolicyResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}
