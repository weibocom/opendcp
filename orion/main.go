/*
 *  Copyright 2009-2016 Weibo, Inc.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */
package main

import (
	//"weibo.com/opendcp/orion/executor"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	. "weibo.com/opendcp/orion/models"
	_ "weibo.com/opendcp/orion/routers"
)

func main() {
	beego.Run()
}

func init() {
	err := beego.LoadAppConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}

	initOrm()

	beego.SetLogger("file", `{"filename":"logs/orion.log"}`)
}

func initOrm() {
	dbUrl := beego.AppConfig.String("database_url")
	dbpoolsizestr := beego.AppConfig.String("database_poolsize")
	if dbUrl == "" {
		panic("db_url not found in config...")
	}

	orm.Debug = true

	orm.RegisterDriver("mysql", orm.DRMySQL)

	dbpoolsize, _ := strconv.Atoi(dbpoolsizestr)
	orm.RegisterDataBase("default", "mysql", dbUrl, dbpoolsize)

	//register model
	orm.RegisterModel(&(Cluster{}), &(Service{}), &(Pool{}), &(Node{}))
	orm.RegisterModel(&(FlowImpl{}), &(Flow{}), &(FlowBatch{}), &(NodeState{}))
	orm.RegisterModel(&(RemoteStep{}), &(RemoteAction{}), &(RemoteActionImpl{}))
}
