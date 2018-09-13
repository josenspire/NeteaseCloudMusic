package utils

import (
	"math"
	"math/rand"
	"time"
	"strconv"
)

type BaseCookie struct {
	baseCookie string
}

func randomString(pattern string, length int) string {

	var newStr string

	for i := 0; i <= length; i++ {
		r := rand.New(rand.NewSource(time.Now().Unix() * int64(i)))
		pi := int(math.Floor(r.Float64() * float64(len(pattern))))

		newStr += pattern[pi : pi+1]
	}
	return newStr
}

func (bc *BaseCookie) GenerateBaseCookie() {
	jsessionid := randomString("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKMNOPQRSTUVWXYZ\\/+", 176) + ":" + strconv.FormatInt(time.Now().UnixNano(), 10)
	nuid := randomString("0123456789abcdefghijklmnopqrstuvwxyz", 32)

	baseCookie := `JSESSIONID-WYYY=` + jsessionid + ";" + `_iuqxldmzr_=32; _ntes_nnid=` + nuid + "," + strconv.FormatInt(time.Now().UnixNano(), 10) + ";_ntes_nuid=" + nuid
	bc.baseCookie = baseCookie
}