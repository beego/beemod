package session

import "github.com/astaxie/beego/session"

type InvokerCfg struct {
	Debug     bool
	Mode      string
	MangerCfg *session.ManagerConfig
}

var DefaultInvokerCfg = InvokerCfg{
	Debug:     false,
	Mode:      "memory",
	MangerCfg: DefaultManagerConfig,
}

var DefaultManagerConfig = &session.ManagerConfig{
	CookieName:      "gosessionid",
	Gclifetime:      10,
	EnableSetCookie: true,
}
