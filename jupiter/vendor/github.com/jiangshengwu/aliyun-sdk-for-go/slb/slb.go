// SLB API package
package slb

import (
	"github.com/jiangshengwu/aliyun-sdk-for-go/util"
)

const (
	// SLB API Host
	SLBHost string = "https://slb.aliyuncs.com/?"

	// All SLB APIs only support GET method
	SLBHttpMethod = "GET"

	// SDK only supports JSON format
	Format = "JSON"

	Version          = "2014-05-15"
	SignatureMethod  = "HMAC-SHA1"
	SignatureVersion = "1.0"
)

// struct for SLB client
type SlbClient struct {
	util.Client

	// Access to API call from this client
	LoadBalancer      LoadBalancerService
	ServerCertificate ServerCertificateService
	Listener          ListenerService
	BackendServer     BackendServerService
}

// Initialize an SLB client
func NewClient(accessKeyId string, accessKeySecret string, resourceOwnerAccount string) *SlbClient {
	client := &SlbClient{}
	client.Common = &util.CommonParam{}
	client.Common.AccessKeyId = accessKeyId
	client.Common.AccessKeySecret = accessKeySecret
	client.Common.ResourceOwnerAccount = resourceOwnerAccount
	client.Common.Host = SLBHost
	ps := map[string]interface{}{
		"HttpMethod":       SLBHttpMethod,
		"Format":           Format,
		"Version":          Version,
		"AccessKeyId":      accessKeyId,
		"SignatureMethod":  SignatureMethod,
		"SignatureVersion": SignatureVersion,
	}
	client.Common.Attr = ps

	client.LoadBalancer = &LoadBalancerOperator{client.Common}
	client.ServerCertificate = &ServerCertificateOperator{client.Common}
	client.Listener = &ListenerOperator{client.Common}
	client.BackendServer = &BackendServerOperator{client.Common}
	return client
}

func (client *SlbClient) GetClientName() string {
	return "SLB Client"
}

func (client *SlbClient) GetVersion() string {
	return client.Common.Attr["Version"].(string)
}

func (client *SlbClient) GetSignatureMethod() string {
	return client.Common.Attr["SignatureMethod"].(string)
}

func (client *SlbClient) GetSignatureVersion() string {
	return client.Common.Attr["SignatureVersion"].(string)
}
