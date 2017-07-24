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

package main

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
	. "weibo.com/opendcp/orion/models"
	_ "weibo.com/opendcp/orion/routers"
	"weibo.com/opendcp/orion/sched"
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

	if err := sched.Initial(); err != nil {
		panic(err)
	}
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
	orm.RegisterModel(&(Cluster{}), &(Service{}), &(Pool{}), &(Node{}), &(Logs{}))
	orm.RegisterModel(&(FlowImpl{}), &(Flow{}), &(FlowBatch{}), &(NodeState{}))
	orm.RegisterModel(&(RemoteStep{}), &(RemoteAction{}), &(RemoteActionImpl{}))
	orm.RegisterModel(&(CronItem{}), &(DependItem{}), &(ExecTask{}))
}
