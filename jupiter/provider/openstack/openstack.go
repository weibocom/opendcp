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
)

//1.由于接口完全是阿里云的接口，已经实现的函数无法实现相应功能
//2.无法实现功能的方法如何处理

type openstackProvider struct {
	client *gophercloud.ServiceClient
	lock   sync.Mutex
}

func init(){
	provider.RegisterProviderDriver("openstack", new)
	fmt.Println("openstack init() execute")
}

//列出所有server
//openstack不需要提供pageNumber和pageSize,该如何处理
//返回的示例中包含所有信息，之后根据需要进行适当的删减
//要求：搞清楚前端调用时到底需要哪些参数，以什么顺序排列
func (driver openstackProvider) List(regionId string, pageNumber int, pageSize int) (*models.ListInstancesResponse, error) {
	opts1 := servers.ListOpts{}
	pager := servers.List(driver.client, opts1)
	var listInstancesResp models.ListInstancesResponse
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		serverList, _ := servers.ExtractServers(page)
		for _, instanceOP := range serverList {
			var instance models.InstanceAllIn
			instance.InstanceId = instanceOP.ID
			instance.TenantID = instanceOP.TenantID
			instance.UserID = instanceOP.UserID
			instance.Name = instanceOP.Name
			instance.Updated = instanceOP.Updated
			instance.Created = instanceOP.Created
			instance.HostID = instanceOP.HostID
			instance.Status = instanceOP.Status
			instance.Progress = instanceOP.Progress
			instance.AccessIPv4 = instanceOP.AccessIPv4
			instance.AccessIPv6 = instanceOP.AccessIPv6
			//instance.Image = instanceOP.Image
			//instance.Flavor = instanceOP.Flavor
			//instance.Addresses = instanceOP.Addresses
			//instance.Metadata = instanceOP.Metadata
			//instance.Links = instanceOP.Links
			//instance.KeyName = instanceOP.KeyName
			//instance.AdminPass = instanceOP.AdminPass
			//instance.SecurityGroups = instanceOP.SecurityGroups
			listInstancesResp.Reservations = append(listInstancesResp.Reservations, instance)
		}
		return  true, nil
	})



	return &listInstancesResp, err
}

//创建实例代码待做
//func (driver openstackProvider) Create(cluster *models.Cluster, number int) ([]string, []error) {
//	client, err :=
//		openstack.NewComputeV2(driver, gophercloud.EndpointOpts{
//			Region: cluster.Zone.Id,
//		})
//	if err != nil {
//		return nil, err
//	}
//	createdInstances := make(chan string, number)
//	createdError := make(chan error, number)
//	for i := 0; i < number; i++ {
//		go func(i int) {
//			server, err := servers.Create(client, servers.CreateOpts{
//				Name:      "My new server!",
//				ImageRef:  "d96e4977-3e0b-4a39-afe7-4641b5e63b3d",
//				FlavorRef: "7c307b7f-4a1e-4e4e-8a42-a36b1ac3c5f5",
//				AvailabilityZone: "nova:75-29-208-yf-core.jpool.sinaimg.cn",
//				Networks: []servers.Network{{UUID : "e9634c8b-0e14-4c2f-83ec-43bd45689f8a"}},
//
//
//
//			}).Extract()
//			if err != nil {
//				for i := 0; i < 3; i++ {
//					delete(params, "Signature")
//					result, err = driver.client.Instance.CreateInstance(params)
//					if err == nil {
//						createdInstances <- result.InstanceId
//						return
//					}
//				}
//				createdError <- err
//				return
//			}
//			createdInstances <- result.InstanceId
//		}(i)
//	}
//	instanceIds := make([]string, 0)
//	errs := make([]error, 0)
//	for i := 0; i < number; i++ {
//		select {
//		case instanceId := <-createdInstances:
//			instanceIds = append(instanceIds, instanceId)
//		case err := <-createdError:
//			errs = append(errs, err)
//		}
//	}
//	return instanceIds, errs
//}

//func buildCreateRequest(input *models.Cluster) map[string]interface{} {
//	params := make(map[string]interface{})
//	params["RegionId"] = input.Zone.RegionName
//	params["ZoneId"] = input.Zone.ZoneName
//	params["ImageId"] = input.ImageId
//	params["InstanceType"] = input.InstanceType
//	params["SecurityGroupId"] = input.Network.SecurityGroup
//	params["Password"] = conf.Config.Password
//	params["SystemDisk.Category"] = input.SystemDiskCategory
//	for i := 1; i <= input.DataDiskNum; i++ {
//		params["DataDisk."+strconv.Itoa(i)+".Size"] = strconv.Itoa(input.DataDiskSize)
//		params["DataDisk."+strconv.Itoa(i)+".Category"] = input.DataDiskCategory
//	}
//	if strings.EqualFold(input.Zone.ZoneName, CN_BEIJING_C) {
//		params["IoOptimized"] = IO_OPTIMIZED
//	}
//	if len(input.Network.VpcId) > 0 {
//		params["VSwitchId"] = input.Network.SubnetId
//	}
//	if len(input.Network.VpcId) <= 0 {
//		params["InternetChargeType"] = input.Network.InternetChargeType
//		params["InternetMaxBandwidthOut"] = strconv.Itoa(input.Network.InternetMaxBandwidthOut)
//	}
//	return params
//}


func (driver openstackProvider) GetInstance(instanceId string) (*models.Instance, error) {

	server, err := servers.Get(driver.client, instanceId).Extract()
	if err != nil {
		return nil, err
	}
	var instance models.Instance

	instance.InstanceId = server.ID
	instance.Provider = "openstack"
	instance.CreateTime, _ = time.ParseInLocation("2006-01-02 15:04:05", server.Created, time.Local)
	tmp := server.Image["id"]
	instance.ImageId = tmp.(*string)
	//InstanceType
	//VpcId
	//subnetId
	//SecurityGroupsId
	//私有Ip和公有Ip替换为IPV4和IPV6
	instance.AccessIPv4 = server.AccessIPv4
	instance.AccessIPv6 = server.AccessIPv6
	instance.Name = server.Name
	instance.TenantID = server.TenantID
	instance.UserID = server.UserID
	return &instance, err
}

//列出镜像列表
//这里使用的镜像是阿里云的镜像，之后根据情况添加openstack镜像的相关参数
func (driver openstackProvider) ListImages(regionId string, snapshotId string, pageSize int, pageNumber int) (*models.ImagesResp, error) {


	opts1 := images.ListOpts{}
	pager := images.ListDetail(driver.client, opts1)
	var imageResp models.ImagesResp
	timages := make([]models.Image, 0)
	pager.EachPage(func(page pagination.Page) (bool, error) {
		imageList, err := images.ExtractImages(page)
		for _, imageOp := range imageList {
			image := models.Image{
				//Architecture: imageOp.
				CreationDate: imageOp.Created,
				//Description: imageOp.
				ImageId: imageOp.ID,
				Name: imageOp.Name,
				//OwnerId: imageOp.
				//ProductCodes
				State: imageOp.Status,

			}
			timages = append(timages, image)
		}

		return true, err
	})
	imageResp.Images = timages
	return &imageResp, nil
}

func (driver openstackProvider) Start(instanceId string) (bool, error) {


	err := startstop.Start(driver.client, instanceId).ExtractErr()

	return true, err
}

func (driver openstackProvider) Stop(instanceId string) (bool, error) {


	err1 := startstop.Stop(driver.client, instanceId).ExtractErr()

	return true, err1
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
		return time.Now(), err
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
	if err != nil {
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
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: "http://10.39.59.27:5000/v3/auth",
		Username: "admin",
		Password: "ZYGL32NDG7JS8IGC",
		DomainName: "default",
	}

	provider, err := openstack.AuthenticatedClient(opts)

	client, err :=
		openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
			Region: "RegionOne",
		})

	ret := openstackProvider{
		client: client,
	}
	return ret, err
}


















