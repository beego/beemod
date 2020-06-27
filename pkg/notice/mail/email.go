// Author: SDing <deen.job@qq.com>
// Date: 2020/6/27 - 1:28 下午

package mail

import (
	"errors"
	"gopkg.in/gomail.v2"
	"strconv"
	"strings"
)

type EMailServe struct{}

type EMail struct {
	// 接收邮件用户 可以多个邮箱地址
	Users []string

	// 邮件主题
	Subject string
	// 邮件内容
	Body   string
	Dialer *gomail.Dialer
}

func New() *EMailServe {
	return &EMailServe{}
}
func (ms EMailServe) BuildDialer(params map[string]string) (*gomail.Dialer, error) {
	for _, v := range params {
		if len(v) == 0 {
			return nil, errors.New("email serve config set failed check param is null")
		}
	}
	port, err := strconv.Atoi(params["ServerPort"])
	if err != nil {
		return nil, errors.New("init email server port error")
	}
	return gomail.NewDialer(params["ServerHost"], port, params["FromEmail"], params["FromPassword"]), nil
}

func NewMail(usersEmailAddr, subject, body string, dialer *gomail.Dialer) (*EMail, error) {
	if len(usersEmailAddr) <= 0 {
		return nil, errors.New("email addr is null")
	}

	if len(subject) <= 0 || len(body) <= 0 {
		return nil, errors.New("email subject or body is null")
	}
	return &EMail{
		Users:   strings.Split(usersEmailAddr, ","),
		Subject: subject,
		Body:    body,
		Dialer:  dialer,
	}, nil
}

func (ms *EMail) Send() error {
	m := gomail.NewMessage()
	m.SetHeader("From", ms.Dialer.Username) //发件人
	m.SetHeader("To", ms.Users...)          //收件人
	m.SetHeader("Subject", ms.Subject)      //邮件标题
	m.SetBody("text/html", ms.Body)         //邮件内容
	err := ms.Dialer.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}
