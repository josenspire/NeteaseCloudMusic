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
)

type IPlayListOperator interface {
	GetCategoryList() interface{}
	GetHottagsList() interface{}
	GetNetRecommendList() interface{}
	GetHighQualityList() interface{}
	GetPlaylistDetail() interface{}
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

type PlayParams struct {
	Cookies []*http.Cookie
	FeatureParams
	DetailParams
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
