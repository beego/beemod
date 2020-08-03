package main

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beemod"
	"github.com/beego/beemod/pkg/oauth"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

var config = `
[beego.oauth2.dev]
	state = "beego"
	appID = "sfdhjaksfddsajks"
	appSecret = "mnjUYj8rlXXKGS2RNgsdad7lygWrjJzjD5"
	authURL = "http://oauthadmin.yitum.com/user/login"
	tokenURL = "http://oauthadmin.yitum.com/api/v1/oauth/token"
	redirectURI = "http://localhost:8000/api/code"
	userInfoURL = "http://oauthadmin.yitum.com/api/v1/oauth/user"
	Scopes=[]
`

type UserInfo struct {
	Uid         int    ` json:"uid" `         // uid
	Nickname    string `json:"nickname" `     // nickname
	Email       string ` json:"email" `       // email
	Avatar      string ` json:"avatar"`       // avatar
	Password    string ` json:"password" `    // password
	Status      int64  ` json:"status" `      // status
	Gender      int64  ` json:"gender"`       // gender
	Birthday    int64  ` json:"birthday" `    // birthday
	LastLoginIp string ` json:"lastLoginIp" ` // last_login_ip
}

type RespUser struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data UserInfo `json:"data"`
}

func main() {
	err := beemod.Register(
		oauth.DefaultBuild,
	).SetCfg([]byte(config), "toml").Run()
	if err != nil {
		panic("register err:" + err.Error())
	}
	client := oauth.Invoker("dev")

	http.HandleFunc("/api/login", func(writer http.ResponseWriter, request *http.Request) {
		page := client.LoginPage()
		http.Redirect(writer, request, page, http.StatusFound)
	})

	http.HandleFunc("/api/code", func(writer http.ResponseWriter, request *http.Request) {
		query := request.URL.Query()
		code := query["code"][0]
		state := query["state"][0]
		token, _ := client.GetAccessToken(state, code)
		var user RespUser
		rsp, _ := client.GetUserInfo(token, user)
		if err := mapstructure.Decode(rsp, &user); err != nil {
			fmt.Println(err)
		}
		writer.Header().Set("content-type", "text/json")
		msg, _ := json.Marshal(user.Data)
		_, _ = writer.Write(msg)
	})

	_ = http.ListenAndServe(":8000", nil)
}
