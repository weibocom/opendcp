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

package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"weibo.com/opendcp/jupiter/service/cluster"
	"fmt"
)

// Operations about cluster
type InitController struct {
	BaseController
}


// @Title .
// @Description.
// @router / [get]
func (initController *InitController) InitDB() {
	resp := ApiResponse{}

	bizId := initController.Ctx.Input.Header("X-Biz-ID")
	bid, err := strconv.Atoi(bizId)
	if bizId=="" || err != nil {
		beego.Error("Get X-Biz-ID err!")
		initController.RespInputError()
		return
	}

	credit,_ := beego.AppConfig.Int("credit")
	provider := "'aliyun'"
	key_id := "''"
	key_secret := "''"

	sqlAccount := "insert into account(biz_id,credit,provider,key_id,key_secret) values(%d,%d,%s,%s,%s)"

	sql3 :=
		"insert into cluster(`name`,provider,lastest_part_num,`desc`,create_time,delete_time," +
			"cpu,ram,instance_type,image_id,post_script,key_name,network_id,zone_id,system_disk_category," +
			"data_disk_size,data_disk_num,data_disk_category,biz_id) " +
		"values('16Core16G经典网','aliyun',0,'',NOW(),NULL,16,16,'ecs.c2.medium','centos7u2_64_40G_cloudinit_20160728.raw','','key',1,1,'cloud_efficiency',100,1,'cloud_efficiency',%d)"

	sql2 :=
		"insert into cluster(`name`,provider,lastest_part_num,`desc`,create_time,delete_time," +
			"cpu,ram,instance_type,image_id,post_script,key_name,network_id,zone_id,system_disk_category," +
			"data_disk_size,data_disk_num,data_disk_category,biz_id) " +
			"values('4Core8G经典网','aliyun',0,'',NOW(),NULL,4,8,'ecs.n2.large','centos7u2_64_40G_cloudinit_20160728.raw','','key',1,1,'cloud_efficiency',100,1,'cloud_efficiency',%d)"
	sql1 :=
		"insert into cluster(`name`,provider,lastest_part_num,`desc`,create_time,delete_time," +
			"cpu,ram,instance_type,image_id,post_script,key_name,network_id,zone_id,system_disk_category," +
			"data_disk_size,data_disk_num,data_disk_category,biz_id) " +
			"values('1Core1G经典网','aliyun',0,'',NOW(),NULL,1,1,'ecs.n1.tiny','centos7u2_64_40G_cloudinit_20160728.raw','','key',1,1,'cloud_efficiency',100,1,'cloud_efficiency',%d)"


	sqlCluster := make([]string,3)
	sqlCluster[0] = sql1
	sqlCluster[1] = sql2
	sqlCluster[2] = sql3


	sqlBill := "insert into bill(cluster_id,costs,credit) values(%d,0,0)"
	delBill := "delete from bill where cluster_id=%d"


	deleteSql := "delete from %s where biz_id=%d"



	//删除数据
	_,err =cluster.OperateBysql(fmt.Sprintf(deleteSql,"cluster",bid))
	if err != nil {
		beego.Error("delete data from cluster err: ", err)
		initController.RespServiceError(err)
		return
	}

	_,err =cluster.OperateBysql(fmt.Sprintf(deleteSql,"account",bid))
	if err != nil {
		beego.Error("delete data from account err: ", err)
		initController.RespServiceError(err)
		return
	}

	//插入数据
	for _,sql := range sqlCluster {
		id64,err := cluster.OperateBysql(fmt.Sprintf(sql,bid))
		if err != nil {
			beego.Error("insert data for cluster err: ", err)
			initController.RespServiceError(err)
			return
		}
		id := int(id64)

		cluster.OperateBysql(fmt.Sprintf(delBill,id))

		_,err = cluster.OperateBysql(fmt.Sprintf(sqlBill,id))
		if err != nil {
			beego.Error("insert data for bill err: ", err)
			initController.RespServiceError(err)
			return
		}

	}

	_,err = cluster.OperateBysql(fmt.Sprintf(sqlAccount,bid,credit,provider,key_id,key_secret))
	if err != nil {
		beego.Error("insert data for account err: ", err)
		initController.RespServiceError(err)
		return
	}

	resp.Content = nil
	initController.ApiResponse = resp
	initController.Status = SERVICE_SUCCESS
	initController.RespJsonWithStatus()
}


