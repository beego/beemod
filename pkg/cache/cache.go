package cache

import (
	"encoding/json"
	cache2 "github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/beego-dev/beemod/pkg/cache/standard"
	"github.com/beego-dev/beemod/pkg/module"
	"github.com/spf13/viper"
	"sync"
)

var defaultInvoker = &descriptor{
	Name: module.CacheName,
	Key:  module.ConfigPrefix + module.CacheName,
}

type descriptor struct {
	Name  string
	Key   string
	store sync.Map
	cfg   map[string]interface{}
}

type Client struct {
	standard.Cache
	cfg InvokerCfg
}

func DefaultBuild() module.Invoker {
	return defaultInvoker
}

func (c *descriptor) InitCfg(cfg []byte, cfgType string) error {
	switch cfgType {
	case "toml":
		if err := viper.UnmarshalKey(c.Key, &c.cfg); err != nil {
			return err
		}
		for name, v := range c.cfg {
			//parse to json str
			config, err := json.Marshal(v)
			if err != nil {
				return err
			}
			c.cfg[name] = string(config)
		}
	case "ini":
		panic("not implement ini")
	case "json":

	}
	return nil
}

func (c *descriptor) Run() error {
	for name, cfg := range c.cfg {
		cache, err := cache2.NewCache(name, cfg.(string))
		if err != nil {
			panic(err.Error())
		}
		c := &Client{
			cache,
			InvokerCfg{
				AdapterName: name,
				ConfigJson:  cfg.(string),
			},
		}
		defaultInvoker.store.Store(name, c)
	}
	return nil
}

func Invoker(name string) *Client {
	obj, ok := defaultInvoker.store.Load(name)
	if !ok {
		return nil
	}
	return obj.(*Client)
}
