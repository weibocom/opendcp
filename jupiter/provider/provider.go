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

var registeredPlugins = map[string](ProviderDriverFunc){}

func RegisterProviderDriver(name string, f ProviderDriverFunc) {
	registeredPlugins[name] = f
}

func New(name string) (ProviderDriver, error) {
	if name == "" {
		return nil, fmt.Errorf("the provider cannot be null.")
	}
	f, ok := registeredPlugins[name]
	if !ok {
		return nil, fmt.Errorf("unknown backend provider driver: %s", name)
	}
	return f()
}

func ListDrivers() []string {
	drivers := make([]string, 0, len(registeredPlugins))
	for name := range registeredPlugins {
		drivers = append(drivers, name)
	}
	sort.Strings(drivers)
	return drivers
}
