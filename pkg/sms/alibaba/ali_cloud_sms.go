// Author: SDing <deen.job@qq.com>
// Date: 2020/6/28 - 4:17 PM

package alibaba

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/beego/beemod/pkg/sms/standard"
)

// AliYun SDK
type Client struct {
	Client *dysmsapi.Client
	//Request *dysmsapi.SendSmsRequest
}

// New AliYun SMS Client
func New(area, accessKeyId, accessSecret string) (standard.Sms, error) {
	client, err := dysmsapi.NewClientWithAccessKey(area, accessKeyId, accessSecret)
	if err != nil {
		return nil, err
	}
	return &Client{Client: client}, nil
}

// AliYun SMS Send
func (ali *Client) Send(param map[string]string) (*standard.Response, error) {
	request := dysmsapi.CreateSendSmsRequest()
	for k, v := range param {
		switch k {
		case "OutId":
			request.OutId = v
		case "Scheme":
			request.Scheme = v
		case "SignName":
			request.SignName = v
		case "PhoneNumbers":
			request.PhoneNumbers = v
		case "TemplateCode":
			request.TemplateCode = v
		case "TemplateParam":
			request.TemplateParam = v
		case "SmsUpExtendCode":
			request.SmsUpExtendCode = v
		}
	}
	sms, err := ali.Client.SendSms(request)
	if err != nil {
		return nil, err
	}
	// unified return result
	return standard.RespAli(sms), nil
}
