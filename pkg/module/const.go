package module

const (
	ConfigPrefix   = "beego."
	OssName        = "oss"
	NoticeMailName = "notice.MAIL"
)

// order invokers
var OrderInvokers = []invokerAttr{
	{OssName},
	{Name: NoticeMailName},
}

type invokerAttr struct {
	Name string
}
