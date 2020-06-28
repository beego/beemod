package module

const ConfigPrefix = "beego."

const (
	OssName     = "oss"
	SessionName = "session"
  Oauth2Name  = "oauth2"
)

// order invokers
var OrderInvokers = []invokerAttr{
	{OssName},
	{Oauth2Name},
	{SessionName},
}

type invokerAttr struct {
	Name string
}
