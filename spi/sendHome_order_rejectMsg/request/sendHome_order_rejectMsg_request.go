package sendHome_order_rejectMsg_request

import (
	"doudian.com/open/sdk_golang/core"
	"doudian.com/open/sdk_golang/spi/sendHome_order_rejectMsg/response"
	"doudian.com/open/sdk_golang/utils"
)

type SendHomeOrderRejectMsgRequest struct {
	doudian_sdk.BaseDoudianOpSpiRequest
	Param *SendHomeOrderRejectMsgParam 
	Response *sendHome_order_rejectMsg_response.SendHomeOrderRejectMsgResponse 
}
func (c *SendHomeOrderRejectMsgRequest) GetParamJsonObject() interface{}{
	return c.Param
}


func (c *SendHomeOrderRejectMsgRequest) GetResponseObject() interface{}{
	return c.Response
}


func (c *SendHomeOrderRejectMsgRequest) Execute() (interface{}, error){
	return c.GetClient().Request(c)
}


func (c *SendHomeOrderRejectMsgRequest) ResponseJson() (string, error){
	responseObj, err := c.Execute()
	if err != nil {
		return "", err
	}
	return utils.MarshalNoErr(responseObj), nil
}


func New() *SendHomeOrderRejectMsgRequest{
	
	request := new(SendHomeOrderRejectMsgRequest)
	request.SetClient(doudian_sdk.DefaultDoudianOpSpiClient)
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetSpiParam(new(doudian_sdk.DoudianOpSpiParam))
	request.Param = new(SendHomeOrderRejectMsgParam)
	response := new(sendHome_order_rejectMsg_response.SendHomeOrderRejectMsgResponse)
	response.Data = new(sendHome_order_rejectMsg_response.SendHomeOrderRejectMsgData)
	request.Response = response
	return request
}


type SendHomeOrderRejectMsgParam struct {
	// 业务信息:包含订单id、门店id、总户id
	MsgBody string `json:"msg_body"`
	// 消息类型4
	SpiType int64 `json:"spi_type"`
	// 消息唯一键
	MsgId string `json:"msg_id"`
}
