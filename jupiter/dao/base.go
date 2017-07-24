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
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"net"
	"sync"
	"time"
)

const BILL_TABLE = "bill"
const INSTANCE_TABLE = "instance"
const ORGANIZATION_TABLE = "organization"
const CLUSTER_TABLE = "cluster"
const INSTANCE_ORGANIZATION_TABLE = "instance_organization"
const NETWORK_TABLE = "network"
const ZONE_TABLE = "zone"
const DETAIL_TABLE = "detail"

var globalOrm orm.Ormer
var once sync.Once

//InitDB initializes the database
func InitDB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	username := beego.AppConfig.String("mysqluser")
	password := beego.AppConfig.String("mysqlpass")
	addr := beego.AppConfig.String("mysqladdr")
	port := beego.AppConfig.String("mysqlport")

	//log.Debugf("db url: %s:%s, db user: %s", addr, port, username)
	orm.Debug = true
	dbStr := username + ":" + password + "@tcp(" + addr + ":" + port + ")/jupiter?charset=utf8&loc=Local"
	ch := make(chan int, 1)
	go func() {
		var err error
		var c net.Conn
		for {
			c, err = net.DialTimeout("tcp", addr+":"+port, 20*time.Second)
			if err == nil {
				c.Close()
				ch <- 1
			} else {
				//log.Errorf("failed to connect to db, retry after 2 seconds :%v", err)
				time.Sleep(2 * time.Second)
			}
		}
	}()
	select {
	case <-ch:
	case <-time.After(60 * time.Second):
		panic("Failed to connect to DB after 60 seconds")
	}
	err := orm.RegisterDataBase("default", "mysql", dbStr)
	if err != nil {
		panic(err)
	}
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
	}
	orm.Debug = true
}

// GetOrmer :set ormer singleton
func GetOrmer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}
