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
		crypto.CreateSecretKey(16)
	}

	Convey("Subject: Crypto CreateSecretKey Test Station Endpoint\n", t, func() {
		Convey("Create SecretKey Should Be A String Which Len Is 16", func() {
			So(len(crypto.SecretKey), ShouldEqual, 16)
		})
	})
}

func TestEncrypt(t *testing.T) {
	const (
		originData      = `{"Username":"13631270438","Password":"e10adc3949ba59abbe56e057f20f883e"}`
		expectEncText   = "MDdJP8+CeWRWJDscZ19LenNYNt3TKjPu0RmMdK5YLgUYMXmW5kYV7bFs1FSEUCZ+ZEWwM30lzdoqQu0+QprEhEVWk76HQpSSF4gF5htnQupxbgYODtaECJfNDD0oY9DcSmjX0vTUH7IIXlBx8tM1Bw=="
		expectEncSeckey = "2e983589cf245726cae4d87690680ec0f58b30948bd99e6698f1d9270bfd12d869c9a54e0ae8885801ab01d16c60bc39420a102907c509a9671a8338932bfd500d3d1560cb2ffaa3e308c8b962a62e1d4c0ffbaf044ca6b41ea8932ad88b1d8355c1e48984c25af6f9ef3dd2ffad216aaeb7cdf8dba533fcef099286ce98e617"
	)
	crypto := utils.Crypto{
		SecretKey: "KLanfgDsc2WD8F2q",
	}

	encText, encSecKey, _ := crypto.Encrypt(originData)

	Convey("Subject: Crypto Encrypt Test Station Endpoint\n", t, func() {
		Convey("Encrypt should encrypt the data then return base64 encText & encSecKey", func() {
			So(encText, ShouldEqual, expectEncText)
			So(encSecKey, ShouldEqual, expectEncSeckey)
		})
	})
}

// func TestDecrypt(t *testing.T) {
// 	const (
// 		originData            = `{"Username":"13631270438","Password":"e10adc3949ba59abbe56e057f20f883e"}`
// 		expectationEncryptStr = "MDdJP8+CeWRWJDscZ19LenNYNt3TKjPu0RmMdK5YLgUYMXmW5kYV7bFs1FSEUCZ+ZEWwM30lzdoqQu0+QprEhEVWk76HQpSSF4gF5htnQupxbgYODtaECJfNDD0oY9DcSmjX0vTUH7IIXlBx8tM1Bw=="
// 	)
// 	crypto := utils.Crypto{
// 		SecretKey: "KLanfgDsc2WD8F2q",
// 	}
//
// 	encText, _, _ := crypto.Encrypt(originData)
//
// 	actualResult, _ := crypto.Decrypt(encText)
//
// 	Convey("Subject: Crypto Decrypt Test Station Endpoint\n", t, func() {
// 		Convey("Encrypt Should encrypt originData and return encrypt result", func() {
// 			So(encText, ShouldEqual, expectationEncryptStr)
// 		})
// 		Convey("Decrypt should decrypt the encrypt data and return string result", func() {
// 			So(actualResult, ShouldEqual, originData)
// 		})
// 	})
// }

func TestRSAEncrypt(t *testing.T) {
	const originData string = `KLanfgDsc2WD8F2q`
	const expectation string = `2e983589cf245726cae4d87690680ec0f58b30948bd99e6698f1d9270bfd12d869c9a54e0ae8885801ab01d16c60bc39420a102907c509a9671a8338932bfd500d3d1560cb2ffaa3e308c8b962a62e1d4c0ffbaf044ca6b41ea8932ad88b1d8355c1e48984c25af6f9ef3dd2ffad216aaeb7cdf8dba533fcef099286ce98e617`

	const (
		pubKey  = "010001"
		modulus = "00e0b509f6259df8642dbc35662901477df22677ec152b5ff68ace615bb7b725152b3ab17a876aea8a5aa76d2e417629ec4ee341f56135fccf695280104e0312ecbda92557c93870114af6c9d05c4f7f0c3685b7a46bee255932575cce10b424d813cfe4875d3e82047b97ddef52741d546b8e289dc6935b3ece0462db0a22b8e7"
	)

	crypto := utils.Crypto{}
	actualResult := crypto.RSAEncrypt(originData, pubKey, modulus)

	Convey("Subject: Crypto RSADecrypt Test Station Endpoint\n", t, func() {
		Convey("RSAEncrypt should encrypt originData with rsa and return result", func() {
			So(actualResult, ShouldEqual, expectation)
		})
	})
}
