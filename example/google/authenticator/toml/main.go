/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/8/6 19:09
 */
package main

import (
	"fmt"
	"github.com/beego/beemod"
	"github.com/beego/beemod/pkg/google/authenticator"
)

var config = `
  [beego.authenticator.dev]
	user = "admin123"
	iss = "admin@beego.com"
`

func main() {
	err := beemod.Register(
		authenticator.DefaultBuild,
	).SetCfg([]byte(config), "toml").Run()
	if err != nil {
		panic("register err:" + err.Error())
	}
	obj := authenticator.Invoker("dev")
	key := obj.Auth.GenerateKey()
	fmt.Println(key)
	uri := obj.Auth.ProvisionURI()
	fmt.Println(uri)
	ok, err := obj.Auth.Authenticate("123456")
	fmt.Println(ok)
}
