package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/beego/beemod"
	"github.com/beego/beemod/pkg/database/mysql"
)

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}

type User struct {
	Id   int    `orm:"pk;auto"`
	Name string `orm:""`
}

func main() {
	var err error
	config := `
[beego.mysql.dev]
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
`
	err = beemod.Register(
		mysql.DefaultBuild,
	).SetCfg([]byte(config), "toml").Run()

	if err != nil {
		panic("register err:" + err.Error())
	}
	client := mysql.Invoker("dev")

	//创建表
	_ = orm.RunSyncdb(client.Cfg.AliasName, false, false)

	user := new(User)
	user.Name = "testName"
	fmt.Println(client.O.Insert(user))

}
