package oss

import (
	"errors"
	"github.com/beego-dev/beemod/pkg/module"
	"github.com/beego-dev/beemod/pkg/oss/alioss"
	"github.com/beego-dev/beemod/pkg/oss/file"
	"github.com/beego-dev/beemod/pkg/oss/standard"
	"github.com/satori/uuid"
	"github.com/spf13/viper"
	"strings"
	"sync"
	"time"
)

var defaultInvoker = &descriptor{
	Name: module.OssName,
	Key:  module.ConfigPrefix + module.OssName,
}

type descriptor struct {
	Name  string
	Key   string
	store sync.Map
	cfg   map[string]InvokerCfg
}

type Client struct {
	standard.Oss
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
		for name, _ := range c.cfg {
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
	for _, cfg := range c.cfg {
		if cfg.Mode == "alioss" && cfg.AccessKeyID == "" && cfg.AccessKeySecret == "" {
			return true
		}
	}
	return false
}

func provider(cfg InvokerCfg) (client standard.Oss, err error) {
	if cfg.Mode == "alioss" {
		client, err = alioss.NewOss(cfg.Addr, cfg.AccessKeyID, cfg.AccessKeySecret, cfg.OssBucket, cfg.IsDeleteSrcPath)
	} else if cfg.Mode == "file" {
		client, err = file.NewOss(cfg.CdnName, cfg.FileBucket, cfg.IsDeleteSrcPath)
	} else {
		err = errors.New("oss mode not exist")
		return
	}
	return
}

func (c *Client) ShowImg(img string, style ...string) (url string) {
	if strings.HasPrefix(img, "https://") || strings.HasPrefix(img, "http://") {
		return img
	}
	img = strings.TrimLeft(img, "./")
	switch c.cfg.Mode {
	case "alioss":
		s := ""
		if len(style) > 0 && strings.TrimSpace(style[0]) != "" {
			s = "/" + style[0]
		}
		url = img + s
	case "file":
		url = img
	}
	url = c.cfg.CdnName + url
	return
}

func (c *Client) ShowImgArr(imgs []string, style ...string) (urlArr []string) {
	urlArr = make([]string, 0)
	for _, img := range imgs {
		urlArr = append(urlArr, c.ShowImg(img, style...))
	}
	return
}

func (c *Client) GenerateKey(prefix string) string {
	month := time.Now().Format("200601")
	// object路径开头不能与“/”
	return prefix + "/" + month + "/" + strings.ReplaceAll(uuid.NewV4().String(), "-", "") + ".jpg"
}
