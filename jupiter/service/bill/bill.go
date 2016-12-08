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

package bill

import (
	"time"
	"weibo.com/opendcp/jupiter/dao"
	"weibo.com/opendcp/jupiter/models"
)

func Bill(cluster *models.Cluster, hours int) error {
	bill, err := dao.GetBill(cluster.Id)
	if err != nil {
		return err
	}
	bill.Costs += hours
	bill.Credit -= hours
	err = dao.UpdateBill(bill)
	if err != nil {
		return err
	}
	return nil
}

func CanCreate(cluster *models.Cluster) bool {
	bill, err := dao.GetBill(cluster.Id)
	if err != nil {
		return false
	}
	if bill.Credit <= 0 {
		return false
	}
	return true
}

func GetUsageHours(instanceId string) (int, error) {
	instance, err := dao.GetInstance(instanceId)
	if err != nil {
		return 0, err
	}
	return time.Now().Hour() - instance.CreateTime.Hour() + 1, nil
}

func GetCosts(clusterId int64) (int, error) {
	bill, err := dao.GetBill(clusterId)
	if err != nil {
		return 0, err
	}
	return bill.Costs, nil
}

func GetCredit(clusterId int64) (int, error) {
	bill, err := dao.GetBill(clusterId)
	if err != nil {
		return 0, err
	}
	return bill.Credit, nil
}

func GetBill(clusterId int64) (*models.Bill, error) {
	bill, err := dao.GetBill(clusterId)
	if err != nil {
		return nil, err
	}
	return bill, nil
}

func IncreaseCredit(clusterId int64, hours int) (bool, error) {
	bill, err := dao.GetBill(clusterId)
	if err != nil {
		return false, err
	}
	bill.Credit += hours
	err = dao.UpdateBill(bill)
	if err != nil {
		return false, err
	}
	return true, nil
}

func InsertBill(cluster *models.Cluster) (bool, error) {
	var bill models.Bill
	bill.Cluster = cluster
	bill.Credit = 0
	bill.Costs = 0
	err := dao.InsertBill(&bill)
	if err != nil {
		return false, err
	}
	return true, nil
}
