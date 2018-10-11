package models

import (
	"NeteaseCloudMusic/utils"
	"crypto/md5"
	"fmt"
	"net/http"
)

type Login interface {
	CellphoneLogin(params Params)
}

type Params struct {
	Phone         string `json:"phone"`
	Password      string `json:"password"`
	RememberLogin string `json:"rememberLogin"`
}

type User struct {
	Params
}

const (
	cellphoneLoginUrl = "/weapi/login/cellphone?csrf_token="
)

func (user *User) CellphoneLogin(params Params) (interface{}, []*http.Cookie) {
	data := []byte(params.Password)
	has := md5.Sum(data)
	params.Password = fmt.Sprintf("%x", has)

	reqParams := utils.TransformStructToStr(params)

	response, cookies, _ := utils.NeteaseCloudRequest(cellphoneLoginUrl, reqParams, http.MethodPost)
	return response, cookies
}
