// Author: SDing <deen.job@qq.com>
// Date: 2020/6/28 - 3:07 PM

package qr_code

import (
	"github.com/beego/beemod/pkg/datasource"
	"github.com/beego/beemod/pkg/module"
	qr "github.com/higker/qrcode-go"
	"sync"
)

var defaultInvoker = &descriptor{
	Name: module.QrcodeName,
	Key:  module.ConfigPrefix + module.QrcodeName,
}

type descriptor struct {
	Name  string
	Key   string
	store sync.Map
	cfg   map[string]InvokerCfg
}

type Client struct {
	cfg InvokerCfg
	Of  *QrEntity
}

// Custom QrCode struct
type QrEntity struct {
	avatarX, avatarY, Size int
	foreground             string
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
		qrInter := provider(cfg)
		defaultInvoker.store.Store(name, &Client{cfg: cfg, Of: qrInter})
	}
	return nil
}

func provider(cfg InvokerCfg) *QrEntity {
	return &QrEntity{avatarY: cfg.AvatarY, avatarX: cfg.AvatarX, Size: cfg.Size, foreground: cfg.Foreground}
}

func (qe *QrEntity) NewFg(content, src string) error {
	code, err := qr.New(content, qr.Highest)
	if err != nil {
		return err
	}
	code.DisableBorder(true)
	code.SetForegroundImage(qe.foreground)
	err = code.WriteFile(qe.Size, src)
	if err != nil {
		return err
	}
	return nil
}
func (qe *QrEntity) New(content, src string) error {
	code, err := qr.New(content, qr.Highest)
	if err != nil {
		return err
	}
	code.DisableBorder(true)
	err = code.WriteFile(qe.Size, src)
	if err != nil {
		return err
	}
	return nil
}

func (qe *QrEntity) NewAvatar(content, srcAvatar, src string, size int, enable bool) error {
	code, err := qr.New(content, qr.Highest)
	if err != nil {
		return err
	}
	code.DisableBorder(true)
	if enable {
		code.SetForegroundImage(qe.foreground)
	}
	code.SetAvatar(&qr.Avatar{
		Src:    srcAvatar,
		Width:  size,
		Height: size,
	})
	err = code.WriteFile(qe.Size, src)
	if err != nil {
		return err
	}
	return nil
}
