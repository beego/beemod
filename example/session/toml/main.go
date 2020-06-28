package main

import (
	"github.com/beego-dev/beemod"
	modSession "github.com/beego-dev/beemod/pkg/session"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
)

var config = `
	[beego.session.mysession]
		mode = "memory"
		debug = true
	[beego.session.mysession.mangerCfg]
		cookieName = "gosessionid"
		gclifetime = 10
		enableSetCookie = true
`

func main() {
	err := beemod.Register(
		modSession.DefaultBuild,
	).SetCfg([]byte(config), "toml").Run()
	if err != nil {
		panic("register err:" + err.Error())
	}

	globalSessions := modSession.Invoker("mysession")
	go globalSessions.GC()

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	sess, err := globalSessions.SessionStart(w, r)
	if err != nil {
		log.Fatal("set error,", err)
	}
	defer sess.SessionRelease(w)
	err = sess.Set("username", "astaxie")
	if err != nil {
		log.Fatal("set error,", err)
	}
	if username := sess.Get("username"); username != "astaxie" {
		log.Fatal("get username error")
	}
	if cookiestr := w.Header().Get("Set-Cookie"); cookiestr == "" {
		log.Fatal("setcookie error")
	} else {
		parts := strings.Split(strings.TrimSpace(cookiestr), ";")
		for k, v := range parts {
			nameval := strings.Split(v, "=")
			if k == 0 && nameval[0] != "gosessionid" {
				log.Fatal("error")
			}
		}
	}
}
