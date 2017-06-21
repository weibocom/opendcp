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



package provider

import (
	"fmt"
	"sort"
	"time"

	"weibo.com/opendcp/jupiter/models"
)

type ProviderDriver interface {
	Create(cluster *models.Cluster, number int) (instanceIds []string, errs []error)
	Delete(instanceId string) (time.Time, error)
	Start(instanceId string) (bool, error)
	Stop(instanceId string) (bool, error)
	WaitForInstanceToStop(instanceId string) bool
	WaitToStartInstance(instanceId string) bool
	List(regionId string, pageNumber int, pageSize int) (*models.ListInstancesResponse, error)
	ListInstanceTypes() ([]string, error)
	GetInstance(instanceId string) (*models.Instance, error)
	ListSecurityGroup(regionId string, vpcId string) (*models.SecurityGroupsResp, error)
	ListAvailabilityZones(regionId string) (*models.AvailabilityZonesResp, error)
	ListRegions() (*models.RegionsResp, error)
	ListVpcs(regionId string, pageNumber int, pageSize int) (*models.VpcsResp, error)
	ListSubnets(zoneId string, vpcId string) (*models.SubnetsResp, error)
	ListImages(regionId string, snapshotId string, pageSize int, pageNumber int) (*models.ImagesResp, error)
	GetInstanceType(key string) string
	ListDiskCategory() []string
	ListInternetChargeType() []string
	AllocatePublicIpAddress(instanceId string) (string, error)
}

type ProviderDriverFunc func() (ProviderDriver, error)
type AccountDriverFunc func(int, string) (ProviderDriver, error)

var registeredDefaultPlugins = map[string](ProviderDriverFunc){}
var registeredAccountPlugins = map[string](AccountDriverFunc){}

func RegisterProviderDriver(name string, f ProviderDriverFunc) {
	registeredDefaultPlugins[name] = f
}

func RegisterAccountDriver(name string, f AccountDriverFunc)  {
	registeredAccountPlugins[name] = f
}

func New(name string) (ProviderDriver, error) {
	if name == "" {
		return nil, fmt.Errorf("the provider cannot be null.")
	}
	f, ok := registeredDefaultPlugins[name]
	if !ok {
		return nil, fmt.Errorf("unknown backend provider driver: %s", name)
	}
	return f()
}

func NewByAccount(bizId int, provider string) (ProviderDriver, error) {
	if provider == "" {
		return nil, fmt.Errorf("the provider cannot be null.")
	}
	f, ok := registeredAccountPlugins[provider]
	if !ok {
		return nil, fmt.Errorf("unknown backend provider driver: %s", provider)
	}
	return f(bizId, provider)
}

func ListDrivers() []string {
	drivers := make([]string, 0, len(registeredDefaultPlugins))
	for name := range registeredDefaultPlugins {
		drivers = append(drivers, name)
	}
	sort.Strings(drivers)
	return drivers
}
