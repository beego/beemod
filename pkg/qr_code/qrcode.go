// TODO: qr_code create impl
// Author: SDing <deen.job@qq.com>
// Date: 2020/6/28 - 3:07 PM

package qr_code

import (
	"github.com/beego-dev/beemod/pkg/module"
	qr "github.com/higker/qrcode-go"
	"github.com/spf13/viper"
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

type QrCode struct {
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
func Invoker(name string) *QrCode {
	obj, ok := defaultInvoker.store.Load(name)
	if !ok {
		return nil
	}
	return obj.(*QrCode)
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
		qrInter := provider(cfg)
		defaultInvoker.store.Store(name, &QrCode{cfg: cfg, Of: qrInter})
	}
	return nil
}

func provider(cfg InvokerCfg) *QrEntity {
	return &QrEntity{avatarY: cfg.AvatarY, avatarX: cfg.AvatarX, Size: cfg.Size, foreground: cfg.Foreground}
}

// disabled
func (c *descriptor) IsDisabled() bool {
	for _, cfg := range c.cfg {
		if cfg.Mode == "" {
			return true
		}
	}
	return false
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
