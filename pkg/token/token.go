package token

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/beego/beemod/pkg/datasource"
	"github.com/beego/beemod/pkg/module"
	"sync"

	mysqlToken "github.com/beego/beemod/pkg/token/mysql"
	redis2 "github.com/beego/beemod/pkg/token/redis"
	"github.com/beego/beemod/pkg/token/standard"

	"github.com/beego/beemod/pkg/cache/redis"
	"github.com/beego/beemod/pkg/common"
	"github.com/beego/beemod/pkg/database/mysql"
	"github.com/beego/beemod/pkg/logger"
)

var defaultCallerStore = &callerStore{
	Name: common.ModTokenName,
	Key:  module.ConfigPrefix + module.TokenName,
}

type callerStore struct {
	Name   string
	caller sync.Map
	Key    string
	cfg    map[string]InvokerCfg
}

type Client struct {
	standard.TokenAccessor
	cfg InvokerCfg
}

// default invoker build
func DefaultBuild() module.Invoker {
	return defaultCallerStore
}

func (c *callerStore) InitCfg(ds datasource.Datasource) error {
	c.cfg = make(map[string]InvokerCfg, 0)
	var config InvokerCfg
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
		accessor, err := provider(cfg)
		if err != nil {
			return err
		}
		c := &Client{
			accessor,
			cfg,
		}
		defaultCallerStore.caller.Store(name, c)
	}

	return nil
}

func Invoker(name string) *Client {
	obj, ok := defaultCallerStore.caller.Load(name)

	redis2.InitTokenKeyPattern(obj.(*Client).cfg.RedisTokenKeyPattern)
	standard.InitAccessToken(obj.(*Client).cfg.AccessTokenIss, obj.(*Client).cfg.AccessTokenKey)
	standard.InitAccessTokenExpireInterval(obj.(*Client).cfg.AccessTokenExpireInterval)

	_ = orm.RunSyncdb(obj.(*Client).cfg.Mysql.AliasName, false, false)

	if !ok {
		return nil
	}
	return obj.(*Client)
}

func provider(cfg InvokerCfg) (client standard.TokenAccessor, err error) {
	var loggerClient *logger.Client

	// 如果没有引用的logger，就创建一个
	if len(cfg.LoggerRef) > 0 {
		loggerClient = logger.Invoker(cfg.LoggerRef)
	} else {
		loggerClient = logger.Provider(logger.CallerCfg(cfg.Logger))
	}
	if cfg.Mode == "mysql" {
		return createMysqlAccessor(cfg, loggerClient)
	} else if cfg.Mode == "redis" {
		return createRedisAccessor(cfg, loggerClient)
	} else {
		return nil, errors.New("The token's mode must be redis or mysql: " + cfg.Mode)
	}
}

func createMysqlAccessor(cfg InvokerCfg, loggerClient *logger.Client) (accessor standard.TokenAccessor, err error) {
	var db orm.Ormer
	if len(cfg.MysqlRef) > 0 {
		db = mysql.Invoker(cfg.MysqlRef).O
	} else {
		db, err = mysql.Provider(mysql.CallerCfg(cfg.Mysql))
		if err != nil {
			return
		}
	}
	return mysqlToken.InitTokenAccessor(loggerClient, db), nil
}

func createRedisAccessor(cfg InvokerCfg, loggerClient *logger.Client) (standard.TokenAccessor, error) {
	var redisClient *redis.Client
	if len(cfg.RedisRef) > 0 {
		redisClient = redis.Caller(cfg.RedisRef)
	} else {
		redisClient = redis.Provider(redis.CallerCfg(cfg.Redis))
	}

	return redis2.InitRedisTokenAccessor(loggerClient, redisClient), nil
}
