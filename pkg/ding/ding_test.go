package ding

import (
	"github.com/beego/beemod"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

const configTpl = `
  [beego.ding.myding]
	mode = "file"
	debug = true
  	WebhookUrl = "https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxxxxxxxxxxxxxx"
`

func TestDingConfig(t *testing.T) {
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

func TestDingInit(t *testing.T) {
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
				obj := Invoker("myding")
				c.So(obj, c.ShouldNotBeNil)
			})
		})
	})
}

func TestDingInstance(t *testing.T) {
	var (
		err    error
		obj    *Client
		config string
		res    string
	)
	c.Convey("Define configuration", t, func() {
		config = configTpl
		c.Convey("Parse configuration", func() {
			err = beemod.Register(DefaultBuild).SetCfg([]byte(config), "toml").Run()
			c.So(err, c.ShouldBeNil)
			c.Convey("Set configuration group (initialization)", func() {
				obj = Invoker("myding")
				c.So(obj, c.ShouldNotBeNil)
				c.Convey("testing method", func() {
					res, err = obj.SendMsg("TESTa")
					c.So(err, c.ShouldBeNil)
					t.Log(res)
				})
			})
		})
	})
}
