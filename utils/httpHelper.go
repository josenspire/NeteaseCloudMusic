package utils

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	host             = "http://music.163.com"
	ConnectTimeOut   = time.Second * 30
	ReadWriteTimeOut = time.Second * 30
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

func setupHeader(request *httplib.BeegoHTTPRequest) {
	headers := map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4",
		"Connection":      "keep-alive",
		"Content-Type":    "application/x-www-form-urlencoded",
		"Referer":         "http://music.163.com",
		"Host":            "music.163.com",
		"Cookie":          "appver=2.0.2",
		"User-Agent":      randomUserAgent(),
	}
	for key, value := range headers {
		request.Header(key, value)
	}
}

func setupCookies(request *httplib.BeegoHTTPRequest, cookies []*http.Cookie) {
	for _, cookie := range cookies {
		request.SetCookie(cookie)
	}

}

func NeteaseCloudRequest(baseUrl string, params map[string]interface{}, cookies []*http.Cookie, method string) (interface{}, []*http.Cookie, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	params["csrf_token"] = GetCookieValueByName(cookies, "__csrf")
	beego.Info("[REQUEST]", params)

	req := httplib.NewBeegoRequest(host+baseUrl, method).SetTimeout(ConnectTimeOut, ReadWriteTimeOut).SetEnableCookie(true)
	baseCookie := GenerateBaseCookie()

	// setup request headers
	setupHeader(req)
	// setup request cookies
	setupCookies(req, append(baseCookie, cookies...))

	crypto := Crypto{}
	reqParams, _ := json.Marshal(params)
	encText, encSecKey, err := crypto.Encrypt(string(reqParams))
	if err != nil {
		return nil, nil, err
	}
	paramsBody := "params=" + url.QueryEscape(encText) + "&encSecKey=" + encSecKey
	req.Body(paramsBody)

	var jsonObj interface{}
	req.ToJSON(&jsonObj)

	res, _ := req.Response()
	neteaseCookies := res.Cookies()

	replaceCookieDomain(neteaseCookies)
	return jsonObj, neteaseCookies, nil
}

// replace response's domain scope
func replaceCookieDomain(cookies []*http.Cookie) {
	for _, cookie := range cookies {
		cookie.Domain = strings.Replace(cookie.Domain, ".music.163.com", "", -1)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error happend: %s", err.Error())
		return
	}
}
