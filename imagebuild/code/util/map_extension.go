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
	"reflect"
	"strings"
)

/**
map扩展
 */
func GetOrDefault(kvs map[string]string, key string, defaultV string) string {
	if _, ok := kvs[key]; !ok {
		return defaultV
	}

	return kvs[key]
}

func DefaultValue(kvs map[string]interface{}, key string) interface{} {
	if _, ok := kvs[key]; !ok {
		return ""
	}
	value := kvs[key]
	switch value.(type) {
		case string:
			value = strings.Replace(value.(string), "'", "&#039;", -1)
	}
	return value
}

func DefaultEmptyArray(kvs map[string]interface{}, key string) interface{} {
	if _, ok := kvs[key]; !ok {
		return make([]string, 0)
	}

	if strings.HasPrefix(reflect.TypeOf(kvs[key]).String(), "[]") {
		return kvs[key]
	}

	ret := make([]interface{}, 1)
	ret[0] = kvs[key]
	return ret
}

func Transform(src map[string][]string) map[string]interface{} {
	new := make(map[string]interface{}, 0)
	for key, value := range src {
		if len(value) == 1 {
			new[key] = value[0]
		} else {
			new[key] = value
		}
	}

	return new
}

func ContainsKey(src map[string]interface{}, key string) bool {
	if _, ok := src[key]; !ok {
		return false
	}

	return true
}

func IsMap(param interface{}) bool {
	if param == nil {
		return false
	}

	if strings.HasPrefix(reflect.TypeOf(param).String(), "map[") {
		return true
	}

	return false
}
