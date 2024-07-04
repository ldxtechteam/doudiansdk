package sendHome_afterSale_modifyRefundMsg_request

import (
	"doudian.com/open/sdk_golang/core"
	"doudian.com/open/sdk_golang/spi/sendHome_afterSale_modifyRefundMsg/response"
	"doudian.com/open/sdk_golang/utils"
)

type SendHomeAfterSaleModifyRefundMsgRequest struct {
	doudian_sdk.BaseDoudianOpSpiRequest
	Param *SendHomeAfterSaleModifyRefundMsgParam 
	Response *sendHome_afterSale_modifyRefundMsg_response.SendHomeAfterSaleModifyRefundMsgResponse 
}
func (c *SendHomeAfterSaleModifyRefundMsgRequest) GetParamJsonObject() interface{}{
	return c.Param
}


func (c *SendHomeAfterSaleModifyRefundMsgRequest) GetResponseObject() interface{}{
	return c.Response
}


func (c *SendHomeAfterSaleModifyRefundMsgRequest) Execute() (interface{}, error){
	return c.GetClient().Request(c)
}


func (c *SendHomeAfterSaleModifyRefundMsgRequest) ResponseJson() (string, error){
	responseObj, err := c.Execute()
	if err != nil {
		return "", err
	}
	return utils.MarshalNoErr(responseObj), nil
}


func New() *SendHomeAfterSaleModifyRefundMsgRequest{
	
	request := new(SendHomeAfterSaleModifyRefundMsgRequest)
	request.SetClient(doudian_sdk.DefaultDoudianOpSpiClient)
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetSpiParam(new(doudian_sdk.DoudianOpSpiParam))
	request.Param = new(SendHomeAfterSaleModifyRefundMsgParam)
	response := new(sendHome_afterSale_modifyRefundMsg_response.SendHomeAfterSaleModifyRefundMsgResponse)
	response.Data = new(sendHome_afterSale_modifyRefundMsg_response.SendHomeAfterSaleModifyRefundMsgData)
	request.Response = response
	return request
}


type SendHomeAfterSaleModifyRefundMsgParam struct {
	// after_sale_id:售后单idstore_id:门店idshop_id:店铺idorder_id:订单IDreason_code:申请售后原因apply_time：售后申请时间refund_type：退款类型。1：店铺单退2：sku单退part_type：是否是自定义金额退款 1: 金额全部退款 2: 金额部分退款refund_amount： 申请退款总金额refund_post_amount：申请退款的配送费refund_packing_amount：申请退款的打包费evidence：退款凭证refund_sku_list：退款商品（列表）
	MsgBody string `json:"msg_body"`
	// 唯一id
	MsgId string `json:"msg_id"`
	// 消息类型
	SpiType int64 `json:"spi_type"`
}
