// Author: SDing <deen.job@qq.com>
// Date: 2020/6/27 - 9:15 下午

package mail

import (
	"github.com/beego-dev/beemod/pkg/notice"
	"gopkg.in/gomail.v2"
	"testing"
)

func TestDialer_Push(t *testing.T) {
	message := MAIL()
	dialer := New(message)
	// 用户发送其他通知也一样使用这种 notice.Push
	err := notice.Push(notice.EMAIL, dialer)
	if err != nil {
		t.Fatal(err)
	}
}

func MAIL() *gomail.Message {
	message := gomail.NewMessage()
	message.SetHeader("From", Config.FromEmail)                                //发件人
	message.SetHeader("To", "deen.job@qq.com")                                 //收件人
	message.SetAddressHeader("Cc", "coding1618@gmail.com", "beemod-mail-test") //抄送人
	message.SetHeader("Subject", "Hello!")                                     //邮件标题
	message.SetBody("text/html", "使用beemod-mail测试发送邮件!")                       //邮件内容
	message.Attach("./logo.png")                                               //邮件附件
	return message
}
