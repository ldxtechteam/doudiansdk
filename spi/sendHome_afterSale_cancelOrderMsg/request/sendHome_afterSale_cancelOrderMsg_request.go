package sendHome_afterSale_cancelOrderMsg_request

import (
	"github.com/ldxtechteam/doudiansdk/utils"
)

type SendHomeAfterSaleCancelOrderMsgRequest struct {
	doudian_sdk.BaseDoudianOpSpiRequest
	Param    *SendHomeAfterSaleCancelOrderMsgParam
	Response *sendHome_afterSale_cancelOrderMsg_response.SendHomeAfterSaleCancelOrderMsgResponse
}

func (c *SendHomeAfterSaleCancelOrderMsgRequest) GetParamJsonObject() interface{} {
	return c.Param
}

func (c *SendHomeAfterSaleCancelOrderMsgRequest) GetResponseObject() interface{} {
	return c.Response
}

func (c *SendHomeAfterSaleCancelOrderMsgRequest) Execute() (interface{}, error) {
	return c.GetClient().Request(c)
}

func (c *SendHomeAfterSaleCancelOrderMsgRequest) ResponseJson() (string, error) {
	responseObj, err := c.Execute()
	if err != nil {
		return "", err
	}
	return utils.MarshalNoErr(responseObj), nil
}

func New() *SendHomeAfterSaleCancelOrderMsgRequest {

	request := new(SendHomeAfterSaleCancelOrderMsgRequest)
	request.SetClient(doudian_sdk.DefaultDoudianOpSpiClient)
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetSpiParam(new(doudian_sdk.DoudianOpSpiParam))
	request.Param = new(SendHomeAfterSaleCancelOrderMsgParam)
	response := new(sendHome_afterSale_cancelOrderMsg_response.SendHomeAfterSaleCancelOrderMsgResponse)
	response.Data = new(sendHome_afterSale_cancelOrderMsg_response.SendHomeAfterSaleCancelOrderMsgData)
	request.Response = response
	return request
}

type SendHomeAfterSaleCancelOrderMsgParam struct {
	// shop_id:总户idstore_id:门店idorder_id:订单id
	MsgBody string `json:"msg_body"`
	// 消息id
	MsgId string `json:"msg_id"`
	// 消息类型
	SpiType int64 `json:"spi_type"`
}
