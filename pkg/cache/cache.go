package cache

import (
	"encoding/json"
	cache2 "github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/beego/beemod/pkg/cache/standard"
	"github.com/beego/beemod/pkg/datasource"
	"github.com/beego/beemod/pkg/module"
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
	cfg   map[string]InvokerCfg
}

type Client struct {
	standard.Cache
	cfg InvokerCfg
}

func DefaultBuild() module.Invoker {
	return defaultInvoker
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
		jsonbyte, err := json.Marshal(cfg)
		if err != nil {
			panic(err)
		}
		cache, err := cache2.NewCache(name, string(jsonbyte))
		if err != nil {
			panic(err.Error())
		}
		c := &Client{
			cache,
			cfg,
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
