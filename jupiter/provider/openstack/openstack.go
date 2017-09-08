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
package openstack

import (
	"fmt"
	"time"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
	"github.com/rackspace/gophercloud/pagination"
	"github.com/rackspace/gophercloud/openstack/compute/v2/images"
	"weibo.com/opendcp/jupiter/provider"
	"sync"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/startstop"
	"weibo.com/opendcp/jupiter/models"
	"github.com/rackspace/gophercloud/openstack/compute/v2/flavors"
	"github.com/rackspace/gophercloud/openstack/networking/v2/networks"
	"weibo.com/opendcp/jupiter/conf"
	"strconv"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/bootfromvolume"
	"github.com/rackspace/gophercloud/openstack/blockstorage/v1/volumetypes"
)

//openstack provider的实现引用了OpenStack官方的Go SDK
//具体API可查看 http://gophercloud.io/docs/compute/

type openstackProvider struct {
	client *gophercloud.ServiceClient
	lock   sync.Mutex
}

func init(){
	provider.RegisterProviderDriver("openstack", new)
}

//instanceTypesInOpenStack 用于建立flavor ID到flavor Name之间的映射
var instanceTypesInOpenStack = map[string]string{}

//本地缓存每个Flavor Id 对应的Vcpu、Ram、Disk、网络
var VcpuInOpenStack = map[string]string{}
var RamInOpenStack = map[string]string{}
var DiskInOpenStack = map[string]string{}
var networksInOpenStack = map[string]string{}


//列出所有server
func (driver openstackProvider) List(regionId string, pageNumber int, pageSize int) (*models.ListInstancesResponse, error) {
	opts1 := servers.ListOpts{}
	pager := servers.List(driver.client, opts1)
	var listInstancesResp models.ListInstancesResponse
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		serverList, _ := servers.ExtractServers(page)
		for _, instanceOP := range serverList {
			var instance models.InstanceAllIn
			instance.InstanceId = instanceOP.ID
			instance.TenantId = instanceOP.TenantID
			instance.UserId = instanceOP.UserID
			instance.Name = instanceOP.Name
			instance.Status = instanceOP.Status
			listInstancesResp.Reservations = append(listInstancesResp.Reservations, instance)
		}
		return  true, nil
	})
	return &listInstancesResp, err
}


//列出可选实例类型
func (driver openstackProvider) ListInstanceTypes() ([]string, error){

	var instanceTypesList []string
	opts := flavors.ListOpts{}
	pager := flavors.ListDetail(driver.client, opts)
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		flavorList, err := flavors.ExtractFlavors(page)
		for _, flavor := range flavorList {
			name := flavor.Name
			id := flavor.ID
			instanceType := fmt.Sprintf("%s#%s", id, name)
			instanceTypesList = append(instanceTypesList, instanceType)
			instanceTypesInOpenStack[flavor.ID] = flavor.Name
			RamInOpenStack[flavor.ID] = strconv.Itoa(flavor.RAM)
			VcpuInOpenStack[flavor.ID] = strconv.Itoa(flavor.VCPUs)
			DiskInOpenStack[flavor.ID] = strconv.Itoa(flavor.Disk)
		}
		return true, err
	})
	return instanceTypesList, err
}
//输入flavor ID，返回该flavor ID对应的配置信息
//返回的格式为"instanceType#cpu#ram"
func (driver openstackProvider) GetInstanceType(key string) string{
	instanceType := instanceTypesInOpenStack[key]
	ram := RamInOpenStack[key]
	cpu := VcpuInOpenStack[key]

	return fmt.Sprintf("%s#%s#%s",instanceType, cpu, ram)
}
func (driver openstackProvider) ListSecurityGroup(regionId string, vpcId string) (*models.SecurityGroupsResp, error){
	return nil, nil
}

func (driver openstackProvider) ListAvailabilityZones(regionId string) (*models.AvailabilityZonesResp, error){
	return nil, nil
}

func (driver openstackProvider) ListRegions() (*models.RegionsResp, error){
	return nil, nil
}

//列出所有openstack的网络。vpc指代network，其中vpcId对应networkID,vpcState对应networkName
//由于没有专门为OpenStack提供的方法和结构，故暂时这样处理
func (driver openstackProvider) ListVpcs(regionId string, pageNumber int, pageSize int) (*models.VpcsResp, error){

	url := fmt.Sprintf("http://%s:%s/v3",conf.Config.OpIp, conf.Config.OpPort)
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: url,
		Username: conf.Config.OpUserName,
		Password: conf.Config.OpPassWord,
		DomainName: "default",
	}
	provider, err := openstack.AuthenticatedClient(opts)

	if(err != nil){
		return nil, err
	}
	client, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
		Name:   "neutron",
		Region: "RegionOne",
	})
	opts1 := networks.ListOpts{}
	// Retrieve a pager (i.e. a paginated collection)
	pager := networks.List(client, opts1)

	var vpcsResp models.VpcsResp

	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		networkList, err := networks.ExtractNetworks(page)
		for _, network := range networkList {
			// "n" will be a networks.Network
			var vpc models.Vpc
			vpc.VpcId = network.ID
			vpc.State = network.Name
			vpcsResp.Vpcs = append(vpcsResp.Vpcs, vpc)
			networksInOpenStack[network.Name] = network.ID
		}
		return true, err
	})
	return &vpcsResp, err
}

func (driver openstackProvider) ListSubnets(zoneId string, vpcId string) (*models.SubnetsResp, error){
	return nil, nil
}



func (driver openstackProvider) ListDiskCategory() []string{
	return nil
}

func (driver openstackProvider) ListInternetChargeType() []string{
	return nil
}

//获取OpenStack实例对应的ID
func (driver openstackProvider) AllocatePublicIpAddress(instanceId string) (string, error){
	server, err := servers.Get(driver.client, instanceId).Extract()
	for _, address := range server.Addresses {
		tmp, ok := address.([]interface{})
		if !ok {
			return "", fmt.Errorf("get instance ip address failed!")
		}
		tmp1, ok := tmp[0].(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("get instance ip address failed!")
		}
		ip, ok := tmp1["addr"].(string)
		if !ok {
			return "", fmt.Errorf("get instance ip address failed!")
		}
		return ip, err
	}
	return "", err

}

//创建实例，如果有存储节点，则将卷作为启动盘启动实例，否则从nova启动实例
func (driver openstackProvider) Create(cluster *models.Cluster, number int) ([]string, []error) {

	//该方法列出所有存储节点的卷类型，如果访问成功则表明有存储节点
	pager := volumetypes.List(driver.client)
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		return true, nil
	})
	if err != nil{
		return driver.CreateInstanceFromVolumes(cluster, number)
	}

	return driver.CreateInstanceFromImages(cluster, number)
}

//通过卷作为启动盘创建实例
func (driver openstackProvider) CreateInstanceFromVolumes(cluster *models.Cluster, number int) ([]string, []error) {

	createdInstances := make(chan string, number)
	createdError := make(chan error, number)

	for i := 0; i < number; i++ {
		go func(i int) {

			//获取Flavor的硬盘大小，并将其作为卷大小
			diskSize, _:= strconv.Atoi( DiskInOpenStack[cluster.FlavorId])
			bd := []bootfromvolume.BlockDevice{
				bootfromvolume.BlockDevice{
					BootIndex:           0,
					DeleteOnTermination: true,
					DestinationType:     "volume",
					UUID:                cluster.ImageId,
					SourceType:          bootfromvolume.Image,
					VolumeSize:          diskSize,
				},
			}

			serverCreateOpts := servers.CreateOpts{
				Name: cluster.Name,
				FlavorRef:        cluster.FlavorId,
				AvailabilityZone: "nova",
				Networks:         []servers.Network{{UUID: cluster.Network.VpcId}},
			}
			server, err := bootfromvolume.Create(driver.client, bootfromvolume.CreateOptsExt{
				serverCreateOpts,
				bd,
			}).Extract()

			if err != nil {
				for i := 0; i < 3; i++ {
					server, err := bootfromvolume.Create(driver.client, bootfromvolume.CreateOptsExt{
						serverCreateOpts,
						bd,
					}).Extract()
					if err == nil {
						createdInstances <- server.ID
						return
					}
				}
				createdError <- err
				return
			}
			createdInstances <- server.ID
		}(i)
	}
	instanceIds := make([]string, 0)
	errs := make([]error, 0)
	for i := 0; i < number; i++ {
		select {
		case instanceId := <-createdInstances:
			instanceIds = append(instanceIds, instanceId)
		case err := <-createdError:
			errs = append(errs, err)
		}
	}
	return instanceIds, errs

}

//通过镜像创建实例
func (driver openstackProvider) CreateInstanceFromImages(cluster *models.Cluster, number int) ([]string, []error){

	createdInstances := make(chan string, number)
	createdError := make(chan error, number)

	for i := 0; i < number; i++ {
		go func(i int) {
			result, err := servers.Create(driver.client, servers.CreateOpts{
				Name:      cluster.Name ,
				ImageRef:  cluster.ImageId,
				FlavorRef: cluster.FlavorId,
				AvailabilityZone: cluster.Zone.ZoneName,
				Networks: []servers.Network{{UUID: cluster.Network.VpcId}},
			}).Extract()
			if err != nil {
				for i := 0; i < 3; i++ {
					result, err := servers.Create(driver.client, servers.CreateOpts{
						Name:      cluster.Name ,
						ImageRef:  cluster.ImageId,
						FlavorRef: cluster.FlavorId,
						AvailabilityZone: cluster.Zone.ZoneName,
						Networks: []servers.Network{{UUID: cluster.Network.VpcId}},
					}).Extract()
					if err == nil {
						createdInstances <- result.ID
						return
					}
				}
				createdError <- err
				return
			}
			createdInstances <- result.ID
		}(i)
	}
	instanceIds := make([]string, 0)
	errs := make([]error, 0)
	for i := 0; i < number; i++ {
		select {
		case instanceId := <-createdInstances:
			instanceIds = append(instanceIds, instanceId)
		case err := <-createdError:
			errs = append(errs, err)
		}
	}
	return instanceIds, errs
}


func (driver openstackProvider) GetInstance(instanceId string) (*models.Instance, error) {

	server, err := servers.Get(driver.client, instanceId).Extract()
	if err != nil {
		return nil, err
	}
	var instance models.Instance
	instance.InstanceId = server.ID
	instance.Provider = "openstack"
	instance.CreateTime, _ = time.ParseInLocation("2006-01-02 15:04:05", server.Created, time.Local)
	var ok bool
	if instance.ImageId, ok = server.Image["id"].(string); !ok {
		return nil, error("could't get instance ")
	}
	//InstanceType
	//VpcId
	//subnetId
	//SecurityGroupsId
	//私有Ip和公有Ip替换为IPV4和IPV6
	instance.Name = server.Name
	instance.TenantID = server.TenantID
	instance.UserID = server.UserID
	return &instance, err
}

//列出可选镜像列表
func (driver openstackProvider) ListImages(regionId string, snapshotId string, pageSize int, pageNumber int) (*models.ImagesResp, error) {
	opts1 := images.ListOpts{}
	pager := images.ListDetail(driver.client, opts1)
	var imageResp models.ImagesResp
	timages := make([]models.Image, 0)
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		imageList, err := images.ExtractImages(page)
		for _, imageOp := range imageList {
			image := models.Image{
				CreationDate: imageOp.Created,
				Description: imageOp.Name,
				ImageId: imageOp.ID,
				Name: imageOp.Name,
				State: imageOp.Status,

			}
			timages = append(timages, image)
		}

		return true, err
	})
	imageResp.Images = timages
	return &imageResp, err
}

func (driver openstackProvider) Start(instanceId string) (bool, error) {


	err := startstop.Start(driver.client, instanceId).ExtractErr()

	return true, err

}

func (driver openstackProvider) Stop(instanceId string) (bool, error) {


	err := startstop.Stop(driver.client, instanceId).ExtractErr()

	return true, err
}

//删除实例
func (driver openstackProvider) Delete(instanceId string) (time.Time, error) {


	server, err := servers.Get(driver.client, instanceId).Extract()

	if err != nil {
		return time.Now(), err
	}
	if server.Status != "Stopped" {
		startstop.Stop(driver.client, instanceId).ExtractErr()

		waitForSpecific(func() bool {
			server, err := servers.Get(driver.client, instanceId).Extract()
			if err != nil {
				return false
			}
			return server.Status == "Stopped"
		}, 10, 6*time.Second)
	}
	time.Sleep(5 * time.Second)
	result := servers.Delete(driver.client, instanceId)

	if result.Err != nil {
		return time.Now(), result.Err
	}
	return time.Now(), nil
}

func (driver openstackProvider) WaitForInstanceToStop(instanceId string) bool {
	st, err := driver.GetState(instanceId)
	if err != nil {
		return false
	}
	return st == models.Stopped
}

func (driver openstackProvider) WaitToStartInstance(instanceId string) bool {
	st, err := driver.GetState(instanceId)
	if err != nil{
		return false
	}
	return st == models.Running
}

func (driver openstackProvider) GetState(instanceId string) (models.InstanceState, error) {

	server, err := servers.Get(driver.client, instanceId).Extract()
	if err != nil {
		return models.StateError, err
	}
	switch server.Status {
	case "ACTIVE":
		return models.Running, nil
	case "BUILD":
		return models.Starting, nil
	case "STOPPED":
		return models.Stopped, nil
	case "PAUSED":
		return models.Stopping, nil
	default:
		return models.None, nil
	}
}

func waitForSpecific(f func() bool, maxAttempts int, waitInterval time.Duration) error {
	for i := 0; i < maxAttempts; i++ {
		if f() {
			return nil
		}
		time.Sleep(waitInterval)
	}
	return fmt.Errorf("Maximum number of retries (%d) exceeded", maxAttempts)
}



func new() (provider.ProviderDriver, error){

	return newProvider()
}

func newProvider() (provider.ProviderDriver, error){

	url := fmt.Sprintf("http://%s:%s/v3",conf.Config.OpIp, conf.Config.OpPort)
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: url,
		Username: conf.Config.OpUserName,
		Password: conf.Config.OpPassWord,
		DomainName: "default",
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil{
		return nil, err
	}
	client, err :=
		openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
			Region: "RegionOne",
		})

	ret := openstackProvider{
		client: client,
	}
	return ret, err
}


















