// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"hometools/ht_server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/family",
			beego.NSInclude(
				&controllers.FamilyController{},
			),
		),
		beego.NSRouter("/user",&controllers.UserController{}),
		beego.NSRouter("/family",&controllers.FamilyController{}),
	)
	beego.AddNamespace(ns)
}
