// Author: SDing <deen.job@qq.com>
// Date: 2020/7/8 - 5:46 PM

package main

import (
	"fmt"
	"github.com/beego-dev/beemod"
	"github.com/beego-dev/beemod/pkg/qr_code"
)

var qrConfig = `
	[beego.qr_code.dev]
		debug = false
		mode = "qr_code"
		avatarX = 40
		avatarY = 40
		size = 280
		foreground = "/Users/ding/Desktop/qrcode/example/static/f.jpg"
`

func main() {
	// use example:  https://www.yuque.com/beego.dev/iuzxi0/zki6gh
	err := beemod.Register(
		qr_code.DefaultBuild,
	).SetCfg([]byte(qrConfig), "toml").Run()
	if err != nil {
		panic("register err:" + err.Error())
	}
	qrc := qr_code.Invoker("dev")
	err = qrc.Of.New("Hello BeeGo", "./qr.png")
	err = qrc.Of.NewFg("Hello BeeGo", "./qr_fg.png")
	err = qrc.Of.NewAvatar("Hello BeeGo", "./fg.png", "./qr_avatar.png", 60, true)
	if err != nil {
		fmt.Println(err)
	}
}
