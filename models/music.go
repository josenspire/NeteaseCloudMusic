package models

import (
	"NeteaseCloudMusic/utils"
	"net/http"
)

const (
	GetMusicUrl = `/weapi/song/enhance/player/url`
	SearchMusic = `/weapi/search/get`
)

type IMusicOperator interface {
	GetMusicUrl() interface{}
	SearchMusic() interface{}
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
