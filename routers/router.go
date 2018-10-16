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
	"NeteaseCloudMusic/models"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "get:RenderIndex")

	idx := beego.NewNamespace("/test",
		// api cache checking
		beego.NSBefore(models.ReadApiCache),
		beego.NSRouter("/cellphone", &controllers.IndexController{}, "get:CellphoneLogin"),
		beego.NSRouter("/update", &controllers.IndexController{}, "get:UpdateProfile"),
		beego.NSRouter("/playList", &controllers.IndexController{}, "get:GetPlayList"),
	)

	ns := beego.NewNamespace("/v1/api",
		// api cache checking
		beego.NSBefore(models.ReadApiCache),

		beego.NSNamespace("/user",
			beego.NSRouter("/cellphone", &controllers.UserController{}, "post:CellphoneLogin"),
			beego.NSRouter("/refreshLogin", &controllers.UserController{}, "get:RefreshLogin"),
			// TODO: uncompleted
			beego.NSRouter("/detail", &controllers.UserController{}, "get:GetUserDetail"),

			beego.NSRouter("/subcount", &controllers.UserController{}, "get:GetUserSubscriptCount"),
			beego.NSRouter("/update", &controllers.UserController{}, "post:UpdateProfile"),

			// play
			beego.NSRouter("/playList", &controllers.UserController{}, "post:GetPlayList"),
		),
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/search",
			beego.NSRouter("/hot", &controllers.SearchController{}, "post:Search"),
		),
	)

	// register namespace
	beego.AddNamespace(ns)
	beego.AddNamespace(idx)
}
