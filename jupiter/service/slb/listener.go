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

func CreateLoadBalancerHTTPListener(listener models.Listener) (slb.CreateLoadBalancerHTTPListenerResponse, error) {
	cli := GetSlbClientByKeyId(listener.KeyId)
	params := make(map[string]interface{})
	params["LoadBalancerId"] = listener.LoadBalancerId
	params["ListenerPort"] = listener.ListenerPort
	params["BackendServerPort"] = listener.BackendServerPort
	params["Bandwidth"] = listener.Bandwidth
	params["StickySession"] = listener.StickySession
	params["HealthCheck"] = listener.HealthCheck
	if listener.Scheduler != "" {
		params["Scheduler"] = listener.Scheduler
	}
	if params["StickySession"] == "on" {
		params["StickySessionType"] = listener.StickySessionType
		if params["StickySessionType"] == "insert" {
			params["CookieTimeout"] = listener.CookieTimeout
		} else if params["StickySessionType"] == "server" {
			params["Cookie"] = listener.Cookie
		}
	}
	if params["HealthCheck"] == "on" {
		params["HealthCheckDomain"] = listener.HealthCheckDomain
		params["HealthCheckURI"] = listener.HealthCheckURI
		params["HealthCheckConnectPort"] = listener.HealthCheckConnectPort
		params["HealthyThreshold"] = listener.HealthyThreshold
		params["UnhealthyThreshold"] = listener.UnhealthyThreshold
		params["HealthCheckTimeout"] = listener.HealthCheckTimeout
		params["HealthCheckInterval"] = listener.HealthCheckInterval
		params["HealthCheckHttpCode"] = listener.HealthCheckHttpCode
	}
	return cli.Listener.CreateLoadBalancerHTTPListener(params)
}

func CreateLoadBalancerHTTPSListener(listener models.Listener) (slb.CreateLoadBalancerHTTPSListenerResponse, error) {
	cli := GetSlbClientByKeyId(listener.KeyId)
	params := make(map[string]interface{})
	params["LoadBalancerId"] = listener.LoadBalancerId
	params["ListenerPort"] = listener.ListenerPort
	params["BackendServerPort"] = listener.BackendServerPort
	params["Bandwidth"] = listener.Bandwidth
	params["StickySession"] = listener.StickySession
	params["HealthCheck"] = listener.HealthCheck
	params["ServerCertificateId"] = listener.ServerCertificateId
	if listener.Scheduler != "" {
		params["Scheduler"] = listener.Scheduler
	}
	if params["StickySession"] == "on" {
		params["StickySessionType"] = listener.StickySessionType
		if params["StickySessionType"] == "insert" {
			params["CookieTimeout"] = listener.CookieTimeout
		} else if params["StickySessionType"] == "server" {
			params["Cookie"] = listener.Cookie
		}
	}
	if params["HealthCheck"] == "on" {
		params["HealthCheckDomain"] = listener.HealthCheckDomain
		params["HealthCheckURI"] = listener.HealthCheckURI
		params["HealthCheckConnectPort"] = listener.HealthCheckConnectPort
		params["HealthyThreshold"] = listener.HealthyThreshold
		params["UnhealthyThreshold"] = listener.UnhealthyThreshold
		params["HealthCheckTimeout"] = listener.HealthCheckTimeout
		params["HealthCheckInterval"] = listener.HealthCheckInterval
		params["HealthCheckHttpCode"] = listener.HealthCheckHttpCode
	}
	return cli.Listener.CreateLoadBalancerHTTPSListener(params)
}

func CreateLoadBalancerTCPListener(listener models.Listener) (slb.CreateLoadBalancerTCPListenerResponse, error) {
	cli := GetSlbClientByKeyId(listener.KeyId)
	params := make(map[string]interface{})
	params["LoadBalancerId"] = listener.LoadBalancerId
	params["ListenerPort"] = listener.ListenerPort
	params["BackendServerPort"] = listener.BackendServerPort
	params["Bandwidth"] = listener.Bandwidth
	params["HealthCheck"] = listener.HealthCheck
	params["PersistenceTimeout"] = listener.PersistenceTimeout
	if listener.Scheduler != "" {
		params["Scheduler"] = listener.Scheduler
	}
	if params["HealthCheck"] == "on" {
		params["HealthCheckDomain"] = listener.HealthCheckDomain
		params["HealthCheckURI"] = listener.HealthCheckURI
		params["HealthCheckConnectPort"] = listener.HealthCheckConnectPort
		params["HealthyThreshold"] = listener.HealthyThreshold
		params["UnhealthyThreshold"] = listener.UnhealthyThreshold
		params["HealthCheckTimeout"] = listener.HealthCheckTimeout
		params["HealthCheckInterval"] = listener.HealthCheckInterval
		params["HealthCheckHttpCode"] = listener.HealthCheckHttpCode
	}
	return cli.Listener.CreateLoadBalancerTCPListener(params)
}

func CreateLoadBalancerUDPListener(listener models.Listener) (slb.CreateLoadBalancerUDPListenerResponse, error) {
	cli := GetSlbClientByKeyId(listener.KeyId)
	params := make(map[string]interface{})
	params["LoadBalancerId"] = listener.LoadBalancerId
	params["ListenerPort"] = listener.ListenerPort
	params["BackendServerPort"] = listener.BackendServerPort
	params["Bandwidth"] = listener.Bandwidth
	params["HealthCheck"] = listener.HealthCheck
	params["PersistenceTimeout"] = listener.PersistenceTimeout
	if listener.Scheduler != "" {
		params["Scheduler"] = listener.Scheduler
	}
	if params["HealthCheck"] == "on" {
		params["HealthCheckDomain"] = listener.HealthCheckDomain
		params["HealthCheckURI"] = listener.HealthCheckURI
		params["HealthCheckConnectPort"] = listener.HealthCheckConnectPort
		params["HealthyThreshold"] = listener.HealthyThreshold
		params["UnhealthyThreshold"] = listener.UnhealthyThreshold
		params["HealthCheckTimeout"] = listener.HealthCheckTimeout
		params["HealthCheckInterval"] = listener.HealthCheckInterval
		params["HealthCheckHttpCode"] = listener.HealthCheckHttpCode
	}
	return cli.Listener.CreateLoadBalancerUDPListener(params)
}

func DeleteLoadBalancerListener(keyId string, loadBalancerId string, listenerPort int) (slb.DeleteLoadBalancerListenerResponse, error) {
	cli := GetSlbClientByKeyId(keyId)
	params := make(map[string]interface{})
	params["LoadBalancerId"] = loadBalancerId
	params["ListenerPort"] = listenerPort
	return cli.Listener.DeleteLoadBalancerListener(params)
}

func StartLoadBalancerListener(keyId string, loadBalancerId string, listenerPort int) (slb.StartLoadBalancerListenerResponse, error) {
	cli := GetSlbClientByKeyId(keyId)
	params := make(map[string]interface{})
	params["LoadBalancerId"] = loadBalancerId
	params["ListenerPort"] = listenerPort
	return cli.Listener.StartLoadBalancerListener(params)
}

func StopLoadBalancerListener(keyId string, loadBalancerId string, listenerPort int) (slb.StopLoadBalancerListenerResponse, error) {
	cli := GetSlbClientByKeyId(keyId)
	params := make(map[string]interface{})
	params["LoadBalancerId"] = loadBalancerId
	params["ListenerPort"] = listenerPort
	return cli.Listener.StopLoadBalancerListener(params)
}

func SetListenerAccessControlStatus(keyId string, loadBalancerId string, listenerPort int, accessControlStatus string) (slb.SetListenerAccessControlStatusResponse, error) {
	cli := GetSlbClientByKeyId(keyId)
	params := make(map[string]interface{})
	params["LoadBalancerId"] = loadBalancerId
	params["ListenerPort"] = listenerPort
	params["AccessControlStatus"] = accessControlStatus
	return cli.Listener.SetListenerAccessControlStatus(params)
}

func AddListenerWhiteListItem(keyId string, loadBalancerId string, listenerPort int, sourceItems string) (slb.AddListenerWhiteListItemResponse, error) {
	cli := GetSlbClientByKeyId(keyId)
	params := make(map[string]interface{})
	params["LoadBalancerId"] = loadBalancerId
	params["ListenerPort"] = listenerPort
	params["SourceItems"] = sourceItems
	return cli.Listener.AddListenerWhiteListItem(params)
}

func RemoveListenerWhiteListItem(keyId string, loadBalancerId string, listenerPort int, sourceItems string) (slb.RemoveListenerWhiteListItemResponse, error) {
	cli := GetSlbClientByKeyId(keyId)
	params := make(map[string]interface{})
	params["LoadBalancerId"] = loadBalancerId
	params["ListenerPort"] = listenerPort
	params["SourceItems"] = sourceItems
	return cli.Listener.RemoveListenerWhiteListItem(params)
}

func DescribeListenerAccessControlAttribute(keyId string, loadBalancerId string, listenerPort int) (slb.DescribeListenerAccessControlAttributeResponse, error) {
	cli := GetSlbClientByKeyId(keyId)
	params := make(map[string]interface{})
	params["LoadBalancerId"] = loadBalancerId
	params["ListenerPort"] = listenerPort
	return cli.Listener.DescribeListenerAccessControlAttribute(params)
}

func SetLoadBalancerHTTPListenerAttribute(listener models.Listener) (slb.SetLoadBalancerHTTPListenerAttributeResponse, error) {
	cli := GetSlbClientByKeyId(listener.KeyId)
	params := make(map[string]interface{})
	params["LoadBalancerId"] = listener.LoadBalancerId
	params["ListenerPort"] = listener.ListenerPort
	params["Bandwidth"] = listener.Bandwidth
	params["StickySession"] = listener.StickySession
	params["HealthCheck"] = listener.HealthCheck
	if listener.Scheduler != "" {
		params["Scheduler"] = listener.Scheduler
	}
	if params["StickySession"] == "on" {
		params["StickySessionType"] = listener.StickySessionType
		if params["StickySessionType"] == "insert" {
			params["CookieTimeout"] = listener.CookieTimeout
		} else if params["StickySessionType"] == "server" {
			params["Cookie"] = listener.Cookie
		}
	}
	if params["HealthCheck"] == "on" {
		params["HealthCheckDomain"] = listener.HealthCheckDomain
		params["HealthCheckURI"] = listener.HealthCheckURI
		params["HealthCheckConnectPort"] = listener.HealthCheckConnectPort
		params["HealthyThreshold"] = listener.HealthyThreshold
		params["UnhealthyThreshold"] = listener.UnhealthyThreshold
		params["HealthCheckTimeout"] = listener.HealthCheckTimeout
		params["HealthCheckInterval"] = listener.HealthCheckInterval
		params["HealthCheckHttpCode"] = listener.HealthCheckHttpCode
	}
	return cli.Listener.SetLoadBalancerHTTPListenerAttribute(params)
}

func SetLoadBalancerHTTPSListenerAttribute(listener models.Listener) (slb.SetLoadBalancerHTTPSListenerAttributeResponse, error) {
	cli := GetSlbClientByKeyId(listener.KeyId)
	params := make(map[string]interface{})
	params["LoadBalancerId"] = listener.LoadBalancerId
	params["ListenerPort"] = listener.ListenerPort
	params["Bandwidth"] = listener.Bandwidth
	params["StickySession"] = listener.StickySession
	params["HealthCheck"] = listener.HealthCheck
	params["ServerCertificateId"] = listener.ServerCertificateId
	if listener.Scheduler != "" {
		params["Scheduler"] = listener.Scheduler
	}
	if params["StickySession"] == "on" {
		params["StickySessionType"] = listener.StickySessionType
		if params["StickySessionType"] == "insert" {
			params["CookieTimeout"] = listener.CookieTimeout
		} else if params["StickySessionType"] == "server" {
			params["Cookie"] = listener.Cookie
		}
	}
	if params["HealthCheck"] == "on" {
		params["HealthCheckDomain"] = listener.HealthCheckDomain
		params["HealthCheckURI"] = listener.HealthCheckURI
		params["HealthCheckConnectPort"] = listener.HealthCheckConnectPort
		params["HealthyThreshold"] = listener.HealthyThreshold
		params["UnhealthyThreshold"] = listener.UnhealthyThreshold
		params["HealthCheckTimeout"] = listener.HealthCheckTimeout
		params["HealthCheckInterval"] = listener.HealthCheckInterval
		params["HealthCheckHttpCode"] = listener.HealthCheckHttpCode
	}
	return cli.Listener.SetLoadBalancerHTTPSListenerAttribute(params)
}

func SetLoadBalancerTCPListenerAttribute(listener models.Listener) (slb.SetLoadBalancerTCPListenerAttributeResponse, error) {
	cli := GetSlbClientByKeyId(listener.KeyId)
	params := make(map[string]interface{})
	params["LoadBalancerId"] = listener.LoadBalancerId
	params["ListenerPort"] = listener.ListenerPort
	params["Bandwidth"] = listener.Bandwidth
	params["HealthCheck"] = listener.HealthCheck
	params["PersistenceTimeout"] = listener.PersistenceTimeout
	if listener.Scheduler != "" {
		params["Scheduler"] = listener.Scheduler
	}
	if params["HealthCheck"] == "on" {
		params["HealthCheckDomain"] = listener.HealthCheckDomain
		params["HealthCheckURI"] = listener.HealthCheckURI
		params["HealthCheckConnectPort"] = listener.HealthCheckConnectPort
		params["HealthyThreshold"] = listener.HealthyThreshold
		params["UnhealthyThreshold"] = listener.UnhealthyThreshold
		params["HealthCheckTimeout"] = listener.HealthCheckTimeout
		params["HealthCheckInterval"] = listener.HealthCheckInterval
		params["HealthCheckHttpCode"] = listener.HealthCheckHttpCode
	}
	return cli.Listener.SetLoadBalancerTCPListenerAttribute(params)
}

func SetLoadBalancerUDPListenerAttribute(listener models.Listener) (slb.SetLoadBalancerUDPListenerAttributeResponse, error) {
	cli := GetSlbClientByKeyId(listener.KeyId)
	params := make(map[string]interface{})
	params["LoadBalancerId"] = listener.LoadBalancerId
	params["ListenerPort"] = listener.ListenerPort
	params["Bandwidth"] = listener.Bandwidth
	params["HealthCheck"] = listener.HealthCheck
	params["PersistenceTimeout"] = listener.PersistenceTimeout
	if listener.Scheduler != "" {
		params["Scheduler"] = listener.Scheduler
	}
	if params["HealthCheck"] == "on" {
		params["HealthCheckDomain"] = listener.HealthCheckDomain
		params["HealthCheckURI"] = listener.HealthCheckURI
		params["HealthCheckConnectPort"] = listener.HealthCheckConnectPort
		params["HealthyThreshold"] = listener.HealthyThreshold
		params["UnhealthyThreshold"] = listener.UnhealthyThreshold
		params["HealthCheckTimeout"] = listener.HealthCheckTimeout
		params["HealthCheckInterval"] = listener.HealthCheckInterval
		params["HealthCheckHttpCode"] = listener.HealthCheckHttpCode
	}
	return cli.Listener.SetLoadBalancerUDPListenerAttribute(params)
}

func DescribeLoadBalancerHTTPListenerAttribute(loadBalancerId string, listenerPort int, keyId string) (slb.DescribeLoadBalancerHTTPListenerAttributeResponse, error) {
	cli := GetSlbClientByKeyId(keyId)
	params := make(map[string]interface{})
	params["LoadBalancerId"] = loadBalancerId
	params["ListenerPort"] = listenerPort
	return cli.Listener.DescribeLoadBalancerHTTPListenerAttribute(params)
}

func DescribeLoadBalancerHTTPSListenerAttribute(loadBalancerId string, listenerPort int, keyId string) (slb.DescribeLoadBalancerHTTPSListenerAttributeResponse, error) {
	cli := GetSlbClientByKeyId(keyId)
	params := make(map[string]interface{})
	params["LoadBalancerId"] = loadBalancerId
	params["ListenerPort"] = listenerPort
	return cli.Listener.DescribeLoadBalancerHTTPSListenerAttribute(params)
}

func DescribeLoadBalancerTCPListenerAttribute(loadBalancerId string, listenerPort int, keyId string) (slb.DescribeLoadBalancerTCPListenerAttributeResponse, error) {
	cli := GetSlbClientByKeyId(keyId)
	params := make(map[string]interface{})
	params["LoadBalancerId"] = loadBalancerId
	params["ListenerPort"] = listenerPort
	return cli.Listener.DescribeLoadBalancerTCPListenerAttribute(params)
}

func DescribeLoadBalancerUDPListenerAttribute(loadBalancerId string, listenerPort int, keyId string) (slb.DescribeLoadBalancerUDPListenerAttributeResponse, error) {
	cli := GetSlbClientByKeyId(keyId)
	params := make(map[string]interface{})
	params["LoadBalancerId"] = loadBalancerId
	params["ListenerPort"] = listenerPort
	return cli.Listener.DescribeLoadBalancerUDPListenerAttribute(params)
}
