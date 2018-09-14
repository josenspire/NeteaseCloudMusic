package controllers

import (
	"fmt"
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

func (i *IndexController) RequestTesting() {
	input := i.Ctx.Input.Cookie

	fmt.Println(input)

	i.Data["json"] = input
	i.ServeJSON()
}
