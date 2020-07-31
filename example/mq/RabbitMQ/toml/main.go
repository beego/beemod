/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/7/31 15:04
 */
package main

import (
	"github.com/beego/beemod"
	"github.com/beego/beemod/pkg/mq/rabbitmq"
	"github.com/streadway/amqp"
	"log"
	"time"
)

const config = `
  [beego.rabbitmq.dev]
	host = "amqp://guest:guest@127.0.0.1:5672/"
`

func main() {
	err := beemod.Register(
		rabbitmq.DefaultBuild,
	).SetCfg([]byte(config), "toml").Run()
	if err != nil {
		panic("register err:" + err.Error())
	}
	obj := rabbitmq.Invoker("dev")

	mq := obj.NewRabbitMQ("simple", "test")
	go mq.Receive(func(delivery amqp.Delivery) {
		log.Print(string(delivery.Body) + "\n")
	})
	for {
		_ = mq.Publish("test")
		time.Sleep(3 * time.Second)
	}

}
