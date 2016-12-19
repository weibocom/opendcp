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


package organization

import (
	"weibo.com/opendcp/jupiter/dao"
	"weibo.com/opendcp/jupiter/models"
)

func ListAll() ([]models.Organization, error) {
	organizations, err := dao.ListOrganizations()
	if err != nil {
		return nil, err
	}
	return organizations, nil
}

func GetOrganization(organizationId int64) (*models.Organization, error) {
	organization, err := dao.GetOrganization(organizationId)
	if err != nil {
		return nil, err
	}
	return organization, nil
}

func GetInstancesByOrganizationId(organizationId int64) ([]models.Instance, error) {
	instances := make([]models.Instance, 0)
	instancesInOrganization, err := dao.GetInstancesInOrganization(organizationId)
	if err != nil {
		return nil, err
	}
	for _, v := range instancesInOrganization {
		instances = append(instances, *(v.Instance))
	}
	return instances, nil
}

func New(organization *models.Organization) error {
	err := dao.InsertOrganization(organization)
	return err
}

func Delete(organizationId int64) error {
	err := dao.DeleteOrganization(organizationId)
	return err
}

func GetClusters(organizationId int64) ([]models.Cluster, error) {
	clusters, err := dao.GetClustersByOrganizationId(organizationId)
	if err != nil {
		return nil, err
	}
	return clusters, nil
}
