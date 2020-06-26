package module

const ConfigPrefix = "beego."

const OssName = "oss"

// order invokers
var OrderInvokers = []invokerAttr{
	{OssName},
}

type invokerAttr struct {
	Name string
}
