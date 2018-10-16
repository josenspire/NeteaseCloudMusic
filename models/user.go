package models

import (
	"NeteaseCloudMusic/utils"
	"crypto/md5"
	"fmt"
	"net/http"
)

const (
	CellphoneLoginUrl         = `/weapi/login/cellphone?csrf_token=`
	RefreshLoginUrl           = `/weapi/login/token/refresh`
	Logout                    = `/weapi/login/token/logout`
	GetUserDetail             = `/weapi/v1/user/detail`
	GetUserAccountInformation = `/weapi/subcount`
	UpdateUserInformation     = `/weapi/user/profile/update`
	PlayList                  = `/weapi/user/playlist`
)

type IUserOperation interface {
	CellphoneLogin() (interface{}, []*http.Cookie)
	RefreshLoginStatus() (interface{}, []*http.Cookie)
	Logout() (interface{}, []*http.Cookie)
	QueryUserStatus() (interface{}, []*http.Cookie)
	GetUserDetail(uid string) (interface{}, []*http.Cookie)
	UpdateProfile() (interface{}, []*http.Cookie)
	GetPlayList() (interface{}, []*http.Cookie)
}

type CellphoneLoginParams struct {
	Phone         string `json:"phone"`
	Password      string `json:"password"`
	RememberLogin string `json:"rememberLogin"`
}
type UserProfileParams struct {
	Nickname  string `json:"nickname"`  // user nickname
	Signature string `json:"signature"` // user signature
	Gender    string `json:gender`      // gender- 0: privary， 1: boy, 2: girl
	Birthday  string `json:birthday`    // type as Unix timestamp
	Province  string `json.province`    // province id
	City      string `json.city`        // city id
}
type PlayListParams struct {
	Uid    string `json:"uid"`    // user id
	Offset int    `json:"offset"` // list offset
	Limit  int    `json:"limit"`  // count limit
}

type User struct {
	CellphoneLoginParams
	Cookies []*http.Cookie
	UserProfileParams
	PlayListParams
}

func (user *User) CellphoneLogin() (interface{}, []*http.Cookie) {
	params := user.CellphoneLoginParams
	data := []byte(params.Password)
	has := md5.Sum(data)
	params.Password = fmt.Sprintf("%x", has)

	reqParams, _ := utils.TransformStructToJSONMap(params)

	response, cookies, _ := utils.NeteaseCloudRequest(CellphoneLoginUrl, reqParams, user.Cookies, http.MethodPost)
	return response, cookies
}

func (user *User) RefreshLoginStatus() (interface{}, []*http.Cookie) {
	response, cookies, _ := utils.NeteaseCloudRequest(RefreshLoginUrl, nil, user.Cookies, http.MethodPost)
	return response, cookies
}
func (user *User) Logout() (interface{}, []*http.Cookie) {
	// response, cookies, _ := utils.NeteaseCloudRequest(RefreshLoginUrl, nil, user.Cookies, http.MethodPost)
	// return response, cookies
	return nil, nil
}
func (user *User) QueryUserStatus() (interface{}, []*http.Cookie) {
	// responseBody := ReadCacheDataByKey(CellphoneLoginUrl)
	// if responseBody == nil {
	// 	return "Get Status Fail, Please login first", nil
	// }
	return nil, nil
}
func (user *User) GetUserDetail(uid string) (interface{}, []*http.Cookie) {
	response, cookies, _ := utils.NeteaseCloudRequest(GetUserDetail+uid, nil, user.Cookies, http.MethodPost)
	return response, cookies
}
func (user *User) GetUserSubscriptCount() (interface{}, []*http.Cookie) {
	response, cookies, _ := utils.NeteaseCloudRequest(GetUserAccountInformation, nil, user.Cookies, http.MethodPost)
	return response, cookies
}
func (user *User) UpdateProfile() (interface{}, []*http.Cookie) {
	reqParams, _ := utils.TransformStructToJSONMap(user.UserProfileParams)
	response, cookies, _ := utils.NeteaseCloudRequest(UpdateUserInformation, reqParams, user.Cookies, http.MethodPost)
	return response, cookies
}
func (user *User) GetPlayList() (interface{}, []*http.Cookie) {
	reqParams, _ := utils.TransformStructToJSONMap(user.PlayListParams)
	response, cookies, _ := utils.NeteaseCloudRequest(PlayList, reqParams, user.Cookies, http.MethodPost)
	return response, cookies
}
