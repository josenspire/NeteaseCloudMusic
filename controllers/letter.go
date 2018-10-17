package controllers

import (
	"NeteaseCloudMusic/models"
	"github.com/astaxie/beego"
)

type LetterController struct {
	beego.Controller
}

// @Title GetPrivateLetter
// @Description Get User's Play List
// @Param    uid      query    string    true        "User id"
// @Success 200 {string}
// @router /privateLetter [post]
func (l *LetterController) GetPrivateLetter() {
	letterParams := models.LetterParams{}
	l.Ctx.Input.Bind(&letterParams.Uid, "uid")
	// l.Ctx.Input.Bind(&letterParams.Offset, "offset")
	// l.Ctx.Input.Bind(&letterParams.Limit, "limit")

	letter := &models.Letter{Cookies: l.Ctx.Request.Cookies(), LetterParams: letterParams}
	result := letter.GetPrivateLetter()

	models.WriteApiCache(l.Ctx, result)

	l.Data["json"] = result
	l.ServeJSON()
}

func (l *LetterController) SendPrivateText() {
	sendLetterParams := models.SendLetterParams{Type: `text`}
	l.Ctx.Input.Bind(&sendLetterParams.Msg, "msg")
	sendLetterParams.UserIds = "[" + l.Ctx.Input.Query("userIds") + "]"

	letter := &models.Letter{Cookies: l.Ctx.Request.Cookies(), SendLetterParams: sendLetterParams}
	result := letter.SendPrivateText()

	models.WriteApiCache(l.Ctx, result)

	l.Data["json"] = result
	l.ServeJSON()
}

func (l *LetterController) SendWithPlaylist() {
	sendLetterParams := models.SendLetterParams{Type: `playlist`}
	l.Ctx.Input.Bind(&sendLetterParams.Id, "playlistId")
	l.Ctx.Input.Bind(&sendLetterParams.Msg, "msg")
	sendLetterParams.UserIds = "[" + l.Ctx.Input.Query("userIds") + "]"

	letter := &models.Letter{Cookies: l.Ctx.Request.Cookies(), SendLetterParams: sendLetterParams}
	result := letter.SendWithPlaylist()

	models.WriteApiCache(l.Ctx, result)

	l.Data["json"] = result
	l.ServeJSON()
}
