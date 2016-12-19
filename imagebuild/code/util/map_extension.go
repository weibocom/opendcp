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
