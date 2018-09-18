package controllers

import (
	"NeteaseCloudMusic/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) CellphoneLogin() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)

	fmt.Println(user)

	resultByte := models.Login(user)

	// var response utils.Response
	// json.Unmarshal(resultByte, &response)
	//
	// decodeStr := response.Data[:]
	fmt.Println("-----------------", resultByte)

	// result, _ := base64.StdEncoding.DecodeString(string(decodeStr))

	u.Data["json"] = resultByte
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
