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
	"html/template"
	"strings"
	"bytes"
)
/**
string扩展
 */
func IsEmpty(src string) bool {
	if len(strings.TrimSpace(src)) == 0 {
		return true
	}

	return false
}

func StartWith(src string, dst string) bool {
	return strings.HasPrefix(src, dst)
}

func EndWith(src string, dst string) bool {
	return strings.HasSuffix(src, dst)
}

func Unescaped(x string) interface{} {
	return template.HTML(x)
}

func Escaped(x string) interface{} {
	return template.HTMLEscapeString(x)
}

func IsString(src interface{}) bool {
	switch src.(type) {
	case string:
		return true
	}

	return false
}

func ConvertToHump(src string) string {
	strs := strings.Split(src, "_")
	var result bytes.Buffer
	for _,item := range  strs {
		chars := []rune(item)
		head:=strings.ToUpper(string(chars[0]))
		tail:=string(chars[1:])
		result.WriteString(head)
		result.WriteString(tail)
	}
	return result.String()
}
