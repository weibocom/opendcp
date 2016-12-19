package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "weibo.com/opendcp/jupiter/docs"
	_ "weibo.com/opendcp/jupiter/routers"

	"github.com/astaxie/beego"
	"weibo.com/opendcp/jupiter/conf"
	"weibo.com/opendcp/jupiter/dao"
	"weibo.com/opendcp/jupiter/future"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.SetLogger("file", `{"filename":"logs/jupiter.log"}`)
	conf.InitConf()
	future.InitExec()
	dao.InitDB()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
