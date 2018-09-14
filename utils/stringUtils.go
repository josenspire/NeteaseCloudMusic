package utils

import (
	"math"
	"math/rand"
	"time"
)

func GenerateRandomString(originStr string, length int) string {
	target := ""
	for i := 0; i < length; i++ {
		r := rand.New(rand.NewSource(time.Now().Unix() * int64(i)))
		pi := uint8(math.Floor(r.Float64() * float64(len(originStr))))

		target += string(originStr[pi])
	}

	return target
}
