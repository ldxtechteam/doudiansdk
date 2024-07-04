package sendHome_afterSale_closeRefundMsg_request

import (
	"doudian.com/open/sdk_golang/core"
	"doudian.com/open/sdk_golang/spi/sendHome_afterSale_closeRefundMsg/response"
	"doudian.com/open/sdk_golang/utils"
)

type SendHomeAfterSaleCloseRefundMsgRequest struct {
	doudian_sdk.BaseDoudianOpSpiRequest
	Param *SendHomeAfterSaleCloseRefundMsgParam 
	Response *sendHome_afterSale_closeRefundMsg_response.SendHomeAfterSaleCloseRefundMsgResponse 
}
func (c *SendHomeAfterSaleCloseRefundMsgRequest) GetParamJsonObject() interface{}{
	return c.Param
}


func (c *SendHomeAfterSaleCloseRefundMsgRequest) GetResponseObject() interface{}{
	return c.Response
}


func (c *SendHomeAfterSaleCloseRefundMsgRequest) Execute() (interface{}, error){
	return c.GetClient().Request(c)
}


func (c *SendHomeAfterSaleCloseRefundMsgRequest) ResponseJson() (string, error){
	responseObj, err := c.Execute()
	if err != nil {
		return "", err
	}
	return utils.MarshalNoErr(responseObj), nil
}


func New() *SendHomeAfterSaleCloseRefundMsgRequest{
	
	request := new(SendHomeAfterSaleCloseRefundMsgRequest)
	request.SetClient(doudian_sdk.DefaultDoudianOpSpiClient)
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetSpiParam(new(doudian_sdk.DoudianOpSpiParam))
	request.Param = new(SendHomeAfterSaleCloseRefundMsgParam)
	response := new(sendHome_afterSale_closeRefundMsg_response.SendHomeAfterSaleCloseRefundMsgResponse)
	response.Data = new(sendHome_afterSale_closeRefundMsg_response.SendHomeAfterSaleCloseRefundMsgData)
	request.Response = response
	return request
}


type SendHomeAfterSaleCloseRefundMsgParam struct {
	// after_sale_id:售后单idstore_id:门店idshop_id:店铺idorder_id:订单IDclose_time:退款关闭时间update_time:售后单更新时间
	MsgBody string `json:"msg_body"`
	// 消息类型
	SpiType int64 `json:"spi_type"`
	// 消息唯一键
	MsgId string `json:"msg_id"`
}
