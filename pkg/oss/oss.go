package oss

import (
	"errors"
	"github.com/beego/beemod/pkg/datasource"
	"github.com/beego/beemod/pkg/module"
	"github.com/beego/beemod/pkg/oss/alioss"
	"github.com/beego/beemod/pkg/oss/file"
	"github.com/beego/beemod/pkg/oss/standard"
	"github.com/satori/uuid"
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
