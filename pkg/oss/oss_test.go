package oss

import (
	"github.com/beego/beemod"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

const configTpl = `
	[beego.oss.myoss]
		mode = "file"
		debug = true
        isDeleteSrcPath = false
        cdnName = "http://127.0.0.1:8080/oss/"
        fileBucket = "oss"
`

func TestOssConfig(t *testing.T) {
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

func TestOssInit(t *testing.T) {
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
				obj := Invoker("myoss")
				c.So(obj, c.ShouldNotBeNil)
			})
		})
	})
}

func TestOssInstance(t *testing.T) {
	var (
		err     error
		obj     *Client
		config  string
		dstPath string
	)
	c.Convey("Define configuration", t, func() {
		config = configTpl
		c.Convey("Parse configuration", func() {
			err = beemod.Register(DefaultBuild).SetCfg([]byte(config), "toml").Run()
			c.So(err, c.ShouldBeNil)
			c.Convey("Set configuration group (initialization)", func() {
				obj = Invoker("myoss")
				c.So(obj, c.ShouldNotBeNil)
				c.Convey("testing method", func() {
					dstPath = obj.GenerateKey("mock")
					err = obj.PutObjectFromFile(dstPath, "../../example/oss/image/oss.jpg")
					c.So(err, c.ShouldBeNil)
				})
			})
		})
	})
}
