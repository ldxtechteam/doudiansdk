package sendHome_order_PayMsg_request

import (
	"github.com/ldxtechteam/doudiansdk/utils"
)

type SendHomeOrderPayMsgRequest struct {
	doudian_sdk.BaseDoudianOpSpiRequest
	Param    *SendHomeOrderPayMsgParam
	Response *sendHome_order_PayMsg_response.SendHomeOrderPayMsgResponse
}

func (c *SendHomeOrderPayMsgRequest) GetParamJsonObject() interface{} {
	return c.Param
}

func (c *SendHomeOrderPayMsgRequest) GetResponseObject() interface{} {
	return c.Response
}

func (c *SendHomeOrderPayMsgRequest) Execute() (interface{}, error) {
	return c.GetClient().Request(c)
}

func (c *SendHomeOrderPayMsgRequest) ResponseJson() (string, error) {
	responseObj, err := c.Execute()
	if err != nil {
		return "", err
	}
	return utils.MarshalNoErr(responseObj), nil
}

func New() *SendHomeOrderPayMsgRequest {

	request := new(SendHomeOrderPayMsgRequest)
	request.SetClient(doudian_sdk.DefaultDoudianOpSpiClient)
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetSpiParam(new(doudian_sdk.DoudianOpSpiParam))
	request.Param = new(SendHomeOrderPayMsgParam)
	response := new(sendHome_order_PayMsg_response.SendHomeOrderPayMsgResponse)
	response.Data = new(sendHome_order_PayMsg_response.SendHomeOrderPayMsgData)
	request.Response = response
	return request
}

type SendHomeOrderPayMsgParam struct {
	// 业务参数：内容和[订单详情]接口一致，可看订单详情接口
	MsgBody string `json:"msg_body"`
	// 消息类型1
	SpiType int64 `json:"spi_type"`
	// 消息唯一id
	MsgId string `json:"msg_id"`
}
