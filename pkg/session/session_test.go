package session

import (
	"github.com/beego-dev/beemod"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var config = `
	[beego.session.mysession]
		mode = "cookie"
		debug = true
	[beego.session.mysession.mangerCfg]
		cookieName = "gosessionid"
		gclifetime = 3600
		enableSetCookie = true
		providerConfig= "{\"cookieName\":\"gosessionid\",\"securityKey\":\"beegocookiehashkey\"}"
`

func TestCookie(t *testing.T) {
	err := beemod.Register(
		DefaultBuild,
	).SetCfg([]byte(config), "toml").Run()
	if err != nil {
		panic("register err:" + err.Error())
	}

	globalSessions := Invoker("mysession")
	go globalSessions.GC()

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	sess, err := globalSessions.SessionStart(w, r)
	if err != nil {
		t.Fatal("set error,", err)
	}
	defer sess.SessionRelease(w)
	err = sess.Set("username", "astaxie")
	if err != nil {
		t.Fatal("set error,", err)
	}
	if username := sess.Get("username"); username != "astaxie" {
		t.Fatal("get username error")
	}
	if cookiestr := w.Header().Get("Set-Cookie"); cookiestr == "" {
		t.Fatal("setcookie error")
	} else {
		parts := strings.Split(strings.TrimSpace(cookiestr), ";")
		for k, v := range parts {
			nameval := strings.Split(v, "=")
			if k == 0 && nameval[0] != "gosessionid" {
				t.Fatal("error")
			}
		}
	}
}
