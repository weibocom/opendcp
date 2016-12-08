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
