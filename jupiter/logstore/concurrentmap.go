package logstore

import (
	"errors"
	"reflect"
	"sync"
)

type ConcurrentMap struct {
	kType  reflect.Type
	vType  reflect.Type
	oriMap map[interface{}]interface{}
	mutex  *sync.RWMutex
}

func NewConcurrentMap(capcity int, kType reflect.Type, vType reflect.Type) *ConcurrentMap {
	cm := &ConcurrentMap{
		kType:  kType,
		vType:  vType,
		oriMap: make(map[interface{}]interface{}, capcity),
		mutex:  new(sync.RWMutex),
	}

	return cm
}

func (cm *ConcurrentMap) ValueType() reflect.Type {
	return cm.vType
}

func (cm *ConcurrentMap) KeyType() reflect.Type {
	return cm.kType
}

func (cm *ConcurrentMap) Get(k interface{}) interface{} {
	defer cm.mutex.RUnlock()
	cm.mutex.RLock()
	if cm.oriMap == nil {
		return nil
	}
	elem := cm.oriMap[k]
	return elem
}

func (cm *ConcurrentMap) Put(k interface{}, v interface{}) (bool, error) {
	defer cm.mutex.Unlock()
	if ok := cm.isAcceptType(k, v); !ok {
		return false, errors.New("type is not pair")
	}

	cm.mutex.Lock()
	if cm.oriMap == nil {
		return false, errors.New("map capcity is 0")
	}

	cm.oriMap[k] = v
	return true, nil
}

func (cm *ConcurrentMap) Remove(k interface{}) {
	defer cm.mutex.Unlock()
	cm.mutex.Lock()
	if cm.oriMap == nil {
		return
	}
	delete(cm.oriMap, k)
}

func (cm *ConcurrentMap) Size() int {
	return len(cm.oriMap)
}

func (cm *ConcurrentMap) isAcceptType(k interface{}, v interface{}) bool {
	if reflect.TypeOf(k) != cm.kType || reflect.TypeOf(v) != cm.vType {
		return false
	}

	return true
}
