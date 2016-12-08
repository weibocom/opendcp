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

type Listener struct {
	LoadBalancerId    string `json:"LoadBalancerId"`    //'负载均衡实例id'
	ListenerPort      int    `json:"ListenerPort"`      //'负载均衡实例前端使用的端口。取值：1-65535'
	BackendServerPort int    `json:"BackendServerPort"` //'负载均衡实例后端使用的端口。取值：1-65535'
	Bandwidth         int    `json:"Bandwidth"`         //'监听的带宽峰值。取值：-1 | 1-1000Mbps'
	Scheduler         string `json:"Scheduler"`         //'调度算法。取值：wrr | wlc。默认值：wrr'
	XForwardedFor     string `json:"XForwardedFor"`     //'是否开启通过X-Forwarded-For的方式获取来访者真实IP。取值：on | off。默认值：on'

	HealthCheck            string `json:"HealthCheck"`            //'是否开启健康检查。取值：on | off'
	HealthCheckDomain      string `json:"HealthCheckDomain"`      //'用于健康检查的域名。取值：$_ip | 用户自定义字符串|空'
	HealthCheckURI         string `json:"HealthCheckURI"`         //'当HealthCheck为on时为必选；用于健康检查的URI。'
	HealthCheckConnectPort int    `json:"HealthCheckConnectPort"` //'当HealthCheck为on时为必选；进行健康检查时使用的端口。取值：1-65535，或者-520(表示使用后端服务端口)'
	HealthyThreshold       int    `json:"HealthyThreshold"`       //'当HealthCheck为on时为必选；判定健康检查结果为success的阈值。取值：1-10'
	UnhealthyThreshold     int    `json:"UnhealthyThreshold"`     //'当HealthCheck为on时为必选；判定健康检查结果为fail的阈值。取值：1-10'
	HealthCheckTimeout     int    `json:"HealthCheckTimeout"`     //'当HealthCheck为on时为必选；每次健康检查响应的最大超时时间。取值：1-50（单位为秒）'
	HealthCheckInterval    int    `json:"HealthCheckInterval"`    //'当HealthCheck为on时为必选；进行健康检查的时间间隔。取值：1-5（单位为秒）'

	HealthCheckType     string `json:"HealthCheckType"`     //'健康检查类型。取值：tcp | http。默认值：tcp '
	HealthCheckHttpCode string `json:"HealthCheckHttpCode"` //'健康检查正常的http状态码，多个http状态码间用”,”分割。取值：http_2xx | http_3xx | http_4xx | http_5xx'

	//HTTP、HTTPS
	StickySession     string `json:"StickySession"`     //'是否开启会话保持。取值：on | off'
	StickySessionType string `json:"StickySessionType"` //'在StickySession为on时为必选；取值：insert | server设置为insert表示由负载均衡插入，设置为server表示负载均衡从后端服务器学习。'
	CookieTimeout     int    `json:"CookieTimeout"`     //'在StickySession为on且StickySessionType为insert时为必选；cookie超时时间。取值：1-86400（单位为秒）'
	Cookie            string `json:"Cookie"`            //'在StickySession为on且StickySessionType为server时为必选；服务器上配置的cookie'

	//HTTPS
	ServerCertificateId string `json:"ServerCertificateId"` //'安全证书的ID'

	//TCP、UDP
	PersistenceTimeout int `json:"PersistenceTimeout"` //'连接持久化的超时时间。取值：0-1000（单位为秒） 默认值：0，表示关闭。'

	Protocol string //'listener使用的协议'

	AccessControlStatus string //'是否开启访问控制。取值：open_white_list | close'
	SourceItems         string //'访问控制列表。支持ip地址或ip地址段的输入，多个ip地址或ip地址段间用”,”分割。'
}
