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
