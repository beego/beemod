package module

import (
	"context"
	"github.com/beego/beemod/pkg/datasource"
	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
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

// Invoker Enable
type InvokerDisable interface {
	IsDisabled() bool
}

// Invoker Background
type InvokerBackground interface {
	RunBackground(ctx context.Context) error
}

func IsDisabled(invoker Invoker) bool {
	instance, ok := invoker.(InvokerDisable)
	return ok && instance.IsDisabled()
}

type InvokerFunc func() Invoker

type ConfigStore struct {
	Ini   *ini.File
	Viper *viper.Viper
}
