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

package instance

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/rs/xid"
	"strings"
	"time"
	"weibo.com/opendcp/jupiter/conf"
	"weibo.com/opendcp/jupiter/dao"
	"weibo.com/opendcp/jupiter/logstore"
	"weibo.com/opendcp/jupiter/models"
	"weibo.com/opendcp/jupiter/provider"
	"weibo.com/opendcp/jupiter/response"
	"weibo.com/opendcp/jupiter/service/bill"
	"weibo.com/opendcp/jupiter/ssh"
	"os/exec"
)

const (
	PhyDev = "phydev"
	OPENSTACK = "openstack"
)

func CreateOne(cluster *models.Cluster) (string, error) {
	providerDriver, err := provider.New(cluster.Provider)
	if err != nil {
		return "", err
	}
	instanceIds, errs := providerDriver.Create(cluster, 1)
	if len(errs) != 0 {
		return "", errs[0]
	}
	//if errs != nil {
	//	return "", errs[0]
	//}
	fmt.Println("getting instance")
	ins, err := providerDriver.GetInstance(instanceIds[0])
	if err != nil {
		return "", err
	}
	fmt.Println("insert into database")
	if err := dao.InsertInstance(ins); err != nil {
		return "", err
	}
	fmt.Println("have inserted")
	return instanceIds[0], nil
}

func StartOne(instanceId string) (bool, error) {
	ins, err := GetInstanceById(instanceId)
	if err != nil {
		return false, err
	}
	providerDriver, err := provider.New(ins.Provider)
	if err != nil {
		return false, err
	}
	isStart, err := providerDriver.Start(ins.InstanceId)
	if err != nil {
		return false, err
	}
	return isStart, nil
}

func StopOne(instanceId string) (bool, error) {
	ins, err := GetInstanceById(instanceId)
	if err != nil {
		return false, err
	}
	providerDriver, err := provider.New(ins.Provider)
	if err != nil {
		return false, err
	}
	isStop, err := providerDriver.Stop(ins.InstanceId)
	if err != nil {
		return false, err
	}
	return isStop, nil
}

func DeleteOne(instanceId, correlationId string) error {
	logstore.Info(correlationId, instanceId,"1. Begin to delete instance")
	err := dao.UpdateDeletingStatus(instanceId)
	if err != nil {
		logstore.Error(correlationId, instanceId, "update deleting status err:", err)
		return err
	}
	ins, err := dao.GetInstance(instanceId)
	if err != nil {
		logstore.Error(correlationId, instanceId, "get instance in db err:", err)
		return err
	}
	logstore.Info(correlationId, instanceId,"2. Delete instance from cloud")
	if ins.Provider != PhyDev {
		providerDriver, err := provider.New(ins.Provider)
		if err != nil {
			logstore.Error(correlationId, instanceId, err)
			return err
		}
		_, err = providerDriver.Delete(instanceId)
		if err != nil {
			if strings.Contains(err.Error(), "InvalidInstanceId.NotFound") {
				//实例已经被删除，可能在其他系统中删除的，需要继续往下走，删除系统数据库的记录
				logstore.Info(correlationId, instanceId, "the instance already deleted, err:", err)
			} else {
				logstore.Error(correlationId, instanceId, "delete instance, err:", err)
				return err
			}
		}
		logstore.Info(correlationId, instanceId, "delete instance", instanceId, "success")
		usageHours, err := bill.GetUsageHours(instanceId)
		cluster, err := GetCluster(instanceId)
		if err != nil {
			logstore.Error(correlationId, instanceId, "get cluster, err:", err)
			return err
		}
		err = bill.Bill(cluster, usageHours)
		if err != nil {
			logstore.Error(correlationId, instanceId, "update bill, err:", err)
			return err
		}
	}
	logstore.Info(correlationId, instanceId,"3. Update db instance status to deleted")
	err = dao.UpdateDeletedStatus(instanceId)
	if err != nil {
		logstore.Error(correlationId, instanceId, "update deleted status, err:", err)
		return err
	}
	logstore.Info(correlationId, instanceId, "update instance status in DB success", instanceId, "success")
	return nil
}

func GetCluster(instanceId string) (*models.Cluster, error) {
	cluster, err := dao.GetClusterByInstanceId(instanceId)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

func GetInstanceByIp(ip string) (*models.Instance, error) {
	var instance *models.Instance
	instance, err := dao.GetInstanceByPrivateIp(ip)
	if err != nil {
		instance, err = dao.GetInstanceByPublicIp(ip)
		if err != nil {
			return nil, err
		}
	}
	return instance, nil
}

func GetInstanceById(instanceId string) (*models.Instance, error) {
	instance, err := dao.GetInstance(instanceId)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func GetInstancesStatus(instancesIds []string) ([]models.StatusResp, error) {
	var results []models.StatusResp
	for i := 0; i < len(instancesIds); i++ {
		instance, err := GetInstanceById(instancesIds[i])
		var tmpInstance models.StatusResp
		tmpInstance.InstanceId = instancesIds[i]
		if err != nil {
			tmpInstance.Status = models.StatusError
			results = append(results, tmpInstance)
			continue
		}
		tmpInstance.Status = instance.Status
		if len(instance.PrivateIpAddress) > 0 {
			tmpInstance.IpAddress = instance.PrivateIpAddress
		}
		if len(instance.PublicIpAddress) > 0 {
			tmpInstance.IpAddress = instance.PublicIpAddress
		}

		results = append(results, tmpInstance)
	}
	return results, nil
}

func GetProviders() ([]string, error) {
	return provider.ListDrivers(), nil
}

func GetRegions(providerName string) ([]models.Region, error) {
	providerDriver, err := provider.New(providerName)
	if err != nil {
		return nil, err
	}
	ret, err := providerDriver.ListRegions()
	if err != nil {
		return nil, err
	}
	return ret.Regions, nil
}

func GetZones(providerName string, regionId string) ([]models.AvailabilityZone, error) {
	providerDriver, err := provider.New(providerName)
	if err != nil {
		return nil, err
	}
	ret, err := providerDriver.ListAvailabilityZones(regionId)
	if err != nil {
		return nil, err
	}
	return ret.AvailabilityZones, nil
}

func GetVpcs(providerName string, regionId string, pageNumber int, pageSize int) ([]models.Vpc, error) {
	providerDriver, err := provider.New(providerName)
	if err != nil {
		return nil, err
	}
	ret, err := providerDriver.ListVpcs(regionId, pageNumber, pageSize)
	if err != nil {
		return nil, err
	}
	return ret.Vpcs, nil
}

func GetSubnets(providerName string, zoneId string, vpcId string) ([]models.Subnet, error) {
	providerDriver, err := provider.New(providerName)
	if err != nil {
		return nil, err
	}
	ret, err := providerDriver.ListSubnets(zoneId, vpcId)
	if err != nil {
		return nil, err
	}
	return ret.Subnets, nil
}

func GetImages(providerName string, regionId string) ([]models.Image, error) {
	providerDriver, err := provider.New(providerName)
	if err != nil {
		return nil, err
	}
	ret, err := providerDriver.ListImages(regionId, "", 50, 1)
	if err != nil {
		return nil, err
	}
	return ret.Images, nil
}

func ListInstanceTypes(providerName string) ([]string, error) {
	providerDriver, err := provider.New(providerName)
	if err != nil {
		return nil, err
	}
	ret, err := providerDriver.ListInstanceTypes()
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func ListInternetChargeTypes(providerName string) ([]string, error) {
	providerDriver, err := provider.New(providerName)
	if err != nil {
		return nil, err
	}
	return providerDriver.ListInternetChargeType(), nil
}

func ListDiskCategory(providerName string) ([]string, error) {
	providerDriver, err := provider.New(providerName)
	if err != nil {
		return nil, err
	}
	return providerDriver.ListDiskCategory(), nil
}

func GetSecurityGroup(providerName string, regionId string, vpcId string) ([]models.SecurityGroup, error) {
	providerDriver, err := provider.New(providerName)
	if err != nil {
		return nil, err
	}
	ret, err := providerDriver.ListSecurityGroup(regionId, vpcId)
	if err != nil {
		return nil, err
	}
	return ret.SecurityGroups, nil
}

func ListInstances() ([]models.Instance, error) {
	instances, err := dao.ListInstances()
	if err != nil {
		return nil, err
	}
	return instances, nil
}

func ListInstancesByClusterId(clusterId int64) ([]models.Instance, error) {
	instances, err := dao.ListInstancesByClusterId(clusterId)
	if err != nil {
		return nil, err
	}
	return instances, nil
}

func StartSshService(instanceId string, ip string, password string, correlationId string) error {
	sshCli, err := getSSHClient(ip, "", password)
	if err != nil {
		return err
	}
	err = sshCli.StoreSSHKey(instanceId)
	if err != nil {
		return err
	}
	logstore.Info(correlationId, instanceId, "ssh key pair end for instance: ", instanceId)
	return nil
}

func getSSHClient(ip string, path string, password string) (*ssh.Client, error) {
	var auth ssh.Auth
	if path == "" {
		auth = ssh.Auth{
			Passwords: []string{password},
		}
	} else {
		auth = ssh.Auth{
			Keys: []string{path},
		}
	}
	port := 22
	sshCli, err := ssh.NewClient("root", ip, port, &auth)
	if err != nil {
		return nil, err
	}
	return sshCli, nil
}

func QueryLogByCorrelationIdAndInstanceId(instanceId string, correlationId string) (string, error) {
	store := logstore.Store{}
	logInfo := store.QueryLogByCorrelationIdAndInstanceId(instanceId, correlationId)
	jupiterLog := logInfo.Message
	url := conf.Config.Ansible.Url + "/api/getlog"
	ip, err := dao.GetIpByInstanceId(instanceId)
	if err != nil {
		return "", err
	}
	body := "{\"host\": \"%s\", \"source\":\"jupiter\"}"
	body = fmt.Sprintf(body, ip)
	raw, err := response.CallApi(body, "POST", url, correlationId)
	if err != nil {
		logstore.Error(correlationId, instanceId, "Error when getting log for", instanceId, "err:", err)
		return "<ERROR> Call octans error", err
	}
	type octansResp struct {
		Content struct {
			Log []string
		}
	}
	resp := &octansResp{}
	err = json.Unmarshal([]byte(raw), &resp)
	if err != nil {
		logstore.Error(correlationId, instanceId, "Error when parsing log for", instanceId, "err:", err)
		return "<ERROR> Call octans error", err
	}
	return jupiterLog + "\n" + strings.Join(resp.Content.Log, "\n"), nil
}

func QueryLogByInstanceId(instanceId string) (string, error) {
	store := logstore.Store{}
	logInfo := store.QueryLogByInstanceId(instanceId)
	jupiterLog := logInfo.Message
	return jupiterLog, nil
}

func InputPhyDev(ins models.Instance) (models.Instance, error) {
	clusters, err := dao.GetClustersByProvider(PhyDev)
	if err != nil {
		return ins, err
	}
	var cluster models.Cluster
	if len(clusters) == 0 {
		cluster = models.Cluster{
			Name:       "Physical device",
			Provider:   "phydev",
			Desc:       "About physical device",
			CreateTime: time.Now(),
			Network:    &models.Network{},
			Zone:       &models.Zone{},
		}
		dao.InsertCluster(&cluster)
		_, err = bill.InsertBill(&cluster)
		if err != nil {
			return ins, err
		}
		ins.Cluster = &cluster
	} else {
		ins.Cluster = &clusters[0]
	}
	guid := xid.New()
	instanceId := "i-" + guid.String()
	ins.InstanceId = instanceId
	ins.Provider = PhyDev
	ins.Status = models.Initing
	if err := dao.InsertInstance(&ins); err != nil {
		return ins, err
	}
	return ins, nil
}

func UploadSshKey(instanceId string, sshKey models.SshKey) (models.SshKey, error) {
	err := dao.UpdateSshKey(instanceId, sshKey.PublicKey, sshKey.PrivateKey)
	return sshKey, err
}

func UpdateInstanceStatus(instanceId string, status models.InstanceStatus) (models.InstanceStatus, error) {
	err := dao.UpdateInstanceStatusByInstanceId(instanceId, status)
	if err != nil {
		return status, err
	}
	return status, nil
}

func ManageDev(ip, password, instanceId, correlationId string) (ssh.Output, error) {
	logstore.Info(correlationId, instanceId, "(1) Begin to generate the ssk keys and init the ssh connection")
	sshErr := StartSshService(instanceId, ip, password, correlationId)
	if sshErr != nil {
		logstore.Error(correlationId, instanceId, "ssh instance: ", instanceId, "failed: ", sshErr)
		dao.UpdateInstanceStatus(ip, models.InitTimeout)
		return ssh.Output{}, sshErr
	}
	logstore.Info(correlationId, instanceId, "Store the ssk keys finished.")
	cli, err := getSSHClient(ip, "", password)

	octanUrl := conf.Config.Ansible.GetOctansUrl

	ins, err := dao.GetInstance(instanceId)
	if err != nil {
		return ssh.Output{}, err
	}

	if strings.EqualFold(ins.Provider, OPENSTACK){
		octanUrl = fmt.Sprintf(octanUrl + "/manage_device_%s.sh", OPENSTACK)
	} else {
		octanUrl = fmt.Sprintf(octanUrl + "/manage_device.sh")
	}

	cmd := fmt.Sprintf("curl %s -o /root/manage_device.sh && chmod +x /root/manage_device.sh", octanUrl)
	logstore.Info(correlationId,instanceId,"(2) Download the init script in instance:"+cmd)

	ret, err := cli.Run(cmd)
	if err != nil {
		dao.UpdateInstanceStatus(ip, models.StatusError)
		result := fmt.Sprintf("Exec cmd %s fail: %s", cmd, err)
		logstore.Error(correlationId,instanceId,result)
		return ssh.Output{}, err
	}
	logstore.Info(correlationId, instanceId, "Download script result:",ret)
	dbAddr := beego.AppConfig.String("host")
	jupiterAddr := beego.AppConfig.String("host")
	cmd = fmt.Sprintf("sh /root/manage_device.sh mysql://%s:%s@%s:%s/octans?charset=utf8  http://%s:8083/v1/instance/sshkey/ %s:8083 %s %s %s > /root/result.out",
		beego.AppConfig.String("mysqluser"), beego.AppConfig.String("mysqlpass"), dbAddr, beego.AppConfig.String("mysqlport"), jupiterAddr, jupiterAddr, instanceId, ip, beego.AppConfig.String("harbor_registry"))
	cmdOut := fmt.Sprintf("sh /root/manage_device.sh mysql://****:****@%s:%s/octans?charset=utf8  http://%s:8083/v1/instance/sshkey/ %s:8083 %s %s %s > /root/result.out",
		  dbAddr, beego.AppConfig.String("mysqlport"), jupiterAddr, jupiterAddr, instanceId, ip, beego.AppConfig.String("harbor_registry"))
	logstore.Info(correlationId, instanceId, "(3) Execute init operaration in instance："+cmdOut)
	ret, err = cli.Run(cmd)
	if err != nil {
		dao.UpdateInstanceStatus(ip, models.StatusError)
		result := fmt.Sprintf("Exec cmd [ %s ] fail: %s", cmd, err)
		logstore.Error(correlationId,instanceId,result)
		return ssh.Output{}, err
	}
	logstore.Info(correlationId, instanceId, "Execute script result:", ret)
	return ret, nil
}


func ChangeOpenStackConf(OpConf *models.OpenStackConf) error{
	//修改hosts文件的controller域名

	op := fmt.Sprintf("awk '{if($2==\"controller\") {$1=\"%s\"} print}' /etc/hosts > /etc/hostsbak",  OpConf.OpIp)
	cmd := exec.Command("/bin/sh", "-c", op)
	err := cmd.Run()
	if err != nil{
		return err
	}
	cmd = exec.Command("/bin/sh", "-c", "cp /etc/hostsbak /etc/hosts")
	err = cmd.Run()
	if err != nil{
		return err
	}
	cmd = exec.Command("/bin/sh", "-c", "rm /etc/hostsbak")
	err = cmd.Run()
	if err != nil{
		return err
	}

	conf.Config.OpIp = OpConf.OpIp
	conf.Config.OpPort = OpConf.OpPort
	conf.Config.OpUserName = OpConf.OpUserName
	conf.Config.OpPassWord = OpConf.OpPassWord
	return err

}

func ListAllInstances() ([]models.Instance, error) {
	instances, err := dao.GetAllRunningInstance()
	if err != nil {
		return nil, err
	}
	return instances, nil
}


func GetClusterInstances(clusterId int64)  ([]models.Instance, error){
	instances, err := dao.GetAllInstanceByClusterId(clusterId)
	if err != nil {
		return nil, err
	}
	return instances, err
}

