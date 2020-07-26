package session

import (
	"github.com/astaxie/beego/session"
	"github.com/beego/beemod/pkg/datasource"
	"github.com/beego/beemod/pkg/module"
	"sync"
)

var defaultInvoker = &descriptor{
	Name: module.SessionName,
	Key:  module.ConfigPrefix + module.SessionName,
}

type descriptor struct {
	Name  string
	Key   string
	store sync.Map
	cfg   map[string]InvokerCfg
}

// default invoker build
func DefaultBuild() module.Invoker {
	return defaultInvoker
}

// invoker
func Invoker(name string) *session.Manager {
	obj, ok := defaultInvoker.store.Load(name)
	if !ok {
		return nil
	}
	return obj.(*session.Manager)
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
		mng, err := provider(cfg)
		if err != nil {
			return err
		}

		defaultInvoker.store.Store(name, mng)
	}
	return nil
}

// disabled
func (c *descriptor) IsDisabled() bool {
	for _, cfg := range c.cfg {
		if cfg.MangerCfg == nil {
			return true
		}

		if cfg.Mode != "memory" && len(cfg.MangerCfg.ProviderConfig) <= 0 {
			return true
		}
	}
	return false
}

func provider(cfg InvokerCfg) (manager *session.Manager, err error) {
	manager, err = session.NewManager(cfg.Mode, cfg.MangerCfg)
	return
}
