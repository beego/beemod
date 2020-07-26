package mysql

import (
	"github.com/BurntSushi/toml"
	"github.com/astaxie/beego/orm"
	"github.com/beego/beemod/pkg/common"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var defaultCaller = &callerStore{
	Name: common.ModMysqlName,
}

type callerStore struct {
	Name   string
	caller sync.Map
	cfg    Cfg
}

func Register() common.Caller {
	return defaultCaller
}

func Caller(name string) orm.Ormer {
	obj, ok := defaultCaller.caller.Load(name)
	if !ok {
		return nil
	}
	return obj.(orm.Ormer)
}

func (c *callerStore) InitCfg(cfg []byte) error {
	if err := toml.Unmarshal(cfg, &c.cfg); err != nil {
		return err
	}
	return nil
}

func (c *callerStore) InitCaller() error {
	for name, cfg := range c.cfg.Muses.Mysql {
		db, err := Provider(cfg)
		if err != nil {
			return err
		}
		defaultCaller.caller.Store(name, db)
	}
	return nil
}

func Provider(cfg CallerCfg) (o orm.Ormer, err error) {
	err = orm.RegisterDriver("mysql", orm.DRMySQL)
	err = orm.RegisterDataBase(cfg.AliasName, "mysql", cfg.Username+":"+cfg.Password+"@"+cfg.Network+"("+cfg.Addr+")/"+cfg.Db+
		"?charset="+cfg.Charset+"&parseTime="+cfg.ParseTime+"&loc="+cfg.Loc,
		cfg.MaxIdleConns,
		cfg.MaxOpenConns)

	o = orm.NewOrm()
	err = o.Using(cfg.AliasName) // 默认使用 default，你可以指定为其他数据库
	return
}
