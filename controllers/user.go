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
	defaultCookies := setupDefaultCookie()
	user := &models.User{CellphoneLoginParams: resParams, Cookies: append(u.Ctx.Request.Cookies(), defaultCookies...)}
	result, cookies := user.CellphoneLogin()
	models.WriteApiCache(u.Ctx, result)

	setupResponseCookies(u.Ctx.ResponseWriter, append(cookies, defaultCookies...))

	u.Data["json"] = result
	u.ServeJSON()
}

// @Title refresh login
// @Description Refresh login status
// @Success 200 {string} refresh success
// @router /refreshLogin [post]
func (u *UserController) RefreshLogin() {
	user := &models.User{Cookies: u.Ctx.Request.Cookies()}
	result, cookies := user.RefreshLoginStatus()

	models.WriteApiCache(u.Ctx, result)
	setupResponseCookies(u.Ctx.ResponseWriter, cookies)

	u.Data["json"] = result
	u.ServeJSON()
}

func (u *UserController) QueryUserStatus() {
	user := &models.User{}
	result, _ := user.QueryUserStatus()
	u.Data["json"] = result
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	// user := &models.User{Cookies: u.Ctx.Request.Cookies()}
	// result, cookies := user.Logout()
	//
	// WriteApiCache(u.Ctx, result)
	// setupResponseCookies(u.Ctx.ResponseWriter, cookies)

	u.Data["json"] = "logout success"
	u.ServeJSON()
}

func setupResponseCookies(rw http.ResponseWriter, cookies []*http.Cookie) {
	for _, cookie := range cookies {
		http.SetCookie(rw, cookie)
	}
}

func setupDefaultCookie() []*http.Cookie {
	cookies := make([]*http.Cookie, 4)
	cookies[0] = &http.Cookie{Name: "appver", Value: "1.5.9"}
	cookies[1] = &http.Cookie{Name: "os", Value: "osx"}
	cookies[2] = &http.Cookie{Name: "channel", Value: "netease"}
	cookies[3] = &http.Cookie{Name: "osver", Value: "%e7%89%88%e6%9c%ac+10.13.2%ef%bc%88%e7%89%88%e5%8f%b7+17C88%ef%bc%89"}
	return cookies
}
