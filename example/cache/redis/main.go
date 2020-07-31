package main

import (
	"github.com/beego/beemod"
	"github.com/beego/beemod/pkg/cache/redis"
)

var config = `
	[beego.redis.dev]
		addr = "127.0.0.1:6379"
		password = ""
`

func main() {
	err := beemod.Register(redis.DefaultBuild).SetCfg([]byte(config), "toml").Run()

	if err != nil {
		panic("register err:" + err.Error())
	}
	obj := redis.Invoker("dev")

	_, _ = obj.Set("a", "a", 60)
}
