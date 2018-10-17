package controllers

import (
	"NeteaseCloudMusic/models"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type ArtistController struct {
	beego.Controller
}

func (a *ArtistController) GetArtistList() {
	artistListParams := models.ArtistListParams{CategoryCode: "1001", Initial: 69, Total: "true", Offset: 0, Limit: 30}
	if cat := a.Input().Get("cat"); cat != "" {
		artistListParams.CategoryCode = cat
	}
	if initial := a.Input().Get("initial"); initial != "" {
		runes := []rune("a")
		// TODO
		fmt.Println("===========", runes, int(runes[0]), int32(runes[0]))
		artistListParams.Initial = int(runes[0])
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
