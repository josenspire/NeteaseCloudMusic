package models

import (
	"NeteaseCloudMusic/utils"
	"crypto/md5"
	"fmt"
	"net/http"
)

const (
	cellphoneLoginUrl = `/weapi/login/cellphone?csrf_token=`
	refreshLoginUrl   = `/weapi/login/token/refresh?csrf_token=`
)

type IUserOperation interface {
	CellphoneLogin() (interface{}, []*http.Cookie)
	RefreshLoginStatus() (interface{}, []*http.Cookie)
}

type CellphoneLoginParams struct {
	Phone         string `json:"phone"`
	Password      string `json:"password"`
	RememberLogin string `json:"rememberLogin"`
}

type User struct {
	CellphoneLoginParams
	Cookies []*http.Cookie
	// Cookies ReqCookies
}

func (user *User) CellphoneLogin() (interface{}, []*http.Cookie) {
	params := user.CellphoneLoginParams
	data := []byte(params.Password)
	has := md5.Sum(data)
	params.Password = fmt.Sprintf("%x", has)

	reqParams := utils.TransformStructToStr(params)

	response, cookies, _ := utils.NeteaseCloudRequest(cellphoneLoginUrl, reqParams, user.Cookies, http.MethodPost)
	return response, cookies
}

func (user *User) RefreshLoginStatus() (interface{}, []*http.Cookie) {
	response, cookies, _ := utils.NeteaseCloudRequest(refreshLoginUrl, "", user.Cookies, http.MethodPost)
	fmt.Println("--------", cookies)
	return response, cookies
}
