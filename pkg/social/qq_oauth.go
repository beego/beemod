package social

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type QQClient struct {
	AppID       string
	AppKey      string
	RedirectURI string
	Name        string
	OpenID      string
}
type QQ_Error struct {
	Error             int    `json:"error"`
	Error_description string `json:"error_description"`
}
type AuthInfo struct {
	ClientID string `json:"client_id"`
	OpenID   string `json:"openid"`
	UnionID  string `json:"unionid"`
}

// UserInfo get_user_info api response
type QQUserInfo struct {
	Ret             int64  `json:"ret"`                // 返回代码
	Msg             string `json:"msg"`                // 错误信息
	Nickname        string `json:"nickname"`           // QQ空间的昵称
	FigType         string `json:"figureurl_type"`     // 头像类型
	Fig             string `json:"figureurl"`          // 30×30 像素空间头像
	Fig1            string `json:"figureurl_1"`        // 50×50 像素空间头像
	Fig2            string `json:"figureurl_2"`        // 100×100 像素空间头像
	FigQQ           string `json:"figureurl_qq"`       // 640×640 QQ头像
	FigQQ1          string `json:"figureurl_qq_1"`     // 40×40 QQ头像
	FigQQ2          string `json:"figureurl_qq_2"`     // 100×100 QQ头像（可能为空）
	Gender          string `json:"gender"`             // 性别（未设置则返回“男”）
	GenderType      int64  `json:"gender_type"`        // 性别类型
	Province        string `json:"province"`           // 省份
	City            string `json:"city"`               // 城市
	Year            string `json:"year"`               // 出生年份
	Constellation   string `json:"constellation"`      // 星座
	IsYellowVIP     string `json:"is_yellow_vip"`      // 是否为黄钻用户
	IsYellowYearVIP string `json:"is_yellow_year_vip"` // 是否年费黄钻用户
	YellowVIPLevel  string `json:"yellow_vip_level"`   // 黄钻登机
	VIP             string `json:"vip"`                // 是非VIP
	Level           string `json:"level"`              // VIP等级
	IsLost          int64  `json:"is_lost"`            // 未知
}

func errorUnmarshal(b []byte, res interface{}) error {
	return json.Unmarshal(b[9:len(b)-3], &res)

}

//get login page
func (c *QQClient) LoginPage(state string) string {
	return "https://graph.qq.com/oauth2.0/authorize?response_type=code&client_id=" + c.AppID + "&redirect_uri=" + c.RedirectURI + "&state=" + state
}

//get token
func (c *QQClient) GetAccessToken(code string) (*BasicTokenInfo, error) {

	params := make(url.Values)
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", c.AppID)
	params.Add("client_secret", c.AppKey)
	params.Add("code", code)
	params.Add("redirect_uri", c.RedirectURI)
	fmt.Printf("%s\n\n", "https://graph.qq.com/oauth2.0/token?"+params.Encode())

	resp, err := http.Get("https://graph.qq.com/oauth2.0/token?" + params.Encode())
	if err != nil {
		return nil, err
	}
	var b []byte
	if b, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	urls, err := url.ParseQuery(string(b))
	if err != nil {
		return nil, err

	}

	token := urls.Get("access_token")
	if token == "" {
		e := &errMsg{}
		if err = errorUnmarshal(b, e); err != nil {
			return nil, err
		}
		err = errors.New(fmt.Sprintf("Error code %d: %s", e.Error, e.ErrorMsg))
	}
	basicToken := &BasicTokenInfo{}
	basicToken.AccessToken = token
	return basicToken, nil
}

//get OpenID
func (c *QQClient) GetOpenID(accessToken string) (*AuthInfo, error) {
	resp, err := http.Get(fmt.Sprintf("https://graph.qq.com/oauth2.0/me?access_token=%s&unionid=1", accessToken))

	var b []byte
	if b, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	l := len(b)

	res := &AuthInfo{}
	err = json.Unmarshal(b[9:l-3], res)
	if err != nil {
		return nil, err
	}

	if res.OpenID == "" {
		e := &errMsg{}
		if err = errorUnmarshal(b, e); err != nil {
			return nil, err
		}
		err = errors.New(fmt.Sprintf("Error code %d: %s", e.Error, e.ErrorMsg))
		return nil, err
	}
	c.OpenID = res.OpenID
	return res, nil
}

//get userInfo
func (c *QQClient) GetUserInfo(accessToken string) (*BasicUserInfo, error) {
	resp, err := http.Get(fmt.Sprintf("https://graph.qq.com/user/get_user_info?access_token=%s&oauth_consumer_key=%s&openid=%s", accessToken, c.AppID, c.OpenID))

	var b []byte
	if b, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}
	res := &QQUserInfo{}
	if err = json.Unmarshal(b, &res); err != nil {
		return nil, err
	}

	if res.Ret != 0 {
		err = fmt.Errorf("Error code %d: %s ", res.Ret, res.Msg)
		return nil, err

	}
	basicUser := &BasicUserInfo{
		NickName: res.Nickname,
		HeadIcon: res.Fig,
	}
	return basicUser, nil
}
func (c *QQClient) GetType() string {
	return c.Name
}
func NewQQauth2Service(appID, appKey, redirectURI string) SocialService {
	return &QQClient{
		AppID:       appID,
		AppKey:      appKey,
		RedirectURI: redirectURI,
		Name:        "qq",
	}
}
