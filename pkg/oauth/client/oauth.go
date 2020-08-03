package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beemod/pkg/datasource"
	"github.com/beego/beemod/pkg/module"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"sync"
)

type BasicUserInfo struct {
	NickName string
	HeadIcon string
}

type BasicTokenInfo struct {
	AccessToken string
}

type OAuthService interface {
	LoginPage(option ...oauth2.AuthCodeOption) string
	GetAccessToken(state, code string, option ...oauth2.AuthCodeOption) (token *oauth2.Token, err error)
	GetUserInfo(token *oauth2.Token, info interface{}) (user interface{}, err error)
	GetCfg() InvokerCfg
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
	o   *oauth2.Config
	cfg InvokerCfg
}

// default invoker build
func DefaultBuild() module.Invoker {
	return defaultInvoker
}

// invoker
func Invoker(name string) OAuthService {
	obj, ok := defaultInvoker.store.Load(name)
	if !ok {
		return nil
	}
	return obj.(OAuthService)
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

func provider(cfg InvokerCfg) (client *oauth2.Config, err error) {
	client = &oauth2.Config{
		ClientID:     cfg.AppID,
		ClientSecret: cfg.AppSecret,
		RedirectURL:  cfg.RedirectURI,
		Endpoint: oauth2.Endpoint{
			AuthURL:  cfg.AuthURL,
			TokenURL: cfg.TokenURL,
		},
		Scopes: cfg.Scopes,
	}
	return
}

func (c *Client) LoginPage(option ...oauth2.AuthCodeOption) string {
	return c.o.AuthCodeURL(c.cfg.State, option...)
}

func (c *Client) GetAccessToken(state, code string, option ...oauth2.AuthCodeOption) (token *oauth2.Token, err error) {
	if state != c.cfg.State {
		err = errors.New(fmt.Sprintf("invalid oauth state, expected '%s', got '%s'\n", c.cfg.State, state))
	} else {
		token, err = c.o.Exchange(context.TODO(), code, option...)
	}
	return
}

func (c *Client) GetUserInfo(token *oauth2.Token, info interface{}) (user interface{}, err error) {
	var (
		resp *http.Response
		body []byte
	)
	resp, err = c.o.Client(context.Background(), token).Get(c.cfg.UserInfoURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &info)
	if err != nil {
		return
	}
	user = info
	return
}

func (c *Client) GetCfg() InvokerCfg {
	return c.cfg
}
