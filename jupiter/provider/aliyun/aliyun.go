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



package aliyun

import (
	"strconv"
	"strings"
	"sync"
	"time"

	"fmt"
	"github.com/jiangshengwu/aliyun-sdk-for-go/ecs"
	"weibo.com/opendcp/jupiter/conf"
	"weibo.com/opendcp/jupiter/models"
	"weibo.com/opendcp/jupiter/provider"
	"errors"
)

func init() {
	provider.RegisterProviderDriver("aliyun", new)
}

const (
	CN_BEIJING_C = "cn-beijing-c"
	IO_OPTIMIZED = "optimized"
	TIMES_DEL 	 = 2
)

var instanceTypesInAliyun = map[string]string{
	"1Core-1GB":    "ecs.n1.tiny",
	"1Core-2GB":    "ecs.n1.small",
	"4Cores-8GB":   "ecs.n1.large",
	"16Cores-16GB": "ecs.c2.medium",
	"16Cores-64GB": "ecs.n2.3xlarge",
}

var dataCategoryInAliyun = []string{
	"cloud_efficiency",
	"cloud_ssd",
}

var internetChargeTypeInAliyun = []string{
	"PayByBandwidth",
	"PayByTraffic",
}

type aliyunProvider struct {
	client *ecs.EcsClient
	lock   sync.Mutex
}

func (driver aliyunProvider) ListDiskCategory() []string {
	return dataCategoryInAliyun
}

func (driver aliyunProvider) ListInternetChargeType() []string {
	return internetChargeTypeInAliyun
}

func (driver aliyunProvider) GetInstanceType(key string) string {
	return instanceTypesInAliyun[key]
}

func (driver aliyunProvider) Create(cluster *models.Cluster, number int) ([]string, []error) {
	createdInstances := make(chan string, number)
	createdError := make(chan error, number)
	for i := 0; i < number; i++ {
		go func(i int) {
			params := buildCreateRequest(cluster)
			result, err := driver.client.Instance.CreateInstance(params)
			if err != nil {
				for i := 0; i < 3; i++ {
					delete(params, "Signature")
					result, err = driver.client.Instance.CreateInstance(params)
					if err == nil {
						createdInstances <- result.InstanceId
						return
					}
				}
				createdError <- err
				return
			}
			createdInstances <- result.InstanceId
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

func buildCreateRequest(input *models.Cluster) map[string]interface{} {
	params := make(map[string]interface{})
	params["RegionId"] = input.Zone.RegionName
	params["ZoneId"] = input.Zone.ZoneName
	params["ImageId"] = input.ImageId
	params["InstanceType"] = input.InstanceType
	params["SecurityGroupId"] = input.Network.SecurityGroup
	params["Password"] = conf.Config.Password
	params["SystemDisk.Category"] = input.SystemDiskCategory
	for i := 1; i <= input.DataDiskNum; i++ {
		params["DataDisk."+strconv.Itoa(i)+".Size"] = strconv.Itoa(input.DataDiskSize)
		params["DataDisk."+strconv.Itoa(i)+".Category"] = input.DataDiskCategory
	}
	if strings.EqualFold(input.Zone.ZoneName, CN_BEIJING_C) {
		params["IoOptimized"] = IO_OPTIMIZED
	}
	if len(input.Network.VpcId) > 0 {
		params["VSwitchId"] = input.Network.SubnetId
	}
	if len(input.Network.VpcId) <= 0 {
		params["InternetChargeType"] = input.Network.InternetChargeType
		params["InternetMaxBandwidthOut"] = strconv.Itoa(input.Network.InternetMaxBandwidthOut)
	}
	return params
}

//判断包年包月的数字是否是1 - 9，12，24，36
func checkPreBuyMonth(month int) bool {
	return (month > 0 && month < 10) || month == 12 || month == 24 || month == 36
}

func (driver aliyunProvider) GetInstance(instanceId string) (*models.Instance, error) {
	insAttr, err := driver.client.Instance.DescribeInstanceAttribute(map[string]interface{}{
		"InstanceId": instanceId,
	})
	if err != nil {
		return nil, err
	}
	var instance models.Instance
	instance.InstanceId = insAttr.InstanceId
	instance.Provider = "aliyun"
	instance.CreateTime, _ = time.ParseInLocation("2006-01-02 15:04:05", insAttr.CreationTime, time.Local)
	//instance.Cpu = insAttr.CPU
	//instance.Ram = insAttr.Memory
	instance.ImageId = insAttr.ImageId
	instance.InstanceType = insAttr.InstanceType
	instance.VpcId = insAttr.VpcAttributes.VpcId
	instance.SubnetId = insAttr.VpcAttributes.VSwitchId
	instance.SecurityGroupId = strings.Join(insAttr.AllSecurityGroupIds.AllSecurityGroupId, ",")
	instance.PrivateIpAddress = strings.Join(insAttr.VpcAttributes.PrivateIpAddress.AllIpAddress, ",")
	instance.PublicIpAddress = strings.Join(insAttr.PublicIpAddress.AllIpAddress, ",")
	instance.RegionId = insAttr.RegionId
	instance.ZoneId = insAttr.ZoneId
	instance.NatIpAddress = insAttr.VpcAttributes.NatIpAddress
	instance.CostWay = insAttr.InstanceChargeType
	return &instance, nil
}

func (driver aliyunProvider) WaitForInstanceToStop(instanceId string) bool {
	st, err := driver.GetState(instanceId)
	if err != nil {
		return false
	}
	return st == models.Stopped
}

func (driver aliyunProvider) WaitToStartInstance(instanceId string) bool {
	st, err := driver.GetState(instanceId)
	if err != nil {
		return false
	}
	return st == models.Running
}

func (driver aliyunProvider) GetState(instanceId string) (models.InstanceState, error) {
	statusResp, err := driver.client.Instance.DescribeInstanceAttribute(map[string]interface{}{
		"InstanceId": instanceId,
	})
	if err != nil {
		return models.StateError, err
	}
	switch statusResp.Status {
	case "Running":
		return models.Running, nil
	case "Starting":
		return models.Starting, nil
	case "Stopped":
		return models.Stopped, nil
	case "Stopping":
		return models.Stopping, nil
	default:
		return models.None, nil
	}
}

func (driver aliyunProvider) Delete(instanceId string) (time.Time, error) {
	result, err := driver.client.Instance.DescribeInstanceAttribute(map[string]interface{}{
		"InstanceId": instanceId,
	})
	if err != nil {
		return time.Now(), err
	}

	times := 1
	if result.Status != "Stopped" {
		_, _ = driver.client.Instance.StopInstance(map[string]interface{}{
			"InstanceId": instanceId,
			"ForceStop":  "true",
		})
		waitForSpecific(func() bool {
			times++
			status, err := driver.client.Instance.DescribeInstanceAttribute(map[string]interface{}{
				"InstanceId": instanceId,
			})
			if err != nil {
				return false
			}
			return status.Status == "Stopped"
		}, 20, 6*time.Second)
		if times >= 20 {
			return time.Now(), errors.New("It already has timed out to wait instance to stop")
		}
	}

	for i:=1; i<=TIMES_DEL; i++ {
		_, err = driver.client.Instance.DeleteInstance(map[string]interface{}{
			"InstanceId": instanceId,
		})
		if err == nil {     		//删除成功返回
			return time.Now(), nil
		} else {
			if i == TIMES_DEL {    //重试删除TIMES_DEL次后失败
				msg := fmt.Sprintf("Retry to delete instance %d times failed, err: %s", TIMES_DEL, err.Error())
				return time.Now(), errors.New(msg)
			}
		}
	}

	return time.Now(), err
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

func (driver aliyunProvider) Start(instanceId string) (bool, error) {
	params := map[string]interface{}{
		"InstanceId": instanceId,
	}
	_, err := driver.client.Instance.StartInstance(params)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (driver aliyunProvider) Stop(instanceId string) (bool, error) {
	params := map[string]interface{}{
		"InstanceId": instanceId,
	}
	_, err := driver.client.Instance.StartInstance(params)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (driver aliyunProvider) List(regionId string, pageNumber int, pageSize int) (*models.ListInstancesResponse, error) {
	resp, err := driver.client.Instance.DescribeInstances(map[string]interface{}{
		"RegionId":   regionId,
		"PageNumber": strconv.Itoa(pageNumber),
		"PageSize":   strconv.Itoa(pageSize),
	})
	if err != nil {
		return nil, err
	}
	var listInstancesResp models.ListInstancesResponse
	for _, instanceAli := range resp.AllInstances.AllInstance {
		var instance models.InstanceAllIn
		instance.InstanceId = instanceAli.InstanceId
		instance.InstanceName = instanceAli.InstanceName
		instance.Description = instanceAli.Description
		instance.ImageId = instanceAli.ImageId
		instance.RegionId = instanceAli.RegionId
		instance.ZoneId = instanceAli.ZoneId
		instance.InstanceType = instanceAli.InstanceType
		instance.Status = instanceAli.Status
		instance.CreationTime = instanceAli.CreationTime
		instance.ExpiredTime = instanceAli.ExpiredTime
		instance.AllSecurityGroupIds.AllSecurityGroupId = instanceAli.AllSecurityGroupIds.AllSecurityGroupId
		instance.PublicIpAddress.AllIpAddress = instanceAli.PublicIpAddress.AllIpAddress
		instance.InnerIpAddress.AllIpAddress = instanceAli.InnerIpAddress.AllIpAddress
		instance.VpcAttributes.NatIpAddress = instanceAli.VpcAttributes.NatIpAddress
		instance.VpcAttributes.PrivateIpAddress.AllIpAddress = instanceAli.VpcAttributes.PrivateIpAddress.AllIpAddress
		instance.VpcAttributes.VpcId = instanceAli.VpcAttributes.VpcId
		instance.VpcAttributes.VSwitchId = instanceAli.VpcAttributes.VSwitchId
		instance.EipAddress.AllocationId = instanceAli.EipAddress.AllocationId
		instance.EipAddress.Bandwidth = instanceAli.EipAddress.Bandwidth
		instance.EipAddress.InternetChargeType = instanceAli.EipAddress.InternetChargeType
		instance.EipAddress.IpAddress = instanceAli.EipAddress.IpAddress
		listInstancesResp.Reservations = append(listInstancesResp.Reservations, instance)
	}
	return &listInstancesResp, nil
}

func (aliyunProvider) ListInstanceTypes() ([]string, error) {
	instanceTypes := make([]string, 0, len(instanceTypesInAliyun))
	for instanceType := range instanceTypesInAliyun {
		instanceTypes = append(instanceTypes, instanceType)
	}
	return instanceTypes, nil
}

func (aliyunProvider) CreateSecurityGroup(input *models.CreateSecurityGroupParam) (*models.SecurityGroupResp, error) {
	return nil, nil
}

func (driver aliyunProvider) ListSecurityGroup(regionId string, vpcId string) (*models.SecurityGroupsResp, error) {
	securityGroups := make([]ecs.SecurityGroupItemType, 0)
	i := 1
	pageSize := 50
	resp, err := driver.client.SecurityGroup.DescribeSecurityGroups(map[string]interface{}{
		"RegionId":   regionId,
		"VpcId":      vpcId,
		"PageNumber": strconv.Itoa(i),
		"PageSize":   strconv.Itoa(pageSize),
	})

	if err != nil {
		return nil, err
	}
	totalCount := resp.TotalCount        //获取返回的列表总条目数
	pageCount := totalCount/pageSize + 1 //算出实际页数
	for j := 0; j < pageCount; j++ {
		resp, err := driver.client.SecurityGroup.DescribeSecurityGroups(map[string]interface{}{
			"RegionId":   regionId,
			"VpcId":      vpcId,
			"PageNumber": strconv.Itoa(j + 1),
			"PageSize":   strconv.Itoa(pageSize),
		})
		if err != nil {
			return nil, err
		}

		groupCount := len(resp.AllGroups.AllGroup)
		for i := 0; i < groupCount; i++ {
			securityGroups = append(securityGroups, resp.AllGroups.AllGroup[i])
		}
	}
	var securityGR models.SecurityGroupsResp
	for _, securityGroupsAli := range securityGroups {
		var securityGroup models.SecurityGroup
		securityGroup.GroupId = securityGroupsAli.SecurityGroupId
		securityGroup.GroupName = securityGroupsAli.SecurityGroupName
		securityGroup.Description = securityGroupsAli.Description
		securityGroup.VpcId = securityGroupsAli.VpcId
		securityGR.SecurityGroups = append(securityGR.SecurityGroups, securityGroup)
	}
	return &securityGR, nil
}
func (aliyunProvider) AuthorizeSecurityGroupIngress(input *models.AuthorizeSecurityGroupIngress) (bool, error) {
	return true, nil
}
func (driver aliyunProvider) ListAvailabilityZones(regionId string) (*models.AvailabilityZonesResp, error) {
	result, err := driver.client.Region.DescribeZones(map[string]interface{}{
		"RegionId": regionId,
	})
	if err != nil {
		return nil, err
	}
	var zoneresp models.AvailabilityZonesResp
	for _, zoneAli := range result.AllZones.AllZone {
		var zone models.AvailabilityZone
		zone.ZoneName = zoneAli.ZoneId
		zone.RegionName = zoneAli.LocalName
		zoneresp.AvailabilityZones = append(zoneresp.AvailabilityZones, zone)
	}
	return &zoneresp, nil
}
func (driver aliyunProvider) ListRegions() (*models.RegionsResp, error) {
	result, err := driver.client.Region.DescribeRegions(map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	var regionresp models.RegionsResp
	for _, regionAli := range result.AllRegions.AllRegion {
		var region models.Region
		region.RegionName = regionAli.RegionId
		regionresp.Regions = append(regionresp.Regions, region)
	}
	return &regionresp, nil
}
func (driver aliyunProvider) ListVpcs(regionId string, pageNumber int, pageSize int) (*models.VpcsResp, error) {
	vpcs := make([]ecs.VpcSetType, 0)
	resp, err := driver.client.Vpc.DescribeVpcs(map[string]interface{}{
		"RegionId":   regionId,
		"PageNumber": pageNumber,
		"PageSize":   pageSize,
	})
	if err != nil {
		return nil, err
	}
	totalCount := resp.TotalCount
	pageCount := totalCount/pageSize + 1 //算出实际页数
	for j := 0; j < pageCount; j++ {
		resp, err := driver.client.Vpc.DescribeVpcs(map[string]interface{}{
			"RegionId":   regionId,
			"PageNumber": strconv.Itoa(j + 1),
			"PageSize":   strconv.Itoa(pageSize),
		})
		if err != nil {
			return nil, err
		}
		VpcCount := len(resp.Vpcs.Vpc) //vpc个数
		for k := 0; k < VpcCount; k++ {
			vpcs = append(vpcs, resp.Vpcs.Vpc[k])
		}
	}
	var vpcsResp models.VpcsResp

	for _, vpcAli := range vpcs {
		var vpc models.Vpc
		vpc.VpcId = vpcAli.VpcId
		vpc.CidrBlock = vpcAli.CidrBlock
		vpc.State = vpcAli.Status
		vpcsResp.Vpcs = append(vpcsResp.Vpcs, vpc)
	}
	return &vpcsResp, nil
}
func (driver aliyunProvider) ListSubnets(zoneId string, vpcId string) (*models.SubnetsResp, error) {
	vswitches := make([]ecs.VSwitchType, 0)
	i := 1
	pageSize := 50
	resp, err := driver.client.VSwitch.DescribeVSwitches(map[string]interface{}{
		"ZoneId":     zoneId,
		"VpcId":      vpcId,
		"PageNumber": strconv.Itoa(i),
		"PageSize":   strconv.Itoa(pageSize),
	})

	if err != nil {
		return nil, err
	}
	totalCount := resp.TotalCount        //获取返回的列表总条目数
	pageCount := totalCount/pageSize + 1 //算出实际页数
	for j := 0; j < pageCount; j++ {
		resp, err := driver.client.VSwitch.DescribeVSwitches(map[string]interface{}{
			"ZoneId":     zoneId,
			"VpcId":      vpcId,
			"PageNumber": strconv.Itoa(j + 1),
			"PageSize":   strconv.Itoa(pageSize),
		})
		if err != nil {
			return nil, err
		}
		switchCount := len(resp.AllVSwitches.AllVSwitch) //每个vpc中交换机个数
		for p := 0; p < switchCount; p++ {
			vswitches = append(vswitches, resp.AllVSwitches.AllVSwitch[p])
		}
	}
	var subnetsResp models.SubnetsResp
	for _, subnetAli := range vswitches {
		var subnet models.Subnet
		subnet.State = subnetAli.Status
		subnet.CidrBlock = subnetAli.CidrBlock
		subnet.SubnetId = subnetAli.VSwitchId
		subnet.AvailabilityZone = subnetAli.ZoneId
		subnet.VpcId = subnetAli.VpcId
		subnetsResp.Subnets = append(subnetsResp.Subnets, subnet)
	}
	return &subnetsResp, nil
}
func (aliyunProvider) CreateSubnet(input *models.CreateSubnetParam) (*models.SubnetResp, error) {
	return nil, nil
}
func (driver aliyunProvider) ListImages(regionId string, snapshotId string, pageSize int, pageNumber int) (*models.ImagesResp, error) {
	result, err := driver.client.Image.DescribeImages(map[string]interface{}{
		"RegionId":   regionId,
		"SnapshotId": snapshotId,
		"PageSize":   pageSize,
		"PageNumber": strconv.Itoa(pageNumber),
	})
	if err != nil {
		return nil, err
	}
	var imageResp models.ImagesResp
	timages := make([]models.Image, 0)
	for _, imageAli := range result.AllImages.AllImage {
		image := models.Image{
			Architecture: imageAli.Architecture,
			CreationDate: imageAli.CreationTime,
			Description:  imageAli.Description,
			ImageId:      imageAli.ImageId,
			Name:         imageAli.ImageName,
			OwnerId:      imageAli.ImageOwnerAlias,
			ProductCodes: []models.ProductCode{models.ProductCode{ProductCodeType: imageAli.ProductCode,
				ProductCodeId: ""}},
			State: imageAli.Status,
		}
		timages = append(timages, image)
	}
	imageResp.Images = timages
	return &imageResp, nil
}

func (driver aliyunProvider) AllocatePublicIpAddress(instanceId string) (string, error) {
	publicIpAddress, err := driver.client.Network.AllocatePublicIpAddress(map[string]interface{}{
		"InstanceId": instanceId,
	})
	if err != nil {
		return "", err
	}
	return publicIpAddress.IpAddress, nil
}

func (aliyunProvider) CreateVpc(input *models.CreateVpcParam) (*models.VpcResp, error) {
	return nil, nil
}
func (aliyunProvider) CreateGateway() (*models.GatewayResp, error) {
	return nil, nil
}
func (aliyunProvider) AttachGateway(input *models.AttachGateway) (bool, error) {
	return true, nil
}

func new() (provider.ProviderDriver, error) {
	return newProvider()
}

func newProvider() (provider.ProviderDriver, error) {
	client := ecs.NewClient(
		conf.Config.KeyId,
		conf.Config.KeySecret,
		"",
	)
	ret := aliyunProvider{
		client: client,
	}
	return ret, nil
}
