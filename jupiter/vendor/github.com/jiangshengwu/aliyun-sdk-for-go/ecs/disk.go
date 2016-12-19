package ecs

import "github.com/jiangshengwu/aliyun-sdk-for-go/util"

type DiskService interface {
	CreateDisk(params map[string]interface{}) (CreateDiskResponse, error)
	DescribeDisks(params map[string]interface{}) (DescribeDisksResponse, error)
	AttachDisk(params map[string]interface{}) (AttachDiskResponse, error)
	DetachDisk(params map[string]interface{}) (DetachDiskResponse, error)
	ModifyDiskAttribute(params map[string]interface{}) (ModifyDiskAttributeResponse, error)
	DeleteDisk(params map[string]interface{}) (DeleteDiskResponse, error)
	ReInitDisk(params map[string]interface{}) (ReInitDiskResponse, error)
	ResetDisk(params map[string]interface{}) (ResetDiskResponse, error)
	ReplaceSystemDisk(params map[string]interface{}) (ReplaceSystemDiskResponse, error)
	ResizeDisk(params map[string]interface{}) (ResizeDiskResponse, error)
}

type DiskOperator struct {
	Common *util.CommonParam
}

// Response struct for CreateDisk
type CreateDiskResponse struct {
	util.ErrorResponse
	DiskId string `json:"DiskId"`
}

// Response struct for DescribeDisks
type DescribeDisksResponse struct {
	util.PageResponse
	util.ErrorResponse
	AllDisks Disks `json:"Disks"`
}

type Disks struct {
	AllDisk []DiskItemType `json:"Disk"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&diskitemtype
type DiskItemType struct {
	DiskId             string         `json:"DiskId"`
	RegionId           string         `json:"RegionId"`
	ZoneId             string         `json:"ZoneId"`
	DiskName           string         `json:"DiskName"`
	Description        string         `json:"Description"`
	Type               string         `json:"Type"`
	Category           string         `json:"Category"`
	Size               int            `json:"Size"`
	ImageId            string         `json:"ImageId"`
	SourceSnapshotId   string         `json:"SourceSnapshotId"`
	ProductCode        string         `json:"ProductCode"`
	Portable           bool           `json:"Portable"`
	Status             string         `json:"Status"`
	AllOperationLocks  OperationLocks `json:"OperationLocks"`
	InstanceId         string         `json:"InstanceId"`
	Device             string         `json:"Device"`
	DeleteWithInstance bool           `json:"DeleteWithInstance"`
	DeleteAutoSnapshot bool           `json:"DeleteAutoSnapshot"`
	EnableAutoSnapshot bool           `json:"EnableAutoSnapshot"`
	CreationTime       string         `json:"CreationTime"`
	AttachedTime       string         `json:"AttachedTime"`
	DetachedTime       string         `json:"DetachedTime"`
}

// Response struct for AttachDisk
type AttachDiskResponse struct {
	util.ErrorResponse
}

// Response struct for DetachDisk
type DetachDiskResponse struct {
	util.ErrorResponse
}

// Response struct for ModifyDiskAttribute
type ModifyDiskAttributeResponse struct {
	util.ErrorResponse
}

// Response struct for DeleteDisk
type DeleteDiskResponse struct {
	util.ErrorResponse
}

// Response struct for ReInitDisk
type ReInitDiskResponse struct {
	util.ErrorResponse
}

// Response struct for ResetDisk
type ResetDiskResponse struct {
	util.ErrorResponse
}

// Response struct for ReplaceSystemDisk
type ReplaceSystemDiskResponse struct {
	util.ErrorResponse
	DiskId string `json:"DiskId"`
}

// Response struct for ResizeDisk
type ResizeDiskResponse struct {
	util.ErrorResponse
}

func (op *DiskOperator) CreateDisk(params map[string]interface{}) (CreateDiskResponse, error) {
	var resp CreateDiskResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *DiskOperator) DescribeDisks(params map[string]interface{}) (DescribeDisksResponse, error) {
	var resp DescribeDisksResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *DiskOperator) AttachDisk(params map[string]interface{}) (AttachDiskResponse, error) {
	var resp AttachDiskResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *DiskOperator) DetachDisk(params map[string]interface{}) (DetachDiskResponse, error) {
	var resp DetachDiskResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *DiskOperator) ModifyDiskAttribute(params map[string]interface{}) (ModifyDiskAttributeResponse, error) {
	var resp ModifyDiskAttributeResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *DiskOperator) DeleteDisk(params map[string]interface{}) (DeleteDiskResponse, error) {
	var resp DeleteDiskResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *DiskOperator) ReInitDisk(params map[string]interface{}) (ReInitDiskResponse, error) {
	var resp ReInitDiskResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *DiskOperator) ResetDisk(params map[string]interface{}) (ResetDiskResponse, error) {
	var resp ResetDiskResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *DiskOperator) ReplaceSystemDisk(params map[string]interface{}) (ReplaceSystemDiskResponse, error) {
	var resp ReplaceSystemDiskResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}

func (op *DiskOperator) ResizeDisk(params map[string]interface{}) (ResizeDiskResponse, error) {
	var resp ResizeDiskResponse
	err := op.Common.Request(util.GetFuncName(1), params, &resp)
	return resp, err
}
