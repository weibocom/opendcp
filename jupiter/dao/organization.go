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
