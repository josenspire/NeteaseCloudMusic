package test

import (
	"NeteaseCloudMusic/utils"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGenerateBaseCookie(t *testing.T) {
	baseCookies := utils.GenerateBaseCookie()

	Convey("Subject: CookieUtils Test Station Endpoint\n", t, func() {
		Convey("Generate BaseCookie should return a cookie array which len is 3", func() {
			So(len(baseCookies), ShouldEqual, 4)
		})
		Convey("Generate BaseCookie should include `JSESSIONID-WYYY`, `_iuqxldmzr_`, `_ntes_nnid`, `_ntes_nuid`", func() {
			So(baseCookies[0].Name, ShouldEqual, "JSESSIONID-WYYY")
			So(baseCookies[1].Name, ShouldEqual, "_iuqxldmzr_")
			So(baseCookies[2].Name, ShouldEqual, "_ntes_nnid")
			So(baseCookies[3].Name, ShouldEqual, "_ntes_nuid")
		})
	})
}
