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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unsafe"
	"weibo.com/opendcp/jupiter/dao"
	"weibo.com/opendcp/jupiter/future"
	"weibo.com/opendcp/jupiter/models"
	"weibo.com/opendcp/jupiter/provider"
	"weibo.com/opendcp/jupiter/service/bill"
	"weibo.com/opendcp/jupiter/service/instance"
	"weibo.com/opendcp/jupiter/service/task"
)

var its task.InstanceTaskService

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
	if cluster.Provider == "aliyun" {
		instanceTypeModel := cluster.InstanceType
		validNumber := regexp.MustCompile("[0-9]")
		cpuAndRam := validNumber.FindAllString(instanceTypeModel, -1)
		cluster.InstanceType = providerDriver.GetInstanceType(instanceTypeModel)
		cpu, _ := strconv.Atoi(cpuAndRam[0])
		ram, _ := strconv.Atoi(cpuAndRam[1])
		cluster.Cpu = cpu
		cluster.Ram = ram
	} else if cluster.Provider == "openstack" {
		cluster.FlavorId = cluster.InstanceType
		resp := providerDriver.GetInstanceType(cluster.InstanceType)
		strs := strings.Split(resp, "#")
		cluster.InstanceType = strs[0]
		cluster.Cpu, _ = strconv.Atoi(strs[1])
		ram, _ := strconv.Atoi(strs[2])
		cluster.Ram = ram / 1024
	}
	id, err := dao.InsertCluster(cluster)
	_, err = bill.InsertBill(cluster)
	return id, err
}

func DeleteCluster(clusterId int64) (bool, error) {
	instances, err := dao.ListInstancesByClusterId(clusterId)
	if len(instances) > 0 {
		return false, errors.New("Can't delete this cluster, because still exsit instance by cluster model created.")
	}
	isDeleted, err := dao.DeleteCluster(clusterId)
	return isDeleted, err
}

/*
func Expand(cluster *models.Cluster, num int, correlationId string) ([]string, error) {
	providerDriver, err := provider.New(cluster.Provider)
	if err != nil {
		return nil, err
	}
	beego.Info("First. Begin to create instances from cloud")
	instanceIds, errs := providerDriver.Create(cluster, num)
	if len(instanceIds) == 0 {
		return nil, errs[0]
	}
	if len(errs) > 0 {
		beego.Error("Expand failed number is", len(errs), "errors is", errs)
	}
	beego.Info("The instance ids is", instanceIds)
	beego.Info("Second. Begin to start and init instances ----")
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
			logstore.Info(correlationId, instanceIds[i], "1. Begin to insert instances into db")
			ins, err := providerDriver.GetInstance(instanceIds[i])
			if err != nil {
				logstore.Error(correlationId, instanceIds[i], "get instance info error:", err)
				c <- i
			}
			logstore.Info(correlationId, instanceIds[i], "get instance info successfully")
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
			logstore.Info(correlationId, instanceIds[i], "insert instance into db successfully")
			logstore.Info(correlationId, instanceIds[i], "2. Begin start instance in future")
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
	go UpdateInstanceDetail()
	return instanceIds, nil
}*/

func Expand(cluster *models.Cluster, num int, correlationId string) ([]string, error) {
	tasks := make([]models.InstanceItem, num)
	taskId := fmt.Sprintf("TASKS-[%s]-[%d]", time.Now().Format("15:04:05"), len(tasks))
	beego.Info("First. Begin to create instance task, task id:", taskId)

	for i := 0; i < num; i++ {
		task := models.InstanceItem{
			TaskId:        taskId,
			CorrelationId: correlationId,
			Cluster:       cluster,
			CreateTime:    time.Now(),
		}
		tasks[i] = task
	}

	err := its.CreateTasks(tasks)
	if err != nil {
		beego.Error("Create instance tasks err:", err)
		return nil, err
	}

	its.WaitTasksComplete(tasks)

	go UpdateInstanceDetail()

	instanceIds := make([]string, num)
	for index, task := range tasks {
		instanceIds[index] = task.InstanceId
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

func UpdateInstanceDetail() error {
	instanceInfo, err := GetLatestDetail()
	if err != nil {
		return err
	}

	instanceData, err := json.Marshal(instanceInfo)
	if err != nil {
		return err
	}

	detail := &models.Detail{
		InstanceNumber: string(instanceData),
		RunningTime:    time.Now(),
	}

	err = dao.InsertDetail(detail)
	return err
}

func GetRecentDetail(beginTime, endTime time.Time) ([]models.Detail, error) {
	begin := beginTime.Format("2006-01-02 15:04:05")
	end := endTime.Format("2006-01-02 15:04:05")
	details, err := dao.GetDetailByTimePeriod(begin, end)
	if err != nil {
		return nil, err
	}

	return details, nil
}

func GetRecentInstanceDetail(hour int) ([]models.InstanceDetail, error) {
	endTime := time.Now()
	beginTime := endTime.Add(-time.Duration(hour) * time.Hour)
	beego.Info("Get the instances number from begin time", beginTime, "to end time", endTime)
	details, err := GetRecentDetail(beginTime, endTime)
	if err != nil {
		return nil, err
	}
	insDetails := make([]models.InstanceDetail, 0)
	for _, v := range details {
		bytes := *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&v.InstanceNumber))))
		number := make(map[string]int)
		err := json.Unmarshal(bytes, &number)
		if err != nil {
			return nil, err
		}
		insDetail := models.InstanceDetail{
			InstanceNumber: number,
			RunningTime:    GetCstTime(v.RunningTime),
		}
		insDetails = append(insDetails, insDetail)
	}
	return insDetails, nil
}

func GetLatestDetail() (map[string]int, error) {
	instanceInfo := make(map[string]int)
	allIns, err := instance.ListAllInstances()
	if err != nil {
		return nil, err
	}
	instanceInfo["total"] = len(allIns)

	providers, err := ListProviders()
	if err != nil {
		return nil, err
	}
	for _, p := range providers {
		providerIns, err := dao.GetInstancesByProvider(p)
		if err != nil {
			return nil, err
		}
		instanceInfo[p] = len(providerIns)
	}
	return instanceInfo, nil
}

func GetLatestInstanceDetail() ([]models.InstanceDetail, error) {
	details := make([]models.InstanceDetail, 0)
	instanceInfo, err := GetLatestDetail()
	if err != nil {
		return nil, err
	}

	instanceDetail := models.InstanceDetail{
		InstanceNumber: instanceInfo,
		RunningTime:    GetCstTime(time.Now()),
	}
	details = append(details, instanceDetail)
	return details, nil
}

func GetPastInstanceDetail(specificTime string) (*models.InstanceDetail, error) {
	theTime, err := time.ParseInLocation("2006-01-02 15:04:05", specificTime, time.Local)
	convertTime := theTime.Add(-time.Duration(8) * time.Hour)
	if err != nil {
		return nil, err
	}
	detail, err := dao.GetDetailByTime(convertTime)
	if err != nil {
		return nil, err
	}
	bytes := *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&detail.InstanceNumber))))
	number := make(map[string]int)
	err = json.Unmarshal(bytes, &number)
	if err != nil {
		return nil, err
	}
	insDetail := &models.InstanceDetail{
		InstanceNumber: number,
		RunningTime:    GetCstTime(detail.RunningTime),
	}
	return insDetail, nil
}

func GetCstTime(converTime time.Time) string {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return converTime.In(loc).Format("2006-01-02 15:04:05")
}

func InitInstanceDetailCron() {
	detailCron := future.NewCronbFuture("Instance detail task", future.DETAIL_INTERVAL, UpdateInstanceDetail)
	if detailCron != nil {
		future.Exec.Submit(detailCron)
	}
}

func ListProviders() ([]string, error) {
	providers, err := dao.GetProviders()
	if err != nil {
		return nil, err
	}
	return providers, nil

}
