package controllers

import (
	"NeteaseCloudMusic/models"
	"github.com/astaxie/beego"
	"strconv"
)

type PlaylistController struct {
	beego.Controller
}

// @Title GetCategoryList
// @Description Get playlist catalogue
// @Success 200 {string}
// @router /catlist [get]
func (p *PlaylistController) GetCategoryList() {
	playList := models.PlayParams{Cookies: p.Ctx.Request.Cookies()}
	result := playList.GetCategoryList()
	models.WriteApiCache(p.Ctx, result)

	p.Data["json"] = result
	p.ServeJSON()
}

// @Title GetHottagsList
// @Description Get hot playlist catalogue
// @Success 200 {string}
// @router /hot [get]
func (p *PlaylistController) GetHottagsList() {
	playList := models.PlayParams{Cookies: p.Ctx.Request.Cookies()}
	result := playList.GetHottagsList()
	models.WriteApiCache(p.Ctx, result)

	p.Data["json"] = result
	p.ServeJSON()
}

// @Title GetNetRecommendList
// @Description Get net friend recommend playlist catalogue (net friend recommend)
// @Params cat   	query    string    false        "cat: `全部`, e.g. "华语 "、"古风 " 、"欧美 "、"流行 ", 默认为 `全部`"
// @Params order  	query    string    false        "can select `new` or `hot`, default is `hot`"
// @Params offset  	query    string    false        "result data offset"
// @Params limit  	query    int	   false        "result item limit"
// @Params total  	query    int	   false        "unknow"
// @Success 200 {string}
// @router /netRecommend [get]
func (p *PlaylistController) GetNetRecommendList() {
	pageLimitParams := models.PageLimitParams{Offset: 0, Limit: 50, Total: `true`}
	featureParams := models.FeatureParams{Cat: `全部`, Order: `hot`, PageLimitParams: pageLimitParams}

	playList := models.PlayParams{Cookies: p.Ctx.Request.Cookies(), FeatureParams: featureParams}
	result := playList.GetNetRecommendList()
	models.WriteApiCache(p.Ctx, result)

	p.Data["json"] = result
	p.ServeJSON()
}

// @Title GetHighQualityList
// @Description Get high quality playlist catalogue (featured)
// @Params cat   	query    string    false        "cat: `全部`, e.g. "华语 "、"古风 " 、"欧美 "、"流行 ", 默认为 `全部`"
// @Params offset  	query    string    false        "result data offset, default is `0`"
// @Params limit  	query    int	   false        "result item limit, default is `20`"
// @Success 200 {string}
// @router /highQuality [get]
func (p *PlaylistController) GetHighQualityList() {
	pageLimitParams := models.PageLimitParams{Offset: 0, Limit: 20}
	featureParams := models.FeatureParams{Cat: `全部`, PageLimitParams: pageLimitParams}

	playList := models.PlayParams{Cookies: p.Ctx.Request.Cookies(), FeatureParams: featureParams}
	result := playList.GetHighQualityList()
	models.WriteApiCache(p.Ctx, result)

	p.Data["json"] = result
	p.ServeJSON()
}

// @Title GetPlaylistDetail
// @Description Get playlist detail information
// @Params id   query    string    true        "play list id"
// @Params s  	query    string    false       "歌单最近 s 个收藏着"
// @Params n  	query    int	   false       "unknow"
// @Success 200 {string}
// @router /detail [get]
func (p *PlaylistController) GetPlaylistDetail() {
	detailParams := models.DetailParams{N: 100000, S: 8}

	if id := p.Input().Get("id"); id == "" {
		p.Data["json"] = "Params error, please check your request"
		p.ServeJSON()
	} else {
		detailParams.Id = id
	}
	if n := p.Input().Get("n"); n != "" {
		detailParams.N, _ = strconv.Atoi(n)
	}
	if s := p.Input().Get("s"); s != "" {
		detailParams.S, _ = strconv.Atoi(s)
	}
	playList := models.PlayParams{Cookies: p.Ctx.Request.Cookies(), DetailParams: detailParams}
	result := playList.GetPlaylistDetail()
	models.WriteApiCache(p.Ctx, result)

	p.Data["json"] = result
	p.ServeJSON()
}
