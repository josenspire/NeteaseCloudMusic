package models

import (
	"NeteaseCloudMusic/utils"
	"net/http"
)

const (
	GetMusicUrl      = `/weapi/song/enhance/player/url`
	SearchMusic      = `/weapi/search/get`
	GetHotSearchList = `/weapi/search/hot`
	SearchSuggest    = `/weapi/search/suggest/web`
	GetMusicLyric    = `/weapi/song/lyric?os=osx&id=`
)

type IMusicOperator interface {
	GetMusicUrl() interface{}
	SearchMusic() interface{}
	GetHotSearchList() interface{}
	GetSearchSuggest(keywords string) interface{}
	GetLyric(id string) interface{}
	GetMusicDetail(s string) interface{}
}

type MusicUrlParams struct {
	Ids     string `json:"ids"` // music ids
	Bitrate string `json:"br"`  // bitrate, default `999000`
}

type SearchParams struct {
	Keywords string `json:"s"`
	Type     int    `json:"type"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
}

type Music struct {
	Cookies []*http.Cookie
	MusicUrlParams
	SearchParams
}

func (music *Music) GetMusicUrl() interface{} {
	reqParams, _ := utils.TransformStructToJSONMap(music.MusicUrlParams)
	response, _, _ := utils.NeteaseCloudRequest(GetMusicUrl, reqParams, music.Cookies, http.MethodPost)
	return response
}
func (music *Music) SearchMusic() interface{} {
	reqParams, _ := utils.TransformStructToJSONMap(music.SearchParams)
	response, _, _ := utils.NeteaseCloudRequest(SearchMusic, reqParams, music.Cookies, http.MethodPost)
	return response
}
func (music *Music) GetHotSearchList() interface{} {
	reqParams := make(map[string]interface{})
	reqParams["type"] = "1111"
	response, _, _ := utils.NeteaseCloudRequest(GetHotSearchList, reqParams, music.Cookies, http.MethodPost)
	return response
}
func (music *Music) GetSearchSuggest(keywords string) interface{} {
	reqParams := make(map[string]interface{})
	reqParams["s"] = keywords
	response, _, _ := utils.NeteaseCloudRequest(SearchSuggest, reqParams, music.Cookies, http.MethodPost)
	return response
}
func (music *Music) GetLyric(id string) interface{} {
	postfix := `&lv=-1&kv=-1&tv=-1`
	response, _, _ := utils.NeteaseCloudRequest(GetMusicLyric+id+postfix, nil, music.Cookies, http.MethodPost)
	return response
}
func (music *Music) GetMusicDetail(s string) interface{} {
	// TODO
	postfix := `&lv=-1&kv=-1&tv=-1`
	response, _, _ := utils.NeteaseCloudRequest(GetMusicLyric+id+postfix, nil, music.Cookies, http.MethodPost)
	return response
}
