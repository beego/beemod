// Author: SDing <deen.job@qq.com>
// Date: 2020/6/28 - 5:40 PM

package main

import (
	"fmt"
	"github.com/beego/beemod"
	"github.com/beego/beemod/pkg/sms"
)

// custom config toml template
var config = `
  [beego.sms.ali]
	debug = false
	mode = "alibaba"
	area = "ap-guangzhou"
	accessKeyId = ""
	accessSecret = ""
`

func main() {
	err := beemod.Register(
		sms.DefaultBuild,
	).SetCfg([]byte(config), "toml").Run()
	if err != nil {
		panic("register err:" + err.Error())
	}
	client := sms.Invoker("my")
	// send alibaba sms
	// see aliYun sdk api
	// https://api.aliyun.com/new#/?product=Dysmsapi&version=2017-05-25&api=SendSms&params={%22RegionId%22:%22cn-hangzhou%22,%22PhoneNumbers%22:%2219999999999%22,%22SignName%22:%22SignName%22,%22TemplateCode%22:%2210001%22,%22TemplateParam%22:%22json%20format%22,%22SmsUpExtendCode%22:%22xxxx%22,%22OutId%22:%22xxxx%22}&tab=DEMO&lang=GO
	param := map[string]string{
		"Scheme":          "https",
		"PhoneNumbers":    "199999999999",
		"SignName":        "you_sign",
		"TemplateCode":    "tlp_code",
		"TemplateParam":   "tlp_param",
		"SmsUpExtendCode": "0000",
		"OutId":           "0000",
	}
	fmt.Println(param)
	result, err := client.Push(param)
	if err != nil {
		fmt.Println(err)
	}
	// response
	fmt.Println(result.AliSmsResponse)
}
