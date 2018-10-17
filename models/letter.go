package models

import (
	"NeteaseCloudMusic/utils"
	"net/http"
)

const (
	GetPrivateLetter  = `/weapi/msg/private/users`
	SendPrivateLetter = `/weapi/msg/private/send`
)

type ILetterOperator interface {
	GetPrivateLetter() interface{}
	SendPrivateText() interface{}
	SendWithPlaylist() interface{}
}

type SendLetterParams struct {
	UserIds string `json:"userIds"` // one or more, split by `,`
	Msg     string `json:"msg"`     // message
	Type    string `json:"type"`    // default `text`
	Id      string `json:"id"`      // playlist id
}

type LetterParams struct {
	Uid string `json:"uid"`
}

type Letter struct {
	Cookies []*http.Cookie
	LetterParams
	SendLetterParams
}

func (letter *Letter) GetPrivateLetter() interface{} {
	reqParams, _ := utils.TransformStructToJSONMap(letter.LetterParams)
	response, _, _ := utils.NeteaseCloudRequest(GetPrivateLetter, reqParams, letter.Cookies, http.MethodPost)
	return response
}
func (letter *Letter) SendPrivateText() interface{} {
	reqParams, _ := utils.TransformStructToJSONMap(letter.SendLetterParams)
	response, _, _ := utils.NeteaseCloudRequest(SendPrivateLetter, reqParams, letter.Cookies, http.MethodPost)
	return response
}
func (letter *Letter) SendWithPlaylist() interface{} {
	reqParams, _ := utils.TransformStructToJSONMap(letter.SendLetterParams)
	response, _, _ := utils.NeteaseCloudRequest(SendPrivateLetter, reqParams, letter.Cookies, http.MethodPost)
	return response
}
