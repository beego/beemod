/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/7/31 9:49
 */
package rabbitmq

type InvokerCfg struct {
	Host      string `ini:"host"`

}

/**
type:
	simple
	pubSub
	routing
	topic
*/
var DefaultInvokerCfg = InvokerCfg{
	Host:      "",
}
