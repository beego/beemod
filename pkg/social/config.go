package social

// example
/**
[beemod.oauth2.qq]
	debug = true
	mode  = "qq"
	app_id  = "app_id"
	app_secret  = "app_secret"
	redirectURI = "www.beego.com"
[beemod.oauth2.wx]
	debug = true
	mode  = "wx"
	app_id  = "app_id"
	app_secret  = "app_secret"
	redirectURI = "www.beego.com"
[beemod.oauth2.github]
	debug = true
	mode  = "github"
	app_id  = "app_id"
	app_secret  = "app_secret"
	redirectURI = "www.beego.com"
*/
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
