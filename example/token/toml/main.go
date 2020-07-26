package main

import (
	"fmt"
	"github.com/beego-dev/beemod/pkg/token"
	"time"
)

func main() {
	var err error
	config := `


[muses.token.default]
	mode = "redis"
	mysqlTableName  = "access_token"
	redisTokenKeyPattern = "/egoshop/token/%d"
	accessTokenExpireInterval = 604800
	accessTokenIss           = "github.com/goecology/egoshop"
	accessTokenKey           = "ecologysK#xo"


[muses.token.default.redis]
	mode = "redis"
    debug = true
    addr = "127.0.0.1:6379"
    network = "tcp"
    db = 0
    password = ""
[muses.token.default.Logger]
	debug = true
    level = "debug"
    path = "./token.log"
`

	store := token.Register()
	err = store.InitCfg([]byte(config))
	if err != nil {
		panic("InitCfg err:" + err.Error())
	}

	err = store.InitCaller()
	if err != nil {
		panic("InitCfg err:" + err.Error())
	}

	Client := token.Caller("default")

	accessToken, err := Client.CreateAccessToken(123123, time.Now().Unix())
	if err != nil {
		panic("CreateAccessToken err:" + err.Error())
	}

	fmt.Println(Client.DecodeAccessToken(accessToken.AccessToken))

	fmt.Println(accessToken)
}
