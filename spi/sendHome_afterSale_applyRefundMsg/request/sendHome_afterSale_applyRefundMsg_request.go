package sendHome_afterSale_applyRefundMsg_request

import (
	"github.com/ldxtechteam/doudiansdk/core"
	"github.com/ldxtechteam/doudiansdk/spi/sendHome_afterSale_applyRefundMsg/response"
	"github.com/ldxtechteam/doudiansdk/utils"
)

type SendHomeAfterSaleApplyRefundMsgRequest struct {
	doudian_sdk.BaseDoudianOpSpiRequest
	Param    *SendHomeAfterSaleApplyRefundMsgParam
	Response *sendHome_afterSale_applyRefundMsg_response.SendHomeAfterSaleApplyRefundMsgResponse
}

func (c *SendHomeAfterSaleApplyRefundMsgRequest) GetParamJsonObject() interface{} {
	return c.Param
}

func (c *SendHomeAfterSaleApplyRefundMsgRequest) GetResponseObject() interface{} {
	return c.Response
}

func (c *SendHomeAfterSaleApplyRefundMsgRequest) Execute() (interface{}, error) {
	return c.GetClient().Request(c)
}

func (c *SendHomeAfterSaleApplyRefundMsgRequest) ResponseJson() (string, error) {
	responseObj, err := c.Execute()
	if err != nil {
		return "", err
	}
	return utils.MarshalNoErr(responseObj), nil
}

func New() *SendHomeAfterSaleApplyRefundMsgRequest {

	request := new(SendHomeAfterSaleApplyRefundMsgRequest)
	request.SetClient(doudian_sdk.DefaultDoudianOpSpiClient)
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetSpiParam(new(doudian_sdk.DoudianOpSpiParam))
	request.Param = new(SendHomeAfterSaleApplyRefundMsgParam)
	response := new(sendHome_afterSale_applyRefundMsg_response.SendHomeAfterSaleApplyRefundMsgResponse)
	response.Data = new(sendHome_afterSale_applyRefundMsg_response.SendHomeAfterSaleApplyRefundMsgData)
	request.Response = response
	return request
}

type SendHomeAfterSaleApplyRefundMsgParam struct {
	// after_sale_id:售后单idstore_id:门店idshop_id:店铺idorder_id:订单IDreason_code:申请售后原因apply_time：售后申请时间refund_type：退款类型。1：店铺单退2：sku单退part_type：是否是自定义金额退款 1: 金额全部退款 2: 金额部分退款refund_amount： 申请退款总金额refund_post_amount：申请退款的配送费refund_packing_amount：申请退款的打包费evidence：退款凭证refund_sku_list：退款商品（列表）
	MsgBody string `json:"msg_body"`
	// 消息类型5
	SpiType int64 `json:"spi_type"`
	// 唯一id
	MsgId string `json:"msg_id"`
}
