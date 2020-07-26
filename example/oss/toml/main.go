package main

import (
	"github.com/beego/beemod"
	"github.com/beego/beemod/pkg/oss"
)

var config = `
	[beego.oss.myoss]
		mode = "file"
		debug = true
        isDeleteSrcPath = false
        cdnName = "http://127.0.0.1:8080/oss/"
        fileBucket = "oss"
`

func main() {
	err := beemod.Register(
		oss.DefaultBuild,
	).SetCfg([]byte(config), "toml").Run()
	if err != nil {
		panic("register err:" + err.Error())
	}

	ossClient := oss.Invoker("myoss")
	dstPath := ossClient.GenerateKey("mock")

	err = ossClient.PutObjectFromFile(dstPath, "../image/oss.jpg")
	if err != nil {
		panic(err)
	}
}
