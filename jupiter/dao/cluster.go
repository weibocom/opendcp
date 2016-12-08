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

func GetClusterById(clusterId int64) (*models.Cluster, error) {
	o := GetOrmer()
	var cluster models.Cluster
	err := o.QueryTable(CLUSTER_TABLE).RelatedSel().Filter("id", clusterId).One(&cluster)
	if err != nil {
		return nil, err
	}
	return &cluster, nil
}

func InsertCluster(cluster *models.Cluster) (int64, error) {
	o := GetOrmer()
	zoneId, err := InsertZone(cluster.Zone)
	if err != nil {
		return 0, err
	}
	networkId, err := InsertNetwork(cluster.Network)
	if err != nil {
		return 0, err
	}
	cluster.Network.Id = networkId
	cluster.Zone.Id = zoneId
	id, err := o.Insert(cluster)
	return id, err
}

func InsertNetwork(network *models.Network) (int64, error) {
	o := GetOrmer()
	id, err := o.Insert(network)
	if err != nil {
		var networkModel models.Network
		o.QueryTable(NETWORK_TABLE).Filter("subnet_id", network.SubnetId).
			Filter("security_group", network.SecurityGroup).
			Filter("internet_charge_type", network.InternetChargeType).
			Filter("internet_max_bandwidth_out", network.InternetMaxBandwidthOut).RelatedSel().One(&networkModel)
		id = networkModel.Id
	}
	return id, nil
}

func InsertZone(zone *models.Zone) (int64, error) {
	o := GetOrmer()
	id, err := o.Insert(zone)
	if err != nil {
		var zoneModel models.Zone
		o.QueryTable(ZONE_TABLE).Filter("zone_name", zone.ZoneName).RelatedSel().One(&zoneModel)
		id = zoneModel.Id
	}
	return id, nil
}

func DeleteCluster(clusterId int64) (bool, error) {
	o := GetOrmer()
	_, err := o.QueryTable(CLUSTER_TABLE).Filter("id", clusterId).Delete()
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetClusters() ([]models.Cluster, error) {
	o := GetOrmer()
	var clusters []models.Cluster
	_, err := o.QueryTable(CLUSTER_TABLE).RelatedSel().OrderBy("-id").All(&clusters)
	if err != nil {
		return nil, err
	}
	return clusters, nil
}
