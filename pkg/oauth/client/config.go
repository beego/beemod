package client

type InvokerCfg struct {
	Debug       bool     `ini:"debug"`
	AppID       string   `ini:"appId"`
	AppSecret   string   `ini:"appSecret"`
	RedirectURI string   `ini:"redirectUri"`
	AuthURL     string   `ini:"authUrl"`
	TokenURL    string   `ini:"tokenUrl"`
	State       string   `ini:"state"`
	UserInfoURL string   `ini:"userInfoUrl"`
	Scopes      []string `ini:"scopes"`
}

var DefaultInvokerCfg = InvokerCfg{
	Debug:       false,
	AppID:       "",
	AppSecret:   "",
	RedirectURI: "",
	AuthURL:     "",
	TokenURL:    "",
	State:       "",
	UserInfoURL: "",
	Scopes:      nil,
}
