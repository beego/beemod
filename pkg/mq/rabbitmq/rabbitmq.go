/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/7/31 9:49
 */
package rabbitmq

import (
	"fmt"
	"github.com/beego/beemod/pkg/datasource"
	"github.com/beego/beemod/pkg/module"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

var defaultInvoker = &descriptor{
	Name: module.RabbitmqName,
	Key:  module.ConfigPrefix + module.RabbitmqName,
}

type descriptor struct {
	Name  string
	Key   string
	store sync.Map
	cfg   map[string]InvokerCfg
}

type Client struct {
	cfg InvokerCfg
}

// default invoker build
func DefaultBuild() module.Invoker {
	return defaultInvoker
}

// invoker
func Invoker(name string) *Client {
	obj, ok := defaultInvoker.store.Load(name)
	if !ok {
		return nil
	}
	return obj.(*Client)
}

func (c *descriptor) InitCfg(ds datasource.Datasource) error {
	c.cfg = make(map[string]InvokerCfg, 0)
	ds.Range(c.Key, func(key string, name string) bool {
		config := DefaultInvokerCfg
		if err := ds.Unmarshal(key, &config); err != nil {
			return false
		}
		c.cfg[name] = config
		return true
	})
	return nil
}

func (c *descriptor) Run() error {
	for name, cfg := range c.cfg {
		c := provider(cfg)
		defaultInvoker.store.Store(name, c)
	}
	return nil
}

func provider(cfg InvokerCfg) (c *Client) {
	c = &Client{cfg: cfg}
	return
}

type RabbitImp interface {
	Publish(string) error
	Receive(func(amqp.Delivery))
	Destroy()
}

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind Key 名称
	Key string
	//连接信息
	Mqurl string
}

//if pubSub				NewRabbitMQ("pubSub",exchangeName)
//if routing  			NewRabbitMQ("routing",exchangeName,key)
//if topic  			NewRabbitMQ("topic",exchangeName,key)
//if simple and work  	NewRabbitMQ("simple",queueName)

func (c *Client) NewRabbitMQ(typeMq string, names ...string) RabbitImp {
	var (
		err error
		mq  *RabbitMQ
	)
	switch typeMq {
	case "pubSub":
		mq = &RabbitMQ{Exchange: names[0], Mqurl: c.cfg.Host}
		mq.conn, err = amqp.Dial(mq.Mqurl)
		mq.failOnErr(err, "failed to connect rabbitmq!")
		mq.channel, err = mq.conn.Channel()
		mq.failOnErr(err, "failed to open a channel")
		return &PubSub{mq}
	case "routing":
		mq = &RabbitMQ{Exchange: names[0], Key: names[1], Mqurl: c.cfg.Host}
		mq.conn, err = amqp.Dial(mq.Mqurl)
		mq.failOnErr(err, "failed to connect rabbitmq!")
		mq.channel, err = mq.conn.Channel()
		mq.failOnErr(err, "failed to open a channel")
		return &Routing{mq}
	case "topic":
		mq = &RabbitMQ{Exchange: names[0], Key: names[1], Mqurl: c.cfg.Host}
		mq.conn, err = amqp.Dial(mq.Mqurl)
		mq.failOnErr(err, "failed to connect rabbitmq!")
		mq.channel, err = mq.conn.Channel()
		mq.failOnErr(err, "failed to open a channel")
		return &Topic{mq}
	default:
		mq = &RabbitMQ{QueueName: names[0], Mqurl: c.cfg.Host}
		mq.conn, err = amqp.Dial(mq.Mqurl)
		mq.failOnErr(err, "failed to connect rabbitmq!")
		mq.channel, err = mq.conn.Channel()
		mq.failOnErr(err, "failed to open a channel")
		return &Simple{mq}
	}
}

func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		msg := fmt.Sprintf("%s:%s", message, err)
		log.Fatalf(msg)
	}
	return
}

type Simple struct {
	*RabbitMQ
}

func (r *Simple) Publish(message string) (err error) {
	_, err = r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exchange")
	err = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	return
}

func (r *Simple) Receive(f func(amqp.Delivery)) {
	var (
		err error
		q   amqp.Queue
		msg <-chan amqp.Delivery
	)
	q, err = r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exchange")

	msg, err = r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")

	forever := make(chan bool)
	go func() {
		for d := range msg {
			f(d)
		}
	}()

	fmt.Println("To exit, press CTRL+C")
	<-forever
}

func (r *Simple) Destroy() {
	_ = r.channel.Close()
	_ = r.conn.Close()
}

type PubSub struct {
	*RabbitMQ
}

func (r *PubSub) Publish(message string) (err error) {
	err = r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	r.failOnErr(err, "Failed to declare an exchange")

	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	return
}

func (r *PubSub) Receive(f func(amqp.Delivery)) {
	var (
		err error
		q   amqp.Queue
		msg <-chan amqp.Delivery
	)
	err = r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"fanout",
		true,
		false,
		//YES表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exchange")

	q, err = r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")

	err = r.channel.QueueBind(
		q.Name,
		"",
		r.Exchange,
		false,
		nil)

	msg, err = r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msg {
			f(d)
		}
	}()

	fmt.Println("To exit, press CTRL+C")
	<-forever
}

func (r *PubSub) Destroy() {
	_ = r.channel.Close()
	_ = r.conn.Close()
}

type Routing struct {
	*RabbitMQ
}

func (r *Routing) Publish(message string) (err error) {
	err = r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	r.failOnErr(err, "Failed to declare an exchange")

	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	return
}

func (r *Routing) Receive(f func(amqp.Delivery)) {
	var (
		err error
		q   amqp.Queue
		msg <-chan amqp.Delivery
	)
	err = r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exchange")
	q, err = r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")

	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil)

	msg, err = r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msg {
			f(d)
		}
	}()

	fmt.Println("To exit, press CTRL+C")
	<-forever
}

func (r *Routing) Destroy() {
	_ = r.channel.Close()
	_ = r.conn.Close()
}

type Topic struct {
	*RabbitMQ
}

func (r *Topic) Publish(message string) (err error) {
	err = r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	r.failOnErr(err, "Failed to declare an exchange")

	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	return
}

func (r *Topic) Receive(f func(amqp.Delivery)) {
	var (
		err error
		q   amqp.Queue
		msg <-chan amqp.Delivery
	)
	err = r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exchange")
	q, err = r.channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")

	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil)

	//消费消息
	msg, err = r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msg {
			f(d)
		}
	}()

	fmt.Println("To exit, press CTRL+C")
	<-forever
}

func (r *Topic) Destroy() {
	_ = r.channel.Close()
	_ = r.conn.Close()
}
