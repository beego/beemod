package module

const ConfigPrefix = "beego."

const OssName = "oss"
const CacheName = "cache"

// order invokers
var OrderInvokers = []invokerAttr{
	{OssName},
	{CacheName},
}

type invokerAttr struct {
	Name string
}
