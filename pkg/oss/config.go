package oss

type InvokerCfg struct {
	Debug           bool   `ini:"debug"`
	Mode            string `ini:"mode"`
	Addr            string `ini:"addr"`
	AccessKeyID     string `ini:"accessKeyId"`
	AccessKeySecret string `ini:"accessKeySecret"`
	CdnName         string `ini:"cdnName"`
	OssBucket       string `ini:"ossBucket"`
	FileBucket      string `ini:"fileBucket"`
	IsDeleteSrcPath bool   `ini:"isDeleteSrcPath"`
}

var DefaultInvokerCfg = InvokerCfg{
	Debug:           false,
	Mode:            "file",
	Addr:            "",
	AccessKeyID:     "",
	AccessKeySecret: "",
	CdnName:         "",
	OssBucket:       "",
	FileBucket:      ".",
	IsDeleteSrcPath: false,
}
