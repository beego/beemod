package session

import (
	"github.com/astaxie/beego/session"
	"github.com/beego/beemod"
	c "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const configTpl = `
	[beego.session.mysession]
		mode = "cookie"
		debug = true
	[beego.session.mysession.mangerCfg]
		cookieName = "gosessionid"
		gclifetime = 3600
		enableSetCookie = true
		providerConfig= "{\"cookieName\":\"gosessionid\",\"securityKey\":\"beegocookiehashkey\"}"
`

func TestSessionConfig(t *testing.T) {
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

func TestSessionInit(t *testing.T) {
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
				obj := Invoker("mysession")
				c.So(obj, c.ShouldNotBeNil)
			})
		})
	})
}

func TestSessionInstance(t *testing.T) {
	var (
		err    error
		obj    *session.Manager
		config string
	)
	c.Convey("Define configuration", t, func() {
		config = configTpl
		c.Convey("Parse configuration", func() {
			err = beemod.Register(DefaultBuild).SetCfg([]byte(config), "toml").Run()
			c.So(err, c.ShouldBeNil)
			c.Convey("Set configuration group (initialization)", func() {
				obj = Invoker("mysession")
				c.So(obj, c.ShouldNotBeNil)
				c.Convey("testing method", func() {
					go obj.GC()
					r, _ := http.NewRequest("GET", "/", nil)
					w := httptest.NewRecorder()
					sess, err := obj.SessionStart(w, r)
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
					c.So(err, c.ShouldBeNil)
				})
			})
		})
	})
}
