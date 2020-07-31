package social

type InvokerCfg struct {
	Debug       bool `ini:"debug"`
	Mode        string `ini:"mode"`
	AppID       string `ini:"appId"`
	AppSecret   string `ini:"appSecret"`
	RedirectURI string `ini:"redirectUri"`
}

var DefaultInvokerCfg = InvokerCfg{
	Debug:     false,
	Mode:      "",
	AppID:     "",
	AppSecret: "",
}
