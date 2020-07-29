// TODO: logger-test-example
// Author: SDing <deen.job@qq.com>
// Date: 2020/6/28 - 1:06 PM

package main

import (
	"github.com/beego/beemod"
	"github.com/beego/beemod/pkg/logger"
)

// custom config toml template
var config = `
  [beego.logger.dev]
    level = 7
    path = "token.log"
`

func main() {
	err := beemod.Register(
		logger.DefaultBuild,
	).SetCfg([]byte(config), "toml").Run()
	if err != nil {
		panic("register err:" + err.Error())
	}
	client := logger.Invoker("dev")
	client.BeeLogger.Error("错误")
}
