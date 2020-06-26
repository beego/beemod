package module

const ConfigPrefix = "beego."

const (
	OssName     = "oss"
	SessionName = "session"
)

// order invokers
var OrderInvokers = []invokerAttr{
	{OssName},
	{SessionName},
}

type invokerAttr struct {
	Name string
}
