package module

const ConfigPrefix = "beego."

const (
	OssName     = "oss"
	SessionName = "session"
	Oauth2Name  = "oauth2"
	SmsName     = "sms"
	MailName    = "mail"
	DingName    = "ding"
	CacheName   = "cache"
	QrcodeName  = "qrcode"
	TokenName   = "token"
)

// order invokers
var OrderInvokers = []invokerAttr{
	{OssName},
	{CacheName},
	{DingName},
	{Oauth2Name},
	{SessionName},
	{SmsName},
	{MailName},
	{QrcodeName},
	{TokenName},
}

type invokerAttr struct {
	Name string
}
