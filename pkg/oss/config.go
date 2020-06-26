package oss

// example
/**
[beemod.oss.myoss]
	debug = true
	mode  = "file"
*/
type InvokerCfg struct {
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

var DefaultInvokerCfg = InvokerCfg{
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
