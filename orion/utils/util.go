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
