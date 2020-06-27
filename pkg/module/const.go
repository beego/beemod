package module

const ConfigPrefix = "beego."

const OssName = "oss"
const DingName = "ding"

// order invokers
var OrderInvokers = []invokerAttr{
	{OssName},
  {DingName},
}

type invokerAttr struct {
	Name string
}
