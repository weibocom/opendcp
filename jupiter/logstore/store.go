package logstore

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"weibo.com/opendcp/jupiter/dao"
)

type Store struct {
	LogInfo *LogInfo
}

var isStore = true

const (
	WarnLevel  = "[Warn] "
	InfoLevel  = "[Info] "
	DebugLevel = "[Debug] "
	ErrorLevel = "[Error] "
)

func (store *Store) Start(instanceId string, correlation_id string, ip string) {
	store.LogInfo = NewLogInfo(instanceId, correlation_id, ip)
}

func (store *Store) Info(v ...interface{}) {
	store.LogMessage(InfoLevel, v...)
}

func (store *Store) Warn(v ...interface{}) {
	store.LogMessage(WarnLevel, v...)
}

func (store *Store) Debug(v ...interface{}) {
	store.LogMessage(DebugLevel, v...)
}

func (store *Store) Error(v ...interface{}) {
	store.LogMessage(ErrorLevel, v...)
}

func (store *Store) End() (result bool) {
	defer func() {
		if err := recover(); err != nil {
			beego.Error("[Store] endLog error!", err)
		}
	}()

	result = true
	if ok := store.check(); ok {
		defer store.LogInfo.Mutex.Unlock()
		store.LogInfo.Mutex.Lock()

		//删除原有的数据
		store.DeleteLog(store.LogInfo.InstanceId, store.LogInfo.CorrelationId)

		orm := dao.GetOrmer()
		_, err := orm.Insert(store.LogInfo)
		if err != nil {
			result = false
			beego.Error("[Store] insert fail!", err)
		} else {
			store.LogInfo.Message = ""
			beego.Info("[Store] insert success!")
		}
	}

	return result
}

func (store *Store) Flush() (result bool) {
	defer func() {
		if err := recover(); err != nil {
			beego.Error("[Store] flush error!", err)
		}
	}()

	result = false
	if ok := store.check(); ok {
		defer store.LogInfo.Mutex.Unlock()
		store.LogInfo.Mutex.Lock()

		orm := dao.GetOrmer()
		sql := ""
		if len(store.LogInfo.Ip) == 0 {
			sql = fmt.Sprintf("INSERT INTO %s(instance_id, correlation_id, message, ip) values('%s', '%s', '%s', '%s') ON DUPLICATE KEY UPDATE message = CONCAT(message, '%s')",
				LOG_TABLE, store.LogInfo.InstanceId, store.LogInfo.CorrelationId, store.LogInfo.Message, store.LogInfo.Ip, store.LogInfo.Message)
		} else {
			sql = fmt.Sprintf("INSERT INTO %s(instance_id, correlation_id, message, ip) values('%s', '%s', '%s', '%s') ON DUPLICATE KEY UPDATE message = CONCAT(message, '%s'), ip = '%s'",
				LOG_TABLE, store.LogInfo.InstanceId, store.LogInfo.CorrelationId, store.LogInfo.Message, store.LogInfo.Ip, store.LogInfo.Message, store.LogInfo.Ip)

			//_, err = orm.Raw("INSERT INTO ?(instance_id, correlation_id, message, ip) values(?, '?', '?', '?') ON DUPLICATE KEY UPDATE message = CONCAT(message, '?'), ip = '?'",
			//	LOG_TABLE, store.LogInfo.InstanceId, store.LogInfo.CorrelationId, store.LogInfo.InnerIp, store.LogInfo.Message, store.LogInfo.Message, store.LogInfo.InnerIp).Exec()
		}
		_, err := orm.Raw(sql).Exec()
		if err != nil {
			beego.Error("[Store] update fail! sql= "+sql, err)
		} else {
			result = true
			store.LogInfo.Message = ""
			//beego.Info("[Store] update success!")
		}
	}

	return result
}

func (store *Store) QueryLogByCorrelationIdAndInstanceId(instanceId string, correlation_id string) (result *LogInfo) {
	defer func() {
		if err := recover(); err != nil {
			beego.Error("[Store] queryLog error!", err)
		}
	}()

	result = &LogInfo{}
	if len(instanceId) > 0 && len(correlation_id) > 0 {
		orm := dao.GetOrmer()
		err := orm.QueryTable(LOG_TABLE).Filter("instance_id", instanceId).Filter("correlation_id", correlation_id).One(result)
		if err != nil {
			beego.Error("[Store] queryLog error!", err)
		}
	}

	return result
}

func (store *Store) QueryLogByInstanceId(instanceId string) (result *LogInfo) {
	defer func() {
		if err := recover(); err != nil {
			beego.Error("[Store] queryLog error!", err)
		}
	}()

	result = &LogInfo{}
	if len(instanceId) > 0 {
		orm := dao.GetOrmer()
		err := orm.QueryTable(LOG_TABLE).Filter("instance_id", instanceId).One(result)
		if err != nil {
			beego.Error("[Store] queryLog error!", err)
		}
	}

	return result
}

func (store *Store) DeleteLog(instanceId string, correlation_id string) (result bool) {
	defer func() {
		if err := recover(); err != nil {
			beego.Error("[Store] deleteLog error!", err)
		}
	}()

	result = false
	if len(instanceId) > 0 && len(correlation_id) > 0 {
		orm := dao.GetOrmer()
		num, err := orm.QueryTable(LOG_TABLE).Filter("instance_id", instanceId).Filter("correlation_id", correlation_id).Delete()
		if err != nil {
			beego.Warn("[Store] deleteLog error!", err)
		}

		result = num > 0
	}

	return result
}

func (store *Store) check() (result bool) {
	result = true
	if (store.LogInfo == nil || len(store.LogInfo.CorrelationId) == 0 || store.LogInfo.InstanceId == "") || !isStore {
		result = false
	}

	return result
}

func (store *Store) LogMessage(level string, v ...interface{}) (result bool) {
	result = false
	if ok := store.check(); ok {
		defer store.LogInfo.Mutex.Unlock()
		store.LogInfo.Mutex.Lock()

		store.LogInfo.Message += store.genMessage(level, v...)
		result = true
	}

	return result
}

func (store *Store) genMessage(level string, v ...interface{}) string {
	msg := fmt.Sprintf(store.generateFmtStr(len(v)), v...)
	return store.timeNow() + level + msg + "\n"
}

func (store *Store) timeNow() string {
	return time.Now().Format("2006-01-02 15:04:05") + "\t"
}

func (store *Store) generateFmtStr(n int) string {
	return strings.Repeat("%v ", n)
}
