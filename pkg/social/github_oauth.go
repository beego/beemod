package social

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type GithubClient struct {
	AppID       string
	AppSecret   string
	AccessToken string
	OpenID      string
	Name        string
}

type githubGetTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type githubGetInfoResponse struct {
	Nickname   string `json:"login"`
	HeadImgurl string `json:"avatar_url"`
}

func NewGithubOauth2Service(app_id, app_secret string) SocialService {
	return &GithubClient{
		AppID:     app_id,
		AppSecret: app_secret,
		Name:      "github",
	}
}

func (c *GithubClient) GetAccessToken(code string) (*BasicTokenInfo, error) {
	var openIdUrl = "https://github.com/login/oauth/access_token"
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, openIdUrl, nil); err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	q := req.URL.Query()
	q.Add("code", code)
	q.Add("client_id", c.AppID)
	q.Add("client_secret", c.AppSecret)
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

	var ret githubGetTokenResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	if ret.AccessToken == "" {
		return nil, errors.New(fmt.Sprintf("GetAccessToken error = %s", ret.AccessToken))
	}
	basicToken := BasicTokenInfo{}
	basicToken.AccessToken = ret.AccessToken
	return &basicToken, nil
}

func (c *GithubClient) GetUserInfo(accessToken string) (*BasicUserInfo, error) {
	// 形成请求
	var openIdUrl = "https://api.weixin.qq.com/sns/userinfo"
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, openIdUrl, nil); err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "token "+accessToken)
	// 发送请求并获取响应
	var client = http.Client{
		Timeout: 10 * time.Second,
	}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var ret githubGetInfoResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}

	if ret.Nickname == "" {
		return nil, errors.New(fmt.Sprintf("GetUserInfo error = %s", ret.Nickname))
	}

	basicUser := BasicUserInfo{
		NickName: ret.Nickname,
		HeadIcon: ret.HeadImgurl,
	}

	return &basicUser, nil
}

func (c *GithubClient) GetType() string {
	return c.Name
}
