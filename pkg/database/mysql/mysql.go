/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/7/27 12:23
*/
package mysql

import (
	"github.com/astaxie/beego/orm"
	"github.com/beego/beemod/pkg/datasource"
	"github.com/beego/beemod/pkg/module"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var defaultCaller = &callerStore{
	Name: module.MysqlName,
	Key:  module.ConfigPrefix + module.MysqlName,
}

type callerStore struct {
	Name   string
	caller sync.Map
	cfg    map[string]CallerCfg
	Key    string
}

type Client struct {
	Cfg CallerCfg
	O   orm.Ormer
}

// default invoker build
func DefaultBuild() module.Invoker {
	return defaultCaller
}

func Invoker(name string) *Client {
	obj, ok := defaultCaller.caller.Load(name)
	if !ok {
		return nil
	}
	return obj.(*Client)
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
		o, err := Provider(cfg)
		if err != nil {
			return err
		}
		c := &Client{
			cfg,
			o,
		}
		defaultCaller.caller.Store(name, c)
	}

	return nil
}

func Provider(cfg CallerCfg) (o orm.Ormer, err error) {
	err = orm.RegisterDriver("mysql", orm.DRMySQL)
	err = orm.RegisterDataBase(cfg.AliasName, "mysql",
		cfg.Username+":"+cfg.Password+"@"+cfg.Network+"("+cfg.Addr+")/"+cfg.Db+
			"?charset="+cfg.Charset+"&parseTime="+cfg.ParseTime+"&loc="+cfg.Loc,
		cfg.MaxIdleConns,
		cfg.MaxOpenConns)

	o = orm.NewOrm()
	err = o.Using(cfg.AliasName) // 默认使用 default，你可以指定为其他数据库
	return
}
