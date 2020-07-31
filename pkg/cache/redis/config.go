package redis

import (
	"github.com/beego/beemod/pkg/module"
)

type CallerCfg struct {
	Debug bool

	Network        string `ini:"network"` // tcp
	Addr           string `ini:"addr"` // 127.0.0.1:6379
	DB             int `ini:"db"`
	Password       string `ini:"password"`
	ConnectTimeout module.Duration `ini:"connectTimeout"`
	ReadTimeout    module.Duration `ini:"readTimeout"`
	WriteTimeout   module.Duration `ini:"writeTimeout"`

	MaxIdle     int `ini:"maxIdle"`
	MaxActive   int `ini:"maxActive"`
	IdleTimeout module.Duration `ini:"idleTimeout"`
	Wait        bool `ini:"wait"`
}
