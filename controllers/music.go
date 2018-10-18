package controllers

import (
	"NeteaseCloudMusic/models"
	"github.com/astaxie/beego"
	"strconv"
)

type MusicController struct {
	beego.Controller
}

// @Title GetMusicUrl
// @Description Query music url by music id
// @Params ids   	query    string    false        "support multi ids, e.g. `ids=405998841,33894312`"
// @Params br  	query    string    false        "result data offset, default is `0`"
// @Success 200 {string}
// @router /musicUrl [get]
func (m *MusicController) GetMusicUrl() {
	musicUrlParams := models.MusicUrlParams{Bitrate: `999000`}
	if ids := m.Input().Get("ids"); ids == "" {
		m.Data["json"] = "Params error, please check your request"
		m.ServeJSON()
	} else {
		musicUrlParams.Ids = "[" + ids + "]"
	}
	music := models.Music{Cookies: m.Ctx.Request.Cookies(), MusicUrlParams: musicUrlParams}
	result := music.GetMusicUrl()
	models.WriteApiCache(m.Ctx, result)

	m.Data["json"] = result
	m.ServeJSON()
}

// @Title SearchMusic
// @Description Search music
// @Params keywords   	 query    string    true    "search criteria keywords"
// @Params type		   	 query    int	    false   "search type, 默认为 1 即单曲 , 取值意义 : 1: 单曲, 10: 专辑, 100: 歌手, 1000: 歌单, 1002: 用户, 1004: MV, 1006: 歌词, 1009: 电台, 1014: 视频"
// @Params offset   	 query    int       false   "items offset"
// @Params limit   	 	 query    int       false   "items limit"
// @Success 200 {string}
// @router /search [get]
func (m *MusicController) SearchMusic() {
	searchParams := models.SearchParams{Offset: 0, Limit: 30, Type: 1}
	if keywords := m.Input().Get("keywords"); keywords == "" {
		m.Data["json"] = "Params error, please check your request"
		m.ServeJSON()
	} else {
		searchParams.Keywords = keywords
	}
	if mType := m.Input().Get("type"); mType != "" {
		searchParams.Type, _ = strconv.Atoi(mType)
	}
	if offset := m.Input().Get("offset"); offset != "" {
		searchParams.Offset, _ = strconv.Atoi(offset)
	}
	if limit := m.Input().Get("limit"); limit != "" {
		searchParams.Limit, _ = strconv.Atoi(limit)
	}
	music := models.Music{Cookies: m.Ctx.Request.Cookies(), SearchParams: searchParams}
	result := music.SearchMusic()
	models.WriteApiCache(m.Ctx, result)

	m.Data["json"] = result
	m.ServeJSON()
}
