package coinmarketcap

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

func btoa(str []byte) string {
	return base64.StdEncoding.EncodeToString(str)
}

func charCodeAt(a string, i int) int {
	return int(a[i])
}

// FromCharCode https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/FromCharCode
func fromCharCode(c int) string {
	return string(rune(c))
}

// >>> operator in golang
func trippleShift(num, t int) int {
	overflow := int32(num)
	return int(uint32(overflow) >> t)
}

func shaPassword(password string) string {
	sha512Hasher := sha512.New()
	sha512Hasher.Write([]byte(password))
	return hex.EncodeToString(sha512Hasher.Sum(nil))
}
