// ECS API package
package ecs

import (
	"github.com/jiangshengwu/aliyun-sdk-for-go/util"
)

const (
	// ECS API Host
	ECSHost string = "https://ecs.aliyuncs.com/?"

	// All ECS APIs only support GET method
	ECSHttpMethod = "GET"

	// SDK only supports JSON format
	Format = "JSON"

	Version          = "2014-05-26"
	SignatureMethod  = "HMAC-SHA1"
	SignatureVersion = "1.0"
)

// struct for ECS client
type EcsClient struct {
	util.Client

	// Access to API call from this client
	Region        RegionService
	SecurityGroup SecurityGroupService
	Instance      InstanceService
	Other         OtherService
	Image         ImageService
	Snapshot      SnapshotService
	Disk          DiskService
	Network       NetworkService
	Monitor       MonitorService
	Vpc           VpcService
	VRouter       VRouterService
	VSwitch       VSwitchService
	Route         RouteService
}

// Initialize an ECS client
func NewClient(accessKeyId string, accessKeySecret string, resourceOwnerAccount string) *EcsClient {
	client := &EcsClient{}
	client.Common = &util.CommonParam{}
	client.Common.AccessKeyId = accessKeyId
	client.Common.AccessKeySecret = accessKeySecret
	client.Common.ResourceOwnerAccount = resourceOwnerAccount
	client.Common.Host = ECSHost
	ps := map[string]interface{}{
		"HttpMethod":       ECSHttpMethod,
		"Format":           Format,
		"Version":          Version,
		"AccessKeyId":      accessKeyId,
		"SignatureMethod":  SignatureMethod,
		"SignatureVersion": SignatureVersion,
	}
	client.Common.Attr = ps

	client.Region = &RegionOperator{client.Common}
	client.SecurityGroup = &SecurityGroupOperator{client.Common}
	client.Instance = &InstanceOperator{client.Common}
	client.Other = &OtherOperator{client.Common}
	client.Image = &ImageOperator{client.Common}
	client.Snapshot = &SnapshotOperator{client.Common}
	client.Disk = &DiskOperator{client.Common}
	client.Network = &NetworkOperator{client.Common}
	client.Monitor = &MonitorOperator{client.Common}
	client.Vpc = &VpcOperator{client.Common}
	client.VRouter = &VRouterOperator{client.Common}
	client.VSwitch = &VSwitchOperator{client.Common}
	client.Route = &RouteOperator{client.Common}
	return client
}

func (client *EcsClient) GetClientName() string {
	return "ECS Client"
}

func (client *EcsClient) GetVersion() string {
	return client.Common.Attr["Version"].(string)
}

func (client *EcsClient) GetSignatureMethod() string {
	return client.Common.Attr["SignatureMethod"].(string)
}

func (client *EcsClient) GetSignatureVersion() string {
	return client.Common.Attr["SignatureVersion"].(string)
}
