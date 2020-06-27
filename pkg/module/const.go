package module

const (
	ConfigPrefix = "beego."
	OssName      = "oss"
	NoticeName   = "notice"
)

// order invokers
var OrderInvokers = []invokerAttr{
	{OssName},
	{Name: NoticeName},
}

type invokerAttr struct {
	Name string
}
