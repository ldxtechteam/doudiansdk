package sendHome_order_finishMsg_request

import (
	"github.com/ldxtechteam/doudiansdk/utils"
)

type SendHomeOrderFinishMsgRequest struct {
	doudian_sdk.BaseDoudianOpSpiRequest
	Param    *SendHomeOrderFinishMsgParam
	Response *sendHome_order_finishMsg_response.SendHomeOrderFinishMsgResponse
}

func (c *SendHomeOrderFinishMsgRequest) GetParamJsonObject() interface{} {
	return c.Param
}

func (c *SendHomeOrderFinishMsgRequest) GetResponseObject() interface{} {
	return c.Response
}

func (c *SendHomeOrderFinishMsgRequest) Execute() (interface{}, error) {
	return c.GetClient().Request(c)
}

func (c *SendHomeOrderFinishMsgRequest) ResponseJson() (string, error) {
	responseObj, err := c.Execute()
	if err != nil {
		return "", err
	}
	return utils.MarshalNoErr(responseObj), nil
}

func New() *SendHomeOrderFinishMsgRequest {

	request := new(SendHomeOrderFinishMsgRequest)
	request.SetClient(doudian_sdk.DefaultDoudianOpSpiClient)
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetSpiParam(new(doudian_sdk.DoudianOpSpiParam))
	request.Param = new(SendHomeOrderFinishMsgParam)
	response := new(sendHome_order_finishMsg_response.SendHomeOrderFinishMsgResponse)
	response.Data = new(sendHome_order_finishMsg_response.SendHomeOrderFinishMsgData)
	request.Response = response
	return request
}

type SendHomeOrderFinishMsgParam struct {
	// 消息具体内容
	MsgBody string `json:"msg_body"`
	// 消息类型2 表示订单完成
	SpiType int64 `json:"spi_type"`
	// 消息唯一id
	MsgId string `json:"msg_id"`
}
