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
	"weibo.com/opendcp/jupiter/models"
)

func CreateLoadBalancer(loadBalancer models.LoadBalancer) (slb.CreateLoadBalancerResponse, error) {
	cli := GetSlbClient()
	params := make(map[string]interface{})
	params["RegionId"] = loadBalancer.RegionId
	params["LoadBalancerName"] = loadBalancer.LoadBalancerName
	params["AddressType"] = loadBalancer.AddressType
	params["VSwitchId"] = loadBalancer.VSwitchId
	params["InternetChargeType"] = loadBalancer.InternetChargeType
	params["Bandwidth"] = loadBalancer.Bandwidth

	return cli.LoadBalancer.CreateLoadBalancer(params)
}

func ModifyLoadBalancerInternetSpec(loadBalancerId string, internetChargeType string, bandwidth int) (slb.ModifyLoadBalancerInternetSpecResponse, error) {
	cli := GetSlbClient()
	params := make(map[string]interface{})
	params["LoadBalancerId"] = loadBalancerId
	if internetChargeType != "" {
		params["InternetChargeType"] = internetChargeType
	}
	if bandwidth > 0 {
		params["Bandwidth"] = string(bandwidth)
	}
	return cli.LoadBalancer.ModifyLoadBalancerInternetSpec(params)
}

func DeleteLoadBalancer(loadBalancerId string) (slb.DeleteLoadBalancerResponse, error) {
	cli := GetSlbClient()
	return cli.LoadBalancer.DeleteLoadBalancer(map[string]interface{}{
		"LoadBalancerId": loadBalancerId,
	})
}

func SetLoadBalancerStatus(loadBalancerId string, loadBalancerStatus string) (slb.SetLoadBalancerStatusResponse, error) {
	cli := GetSlbClient()
	return cli.LoadBalancer.SetLoadBalancerStatus(map[string]interface{}{
		"LoadBalancerId":     loadBalancerId,
		"LoadBalancerStatus": loadBalancerStatus,
	})
}

func SetLoadBalancerName(loadBalancerId string, loadBalancerName string) (slb.SetLoadBalancerNameResponse, error) {
	cli := GetSlbClient()
	return cli.LoadBalancer.SetLoadBalancerName(map[string]interface{}{
		"LoadBalancerId":   loadBalancerId,
		"LoadBalancerName": loadBalancerName,
	})
}

func DescribeLoadBalancers(loadBalancer models.LoadBalancer) (slb.DescribeLoadBalancersResponse, error) {
	cli := GetSlbClient()
	params := make(map[string]interface{})
	params["RegionId"] = loadBalancer.RegionId

	return cli.LoadBalancer.DescribeLoadBalancers(params)
}

func DescribeLoadBalancerAttribute(loadBalancerId string) (slb.DescribeLoadBalancerAttributeResponse, error) {
	cli := GetSlbClient()
	return cli.LoadBalancer.DescribeLoadBalancerAttribute(map[string]interface{}{
		"LoadBalancerId": loadBalancerId,
	})
}

func DescribeRegions() (slb.DescribeRegionsResponse, error) {
	cli := GetSlbClient()
	params := make(map[string]interface{})
	return cli.LoadBalancer.DescribeRegions(params)
}
