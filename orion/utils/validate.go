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
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
	"unicode"

	log "github.com/astaxie/beego"
)

const tempplate_path string = "orion/template.json"

type ValidateUtil struct {
	template map[string]interface{}
}

var validate *ValidateUtil
var onceForValidate sync.Once

func GetValidateUtil() *ValidateUtil {
	if validate == nil {
		onceForValidate.Do(func() {
			if validate == nil {
				path, err := filepath.Abs(tempplate_path)
				if err != nil {
					log.Debug("filepath abs with err:", err)
					panic(err)
				}
				fmt.Println("path:", path)
				bytes, err := ioutil.ReadFile(path)
				if err != nil {
					log.Debug("ioutil ReadFile with err:", err)
					panic(err)
				}
				template := make(map[string]interface{})
				err = json.Unmarshal(bytes, &template)
				if err != nil {
					log.Debug("json Unmarshal with err:", err)
					panic(err)
				}
				validate = &ValidateUtil{template: template}
			}
		})
	}
	return validate
}

/**
判断jsonString是否符合模板样式
*/
func (validate *ValidateUtil) ValidateString(jsonString string) bool {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &jsonMap)
	if err != nil {
		return false
	}
	return validate.ValidateMap(jsonMap)
}

/**
判断jsonMap是否符合模板样式
*/
func (validate *ValidateUtil) ValidateMap(jsonMap map[string]interface{}) bool {
	result := false
	for k, v := range validate.template {
		if jsonMap[k] == nil {
			break
		}
		for _, item := range v.([]interface{}) {
			action := item.(map[string]interface{})
			if action["module"] == jsonMap[k].(map[string]interface{})["module"] && jsonMap[k].(map[string]interface{})[action["diff"].(string)] != "" {
				result = true
				break
			}
		}
	}
	return result
}

/**
是否包含中文
*/
func (validate *ValidateUtil) IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}
