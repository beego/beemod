package module

const ConfigPrefix = "beego."

const OssName = "oss"

const Oauth2Name = "oauth2"

// order invokers
var OrderInvokers = []invokerAttr{
	{OssName},
	{Oauth2Name},
}

type invokerAttr struct {
	Name string
}
