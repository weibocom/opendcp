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



package cluster

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
	"weibo.com/opendcp/jupiter/dao"
	"weibo.com/opendcp/jupiter/future"
	"weibo.com/opendcp/jupiter/models"
	"weibo.com/opendcp/jupiter/provider"
	"weibo.com/opendcp/jupiter/service/bill"
	"weibo.com/opendcp/jupiter/logstore"
	"errors"
	"github.com/astaxie/beego"
	"strings"
)

func GetCluster(clusterId int64) (*models.Cluster, error) {
	cluster, err := dao.GetClusterById(clusterId)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

func CreateCluster(cluster *models.Cluster) (int64, error) {
	cluster.CreateTime = time.Now()
	providerDriver, err := provider.New(cluster.Provider)
	if err != nil {
		return 0, err
	}

	//在openstack中，Flavor的名称就是对应的类型，不需要再进行转换
	if(cluster.Provider == "aliyun") {
		instanceTypeModel := cluster.InstanceType
		validNumber := regexp.MustCompile("[0-9]")
		cpuAndRam := validNumber.FindAllString(instanceTypeModel, -1)
		cluster.InstanceType = providerDriver.GetInstanceType(instanceTypeModel)
		cpu, _ := strconv.Atoi(cpuAndRam[0])
		ram, _ := strconv.Atoi(cpuAndRam[1])
		cluster.Cpu = cpu
		cluster.Ram = ram
	}else if(cluster.Provider == "openstack"){
		cluster.FlavorId = cluster.InstanceType
		resp := providerDriver.GetInstanceType(cluster.InstanceType)
		strs := strings.Split(resp, '#')
		cluster.InstanceType = strs[0]
		cluster.Cpu = strs[1]
		cluster.Ram = strs[2]
	}

	id, err := dao.InsertCluster(cluster)
	_, err = bill.InsertBill(cluster)
	return id, err
}

func DeleteCluster(clusterId int64) (bool, error) {
	instances, err := dao.ListInstancesByClusterId(clusterId)
	if  len(instances) > 0 {
		return false, errors.New("Can't delete this cluster, because still exsit instance by cluster model created.")
	}
	isDeleted, err := dao.DeleteCluster(clusterId)
	return isDeleted, err
}

func Expand(cluster *models.Cluster, num int, correlationId string) ([]string, error) {
	providerDriver, err := provider.New(cluster.Provider)
	if err != nil {
		return nil, err
	}
	instanceIds, errs := providerDriver.Create(cluster, num)
	if len(instanceIds) == 0 {
		return nil, errs[0]
	}
	if len(errs) > 0 {
		beego.Error("Expand failed number is", len(errs), "errors is", errs)
	}
	beego.Info("The instance ids is", instanceIds)
	c := make(chan int)
	for i := 0; i < len(instanceIds); i++ {
		go func(i int) {
			defer func() {
				if r := recover(); r != nil {
					logstore.Error(correlationId, instanceIds[i], "Recovered from err:", r)
					switch r.(type) {
					case error:
						err = r.(error)
					default:
						err = errors.New(fmt.Sprint("Expand machine faile:", r))
					}
				}
			}()
			ins, err := providerDriver.GetInstance(instanceIds[i])
			if err != nil {
				logstore.Error(correlationId, instanceIds[i], "get instance info error:", err)
				c <- i
			}
			ins.Cluster = cluster
			ins.Cpu = cluster.Cpu
			ins.Ram = cluster.Ram
			ins.DataDiskCategory = cluster.DataDiskCategory
			ins.DataDiskSize = cluster.DataDiskSize
			ins.DataDiskNum = cluster.DataDiskNum
			ins.SystemDiskCategory = cluster.SystemDiskCategory
			ins.InstanceType = cluster.InstanceType
			ins.Status = models.Pending
			if err := dao.InsertInstance(ins); err != nil {
				logstore.Error(correlationId, instanceIds[i], "insert instance to db error:", err)
				c <- i
			}
			startFuture := future.NewStartFuture(instanceIds[i], cluster.Provider, true, ins.PrivateIpAddress, correlationId)
			future.Exec.Submit(startFuture)
			c <- i
		}(i)
	}
	for i := 0; i < len(instanceIds); i++ {
		select {
		case <-c:
		}
	}
	return instanceIds, nil
}

func ListClusters() ([]models.Cluster, error) {
	clusters, err := dao.GetClusters()
	if err != nil {
		return nil, err
	}
	return clusters, nil
}
