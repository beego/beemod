/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/7/31 12:00
 */
package rabbitmq

import (
	"github.com/beego/beemod"
	c "github.com/smartystreets/goconvey/convey"
	"github.com/streadway/amqp"
	"testing"
)

const configTpl = `
  [beego.rabbitmq.dev]
	host = "amqp://guest:guest@127.0.0.1:5672/"
`

func TestRabbitMQConfig(t *testing.T) {
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

func TestRabbitMQInit(t *testing.T) {
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

func TestRabbitMQInstanceSimple(t *testing.T) {
	var (
		err    error
		obj    *Client
		config string
		ch     chan int
	)
	c.Convey("Define configuration", t, func() {
		config = configTpl
		c.Convey("Parse configuration", func() {
			err = beemod.Register(DefaultBuild).SetCfg([]byte(config), "toml").Run()
			c.So(err, c.ShouldBeNil)
			c.Convey("Set configuration group (initialization)", func() {
				obj = Invoker("dev")
				c.So(obj, c.ShouldNotBeNil)
				ch = make(chan int, 1)
				c.Convey("testing method", func() {
					mq := obj.NewRabbitMQ("simple", "test")
					go mq.Receive(func(delivery amqp.Delivery) {
						t.Log(string(delivery.Body))
						ch <- 1
					})
					err = mq.Publish("test")
					select {
					case <-ch:
						mq.Destroy()
						c.So(true, c.ShouldBeTrue)
					}
				})
			})
		})
	})
}