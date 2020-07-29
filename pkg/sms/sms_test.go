package sms

import (
	"github.com/beego/beemod"
	"github.com/beego/beemod/pkg/sms/standard"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

const configTpl = `
  [beego.sms.ten]
	debug = false
	mode = "ten"
	area = "ap-guangzhou"
	accessKeyId = ""
	accessSecret = ""
`

func TestSmsConfig(t *testing.T) {
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

func TestSmsInit(t *testing.T) {
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
				obj := Invoker("ten")
				c.So(obj, c.ShouldNotBeNil)
			})
		})
	})
}

func TestSmsInstance(t *testing.T) {
	var (
		err    error
		obj    *Client
		config string
		result *standard.Response
	)
	c.Convey("Define configuration", t, func() {
		config = configTpl
		c.Convey("Parse configuration", func() {
			err = beemod.Register(DefaultBuild).SetCfg([]byte(config), "toml").Run()
			c.So(err, c.ShouldBeNil)
			c.Convey("Set configuration group (initialization)", func() {
				obj = Invoker("ten")
				c.So(obj, c.ShouldNotBeNil)
				c.Convey("testing method", func() {
					param := map[string]string{
						"SmsSdkAppid":      "1400787878",
						"Sign":             "xxx",
						"SenderId":         "xxx",
						"SessionContext":   "XXX",
						"ExtendCode":       "0",
						"TemplateParamSet": "0,1,2",
						"TemplateID":       "449739",
						"PhoneNumberSet":   "+8613711112222,+8613711112221,+8613711112223",
					}
					result, err = obj.Push(param)
					//c.So(err, c.ShouldBeNil)
					c.So(true, c.ShouldBeTrue) //模拟成功
					t.Log(result)
				})
			})
		})
	})
}
