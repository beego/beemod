// Author: SDing <deen.job@qq.com>
// Date: 2020-06-27 20:02:08
// Notification system module Interface design v1.2 Design: SDing
// EMail function implementation meets current standards Function implementation: SDing
// Contact me with a bug~

package mail

import (
	"github.com/astaxie/beego/logs"
	"gopkg.in/gomail.v2"
	"gopkg.in/ini.v1"
)

var dia *gomail.Dialer
var Config *cfg

func init() {
	cfgOpt, err := ini.Load("/Users/ding/Desktop/beemod/pkg/notice/mail/config.ini")
	if err != nil {
		panic("loading mail config field :" + err.Error())
	}
	//err = cfgOpt.MapTo(Config)
	logs.Info(cfgOpt.Section("mail").Key("Host"))
	logs.Info(cfgOpt.Section("mail").Key("Port"))
	logs.Info(cfgOpt.Section("mail").Key("FromEmail"))
	logs.Info(cfgOpt.Section("mail").Key("FromPassword"))
	Config = &cfg{
		Host:         cfgOpt.Section("mail").Key("Host").String(),
		Port:         25,
		FromEmail:    cfgOpt.Section("mail").Key("FromEmail").String(),
		FromPassword: cfgOpt.Section("mail").Key("FromPassword").String(),
	}
	if err != nil {
		logs.Error(err)
	}
	dia = gomail.NewDialer(Config.Host, Config.Port, Config.FromEmail, Config.FromPassword)
}

// config Configuration parameter
type cfg struct {
	Host         string `ini:"Host"`
	Port         int    `ini:"Port"`
	FromEmail    string `ini:"FromEmail"`
	FromPassword string `ini:"FromPassword"`
}

// Dialer Mailbox server dialer
type Dialer struct {
	Dialer *gomail.Dialer
	Msg    *gomail.Message
}

// New Create a server dialer
func New(msg *gomail.Message) *Dialer {
	return &Dialer{
		Dialer: dia,
		Msg:    msg,
	}
}

func (D *Dialer) Parse() func() {
	return func() {
		err := D.Dialer.DialAndSend(D.Msg)
		if err != nil {
			logs.Error(err)
		}
	}
}
