package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:ClusterController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:ClusterController"],
		beego.ControllerComments{
			Method: "GetClusters",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:ClusterController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:ClusterController"],
		beego.ControllerComments{
			Method: "GetClusterInfo",
			Router: `/:clusterId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:ClusterController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:ClusterController"],
		beego.ControllerComments{
			Method: "CreateCluster",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:ClusterController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:ClusterController"],
		beego.ControllerComments{
			Method: "DeleteCluster",
			Router: `/:clusterId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:ClusterController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:ClusterController"],
		beego.ControllerComments{
			Method: "ExpandInstances",
			Router: `/:clusterId/expand/:number`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:ClusterController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:ClusterController"],
		beego.ControllerComments{
			Method: "GetInstancesNumber",
			Router: `/number/:hour`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:ClusterController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:ClusterController"],
		beego.ControllerComments{
			Method: "GetPastInstancesNumber",
			Router: `/oldnumber/:time`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:ClusterController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:ClusterController"],
		beego.ControllerComments{
			Method: "UpdateInstanceInfo",
			Router: `/update`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:CredentialController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:CredentialController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:CredentialController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:CredentialController"],
		beego.ControllerComments{
			Method: "Authorize",
			Router: `/authorization`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "CreateInstance",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "StartInstance",
			Router: `/start/:instanceId`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "GetInstance",
			Router: `/:instanceId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "GetInstancesStatus",
			Router: `status/:instanceIds`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "UpdateInstanceStatus",
			Router: `/status`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "DeleteMulti",
			Router: `/:instanceIds`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "DownloadKey",
			Router: `sshkey/:ip`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "UploadKey",
			Router: `sshkey/:instanceId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "GetProviders",
			Router: `/provider`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "GetRegionIds",
			Router: `/regions/:provider`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "GetZones",
			Router: `/zones/:provider/:regionId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "GetVpcs",
			Router: `/vpc/:provider/:regionId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "GetSubnets",
			Router: `/subnet/:provider/:zoneId/:vpcId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "ListInstanceTypes",
			Router: `/type/:provider`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "ListInternetChargeType",
			Router: `/charge/:provider`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "ListDiskCategory",
			Router: `/disk_category/:provider`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "GetImages",
			Router: `/image/:provider/:regionId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "GetSecurityGroup",
			Router: `/security_group/:provider/:regionId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "ListAllInstances",
			Router: `/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "ListInstancesByClusterId",
			Router: `/cluster/:clusterId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "QueryLogByCorrelationIdAndInstanceId",
			Router: `/log/:correlationId/:instanceId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "QueryLogByInstanceId",
			Router: `/log/:instanceId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "UploadPhyDevInfo",
			Router: `/phydev`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:InstanceController"],
		beego.ControllerComments{
			Method: "ManagePhyDev",
			Router: `/phydev`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"],
		beego.ControllerComments{
			Method: "ListAllOrganizations",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"],
		beego.ControllerComments{
			Method: "GetOrganizationById",
			Router: `/:organizationId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"],
		beego.ControllerComments{
			Method: "GetInstancesByOrganizationId",
			Router: `/instances/:organizationId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"],
		beego.ControllerComments{
			Method: "CreateOrganization",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"],
		beego.ControllerComments{
			Method: "DeleteOrganization",
			Router: `/:organizationId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"],
		beego.ControllerComments{
			Method: "GetClustersByOrganizationId",
			Router: `/cluster/:organizationId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"],
		beego.ControllerComments{
			Method: "GetUsage",
			Router: `/usage/:clusterId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"],
		beego.ControllerComments{
			Method: "IncreaseCredit",
			Router: `/credit/:clusterId/:hours`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"],
		beego.ControllerComments{
			Method: "GetCredit",
			Router: `/credit/:clusterId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:OrganizationController"],
		beego.ControllerComments{
			Method: "GetBill",
			Router: `/bill/:clusterId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "CreateLoadBalancer",
			Router: `/loadbalancer`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "SetLoadBalancerStatus",
			Router: `/loadbalancer/status`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "DeleteLoadBalancer",
			Router: `/loadbalancer`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "GetLoadBalancers",
			Router: `/list/:regionId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "GetLoadBalancer",
			Router: `/loadbalancer/:loadbalancerid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "AddBackendServers",
			Router: `/backendservers/:loadbalancerid`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "RemoveBackendServers",
			Router: `/backendservers/:loadbalancerid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "SetBackendServers",
			Router: `/backendservers/:loadbalancerid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "DescribeHealthStatus",
			Router: `/backendservers/healthstatus/:loadbalancerid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "SetBackendOfLoadBalance",
			Router: `/backendservers/by_ip`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "AddToLoadBalance",
			Router: `/backendservers/by_ip`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "RemoveFromLoadBalance",
			Router: `/backendservers/by_ip`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "CreateLoadBalancerListener",
			Router: `/listener`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "SetLoadBalancerListenerAttribute",
			Router: `/listener`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "DescribeLoadBalancerListenerAttribute",
			Router: `/listener`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "DeleteLoadBalancerListener",
			Router: `/listener`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "StartLoadBalancerListener",
			Router: `/listener/start`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "StopLoadBalancerListener",
			Router: `/listener/stop`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "SetListenerAccessControlStatus",
			Router: `/listener/whitelist/status`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "AddListenerWhiteListItem",
			Router: `/listener/whitelist`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "RemoveListenerWhiteListItem",
			Router: `/listener/whitelist`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"] = append(beego.GlobalControllerRouter["weibo.com/opendcp/jupiter/controllers:SlbController"],
		beego.ControllerComments{
			Method: "DescribeListenerAccessControlAttribute",
			Router: `/listener/whitelist`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
