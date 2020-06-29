package main

import (
	"github.com/beego-dev/beemod"
	"github.com/beego-dev/beemod/pkg/ding"
  "fmt"
)

var config = `
  [beego.ding.myding]
	mode = "file"
	debug = true
  WebhookUrl = "https://oapi.dingtalk.com/robot/send?access_token="
`

func main() {
	err := beemod.Register(
		ding.DefaultBuild,
	).SetCfg([]byte(config), "toml").Run()
	if err != nil {
		panic("register err:" + err.Error())
	}
  httpClient := ding.Invoker("myding")
  res, _ := httpClient.SendMsg("TESTa")
  fmt.Println(res)
}
