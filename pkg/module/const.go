package module

const ConfigPrefix = "beego."

const (
	OssName     = "oss"
	SessionName = "session"
	Oauth2Name  = "oauth2"
	SmsName     = "sms"
	MailName    = "mail"
  DingName    = "ding"
)


// order invokers
var OrderInvokers = []invokerAttr{
	{OssName},
  {DingName},
	{Oauth2Name},
	{SessionName},
	{SmsName},
	{MailName},
}

type invokerAttr struct {
	Name string
}
