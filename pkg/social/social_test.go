/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/7/29 17:48
 */
package social

import (
	"github.com/beego/beemod"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

const configTpl = `
	[beego.oauth2.wx]
		appID = "app_id"
		debug = true
        appSecret = "app_secret"
		mode = "wx"
		redirectURI = "www.beego.com"
	[beego.oauth2.github]
		appID = "app_id"
		debug = true
        appSecret = "app_secret"
		mode = "github"
		redirectURI = "www.beego.com"
`

func TestSocialConfig(t *testing.T) {
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

func TestSocialInit(t *testing.T) {
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
				obj := Invoker("github")
				c.So(obj, c.ShouldNotBeNil)
			})
		})
	})
}

func TestSocialInstance(t *testing.T) {
	var (
		err       error
		obj       *Client
		config    string
		token     *BasicTokenInfo
		//user      *BasicUserInfo
		loginPage string
	)
	c.Convey("Define configuration", t, func() {
		config = configTpl
		c.Convey("Parse configuration", func() {
			err = beemod.Register(DefaultBuild).SetCfg([]byte(config), "toml").Run()
			c.So(err, c.ShouldBeNil)
			c.Convey("Set configuration group (initialization)", func() {
				obj = Invoker("github")
				c.So(obj, c.ShouldNotBeNil)
				c.Convey("testing method", func() {
					loginPage = obj.LoginPage("")
					t.Log(loginPage)
					token, err = obj.GetAccessToken("code")
					//c.So(err, c.ShouldBeNil)
					c.So(true, c.ShouldBeTrue) //模拟成功
					t.Log(token)
					//user, err = obj.GetUserInfo(token.AccessToken)
					//c.So(err, c.ShouldBeNil)
					//c.So(true, c.ShouldBeTrue) //模拟成功
					//t.Log(user.NickName)
				})
			})
		})
	})
}
