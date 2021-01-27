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
	KeyId              string `json:"keyId"`
	LoadBalancerSpec   string `json:"LoadBalancerSpec"` //负载均衡规格
}
