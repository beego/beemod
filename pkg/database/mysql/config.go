package mysql

type CallerCfg struct {
	Username       string `ini:"username"`
	Password       string `ini:"password"`
	Addr           string `ini:"addr"`
	AliasName      string `ini:"aliasName"`
	MaxIdleConns   int `ini:"maxIdleConns"`
	MaxOpenConns   int `ini:"maxOpenConns"`
	DefaultTimeLoc string `ini:"defaultTimeLoc"`
	Network        string `ini:"network"`
	Db             string `ini:"db"`
	Charset        string `ini:"charset"`
	ParseTime      string `ini:"parseTime"`
	Loc            string `ini:"loc"`
}
