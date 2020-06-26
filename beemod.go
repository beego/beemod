package beemod

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/beego-dev/beemod/pkg/module"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type BeeMod struct {
	invokers    []module.Invoker
	cfgByte     []byte
	cfgType     string
	isSetConfig bool
	filePath    string
	ext         string
	err         error
}

var (
	BEEMOD_DEBUG = false
)

// init debug
func init() {
	if os.Getenv("BEEMOD_DEBUG") == "true" {
		BEEMOD_DEBUG = true
	}
}

// Register Interface
func Register(invokerFuncs ...module.InvokerFunc) (obj *BeeMod) {
	obj = &BeeMod{}
	invokers, err := sortInvokers(invokerFuncs)
	if err != nil {
		obj.err = err
		return
	}
	obj.invokers = invokers
	return
}

// 设置配置
func (m *BeeMod) SetCfg(cfg interface{}, cfgType string) *BeeMod {
	if m.err != nil {
		return m
	}
	var err error
	var cfgByte []byte
	switch cfgInfo := cfg.(type) {
	case string:
		m.filePath = cfgInfo
		err = isPathExist(m.filePath)
		if err != nil {
			m.err = err
			return m
		}

		ext := filepath.Ext(m.filePath)

		if len(ext) <= 1 {
			m.err = errors.New("config file ext is error")
			return m
		}
		m.ext = ext[1:]
		cfgByte, err = parseFile(cfgInfo)
		if err != nil {
			m.err = err
			return m
		}
	case []byte:
		cfgByte = cfgInfo
	default:
		m.err = fmt.Errorf("type is error %s", cfg)
		return m
	}
	m.cfgType = cfgType
	m.cfgByte = cfgByte
	m.isSetConfig = true
	m.initViper()
	return m
}

func (m *BeeMod) Run() (err error) {
	if !m.isSetConfig {
		err = errors.New("bee mod need set config")
		return
	}

	if m.err != nil {
		err = m.err
		return
	}
	for _, invoker := range m.invokers {
		name := getCallerName(invoker)
		logs.Info("module", name, "cfg start")
		if err = invoker.InitCfg(m.cfgByte, m.cfgType); err != nil {
			logs.Info("module", name, "init config error")
			return
		}

		logs.Info("module", name, "cfg end")

		// is invoker enabled
		if module.IsDisabled(invoker) {
			logs.Critical("module", name, "not enabled")
			panic("module" + name + "not enabled")
		}

		logs.Info("module", name, "run start")
		if err = invoker.Run(); err != nil {
			logs.Info("module", name, "run error")
			return
		}
		logs.Info("module", name, "run ok")
	}

	return
}

func (m *BeeMod) initViper() {
	rBytes := bytes.NewReader(m.cfgByte)
	cfgType := m.ext
	if m.cfgType != "" {
		cfgType = m.cfgType
	}
	viper.SetConfigType(cfgType)
	viper.AutomaticEnv() // read in environment variables that match
	err := viper.ReadConfig(rBytes)
	if err != nil {
		logs.Critical("Using config file:", viper.ConfigFileUsed())
	}
	if BEEMOD_DEBUG {
		viper.Debug()
	}
}

func isPathExist(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	return err
}
