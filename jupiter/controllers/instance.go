package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"weibo.com/opendcp/jupiter/models"
	"weibo.com/opendcp/jupiter/service/instance"
	"weibo.com/opendcp/jupiter/ssh"

	"github.com/astaxie/beego"
	"weibo.com/opendcp/jupiter/conf"
	_ "weibo.com/opendcp/jupiter/provider/aliyun"
	_ "weibo.com/opendcp/jupiter/provider/aws"
	"weibo.com/opendcp/jupiter/service/cluster"
	"time"
	"weibo.com/opendcp/jupiter/logstore"
)

const DEFAULT_CPU = 1
const DEFAULT_RAM = 1

// Operations about Instance
type InstanceController struct {
	BaseController
}

type AppendPhyDevRequest struct {
	InstanceList []models.PhyAuth `json:"instancelist"`
}

type AppendPhyDevResponse struct {
	Success int                `json:"success"`
	Failed  int                `json:"failed"`
	Errors  []string        `json:"errors"`
}

// @Title create instance
// @Description create a instance
// @router / [post]
func (ic *InstanceController) CreateInstance() {
	var ob models.Cluster
	err := json.Unmarshal(ic.Ctx.Input.RequestBody, &ob)
	if err != nil {
		beego.Error("Could parase request before crate instance: ", err)
		ic.RespInputError()
	}
	ip, err := instance.CreateOne(&ob)
	if err != nil {
		beego.Error("Create instance error:", err)
		ic.RespServiceError(err)
	}
	resp := ApiResponse{}
	resp.Content = ip
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title start
// @Description start the instances
// @router /start/:instanceId [post]
func (ic *InstanceController) StartInstance() {
	instanceId := ic.GetString("instanceId")
	if instanceId == "" {
		beego.Error("Could parse request before start instance")
		ic.RespMissingParams("instanceId")
		return
	}
	isStart, err := instance.StartOne(instanceId)
	if err != nil {
		beego.Error("Could not start instances", err)
		ic.RespServiceError(err)
	}
	resp := ApiResponse{}
	resp.Content = isStart
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title Get instance
// @Description Get instance information
// @router /:instanceId [get]
func (ic *InstanceController) GetInstance() {
	instanceId := ic.GetString(":instanceId")
	if instanceId == "" {
		beego.Error("Could parse request before get instance")
		ic.RespMissingParams("instanceId")
	}
	ins, err := instance.GetInstanceById(instanceId)
	if err != nil {
		beego.Error("get one instance err: ", err)
		ic.RespServiceError(err)
	}
	resp := ApiResponse{}
	resp.Content = ins
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title Check status
// @Description check instances status
// @router status/:instanceIds [get]
func (ic *InstanceController) GetInstancesStatus() {
	instanceIds := ic.GetString(":instanceIds")
	if instanceIds == "" {
		beego.Error("Could parse request before get instance")
		ic.RespMissingParams("instanceId")
	}
	instanceIdSlice := strings.Split(instanceIds, ",")
	ins, err := instance.GetInstancesStatus(instanceIdSlice)
	if err != nil {
		beego.Error("get multi instance err: ", err)
		ic.RespServiceError(err)
	}
	resp := ApiResponse{}
	resp.Content = ins
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title Update machine status
// @Description Update machine status
// @router /status [post]
func (ic *InstanceController) UpdateInstanceStatus() {
	var insStat models.InstanceIdStatus
	err := json.Unmarshal(ic.Ctx.Input.RequestBody, &insStat)
	if err != nil {
		beego.Error("Could parase request before crate instance: ", err)
		ic.RespInputError()
		return
	}
	status, err := instance.UpdateInstanceStatus(insStat.InstanceId, insStat.Status)
	if err != nil {
		beego.Error("update instance status err: ", err)
		ic.RespServiceError(err)
		return
	}
	resp := ApiResponse{}
	resp.Content = status
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title Delete instances
// @Description Delete many instances.
// @router /:instanceIds [delete]
func (ic *InstanceController) DeleteMulti() {
	correlationId := ic.Ctx.Input.Header("X-CORRELATION-ID")
	if len(correlationId) <= 0 {
		ic.RespMissingParams("X-CORRELATION-ID")
		return
	}
	instanceIds := ic.GetString(":instanceIds")
	instanceIdsArray := strings.Split(instanceIds, ",")
	for i := 0; i < len(instanceIdsArray); i++ {
		go instance.DeleteOne(instanceIdsArray[i], correlationId)
	}
	go func() {
		time.Sleep(time.Second*10)
		cluster.UpdateInstanceDetail()
		time.Sleep(time.Minute)
		cluster.UpdateInstanceDetail()
		time.Sleep(time.Minute*2)
		cluster.UpdateInstanceDetail()
	}()
	resp := ApiResponse{}
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title Download ssh key
// @Description Download ssh key
// @router sshkey/:ip [get]
func (ic *InstanceController) DownloadKey() {
	ip := ic.GetString(":ip")
	if ip == "" {
		ic.RespMissingParams("ip")
		return
	}
	dir := conf.Config.KeyDir
	res, err := instance.GetInstanceByIp(ip)
	if err != nil {
		beego.Error(err)
		return
	}
	var path string
	if len(res.PrivateIpAddress) > 0 {
		path = fmt.Sprintf("%s/%s.pem", dir, strings.Replace(res.PrivateIpAddress, ".", "-", -1))
	} else if len(res.PublicIpAddress) > 0 {
		path = fmt.Sprintf("%s/%s.pem", dir, strings.Replace(res.PublicIpAddress, ".", "-", -1))
	}
	if conf.FileExists(path) {
		os.Remove(path)
	}
	err = ssh.GetSSHKeyFromDb(res, path, true)
	if err != nil {
		beego.Error("get priv key err: ", err)
		ic.Ctx.WriteString("")
		return
	}
	ic.Ctx.Output.Download(path)
	os.Remove(path)
}

// @Title Upload ssh key
// @Description Upload ssh key
// @router sshkey/:instanceId [put]
func (ic *InstanceController) UploadKey() {
	instanceId := ic.GetString(":instanceId")
	if instanceId == "" {
		ic.RespMissingParams("instanceId")
		return
	}
	var sshKey models.SshKey
	err := json.Unmarshal(ic.Ctx.Input.RequestBody, &sshKey)
	if err != nil {
		beego.Error("Could parase request before upload ssh key: ", err)
		ic.RespInputError()
		return
	}
	resp := ApiResponse{}
	result, err := instance.UploadSshKey(instanceId, sshKey)
	if err != nil {
		beego.Error("input phy dev error:", err)
		ic.RespServiceError(err)
		return
	}
	resp.Content = result
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title Get providers
// @Description Get providers
// @router /provider [get]
func (ic *InstanceController) GetProviders() {
	providers, err := instance.GetProviders()
	if err != nil {
		ic.RespServiceError(err)
		return
	}
	resp := ApiResponse{}
	resp.Content = providers
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title Get region
// @Description Get region in provider
// @router /regions/:provider [get]
func (ic *InstanceController) GetRegionIds() {
	provider := ic.GetString(":provider")
	regions, err := instance.GetRegions(provider)
	if err != nil {
		beego.Error(err)
		ic.RespServiceError(err)
		return
	}
	resp := ApiResponse{}
	resp.Content = regions
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title Get available zone
// @Description Get available zone in provider
// @router /zones/:provider/:regionId [get]
func (ic *InstanceController) GetZones() {
	provider := ic.GetString(":provider")
	regionId := ic.GetString(":regionId")
	zones, err := instance.GetZones(provider, regionId)
	if err != nil {
		beego.Error(err)
		ic.RespServiceError(err)
		return
	}
	resp := ApiResponse{}
	resp.Content = zones
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title Get VPCs
// @Description Get VPCs in provider
// @router /vpc/:provider/:regionId [get]
func (ic *InstanceController) GetVpcs() {
	provider := ic.GetString(":provider")
	regionId := ic.GetString(":regionId")
	// pageNumber 起始值为1，pageSize 最大值为50
	vpcs, err := instance.GetVpcs(provider, regionId, 1, 50)
	if err != nil {
		beego.Error(err)
		ic.RespServiceError(err)
		return
	}
	resp := ApiResponse{}
	resp.Content = vpcs
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title Get subnets
// @Description Get subnets in provider
// @router /subnet/:provider/:zoneId/:vpcId [get]
func (ic *InstanceController) GetSubnets() {
	provider := ic.GetString(":provider")
	zoneId := ic.GetString(":zoneId")
	vpcId := ic.GetString(":vpcId")
	subnets, err := instance.GetSubnets(provider, zoneId, vpcId)
	if err != nil {
		beego.Error(err)
		ic.RespServiceError(err)
		return
	}
	resp := ApiResponse{}
	resp.Content = subnets
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title List instance types
// @Description List all instance types
// @router /type/:provider [get]
func (ic *InstanceController) ListInstanceTypes() {
	provider := ic.GetString(":provider")
	instanceType, err := instance.ListInstanceTypes(provider)
	if err != nil {
		beego.Error("get all instance type error: ", err)
		ic.RespServiceError(err)
	}
	resp := ApiResponse{}
	resp.Content = instanceType
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title List internet charge type
// @Description List all instance internet charge type
// @router /charge/:provider [get]
func (ic *InstanceController) ListInternetChargeType() {
	provider := ic.GetString(":provider")
	chargeType, err := instance.ListInternetChargeTypes(provider)
	if err != nil {
		beego.Error("get all internet charge type error: ", err)
		ic.RespServiceError(err)
	}
	resp := ApiResponse{}
	resp.Content = chargeType
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title List disk category
// @Description List disk category in provider
// @router /disk_category/:provider [get]
func (ic *InstanceController) ListDiskCategory() {
	provider := ic.GetString(":provider")
	dataCategory, err := instance.ListDiskCategory(provider)
	if err != nil {
		ic.RespServiceError(err)
	}
	resp := ApiResponse{}
	resp.Content = dataCategory
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title Get images
// @Description Get all images in provider
// @router /image/:provider/:regionId [get]
func (ic *InstanceController) GetImages() {
	provider := ic.GetString(":provider")
	regionId := ic.GetString(":regionId")
	images, err := instance.GetImages(provider, regionId)
	if err != nil {
		beego.Error(err)
		ic.RespServiceError(err)
		return
	}
	resp := ApiResponse{}
	resp.Content = images
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title Get security groups
// @Description Get all security groups in provider
// @router /security_group/:provider/:regionId [get]
func (ic *InstanceController) GetSecurityGroup() {
	provider := ic.GetString(":provider")
	regionId := ic.GetString(":regionId")
	vpcId := ic.GetString("vpcId")
	securityGroup, err := instance.GetSecurityGroup(provider, regionId, vpcId)
	if err != nil {
		beego.Error(err)
		ic.RespServiceError(err)
		return
	}
	resp := ApiResponse{}
	resp.Content = securityGroup
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title List all instances
// @Description List all instances.
// @router /list [get]
func (ic *InstanceController) ListAllInstances() {
	instances, err := instance.ListInstances()
	if err != nil {
		beego.Error("get all instances error: ", err)
		ic.RespServiceError(err)
	}
	resp := ApiResponse{}
	resp.Content = instances
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title List all instances by cluster id
// @Description List all instances in someone cluster
// @router /cluster/:clusterId [get]
func (ic *InstanceController) ListInstancesByClusterId() {
	clusterId, err := ic.GetInt64(":clusterId")
	if err != nil {
		beego.Error("Could parase cluster id: ", err)
		ic.RespInputError()
	}
	instances, err := instance.ListInstancesByClusterId(clusterId)
	if err != nil {
		beego.Error("List instaces in cluster error:", err)
		ic.RespServiceError(err)
	}
	resp := ApiResponse{}
	resp.Content = instances
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title Query instance log
// @Description Query log by correlation id and instance id
// @router /log/:correlationId/:instanceId [get]
func (ic *InstanceController) QueryLogByCorrelationIdAndInstanceId() {
	correlationId := ic.GetString(":correlationId")
	if len(correlationId) <= 0 {
		beego.Error("correlationId is empty!")
		ic.RespInputError()
		return
	}
	instanceId := ic.GetString(":instanceId")
	if len(instanceId) <= 0 {
		beego.Error("params instanceId is empty")
		ic.RespInputError()
		return
	}
	resp := ApiResponse{}
	data, err := instance.QueryLogByCorrelationIdAndInstanceId(instanceId, correlationId)
	if err != nil {
		beego.Error("[ResourceLogApi] getResourceLog result json error!", err)
		resp.Msg = "result to json error"
		ic.ApiResponse = resp
		ic.Status = BAD_REQUEST
		ic.RespJsonWithStatus()
	}
	resp.Content = data
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title Query instance log
// @Description Query log by instance id
// @router /log/:instanceId [get]
func (ic *InstanceController) QueryLogByInstanceId() {
	instanceId := ic.GetString(":instanceId")
	if len(instanceId) <= 0 {
		beego.Error("params instanceId is empty")
		ic.RespInputError()
		return
	}
	resp := ApiResponse{}
	data, err := instance.QueryLogByInstanceId(instanceId)
	if err != nil {
		beego.Error("[ResourceLogApi] getResourceLog result json error!", err)
		resp.Msg = "result to json error"
		ic.ApiResponse = resp
		ic.Status = BAD_REQUEST
		ic.RespJsonWithStatus()
	}
	resp.Content = data
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title Upload machine infomation
// @Description Upload machine information to DB
// @router /phydev [put]
func (ic *InstanceController) UploadPhyDevInfo() {
	var ins models.Instance
	err := json.Unmarshal(ic.Ctx.Input.RequestBody, &ins)
	if err != nil {
		beego.Error("Could parase request before input instance: ", err)
		ic.RespInputError()
		return
	}
	resp := ApiResponse{}
	result, err := instance.InputPhyDev(ins)
	if err != nil {
		beego.Error("input phy dev error:", err)
		ic.RespServiceError(err)
		return
	}
	resp.Content = result
	ic.ApiResponse = resp
	ic.Status = SERVICE_SUCCESS
	ic.RespJsonWithStatus()
}

// @Title manage physical device
// @Description manage physical device
// @router /phydev [post]
func (ic *InstanceController) ManagePhyDev() {
	correlationId := ic.Ctx.Input.Header("X-CORRELATION-ID")
	if len(correlationId) <= 0 {
		ic.RespMissingParams("X-CORRELATION-ID")
		return
	}

	var request AppendPhyDevRequest
	err := json.Unmarshal(ic.Ctx.Input.RequestBody, &request)
	if err != nil {
		beego.Error("Could parase request before input instance: ", err)
		ic.RespInputError()
		return
	}

	// 1. check request form
	for _, info := range request.InstanceList {
		if (strings.TrimSpace(info.PublicIp) == "" && strings.TrimSpace(info.PrivateIp) == "") || strings.TrimSpace(info.Password) == "" {
			beego.Error("Please check request")
			ic.RespInputError()
			return
		}
	}

	// 2. start insert DB
	successCount := 0
	failedCount := 0
	errList := make([]string, 0)
	for _, info := range request.InstanceList {
		ip := info.PublicIp
		if ip == "" {
			ip = info.PrivateIp
		}
		// already in database, skip
		inst, _ := instance.GetInstanceByIp(ip)
		if inst != nil {
			failedCount++
			errList = append(errList, "Instance: "+ip+" is already in DB")
			continue
		}
		var ins models.Instance
		ins.Cpu = DEFAULT_CPU
		ins.Ram = DEFAULT_RAM
		ins.PublicIpAddress = info.PublicIp
		ins.PrivateIpAddress = info.PrivateIp

		logstore.Info(correlationId, ins.InstanceId, "1. Insert the instance into db")
		ins, err = instance.InputPhyDev(ins)

		if err != nil {
			failedCount++
			errList = append(errList, err.Error())
		} else {
			successCount++
			// asynchronous manage
			logstore.Info(correlationId, ins.InstanceId, "Insert the instance into db successfully")
			logstore.Info(correlationId, ins.InstanceId, "2. Begin to execute init operation in the instance")
			go instance.ManageDev(ip, info.Password, ins.InstanceId, correlationId)
		}
	}
	go cluster.UpdateInstanceDetail()

	// 3. response
	resp := ApiResponse{}
	resp.Content = AppendPhyDevResponse{
		Success:successCount,
		Failed: failedCount,
		Errors: errList,
	}
	ic.ApiResponse = resp
	if failedCount == 0 {
		ic.Status = SERVICE_SUCCESS
	} else {
		ic.Status = SERVICE_ERRROR
	}
	ic.RespJsonWithStatus()
}

