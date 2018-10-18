package controllers

import (
	"NeteaseCloudMusic/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type ArtistController struct {
	beego.Controller
}

// @Title GetArtistList
// @Description Get User's Artist List
// @Param    cat      	query    string    true        "category code, default is `1001`"
// @Param    initial    query    string    true        "The expected return structure of name is sorted according to this parameter. e.g. `initial=a`"
// @Param    total      query    string    false       "unknow, `true` or `false`, default is `true`"
// @Param    offset     query    int       false       "page offset, default is 0"
// @Param    limit      query    int       false       "return data limit, default is 30 items"
// @Success 200 {string}
// @router /list [get]
func (a *ArtistController) GetArtistList() {
	listParams := models.ListParams{Total: "true", Offset: 0, Limit: 30}
	artistListParams := models.ArtistListParams{CategoryCode: "1001", Initial: 65}

	if cat := a.Input().Get("cat"); cat != "" {
		artistListParams.CategoryCode = cat
	}
	if initial := a.Input().Get("initial"); initial != "" {
		// Note: string to unicode char:  `runes := []rune(str)`
		runes := []rune(strings.ToUpper(initial))
		artistListParams.Initial = runes[0]
	}
	if offset := a.Input().Get("offset"); offset != "" {
		listParams.Offset, _ = strconv.Atoi(offset)
	}
	if limit := a.Input().Get("limit"); limit != "" {
		listParams.Limit, _ = strconv.Atoi(limit)
	}
	if total := a.Input().Get("total"); total == "" {
		listParams.Total = "false"
	}
	artist := &models.Artist{Cookies: a.Ctx.Request.Cookies(), ArtistListParams: artistListParams, ListParams: listParams}

	result := artist.GetArtistList()
	models.WriteApiCache(a.Ctx, result)

	a.Data["json"] = result
	a.ServeJSON()
}

// @Title SubscriptArtist
// @Description Subscript artist
// @Param    artistId      	query    string    true        "artist id, e.g.`artistId=1204010`"
// @Success 200 {string}
// @router /sub [get]
func (a *ArtistController) SubscriptArtist() {
	if artistId := a.Input().Get("artistId"); artistId == "" {
		a.Data["json"] = "Params error, please check your request"
		a.ServeJSON()
	} else {
		artist := &models.Artist{Cookies: a.Ctx.Request.Cookies()}
		result := artist.SubscriptArtist(artistId)
		models.WriteApiCache(a.Ctx, result)

		a.Data["json"] = result
		a.ServeJSON()
	}
}

// @Title UnSubscriptArtist
// @Description Cancel Subscript Artist
// @Param    artistId      	query    string    true        "artist id, e.g.`artistId=1204010`"
// @Success 200 {string}
// @router /unsub [get]
func (a *ArtistController) UnSubscriptArtist() {
	if artistId := a.Input().Get("artistId"); artistId == "" {
		a.Data["json"] = "Params error, please check your request"
		a.ServeJSON()
	} else {
		artist := &models.Artist{Cookies: a.Ctx.Request.Cookies()}
		result := artist.UnSubscriptArtist(artistId)
		models.WriteApiCache(a.Ctx, result)

		a.Data["json"] = result
		a.ServeJSON()
	}
}

// @Title GetSubscriptArtistList
// @Description Cancel Subscript Artist
// @Param    total      query    string    false       "unknow, `true` or `false`, default is `true`"
// @Param    offset     query    int       false       "page offset, default is 0"
// @Param    limit      query    int       false       "return data limit, default is 30 items"
// @Success 200 {string}
// @router /sublist [get]
func (a *ArtistController) GetSubscriptArtistList() {
	listParams := models.ListParams{Limit: 30, Total: `true`}
	artist := &models.Artist{Cookies: a.Ctx.Request.Cookies(), ListParams: listParams}
	result := artist.GetSubscriptArtistList()
	models.WriteApiCache(a.Ctx, result)

	a.Data["json"] = result
	a.ServeJSON()

}
