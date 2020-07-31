package ding

type InvokerCfg struct {
	Debug      bool `ini:"debug"`
	Mode       string `ini:"mode"`
	WebhookUrl string `ini:"webhookUrl"`
}

var DefaultInvokerCfg = InvokerCfg{
	Debug:      false,
	Mode:       "keyword",
	WebhookUrl: "",
}
