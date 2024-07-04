package sendHome_order_acceptMsg_request

import (
	"github.com/ldxtechteam/doudiansdk/core"
	"github.com/ldxtechteam/doudiansdk/spi/sendHome_order_acceptMsg/response"
	"github.com/ldxtechteam/doudiansdk/utils"
)

type SendHomeOrderAcceptMsgRequest struct {
	doudian_sdk.BaseDoudianOpSpiRequest
	Param    *SendHomeOrderAcceptMsgParam
	Response *sendHome_order_acceptMsg_response.SendHomeOrderAcceptMsgResponse
}

func (c *SendHomeOrderAcceptMsgRequest) GetParamJsonObject() interface{} {
	return c.Param
}

func (c *SendHomeOrderAcceptMsgRequest) GetResponseObject() interface{} {
	return c.Response
}

func (c *SendHomeOrderAcceptMsgRequest) Execute() (interface{}, error) {
	return c.GetClient().Request(c)
}

func (c *SendHomeOrderAcceptMsgRequest) ResponseJson() (string, error) {
	responseObj, err := c.Execute()
	if err != nil {
		return "", err
	}
	return utils.MarshalNoErr(responseObj), nil
}

func New() *SendHomeOrderAcceptMsgRequest {

	request := new(SendHomeOrderAcceptMsgRequest)
	request.SetClient(doudian_sdk.DefaultDoudianOpSpiClient)
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetSpiParam(new(doudian_sdk.DoudianOpSpiParam))
	request.Param = new(SendHomeOrderAcceptMsgParam)
	response := new(sendHome_order_acceptMsg_response.SendHomeOrderAcceptMsgResponse)
	response.Data = new(sendHome_order_acceptMsg_response.SendHomeOrderAcceptMsgData)
	request.Response = response
	return request
}

type SendHomeOrderAcceptMsgParam struct {
	// 业务参数:订单id、门店id、总户id
	MsgBody string `json:"msg_body"`
	// 消息类型
	SpiType int64 `json:"spi_type"`
	// 消息唯一id
	MsgId string `json:"msg_id"`
}
