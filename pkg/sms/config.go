// Author: SDing <deen.job@qq.com>
// Date: 2020/6/28 - 3:16 PM

package sms

type InvokerCfg struct {
	Debug        bool 	`ini:"debug"`
	Mode         string `ini:"mode"`
	Area         string `ini:"area"`
	Domain       string `ini:"domain"`
	Method       string `ini:"method"`
	Sign         string `ini:"sign"`
	TimeOut      int `ini:"time_out"`
	AccessKeyId  string `ini:"accessKeyId"`
	AccessSecret string `ini:"accessSecret"`
}

var DefaultInvokerCfg = InvokerCfg{
	Debug:        true,
	Mode:         "alibaba",
	Area:         "ap-guangzhou",
	AccessKeyId:  "",
	AccessSecret: "",
	// 下面配置是腾讯云特有  使用阿里云时可以不用填写
	// The following configuration is specific to Tencent Cloud when using
	// Ali cloud can be filled out
	Domain:  "",
	Method:  "",
	Sign:    "",
	TimeOut: 5,
}
