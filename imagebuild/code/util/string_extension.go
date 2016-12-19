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
