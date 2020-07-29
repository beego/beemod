package logger

type Cfg struct {
	Muses struct {
		Logger map[string]CallerCfg `toml:"logger"`
	} `toml:"muses"`
}

type CallerCfg struct {
	Level int
	Path  string
}
