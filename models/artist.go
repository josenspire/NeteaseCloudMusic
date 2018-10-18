package models

import (
	"NeteaseCloudMusic/utils"
	"net/http"
)

const (
	GetArtistList     = `/weapi/artist/list`
	SubscriptArtist   = `/weapi/artist/sub`
	UnSubscriptArtist = `/weapi/artist/unsub`
)

type IArtistOperator interface {
	GetArtistList() interface{}
	SubscriptArtist(artistId string) interface{}
	UnSubscriptArtist(artistId string) interface{}
	GetSubscriptArtistList() interface{}
}

// @categoryCode
// @入驻歌手 5001
// @华语男歌手 1001
// @华语女歌手 1002
// @华语组合/乐队 1003
// @欧美男歌手 2001
// @欧美女歌手 2002
// @欧美组合/乐队 2003
// @日本男歌手 6001
// @日本女歌手 6002
// @日本组合/乐队 6003
// @韩国男歌手 7001
// @韩国女歌手 7002
// @韩国组合/乐队 7003
// @其他男歌手 4001
// @其他女歌手 4002
// @其他组合/乐队 4003
type ArtistListParams struct {
	CategoryCode string `json:"categoryCode"` // category code , default is `1001`
	Initial      int32  `json:"initial"`      // 按首字母索引查找参数
}

type ListParams struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Total  string `json:"total"`
}

type Artist struct {
	Cookies []*http.Cookie
	ArtistListParams
	ListParams
}

func (artist *Artist) GetArtistList() interface{} {
	artistListParams, _ := utils.TransformStructToJSONMap(artist.ArtistListParams)
	listParams, _ := utils.TransformStructToJSONMap(artist.ListParams)
	reqParams := utils.MergeMaps(artistListParams, listParams)
	response, _, _ := utils.NeteaseCloudRequest(GetArtistList, reqParams, artist.Cookies, http.MethodPost)
	return response
}
func (artist *Artist) SubscriptArtist(artistId string) interface{} {
	reqParams := make(map[string]interface{})
	reqParams["artistId"] = artistId
	response, _, _ := utils.NeteaseCloudRequest(SubscriptArtist, reqParams, artist.Cookies, http.MethodPost)
	return response
}
func (artist *Artist) UnSubscriptArtist(artistId string) interface{} {
	reqParams := make(map[string]interface{})
	reqParams["artistId"] = artistId
	reqParams["artistIds"] = "[" + artistId + "]"
	response, _, _ := utils.NeteaseCloudRequest(UnSubscriptArtist, reqParams, artist.Cookies, http.MethodPost)
	return response
}
func (artist *Artist) GetSubscriptArtistList() interface{} {
	reqParams, _ := utils.TransformStructToJSONMap(artist.ListParams)
	response, _, _ := utils.NeteaseCloudRequest(GetArtistList, reqParams, artist.Cookies, http.MethodPost)
	return response
}
