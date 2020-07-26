package cache

type InvokerCfg struct {
	Key      string `ini:"key"json:"key"`
	Conn     string `ini:"conn"json:"conn"`
	Password string `ini:"password"json:"password"`
}

var DefaultInvokerCfg = InvokerCfg{
	Key:      "",
	Conn:     "",
	Password: "",
}
