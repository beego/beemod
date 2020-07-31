// Author: SDing <deen.job@qq.com>
// Date: 2020/6/28 - 1:06 PM

package main

import (
	"github.com/beego/beemod"
	"github.com/beego/beemod/pkg/logger"
)

// custom config toml template
var config = `
	type = "file"
    level = 7
    path = "log.log"
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
