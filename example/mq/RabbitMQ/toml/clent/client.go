/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/7/31 19:21
 */
package main

import (
	"github.com/beego/beemod"
	"github.com/beego/beemod/pkg/mq/rabbitmq"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	err := beemod.Register(
		rabbitmq.DefaultBuild,
	).SetCfg("../config.toml", "toml").Run()
	if err != nil {
		panic("register err:" + err.Error())
	}
	obj := rabbitmq.Invoker("dev")

	mq := obj.NewRabbitMQ("simple", "test")
	mq.Receive(func(delivery amqp.Delivery) {
		log.Print(string(delivery.Body) + "\n")
	})
}
