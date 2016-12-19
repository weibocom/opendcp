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
