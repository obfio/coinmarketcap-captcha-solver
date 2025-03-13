package coinmarketcap

func (p *Payload) GenSValue(encodedPayload, sig, salt string) int {
	str := "CMC_register" + sig + encodedPayload + salt
	out := 0
	for i := 0; i < len(str); i++ {
		out += charCodeAt(str, i)
	}
	return out
}
