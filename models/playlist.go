package models

import (
	"NeteaseCloudMusic/utils"
	"net/http"
)

const (
	GetCatalogue       = `/weapi/playlist/catalogue`
	GetHottags         = `/weapi/playlist/hottags`
	GetNetRecommend    = `/weapi/playlist/list`
	GetHighQualityList = `/weapi/playlist/highquality/list`
	GetPlaylistDetail  = `/weapi/v3/playlist/detail`
	CreatePlaylist     = `/weapi/playlist/create`
	SubscribePlaylist  = `/weapi/playlist`
	TrackPlaylistMusic = `/weapi/playlist/manipulate/tracks`
)

type IPlayListOperator interface {
	GetCategoryList() interface{}
	GetHottagsList() interface{}
	GetNetRecommendList() interface{}
	GetHighQualityList() interface{}
	GetPlaylistDetail() interface{}
	CreatePlaylist(name string) interface{}
	SubscribePlaylist(subType string, playlistId string) interface{}
	TrackPlaylistMusic() interface{}
}

type FeatureParams struct {
	Cat   string `json:"cat"`   // category
	Order string `json:"order"` // `new` or `hot`
	PageLimitParams
}

type DetailParams struct {
	Id string `json:"id"`
	N  int    `json:"n"`
	S  int    `json:"s"`
}

type TrackParams struct {
	Op     string `json:"op"`
	Pid    string `json:"pid"`
	Tracks string `json:"trackIds"`
}

type PlayParams struct {
	Cookies []*http.Cookie
	FeatureParams
	DetailParams
	TrackParams
}

func (play *PlayParams) GetCategoryList() interface{} {
	response, _, _ := utils.NeteaseCloudRequest(GetCatalogue, nil, play.Cookies, http.MethodPost)
	return response
}
func (play *PlayParams) GetHottagsList() interface{} {
	response, _, _ := utils.NeteaseCloudRequest(GetHottags, nil, play.Cookies, http.MethodPost)
	return response
}
func (play *PlayParams) GetNetRecommendList() interface{} {
	reqParams, _ := utils.TransformStructToJSONMap(play.FeatureParams)
	response, _, _ := utils.NeteaseCloudRequest(GetNetRecommend, reqParams, play.Cookies, http.MethodPost)
	return response
}
func (play *PlayParams) GetHighQualityList() interface{} {
	reqParams, _ := utils.TransformStructToJSONMap(play.FeatureParams)
	response, _, _ := utils.NeteaseCloudRequest(GetHighQualityList, reqParams, play.Cookies, http.MethodPost)
	return response
}
func (play *PlayParams) GetPlaylistDetail() interface{} {
	reqParams, _ := utils.TransformStructToJSONMap(play.DetailParams)
	response, _, _ := utils.NeteaseCloudRequest(GetPlaylistDetail, reqParams, play.Cookies, http.MethodPost)
	return response
}
func (play *PlayParams) CreatePlaylist(name string) interface{} {
	reqParams := make(map[string]interface{})
	reqParams["name"] = name
	response, _, _ := utils.NeteaseCloudRequest(CreatePlaylist, reqParams, play.Cookies, http.MethodPost)
	return response
}
func (play *PlayParams) SubscribePlaylist(subType string, playlistId string) interface{} {
	reqParams := make(map[string]interface{})
	reqParams["id"] = playlistId
	response, _, _ := utils.NeteaseCloudRequest(SubscribePlaylist+"/"+subType, reqParams, play.Cookies, http.MethodPost)
	return response
}
func (play *PlayParams) TrackPlaylistMusic() interface{} {
	reqParams, _ := utils.TransformStructToJSONMap(play.TrackParams)
	response, _, _ := utils.NeteaseCloudRequest(TrackPlaylistMusic, reqParams, play.Cookies, http.MethodPost)
	return response
}
