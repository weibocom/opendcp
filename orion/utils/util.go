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
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func ToInt(v interface{}) (int, error) {
	switch v.(type) {
	case int:
		return v.(int), nil
	case float64:
		return int(v.(float64)), nil
	case string:
		i, err := strconv.Atoi(v.(string))
		if err == nil {
			return i, nil
		}
	}

	return -1, errors.New("cannot convert to int:" + fmt.Sprintln(v) +
		", type : " + reflect.TypeOf(v).String())
}
