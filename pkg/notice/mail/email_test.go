// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/6/27 - 2:48 下午

package mail

import (
	"testing"
)

func TestEMailServe_BuildDialer(t *testing.T) {
	es := New()
	parmas := make(map[string]string, 5)
	parmas["ServerHost"] = "smtp.qq.com"
	parmas["ServerPort"] = "25"
	parmas["FromEmail"] = "1423119397@qq.com"
	parmas["FromPassword"] = "xxxx"
	d, err := es.BuildDialer(parmas)
	if err != nil {
		t.Fatal(err)
	}
	mail, err := NewMail("deen.job@qq.com", "测试邮件", "邮件内容测试", d)
	if err != nil {
		t.Fatal(err)
	}
	err = mail.Send()
	if err != nil {
		t.Log(err)
	}
}
