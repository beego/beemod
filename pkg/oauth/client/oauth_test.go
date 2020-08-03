/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/8/3 10:00
 */
package client

import (
	"github.com/beego/beemod"
	c "github.com/smartystreets/goconvey/convey"
	"log"
	"testing"
)

const configTpl = `
[beego.oauth2.dev]
	state = "beego"
	appID = "sfdhjaksfddsajks"
	appSecret = "mnjUYj8rlXXKGS2RNgsdad7lygWrjJzjD5"
	authURL = "http://oauthadmin.yitum.com/user/login"
	tokenURL = "http://oauthadmin.yitum.com/api/v1/oauth/token"
	redirectURI = "http://localhost:8000/api/code"
	userInfoURL = "http://oauthadmin.yitum.com/api/v1/oauth/user"
	Scopes=[]
`

func TestOauthConfig(t *testing.T) {
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

func TestOauthInit(t *testing.T) {
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

func TestOauthInstance(t *testing.T) {
	var (
		err    error
		obj    OAuthService
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
					page := obj.LoginPage()
					log.Print(page)
					c.So(err, c.ShouldBeNil)
				})
			})
		})
	})
}
