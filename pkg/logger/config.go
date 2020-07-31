package logger

type CallerCfg struct {
	Level int `ini:"level"`
	Path  string `ini:"path"`
}
