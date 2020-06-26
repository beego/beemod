package main

import (
	"github.com/beego-dev/beemod"
	"github.com/beego-dev/beemod/pkg/oss"
)

var config = `
	[beego.oss.myoss]
		debug = true
        isDeleteSrcPath = false
        cdnName = "http://127.0.0.1:8080/oss/"
        fileBucket = "oss"
`

func main() {
	err := beemod.Register(
		oss.DefaultBuild,
	).SetCfg([]byte(config), "ini").Run()
	if err != nil {
		panic("register err:" + err.Error())
	}
	obj := oss.Invoker("myoss")
	key := obj.GenerateKey("mock")

	err = obj.PutObjectFromFile(key, "../image/oss.jpg")
	if err != nil {
		panic(err)
	}
}
