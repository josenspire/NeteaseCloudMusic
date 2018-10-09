package utils

import (
	"github.com/astaxie/beego/httplib"
	"log"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

const (
	host = "http://music.163.com"
)

var userAgentList = []string{
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
	"Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 5.1.1; Nexus 6 Build/LYZ28E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89;GameHelper",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:46.0) Gecko/20100101 Firefox/46.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:46.0) Gecko/20100101 Firefox/46.0",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0)",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)",
	"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)",
	"Mozilla/5.0 (Windows NT 6.3; Win64, x64; Trident/7.0; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/13.10586",
	"Mozilla/5.0 (iPad; CPU OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
}

func randomUserAgent() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return userAgentList[r.Intn(19)]
}

func SetupHeader(request *httplib.BeegoHTTPRequest) {
	headers := map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4",
		"Connection":      "keep-alive",
		"Content-Type":    "application/x-www-form-urlencoded",
		"Referer":         "http://music.163.com",
		"Host":            "music.163.com",
		"Cookie":          "appver=2.0.2",
		"User-Agent":      randomUserAgent(),
		// "Content-Length":  strconv.Itoa(contentLength),
	}
	for key, value := range headers {
		request.Header(key, value)
	}
}

func NeteaseCloudRequest(baseUrl string, params map[string]interface{}, method string) (string, error) {
	req := httplib.NewBeegoRequest(host+baseUrl, method)
	// req := httplib.NewBeegoRequest("http://localhost:3010/googlemap", method)
	// _params, _ := json.Marshal(params)

	crypto := Crypto{}
	crypto.SecretKey = "KLanfgDsc2WD8F2q"
	_params := `{"phone":"13631270438","password":"e10adc3949ba59abbe56e057f20f883e","rememberLogin":"true"}`
	encText, encSeckey, _ := crypto.Encrypt(_params)

	paramsBody := "params=" + url.QueryEscape(encText) + "&encSecKey=" + encSeckey

	SetupHeader(req)
	req.Body(paramsBody)

	result, _ := req.String()
	return result, nil
}

func generateCookie() string {
	bc := BaseCookie{}
	if strings.EqualFold("", bc.BaseCookie) {
		bc.GenerateBaseCookie()
	}
	const cookieStr = `appver=2.0.3;os=osx; channel=netease;osver=%E7%89%88%E6%9C%AC%2010.13.2%EF%BC%88%E7%89%88%E5%8F%B7%2017C88%EF%BC%89;`
	return cookieStr + bc.BaseCookie
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error happend: %s", err.Error())
		return
	}
}
