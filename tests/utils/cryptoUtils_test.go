package test

import (
	"NeteaseCloudMusic/utils"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../../.."+string(filepath.Separator))))
	beego.TestBeegoInit(appPath)
}

func TestCreateSecretKey(t *testing.T) {
	crypto := utils.Crypto{}

	if strings.EqualFold(crypto.SecretKey, "") {
		crypto.CreateSecretKey(127)
	}

	Convey("Subject: Crypto Test Station Endpoint\n", t, func() {
		Convey("Create SecretKey Should Be A String Which Len Is 127", func() {
			So(len(crypto.SecretKey), ShouldEqual, 127)
		})
	})

}
