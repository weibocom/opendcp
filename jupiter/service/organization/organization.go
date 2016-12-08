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
