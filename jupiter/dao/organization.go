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
)

func GetBill(clusterId int64) (*models.Bill, error) {
	o := GetOrmer()
	var bill models.Bill
	err := o.QueryTable(BILL_TABLE).RelatedSel().Filter("cluster_id", clusterId).One(&bill)
	if err != nil {
		return nil, err
	}
	return &bill, nil
}

func InsertBill(bill *models.Bill) error {
	o := GetOrmer()
	_, err := o.Insert(bill)
	if err != nil {
		return err
	}
	return nil
}

func UpdateBill(bill *models.Bill) error {
	o := GetOrmer()
	_, err := o.Update(bill)
	if err != nil {
		return err
	}
	return nil
}

func ListOrganizations() ([]models.Organization, error) {
	o := GetOrmer()
	var organizations []models.Organization
	_, err := o.QueryTable(ORGANIZATION_TABLE).RelatedSel().OrderBy("create_time").All(&organizations)
	return organizations, err
}

func GetOrganization(organizationId int64) (*models.Organization, error) {
	o := GetOrmer()
	var organization models.Organization
	err := o.QueryTable(ORGANIZATION_TABLE).RelatedSel().Filter("id", organizationId).One(&organization)
	if err != nil {
		return nil, err
	}
	return &organization, nil
}

func GetInstancesInOrganization(organizationId int64) ([]models.InstanceOrganization, error) {
	o := GetOrmer()
	var instancesInOrganization []models.InstanceOrganization
	_, err := o.QueryTable(INSTANCE_ORGANIZATION_TABLE).RelatedSel().Filter("organization_id", organizationId).All(&instancesInOrganization)
	if err != nil {
		return nil, err
	}
	return instancesInOrganization, nil
}

func InsertOrganization(organization *models.Organization) error {
	o := GetOrmer()
	_, err := o.Insert(organization)
	return err
}

func DeleteOrganization(organizationId int64) error {
	o := GetOrmer()
	_, err := o.QueryTable(ORGANIZATION_TABLE).Filter("id", organizationId).Delete()
	return err
}

func GetClustersByOrganizationId(organizationId int64) ([]models.Cluster, error) {
	o := GetOrmer()
	var clusters []models.Cluster
	_, err := o.QueryTable(CLUSTER_TABLE).Filter("organization_id", organizationId).All(&clusters)
	if err != nil {
		return nil, err
	}
	return clusters, nil
}
