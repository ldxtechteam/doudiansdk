package sendHome_afterSale_completeRefundMsg_request

import (
	"github.com/ldxtechteam/doudiansdk/utils"
)

type SendHomeAfterSaleCompleteRefundMsgRequest struct {
	doudian_sdk.BaseDoudianOpSpiRequest
	Param    *SendHomeAfterSaleCompleteRefundMsgParam
	Response *sendHome_afterSale_completeRefundMsg_response.SendHomeAfterSaleCompleteRefundMsgResponse
}

func (c *SendHomeAfterSaleCompleteRefundMsgRequest) GetParamJsonObject() interface{} {
	return c.Param
}

func (c *SendHomeAfterSaleCompleteRefundMsgRequest) GetResponseObject() interface{} {
	return c.Response
}

func (c *SendHomeAfterSaleCompleteRefundMsgRequest) Execute() (interface{}, error) {
	return c.GetClient().Request(c)
}

func (c *SendHomeAfterSaleCompleteRefundMsgRequest) ResponseJson() (string, error) {
	responseObj, err := c.Execute()
	if err != nil {
		return "", err
	}
	return utils.MarshalNoErr(responseObj), nil
}

func New() *SendHomeAfterSaleCompleteRefundMsgRequest {

	request := new(SendHomeAfterSaleCompleteRefundMsgRequest)
	request.SetClient(doudian_sdk.DefaultDoudianOpSpiClient)
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetSpiParam(new(doudian_sdk.DoudianOpSpiParam))
	request.Param = new(SendHomeAfterSaleCompleteRefundMsgParam)
	response := new(sendHome_afterSale_completeRefundMsg_response.SendHomeAfterSaleCompleteRefundMsgResponse)
	response.Data = new(sendHome_afterSale_completeRefundMsg_response.SendHomeAfterSaleCompleteRefundMsgData)
	request.Response = response
	return request
}

type SendHomeAfterSaleCompleteRefundMsgParam struct {
	// after_sale_id:售后单idstore_id:门店idshop_id:店铺idorder_id:订单IDreason_code:申请售后原因apply_time：售后申请时间refund_type：退款类型。1：店铺单退2：sku单退part_type：是否是自定义金额退款 1: 金额全部退款 2: 金额部分退款refund_amount： 申请退款总金额refund_post_amount：申请退款的配送费refund_packing_amount：申请退款的打包费evidence：退款凭证refund_sku_list：退款商品（列表）
	MsgBody string `json:"msg_body"`
	// 消息类型
	SpiType int64 `json:"spi_type"`
	// 消息唯一id
	MsgId string `json:"msg_id"`
}
