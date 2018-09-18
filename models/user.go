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
	Username string
	Password string
}

type Profile struct {
}

const (
	cellphoneLoginUrl = "http://music.163.com/weapi/login/cellphone"
)

func Login(user User) utils.Response {
	data := []byte(user.Password)
	has := md5.Sum(data)

	user.Password = fmt.Sprintf("%x", has)

	fmt.Println(user)

	params := utils.TransformStruct2Map(user)

	response := utils.NeteaseCloudRequest(cellphoneLoginUrl, params, http.MethodPost)
	return response
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
