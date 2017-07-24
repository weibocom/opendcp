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

func GetInstanceIncludeDeleted(instanceId string) (*models.Instance, error) {
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
	if err != nil {
		return err
	}
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
	if err != nil {
		return err
	}
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

func UpdateInstanceStatusByInstanceId(instanceId string, status models.InstanceStatus) error {
	o := GetOrmer()
	instance, err := GetInstance(instanceId)
	if err != nil {
		return err
	}
	instance.Status = status
	_, err = o.Update(instance)
	if err != nil {
		return err
	}
	return nil
}

func GetIpByInstanceId(instanceId string) (string, error) {
	ins, err := GetInstanceIncludeDeleted(instanceId)
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

func UpdateSshKey(instanceId string, publicKey string, privateKey string) error {
	o := GetOrmer()
	ins, err := GetInstance(instanceId)
	ins.PublicKey = publicKey
	ins.PrivateKey = privateKey
	_, err = o.Update(ins)
	return err
}

func GetAllRunningInstance () (instances []models.Instance,err error) {
	o := GetOrmer()
	_,err = o.QueryTable(INSTANCE_TABLE).Exclude("status", models.Deleted).All(&instances)

	if err != nil {
		return nil, err
	}

	return instances, nil
}

func GetAllInstanceByClusterId(clusterId int64) (instances []models.Instance, err error) {
	o := GetOrmer()
	_, err = o.QueryTable(INSTANCE_TABLE).RelatedSel().Filter("cluster", clusterId).Exclude("status", models.Deleted).
		All(&instances)

	if err != nil {
		return nil, err
	}

	return instances, nil
}

func GetInstancesByProvider(provider string) ([]models.Instance, error) {
	o := GetOrmer()
	var instances []models.Instance
	_, err := o.QueryTable(INSTANCE_TABLE).Filter("provider", provider).Exclude("status", models.Deleted).All(&instances)
	if err != nil {
		return nil, err
	}
	return instances, nil
}