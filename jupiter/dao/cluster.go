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
	"fmt"
	"time"
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

func GetClustersByProvider(providerName string) ([]models.Cluster, error) {
	o := GetOrmer()
	var clusters []models.Cluster
	_, err := o.QueryTable(CLUSTER_TABLE).RelatedSel().Filter("provider", providerName).All(&clusters)
	if err != nil {
		return nil, err
	}
	return clusters, nil
}

func InsertDetail(detail *models.Detail) error {
	o := GetOrmer()
	_, err := o.Insert(detail)
	return err
}

func GetDetailByTimePeriod(begin, end string)  (details []models.Detail, err error) {
	o := GetOrmer()
	sql := fmt.Sprintf("SELECT * FROM %s WHERE RUNNING_TIME BETWEEN '%s' AND '%s' ORDER BY RUNNING_TIME", DETAIL_TABLE, begin, end)
	_, err = o.Raw(sql).QueryRows(&details)
	if err != nil {
		return nil, err
	}

	return details, nil
}

func GetDetailByTime(specificTime time.Time) (*models.Detail, error)  {
	o := GetOrmer()
	timeStr := specificTime.Format("2006-01-02 15:04:05")
	var detail models.Detail
	err := o.QueryTable(DETAIL_TABLE).Filter("running_time", timeStr).One(&detail)
	if err != nil {
		return nil, err
	}
	return &detail, err
}

func GetProviders() ([]string, error) {
	o := GetOrmer()
	var providers []string
	sql := fmt.Sprintf("SELECT DISTINCT PROVIDER FROM %s ", CLUSTER_TABLE)
	_, err := o.Raw(sql).QueryRows(&providers)
	if err != nil {
		return nil, err
	}
	return providers,nil
}
