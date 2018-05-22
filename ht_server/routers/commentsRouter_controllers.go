package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["hometools/ht_server/controllers:FamilyController"] = append(beego.GlobalControllerRouter["hometools/ht_server/controllers:FamilyController"],
		beego.ControllerComments{
			Method: "GetFamily",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hometools/ht_server/controllers:FamilyController"] = append(beego.GlobalControllerRouter["hometools/ht_server/controllers:FamilyController"],
		beego.ControllerComments{
			Method: "GetMembers",
			Router: `/:id/members`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hometools/ht_server/controllers:FamilyController"] = append(beego.GlobalControllerRouter["hometools/ht_server/controllers:FamilyController"],
		beego.ControllerComments{
			Method: "DelMember",
			Router: `/:id/members`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hometools/ht_server/controllers:FamilyController"] = append(beego.GlobalControllerRouter["hometools/ht_server/controllers:FamilyController"],
		beego.ControllerComments{
			Method: "AddMember",
			Router: `/members/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hometools/ht_server/controllers:UserController"] = append(beego.GlobalControllerRouter["hometools/ht_server/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUser",
			Router: `/:openid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hometools/ht_server/controllers:UserController"] = append(beego.GlobalControllerRouter["hometools/ht_server/controllers:UserController"],
		beego.ControllerComments{
			Method: "FamilyList",
			Router: `/familylist/:userid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["hometools/ht_server/controllers:UserController"] = append(beego.GlobalControllerRouter["hometools/ht_server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login/:openid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
