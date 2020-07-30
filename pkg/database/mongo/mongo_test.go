/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/7/30 11:14
 */
package mongo

import (
	"github.com/beego/beemod"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

const configTpl = `
[beego.mongo.dev]
    URL   = "127.0.0.1:27017"
	debug = true
    source = "mydb_test"
    user   = "root"
    password   = "root"
`
const (
	db         = "Movies"
	collection = "MovieModel"
)

func TestMysqlConfig(t *testing.T) {
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

func TestMysqlInit(t *testing.T) {
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

func TestMysqlInstance(t *testing.T) {
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
					err = obj.Insert(db, collection, "testData")
					c.So(err, c.ShouldBeNil)
				})
			})
		})
	})
}
