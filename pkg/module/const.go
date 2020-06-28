package module

const (
	ConfigPrefix = "beego."
	OssName      = "oss"
	MailName     = "mail"
)

// order invokers
var OrderInvokers = []invokerAttr{
	{OssName},
	{Name: MailName},
}

type invokerAttr struct {
	Name string
}
