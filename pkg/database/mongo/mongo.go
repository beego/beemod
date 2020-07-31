/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/7/30 10:34
 */
package mongo

import (
	"fmt"
	"github.com/beego/beemod/pkg/datasource"
	"github.com/beego/beemod/pkg/module"
	"github.com/globalsign/mgo"
	"log"
	"sync"
)

type ClientImp interface {
	connect(db, collection string) (*mgo.Session, *mgo.Collection)
	Insert(db, collection string, docs ...interface{}) error
	FindOne(db, collection string, query, selector, result interface{}) error
	FindAll(db, collection string, query, selector, result interface{}) error
	Update(db, collection string, query, update interface{}) error
	Remove(db, collection string, query interface{}) error
}

type Client struct {
	cfg CallerCfg
	m   *mgo.Session
}

type callerStore struct {
	Name   string
	caller sync.Map
	cfg    map[string]CallerCfg
	Key    string
}

var defaultCaller = &callerStore{
	Name: module.MongoName,
	Key:  module.ConfigPrefix + module.MongoName,
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
		o, err := provider(cfg)
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

func Invoker(name string) *Client {
	obj, ok := defaultCaller.caller.Load(name)
	if !ok {
		return nil
	}
	return obj.(*Client)
}

func provider(cfg CallerCfg) (resp *mgo.Session, err error) {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{cfg.URL},
		Source:   cfg.Source,
		Username: cfg.User,
		Password: cfg.Password,
	}
	resp, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalln("create session error ", err)
	}
	mgo.SetLogger(&cLogger{})
	mgo.SetDebug(cfg.Debug)
	// Optional. Switch the session to a monotonic behavior.
	resp.SetMode(mgo.Monotonic, true)
	return resp, err
}

type cLogger struct{}

func (c *cLogger) Output(calldepth int, s string) error {
	fmt.Println("calldepth: ", calldepth, ", s: ", s)
	return nil
}

func (c *Client) connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	s := c.m.Copy()
	return s, s.DB(db).C(collection)
}

func (c *Client) Insert(db, collection string, docs ...interface{}) error {
	ms, mc := c.connect(db, collection)
	defer ms.Close()
	return mc.Insert(docs...)
}

func (c *Client) FindOne(db, collection string, query, selector, result interface{}) error {
	ms, mc := c.connect(db, collection)
	defer ms.Close()
	return mc.Find(query).Select(selector).One(result)
}

func (c *Client) FindAll(db, collection string, query, selector, result interface{}) error {
	ms, mc := c.connect(db, collection)
	defer ms.Close()
	return mc.Find(query).Select(selector).All(result)
}

func (c *Client) Update(db, collection string, query, update interface{}) error {
	ms, mc := c.connect(db, collection)
	defer ms.Close()
	return mc.Update(query, update)
}

func (c *Client) Remove(db, collection string, query interface{}) error {
	ms, mc := c.connect(db, collection)
	defer ms.Close()
	return mc.Remove(query)
}
