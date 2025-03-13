package main

import (
	"bytes"
	"fmt"
	"github.com/obfio/cmc-solve-image/coinmarketcap"
	"github.com/obfio/cmc-solve-image/solve"
	"image/png"
	"os"
	"strings"
)

func main() {
	client := coinmarketcap.MakeClient("")
	captcha, err := client.GetCaptcha()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", captcha)
	if captcha.Data.CaptchaType != "SLIDE" {
		panic("captcha type is " + captcha.Data.CaptchaType)
	}
	imageBytes, err := client.GetImage(captcha.Data.Path2)
	if err != nil {
		panic(err)
	}
	os.WriteFile(fmt.Sprintf("./examples/%s", strings.Split(captcha.Data.Path2, "/")[7]), imageBytes, 0666)

	// convert bytes to png
	img, err := png.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		panic(err)
	}

	// solve image
	solvedX := solve.SolveImage(img)

	// make the payload
	payload := &coinmarketcap.Payload{}
	payload.FillPayload(solvedX)
	fmt.Printf("%+v\n", payload)

	// get encoded payload
	encodedPayload := payload.Encode(captcha.Data.Ek)

	// get S value
	sValue := payload.GenSValue(encodedPayload, captcha.Data.Sig, captcha.Data.Salt)
	solvedCaptcha, err := client.SolveCaptcha(captcha.Data.Sig, encodedPayload, sValue)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", solvedCaptcha)
	fmt.Printf("\n\n%s\n\n", solvedCaptcha.Data.Token)
}
