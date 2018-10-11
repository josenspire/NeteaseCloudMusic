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
	// uid := u.GetString(":uid")
	// models.DeleteUser(uid)
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
	resParams := models.CellphoneLoginParams{}
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &resParams)
	if err != nil {
		log.Fatal(err.Error())
		u.Data["json"] = "Params error, please check"
		u.ServeJSON()
		return
	}
	user := &models.User{CellphoneLoginParams: resParams}
	result, cookies := user.CellphoneLogin()
	WriteApiCache(u.Ctx, result)

	setupCookies(u.Ctx.ResponseWriter, cookies)

	u.Data["json"] = result
	u.ServeJSON()
}

// @Title refresh login
// @Description Refresh login status
// @Success 200 {string} refresh success
// @router /refreshLogin [post]
func (u *UserController) RefreshLogin() {
	cookies := u.Ctx.Request.Cookies()
	user := &models.User{Cookies: cookies}
	result, cookies := user.RefreshLoginStatus()

	WriteApiCache(u.Ctx, result)
	setupCookies(u.Ctx.ResponseWriter, cookies)

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

func setupCookies(rw http.ResponseWriter, cookies []*http.Cookie) {
	for _, cookie := range cookies {
		http.SetCookie(rw, cookie)
	}
}
