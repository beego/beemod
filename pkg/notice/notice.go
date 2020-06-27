// Author: SDing <deen.job@qq.com>
// Date: 2020-06-27 20:02:08
// Notification system module Interface design v1.2 Design: SDing
package notice

import "github.com/astaxie/beego/logs"

// NoticeType of notification message
type NoticeType int

const (
	_ NoticeType = iota
	SMS
	DING_TALK
	EMAIL
)

// Notification is Other message parsing interface
type Notification interface {
	Parse() func()
}

// PushAction of message push
type PushAction interface{}

func Push(nt NoticeType, n Notification) error {
	switch nt {
	case EMAIL:
		logs.Info("Is mail")
		n.Parse()()
	default:
		logs.Info("Other types")
	}
	return nil
}
