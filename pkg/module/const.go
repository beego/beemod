package module

const ConfigPrefix = "beego."

const (
	OssName     = "oss"
	SessionName = "session"
  Oauth2Name  = "oauth2"
  DingName    = "ding"
)


// order invokers
var OrderInvokers = []invokerAttr{
	{OssName},
  {DingName},
	{Oauth2Name},
	{SessionName},
}

type invokerAttr struct {
	Name string
}
