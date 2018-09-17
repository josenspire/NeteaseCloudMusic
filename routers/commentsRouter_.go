package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:IndexController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "RenderIndex",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"],
		beego.ControllerComments{
			Method:           "CellphoneLogin",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

}
