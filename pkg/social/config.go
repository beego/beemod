package social

type InvokerCfg struct {
	Debug       bool
	Mode        string
	AppID       string
	AppSecret   string
	RedirectURI string
}

var DefaultInvokerCfg = InvokerCfg{
	Debug:     false,
	Mode:      "",
	AppID:     "",
	AppSecret: "",
}
