package redis

import (
	"github.com/beego/beemod/pkg/datasource"
	"github.com/beego/beemod/pkg/module"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/beego/beemod/pkg/common"
)

var defaultCaller = &callerStore{
	Name: common.ModRedisName,
	Key:  module.ConfigPrefix + module.RedisName,
}

type callerStore struct {
	Name  string
	store sync.Map
	cfg   map[string]CallerCfg
	Key   string
}

type Client struct {
	pool *redis.Pool
}

func DefaultBuild() module.Invoker {
	return defaultCaller
}

func (c *callerStore) InitCfg(ds datasource.Datasource) error {
	c.cfg = make(map[string]CallerCfg, 0)
	ds.Range(c.Key, func(key string, name string) bool {
		config := CallerCfg{}
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
		db := Provider(cfg)
		defaultCaller.store.Store(name, db)
	}
	return nil
}

func Caller(name string) *Client {
	obj, ok := defaultCaller.store.Load(name)
	if !ok {
		return nil
	}
	return obj.(*Client)
}

func Invoker(name string) *Client {
	obj, ok := defaultCaller.store.Load(name)
	if !ok {
		return nil
	}
	return obj.(*Client)
}

func Provider(cfg CallerCfg) (resp *Client) {
	dialOptions := []redis.DialOption{
		redis.DialConnectTimeout(cfg.ConnectTimeout.Duration),
		redis.DialReadTimeout(cfg.ReadTimeout.Duration),
		redis.DialWriteTimeout(cfg.WriteTimeout.Duration),
		redis.DialDatabase(cfg.DB),
		redis.DialPassword(cfg.Password),
	}


	resp = &Client{
		pool: &redis.Pool{
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", cfg.Addr, dialOptions...)
				if err != nil {
					return nil, err
				}

				if cfg.Debug {
					return redis.NewLoggingConn(c, log.New(os.Stderr, "", log.LstdFlags), "redis"), nil
				}
				return c, nil
			},
			// Use the TestOnBorrow function to check the health of an idle connection
			// before the connection is returned to the application. This example PINGs
			// connections that have been idle more than a minute:
			//
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < time.Minute {
					return nil
				}
				_, err := c.Do("PING")
				return err
			},
			MaxIdle:     cfg.MaxActive,
			MaxActive:   cfg.MaxActive,
			IdleTimeout: cfg.IdleTimeout.Duration,
			Wait:        cfg.Wait, // wait until getting connection from pool
		},
	}
	return
}
