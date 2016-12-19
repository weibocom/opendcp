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
