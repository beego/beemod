// Author: SDing <deen.job@qq.com>
// Date: 2020/6/28 - 3:16 PM

package sms

// example
/**
[beemod.sms.my]
	debug = true
	mode  = "alibaba"
*/
type InvokerCfg struct {
	Debug        bool
	Mode         string
	Area         string
	domain       string
	method       string
	sign         string
	timeOut      int
	AccessKeyId  string
	AccessSecret string
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
	domain:  "",
	method:  "",
	sign:    "",
	timeOut: 5,
}
