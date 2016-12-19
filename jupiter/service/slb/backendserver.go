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
