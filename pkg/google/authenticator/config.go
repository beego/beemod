/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/8/6 18:39
 */
package authenticator

type InvokerCfg struct {
	User string `ini:"user"`
	Iss  string `ini:"iss"`
}

var DefaultInvokerCfg = InvokerCfg{
	User: "",
	Iss:  "",
}
