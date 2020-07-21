package session

import (
	"github.com/astaxie/beego/session"
	"github.com/beego-dev/beemod/pkg/module"
	"github.com/spf13/viper"
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
