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
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"strings"
)

/**
加载配置
 */
func LoadConfig(configPath string) map[string]string {
	configMap := make(map[string]string, 0)
	configBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Errorf("read config file %s error, error is %s\n", configPath, err)
		return configMap
	}

	configs := strings.Split(string(configBytes), "\n")
	for _, config := range configs {
		separatorIndex := strings.Index(config, "=")
		if separatorIndex < 0 {
			continue
		}

		configKey := strings.TrimSpace(config[0:separatorIndex])
		if separatorIndex+1 == len(config) {
			configMap[configKey] = ""
			continue
		}

		configValue := strings.TrimSpace(config[separatorIndex+1:])
		configMap[configKey] = configValue
	}

	log.Infof("configs: %s", configMap)

	return configMap
}
