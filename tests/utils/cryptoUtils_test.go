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
		crypto.CreateSecretKey()
	}

	Convey("Subject: Crypto CreateSecretKey Test Station Endpoint\n", t, func() {
		Convey("Create SecretKey Should Be A String Which Len Is 16", func() {
			So(len(crypto.SecretKey), ShouldEqual, 16)
		})
	})
}

func TestEncrypt(t *testing.T) {
	const originData string = `{"Username":"james","Password":"123","CsrfToken":"token123"}`
	crypto := utils.Crypto{}

	actualResult, _ := crypto.Encrypt(originData)

	Convey("Subject: Crypto Encrypt Test Station Endpoint\n", t, func() {
		Convey("Encrypt should encrypt the data then return base64 result", func() {
			So(len(actualResult), ShouldEqual, 152)
		})
	})
}

func TestDecrypt(t *testing.T) {
	const originData string = `{"Username":"james","Password":"123","CsrfToken":"token123"}`
	const expectationEncryptStr string = "DLAFT+WsUAcj3c7DZb9hspl4yyhLWFfh2h+sqswVOEh4ApN/z27wN4hWDVZhRvDDm+RrEc3+OOeVVEwbSzQeDXCLp8MZee9+FT41zZSnZUcyzY9iMbgjTHO2nkKeMcuTGW6+nzt4F7MKG9VprsgdMw=="

	crypto := utils.Crypto{}

	crypto.SecretKey = "1234567890asdfgh"

	encryptResult, _ := crypto.Encrypt(originData)

	actualResult, _ := crypto.Decrypt(encryptResult)

	Convey("Subject: Crypto Decrypt Test Station Endpoint\n", t, func() {
		Convey("Encrypt Should encrypt originData and return encrypt result", func() {
			So(encryptResult, ShouldEqual, expectationEncryptStr)
		})
		Convey("Decrypt should decrypt the encrypt data and return string result", func() {
			So(actualResult, ShouldEqual, originData)
		})
	})
}
