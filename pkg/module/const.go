package module

const ConfigPrefix = "beego."

const (
	OssName     = "oss"
	SessionName = "session"
	Oauth2Name  = "oauth2"
	SmsName     = "sms"
	MailName    = "mail"
)

// order invokers
var OrderInvokers = []invokerAttr{
	{OssName},
	{Oauth2Name},
	{SessionName},
	{SmsName},
	{MailName},
}

type invokerAttr struct {
	Name string
}
