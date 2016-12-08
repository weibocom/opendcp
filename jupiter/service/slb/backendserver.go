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
)

func AddBackendServers(loadBalancerId string, backendServers string) (slb.AddBackendServersResponse, error) {
	cli := GetSlbClient()
	params := make(map[string]interface{})
	params["LoadBalancerId"] = loadBalancerId
	params["BackendServers"] = backendServers
	return cli.BackendServer.AddBackendServers(params)
}

func RemoveBackendServers(loadBalancerId string, backendServers string) (slb.RemoveBackendServersResponse, error) {
	cli := GetSlbClient()
	params := make(map[string]interface{})
	params["LoadBalancerId"] = loadBalancerId
	params["BackendServers"] = backendServers
	return cli.BackendServer.RemoveBackendServers(params)
}

func SetBackendServers(loadBalancerId string, backendServers string) (slb.SetBackendServersResponse, error) {
	cli := GetSlbClient()
	params := make(map[string]interface{})
	params["LoadBalancerId"] = loadBalancerId
	params["BackendServers"] = backendServers
	return cli.BackendServer.SetBackendServers(params)
}

func DescribeHealthStatus(loadBalancerId string, listenerPort ...int) (slb.DescribeHealthStatusResponse, error) {
	cli := GetSlbClient()
	params := make(map[string]interface{})
	params["LoadBalancerId"] = loadBalancerId
	if len(listenerPort) == 1 {
		params["ListenerPort"] = listenerPort[0]
	}
	return cli.BackendServer.DescribeHealthStatus(params)
}
