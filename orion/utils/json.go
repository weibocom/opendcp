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

package utils

import (
	"encoding/json"
)

var (
	Json = &Helper{}
)

type Helper struct {
}

func (h *Helper) ToArray(text string) ([]interface{}, error) {
	arr := make([]interface{}, 0)
	err := json.Unmarshal([]byte(text), &arr)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

func (h *Helper) ToMap(text string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(text), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
