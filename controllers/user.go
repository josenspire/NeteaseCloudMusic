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
func (u *UserController) Login() {
	resParams := models.CellphoneLoginParams{}
	u.Ctx.Input.Bind(&resParams.Phone, "phone")
	u.Ctx.Input.Bind(&resParams.Password, "password")
	u.Ctx.Input.Bind(&resParams.RememberLogin, "rememberLogin")

	defaultCookies := setupDefaultCookie()
	user := &models.User{CellphoneLoginParams: resParams, Cookies: append(u.Ctx.Request.Cookies(), defaultCookies...)}
	result, cookies := user.CellphoneLogin()
	models.WriteApiCache(u.Ctx, result)

	setupResponseCookies(u.Ctx.ResponseWriter, append(cookies, defaultCookies...))

	u.Data["json"] = result
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
	user := &models.User{Cookies: u.Ctx.Request.Cookies()}
	result, cookies := user.Logout()

	models.WriteApiCache(u.Ctx, result)
	setupResponseCookies(u.Ctx.ResponseWriter, cookies)

	u.Data["json"] = "logout success"
	u.ServeJSON()
}

// @Title GetUserDetailInformation
// @Description Get user detail information
// @Success 200 {string}
// @router /detail [get]
func (u *UserController) GetUserDetail() {
	uid := u.Ctx.Input.Param("uid")
	user := &models.User{Cookies: u.Ctx.Request.Cookies()}
	result, cookies := user.GetUserDetail(uid)

	models.WriteApiCache(u.Ctx, result)
	setupResponseCookies(u.Ctx.ResponseWriter, cookies)

	u.Data["json"] = result
	u.ServeJSON()
}

// @Title GetUserSubscriptCount
// @Description Get user subscript count information
// @Success 200 {string}
// @router /subcount [get]
func (u *UserController) GetUserSubscriptCount() {
	user := &models.User{Cookies: u.Ctx.Request.Cookies()}
	result, cookies := user.GetUserSubscriptCount()

	models.WriteApiCache(u.Ctx, result)
	setupResponseCookies(u.Ctx.ResponseWriter, cookies)

	u.Data["json"] = result
	u.ServeJSON()
}

// @Title UpdateUserInformation
// @Description Update user profile information
// @Param    nickname      query    string    true        "User nickname"
// @Param    signature     query    string    true        "User signature"
// @Param    gender    	   query    string    true        "User gender"
// @Param    birthday      query    string    true        "User birthday"
// @Param    province      query    string    true        "User province"
// @Param    city    	   query    string    true        "User city"
// @Success 200 {string}
// @router /update [post]
func (u *UserController) UpdateProfile() {
	userProfile := models.UserProfileParams{}
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &userProfile); err != nil {
		u.Ctx.Abort(http.StatusBadRequest, "Params error, please check your request params")
	}
	user := &models.User{Cookies: u.Ctx.Request.Cookies(), UserProfileParams: userProfile}
	result, cookies := user.UpdateProfile()

	models.WriteApiCache(u.Ctx, result)
	setupResponseCookies(u.Ctx.ResponseWriter, cookies)

	u.Data["json"] = result
	u.ServeJSON()
}

// @Title GetPlayList
// @Description Get User's Play List
// @Param    uid        query    string    true        "User id"
// @Param    offset     query    int    true        "Play list offset, default 0"
// @Param    limit      query    int    true        "List limit"
// @Success 200 {string}
// @router /playList [post]
func (u *UserController) GetPlayList() {
	playListParams := models.PlayListParams{}
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &playListParams); err != nil {
		u.Ctx.Abort(http.StatusBadRequest, "Params error, please check your request params")
	}
	user := &models.User{Cookies: u.Ctx.Request.Cookies(), PlayListParams: playListParams}
	result, cookies := user.GetPlayList()

	models.WriteApiCache(u.Ctx, result)
	setupResponseCookies(u.Ctx.ResponseWriter, cookies)

	u.Data["json"] = result
	u.ServeJSON()
}

// @Title GetPlayRecord
// @Description Get User's Play List
// @Param    uid      query    string    true        "User id"
// @Param    type     query    int       true        "Return list type, (-1: include 0&1, 0: allData, 1: weekData)"
// @Success 200 {string}
// @router /playRecord [post]
func (u *UserController) GetPlayRecord() {
	playListParams := models.PlayListParams{}
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &playListParams); err != nil {
		u.Ctx.Abort(http.StatusBadRequest, "Params error, please check your request params")
	}
	user := &models.User{Cookies: u.Ctx.Request.Cookies()}
	result, cookies := user.GetPlayRecord()

	models.WriteApiCache(u.Ctx, result)
	setupResponseCookies(u.Ctx.ResponseWriter, cookies)

	u.Data["json"] = result
	u.ServeJSON()
}

// @Title GetDjradioList
// @Description Get User's Play List
// @Param    uid      query    string    true        "User id"
// @Param    type     query    int       true        "Return list type, (-1: include 0&1, 0: allData, 1: weekData)"
// @Success 200 {string}
// @router /djradio [post]
func (u *UserController) GetDjradioList() {
	user := &models.User{Cookies: u.Ctx.Request.Cookies()}

	result := user.GetDjradioList(u.Ctx.Input.Query("uid"))
	models.WriteApiCache(u.Ctx, result)

	u.Data["json"] = result
	u.ServeJSON()
}

// @Title GetDjradioSubedList
// @Description Get User's Play List
// @Param    uid      query    string    true        "User id"
// @Param    type     query    int       true        "Return list type, (-1: include 0&1, 0: allData, 1: weekData)"
// @Success 200 {string}
// @router /djradioSubed [post]
func (u *UserController) GetDjradioSubedList() {
	djradioParams := models.DjradioParams{}
	u.Ctx.Input.Bind(&djradioParams.Offset, "offset")
	u.Ctx.Input.Bind(&djradioParams.Limit, "limit")
	u.Ctx.Input.Bind(&djradioParams.Total, "total")
	user := &models.User{Cookies: u.Ctx.Request.Cookies(), DjradioParams: djradioParams}

	result := user.GetDjradioSubedList()
	models.WriteApiCache(u.Ctx, result)

	u.Data["json"] = result
	u.ServeJSON()
}

// @Title GetFollows
// @Description Get User's Play List
// @Param    uid      query    string    true        "User id"
// @Param    type     query    int       true        "Return list type, (-1: include 0&1, 0: allData, 1: weekData)"
// @Success 200 {string}
// @router /follows [post]
func (u *UserController) GetFollows() {
	followParams := models.FollowParams{}
	uid := u.Ctx.Input.Query("uid")
	u.Ctx.Input.Bind(&followParams.Offset, "offset")
	u.Ctx.Input.Bind(&followParams.Limit, "limit")
	user := &models.User{Cookies: u.Ctx.Request.Cookies(), FollowParams: followParams}

	result := user.GetFollows(uid)
	models.WriteApiCache(u.Ctx, result)

	u.Data["json"] = result
	u.ServeJSON()
}

// @Title GetFolloweds
// @Description Get User's Play List
// @Param    uid      query    string    true        "User id"
// @Param    type     query    int       true        "Return list type, (-1: include 0&1, 0: allData, 1: weekData)"
// @Success 200 {string}
// @router /followeds [post]
func (u *UserController) GetFolloweds() {
	followParams := models.FollowParams{}
	userId := u.Ctx.Input.Query("uid")
	u.Ctx.Input.Bind(&followParams.Offset, "offset")
	u.Ctx.Input.Bind(&followParams.Limit, "limit")
	user := &models.User{Cookies: u.Ctx.Request.Cookies(), FollowParams: followParams}

	result := user.GetFolloweds(userId)
	models.WriteApiCache(u.Ctx, result)

	u.Data["json"] = result
	u.ServeJSON()
}

// @Title GetFolloweds
// @Description Get User's Play List
// @Param    uid      query    string    true        "User id"
// @Param    type     query    int       true        "Return list type, (-1: include 0&1, 0: allData, 1: weekData)"
// @Success 200 {string}
// @router /followeds [post]
func (u *UserController) GetEvent() {
	eventParams := models.EventParams{Time: -1, GetCounts: true}
	uid := u.Ctx.Input.Query("uid")
	user := &models.User{Cookies: u.Ctx.Request.Cookies(), EventParams: eventParams}

	result := user.GetEvent(uid)
	models.WriteApiCache(u.Ctx, result)

	u.Data["json"] = result
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
