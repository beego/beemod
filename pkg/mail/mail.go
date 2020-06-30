// TODO: mail mod impl
// Author: SDing <deen.job@qq.com>
// Date: 2020/6/28 - 12:07

package mail

import (
	"errors"
	"github.com/beego-dev/beemod/pkg/module"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
	"sync"
)

var defaultInvoker = &descriptor{
	Name: module.MailName,
	Key:  module.ConfigPrefix + module.MailName,
}

type descriptor struct {
	Name  string
	Key   string
	store sync.Map
	cfg   map[string]InvokerCfg
}

type Client struct {
	cfg    InvokerCfg
	dialer *gomail.Dialer
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
		dialer := provider(cfg)
		c := &Client{cfg: cfg, dialer: dialer}
		defaultInvoker.store.Store(name, c)
	}
	return nil
}

// disabled
func (c *descriptor) IsDisabled() bool {
	for _, cfg := range c.cfg {
		if cfg.Mode == "" {
			return true
		}
	}
	return false
}

func provider(cfg InvokerCfg) (dialer *gomail.Dialer) {
	return gomail.NewDialer(cfg.Host, cfg.Port, cfg.FromEmail, cfg.FromPassword)
}

func (c *Client) Push(msg *gomail.Message) error {
	// fmt.Println(c.dialer)
	err := c.dialer.DialAndSend(msg)
	if err != nil {
		return errors.New("beemod-mail : send email failed :" + err.Error())
	}
	return nil
}
