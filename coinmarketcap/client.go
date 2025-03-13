package coinmarketcap

import (
	tlsclient "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
)

type Client struct {
	HTTPClient tlsclient.HttpClient
	Proxy      string
}

func MakeClient(proxy string) *Client {
	opts := []tlsclient.HttpClientOption{
		tlsclient.WithTimeoutSeconds(10),
		tlsclient.WithProxyUrl(proxy),
		tlsclient.WithInsecureSkipVerify(),
		tlsclient.WithClientProfile(profiles.Okhttp4Android12),
	}
	h, err := tlsclient.NewHttpClient(nil, opts...)
	if err != nil {
		panic(err)
	}
	c := &Client{
		HTTPClient: h,
		Proxy:      proxy,
	}
	return c
}
