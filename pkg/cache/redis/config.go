package redis

import (
	"github.com/beego/beemod/pkg/module"
)

type CallerCfg struct {
	Debug bool

	Network        string // tcp
	Addr           string // 127.0.0.1:6379
	DB             int
	Password       string
	ConnectTimeout module.Duration
	ReadTimeout    module.Duration
	WriteTimeout   module.Duration

	MaxIdle     int
	MaxActive   int
	IdleTimeout module.Duration
	Wait        bool
}
