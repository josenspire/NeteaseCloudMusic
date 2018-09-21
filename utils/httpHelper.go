package utils

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
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

func NeteaseCloudRequest(baseUrl string, params map[string]interface{}, method string) (string, error) {
	// crypto := Crypto{}
	// crypto.SecretKey = "BpLnfgDsc2WD8F2q"
	// _params := `{"phone":"13631270436","password":"e10adc3949ba59abbe56e057f20f883e","rememberLogin":"true","csrf_token":""}`
	// encText, encSeckey, err := crypto.Encrypt(_params)
	// checkError(err)
	// var paramsMap = `params=` + url.QueryEscape(encText) + `&encSeckey=` + encSeckey

	// var paramsBody = "params=PaBwf0ljoojLjWSjRWn6mKPWndhYwSLDhHnEUbkSdjpXCHb6ACx08uuTXcnqjmhhvjBIeClm%2fTcqyjEiwKsIFIfnD5%2fUYCulG8c4LjzuKpwToPYSiYaMFxE6aq02CI5BEOsJklkviywLaS95l37OmXPS40Kxu7KuFMke0FyQeOXfo6JPD0Vz6qsht34Kts2F&encSecKey=2e983589cf245726cae4d87690680ec0f58b30948bd99e6698f1d9270bfd12d869c9a54e0ae8885801ab01d16c60bc39420a102907c509a9671a8338932bfd500d3d1560cb2ffaa3e308c8b962a62e1d4c0ffbaf044ca6b41ea8932ad88b1d8355c1e48984c25af6f9ef3dd2ffad216aaeb7cdf8dba533fcef099286ce98e617"

	encText := "PaBwf0ljoojLjWSjRWn6mKPWndhYwSLDhHnEUbkSdjpXCHb6ACx08uuTXcnqjmhhvjBIeClm/TcqyjEiwKsIFIfnD5/UYCulG8c4LjzuKpwToPYSiYaMFxE6aq02CI5BEOsJklkviywLaS95l37OmXPS40Kxu7KuFMke0FyQeOXfo6JPD0Vz6qsht34Kts2F"
	encSeckey := "2e983589cf245726cae4d87690680ec0f58b30948bd99e6698f1d9270bfd12d869c9a54e0ae8885801ab01d16c60bc39420a102907c509a9671a8338932bfd500d3d1560cb2ffaa3e308c8b962a62e1d4c0ffbaf044ca6b41ea8932ad88b1d8355c1e48984c25af6f9ef3dd2ffad216aaeb7cdf8dba533fcef099286ce98e617"

	client := &http.Client{}

	form := url.Values{}
	form.Set("params", encText)
	form.Set("encSecKey", encSeckey)
	body := strings.NewReader(form.Encode())
	req, _ := http.NewRequest("POST", host+baseUrl, body)
	SetupRequestHeader(req)

	response, err := client.Do(req)
	checkError(err)

	defer response.Body.Close()
	resBody, resErr := ioutil.ReadAll(response.Body)
	checkError(resErr)

	fmt.Println(string(resBody))
	return string(resBody), nil
}

func NeteaseCloudRequest1(baseUrl string, params map[string]interface{}, method string) (string, error) {
	// crypto := Crypto{}
	// crypto.SecretKey = "BpLnfgDsc2WD8F2q"
	// _params := `{"phone":"13631270436","password":"e10adc3949ba59abbe56e057f20f883e","rememberLogin":"true","csrf_token":""}`
	// encText, encSeckey, err := crypto.Encrypt(_params)
	// checkError(err)
	// var paramsMap = `params=` + url.QueryEscape(encText) + `&encSeckey=` + encSeckey

	var paramsBody = "params=PaBwf0ljoojLjWSjRWn6mKPWndhYwSLDhHnEUbkSdjpXCHb6ACx08uuTXcnqjmhhvjBIeClm%2fTcqyjEiwKsIFIfnD5%2fUYCulG8c4LjzuKpwToPYSiYaMFxE6aq02CI5BEOsJklkviywLaS95l37OmXPS40Kxu7KuFMke0FyQeOXfo6JPD0Vz6qsht34Kts2F&encSecKey=2e983589cf245726cae4d87690680ec0f58b30948bd99e6698f1d9270bfd12d869c9a54e0ae8885801ab01d16c60bc39420a102907c509a9671a8338932bfd500d3d1560cb2ffaa3e308c8b962a62e1d4c0ffbaf044ca6b41ea8932ad88b1d8355c1e48984c25af6f9ef3dd2ffad216aaeb7cdf8dba533fcef099286ce98e617"

	client := &http.Client{}
	req, err := http.NewRequest(method, host+baseUrl, strings.NewReader(paramsBody))
	checkError(err)
	SetupRequestHeader(req)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	resHost := resp.Request.Host // 有的代理IP被DNS劫持，不干净
	if !strings.Contains(resHost, "163") {
		return "", errors.New("Request error")
	}

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Println(string(body))

	return string(body), nil

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// checkError(err)
	// defer resp.Body.Close()

	// resHost := resp.Request.Host // 有的代理IP被DNS劫持，不干净
	// if !strings.Contains(resHost, "163") {
	// 	return "", errors.New("Request error")
	// }

	// fmt.Println("request Body: ", req)

	// statusCode := resp.StatusCode
	// hea := resp.Header

	// body, _ := ioutil.ReadAll(resp.Body)
	//
	// responseBody := string(body)
	// fmt.Println("=============", responseBody)

	// return responseBody, nil
}

func SetupRequestHeader(request *http.Request) {
	headers := map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4",
		"Connection":      "keep-alive",
		"Content-Type":    "application/x-www-form-urlencoded",
		"Referer":         "http://music.163.com",
		"Host":            "music.163.com",
		"Cookie":          "appver=2.0.2",
		"User-Agent":      randomUserAgent(),
		// "Content-Length":  len,
	}
	for key, value := range headers {
		request.Header.Set(key, value)
	}
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

func NeteaseCloudRequest2(baseUrl string, params map[string]interface{}, method string) (string, error) {
	req := httplib.NewBeegoRequest(host+baseUrl, method)
	// req := httplib.NewBeegoRequest("http://localhost:3010/googlemap", method)
	// _params, _ := json.Marshal(params)

	crypto := Crypto{}
	crypto.SecretKey = "KLanfgDsc2WD8F2q"
	_params := `{"phone":"13631270438","password":"e10adc3949ba59abbe56e057f20f883e","rememberLogin":"true"}`
	encText, encSeckey, _ := crypto.Encrypt(_params)

	var paramsBody string = "params=" + url.QueryEscape(encText) + "&encSeckey=" + encSeckey

	// _encText := "PaBwf0ljoojLjWSjRWn6mKPWndhYwSLDhHnEUbkSdjpXCHb6ACx08uuTXcnqjmhhvjBIeClm/TcqyjEiwKsIFIfnD5/UYCulG8c4LjzuKpwToPYSiYaMFxE6aq02CI5BEOsJklkviywLaS95l37OmXPS40Kxu7KuFMke0FyQeOXfo6JPD0Vz6qsht34Kts2F"
	// _encSeckey := "2e983589cf245726cae4d87690680ec0f58b30948bd99e6698f1d9270bfd12d869c9a54e0ae8885801ab01d16c60bc39420a102907c509a9671a8338932bfd500d3d1560cb2ffaa3e308c8b962a62e1d4c0ffbaf044ca6b41ea8932ad88b1d8355c1e48984c25af6f9ef3dd2ffad216aaeb7cdf8dba533fcef099286ce98e617"
	// paramsBody = "params=" + url.QueryEscape(_encText) + "&encSecKey=" + _encSeckey

	var _paramsBody = "params=PaBwf0ljoojLjWSjRWn6mKPWndhYwSLDhHnEUbkSdjpXCHb6ACx08uuTXcnqjmhhvjBIeClm%2fTcqyjEiwKsIFIfnD5%2fUYCulG8c4LjzuKpwToPYSiYaMFxE6aq02CI5BEOsJklkviywLaS95l37OmXPS40Kxu7KuFMke0FyQeOXfo6JPD0Vz6qsht34Kts2F&encSecKey=2e983589cf245726cae4d87690680ec0f58b30948bd99e6698f1d9270bfd12d869c9a54e0ae8885801ab01d16c60bc39420a102907c509a9671a8338932bfd500d3d1560cb2ffaa3e308c8b962a62e1d4c0ffbaf044ca6b41ea8932ad88b1d8355c1e48984c25af6f9ef3dd2ffad216aaeb7cdf8dba533fcef099286ce98e617"

	fmt.Println("Is Equal:", strings.EqualFold(paramsBody, _paramsBody))
	fmt.Println("Is Compare:", paramsBody, _paramsBody)

	SetupHeader(req)
	req.Body(_paramsBody)

	fmt.Println("===========Request Body:", req.GetRequest().Body)
	result, _ := req.String()

	fmt.Println("===========Response Body: ", result)
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
