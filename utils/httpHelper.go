package utils

import (
	"fmt"
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

// func NeteaseCloudRequest(baseUrl string, params map[string]interface{}, method string) (string, error) {
//
// 	crypto := Crypto{}
// 	crypto.SecretKey = "BpLnfgDsc2WD8F2q"
// 	_params := `{"phone":"13631270436","password":"e10adc3949ba59abbe56e057f20f883e","rememberLogin":"true","csrf_token":""}`
//
// 	encText, encSeckey, err := crypto.Encrypt(_params)
// 	checkError(err)
//
// 	var paramsMap = `params=` + url.QueryEscape(encText) + `&encSeckey=` + url.QueryEscape(encSeckey)
//
// 	byte, err := json.Marshal(paramsMap)
// 	checkError(err)
//
// 	req, err := http.NewRequest(method, host+baseUrl, bytes.NewBuffer(byte))
// 	SetupHeader(req, strconv.Itoa(len(byte)))
//
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		checkError(err)
// 	}
// 	defer resp.Body.Close()
//
// 	resHost := resp.Request.Host // 有的代理IP被DNS劫持，不干净
// 	if !strings.Contains(resHost, "163") {
// 		return "", errors.New("Request error")
// 	}
//
// 	fmt.Println(req.Body)
//
// 	// statusCode := resp.StatusCode
// 	// hea := resp.Header
//
// 	body, _ := ioutil.ReadAll(resp.Body)
//
// 	return string(body), nil
// }

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
		// "Content-Length":  len,
	}
	for key, value := range headers {
		request.Header(key, value)
	}
}

func NeteaseCloudRequest(baseUrl string, params map[string]interface{}, method string) (string, error) {
	req := httplib.NewBeegoRequest(host+baseUrl, method)

	// _params, _ := json.Marshal(params)

	crypto := Crypto{}
	crypto.SecretKey = "KLanfgDsc2WD8F2q"
	_params := `{"phone":"13631270438","password":"e10adc3949ba59abbe56e057f20f883e","rememberLogin":"true"}`

	encText, encSeckey, err := crypto.Encrypt(_params)
	checkError(err)

	// paramsBody := "params=PaBwf0ljoojLjWSjRWn6mKPWndhYwSLDhHnEUbkSdjpXCHb6ACx08uuTXcnqjmhhvjBIeClm%2FTcqyjEiwKsIFIfnD5%2FUYCulG8c4LjzuKpwToPYSiYaMFxE6aq02CI5BEOsJklkviywLaS95l37OmXPS40Kxu7KuFMke0FyQeOXfo6JPD0Vz6qsht34Kts2F&encSecKey=2e983589cf245726cae4d87690680ec0f58b30948bd99e6698f1d9270bfd12d869c9a54e0ae8885801ab01d16c60bc39420a102907c509a9671a8338932bfd500d3d1560cb2ffaa3e308c8b962a62e1d4c0ffbaf044ca6b41ea8932ad88b1d8355c1e48984c25af6f9ef3dd2ffad216aaeb7cdf8dba533fcef099286ce98e617"

	paramsBody := "params=" + url.QueryEscape(encText) + "&encSeckey=" + url.QueryEscape(encSeckey)
	fmt.Println("=========", paramsBody)

	// paramsMap = "params=ufNHvI59iN8TFbhtq9L%2F%2BKDk3QryYKHD7RBZBmPL2mS2KgGpEBSQ0wio8D15ZAwCt8PZTW97OO9hC1LQdirkliPA3VSJgTr2GdBimBGW2NNfcUGq7D%2FBNdG0w0qx1sssPyiEzgh0oGRek3ljZxCgfnXMj5uKLPmu497yn%2FqAumVIWtWGFSLnIKW2U7868cWQgf8nskis1gNmxYtTId6AfA%3D%3D&encSecKey=0d3955198b0c79302c116df6196580b6d464d1a06c2f7aa7a798ac4c7d7143ef7edf25dbc3a85d9b420380903b482b8ae1416e35b13cb8c375a80011602ed260ac709fb2873531f2e099ebbfa49d53d6531083f43b7eb65a4acadbbb9a105a5f4df1e3671d85c802183895c254ccf034630d89adb0e89ff5bfc35c36f4c15b4e"
	// paramsMap := "params=%2BcVGZTV4RnOAXyBj4zhb%2FV6MaIIWRJQAupMkqHmTGiF3AxMTQnGitQZkF5EbZJGu%2ByoruGwdvnd%2Blbr5K90s2ndKr00qkjqeXt%2Bfz00yFG9QJyf1fzuAo7pYxFB7DspLHWwpGGFsbaihzzt63FpOPDoQIXkaOb8G2awBnXu4M0n8pRPo2b5JUuP%2BJzZ3icfhbU4OA1OUm%2B2DJNFcaqExjw%3D%3D&encSecKey=4bfbe6157d2674355339fb2d6a697f664ffa4c8b43cc105d0d437f674b3f77aaf361507060bdea58f610aef07929fd8c3cec616f5510ecf7b8f3b4c8bb013dd66c0e99a3a523cc1e06be5f2112f1565c0037c089a6d886161f31cf3d4c5e5494083504c8d3a1ab48aceed51a877e006e2446d09c3d1e0078aaa7d042db9c78e8"

	SetupHeader(req)
	req.Body(string(paramsBody))
	result, _ := req.String()

	fmt.Println("======================:", req.GetRequest())
	fmt.Println("===========Request Body: ", req.GetRequest().Body)
	fmt.Println("===========Response Body: ", result)

	req.DoRequest()
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
