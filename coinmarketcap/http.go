package coinmarketcap

import (
	"encoding/json"
	"errors"
	"fmt"
	http "github.com/bogdanfinn/fhttp"
	"io"
	"net/url"
	"strings"
)

func (c *Client) GetCaptcha() (*Captcha, error) {

	req, err := http.NewRequest("POST", "https://api.commonservice.io/gateway-api/v1/public/antibot/getCaptcha", strings.NewReader("bizId=CMC_register&sv=20220812&snv=1.4.7&lang=en&securityCheckResponseValidateId=CMC_register&clientType=android"))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Fvideo-Id", "xxx")
	req.Header.Set("User-Agent", "Binance/1.0.0 Mozilla/5.0 (Linux; Android 12; SM-A528B Build/V417IR; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/91.0.4472.114 Safari/537.36 Captcha/1.0.0 (Android 1.0.0) SecVersion/1.4.7")
	req.Header.Set("Content-Type", "text/plain; charset=UTF-8")
	req.Header.Set("Captcha-Sdk-Version", "1.0.0")
	req.Header.Set("Device-Info", "eyJzY3JlZW5fcmVzb2x1dGlvbiI6IjYwMCwxMDY3IiwiYXZhaWxhYmxlX3NjcmVlbl9yZXNvbHV0aW9uIjoiNjAwLDEwNjciLCJzeXN0ZW1fdmVyc2lvbiI6InVua25vd24iLCJicmFuZF9tb2RlbCI6InVua25vd24iLCJ0aW1lem9uZSI6IkFtZXJpY2EvTmV3X1lvcmsiLCJ0aW1lem9uZU9mZnNldCI6MjQwLCJ1c2VyX2FnZW50IjoiQmluYW5jZS8xLjAuMCBNb3ppbGxhLzUuMCAoTGludXg7IEFuZHJvaWQgMTI7IFNNLUE1MjhCIEJ1aWxkL1Y0MTdJUjsgd3YpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIFZlcnNpb24vNC4wIENocm9tZS85MS4wLjQ0NzIuMTE0IFNhZmFyaS81MzcuMzYgQ2FwdGNoYS8xLjAuMCAoQW5kcm9pZCAxLjAuMCkgU2VjVmVyc2lvbi8xLjQuNyIsImxpc3RfcGx1Z2luIjoiIiwicGxhdGZvcm0iOiJMaW51eCBpNjg2Iiwid2ViZ2xfdmVuZG9yIjoidW5rbm93biIsIndlYmdsX3JlbmRlcmVyIjoidW5rbm93biJ9")
	req.Header.Set("Bnc-Uuid", "xxx")
	req.Header.Set("Clienttype", "android")
	req.Header.Set("X-Captcha-Se", "true")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", "https://staticrecap.cgicgi.io")
	req.Header.Set("X-Requested-With", "com.coinmarketcap.android")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://static://staticrecap.cgicgi.io/")
	req.Header.Set("Accept-Language", "en,en-US;q=0.9")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	f := &Captcha{}
	err = json.Unmarshal(b, f)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (c *Client) GetImage(path string) ([]byte, error) {
	req, err := http.NewRequest("GET", "https://staticrecap.cgicgi.io"+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Binance/1.0.0 Mozilla/5.0 (Linux; Android 12; SM-A528B Build/V417IR; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/91.0.4472.114 Safari/537.36 Captcha/1.0.0 (Android 1.0.0) SecVersion/1.4.7")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", "https://staticrecap.cgicgi.io")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (c *Client) SolveCaptcha(sig, encodedPayload string, sValue int) (*SolveResponse, error) {
	//bizId=CMC_register&sv=20220812&snv=1.4.7&lang=en&securityCheckResponseValidateId=CMC_register&clientType=android&sig=pUBDZytKUt6D1SSxq7H8JQEtYSiU4ywACNPeOnJT1DwXKt1c&data
	t := url.Values{}
	t.Set("data", encodedPayload)
	body := fmt.Sprintf(`bizId=CMC_register&sv=20220812&snv=1.4.7&lang=en&securityCheckResponseValidateId=CMC_register&clientType=android&sig=%s&data=%s&s=%v`, sig, strings.Split(t.Encode(), "=")[1], sValue)
	req, err := http.NewRequest("POST", "https://api.commonservice.io/gateway-api/v1/public/antibot/validateCaptcha", strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Fvideo-Id", "xxx")
	req.Header.Set("User-Agent", "Binance/1.0.0 Mozilla/5.0 (Linux; Android 12; SM-A528B Build/V417IR; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/91.0.4472.114 Safari/537.36 Captcha/1.0.0 (Android 1.0.0) SecVersion/1.4.7")
	req.Header.Set("Content-Type", "text/plain; charset=UTF-8")
	req.Header.Set("Captcha-Sdk-Version", "1.0.0")
	req.Header.Set("Device-Info", "eyJzY3JlZW5fcmVzb2x1dGlvbiI6IjYwMCwxMDY3IiwiYXZhaWxhYmxlX3NjcmVlbl9yZXNvbHV0aW9uIjoiNjAwLDEwNjciLCJzeXN0ZW1fdmVyc2lvbiI6InVua25vd24iLCJicmFuZF9tb2RlbCI6InVua25vd24iLCJ0aW1lem9uZSI6IkFtZXJpY2EvTmV3X1lvcmsiLCJ0aW1lem9uZU9mZnNldCI6MjQwLCJ1c2VyX2FnZW50IjoiQmluYW5jZS8xLjAuMCBNb3ppbGxhLzUuMCAoTGludXg7IEFuZHJvaWQgMTI7IFNNLUE1MjhCIEJ1aWxkL1Y0MTdJUjsgd3YpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIFZlcnNpb24vNC4wIENocm9tZS85MS4wLjQ0NzIuMTE0IFNhZmFyaS81MzcuMzYgQ2FwdGNoYS8xLjAuMCAoQW5kcm9pZCAxLjAuMCkgU2VjVmVyc2lvbi8xLjQuNyIsImxpc3RfcGx1Z2luIjoiIiwicGxhdGZvcm0iOiJMaW51eCBpNjg2Iiwid2ViZ2xfdmVuZG9yIjoidW5rbm93biIsIndlYmdsX3JlbmRlcmVyIjoidW5rbm93biJ9")
	req.Header.Set("Bnc-Uuid", "xxx")
	req.Header.Set("Clienttype", "android")
	req.Header.Set("X-Captcha-Se", "true")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", "https://staticrecap.cgicgi.io")
	req.Header.Set("X-Requested-With", "com.coinmarketcap.android")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://static://staticrecap.cgicgi.io/")
	req.Header.Set("Accept-Language", "en,en-US;q=0.9")
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("invalid status code")
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	f := &SolveResponse{}
	err = json.Unmarshal(b, &f)
	if err != nil {
		return nil, err
	}
	if !f.Success || f.Code != "000000" || f.Data.Token == "" {
		return nil, errors.New(string(b))
	}
	return f, nil
}
