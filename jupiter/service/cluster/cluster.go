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
	instanceTypeModel := cluster.InstanceType
	validNumber := regexp.MustCompile("[0-9]")
	cpuAndRam := validNumber.FindAllString(instanceTypeModel, -1)
	cluster.InstanceType = providerDriver.GetInstanceType(instanceTypeModel)
	cpu, _ := strconv.Atoi(cpuAndRam[0])
	ram, _ := strconv.Atoi(cpuAndRam[1])
	cluster.Cpu = cpu
	cluster.Ram = ram
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
