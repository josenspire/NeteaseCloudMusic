package utils

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"log"
	"math"
	"math/rand"
	"strings"
	"time"
)

type Response struct {
	Code    int
	Data    []byte
	Message string
}

func (rs *Response) SetResponse(code int, data []byte, message string) {
	rs.Code = code
	rs.Data = data
	rs.Message = message
}

func (rs *Response) PrintResponse() {
	fmt.Println(rs.Code)
	fmt.Println(rs.Data)
	fmt.Println(rs.Message)
}

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
	num := math.Floor(r.Float64() * float64(len(userAgentList)))
	return userAgentList[int(num)]
}

// func NeteaseCloudRequest(url string, params interface{}, method string) ([]byte, error) {
// 	response := Response{}
//
// 	byte, err := json.Marshal(params)
// 	checkError(err)
//
// 	req, err := http.NewRequest(method, url, bytes.NewBuffer(byte))
//
// 	SetupHeader(req)
//
// 	// if params.Cookie != "" {
// 	// 	cookie = cookie + "; " + params.Cookie
// 	// }
// 	// req.Header.Set("Cookie", cookie)
//
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		checkError(err)
// 	}
// 	defer resp.Body.Close()
//
// 	resqHost := resp.Request.Host // 有的代理IP被DNS劫持，不干净
//
// 	if !strings.Contains(resqHost, "163") {
// 		response.SetResponse(http.StatusBadRequest, nil, "Request is error")
// 		return json.Marshal(response)
// 	}
//
// 	statusCode := resp.StatusCode
// 	// hea := resp.Header
//
// 	body, _ := ioutil.ReadAll(resp.Body)
//
// 	response.SetResponse(statusCode, body, "")
// 	return json.Marshal(response)
// }

func NeteaseCloudRequest(url string, params map[string]string, method string) Response {
	req := httplib.NewBeegoRequest(url, method)
	SetupHeader(req)

	for key, value := range params {
		req.Param(key, value)
	}

	res := Response{}
	// req.ToJSON(&res)

	_req := req.GetRequest()
	fmt.Println("=====", _req)

	for k, v := range _req.Header {
		fmt.Println(k, v)
	}

	return res
}

func SetupHeader(req *httplib.BeegoHTTPRequest) {
	bc := BaseCookie{}
	bc.GenerateBaseCookie()

	headers := map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4",
		"Connection":      "keep-alive",
		"Content-Type":    "application/x-www-form-urlencoded",
		"Referer":         "http://music.163.com",
		"Host":            "music.163.com",
		"Cookie":          bc.BaseCookie,
		"User-Agent":      randomUserAgent(),
	}

	for key, value := range headers {
		req.Header(key, value)
	}
}

func generateCookie(cookie *[]string) {
	const cookieStr = `appver=1.5.9;os=osx; channel=netease;osver=%E7%89%88%E6%9C%AC%2010.13.2%EF%BC%88%E7%89%88%E5%8F%B7%2017C88%EF%BC%89`
	cookieArray := strings.Split(cookieStr, ";")
	for _, ck := range cookieArray {
		_ = append(*cookie, ck+";Path=/")
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error happend: %s", err.Error())
		return
	}
}
