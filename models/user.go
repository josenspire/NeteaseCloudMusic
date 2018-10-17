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
	PlayRecord                = `/weapi/v1/play/record`
	DjRadio                   = `/weapi/djradio/get/byuser`
	DjRadioSubed              = `/weapi/djradio/get/subed`
	GetFollows                = `/weapi/user/getfollows`
	GetFolloweds              = `/weapi/user/getfolloweds`
	GetEvent                  = `/weapi/event/get`
)

type IUserOperation interface {
	CellphoneLogin() (interface{}, []*http.Cookie)
	RefreshLoginStatus() (interface{}, []*http.Cookie)
	Logout() (interface{}, []*http.Cookie)
	QueryUserStatus() (interface{}, []*http.Cookie)
	GetUserDetail(uid string) (interface{}, []*http.Cookie)
	UpdateProfile() (interface{}, []*http.Cookie)
	GetPlayList() (interface{}, []*http.Cookie)
	GetPlayRecord() (interface{}, []*http.Cookie)
	GetDjradioList() interface{}
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
	Type   int    `json:"type"`   // data type, 0：allData, 1: weekData, -1: allData & weekData
}

type DjradioParams struct {
	Offset int    `json:"offset"` // list offset
	Limit  int    `json:"limit"`  // count limit
	Total  string `json:"type"`
}

type FollowParams struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Order  string `json:"order"`
}

type EventParams struct {
	Time      int  `json:"time"`
	GetCounts bool `json:"getcounts"`
}

type User struct {
	CellphoneLoginParams
	Cookies []*http.Cookie
	UserProfileParams
	PlayListParams
	DjradioParams
	FollowParams
	EventParams
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
	response, cookies, _ := utils.NeteaseCloudRequest(Logout, nil, user.Cookies, http.MethodPost)
	return response, cookies
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
func (user *User) GetPlayRecord() (interface{}, []*http.Cookie) {
	reqParams, _ := utils.TransformStructToJSONMap(user.PlayListParams)
	response, cookies, _ := utils.NeteaseCloudRequest(PlayRecord, reqParams, user.Cookies, http.MethodPost)
	return response, cookies
}
func (user *User) GetDjradioList(uid string) interface{} {
	reqParams := make(map[string]interface{})
	reqParams["userId"] = uid
	response, _, _ := utils.NeteaseCloudRequest(DjRadio, reqParams, user.Cookies, http.MethodPost)
	return response
}
func (user *User) GetDjradioSubedList() interface{} {
	reqParams, _ := utils.TransformStructToJSONMap(user.DjradioParams)
	response, _, _ := utils.NeteaseCloudRequest(DjRadioSubed, reqParams, user.Cookies, http.MethodPost)
	return response
}
func (user *User) GetFollows(uid string) interface{} {
	reqParams, _ := utils.TransformStructToJSONMap(user.FollowParams)
	response, _, _ := utils.NeteaseCloudRequest(GetFollows+"/"+uid, reqParams, user.Cookies, http.MethodPost)
	return response
}
func (user *User) GetFolloweds(userId string) interface{} {
	reqParams, _ := utils.TransformStructToJSONMap(user.FollowParams)
	reqParams["userId"] = userId
	response, _, _ := utils.NeteaseCloudRequest(GetFolloweds, reqParams, user.Cookies, http.MethodPost)
	return response
}
func (user *User) GetEvent(uid string) interface{} {
	reqParams, _ := utils.TransformStructToJSONMap(user.EventParams)
	response, _, _ := utils.NeteaseCloudRequest(GetEvent+"/"+uid, reqParams, user.Cookies, http.MethodPost)
	return response
}
