package utils

import (
	"NeteaseCloudMusic/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
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
	num := math.Floor(r.Float64() * float64(len(userAgentList)))
	return userAgentList[int(num)]
}

func NeteaseCloudRequest(url string, params models.RequestObject, method string) models.Response {
	response := models.Response{}

	object, error := json.Marshal(params)
	if error != nil {
		log.Fatalf("Format params error: %s", error.Error())
		response.SetResponse(http.StatusBadRequest, nil, "")
		return response
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(object))

	bc := BaseCookie{}
	bc.GenerateBaseCookie()
	cookie := bc.BaseCookie

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "http://music.163.com")
	req.Header.Set("Host", "music.163.com")
	req.Header.Set("User-Agent", randomUserAgent())

	if params.Cookie != "" {
		cookie = cookie + "; " + params.Cookie
	}
	req.Header.Set("Cookie", cookie)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		checkError(err)
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	hea := resp.Header
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(hea)

	response.SetResponse(statusCode, body, "")
	response.PrintResponse()
	return response
}

func checkError(e error) {
	if e != nil {
		log.Fatalf("Error happend: %s", e.Error())
		return
	}
}
