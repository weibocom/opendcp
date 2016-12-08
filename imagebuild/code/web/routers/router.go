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
