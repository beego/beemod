/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/7/29 15:39
*/
package cache

import (
	"github.com/beego/beemod"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

const configTpl = `
	[beego.cache.redis]
		key = "default"
		conn = "127.0.0.1:6379"
		password = ""
`

func TestCacheConfig(t *testing.T) {
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

func TestCacheInit(t *testing.T) {
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
				obj := Invoker("redis")
				c.So(obj, c.ShouldNotBeNil)
			})
		})
	})
}

func TestCacheInstance(t *testing.T) {
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
				obj = Invoker("redis")
				c.So(obj, c.ShouldNotBeNil)
				c.Convey("testing method", func() {
					err = obj.Put("testKey1", 1, 10*time.Second)
					c.So(err, c.ShouldBeNil)
				})
			})
		})
	})
}
