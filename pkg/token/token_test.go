package token

import (
	"github.com/beego/beemod"
	"github.com/beego/beemod/pkg/token/standard"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

const configTpl = `
[beego.token.dev]
	mode = "mysql"
	redisTokenKeyPattern = "/egoshop/token/%d"
	accessTokenExpireInterval = 604800
	accessTokenIss           = "github.com/goecology/egoshop"
	accessTokenKey           = "ecologysK#xo"
[beego.token.dev.mysql]
    debug = true
    level = "panic"
    network = "tcp"
    dialect = "mysql"
    addr = "127.0.0.1:3306"
    username = "root"
    password = ""
    db = "beetest"
    charset = "utf8"
    parseTime = "True"
    loc = "Local"
    maxOpenConns = 30
    maxIdleConns = 10
    connMaxLifetime = "300s"
	aliasName="default"
[beego.token.dev.Logger]
    level = 7
    path = "token.log"
`

func TestTokenConfig(t *testing.T) {
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

func TestTokenInit(t *testing.T) {
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

func TestTokenInstance(t *testing.T) {
	var (
		err         error
		obj         *Client
		config      string
		accessToken standard.AccessTokenTicket
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
					accessToken, err = obj.CreateAccessToken(123123, time.Now().Unix())
					c.So(err, c.ShouldBeNil)
					t.Log(accessToken)
					t.Log(obj.CheckAccessToken(accessToken.AccessToken))
					t.Log(obj.DecodeAccessToken(accessToken.AccessToken))
				})
			})
		})
	})
}
