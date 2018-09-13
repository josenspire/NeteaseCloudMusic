package controllers

import "github.com/astaxie/beego"

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
