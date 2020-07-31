/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/7/31 19:22
 */
package main

import (
	"github.com/beego/beemod"
	"github.com/beego/beemod/pkg/mq/rabbitmq"
	"time"
)

func main() {
	err := beemod.Register(
		rabbitmq.DefaultBuild,
	).SetCfg("./config.toml", "toml").Run()
	if err != nil {
		panic("register err:" + err.Error())
	}
	obj := rabbitmq.Invoker("dev")

	mq := obj.NewRabbitMQ("simple", "test")
	for {
		_ = mq.Publish("test")
		time.Sleep(3 * time.Second)
	}
}
