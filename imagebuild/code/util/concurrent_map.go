/*
 *  Copyright 2009-2016 Weibo, Inc.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package util

import (
	"encoding/json"
	"sync"
)

/**
并发map
 */
type ConcurrentMap struct {
	internalMap map[string]interface{}
	lock        sync.RWMutex
}

func MakeConcurrentMap() *ConcurrentMap {
	m := &ConcurrentMap{internalMap: make(map[string]interface{}, 0)}
	return m
}

func (m *ConcurrentMap) Get(key string) interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if value, ok := m.internalMap[key]; !ok {
		return nil
	} else {
		return value
	}
}

func (m *ConcurrentMap) Put(key string, value interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.internalMap[key] = value
}

func (m *ConcurrentMap) Remove(key string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	delete(m.internalMap, key)
}

func (m *ConcurrentMap) ContainsKey(key string) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if _, ok := m.internalMap[key]; !ok {
		return false
	}

	return true
}

func (m *ConcurrentMap) ToPrettyString() string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	bytes, err := json.MarshalIndent(m.internalMap, "", "  ")
	if err != nil {
		return ""
	}

	return string(bytes)
}

func (m *ConcurrentMap) ToString() string {
	m.lock.RLock()
	defer m.lock.RUnlock()

	bytes, err := json.Marshal(m.internalMap)
	if err != nil {
		return ""
	}

	return string(bytes)
}

type Entry struct {
	Key   string
	Value interface{}
}

func (m *ConcurrentMap) Iterator() <-chan Entry {
	ch := make(chan Entry, len(m.internalMap))

	go func() {
		m.lock.RLock()
		defer m.lock.RUnlock()
		for key, value := range m.internalMap {
			ch <- Entry{Key: key, Value: value}
		}

		close(ch)
	}()

	return ch
}
