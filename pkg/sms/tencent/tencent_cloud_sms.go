// Author: SDing <deen.job@qq.com>
// Date: 2020/6/28 - 4:18 PM

package tencent

import (
	"github.com/beego/beemod/pkg/sms/standard"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711" //import sms
	"strings"
)

// see Tencent SDK API
//https://cloud.tencent.com/document/product/382/43199

// Tencent SDK Api
type Client struct {
	// Credential
	//common.Credential
	// Profile
	//*profile.ClientProfile
	sms *sms.Client
}

// Build Tencent SDK
func New(area, accessKeyId, accessSecret, domain, method, sign string, timeOut int) (standard.Sms, error) {
	credential := common.NewCredential(accessKeyId, accessSecret)
	clientProfile := profile.NewClientProfile()
	clientProfile.HttpProfile.Endpoint = domain
	clientProfile.HttpProfile.ReqMethod = method
	clientProfile.SignMethod = sign
	clientProfile.HttpProfile.ReqTimeout = timeOut
	client, err := sms.NewClient(credential, area, clientProfile)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

// Tencent Send  SMS
func (ten *Client) Send(param map[string]string) (*standard.Response, error) {
	request := sms.NewSendSmsRequest()
	for k, v := range param {
		switch k {
		case "SmsSdkAppid":
			request.SmsSdkAppid = common.StringPtr(v)
		case "Sign":
			request.Sign = common.StringPtr(v)
		case "SenderId":
			request.SenderId = common.StringPtr(v)
		case "SessionContext":
			request.SessionContext = common.StringPtr(v)
		case "ExtendCode":
			request.ExtendCode = common.StringPtr(v)
		case "TemplateParamSet":
			request.TemplateParamSet = common.StringPtrs(strings.Split(v, ","))
		case "TemplateID":
			request.TemplateID = common.StringPtr(v)
		case "PhoneNumberSet":
			request.PhoneNumberSet = common.StringPtrs(strings.Split(v, ","))
		}
	}
	sendSms, err := ten.sms.SendSms(request)
	if err != nil {
		return nil, err
	}
	return standard.RespTen(sendSms), nil
}
