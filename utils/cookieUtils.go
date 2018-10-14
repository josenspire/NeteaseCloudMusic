package utils

import (
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func GenerateBaseCookie() []*http.Cookie {
	randomStr := GenerateRandomString("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKMNOPQRSTUVWXYZ\\/+", 176)
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)

	jsessionid := randomStr + ":" + timestamp
	nuid := GenerateRandomString("0123456789abcdefghijklmnopqrstuvwxyz", 32)

	// cookieStr := `JSESSIONID-WYYY=` + jsessionid + `;_iuqxldmzr_=32;_ntes_nnid=` + nuid + "," + strconv.FormatInt(time.Now().UnixNano(), 10) + `;_ntes_nuid=` + nuid
	baseCookies := make([]*http.Cookie, 4)
	baseCookies[0] = &http.Cookie{Name: `JSESSIONID-WYYY`, Value: url.QueryEscape(jsessionid)}
	baseCookies[1] = &http.Cookie{Name: `_iuqxldmzr_`, Value: "32"}
	baseCookies[2] = &http.Cookie{Name: `_ntes_nnid`, Value: nuid + "," + strconv.FormatInt(time.Now().UnixNano(), 10)}
	baseCookies[3] = &http.Cookie{Name: `_ntes_nuid`, Value: nuid}

	return baseCookies
}
