// SMS implement standard
// Author: SDing <deen.job@qq.com>
// Date: 2020/6/28 - 4:13 PM

package standard

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711" //import sms
)

// Custom sms interface
type Sms interface {
	Send(map[string]string) (*Response, error)
}

// Unified SMS Response
type Response struct {
	AliSmsResponse *dysmsapi.SendSmsResponse
	TenSmsResponse *sms.SendSmsResponse
}

// Set AliYun SMS Response
func RespAli(ali *dysmsapi.SendSmsResponse) *Response {
	return &Response{AliSmsResponse: ali}
}

// Set Tencent SMS Response
func RespTen(ten *sms.SendSmsResponse) *Response {
	return &Response{TenSmsResponse: ten}
}
