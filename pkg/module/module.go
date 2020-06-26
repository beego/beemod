package module

import "context"

// Descriptor
type Descriptor struct {
	Name    string
	Invoker Invoker
}

// Invoker
type Invoker interface {
	// Init cfg returns parse cfg error.
	InitCfg(cfg []byte, cfgType string) error
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
