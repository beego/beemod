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
	RedisName   = "redis"
	QrcodeName  = "qrcode"
	TokenName   = "token"
	MysqlName   = "mysql"
	LogName     = "logger"
	MongoName   = "mongo"
	RabbitmqName   = "rabbitmq"
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
	{MysqlName},
	{LogName},
	{RedisName},
	{MongoName},
	{RabbitmqName},
}

type invokerAttr struct {
	Name string
}
