package social

import (
	"errors"
	"github.com/beego-dev/beemod/pkg/module"
	"github.com/spf13/viper"
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

// disabled
func (c *descriptor) IsDisabled() bool {
	return false
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
