// Author: SDing <deen.job@qq.com>
// Date: 2020/6/28 - 3:07 PM

package sms

import (
	"errors"
	"github.com/beego/beemod/pkg/datasource"
	"github.com/beego/beemod/pkg/module"
	"github.com/beego/beemod/pkg/sms/alibaba"
	"github.com/beego/beemod/pkg/sms/standard"
	"github.com/beego/beemod/pkg/sms/tencent"
	"sync"
)

var defaultInvoker = &descriptor{
	Name: module.SmsName,
	Key:  module.ConfigPrefix + module.SmsName,
}

type descriptor struct {
	Name  string
	Key   string
	store sync.Map
	cfg   map[string]InvokerCfg
}

type Client struct {
	cfg    InvokerCfg
	client standard.Sms
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
		client, err := provider(cfg)
		if err != nil {
			return err
		}
		c := &Client{cfg: cfg, client: client}
		defaultInvoker.store.Store(name, c)
	}
	return nil
}

func provider(cfg InvokerCfg) (client standard.Sms, err error) {
	if cfg.Mode == "alibaba" {
		client, err = alibaba.New(cfg.Area, cfg.AccessKeyId, cfg.AccessSecret)
	} else if cfg.Mode == "ten" {
		client, err = tencent.New(cfg.Area, cfg.AccessKeyId, cfg.AccessSecret, cfg.Domain, cfg.Method, cfg.Sign, cfg.TimeOut)
	} else {
		panic(cfg.Mode + "not impl")
	}
	return
}
func (c *Client) Push(param map[string]string) (*standard.Response, error) {
	// fmt.Println(c.dialer)
	smsResponse, err := c.client.Send(param)
	if err != nil {
		return nil, errors.New("send sms failed :" + err.Error())
	}
	return smsResponse, nil
}
