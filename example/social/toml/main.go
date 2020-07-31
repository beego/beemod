package main

import (
	"fmt"
	"github.com/beego/beemod"
	"github.com/beego/beemod/pkg/social"
)

var config = `
[beemod.oauth2.qq]
	debug = true
	mode  = "qq"
	app_id  = "app_id"
	app_secret  = "app_secret"
	redirectURI = "www.beego.com"
[beemod.oauth2.wx]
	debug = true
	mode  = "wx"
	app_id  = "app_id"
	app_secret  = "app_secret"
	redirectURI = "www.beego.com"
[beemod.oauth2.github]
	debug = true
	mode  = "github"
	app_id  = "app_id"
	app_secret  = "app_secret"
	redirectURI = "www.beego.com"
`

func main() {
	err := beemod.Register(
		social.DefaultBuild,
	).SetCfg([]byte(config), "toml").Run()
	if err != nil {
		panic("register err:" + err.Error())
	}
	client := social.Invoker("github")
	fmt.Println(client)
	a, b := client.GetAccessToken("code")
	fmt.Println(a, b)
}
