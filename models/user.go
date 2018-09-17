package models

import (
	"NeteaseCloudMusic/utils"
	"net/http"
	"reflect"
)

var (
	UserList map[string]*User
)

type User struct {
	Username string
	Password string
}

type Profile struct {
}

const (
	cellphoneLoginUrl = "http://music.163.com/weapi/login/cellphone"
)

func Login(user User) utils.Response {

	t := reflect.TypeOf(user)
	v := reflect.ValueOf(user)

	var params = make(map[string]string)

	for i := 0; i < t.NumField(); i++ {
		params[t.Field(i).Name] = v.Field(i).String()
	}

	response := utils.NeteaseCloudRequest(cellphoneLoginUrl, params, http.MethodPost)
	return response
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
