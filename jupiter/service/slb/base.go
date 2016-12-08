// Copyright 2016 Weibo Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package slb

import (
	"github.com/jiangshengwu/aliyun-sdk-for-go/slb"
	"sync"
	"weibo.com/opendcp/jupiter/conf"
)

var globalSlbClient *slb.SlbClient
var once sync.Once

// GetOrmer :set ormer singleton
func GetSlbClient() *slb.SlbClient {
	once.Do(func() {
		globalSlbClient = slb.NewClient(
			conf.Config.KeyId,
			conf.Config.KeySecret,
			"",
		)
	})
	return globalSlbClient
}
