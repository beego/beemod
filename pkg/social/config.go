package social

// example
/**
[beemod.oauth2.qq]
	debug = true
	mode  = "qq"
	app_id  = "app_id"
	app_secret  = "app_secret"
[beemod.oauth2.wx]
	debug = true
	mode  = "wx"
	app_id  = "app_id"
	app_secret  = "app_secret"
[beemod.oauth2.github]
	debug = true
	mode  = "github"
	app_id  = "app_id"
	app_secret  = "app_secret"
*/
type InvokerCfg struct {
	Debug     bool
	Mode      string
	AppID     string
	AppSecret string
}

var DefaultInvokerCfg = InvokerCfg{
	Debug:     false,
	Mode:      "",
	AppID:     "",
	AppSecret: "",
}
