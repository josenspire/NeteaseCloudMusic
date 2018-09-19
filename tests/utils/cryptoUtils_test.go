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
	const originData string = `{"Username":"13631270436","Password":"123456789"}`
	crypto := utils.Crypto{}

	actualResult, _, _ := crypto.Encrypt(originData)

	Convey("Subject: Crypto Encrypt Test Station Endpoint\n", t, func() {
		Convey("Encrypt should encrypt the data then return base64 result", func() {
			So(len(actualResult), ShouldEqual, 152)
		})
	})
}

func TestDecrypt(t *testing.T) {
	const originData string = `{"Username":"james","Password":"123","CsrfToken":"token123"}`
	const expectationEncryptStr string = "DLAFT%2BWsUAcj3c7DZb9hspl4yyhLWFfh2h%2BsqswVOEh4ApN%2Fz27wN4hWDVZhRvDDm%2BRrEc3%2BOOeVVEwbSzQeDXCLp8MZee9%2BFT41zZSnZUcyzY9iMbgjTHO2nkKeMcuTGW6%2Bnzt4F7MKG9VprsgdMw%3D%3D"

	crypto := utils.Crypto{}

	crypto.SecretKey = "1234567890asdfgh"

	encryptResult, _, _ := crypto.Encrypt(originData)

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

func TestRSAEncrypt(t *testing.T) {
	const originData string = `{"Username": "james", "Password": "123456"}`
	const expectation string = `858130a0a5e6c330d4ca8d439841147f1f131ddb7d91e8b2cd2d00556c7e84702b12e0e3347d387e2dbbe53fcefa04008936075bb46b0c1c99c90dcceda33a14900ec3a74579822eebcd1208a8beee14d69180d131aa64100f81a21436d00ebd48ff69f800282c94420dc8d22bad7ee4cd0f48214f81f0eb1484bd60d644f89e`

	const (
		pubKey = "010001"
		iv     = "0102030405060708"
	)

	crypto := utils.Crypto{}

	actualResult := crypto.RSAEncrypt(originData, pubKey, iv)

	Convey("Subject: Crypto RSADecrypt Test Station Endpoint\n", t, func() {
		Convey("RSAEncrypt should encrypt originData with rsa and return result", func() {
			So(actualResult, ShouldEqual, expectation)
		})
	})
}
