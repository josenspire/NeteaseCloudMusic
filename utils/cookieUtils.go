package utils

import (
	"strconv"
	"strings"
	"time"
)

type BaseCookie struct {
	BaseCookie string
}

func (bc *BaseCookie) GenerateBaseCookie() {
	randomStr := GenerateRandomString("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKMNOPQRSTUVWXYZ\\/+", 176)
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)

	jsessionid := randomStr + ":" + timestamp
	nuid := GenerateRandomString("0123456789abcdefghijklmnopqrstuvwxyz", 32)

	baseCookie := `JSESSIONID-WYYY=` + jsessionid + `;_iuqxldmzr_=32; _ntes_nnid=` + nuid + "," + strconv.FormatInt(time.Now().UnixNano(), 10) + `;_ntes_nuid=` + nuid
	bc.BaseCookie = baseCookie
}

func generateCookie() string {
	bc := BaseCookie{}
	if strings.EqualFold("", bc.BaseCookie) {
		bc.GenerateBaseCookie()
	}
	const cookieStr = `appver=2.0.3;os=osx; channel=netease;osver=%E7%89%88%E6%9C%AC%2010.13.2%EF%BC%88%E7%89%88%E5%8F%B7%2017C88%EF%BC%89;`
	return cookieStr + bc.BaseCookie
}
