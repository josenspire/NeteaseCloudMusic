package utils

import (
	"strconv"
	"time"
)

func GenerateBaseCookie() string {
	randomStr := GenerateRandomString("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKMNOPQRSTUVWXYZ\\/+", 176)
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)

	jsessionid := randomStr + ":" + timestamp
	nuid := GenerateRandomString("0123456789abcdefghijklmnopqrstuvwxyz", 32)

	return `JSESSIONID-WYYY=` + jsessionid + `;_iuqxldmzr_=32; _ntes_nnid=` + nuid + "," + strconv.FormatInt(time.Now().UnixNano(), 10) + `;_ntes_nuid=` + nuid
}
