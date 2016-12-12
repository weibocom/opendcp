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
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	. "weibo.com/opendcp/orion/models"
)

func main() {
	orm.RunCommand()
}

func init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		beego.Error(err)
	}

	orm.RegisterDataBase("default", "mysql", "root:xxxxx@tcp(localhost:3306)/orion?charset=utf8")

	orm.RegisterModel(&(Cluster{}), &(Service{}), &(Pool{}), &(Node{}))
	orm.RegisterModel(&(FlowImpl{}), &(Flow{}), &(FlowBatch{}), &(NodeState{}))
	orm.RegisterModel(&(RemoteStep{}), &(RemoteAction{}), &(RemoteActionImpl{}))
}
