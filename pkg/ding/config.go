package ding

// example
/**
[beemod.oss.myoss]
	debug = true
	mode  = "file"
*/
type InvokerCfg struct {
	Debug           bool
	Mode            string
	WebhookUrl      string
}

var DefaultInvokerCfg = InvokerCfg{
	Debug:           false,
	Mode:            "keyword",
	WebhookUrl:       "",
}
