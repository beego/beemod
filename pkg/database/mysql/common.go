package mysql

type Cfg struct {
	Muses struct {
		Mysql map[string]CallerCfg `toml:"mysql"`
	} `toml:"muses"`
}

type CallerCfg struct {
	Username       string
	Password       string
	Addr           string
	AliasName      string
	MaxIdleConns   int
	MaxOpenConns   int
	DefaultTimeLoc string
	Network        string
	Db             string
	Charset        string
	ParseTime      string
	Loc            string
}
