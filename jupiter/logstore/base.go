package logstore

import (
	"github.com/astaxie/beego/orm"
	"sync"
)

const (
	LOG_TABLE = "log"
)

type LogInfoRequest struct {
	InstanceId string `json:"instance_id"`
	CorrelationId     string `json:"correlationId"`
	Ip         string `json:"ip"`
	Msg        string `json:"message;type(text)"`
	Level      int    `json:"level"`
}

type LogInfo struct {
	InstanceId string      `json:"instance_id" orm:"pk"`
	CorrelationId     string      `json:"correlation_id"`
	Ip         string      `json:"ip"`        //机器ip
	Message    string      `json:"message"`   //日志信息
	Mutex      *sync.Mutex `json:"-" orm:"-"` //互斥锁
}

func NewDefaultLogInfo(instanceId string, correlationId string) (result *LogInfo) {
	result = &LogInfo{}
	result.InstanceId = instanceId
	result.CorrelationId = correlationId
	result.Message = ""
	result.Mutex = new(sync.Mutex)
	return result
}

func NewLogInfo(instanceId string, correlationId string, ip string) (result *LogInfo) {
	result = &LogInfo{}
	result.InstanceId = instanceId
	result.CorrelationId = correlationId
	result.Ip = ip
	result.Message = ""
	result.Mutex = new(sync.Mutex)
	return result
}

func (log *LogInfo) TableName() string {
	return LOG_TABLE
}

func init() {
	orm.RegisterModel(new(LogInfo))
}
