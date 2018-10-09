package controllers

import (
	"NeteaseCloudMusic/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"log"
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
// @Param	cellphone		query 	string	true		"The cellphone for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) CellphoneLogin() {
	var user models.User

	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &user); err != nil {
		log.Fatal(err.Error())
		u.Data["json"] = "Params error, please check"
		u.ServeJSON()
		return
	}

	fmt.Println("controller", user)
	result := models.Login(user)

	u.Data["json"] = result
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
