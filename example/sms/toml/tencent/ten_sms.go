// Author: SDing <deen.job@qq.com>
// Date: 2020/6/28 - 5:40 PM

package main

import (
	"fmt"
	"github.com/beego/beemod"
	"github.com/beego/beemod/pkg/sms"
)

// custom config toml template
var tenConfig = `
  [beego.sms.ten]
	debug = false
	mode = "ten"
	area = "ap-guangzhou"
	accessKeyId = ""
	accessSecret = ""
`

func main() {
	err := beemod.Register(
		sms.DefaultBuild,
	).SetCfg([]byte(tenConfig), "toml").Run()
	if err != nil {
		panic("register err:" + err.Error())
	}
	client := sms.Invoker("ten")
	// send tencent sms
	// see tencent SDK API
	//https://cloud.tencent.com/document/product/382/43199
	param := map[string]string{
		"SmsSdkAppid":      "1400787878",
		"Sign":             "xxx",
		"SenderId":         "xxx",
		"SessionContext":   "XXX",
		"ExtendCode":       "0",
		"TemplateParamSet": "0,1,2",
		"TemplateID":       "449739",
		"PhoneNumberSet":   "+8613711112222,+8613711112221,+8613711112223",
	}
	fmt.Println(param)
	result, err := client.Push(param)
	if err != nil {
		fmt.Println(err)
	}
	// response
	fmt.Println(result.TenSmsResponse)
}
