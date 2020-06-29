package main

import (
	"fmt"
	"github.com/beego-dev/beemod"
	"github.com/beego-dev/beemod/pkg/social"
)

var config = `
	[beego.oauth2.wx]
		appID = "app_id"
		debug = true
        appSecret = "app_secret"
		mode = "wx"
	[beego.oauth2.github]
		appID = "app_id"
		debug = true
        appSecret = "app_secret"
		mode = "github"
`

func main() {
	err := beemod.Register(
		social.DefaultBuild,
	).SetCfg([]byte(config), "toml").Run()
	if err != nil {
		panic("register err:" + err.Error())
	}
	client := social.Invoker("github")
	a, b := client.GetAccessToken("code")
	fmt.Println(a, b)

}
