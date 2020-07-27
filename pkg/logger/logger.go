package logger

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/beego/beemod/pkg/common"
	"github.com/beego/beemod/pkg/datasource"
	"github.com/beego/beemod/pkg/module"
	"sync"
)

var defaultCaller = &callerStore{
	Name: common.ModLoggerName,
	Key:  module.ConfigPrefix + module.LogName,
}

type callerStore struct {
	Name         string
	caller       sync.Map
	cfg          map[string]CallerCfg
	Key          string
}

type Client struct {
	*logs.BeeLogger
	cfg CallerCfg
}

func DefaultBuild() module.Invoker {
	return defaultCaller
}

func (c *callerStore) InitCfg(ds datasource.Datasource) error {
	c.cfg = make(map[string]CallerCfg, 0)
	var config CallerCfg
	ds.Range(c.Key, func(key string, name string) bool {
		if err := ds.Unmarshal(key, &config); err != nil {
			return false
		}
		c.cfg[name] = config
		return true
	})
	return nil
}

func (c *callerStore) Run() error {
	for name, cfg := range c.cfg {
		log := Provider(cfg)
		c := &Client{
			log.BeeLogger,
			cfg,
		}
		defaultCaller.caller.Store(name, c)
	}

	return nil
}

func Invoker(name string) *Client {
	obj, ok := defaultCaller.caller.Load(name)
	if !ok {
		return nil
	}
	return obj.(*Client)
}

func Provider(cfg CallerCfg) *Client {
	log := logs.NewLogger()
	err := log.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"%s","color":true,"level":%v}`, cfg.Path, cfg.Level))
	if err != nil {
		log.Error(err.Error())
	}
	logs.EnableFuncCallDepth(true)

	return &Client{log, cfg}
}
