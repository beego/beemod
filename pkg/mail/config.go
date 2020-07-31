// Author: SDing <deen.job@qq.com>
// Date: 2020/6/28 - 12:08

package mail

type InvokerCfg struct {
	Debug        bool `ini:"debug"`
	Mode         string `ini:"mode"`
	Host         string `ini:"host"`
	Port         int `ini:"port"`
	FromEmail    string `ini:"fromEmail"`
	FromPassword string `ini:"fromPassword"`
}

var DefaultInvokerCfg = InvokerCfg{
	Debug:        false,
	Mode:         "smtp",
	Host:         "smtp.qq.com",
	Port:         25,
	FromEmail:    "",
	FromPassword: "",
}
