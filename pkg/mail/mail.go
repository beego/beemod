// TODO: mail mod impl
// Author: SDing <deen.job@qq.com>
// Date: 2020/6/28 - 12:07

package mail

import (
	"errors"
	"github.com/beego/beemod/pkg/datasource"
	"github.com/beego/beemod/pkg/module"
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
