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

package models

type LoadBalancer struct {
	LoadBalancerId     string `json:"LoadBalancerId"`     //'负载均衡实例id'
	LoadBalancerName   string `json:"LoadBalancerName"`   //'负载均衡实例名称'
	Bandwidth          int    `json:"Bandwidth"`          //'实例带宽峰值,取值：1-1000（单位为Mbps）'
	RegionId           string `json:"RegionId"`           //'负载均衡实例所属的Region编号'
	Address            string `json:"Address"`            //'负载均衡实例服务地址'
	LoadBalancerStatus string `json:"LoadBalancerStatus"` //'负载均衡实例状态，取值：inactive | active | locked'
	InternetChargeType string `json:"InternetChargeType"` //'公网类型实例付费方式。取值：paybybandwidth | paybytraffic'
	AddressType        string `json:"AddressType"`        //'Address类型,取值：internet | intranet'
	NetworkType        string `json:"NetworkType"`        //'负载均衡实例网络类型，取值：vpc | classic'
	VpcId              string `json:"VpcId"`              //vpcid
	VSwitchId          string `json:"VSwitchId"`          //'VSwitchId'
	ServerId           string `json:"ServerId"`           //ECS实例ID
	CreateTime         string `json:"CreateTime"`         //'负载均衡实例创建时间'
	RegionIdAlias      string `json:"RegionIdAlias"`      //'负载均衡实例所属的Region编号别名'
	Reason             string `json:"Reason"`
}
