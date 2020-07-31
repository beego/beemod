package logger

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/beego/beemod/pkg/datasource"
	"github.com/beego/beemod/pkg/module"
	"sync"
)

var defaultCaller = &callerStore{
	Name: module.LogName,
	Key:  module.ConfigPrefix + module.LogName,
}

type callerStore struct {
	Name   string
	caller sync.Map
	cfg    map[string]CallerCfg
	Key    string
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
	var config string
	switch cfg.Type {
	case "file":
		config = fmt.Sprintf(`{"filename":"%s","level":%v,"maxlines":%v,"maxsize":%v,"daily":%v,"maxdays":%v,"rotate":%v,"perm":"%v"}`,
			cfg.Path, cfg.Level, cfg.Maxlines, cfg.Maxsize, cfg.Daily, cfg.Maxdays, cfg.Rotate, cfg.Perm)
	case "multifile":
		config = fmt.Sprintf(`{"filename":"%s","level":%v,"maxlines":%v,"maxsize":%v,"daily":%v,"maxdays":%v,"rotate":%v,"perm":"%v","separate":"%v"}`,
			cfg.Path, cfg.Level, cfg.Maxlines, cfg.Maxsize, cfg.Daily, cfg.Maxdays, cfg.Rotate, cfg.Perm, cfg.Separate)
	case "smtp":
		config = fmt.Sprintf(`{"level":%v,"username":"%v","password":"%v","host":"%v","sendTos":%v,"subject":"%v"}`,
			cfg.Level, cfg.Username, cfg.Password, cfg.Host, cfg.SendTos, cfg.Subject)
	case "conn":
		config = fmt.Sprintf(`{"level":%v,"reconnectOnMsg":%v,"reconnect":%v,"net":"%v","addr":"%v"}`,
			cfg.Level, cfg.ReconnectOnMsg, cfg.Reconnect, cfg.Net, cfg.Addr)
	case "es":
		config = fmt.Sprintf(`{"level":%v,"dsn":"%v"}`,
			cfg.Level, cfg.Dsn)
	case "jianliao":
		config = fmt.Sprintf(`{"level":%v,"authorname":"%v","title":"%v","webhookurl":"%v","redirecturl":"%v","imageurl":"%v"}`,
			cfg.Level, cfg.Authorname, cfg.Title, cfg.Webhookurl, cfg.Redirecturl, cfg.Imageurl)
	case "slack":
		config = fmt.Sprintf(`{"level":%v,"webhookurl":"%v"}`,
			cfg.Level, cfg.SlackWebhookurl)
	case "alils":
		config = fmt.Sprintf(`{"level":%v}`,
			cfg.Level)
	default:
		cfg.Type = "console"
		config = fmt.Sprintf(`{"level":%v,"color":true}`,
			cfg.Level)
	}
	err := log.SetLogger(cfg.Type, config)
	if err != nil {
		log.Error(err.Error())
	}
	logs.EnableFuncCallDepth(true)

	return &Client{log, cfg}
}
