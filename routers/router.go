// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"NeteaseCloudMusic/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "get:RenderIndex")

	beego.Router("/test", &controllers.IndexController{}, "get:RequestTesting")

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			// beego.NSInclude(
			// 	&controllers.UserController{},
			// ),
			beego.NSRouter("/cellphone", &controllers.UserController{}, "post:CellphoneLogin"),
		),
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
