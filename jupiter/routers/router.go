// @APIVersion 1.0.0
// @Title jupiter API
// @Description jupiter has a very cool tools to manage cloud server
// @Contact wenhui16
// @TermsOfServiceUrl http://weibo.com/opendcp/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"weibo.com/opendcp/jupiter/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/cluster",
			beego.NSInclude(
				&controllers.ClusterController{},
			),
		),
		beego.NSNamespace("/organization",
			beego.NSInclude(
				&controllers.OrganizationController{},
			),
		),
		beego.NSNamespace("/instance",
			beego.NSInclude(
				&controllers.InstanceController{},
			),
		),
		beego.NSNamespace("/slb",
			beego.NSInclude(
				&controllers.SlbController{},
			),
		),
		beego.NSNamespace("/account",
			beego.NSInclude(
				&controllers.AccountController{},
			),
		),
		beego.NSNamespace("/init",
			beego.NSInclude(
				&controllers.InitController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
