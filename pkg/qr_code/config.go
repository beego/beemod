// Author: SDing <deen.job@qq.com>
// Date: 2020/7/08 - 4:32 PM

package qr_code

type InvokerCfg struct {
	Debug      bool   `ini:"debug"`
	Mode       string `ini:"mode"`
	AvatarX    int    `ini:"avatarX"`
	AvatarY    int    `ini:"avatarY"`
	Size       int    `ini:"size"`
	Foreground string `ini:"foreground"`
}

var DefaultInvokerCfg = InvokerCfg{
	Debug:   true,
	Mode:    "QrCode",
	AvatarX: 40,
	AvatarY: 40,
	Size:    150,
	//Logo:       "./gopher-500.png",
	Foreground: "",
}
