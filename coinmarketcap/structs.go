package coinmarketcap

type Captcha struct {
	Code string `json:"code"`
	Data struct {
		Sig         string `json:"sig"`
		Salt        string `json:"salt"`
		Path2       string `json:"path2"`
		Ek          string `json:"ek"`
		CaptchaType string `json:"captchaType"`
		Tag         string `json:"tag"`
		Fb          string `json:"fb"`
		I18N        string `json:"i18n"`
	} `json:"data"`
	Success bool `json:"success"`
}

type Payload struct {
	Ev struct {
		Wd     int    `json:"wd"`
		Im     int    `json:"im"`
		De     string `json:"de"`
		Prde   string `json:"prde"`
		Brla   int    `json:"brla"`
		Pl     string `json:"pl"`
		Wiinhe int    `json:"wiinhe"`
		Wiouhe string `json:"wiouhe"`
	} `json:"ev"`
	Be struct {
		Ec struct {
			Ts int `json:"ts"`
			Tm int `json:"tm"`
			Te int `json:"te"`
		} `json:"ec"`
		El []string `json:"el"`
		Th struct {
			El []string `json:"el"`
			Si struct {
				W int `json:"w"`
				H int `json:"h"`
			} `json:"si"`
		} `json:"th"`
	} `json:"be"`
	Dist       int    `json:"dist"`
	ImageWidth string `json:"imageWidth"`
}

type SolveResponse struct {
	Code string `json:"code"`
	Data struct {
		Result int    `json:"result"`
		Tag    string `json:"tag"`
		I18N   string `json:"i18n"`
		Token  string `json:"token"`
	} `json:"data"`
	Success bool `json:"success"`
}
