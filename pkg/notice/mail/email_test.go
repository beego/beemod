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
	parmas["ServerHost"] = "xxxxxxx"
	parmas["ServerPort"] = "xxxxxxx"
	parmas["FromEmail"] = "xxxxxxx"
	parmas["FromPassword"] = "xxxxxxx"
	d, err := es.BuildDialer(parmas)
	if err != nil {
		t.Fatal(err)
	}
	d.DialAndSend()
}
