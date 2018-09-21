package controllers

import (
	"NeteaseCloudMusic/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type SearchController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (s *SearchController) Search() {
	var search models.Search
	json.Unmarshal(s.Ctx.Input.RequestBody, &search)

	search.Type = 1
	search.Limit = 30
	search.Offset = 0

	searchResult := models.SearchSongs(search)
	s.Data["json"] = searchResult
	s.ServeJSON()
}
