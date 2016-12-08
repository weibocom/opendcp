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

package instance

import (
	"weibo.com/opendcp/jupiter/conf"
	"weibo.com/opendcp/jupiter/dao"
	"weibo.com/opendcp/jupiter/models"
	"weibo.com/opendcp/jupiter/provider"
	"weibo.com/opendcp/jupiter/service/bill"
	"weibo.com/opendcp/jupiter/ssh"
	"weibo.com/opendcp/jupiter/logstore"
	"strings"
	"weibo.com/opendcp/jupiter/response"
	"fmt"
	"encoding/json"
)

func CreateOne(cluster *models.Cluster) (string, error) {
	providerDriver, err := provider.New(cluster.Provider)
	if err != nil {
		return "", err
	}
	instanceIds, errs := providerDriver.Create(cluster, 1)
	if errs != nil {
		return "", errs[0]
	}
	ins, err := providerDriver.GetInstance(instanceIds[0])
	if err != nil {
		return "", err
	}
	if err := dao.InsertInstance(ins); err != nil {
		return "", err
	}
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
			return err
		}
		logstore.Error(correlationId, instanceId, "delete instance, err:", err)
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

func StartSshService(instanceId string, ip string, correlationId string) error {
	sshCli, err := getSSHClient(ip, "")
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

func getSSHClient(ip string, path string) (*ssh.Client, error) {
	var auth ssh.Auth
	if path == "" {
		auth = ssh.Auth{
			Passwords: []string{conf.Config.Password},
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
	url:= conf.Config.Ansible.Url + "/api/getlog"
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
	resp := &octansResp {}
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

