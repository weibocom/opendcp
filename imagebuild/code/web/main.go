package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"os"
	"weibo.com/opendcp/imagebuild/code/env"
	"weibo.com/opendcp/imagebuild/code/util"
	"weibo.com/opendcp/imagebuild/code/web/models"
	_ "weibo.com/opendcp/imagebuild/code/web/models"
	_ "weibo.com/opendcp/imagebuild/code/web/routers"
)

func main() {
	args_num := len(os.Args)
	if args_num < 2 {
		fmt.Println("missing basedir or log path..")
		os.Exit(-1)
	}

	serverBaseDir := os.Args[1]
	log.Infof("Basedir: %s", serverBaseDir)

	configs := util.LoadConfig(serverBaseDir + "/globle_config/env")
	ip := ParseEnv("SERVER_IP", configs, true)
	port := ParseEnv("SERVER_PORT", configs, true)
	mysqlHost := ParseEnv("MYSQL_HOST", configs, true)
	mysqlPort := ParseEnv("MYSQL_PORT", configs, true)
	mysqlUser := ParseEnv("MYSQL_USER", configs, true)
	mysqlPassword := ParseEnv("MYSQL_PASSWORD", configs, true)
	logPath := ParseEnv("LOG_PATH", configs, false)
	pluginViewUrl := ParseEnv("PLUGIN_VIEW_URL", configs, true)
	extensionInterfaceUrl := ParseEnv("EXTENSION_INTERFACE_URL", configs, true)
	harborAddress := ParseEnv("HARBOR_ADDRESS", configs, true)
	harborUser := ParseEnv("HARBOR_USER", configs, true)
	harborPassword := ParseEnv("HARBOR_PASSWORD", configs, true)
	clusterlistAddress := ParseEnv("CLUSTERLIST_ADDRESS", configs, true)

	if logPath == "" {
		logPath = "/tmp/imagebuild.log"
	}

	util.LogInit(logPath)

	// init env
	env.InitEnv(harborAddress,
		harborUser,
		harborPassword,
		pluginViewUrl,
		extensionInterfaceUrl,
		logPath,
		ip,
		port,
		serverBaseDir,
		mysqlHost,
		mysqlPort,
		mysqlUser,
		mysqlPassword,
		clusterlistAddress)

	// init server
	models.AppServer.Init(ip, port)

	// register template function
	beego.AddFuncMap("startwith", util.StartWith)
	beego.AddFuncMap("endwith", util.EndWith)
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}

func ParseEnv(key string, configs map[string]string, need bool) string {
	//　首先从环境变量中获取
	value := os.Getenv(key)
	if value == "" {
		value = util.GetOrDefault(configs, key, "")
		if value == "" {
			log.Errorf("Property %s is empty", key)
			if need {
				os.Exit(-1)
			}
		}
	}

	return value
}
