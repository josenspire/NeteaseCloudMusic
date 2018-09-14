package utils

import (
	"strconv"
	"time"
)

type BaseCookie struct {
	BaseCookie string
}

func (bc *BaseCookie) GenerateBaseCookie() {
	jsessionid := GenerateRandomString("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKMNOPQRSTUVWXYZ\\/+", 176) + ":" + strconv.FormatInt(time.Now().UnixNano(), 10)
	nuid := GenerateRandomString("0123456789abcdefghijklmnopqrstuvwxyz", 32)

	baseCookie := `JSESSIONID-WYYY=` + jsessionid + ";" + `_iuqxldmzr_=32; _ntes_nnid=` + nuid + "," + strconv.FormatInt(time.Now().UnixNano(), 10) + ";_ntes_nuid=" + nuid
	bc.BaseCookie = baseCookie
}
