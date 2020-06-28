package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/beego-dev/beemod"
	"github.com/beego-dev/beemod/pkg/cache"
	"time"
)

var config = `
	[beego.cache.redis]
		key = "default"
		conn = ":6379"
		password = ""
`

func main() {
	err := beemod.Register(
		cache.DefaultBuild).SetCfg([]byte(config), "toml").Run()

	if err != nil {
		panic("register err:" + err.Error())
	}
	obj := cache.Invoker("redis")

	// put
	isPut := obj.Put("beemod", 1, time.Second*100)
	logs.Info(isPut)

	isPut = obj.Put("hello", "world", time.Second*100)
	logs.Info(isPut)

	// get
	result := obj.Get("beemod")
	logs.Info(string(result.([]byte)))

	multiResult := obj.GetMulti([]string{"beemod", "hello"})
	for i := range multiResult {
		logs.Info(string(multiResult[i].([]byte)))
	}

	// isExist
	isExist := obj.IsExist("beemod")
	logs.Info(isExist)
}
