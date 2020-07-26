// TODO: mail-test-example
// Author: SDing <deen.job@qq.com>
// Date: 2020/6/28 - 1:06 PM

package main

import (
	"fmt"
	"github.com/beego/beemod"
	"github.com/beego/beemod/pkg/mail"
	"gopkg.in/gomail.v2"
)

// custom config toml template
var config = `
  [beego.mail.my]
	debug = false
	mode = "smtp"
	host = "smtp.qq.com"
	port = 25
	FromEmail = "1423119397@qq.com"
	FromPassword = "If you would like to test this component please use your configuration, thank you"
`

// Your From Email
const FromEmail = "1423119397@qq.com"

func main() {
	err := beemod.Register(
		mail.DefaultBuild,
	).SetCfg([]byte(config), "toml").Run()
	if err != nil {
		panic("register err:" + err.Error())
	}
	client := mail.Invoker("my")
	// form() build email message & push message
	err = client.Push(form())
	if err != nil {
		fmt.Println(err)
	}
}

func form() *gomail.Message {
	message := gomail.NewMessage()
	message.SetHeader("From", FromEmail)                                       //发件人
	message.SetHeader("To", "deen.job@qq.com")                                 //收件人
	message.SetAddressHeader("Cc", "coding1618@gmail.com", "beemod-mail-test") //抄送人
	message.SetHeader("Subject", "Hello!")                                     //邮件标题
	message.SetBody("text/html", "使用beemod-mail测试发送邮件!")                       //邮件内容
	message.Attach("/Users/ding/Desktop/logo.png")                             //邮件附件
	return message
}
