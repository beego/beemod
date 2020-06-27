// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/6/27 - 3:15 下午

package mail

// example
/**
[beemod.oss.mail]
	debug = true
	mode  = "test"
*/
type InvokerMailCfg struct {
	Debug           bool
	Mode            string
	Addr            string
	AccessKeyID     string
	AccessKeySecret string
	CdnName         string
	OssBucket       string
	FileBucket      string
	IsDeleteSrcPath bool
}

var DefaultInvokerMailCfg = InvokerMailCfg{
	Debug:           false,
	Mode:            "file",
	Addr:            "",
	AccessKeyID:     "",
	AccessKeySecret: "",
	CdnName:         "",
	OssBucket:       "",
	FileBucket:      "",
	IsDeleteSrcPath: false,
}
