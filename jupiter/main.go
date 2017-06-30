package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "weibo.com/opendcp/jupiter/docs"
	_ "weibo.com/opendcp/jupiter/routers"

	"github.com/astaxie/beego"
	"weibo.com/opendcp/jupiter/conf"
	"weibo.com/opendcp/jupiter/dao"
	"weibo.com/opendcp/jupiter/future"
	"weibo.com/opendcp/jupiter/service/account"
)

func initCostCron()  {
	costTask := future.NewCronbFuture("Peroidic compute cost", account.CheckCredit)
	if costTask != nil {
		future.Exec.Submit(costTask)
	}
}

func init()  {
	conf.InitConf()
	future.InitExec()
	dao.InitDB()
	initCostCron()
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.SetLogger("file", `{"filename":"logs/jupiter.log"}`)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
