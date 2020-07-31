package main

import (
	"fmt"
	"github.com/beego/beemod"
	"github.com/beego/beemod/pkg/token"
	"time"
)

func main() {
	var err error
	config := `
[beego.token.dev]
	mode = "mysql"
	redisTokenKeyPattern = "/egoshop/token/%d"
	accessTokenExpireInterval = 604800
	accessTokenIss           = "github.com/goecology/egoshop"
	accessTokenKey           = "ecologysK#xo"
[beego.token.dev.mysql]
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
    maxOpenConns = 30
    maxIdleConns = 10
    connMaxLifetime = "300s"
	aliasName="default"
[beego.token.dev.Logger]
    level = 7
    path = "token.log"
`
	err = beemod.Register(
		token.DefaultBuild,
	).SetCfg([]byte(config), "toml").Run()

	if err != nil {
		panic("register err:" + err.Error())
	}

	Client := token.Invoker("dev")

	accessToken, err := Client.CreateAccessToken(123123, time.Now().Unix())
	if err != nil {
		panic("CreateAccessToken err:" + err.Error())
	}

	fmt.Println(accessToken)

	fmt.Println(Client.CheckAccessToken(accessToken.AccessToken))

	fmt.Println(Client.DecodeAccessToken(accessToken.AccessToken))
}
