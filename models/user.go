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
	cellphoneLoginUrl = "/weapi/login/cellphone?csrf_token="
)

func Login(user User) string {
	data := []byte(user.Password)
	has := md5.Sum(data)

	user.Password = fmt.Sprintf("%x", has)

	fmt.Println(user)

	params := utils.TransformStruct2Map(user)

	response, _ := utils.NeteaseCloudRequest(cellphoneLoginUrl, params, http.MethodPost)
	return response
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
