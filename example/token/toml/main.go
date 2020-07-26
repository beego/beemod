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
	mode = "mysql"
	redisTokenKeyPattern = "/egoshop/token/%d"
	accessTokenExpireInterval = 604800
	accessTokenIss           = "github.com/goecology/egoshop"
	accessTokenKey           = "ecologysK#xo"


[muses.token.default.mysql]
    debug = true
    level = "panic"
    network = "tcp"
    dialect = "mysql"
    addr = "127.0.0.1:3306"
    username = "root"
    password = "root"
    db = "beetest"
    charset = "utf8"
    parseTime = "True"
    loc = "Local"
    timeout = "1s"
    readTimeout = "1s"
    writeTimeout = "1s"
    maxOpenConns = 30
    maxIdleConns = 10
    connMaxLifetime = "300s"
	aliasName="default"

[muses.token.default.Logger]
    level = 7
    path = "token.log"
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

	fmt.Println(accessToken)

	fmt.Println(Client.CheckAccessToken(accessToken.AccessToken))

	fmt.Println(Client.DecodeAccessToken(accessToken.AccessToken))
}
