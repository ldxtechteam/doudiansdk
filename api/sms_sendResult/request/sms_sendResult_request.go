package sms_sendResult_request

import (
	"encoding/json"
	"github.com/ldxtechteam/doudiansdk/api/sms_sendResult/response"
	"github.com/ldxtechteam/doudiansdk/core"
)

type SmsSendResultRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SmsSendResultParam
}

func (c *SmsSendResultRequest) GetUrlPath() string {
	return "/sms/sendResult"
}

func New() *SmsSendResultRequest {
	request := &SmsSendResultRequest{
		Param: &SmsSendResultParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *SmsSendResultRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sms_sendResult_response.SmsSendResultResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sms_sendResult_response.SmsSendResultResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *SmsSendResultRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *SmsSendResultRequest) GetParams() *SmsSendResultParam {
	return c.Param
}

type SmsSendResultParam struct {
	// 开始时间-时间戳，单位秒
	FromTime int64 `json:"from_time"`
	// 结束时间-时间戳，单位秒
	ToTime int64 `json:"to_time"`
	// 短信发送渠道，主要做资源隔离
	SmsAccount string `json:"sms_account"`
	// 模版id
	TemplateId *string `json:"template_id"`
	// 发送状态： 未回执：1 发送失败：2 发送成功：3
	Status *int64 `json:"status"`
	// 查询结果大小，默认是10
	Size *int64 `json:"size"`
	// 查询结果页数，从0开始，默认是0
	Page *int64 `json:"page"`
	// 签名内容
	Sign *string `json:"sign"`
	// 既支持明文，又支持密文
	PostTel *string `json:"post_tel"`
	// 消息的唯一标识，可以用于查询短信到达等
	MessageId *string `json:"message_id"`
	// 查询短信类型，默认是查普通文本短信：0是查询所有类型短信，1是查询普通文本短信，2是查询视频短信
	TplType int64 `json:"tpl_type"`
}
