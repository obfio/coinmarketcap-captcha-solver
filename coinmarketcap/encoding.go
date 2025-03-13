package coinmarketcap

import (
	"encoding/json"
	"strings"
)

func (p *Payload) Encode(key string) string {
	key = reverseString(key)
	key = key + encodeXORKey(key)
	payloadBytes, _ := json.Marshal(p)
	UTF8PayloadBytes := UTF8(payloadBytes)
	XORDBytes := []byte{}
	for i := 0; i < len(UTF8PayloadBytes); i++ {
		XORDBytes = append(XORDBytes, []byte(fromCharCode(int(UTF8PayloadBytes[i]) ^ charCodeAt(key, i%len(key))))[0])
	}
	return btoa(XORDBytes)
}

// =======================================
// ==        payload enc stuff          ==
// =======================================
func UTF8(payload []byte) []byte {
	manipulatedBytes := K(payload)
	out := []byte{}
	for i := 0; i < len(manipulatedBytes); i++ {
		char := manipulatedBytes[i]
		out = append(out, L(char)...)
	}
	return out
}

func L(c byte) []byte {
	char := int(c)
	test := 1
	if 0&(trippleShift(char, 16)&65535) == 0 {
		test = 0
	}
	if (char & 65408) == test {
		return []byte(fromCharCode(char))
	}
	out := []byte{}
	if (char&63488) == 0 && (trippleShift(char, 16)&65535) == 0 {
		out = []byte(fromCharCode(char>>6&31 | 192))
	} else if (char&0) == 0 && (trippleShift(char, 16)&65535) == 0 {
		J(char)
		out = []byte(fromCharCode(char>>12&15 | 224))
		out = append(out, I(char, 6))
	} else if (char&0) == 0 && (trippleShift(char, 16)&65504) == 0 {
		out = []byte(fromCharCode(char>>18&7 | 240))
		out = append(out, I(char, 12))
		out = append(out, I(char, 6))
	}
	out = append(out, []byte(fromCharCode(char&63|128))...)
	return out
}

func I(char, num int) byte {
	return byte(char>>num&63 | 128)
}

func J(char int) {
	if char >= 55296 && char <= 57343 {
		panic("not a scalar value")
	}
}

func K(payload []byte) []byte {
	out := []byte{}
	bytesAdded := 0
	payloadLen := len(payload)
	byteValue := 0
	nextByte := 0
	for bytesAdded < payloadLen {
		byteValue = int(payload[bytesAdded])
		bytesAdded++
		if byteValue >= 55296 && byteValue <= 56319 && bytesAdded < payloadLen {
			nextByte = int(payload[bytesAdded])
			bytesAdded++
			bytesAdded++
			if (nextByte & 64512) == 56320 {
				out = append(out, byte(((byteValue&1023)<<10)+(nextByte&1023)+65536))
			} else {
				out = append(out, byte(byteValue))
				bytesAdded--
			}
		} else {
			out = append(out, byte(byteValue))
		}
	}
	return out
}

// =======================================
// ==          XOR key stuff            ==
// =======================================

func encodeXORKey(key string) string {
	output := []string{}
	chars := strings.Split("abcdhijkxy", "")
	for i := 0; i < 4; i++ {
		charPos := 0
		keyPos := 1
		if i == 3 {
			keyPos = 1 + len(key)%4
		}
		for x := 0; x < keyPos; x++ {
			charCodePos := i + x
			if x < len(key) {
				charPos = charPos + charCodeAt(key, charCodePos)
			}
		}
		charPos = charPos * 31
		output = append(output, chars[charPos%len(chars)])
	}
	return strings.Join(output, "")
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
