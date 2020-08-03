package oauth

type InvokerCfg struct {
	Debug       bool   `ini:"debug"`
	AppID       string `ini:"appId"`
	AppSecret   string `ini:"appSecret"`
	RedirectURI string `ini:"redirectUri"`
	AuthURL     string `ini:"authUrl"`
	TokenURL    string `ini:"tokenUrl"`
	State       string `ini:"state"`
	UserInfoURL string `ini:"userInfoUrl"`
}

var DefaultInvokerCfg = InvokerCfg{
	Debug:       false,
	AppID:       "",
	AppSecret:   "",
	RedirectURI: "",
	AuthURL:     "",
	TokenURL:    "",
	State:       "",
}
