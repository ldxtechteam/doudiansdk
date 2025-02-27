package sms_send_request

import (
	"encoding/json"
	"github.com/ldxtechteam/doudiansdk/api/sms_send/response"
	"github.com/ldxtechteam/doudiansdk/core"
)

type SmsSendRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SmsSendParam
}

func (c *SmsSendRequest) GetUrlPath() string {
	return "/sms/send"
}

func New() *SmsSendRequest {
	request := &SmsSendRequest{
		Param: &SmsSendParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *SmsSendRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sms_send_response.SmsSendResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sms_send_response.SmsSendResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *SmsSendRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *SmsSendRequest) GetParams() *SmsSendParam {
	return c.Param
}

type SmsSendParam struct {
	// 短信发送渠道，主要做资源隔离
	SmsAccount string `json:"sms_account"`
	// 签名
	Sign string `json:"sign"`
	// 短信模版id
	TemplateId string `json:"template_id"`
	// 短信模板占位符要替换的值
	TemplateParam string `json:"template_param"`
	// 透传字段，回执的时候原样返回给调用方，最大长度512字符
	Tag *string `json:"tag"`
	// 既支持手机号明文，又支持手机号密文。同时传outbound_id和post_tel，以post_tel为准，不能同时为空
	PostTel *string `json:"post_tel"`
	// 用户自定义扩展码，仅当允许自定义扩展码的时候生效
	UserExtCode *string `json:"user_ext_code"`
	// 外呼id，由/member/getOutboundId接口获得
	OutboundId string `json:"outbound_id"`
}
