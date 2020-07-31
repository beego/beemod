/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/7/29 17:12
 */
package qr_code

import (
	"github.com/beego/beemod"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

const configTpl = `
	[beego.qrcode.dev]
		debug = false
		mode = "qr_code"
		avatarX = 40
		avatarY = 40
		size = 280
		foreground = "./gopher-500.png"
`

func TestQrCodeConfig(t *testing.T) {
	var (
		err    error
		config string
	)
	c.Convey("Define configuration", t, func() {
		config = configTpl
		c.Convey("Parse configuration", func() {
			err = beemod.Register(DefaultBuild).SetCfg([]byte(config), "toml").Run()
			c.So(err, c.ShouldBeNil)
		})
	})
}

func TestQrCodeInit(t *testing.T) {
	var (
		err    error
		config string
	)
	c.Convey("Define configuration", t, func() {
		config = configTpl
		c.Convey("Define configuration", func() {
			err = beemod.Register(DefaultBuild).SetCfg([]byte(config), "toml").Run()
			c.So(err, c.ShouldBeNil)
			c.Convey("Set configuration group (initialization)", func() {
				obj := Invoker("dev")
				c.So(obj, c.ShouldNotBeNil)
			})
		})
	})
}

func TestQrCodeInstance(t *testing.T) {
	var (
		err    error
		obj    *Client
		config string
	)
	c.Convey("Define configuration", t, func() {
		config = configTpl
		c.Convey("Parse configuration", func() {
			err = beemod.Register(DefaultBuild).SetCfg([]byte(config), "toml").Run()
			c.So(err, c.ShouldBeNil)
			c.Convey("Set configuration group (initialization)", func() {
				obj = Invoker("dev")
				c.So(obj, c.ShouldNotBeNil)
				c.Convey("testing method", func() {
					err = obj.Of.New("Hello BeeGo", "./qr.png")
					//err = obj.Of.NewFg("Hello BeeGo", "./qr_fg.png")
					//err = obj.Of.NewAvatar("Hello BeeGo", "./fg.png", "./qr_avatar.png", 60, true)
					c.So(err, c.ShouldBeNil)
				})
			})
		})
	})
}
