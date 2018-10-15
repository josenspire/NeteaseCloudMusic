package controllers

import (
	"NeteaseCloudMusic/models"
	"github.com/astaxie/beego"
)

// Operations about object
type IndexController struct {
	beego.Controller
}

// @Title Get
// @Description Index page
// @Success 200
// @Failure 404 page not found
// @router / [get]
func (i *IndexController) RenderIndex() {
	i.Data["Website"] = "Beego.me"
	i.Data["Email"] = "josenspire@gmail.com"
	i.TplName = "index.tpl"
}

func (i *IndexController) CellphoneLogin() {
	resParams := models.CellphoneLoginParams{}

	i.Ctx.Input.Bind(&resParams.Phone, "phone")
	i.Ctx.Input.Bind(&resParams.Password, "password")
	i.Ctx.Input.Bind(&resParams.RememberLogin, "rememberLogin")

	user := &models.User{CellphoneLoginParams: resParams, Cookies: append(i.Ctx.Request.Cookies(), setupDefaultCookie()...)}
	result, cookies := user.CellphoneLogin()
	WriteApiCache(i.Ctx, result)

	setupResponseCookies(i.Ctx.ResponseWriter, cookies)

	i.Data["json"] = result
	i.ServeJSON()
}
