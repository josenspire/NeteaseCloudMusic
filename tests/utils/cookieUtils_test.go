package test

import (
	"NeteaseCloudMusic/utils"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGenerateBaseCookie(t *testing.T) {
	bc := utils.BaseCookie{}

	bc.GenerateBaseCookie()
	Convey("Subject: CookieUtils Test Station Endpoint\n", t, func() {
		Convey("Generate BaseCookie should setup BaseCookie struct attribute", func() {
			So(bc.BaseCookie, ShouldNotBeEmpty)
		})
		Convey("Generate BaseCookie should include `JSESSIONID-WYYY`", func() {
			So(bc.BaseCookie, ShouldContainSubstring, "JSESSIONID-WYYY")
		})
	})
}
