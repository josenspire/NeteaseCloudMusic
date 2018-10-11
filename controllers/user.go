package controllers

import (
	"NeteaseCloudMusic/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"net/http"
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
// @Param	phone		query 	string	true		"The cellphone for login"
// @Param	password		query 	string	true		"The password for login"
// @Param	rememberLogin	query 	string	false		"Remember login status"
// @Success 200 {json} login success
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
	result, cookies := models.Login(user)
	WriteApiCache(u.Ctx, result)

	for _, cookie := range cookies {
		http.SetCookie(u.Ctx.ResponseWriter, cookie)
	}
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
