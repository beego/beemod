package module

import (
	"github.com/beego/beemod/pkg/datasource"
	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
	"time"
)

// global config
var Config *ConfigStore

// Descriptor
type Descriptor struct {
	Name    string
	Invoker Invoker
}

// InvokerRegister
type Invoker interface {
	// Init cfg returns parse cfg error.
	InitCfg(ds datasource.Datasource) error
	// Init Caller returns init caller error
	Run() error
}

type InvokerFunc func() Invoker

type ConfigStore struct {
	Ini   *ini.File
	Viper *viper.Viper
}

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}
