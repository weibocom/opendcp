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



package routers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"strings"
	"weibo.com/opendcp/imagebuild/code/web/controllers"
)

func init() {

	// Initialized language type list.
	langs := strings.Split(beego.AppConfig.String("lang"), "|")
	log.Infof("langs:%s", langs)
	for _, lang := range langs {
		log.Infof("Loading language:%s", lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			log.Error("Fail to set message file: " + err.Error())
			return
		}
	}

	// 页面
	beego.Router("/view/project", &controllers.ProjectViewController{})
	beego.Router("/view/config", &controllers.ConfigViewController{})
	beego.Router("/view/plugin", &controllers.PluginViewController{})
	beego.Router("/view/clone", &controllers.CloneViewController{})
	beego.Router("/view/info", &controllers.InfoViewController{})
	beego.Router("/view/config/independent", &controllers.IndependentConfigViewController{})

	// beego.Router("/view/list", &controllers.ProjectListViewController{})

	// 接口
	beego.Router("/api/project/new", &controllers.ProjectNewController{})
	beego.Router("/api/project/delete", &controllers.ProjectDeleteController{})
	beego.Router("/api/project/clone", &controllers.ProjectCloneController{})
	beego.Router("/api/project/info", &controllers.ProjectInfoController{})
	beego.Router("/api/project/list", &controllers.ProjectListController{})
	beego.Router("/api/project/save", &controllers.ConfigSaveController{})
	beego.Router("/api/project/build", &controllers.BuildImageController{})
	beego.Router("/api/project/buildHistory", &controllers.BuildImageHistoryController{})

	beego.Router("/api/plugin/extension/interface", &controllers.PluginExtensionInterfaceController{})

}
