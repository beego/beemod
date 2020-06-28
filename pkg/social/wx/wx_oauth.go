package wx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego-dev/beemod/pkg/social"
	"net/http"
	"time"
)

type Client struct {
	AppID       string
	AppSecret   string
	AccessToken string
	OpenID      string
	Name        string
}

type wxGetTokenResponse struct {
	AccessToken string `json:"access_token"`
	Openid      string `json:"openid"`
	ErrorCode   int    `json:"errcode"`
	ErrorMsg    string `json:"errmsg"`
}

type wxGetWxInfoResponse struct {
	Openid     string `json:"openid"`
	Nickname   string `json:"nickname"`
	HeadImgurl string `json:"headimgurl"`
	Unionid    string `json:"unionid"`
	ErrorCode  int    `json:"errcode"`
	ErrorMsg   string `json:"errmsg"`
}

func NewOauth2Service(app_id, app_secret string) *Client {
	return &Client{
		AppID:     app_id,
		AppSecret: app_secret,
		Name:      "wx",
	}
}

func (c *Client) GetAccessToken(code string) (*social.BasicTokenInfo, error) {
	var openIdUrl = "https://api.weixin.qq.com/sns/oauth2/access_token"
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, openIdUrl, nil); err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("code", code)
	q.Add("appid", c.AppID)
	q.Add("secret", c.AppSecret)
	q.Add("grant_type", "authorization_code")
	req.URL.RawQuery = q.Encode()
	// 发送请求并获取响应
	var client = http.Client{
		Timeout: 10 * time.Second,
	}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var ret wxGetTokenResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	if ret.ErrorCode != 0 {
		return nil, errors.New(fmt.Sprintf("GetWxTokenByCode error = %d", ret.ErrorCode))
	}
	basicToken := social.BasicTokenInfo{}
	basicToken.AccessToken = ret.AccessToken
	return &basicToken, nil
}

func (c *Client) GetUserInfo(accessToken string) (*social.BasicUserInfo, error) {
	//return &wxGetWxInfoResponse{Unionid: "wxunionid2", Nickname: "mancangluo2", HeadImgurl: "icon2"}, nil
	// 形成请求
	var openIdUrl = "https://api.weixin.qq.com/sns/userinfo"
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, openIdUrl, nil); err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("access_token", accessToken)
	q.Add("openid", c.OpenID)
	req.URL.RawQuery = q.Encode()
	// 发送请求并获取响应
	var client = http.Client{
		Timeout: 10 * time.Second,
	}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var ret wxGetWxInfoResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}

	if ret.ErrorCode != 0 {
		return nil, errors.New(fmt.Sprintf("GetWxInfoByToken error = %d", ret.ErrorCode))
	}

	basicUser := social.BasicUserInfo{
		NickName: ret.Nickname,
		HeadIcon: ret.HeadImgurl,
	}

	return &basicUser, nil
}

func (c *Client) GetType() string {
	return c.Name
}
