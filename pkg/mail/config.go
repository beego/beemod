// Author: SDing <deen.job@qq.com>
// Date: 2020/6/28 - 12:08

package mail

// example
/**
[beemod.mail.my]
	debug = true
	mode  = "smtp"
*/
type InvokerCfg struct {
	Debug        bool
	Mode         string
	Host         string
	Port         int
	FromEmail    string
	FromPassword string
}

var DefaultInvokerCfg = InvokerCfg{
	Debug:        false,
	Mode:         "smtp",
	Host:         "smtp.qq.com",
	Port:         25,
	FromEmail:    "1423119397@qq.com",
	FromPassword: "",
}
