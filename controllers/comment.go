package controllers

import (
	"NeteaseCloudMusic/models"
	"github.com/astaxie/beego"
	"strconv"
)

type CommentController struct {
	beego.Controller
}

var typeMap = map[string]string{
	"0": "R_SO_4_", // 歌曲
	"1": "R_MV_5_", // mv
	"2": "A_PL_0_", // 歌单
	"3": "R_AL_3_", // 专辑
	"4": "A_DJ_1_", // 电台
}

// @Title GetMusicComment
// @Description  Get music comment
// @Params  id   	 query    string    true    "music id"
// @Params  offset   	 query    string    false    "result data offset, default 0"
// @Params  limit   	 query    string    false    "result data items limit, default 20"
// @Success 200 {string}
// @router /music [get]
func (c *CommentController) GetMusicComment() {
	if rid := c.Input().Get("id"); rid == "" {
		c.Data["json"] = "Params error, please check your request"
	} else {
		pageLimitParams := models.PageLimitParams{Offset: 0, Limit: 20}
		if limit := c.Input().Get("limit"); limit != "" {
			pageLimitParams.Limit, _ = strconv.Atoi(limit)
		}
		if offset := c.Input().Get("offset"); offset != "" {
			pageLimitParams.Offset, _ = strconv.Atoi(offset)
		}
		music := models.Comment{Cookies: c.Ctx.Request.Cookies(), PageLimitParams: pageLimitParams}
		result := music.GetMusicComment(rid)
		models.WriteApiCache(c.Ctx, result)
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// @Title GetAlbumComment
// @Description  Get album comment
// @Params  id   	 query    string    true    "music id"
// @Params  offset   	 query    string    false    "result data offset, default 0"
// @Params  limit   	 query    string    false    "result data items limit, default 20"
// @Success 200 {string}
// @router /album [get]
func (c *CommentController) GetAlbumComment() {
	if rid := c.Input().Get("id"); rid == "" {
		c.Data["json"] = "Params error, please check your request"
	} else {
		pageLimitParams := models.PageLimitParams{Offset: 0, Limit: 20}
		if limit := c.Input().Get("limit"); limit != "" {
			pageLimitParams.Limit, _ = strconv.Atoi(limit)
		}
		if offset := c.Input().Get("offset"); offset != "" {
			pageLimitParams.Offset, _ = strconv.Atoi(offset)
		}
		music := models.Comment{Cookies: c.Ctx.Request.Cookies(), PageLimitParams: pageLimitParams}
		result := music.GetAlbumComment(rid)
		models.WriteApiCache(c.Ctx, result)
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// @Title GetPlaylistComment
// @Description  Get playlist comment
// @Params  id   	 query    string    true    "music id"
// @Params  offset   	 query    string    false    "result data offset, default 0"
// @Params  limit   	 query    string    false    "result data items limit, default 20"
// @Success 200 {string}
// @router /playlist [get]
func (c *CommentController) GetPlaylistComment() {
	if rid := c.Input().Get("id"); rid == "" {
		c.Data["json"] = "Params error, please check your request"
	} else {
		pageLimitParams := models.PageLimitParams{Offset: 0, Limit: 20}
		if limit := c.Input().Get("limit"); limit != "" {
			pageLimitParams.Limit, _ = strconv.Atoi(limit)
		}
		if offset := c.Input().Get("offset"); offset != "" {
			pageLimitParams.Offset, _ = strconv.Atoi(offset)
		}
		music := models.Comment{Cookies: c.Ctx.Request.Cookies(), PageLimitParams: pageLimitParams}
		result := music.GetPlaylistComment(rid)
		models.WriteApiCache(c.Ctx, result)
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// @Title GetMVComment
// @Description  Get mv comment
// @Params  id   	 query    string    true    "music id"
// @Params  offset   	 query    string    false    "result data offset, default 0"
// @Params  limit   	 query    string    false    "result data items limit, default 20"
// @Success 200 {string}
// @router /playlist [get]
func (c *CommentController) GetMVComment() {
	if rid := c.Input().Get("id"); rid == "" {
		c.Data["json"] = "Params error, please check your request"
	} else {
		pageLimitParams := models.PageLimitParams{Offset: 0, Limit: 20}
		if limit := c.Input().Get("limit"); limit != "" {
			pageLimitParams.Limit, _ = strconv.Atoi(limit)
		}
		if offset := c.Input().Get("offset"); offset != "" {
			pageLimitParams.Offset, _ = strconv.Atoi(offset)
		}
		music := models.Comment{Cookies: c.Ctx.Request.Cookies(), PageLimitParams: pageLimitParams}
		result := music.GetMVComment(rid)
		models.WriteApiCache(c.Ctx, result)
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// @Title GetDJComment
// @Description  Get dj comment
// @Params  id   	 query    string    true    "music id"
// @Params  offset   	 query    string    false    "result data offset, default 0"
// @Params  limit   	 query    string    false    "result data items limit, default 20"
// @Success 200 {string}
// @router /dj [get]
func (c *CommentController) GetDJComment() {
	if rid := c.Input().Get("id"); rid == "" {
		c.Data["json"] = "Params error, please check your request"
	} else {
		pageLimitParams := models.PageLimitParams{Offset: 0, Limit: 20}
		if limit := c.Input().Get("limit"); limit != "" {
			pageLimitParams.Limit, _ = strconv.Atoi(limit)
		}
		if offset := c.Input().Get("offset"); offset != "" {
			pageLimitParams.Offset, _ = strconv.Atoi(offset)
		}
		music := models.Comment{Cookies: c.Ctx.Request.Cookies(), PageLimitParams: pageLimitParams}
		result := music.GetDJComment(rid)
		models.WriteApiCache(c.Ctx, result)
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// @Title LikeComment
// @Description  Like one of comment
// @Params  id   	 query    string    true    "music id"
// @Params  cid   	 query    string    true   	"comment id"
// @Params  type   	 query    string    true   	"target type, `music` `mv` `dj` `playlist` `album`"
// @Params  action   query    string    true   	"action: 1: like, 2: unlike"
// @Success 200 {string}
// @router /lick [get]
func (c *CommentController) LikeComment() {
	cid := c.Input().Get("cid")
	id := c.Input().Get("id")
	cType := c.Input().Get("type")
	action := c.Input().Get("action")

	if cid == "" || id == "" || cType == "" {
		c.Data["json"] = "Params error, please check your request"
	} else {
		var _action = ""
		if action == "" || action == "1" {
			_action = "like"
		} else {
			_action = "unlike"
		}
		basicParams := models.BasicParams{ThreadId: typeMap[cType] + id, CommentId: cid}
		music := models.Comment{Cookies: c.Ctx.Request.Cookies(), BasicParams: basicParams}
		result := music.LikeComment(_action)
		models.WriteApiCache(c.Ctx, result)
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// @Title Publish
// @Description  Publish comment
// @Params  id   	 query    string    true     "music/mv/... id"
// @Params  action   query    string    false    "operation: publish or add"
// @Params  type   	 query    string    false    "target type: music, mv, playlist..."
// @Params  content  query	  string 	true	 "comment content"
// @Success 200 {string}
// @router /publish [get]
func (c *CommentController) Publish() {
	id := c.Input().Get("id")
	cType := c.Input().Get("type")
	action := c.Input().Get("action")
	content := c.Input().Get("content")
	commentId := c.Input().Get("commentId")
	if id == "" || id == "" || cType == "" || (action == "1" && content == "") || (action == "2" && commentId == "") {
		c.Data["json"] = "Params error, please check your request"
	} else {
		var (
			_action = ""
		)
		if action == "" || action == "1" {
			_action = "add"
			basicParams := models.BasicParams{ThreadId: typeMap[cType] + id}
			music := models.Comment{Cookies: c.Ctx.Request.Cookies(), BasicParams: basicParams}
			result := music.Publish(content, _action)
			models.WriteApiCache(c.Ctx, result)
			c.Data["json"] = result
		} else {
			_action = "delete"
			basicParams := models.BasicParams{ThreadId: typeMap[cType] + id, CommentId: commentId}
			music := models.Comment{Cookies: c.Ctx.Request.Cookies(), BasicParams: basicParams}
			result := music.Publish("", _action)
			models.WriteApiCache(c.Ctx, result)
			c.Data["json"] = result
		}

	}
	c.ServeJSON()
}
