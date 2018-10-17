package models

import (
	"NeteaseCloudMusic/utils"
	"net/http"
)

const (
	GetArtistList = `/weapi/artist/list`
)

type IArtistOperator interface {
	GetArtistList() interface{}
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
	Offset       int    `json:"offset"`
	Total        string `json:"total"`
	Limit        int    `json:"limit"`
	Initial      int    `json:"initial"` // 按首字母索引查找参数
}

type Artist struct {
	Cookies []*http.Cookie
	ArtistListParams
}

func (artist *Artist) GetArtistList() interface{} {
	reqParams, _ := utils.TransformStructToJSONMap(artist.ArtistListParams)
	response, _, _ := utils.NeteaseCloudRequest(GetArtistList, reqParams, artist.Cookies, http.MethodPost)
	return response
}
