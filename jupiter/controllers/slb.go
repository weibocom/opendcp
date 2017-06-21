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



package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"strings"
	"weibo.com/opendcp/jupiter/models"
	"weibo.com/opendcp/jupiter/service/instance"
	"weibo.com/opendcp/jupiter/service/slb"
	"weibo.com/opendcp/jupiter/provider"
	"strconv"
)

const MAX_IP_NUMBER = 20

// SlbController handles requests to /slb/loadbalancer /slb/listener /slb/backendserver, the parm has to be put
// in the query string as the web framework can not parse the URL if it contains veriadic sectors.
type SlbController struct {
	BaseController
}

// @Title Create a load balancer
// @Description CreateLoadBalancer handles POST /v1/slb/loadbalancer
// @router /loadbalancer [post]
func (sc *SlbController) CreateLoadBalancer() {
	resp := ApiResponse{}
	var loadBalancer models.LoadBalancer
	body := sc.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &loadBalancer)
	if err != nil {
		beego.Error("format loadbalancer json err: ", err)
		sc.RespInputError()
		return
	}
	if loadBalancer.LoadBalancerName != "" {
		if len(loadBalancer.LoadBalancerName) > 80 || len(loadBalancer.LoadBalancerName) < 1 {
			sc.RespInputOverLimited("loadBalancer.LoadBalancerName", "larger 0 and less 81")
			return
		}
	}
	if loadBalancer.AddressType == "internet" {
		if loadBalancer.Bandwidth < 1 || loadBalancer.Bandwidth > 1000 {
			sc.RespInputOverLimited("loadBalancer.Bandwidth", "larger 0 and less 1001")
			return
		}
	}
	if loadBalancer.RegionId == "" {
		sc.RespMissingParams("loadBalancer.RegionId")
		return
	}

	r, err := slb.CreateLoadBalancer(loadBalancer)
	if err != nil {
		beego.Error("Create loadbalancer failed: " + err.Error())
		sc.RespServiceError(err)
		return
	}
	resp.Content = r
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title Change load balancer status
// @Description SetLoadBalancerStatus handles POST /v1/slb/loadbalancer/status
// @router /loadbalancer/status [post]
func (sc *SlbController) SetLoadBalancerStatus() {
	resp := ApiResponse{}
	var loadBalancer models.LoadBalancer
	body := sc.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &loadBalancer)
	if err != nil {
		beego.Error("format loadbalancer json err: ", err)
		sc.RespInputError()
		return
	}
	if loadBalancer.LoadBalancerId == "" {
		sc.RespMissingParams("loadBalancer.LoadBalancerId")
		return
	}

	if loadBalancer.LoadBalancerStatus == "" {
		sc.RespMissingParams("loadBalancer.LoadBalancerStatus")
		return
	}

	r, err := slb.SetLoadBalancerStatus(loadBalancer.LoadBalancerId, loadBalancer.LoadBalancerStatus)
	if err != nil {
		beego.Error("Create loadbalancer failed: " + err.Error())
		sc.RespServiceError(err)
		return
	}
	resp.Content = r
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title Delete load balancer status
// @Description DeleteLoadBalancerStatus handles DELETE /v1/slb/loadbalancer
// @router /loadbalancer [delete]
func (sc *SlbController) DeleteLoadBalancer() {
	resp := ApiResponse{}
	var loadBalancer models.LoadBalancer
	body := sc.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &loadBalancer)
	if err != nil {
		beego.Error("format loadbalancer json err: ", err)
		sc.RespInputError()
		return
	}
	if loadBalancer.LoadBalancerId == "" {
		sc.RespMissingParams("loadBalancer.LoadBalancerId")
		return
	}
	r, err := slb.DeleteLoadBalancer(loadBalancer.LoadBalancerId)
	if err != nil {
		beego.Error("Remove loadbalancer failed: " + err.Error())
		sc.RespServiceError(err)
		return
	}
	resp.Content = r
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title list load balancers
// @Description GetLoadBalancers handles GET /v1/slb/loadbalancers
// @router /list/:regionId [get]
func (sc *SlbController) GetLoadBalancers() {
	resp := ApiResponse{}
	regionId := sc.GetString(":regionId")
	var loadBalancer models.LoadBalancer
	loadBalancer.RegionId = regionId
	if loadBalancer.RegionId == "" {
		sc.RespMissingParams("loadBalancer.RegionId")
		return
	}
	r, err := slb.DescribeLoadBalancers(loadBalancer)
	if err != nil {
		beego.Error("Get loadbalancers failed: " + err.Error())
		sc.RespServiceError(err)
		return
	}
	resp.Content = r.LoadBalancers.LoadBalancer
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title get a load balancer
// @Description GetLoadBalancer handles GET /v1/slb/loadbalancer
// @router /loadbalancer/:loadbalancerid [get]
func (sc *SlbController) GetLoadBalancer() {
	resp := ApiResponse{}
	loadBalancerId := sc.GetString(":loadbalancerid")
	r, err := slb.DescribeLoadBalancerAttribute(loadBalancerId)
	if err != nil {
		beego.Error("Get loadbalancer failed: " + err.Error())
		sc.RespServiceError(err)
		return
	}
	resp.Content = r
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title Add backend servers to load balancer
// @Description AddBackendServers handles GET /v1/slb/backendservers
// @router /backendservers/:loadbalancerid [post]
func (sc *SlbController) AddBackendServers() {
	resp := ApiResponse{}
	body := sc.Ctx.Input.RequestBody
	loadBalancerId := sc.GetString(":loadbalancerid")

	if len(strings.Split(string(body), ",")) > MAX_IP_NUMBER {
		sc.RespInputOverLimited("BackendServerList", "SLB max number of backendServer is 20.")
		return
	}
	r, err := slb.AddBackendServers(loadBalancerId, string(body))
	if err != nil {
		beego.Error("Add backendservers failed: " + err.Error())
		sc.RespServiceError(err)
		return
	}
	resp.Content = r
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title Remove backend servers in load balancer
// @Description RemoveBackendServers handles DELETE /v1/slb/backendservers
// @router /backendservers/:loadbalancerid [delete]
func (sc *SlbController) RemoveBackendServers() {
	resp := ApiResponse{}
	body := sc.Ctx.Input.RequestBody
	loadBalancerId := sc.GetString(":loadbalancerid")

	if len(strings.Split(string(body), ",")) > MAX_IP_NUMBER {
		sc.RespInputOverLimited("BackendServerList", "SLB max number of backendServer is 20.")
		return
	}
	r, err := slb.RemoveBackendServers(loadBalancerId, string(body))
	if err != nil {
		beego.Error("Remove backendservers failed: " + err.Error())
		sc.RespServiceError(err)
		return
	}
	resp.Content = r
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title Set backend servers in load balancer
// @Description SetBackendServers handles PUT /v1/slb/backendservers
// @router /backendservers/:loadbalancerid [put]
func (sc *SlbController) SetBackendServers() {
	resp := ApiResponse{}
	body := sc.Ctx.Input.RequestBody
	loadBalancerId := sc.GetString(":loadbalancerid")

	if len(strings.Split(string(body), ",")) > MAX_IP_NUMBER {
		sc.RespInputOverLimited("BackendServerList", "SLB max number of backendServer is 20.")
		return
	}
	r, err := slb.SetBackendServers(loadBalancerId, string(body))
	if err != nil {
		beego.Error("Remove backendservers failed: " + err.Error())
		sc.RespServiceError(err)
		return
	}
	resp.Content = r
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title check backend servers health status
// @Description DescribeHealthStatus handles GET /v1/slb/backendservers/healthstatus
// @router /backendservers/healthstatus/:loadbalancerid [get]
func (sc *SlbController) DescribeHealthStatus() {
	resp := ApiResponse{}
	loadBalancerId := sc.GetString(":loadbalancerid")

	r, err := slb.DescribeHealthStatus(loadBalancerId)
	if err != nil {
		beego.Error("Remove backendservers failed: " + err.Error())
		sc.RespServiceError(err)
		return
	}
	servers := make([]models.BackendServer, 0)
	providerDriver, _ := provider.New("aliyun")
	for _, v := range r.BackendServers.BackendServer {
		ins, _ := providerDriver.GetInstance(v.ServerId)
		var ip string = ""
		if len(ins.PrivateIpAddress) > 0 {
			ip = ins.PrivateIpAddress
		} else if len(ins.PublicIpAddress) > 0 {
			ip = ins.PublicIpAddress
		}
		servers = append(servers, models.BackendServer{
			ServerId: v.ServerId,
			Weight:   v.Weight,
			ServerHealthStatus: v.ServerHealthStatus,
			Address: ip,
		})
	}
	resp.Content = servers
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title Set backend servers to load balance by ip
// @Description DescribeHealthStatus handles PUT /v1/slb/backendservers/by_ip
// @router /backendservers/by_ip [put]
func (sc *SlbController) SetBackendOfLoadBalance() {
	resp := ApiResponse{}
	body := sc.Ctx.Input.RequestBody
	var setServer models.BackendServerRequest
	err := json.Unmarshal(body, &setServer)
	if err != nil {
		beego.Error("parameter error: ", err)
		sc.RespInputError()
		return
	}
	servers := make([]models.BackendServer, 0)
	if setServer.BackendServerList != nil {
		for _, v := range setServer.BackendServerList {
			servers = append(servers, models.BackendServer{
				ServerId: v.ServerId,
				Weight:   v.Weight,
			})
		}
	}
	if len(servers) > 20 {
		sc.RespInputOverLimited("BackendServerList", "SLB max number of backendServer is 20.")
		return
	}
	if len(servers) > 0 {
		bytes, err := json.Marshal(&servers)
		if err != nil {
			beego.Error("parse servers error: ", err)
			return
		}
		r, err := slb.SetBackendServers(setServer.LoadBalancerId, string(bytes))
		if err != nil {
			beego.Error("set backendservers failed: " + err.Error())
			sc.RespServiceError(err)
			return
		}
		resp.Content = r
	}
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title Add backend servers to load balance by ip
// @Description DescribeHealthStatus handles POST /v1/slb/backendservers/by_ip
// @router /backendservers/by_ip [post]
func (sc *SlbController) AddToLoadBalance() {
	bizId := sc.Ctx.Input.Header("bizId")
	bid, err := strconv.Atoi(bizId)
	if bizId=="" || err != nil {
		beego.Error("Get instance bizId err!")
		sc.RespInputError()
		return
	}

	resp := ApiResponse{}
	body := sc.Ctx.Input.RequestBody
	var addServer models.BackendServerRequest
	err = json.Unmarshal(body, &addServer)
	if err != nil {
		beego.Error("parameter error: ", err)
		sc.RespInputError()
		return
	}
	servers := make([]models.BackendServer, 0)
	if addServer.BackendServerList != nil {
		for _, v := range addServer.BackendServerList {
			ins, err := instance.GetInstanceByIp(v.Address, bid)
			if err != nil {
				beego.Error("get instance failed: ", err)
				sc.RespServiceError(err)
				return
			}
			servers = append(servers, models.BackendServer{
				ServerId: ins.InstanceId,
				Weight:   v.Weight,
			})
		}
	}
	if len(servers) > 20 {
		sc.RespInputOverLimited("BackendServerList", "SLB max number of backendServer is 20.")
		return
	}
	if len(servers) > 0 {
		bytes, err := json.Marshal(&servers)
		if err != nil {
			beego.Error("parse servers error: ", err)
			sc.RespInputError()
		}
		r, err := slb.AddBackendServers(addServer.LoadBalancerId, string(bytes))
		if err != nil {
			beego.Error("add backendservers failed: " + err.Error())
			sc.RespServiceError(err)
			return
		}
		resp.Content = r
	}
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title Delete backend servers to load balance by ip
// @Description DescribeHealthStatus handles DELETE /v1/slb/backendservers/by_ip
// @router /backendservers/by_ip [delete]
func (sc *SlbController) RemoveFromLoadBalance() {
	bizId := sc.Ctx.Input.Header("bizId")
	bid, err := strconv.Atoi(bizId)
	if bizId=="" || err != nil {
		beego.Error("Get instance bizId err!")
		sc.RespInputError()
		return
	}
	resp := ApiResponse{}
	body := sc.Ctx.Input.RequestBody
	var removeServer models.BackendServerRequest
	err = json.Unmarshal(body, &removeServer)
	if err != nil {
		beego.Error("parameter error: ", err)
		sc.RespInputError()
		return
	}
	servers := make([]string, 0)
	if removeServer.BackendServerList != nil {
		for _, v := range removeServer.BackendServerList {
			ins, err := instance.GetInstanceByIp(v.Address, bid)
			if err != nil {
				beego.Error("get resource failed: ", err)
				sc.RespServiceError(err)
				return
			}
			servers = append(servers, ins.InstanceId)
		}
	}
	if len(servers) > 20 {
		sc.RespInputOverLimited("BackendServerList", "SLB max number of backendServer is 20.")
		return
	}
	if len(servers) > 0 {
		bytes, err := json.Marshal(&servers)
		if err != nil {
			beego.Error("parse servers error: ", err)
			sc.RespInputError()
			return
		}
		r, err := slb.RemoveBackendServers(removeServer.LoadBalancerId, string(bytes))
		if err != nil {
			beego.Error("remove backendservers failed: " + err.Error())
			sc.RespServiceError(err)
			return
		}
		resp.Content = r
	}
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title Add backend servers to load balance by ip
// @Description DescribeHealthStatus handles POST /v1/slb/listener
// @router /listener [post]
func (sc *SlbController) CreateLoadBalancerListener() {
	resp := ApiResponse{}
	var listener models.Listener
	body := sc.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &listener)
	if err != nil {
		beego.Error("format listener json err: ", err)
		sc.RespInputError()
		return
	}
	if listener.Bandwidth < 1 || listener.Bandwidth > 1000 {
		if listener.Bandwidth != -1 {
			sc.RespInputOverLimited("listener.Bandwidth", "need to larger 0 and less 1001 or equal -1")
			return
		}
	}
	if listener.StickySession == "on" {
		if listener.CookieTimeout < 1 || listener.CookieTimeout > 86400 {
			sc.RespInputOverLimited("listener.CookieTimeout", "need to larger 0 and less 86401")
			return
		}
	}
	if listener.HealthCheck == "on" {
		if listener.HealthyThreshold < 1 || listener.HealthyThreshold > 10 {
			sc.RespInputOverLimited("listener.HealthyThreshold", "need to larger 0 and less 11")
			return
		}
		if listener.UnhealthyThreshold < 1 || listener.UnhealthyThreshold > 10 {
			sc.RespInputOverLimited("listener.UnhealthyThreshold", "need to larger 0 and less 11")
			return
		}
		if listener.HealthCheckTimeout < 1 || listener.HealthCheckTimeout > 50 {
			sc.RespInputOverLimited("listener.HealthCheckTimeout", "need to larger 0 and less 51")
			return
		}
		if listener.HealthCheckInterval < 1 || listener.HealthCheckInterval > 5 {
			sc.RespInputOverLimited("listener.HealthCheckInterval", "need to larger 0 and less 6")
			return
		}
	}
	protocol := listener.Protocol
	var r interface{}
	switch protocol {
	case "HTTP", "http":
		r, err = slb.CreateLoadBalancerHTTPListener(listener)
	case "HTTPS", "https":
		r, err = slb.CreateLoadBalancerHTTPSListener(listener)
	case "TCP", "tcp":
		if listener.PersistenceTimeout > 1000 {
			sc.RespInputOverLimited("listener.PersistenceTimeout", "need to larger or equal 0 and less 1001")
			return
		}
		r, err = slb.CreateLoadBalancerTCPListener(listener)
	case "UDP", "udp":
		if listener.PersistenceTimeout > 86400 {
			sc.RespInputOverLimited("listener.PersistenceTimeout", "need to larger or equal 0 and less 86401")
			return
		}
		r, err = slb.CreateLoadBalancerUDPListener(listener)
	}
	if err != nil {
		beego.Error("Create listener failed: " + err.Error())
		sc.RespServiceError(err)
		return
	}
	resp.Content = r
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title Set listener attribute
// @Description SetLoadBalancerListenerAttribute handles PUT /v1/slb/listener
// @router /listener [put]
func (sc *SlbController) SetLoadBalancerListenerAttribute() {
	resp := ApiResponse{}
	var listener models.Listener
	body := sc.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &listener)
	if err != nil {
		beego.Error("format listener json err: ", err)
		sc.RespInputError()
		return
	}
	protocol := listener.Protocol
	var r interface{}
	switch protocol {
	case "HTTP", "http":
		r, err = slb.SetLoadBalancerHTTPListenerAttribute(listener)
	case "HTTPS", "https":
		r, err = slb.SetLoadBalancerHTTPSListenerAttribute(listener)
	case "TCP", "tcp":
		r, err = slb.SetLoadBalancerTCPListenerAttribute(listener)
	case "UDP", "udp":
		r, err = slb.SetLoadBalancerUDPListenerAttribute(listener)

	}
	if err != nil {
		beego.Error("Set listener attribute failed: " + err.Error())
		sc.RespServiceError(err)
		return
	}
	resp.Content = r
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title Describe load balancer listener attribute
// @Description DescribeLoadBalancerListenerAttribute handles GET /v1/slb/listener
// @router /listener [get]
func (sc *SlbController) DescribeLoadBalancerListenerAttribute() {
	resp := ApiResponse{}
	loadBalancerId := sc.GetString("LoadBalancerId")
	listenerPort, err := sc.GetInt("ListenerPort")
	if err != nil {
		sc.RespInputError()
		return
	}
	protocol := sc.GetString("Protocol")
	beego.Info(loadBalancerId, listenerPort, protocol)
	var r interface{}
	switch protocol {
	case "HTTP", "http":
		r, err = slb.DescribeLoadBalancerHTTPListenerAttribute(loadBalancerId, listenerPort)
	case "HTTPS", "https":
		r, err = slb.DescribeLoadBalancerHTTPSListenerAttribute(loadBalancerId, listenerPort)
	case "TCP", "tcp":
		r, err = slb.DescribeLoadBalancerTCPListenerAttribute(loadBalancerId, listenerPort)
	case "UDP", "udp":
		r, err = slb.DescribeLoadBalancerUDPListenerAttribute(loadBalancerId, listenerPort)

	}
	if err != nil {
		beego.Error("Decribe listener attribute failed: " + err.Error())
		sc.RespServiceError(err)
		return
	}
	resp.Content = r
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title Delete load balancer listener
// @Description DeleteLoadBalancerListener handles DELETE /v1/slb/listener
// @router /listener [delete]
func (sc *SlbController) DeleteLoadBalancerListener() {
	resp := ApiResponse{}
	var listener models.Listener
	body := sc.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &listener)
	if err != nil {
		beego.Error("format listener json err: ", err)
		sc.RespInputError()
		return
	}
	r, err := slb.DeleteLoadBalancerListener(listener.LoadBalancerId, listener.ListenerPort)

	if err != nil {
		beego.Error("Delete listener failed: " + err.Error())
		sc.RespServiceError(err)
		return
	}
	resp.Content = r
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title start load balancer listener
// @Description StartLoadBalancerListener handles POST /v1/slb/listener/start
// @router /listener/start [post]
func (sc *SlbController) StartLoadBalancerListener() {
	resp := ApiResponse{}
	var listener models.Listener
	body := sc.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &listener)
	if err != nil {
		beego.Error("format listener json err: ", err)
		sc.RespInputError()
		return
	}
	r, err := slb.StartLoadBalancerListener(listener.LoadBalancerId, listener.ListenerPort)

	if err != nil {
		beego.Error("Start listener failed: " + err.Error())
		sc.RespServiceError(err)
		return
	}
	resp.Content = r
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title Stop load balancer listener
// @Description StopLoadBalancerListener handles POST /v1/slb/listener/stop
// @router /listener/stop [post]
func (sc *SlbController) StopLoadBalancerListener() {
	resp := ApiResponse{}
	var listener models.Listener
	body := sc.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &listener)
	if err != nil {
		beego.Error("format listener json err: ", err)
		sc.RespInputError()
		return
	}
	r, err := slb.StopLoadBalancerListener(listener.LoadBalancerId, listener.ListenerPort)

	if err != nil {
		beego.Error("Stop listener failed: " + err.Error())
		sc.RespServiceError(err)
		return
	}
	resp.Content = r
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title On or Down access control
// @Description SetListenerAccessControlStatus handles POST /v1/slb/listener/whitelist/status
// @router /listener/whitelist/status [post]
func (sc *SlbController) SetListenerAccessControlStatus() {
	resp := ApiResponse{}
	var listener models.Listener
	body := sc.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &listener)
	if err != nil {
		beego.Error("format listener json err: ", err)
		sc.RespInputError()
		return
	}
	r, err := slb.SetListenerAccessControlStatus(listener.LoadBalancerId, listener.ListenerPort, listener.AccessControlStatus)

	if err != nil {
		beego.Error("Set listener_access_control_status failed: " + err.Error())
		sc.RespServiceError(err)
		return
	}
	resp.Content = r
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title Add ip in listener access control
// @Description AddListenerWhiteListItem handles POST /v1/slb/listener/whitelist
// @router /listener/whitelist [post]
func (sc *SlbController) AddListenerWhiteListItem() {
	resp := ApiResponse{}
	var listener models.Listener
	body := sc.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &listener)
	if err != nil {
		beego.Error("format listener json err: ", err)
		sc.RespInputError()
		return
	}
	r, err := slb.AddListenerWhiteListItem(listener.LoadBalancerId, listener.ListenerPort, listener.SourceItems)

	if err != nil {
		beego.Error("Add listener_whitelist_item failed: " + err.Error())
		sc.RespServiceError(err)
		return
	}
	resp.Content = r
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title Remove ip in listener access control
// @Description RemoveListenerWhiteListItem handles DELETE /v1/slb/listener/whitelist
// @router /listener/whitelist [delete]
func (sc *SlbController) RemoveListenerWhiteListItem() {
	resp := ApiResponse{}
	var listener models.Listener
	body := sc.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &listener)
	if err != nil {
		beego.Error("format listener json err: ", err)
		sc.RespInputError()
		return
	}
	r, err := slb.RemoveListenerWhiteListItem(listener.LoadBalancerId, listener.ListenerPort, listener.SourceItems)

	if err != nil {
		beego.Error("Set listener_access_control_status failed: " + err.Error())
		sc.RespServiceError(err)
		return
	}
	resp.Content = r
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}

// @Title Describe listener access control attribute
// @Description DescribeListenerAccessControlAttribute handles GET /v1/slb/listener/whitelist
// @router /listener/whitelist [get]
func (sc *SlbController) DescribeListenerAccessControlAttribute() {
	resp := ApiResponse{}
	var listener models.Listener
	body := sc.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &listener)
	if err != nil {
		beego.Error("format listener json err: ", err)
		sc.RespInputError()
	}
	r, err := slb.DescribeListenerAccessControlAttribute(listener.LoadBalancerId, listener.ListenerPort)

	if err != nil {
		beego.Error("Describe listener_access_control_attribute failed: " + err.Error())
		sc.RespServiceError(err)
		return
	}
	resp.Content = r
	sc.ApiResponse = resp
	sc.Status = SERVICE_SUCCESS
	sc.RespJsonWithStatus()
}
