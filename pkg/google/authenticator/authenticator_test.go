/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/8/6 18:51
 */
package authenticator

import (
	"fmt"
	"github.com/beego/beemod"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

const configTpl = `
  [beego.authenticator.dev]
	user = "admin123"
	iss = "admin@beego.com"
`

func TestAuthenticatorConfig(t *testing.T) {
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

func TestAuthenticatorInit(t *testing.T) {
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

func TestAuthenticatorInstance(t *testing.T) {
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
					key := obj.Auth.GenerateKey()
					fmt.Println(key)
					c.So(key, c.ShouldNotBeNil)
					uri := obj.Auth.ProvisionURI()
					fmt.Println(uri)
					c.So(uri, c.ShouldNotBeNil)
					ok, err := obj.Auth.Authenticate("123456")
					fmt.Println(ok)
					c.So(ok, c.ShouldNotBeNil)
					c.So(err, c.ShouldBeNil)
				})
			})
		})
	})
}
