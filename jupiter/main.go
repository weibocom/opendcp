package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "weibo.com/opendcp/jupiter/docs"
	_ "weibo.com/opendcp/jupiter/routers"

	"github.com/astaxie/beego"
	"weibo.com/opendcp/jupiter/conf"
	"weibo.com/opendcp/jupiter/dao"
	"weibo.com/opendcp/jupiter/future"
	"weibo.com/opendcp/jupiter/service/cluster"
	"weibo.com/opendcp/jupiter/service/task"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.SetLogger("file", `{"filename":"logs/jupiter.log"}`)
	conf.InitConf()
	future.InitExec()
	cluster.InitInstanceDetailCron()
	dao.InitDB()
	task.InitInstanceTask()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
