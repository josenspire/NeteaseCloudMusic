// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"NeteaseCloudMusic/controllers"
	"NeteaseCloudMusic/models"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "get:RenderIndex")

	idx := beego.NewNamespace("/test",
		// api cache checking
		beego.NSBefore(models.ReadApiCache),
		beego.NSRouter("/cellphone", &controllers.IndexController{}, "get:CellphoneLogin"),
		beego.NSRouter("/update", &controllers.IndexController{}, "get:UpdateProfile"),
		beego.NSRouter("/playList", &controllers.IndexController{}, "get:GetPlayList"),
		beego.NSRouter("/playRecord", &controllers.IndexController{}, "get:GetPlayRecord"),
		beego.NSRouter("/privateLetter", &controllers.IndexController{}, "get:GetPrivateLetter"),
	)

	ns := beego.NewNamespace("/v1/api",
		// api cache checking
		beego.NSBefore(models.ReadApiCache),

		beego.NSNamespace("/user",
			beego.NSRouter("/login", &controllers.UserController{}, "get:Login"),
			beego.NSRouter("/cellphone", &controllers.UserController{}, "post:CellphoneLogin"),
			beego.NSRouter("/refreshLogin", &controllers.UserController{}, "get:RefreshLogin"),
			beego.NSRouter("/logout", &controllers.UserController{}, "get:Logout"),
			beego.NSRouter("/subcount", &controllers.UserController{}, "get:GetUserSubscriptCount"),
			beego.NSRouter("/update", &controllers.UserController{}, "post:UpdateProfile"),

			// play & djradio
			beego.NSRouter("/playList", &controllers.UserController{}, "post:GetPlayList"),
			beego.NSRouter("/playRecord", &controllers.UserController{}, "post:GetPlayRecord"),
			beego.NSRouter("/djradio", &controllers.UserController{}, "get:GetDjradioList"),
			beego.NSRouter("/djradioSubed", &controllers.UserController{}, "get:GetDjradioSubedList"),

			// follows
			beego.NSRouter("/follows", &controllers.UserController{}, "get:GetFollows"),
			beego.NSRouter("/followeds", &controllers.UserController{}, "get:GetFolloweds"),

			beego.NSRouter("/event", &controllers.UserController{}, "get:GetEvent"),

			// TODO: uncompleted
			beego.NSRouter("/detail", &controllers.UserController{}, "get:GetUserDetail"),
		),
		beego.NSNamespace("/letter",
			beego.NSRouter("/privateList", &controllers.LetterController{}, "get:GetPrivateLetter"),
			beego.NSRouter("/send", &controllers.LetterController{}, "get:SendPrivateText"),
			beego.NSRouter("/sendWithPlaylist", &controllers.LetterController{}, "get:SendWithPlaylist"),
		),

		beego.NSNamespace("/artist",
			beego.NSRouter("/list", &controllers.ArtistController{}, "get:GetArtistList"),
			beego.NSRouter("/sub", &controllers.ArtistController{}, "get:SubscriptArtist"),
			beego.NSRouter("/unsub", &controllers.ArtistController{}, "get:UnSubscriptArtist"),
			beego.NSRouter("/sublist", &controllers.ArtistController{}, "get:GetSubscriptArtistList"),
		),
		beego.NSNamespace("/artist",
			beego.NSRouter("/list", &controllers.ArtistController{}, "get:GetArtistList"),
			beego.NSRouter("/sub", &controllers.ArtistController{}, "get:SubscriptArtist"),
			beego.NSRouter("/unsub", &controllers.ArtistController{}, "get:UnSubscriptArtist"),
			beego.NSRouter("/sublist", &controllers.ArtistController{}, "get:GetSubscriptArtistList"),
		),
		beego.NSNamespace("/playlist",
			beego.NSRouter("/catlist", &controllers.PlaylistController{}, "get:GetCategoryList"),
			beego.NSRouter("/hot", &controllers.PlaylistController{}, "get:GetHottagsList"),
			beego.NSRouter("/netRecommend", &controllers.PlaylistController{}, "get:GetNetRecommendList"),
			beego.NSRouter("/highQuality", &controllers.PlaylistController{}, "get:GetHighQualityList"),
			beego.NSRouter("/detail", &controllers.PlaylistController{}, "get:GetPlaylistDetail"),
			beego.NSRouter("/create", &controllers.PlaylistController{}, "get:CreatePlaylist"),
			beego.NSRouter("/subscribe", &controllers.PlaylistController{}, "get:SubscribePlaylist"),
			beego.NSRouter("/tracks", &controllers.PlaylistController{}, "get:TrackPlaylistMusic"),
		),
		beego.NSNamespace("/music",
			beego.NSRouter("/musicUrl", &controllers.MusicController{}, "get:GetMusicUrl"),
			beego.NSRouter("/search", &controllers.MusicController{}, "get:SearchMusic"),
			beego.NSRouter("/search/hot", &controllers.MusicController{}, "get:GetHotSearchList"),
			beego.NSRouter("/search/suggest", &controllers.MusicController{}, "get:GetSearchSuggest"),
			beego.NSRouter("/lyric", &controllers.MusicController{}, "get:GetLyric"),
			beego.NSRouter("/detail", &controllers.MusicController{}, "get:GetSongDetail"),
		),

		beego.NSNamespace("/comment",
			beego.NSRouter("/music", &controllers.CommentController{}, "get:GetMusicComment"),
			beego.NSRouter("/album", &controllers.CommentController{}, "get:GetAlbumComment"),
			beego.NSRouter("/playlist", &controllers.CommentController{}, "get:GetPlaylistComment"),
			beego.NSRouter("/mv", &controllers.CommentController{}, "get:GetMVComment"),
			beego.NSRouter("/dj", &controllers.CommentController{}, "get:GetDJComment"),
			beego.NSRouter("/like", &controllers.CommentController{}, "get:LikeComment"),
			beego.NSRouter("/publish", &controllers.CommentController{}, "get:Publish"),
		),

		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
	)

	// register namespace
	beego.AddNamespace(ns)
	beego.AddNamespace(idx)
}
