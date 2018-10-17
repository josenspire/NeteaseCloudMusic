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
	models.WriteApiCache(i.Ctx, result)

	setupResponseCookies(i.Ctx.ResponseWriter, cookies)

	i.Data["json"] = result
	i.ServeJSON()
}

func (i *IndexController) UpdateProfile() {
	resParams := models.UserProfileParams{}

	i.Ctx.Input.Bind(&resParams.Nickname, "nickname")
	i.Ctx.Input.Bind(&resParams.Signature, "signature")
	i.Ctx.Input.Bind(&resParams.Gender, "gender")
	i.Ctx.Input.Bind(&resParams.Birthday, "birthday")
	i.Ctx.Input.Bind(&resParams.Province, "province")
	i.Ctx.Input.Bind(&resParams.City, "city")

	user := &models.User{UserProfileParams: resParams, Cookies: append(i.Ctx.Request.Cookies(), setupDefaultCookie()...)}
	result, cookies := user.UpdateProfile()
	models.WriteApiCache(i.Ctx, result)

	setupResponseCookies(i.Ctx.ResponseWriter, cookies)

	i.Data["json"] = result
	i.ServeJSON()
}

func (i *IndexController) GetPlayList() {
	playListParams := models.PlayListParams{}
	i.Ctx.Input.Bind(&playListParams.Uid, "uid")
	i.Ctx.Input.Bind(&playListParams.Limit, "limit")
	i.Ctx.Input.Bind(&playListParams.Offset, "offset")

	user := &models.User{Cookies: i.Ctx.Request.Cookies(), PlayListParams: playListParams}
	result, cookies := user.GetPlayList()

	models.WriteApiCache(i.Ctx, result)
	setupResponseCookies(i.Ctx.ResponseWriter, cookies)

	i.Data["json"] = result
	i.ServeJSON()
}

func (i *IndexController) GetPlayRecord() {
	playListParams := models.PlayListParams{}
	i.Ctx.Input.Bind(&playListParams.Uid, "uid")
	i.Ctx.Input.Bind(&playListParams.Type, "type")

	user := &models.User{Cookies: i.Ctx.Request.Cookies(), PlayListParams: playListParams}
	result, cookies := user.GetPlayRecord()

	models.WriteApiCache(i.Ctx, result)
	setupResponseCookies(i.Ctx.ResponseWriter, cookies)

	i.Data["json"] = result
	i.ServeJSON()
}

func (i *IndexController) GetPrivateLetter() {
	letterParams := models.LetterParams{}
	i.Ctx.Input.Bind(&letterParams.Uid, "uid")
	// i.Ctx.Input.Bind(&letterParams.Offset, "offset")
	// i.Ctx.Input.Bind(&letterParams.Limit, "limit")

	letter := &models.Letter{Cookies: i.Ctx.Request.Cookies(), LetterParams: letterParams}
	result := letter.GetPrivateLetter()

	models.WriteApiCache(i.Ctx, result)

	i.Data["json"] = result
	i.ServeJSON()
}
