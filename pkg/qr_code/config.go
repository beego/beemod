// Author: SDing <deen.job@qq.com>
// Date: 2020/7/08 - 4:32 PM

package qr_code

type InvokerCfg struct {
	Debug            bool
	Mode             string
	AvatarX, AvatarY int
	Size             int
	Foreground       string
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
