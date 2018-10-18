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

func (a *ArtistController) GetArtistList() {
	artistListParams := models.ArtistListParams{CategoryCode: "1001", Initial: 65, Total: "true", Offset: 0, Limit: 30}
	if cat := a.Input().Get("cat"); cat != "" {
		artistListParams.CategoryCode = cat
	}
	if initial := a.Input().Get("initial"); initial != "" {
		// Note: string to unicode char:  `runes := []rune(str)`
		runes := []rune(strings.ToUpper(initial))
		artistListParams.Initial = runes[0]
	}
	if offset := a.Input().Get("offset"); offset != "" {
		artistListParams.Offset, _ = strconv.Atoi(offset)
	}
	if limit := a.Input().Get("limit"); limit != "" {
		artistListParams.Limit, _ = strconv.Atoi(limit)
	}
	if total := a.Input().Get("total"); total == "" {
		artistListParams.Total = "false"
	}
	artist := &models.Artist{Cookies: a.Ctx.Request.Cookies(), ArtistListParams: artistListParams}

	result := artist.GetArtistList()
	models.WriteApiCache(a.Ctx, result)

	a.Data["json"] = result
	a.ServeJSON()
}
