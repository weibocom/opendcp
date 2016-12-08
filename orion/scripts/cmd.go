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
