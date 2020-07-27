package beemod

import (
	"github.com/beego/beemod/pkg/module"
	"reflect"
)

func sortInvokers(invokers []module.InvokerFunc) (invokerSort []module.Invoker, err error) {
	invokerMap := make(map[string]module.Invoker)
	invokerSort = make([]module.Invoker, 0)
	for _, invoker := range invokers {
		obj := invoker()
		name := getCallerName(obj)
		invokerMap[name] = obj
	}
	for _, value := range module.OrderInvokers {
		caller, ok := invokerMap[value.Name]
		if ok {
			// 如果存在于map，加入到排序里的caller sort
			invokerSort = append(invokerSort, caller)
		}
	}
	return
}

func getCallerName(caller module.Invoker) string {
	return reflect.ValueOf(caller).Elem().FieldByName("Name").String()
}
