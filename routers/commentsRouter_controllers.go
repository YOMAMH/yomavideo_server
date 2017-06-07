package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["yomavideo_server/controllers:HomeController"] = append(beego.GlobalControllerRouter["yomavideo_server/controllers:HomeController"],
		beego.ControllerComments{
			Method: "Carouses",
			Router: `/carouses/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["yomavideo_server/controllers:UserController"] = append(beego.GlobalControllerRouter["yomavideo_server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/user/singoIn/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["yomavideo_server/controllers:UserController"] = append(beego.GlobalControllerRouter["yomavideo_server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/user/login/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
