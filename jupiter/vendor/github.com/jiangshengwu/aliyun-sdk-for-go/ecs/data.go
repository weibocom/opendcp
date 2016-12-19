package ecs

type OperationLocks struct {
	AllOperationLock []OperationLocksType `json:"OperationLock"`
}

// See http://docs.aliyun.com/?spm=5176.775974174.2.4.BYfRJ2#/ecs/open-api/datatype&operationlockstype
type OperationLocksType struct {
	LockReason string `json:"LockReason"`
}
