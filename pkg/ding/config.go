package ding

type InvokerCfg struct {
	Debug      bool
	Mode       string
	WebhookUrl string
}

var DefaultInvokerCfg = InvokerCfg{
	Debug:      false,
	Mode:       "keyword",
	WebhookUrl: "",
}
