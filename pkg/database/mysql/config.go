package mysql

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
