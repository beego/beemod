package beemod

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/beego/beemod/pkg/datasource"
	"github.com/beego/beemod/pkg/module"
	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
	"os"
)

type BeeMod struct {
	invokers    []module.Invoker
	cfgType     string
	isSetConfig bool
	err         error
	ds          datasource.Datasource
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

// Set Config
func (m *BeeMod) SetCfg(cfg interface{}, cfgType string) *BeeMod {
	if m.err != nil {
		return m
	}
	module.Config = &module.ConfigStore{}
	switch cfgInfo := cfg.(type) {
	case string:
		if cfgType == "ini" {
			f, err := ini.Load(cfgInfo)
			if err != nil {
				m.err = err
				return m
			}
			m.ds = &datasource.Ini{
				Ini: f,
			}
		} else {
			viper.SetConfigFile(cfgInfo)
			if cfgType != "" {
				viper.SetConfigType(cfgType)
			}
			viper.AutomaticEnv() // read in environment variables that match
			err := viper.ReadInConfig()
			if err != nil {
				logs.Critical("Using config file:", viper.ConfigFileUsed())
				m.err = err
				return m
			}
			m.ds = &datasource.Toml{
				Viper: viper.GetViper(),
			}
		}
	case []byte:
		if cfgType == "ini" {
			f, err := ini.Load(cfgInfo)
			if err != nil {
				m.err = err
				return m
			}
			m.ds = &datasource.Ini{
				Ini: f,
			}
		} else {
			rBytes := bytes.NewReader(cfgInfo)
			if cfgType != "" {
				viper.SetConfigType(cfgType)
			}
			viper.AutomaticEnv() // read in environment variables that match
			err := viper.ReadConfig(rBytes)
			if err != nil {
				logs.Critical("Using config file:", viper.ConfigFileUsed())
				m.err = err
				return m
			}
			m.ds = &datasource.Toml{
				Viper: viper.GetViper(),
			}
		}
	default:
		m.err = fmt.Errorf("type is error %s", cfg)
		return m
	}
	m.cfgType = cfgType
	m.isSetConfig = true
	if BEEMOD_DEBUG {
		viper.Debug()
	}
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
		if err = invoker.InitCfg(m.ds); err != nil {
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
