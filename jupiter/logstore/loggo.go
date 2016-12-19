package logstore

import (
	"reflect"
	"time"

	"github.com/astaxie/beego"
)

var resourceMap = NewConcurrentMap(1000, reflect.TypeOf(string("")), reflect.TypeOf(&Store{}))
var flushTimeOnce = 10 * time.Second //每次日志flush至DB的时间间隔
var isPrintBeego = true
var isAsync = false //异步定时写入日志,暂时不能解决多机器时间序问题,所以暂时不能开启

func Start(correlationId string, instanceId string) bool {
	defer func() {
		if err := recover(); err != nil {
			beego.Error("[LogStore] start err!", err)
		}
	}()

	if len(correlationId) == 0 || len(instanceId) <= 0 {
		beego.Warn("[LogStore] start error, instanceId= " + instanceId + " correlationId= " + correlationId)
		return false
	}

	key := getMapKey(correlationId, instanceId)
	if resourceMap.Get(key) != nil {
		//beego.Debug("[LogStore] AddLogResource is already exist!", key)
		return true
	}

	logInfo := NewDefaultLogInfo(instanceId, correlationId)
	resourceMap.Put(key, &Store{
		LogInfo: logInfo,
	})

	go func() {
		timer := time.NewTicker(flushTimeOnce)
		defer func() {
			timer.Stop()
			beego.Info("[LogStore] timeticker stop")
		}()

		for {
			select {
			case <-timer.C:
				if len(resourceMap.Get(key).(*Store).LogInfo.Message) == 0 {
					resourceMap.Remove(key)
					return
				}
				resourceMap.Get(key).(*Store).Flush()
			}
		}
	}()

	return true
}

func AppendIp(correlationId string, instanceId string, ip string) {
	if ok := Start(correlationId, instanceId); !ok {
		return
	}

	resourceMap.Get(getMapKey(correlationId, instanceId)).(*Store).LogInfo.Ip = ip
}

func Warn(correlationId string, instanceId string, v ...interface{}) {
	if isPrintBeego {
		beego.Warn(v...)
	}

	if !isAsync {
		logInfo := NewDefaultLogInfo(instanceId, correlationId)
		store := &Store{
			LogInfo: logInfo,
		}
		store.Warn(v...)
		store.Flush()
		return
	}

	if ok := Start(correlationId, instanceId); !ok {
		return
	}

	resourceMap.Get(getMapKey(correlationId, instanceId)).(*Store).Warn(v...)
}

func Info(correlationId string, instanceId string, v ...interface{}) {
	if isPrintBeego {
		beego.Info(v...)
	}

	if !isAsync {
		logInfo := NewDefaultLogInfo(instanceId, correlationId)
		store := &Store{
			LogInfo: logInfo,
		}
		store.Info(v...)
		store.Flush()
		return
	}

	if ok := Start(correlationId, instanceId); !ok {
		return
	}

	resourceMap.Get(getMapKey(correlationId, instanceId)).(*Store).Info(v...)
}

func Debug(correlationId string, instanceId string, v ...interface{}) {
	if isPrintBeego {
		beego.Debug(v...)
	}

	if !isAsync {
		logInfo := NewDefaultLogInfo(instanceId, correlationId)
		store := &Store{
			LogInfo: logInfo,
		}
		store.Debug(v...)
		store.Flush()
		return
	}

	if ok := Start(correlationId, instanceId); !ok {
		return
	}

	resourceMap.Get(getMapKey(correlationId, instanceId)).(*Store).Debug(v...)
}

func Error(correlationId string, instanceId string, v ...interface{}) {
	if isPrintBeego {
		beego.Error(v...)
	}

	if !isAsync {
		logInfo := NewDefaultLogInfo(instanceId, correlationId)
		store := &Store{
			LogInfo: logInfo,
		}
		store.Error(v...)
		store.Flush()
		return
	}

	if ok := Start(correlationId, instanceId); !ok {
		return
	}

	resourceMap.Get(getMapKey(correlationId, instanceId)).(*Store).Error(v...)
}

func End(correlationId string, instanceId string) {
	key := getMapKey(correlationId, instanceId)
	if resourceMap.Get(key) == nil {
		fail(key)
		return
	}

	resourceMap.Get(getMapKey(correlationId, instanceId)).(*Store).End()
	resourceMap.Remove(key)
}

/**
 * 日志直接存储
 */
func StoreLog(level string, instanceId string, correlationId string, innerIp string, isSync bool, v ...interface{}) {
	logInfo := NewLogInfo(instanceId, correlationId, innerIp)
	store := &Store{
		LogInfo: logInfo,
	}
	store.LogMessage(level, v...)
	if isSync {
		store.Flush()
	} else {
		go func() {
			store.Flush()
		}()
	}
}

/**
 * 日志串行存储
 */
func StoreLogSync(level string, instanceId string, correlationId string, innerIp string, v ...interface{}) {
	StoreLog(level, instanceId, correlationId, innerIp, true, v...)
}

/**
 * 日志异步存储
 */
func StoreLogAnsync(level string, instanceId string, correlationId string, innerIp string, v ...interface{}) {
	StoreLog(level, instanceId, correlationId, innerIp, false, v...)
}

func fail(correlationId string) {
	beego.Warn("[LogStore] log fail! correlationId= " + correlationId)
}

func getMapKey(correlationId string, instanceId string) string {
	return correlationId + "_" + instanceId
}
