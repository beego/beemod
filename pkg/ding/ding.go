package ding

import (
	"bytes"
	"encoding/json"
	"github.com/beego/beemod/pkg/datasource"
	"github.com/beego/beemod/pkg/module"
	"io/ioutil"
	"net/http"
	"sync"
)

var defaultInvoker = &descriptor{
	Name: module.DingName,
	Key:  module.ConfigPrefix + module.DingName,
}

type descriptor struct {
	Name  string
	Key   string
	store sync.Map
	cfg   map[string]InvokerCfg
}

type Client struct {
	ss  *http.Client
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

// todo with option
func (c *descriptor) Build() module.Invoker {
	return c
}

func (c *descriptor) InitCfg(ds datasource.Datasource) error {
	c.cfg = make(map[string]InvokerCfg, 0)
	config := DefaultInvokerCfg
	ds.Range(c.Key, func(key string, name string) bool {
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
		ss := provider(cfg)
		c := &Client{
			ss,
			cfg,
		}
		defaultInvoker.store.Store(name, c)
	}
	return nil
}

// disabled
func (c *descriptor) IsDisabled() bool {
	for _, cfg := range c.cfg {
		if cfg.WebhookUrl == "" {
			return true
		}
	}
	return false
}

func provider(cfg InvokerCfg) (status *http.Client) {
	client := &http.Client{}
	return client
}

func (c *Client) SendMsg(msg string) (string, error) {
	content := make(map[string]string)
	data := make(map[string]interface{})
	content["content"] = msg
	data["msgtype"] = "text"
	data["text"] = content
	b, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", c.cfg.WebhookUrl, bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	client := c.ss
	req.Header.Set("Content-Type", "application/json") //这个一定要加，不加form的值post不过去，被坑了两小时

	resp, err := client.Do(req) //发送
	if err != nil {
		return "", err
	}
	defer resp.Body.Close() //一定要关闭resp.Body
	rdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(rdata), nil
}
