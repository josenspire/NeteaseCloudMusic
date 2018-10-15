package models

import (
	"NeteaseCloudMusic/utils"
	"crypto/md5"
	"fmt"
	"net/http"
)

const (
	cellphoneLoginUrl = `/weapi/login/cellphone?csrf_token=`
	refreshLoginUrl   = `/weapi/login/token/refresh`
	logout            = `/weapi/login/token/logout`
)

type IUserOperation interface {
	CellphoneLogin() (interface{}, []*http.Cookie)
	RefreshLoginStatus() (interface{}, []*http.Cookie)
	Logout() (interface{}, []*http.Cookie)
	QueryUserStatus() (interface{}, []*http.Cookie)
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

	reqParams, _ := utils.TransformStructToJSONMap(params)

	response, cookies, _ := utils.NeteaseCloudRequest(cellphoneLoginUrl, reqParams, user.Cookies, http.MethodPost)
	return response, cookies
}

func (user *User) RefreshLoginStatus() (interface{}, []*http.Cookie) {
	response, cookies, _ := utils.NeteaseCloudRequest(refreshLoginUrl, nil, user.Cookies, http.MethodPost)
	return response, cookies
}
func (user *User) Logout() (interface{}, []*http.Cookie) {
	// response, cookies, _ := utils.NeteaseCloudRequest(refreshLoginUrl, nil, user.Cookies, http.MethodPost)
	// return response, cookies
	return nil, nil
}
func (user *User) QueryUserStatus() (interface{}, []*http.Cookie) {
	responseBody := ReadCacheDataByKey(cellphoneLoginUrl)
	if responseBody == nil {
		return "Get Status Fail, Please login first", nil
	}
	return responseBody, nil
}
