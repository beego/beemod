package mysql

import (
	"github.com/astaxie/beego/orm"
	"github.com/beego/beemod"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

type User struct {
	Id   int    `orm:"pk;auto"`
	Name string `orm:""`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}

const configTpl = `
[beego.mysql.dev]
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
`

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
					//创建表
					err = orm.RunSyncdb(obj.Cfg.AliasName, false, false)
					user := new(User)
					user.Name = "testName"
					_, err = obj.O.Insert(user)
					c.So(err, c.ShouldBeNil)
				})
			})
		})
	})
}
