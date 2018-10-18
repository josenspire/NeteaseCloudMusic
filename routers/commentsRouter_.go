package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ArtistController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ArtistController"],
		beego.ControllerComments{
			Method:           "GetArtistList",
			Router:           `/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ArtistController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ArtistController"],
		beego.ControllerComments{
			Method:           "SubscriptArtist",
			Router:           `/sub`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ArtistController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ArtistController"],
		beego.ControllerComments{
			Method:           "GetSubscriptArtistList",
			Router:           `/sublist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ArtistController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ArtistController"],
		beego.ControllerComments{
			Method:           "UnSubscriptArtist",
			Router:           `/unsub`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:IndexController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:IndexController"],
		beego.ControllerComments{
			Method:           "RenderIndex",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:LetterController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:LetterController"],
		beego.ControllerComments{
			Method:           "GetPrivateLetter",
			Router:           `/privateLetter`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:MusicController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:MusicController"],
		beego.ControllerComments{
			Method:           "GetMusicUrl",
			Router:           `/musicUrl`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:MusicController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:MusicController"],
		beego.ControllerComments{
			Method:           "SearchMusic",
			Router:           `/search`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:ObjectController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:PlaylistController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:PlaylistController"],
		beego.ControllerComments{
			Method:           "GetCategoryList",
			Router:           `/catlist`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:PlaylistController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:PlaylistController"],
		beego.ControllerComments{
			Method:           "GetPlaylistDetail",
			Router:           `/detail`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:PlaylistController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:PlaylistController"],
		beego.ControllerComments{
			Method:           "GetHighQualityList",
			Router:           `/highQuality`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:PlaylistController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:PlaylistController"],
		beego.ControllerComments{
			Method:           "GetHottagsList",
			Router:           `/hot`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:PlaylistController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:PlaylistController"],
		beego.ControllerComments{
			Method:           "GetNetRecommendList",
			Router:           `/netRecommend`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:SearchController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:SearchController"],
		beego.ControllerComments{
			Method:           "Search",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetUserDetail",
			Router:           `/detail`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetDjradioList",
			Router:           `/djradio`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetDjradioSubedList",
			Router:           `/djradioSubed`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetFolloweds",
			Router:           `/followeds`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetEvent",
			Router:           `/followeds`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetFollows",
			Router:           `/follows`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"],
		beego.ControllerComments{
			Method:           "CellphoneLogin",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetPlayList",
			Router:           `/playList`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetPlayRecord",
			Router:           `/playRecord`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"],
		beego.ControllerComments{
			Method:           "RefreshLogin",
			Router:           `/refreshLogin`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetUserSubscriptCount",
			Router:           `/subcount`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"] = append(beego.GlobalControllerRouter["NeteaseCloudMusic/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UpdateProfile",
			Router:           `/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

}
