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

package dao

import (
	"weibo.com/opendcp/jupiter/models"
	"errors"
)

func GetInstance(instanceId string) (*models.Instance, error) {
	o := GetOrmer()
	var instance models.Instance
	err := o.QueryTable(INSTANCE_TABLE).RelatedSel().Filter("instance_id", instanceId).Exclude("status", models.Deleted).One(&instance)
	if err != nil {
		return nil, err
	}
	return &instance, nil
}

func GetInstanceIncludeDelted(instanceId string) (*models.Instance, error) {
	o := GetOrmer()
	var instance models.Instance
	err := o.QueryTable(INSTANCE_TABLE).RelatedSel().Filter("instance_id", instanceId).One(&instance)
	if err != nil {
		return nil, err
	}
	return &instance, nil
}

func GetClusterByInstanceId(instanceId string) (*models.Cluster, error) {
	instance, err := GetInstance(instanceId)
	if err != nil {
		return nil, err
	}
	return instance.Cluster, nil
}

func UpdateDeletedStatus(instanceId string) error {
	o := GetOrmer()
	instance, err := GetInstance(instanceId)
	if err != nil {
		return err
	}
	instance.Status = models.Deleted
	_, err = o.Update(instance)
	if err != nil {
		return err
	}
	return nil
}


func UpdateDeletingStatus(instanceId string) error {
	o := GetOrmer()
	instance, err := GetInstance(instanceId)
	if err != nil {
		return err
	}
	instance.Status = models.Deleting
	_, err = o.Update(instance)
	if err != nil {
		return err
	}
	return nil
}

func GetInstanceByPrivateIp(ip string) (*models.Instance, error) {
	o := GetOrmer()
	var instance models.Instance
	err := o.QueryTable(INSTANCE_TABLE).RelatedSel().Filter("private_ip_address", ip).Exclude("status", models.Deleted).One(&instance)
	if err != nil {
		return nil, err
	}
	return &instance, nil
}

func GetInstanceByPublicIp(ip string) (*models.Instance, error) {
	o := GetOrmer()
	var instance models.Instance
	err := o.QueryTable(INSTANCE_TABLE).RelatedSel().Filter("public_ip_address", ip).Exclude("status", models.Deleted).One(&instance)
	if err != nil {
		return nil, err
	}
	return &instance, nil
}

func InsertInstance(instance *models.Instance) error {
	o := GetOrmer()
	_, err := o.Insert(instance)
	return err
}

func UpdateInstancePrivateIp(instanceId, private_ip_address string) error {
	o := GetOrmer()
	instance, err := GetInstance(instanceId)
	instance.PrivateIpAddress = private_ip_address
	_, err = o.Update(instance)
	if err != nil {
		return err
	}
	return err
}

func UpdateInstancePublicIp(instanceId, public_ip_address string) error {
	o := GetOrmer()
	instance, err := GetInstance(instanceId)
	instance.PublicIpAddress = public_ip_address
	_, err = o.Update(instance)
	if err != nil {
		return err
	}
	return err
}

func ListInstances() ([]models.Instance, error) {
	o := GetOrmer()
	var instances []models.Instance
	_, err := o.QueryTable(INSTANCE_TABLE).RelatedSel().Exclude("status", models.Deleted).OrderBy("-id").All(&instances)
	if err != nil {
		return nil, err
	}
	return instances, nil
}

func ListInstancesByClusterId(clusterId int64) ([]models.Instance, error) {
	o := GetOrmer()
	var instances []models.Instance
	_, err := o.QueryTable(INSTANCE_TABLE).RelatedSel().Filter("cluster_id", clusterId).Exclude("status", models.Deleted).All(&instances)
	if err != nil {
		return nil, err
	}
	return instances, nil
}

func UpdateInstanceStatus(ip string, status models.InstanceStatus) error {
	o := GetOrmer()
	instance, err := GetInstanceByPrivateIp(ip)
	if err != nil {
		instance, err = GetInstanceByPublicIp(ip)
		if err != nil {
			return err
		}
	}
	instance.Status = status
	_, err = o.Update(instance)
	if err != nil {
		return err
	}
	return nil
}

func GetIpByInstanceId(instanceId string) (string, error) {
	ins, err := GetInstanceIncludeDelted(instanceId)
	if err != nil {
		return "", err
	}
	if len(ins.PrivateIpAddress) > 0 {
		return ins.PrivateIpAddress, nil
	}
	if len(ins.PublicIpAddress) > 0 {
		return ins.PublicIpAddress, nil
	}
	return "", errors.New("The instance no private ip address or public ip address.")
}
