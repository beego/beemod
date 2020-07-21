// TODO: SMS implement
// Author: SDing <deen.job@qq.com>
// Date: 2020/6/28 - 3:07 PM

package sms

import (
	"errors"
	"github.com/beego-dev/beemod/pkg/module"
	"github.com/beego-dev/beemod/pkg/sms/alibaba"
	"github.com/beego-dev/beemod/pkg/sms/standard"
	"github.com/beego-dev/beemod/pkg/sms/tencent"
	"github.com/spf13/viper"
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

// todo with option
func (c *descriptor) Build() module.Invoker {
	return c
}

func (c *descriptor) InitCfg(cfg []byte, cfgType string) error {
	// todo ini cant unmarshal
	switch cfgType {
	case "toml":
		if err := viper.UnmarshalKey(c.Key, &c.cfg); err != nil {
			return err
		}
		// we need assign the default config, so we should unmarshal twice
		for name := range c.cfg {
			config := DefaultInvokerCfg
			if err := viper.UnmarshalKey(c.Key+"."+name, &config); err != nil {
				return err
			}
			c.cfg[name] = config
		}
	case "ini":
		panic("not implement ini")
	}
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

// disabled
func (c *descriptor) IsDisabled() bool {
	for _, cfg := range c.cfg {
		if cfg.Mode == "" && cfg.AccessKeyId == "" && cfg.AccessSecret == "" {
			return true
		}
	}
	return false
}

func provider(cfg InvokerCfg) (client standard.Sms, err error) {
	if cfg.Mode == "alibaba" {
		client, err = alibaba.New(cfg.Area, cfg.AccessKeyId, cfg.AccessSecret)
	} else if cfg.Mode == "ten" {
		client, err = tencent.New(cfg.Area, cfg.AccessKeyId, cfg.AccessSecret, cfg.domain, cfg.method, cfg.sign, cfg.timeOut)
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
