package mail

import (
	"github.com/beego/beemod"
	c "github.com/smartystreets/goconvey/convey"
	"gopkg.in/gomail.v2"
	"testing"
)

const configTpl = `
  [beego.mail.dev]
	debug = false
	mode = "smtp"
	host = "smtp.yeah.net"
	port = 25
	FromEmail = "yangzzzzzz@yeah.net"
	FromPassword = "xxxxxxxxxxxxxxx"
`

const FromEmail = "yangzzzzzz@yeah.net"

func TestMailConfig(t *testing.T) {
	var (
		err    error
		config string
	)
	c.Convey("Define configuration", t, func() {
		config = configTpl
		c.Convey("Parse configuration", func() {
			err = beemod.Register(DefaultBuild).SetCfg([]byte(config), "toml").Run()
			c.So(err, c.ShouldBeNil)
		})
	})
}

func TestMailInit(t *testing.T) {
	var (
		err    error
		config string
	)
	c.Convey("Define configuration", t, func() {
		config = configTpl
		c.Convey("Define configuration", func() {
			err = beemod.Register(DefaultBuild).SetCfg([]byte(config), "toml").Run()
			c.So(err, c.ShouldBeNil)
			c.Convey("Set configuration group (initialization)", func() {
				obj := Invoker("dev")
				c.So(obj, c.ShouldNotBeNil)
			})
		})
	})
}

func TestMailInstance(t *testing.T) {
	var (
		err    error
		obj    *Client
		config string
	)
	c.Convey("Define configuration", t, func() {
		config = configTpl
		c.Convey("Parse configuration", func() {
			err = beemod.Register(DefaultBuild).SetCfg([]byte(config), "toml").Run()
			c.So(err, c.ShouldBeNil)
			c.Convey("Set configuration group (initialization)", func() {
				obj = Invoker("dev")
				c.So(obj, c.ShouldNotBeNil)
				c.Convey("testing method", func() {
					err = obj.Push(form())
					c.So(err, c.ShouldBeNil)
				})
			})
		})
	})
}

func form() *gomail.Message {
	message := gomail.NewMessage()
	message.SetHeader("From", FromEmail)                                       //发件人
	message.SetHeader("To", "2399158611@qq.com")                                 //收件人
//	message.SetAddressHeader("Cc", "yangzzzzzz@yeah.net", "beemod-mail-test") //抄送人
	message.SetHeader("Subject", "Hello!")                                     //邮件标题
	message.SetBody("text/html", "使用beemod-mail测试发送邮件!")                       //邮件内容
	//message.Attach("/Users/ding/Desktop/logo.png")                             //邮件附件
	return message
}
