package test

import (
	"NeteaseCloudMusic/utils"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// func init() {
// 	_, file, _, _ := runtime.Caller(1)
// 	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../../.." + string(filepath.Separator))))
// 	beego.TestBeegoInit(appPath)
// }

func TestGenerateRandomString(t *testing.T) {
	const originStr string = "1234567890asdfghjklqweryuiopzxcvnm"
	actualResult := utils.GenerateRandomString(originStr, 10)

	Convey("Subject: StringUtils Test Station Endpoint\n", t, func() {
		Convey("Generate Random String should return a string witch length is 10", func() {
			So(len(actualResult), ShouldEqual, 10)
		})
		Convey("Generate Random String should return diff string", func() {
			So(actualResult, ShouldNotEqual, "1234567890asdfghjklqweryuiopzxcvnm")
		})
	})

}
