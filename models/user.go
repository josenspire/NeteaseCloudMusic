package models

import (
	"NeteaseCloudMusic/utils"
	"crypto/md5"
	"fmt"
	"net/http"
)

var (
	UserList map[string]*User
)

type User struct {
	Phone         string `json:"phone"`
	Password      string `json:"password"`
	RememberLogin string `json:"rememberLogin"`
}

type Profile struct {
}

const (
	cellphoneLoginUrl = "/weapi/login/cellphone?csrf_token="
)

func Login(user User) (interface{}, []*http.Cookie) {
	data := []byte(user.Password)
	has := md5.Sum(data)
	user.Password = fmt.Sprintf("%x", has)

	params := utils.TransformStructToStr(user)

	response, cookies, _ := utils.NeteaseCloudRequest(cellphoneLoginUrl, params, http.MethodPost)
	return response, cookies
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
