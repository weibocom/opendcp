// Copyright 2016 Weibo Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use sf file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package future

import (
	"time"
	"weibo.com/opendcp/jupiter/conf"
	"weibo.com/opendcp/jupiter/dao"
	"weibo.com/opendcp/jupiter/logstore"
	"weibo.com/opendcp/jupiter/models"
	"weibo.com/opendcp/jupiter/provider"
	"weibo.com/opendcp/jupiter/service/instance"
	"fmt"
)

const (
	INTERVAL  = 60
	TIME4WAIT = 3
)

type StartFuture struct {
	InstanceId    string
	ProviderName  string
	AutoInit      bool
	Ip            string
	CorrelationId string
}

func NewStartFuture(instanceId string, providerName string, autoInit bool, ip, correlationId string) *StartFuture {
	return &StartFuture{
		InstanceId:    instanceId,
		ProviderName:  providerName,
		AutoInit:      autoInit,
		Ip:            ip,
		CorrelationId: correlationId,
	}
}

func (sf *StartFuture) Run() error {
	providerDriver, err := provider.New(sf.ProviderName)
	fmt.Println("get the provider")
	if err != nil {
		return err
	}
	logstore.Info(sf.CorrelationId, sf.InstanceId, "----- Begin start instance in future -----")
	if(sf.ProviderName=="aliyun") {
		for j := 0; j < INTERVAL; j++ {
			logstore.Info(sf.CorrelationId, sf.InstanceId, "wait for instance", sf.InstanceId, "to stop:", j)
			if providerDriver.WaitForInstanceToStop(sf.InstanceId) {
				break
			}
			time.Sleep(TIME4WAIT * time.Second)
		}
		isStart, err := providerDriver.Start(sf.InstanceId)
		if err != nil {
			return err
		}
		logstore.Info(sf.CorrelationId, sf.InstanceId, "Is the machine start?", isStart)
	}
	fmt.Println("get Instance")
	ins, err := providerDriver.GetInstance(sf.InstanceId)
	if err != nil {
		return err
	}
	// 支持专有网和经典网
	if len(ins.PrivateIpAddress) > 0 {
		sf.Ip = ins.PrivateIpAddress
		if err := dao.UpdateInstancePrivateIp(ins.InstanceId, ins.PrivateIpAddress); err != nil {
			return err
		}
	} else {
		fmt.Println("allocate Ip")
		publicIpAddress, err := providerDriver.AllocatePublicIpAddress(sf.InstanceId)
		if err != nil {
			return err
		}
		fmt.Println("get the Ip")
		sf.Ip = publicIpAddress
		if err := dao.UpdateInstancePublicIp(ins.InstanceId, publicIpAddress); err != nil {
			return err
		}
	}
	fmt.Println("allocated IpAd")
	for i := 0; i < 60; i++ {
		time.Sleep(10 * time.Second)
		logstore.Info(sf.CorrelationId, sf.InstanceId, "Wati for instance", sf.InstanceId, "to start", i)
		if providerDriver.WaitToStartInstance(sf.InstanceId) {
			break
		}
	}
	logstore.Info(sf.CorrelationId, sf.InstanceId, "Finished to start instance:", sf.InstanceId, sf.Ip)
	return nil
}

func (sf *StartFuture) Success() {
	dao.UpdateInstanceStatus(sf.Ip, models.Initing)
	/*logstore.Info(sf.CorrelationId, sf.InstanceId, "store ssh key: ", sf.InstanceId, sf.Ip)*/
	//sshErr := instance.StartSshService(sf.InstanceId, sf.Ip, conf.Config.Password, sf.CorrelationId)
	//if sshErr != nil {
	//logstore.Error(sf.CorrelationId, sf.InstanceId, "ssh instance: ", sf.InstanceId, "failed: ", sshErr)
	//dao.UpdateInstanceStatus(sf.Ip, models.InitTimeout)
	//return
	/*}*/
	logstore.Info(sf.CorrelationId, sf.InstanceId, "StartFuture success: ", sf.InstanceId, sf.Ip)
	//roles := []string{
	//conf.Config.Ansible.DefaultRole,
	//}
	if sf.AutoInit {
		//Exec.Submit(NewAnsibleTaskFuture(sf.InstanceId, sf.Ip, roles, sf.CorrelationId))
		instance.ManageDev(sf.Ip, conf.Config.Password, sf.InstanceId, sf.CorrelationId)
	}
}

func (sf *StartFuture) Failure(err error) {
	logstore.Error(sf.CorrelationId, sf.InstanceId, "StartFuture - ", sf.Ip, ":", err)
}

func (sf *StartFuture) ShutDown() {
}
