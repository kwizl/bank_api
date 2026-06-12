package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt Random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

// RandomString Generates random string based on alphabet characters
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "GBP", "USD", "CNY", "JPY"}
	b := len(currencies)
	return currencies[rand.Intn(b)]
}
