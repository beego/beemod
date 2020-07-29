package social

import (
	"errors"
	"github.com/beego/beemod/pkg/datasource"
	"github.com/beego/beemod/pkg/module"
	"sync"
)

type BasicUserInfo struct {
	NickName string
	HeadIcon string
}

type BasicTokenInfo struct {
	AccessToken string
}

type SocialService interface {
	LoginPage(state string) string
	GetAccessToken(code string) (*BasicTokenInfo, error)
	GetUserInfo(accessToken string) (*BasicUserInfo, error)
	GetType() string
}

var defaultInvoker = &descriptor{
	Name: module.Oauth2Name,
	Key:  module.ConfigPrefix + module.Oauth2Name,
}

type descriptor struct {
	Name  string
	Key   string
	store sync.Map
	cfg   map[string]InvokerCfg
}

type Client struct {
	SocialService
	cfg InvokerCfg
}

// default invoker build
func DefaultBuild() module.Invoker {
	return defaultInvoker
}

// invoker
func Invoker(name string) *Client {
	obj, ok := defaultInvoker.store.Load(name)
	if !ok {
		return nil
	}
	return obj.(*Client)
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
		db, err := provider(cfg)
		if err != nil {
			return err
		}
		c := &Client{
			db,
			cfg,
		}
		defaultInvoker.store.Store(name, c)
	}
	return nil
}

func provider(cfg InvokerCfg) (client SocialService, err error) {
	if cfg.Mode == "wx" {
		client = NewWxOauth2Service(cfg.AppID, cfg.AppSecret, cfg.RedirectURI)
	} else if cfg.Mode == "github" {
		client = NewGithubOauth2Service(cfg.AppID, cfg.AppSecret, cfg.RedirectURI)
	} else if cfg.Mode == "qq" {
		client = NewQQauth2Service(cfg.AppID, cfg.AppSecret, cfg.RedirectURI)
	} else {
		err = errors.New("oauth2 mode not exist")
		return
	}
	return
}
