/**
 *    Copyright (C) 2016 Weibo Inc.
 *
 *    This file is part of Opendcp.
 *
 *    Opendcp is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation; version 2 of the License.
 *
 *    Opendcp is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with Opendcp; if not, write to the Free Software
 *    Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301  USA
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
