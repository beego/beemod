package token

import (
	"github.com/beego/beemod/pkg/cache/redis"
	"github.com/beego/beemod/pkg/database/mysql"
	"github.com/beego/beemod/pkg/logger"
)

type MysqlCallerCfg mysql.CallerCfg
type RedisCallerCfg redis.CallerCfg
type LoggerCallerCfg logger.CallerCfg

// CallerCfg是token的配置。
// 需要注意的是，XXXRef是指指定了一个已经在配置文件里面的Caller。
// 比如说，你已经设置了一个mysql数据库myDB用于存储数据，而你又希望同时使用该数据库来存放token的数据，
// 那么你只需将MysqlRef设置为myDB。
// 如果你没有指定Ref，那么token会在初始化的时候依据配置来重新创建一个
type InvokerCfg struct {
	Mode string

	RedisTokenKeyPattern      string `toml:"redisTokenKeyPattern"`
	AccessTokenExpireInterval int64  `toml:"accessTokenExpireInterval"`
	AccessTokenIss            string `toml:"accessTokenIss"`
	AccessTokenKey            string `toml:"accessTokenKey"`

	LoggerRef string
	Logger    LoggerCallerCfg `toml:"logger"`

	MysqlRef string
	Mysql    MysqlCallerCfg `toml:"mysql"`

	RedisRef string
	Redis    RedisCallerCfg `toml:"redis"`
}
