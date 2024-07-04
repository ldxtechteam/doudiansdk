package sendHome_afterSale_auditResultMsg_request

import (
	"github.com/ldxtechteam/doudiansdk/core"
	"github.com/ldxtechteam/doudiansdk/spi/sendHome_afterSale_auditResultMsg/response"
	"github.com/ldxtechteam/doudiansdk/utils"
)

type SendHomeAfterSaleAuditResultMsgRequest struct {
	doudian_sdk.BaseDoudianOpSpiRequest
	Param    *SendHomeAfterSaleAuditResultMsgParam
	Response *sendHome_afterSale_auditResultMsg_response.SendHomeAfterSaleAuditResultMsgResponse
}

func (c *SendHomeAfterSaleAuditResultMsgRequest) GetParamJsonObject() interface{} {
	return c.Param
}

func (c *SendHomeAfterSaleAuditResultMsgRequest) GetResponseObject() interface{} {
	return c.Response
}

func (c *SendHomeAfterSaleAuditResultMsgRequest) Execute() (interface{}, error) {
	return c.GetClient().Request(c)
}

func (c *SendHomeAfterSaleAuditResultMsgRequest) ResponseJson() (string, error) {
	responseObj, err := c.Execute()
	if err != nil {
		return "", err
	}
	return utils.MarshalNoErr(responseObj), nil
}

func New() *SendHomeAfterSaleAuditResultMsgRequest {

	request := new(SendHomeAfterSaleAuditResultMsgRequest)
	request.SetClient(doudian_sdk.DefaultDoudianOpSpiClient)
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetSpiParam(new(doudian_sdk.DoudianOpSpiParam))
	request.Param = new(SendHomeAfterSaleAuditResultMsgParam)
	response := new(sendHome_afterSale_auditResultMsg_response.SendHomeAfterSaleAuditResultMsgResponse)
	response.Data = new(sendHome_afterSale_auditResultMsg_response.SendHomeAfterSaleAuditResultMsgData)
	request.Response = response
	return request
}

type SendHomeAfterSaleAuditResultMsgParam struct {
	// after_sale_id:售后单idstore_id:门店idshop_id:店铺idorder_id:订单IDresult：是否同意退款,1同意,2拒绝reject_reason_code：拒绝原因update_time：售后单更新时间，秒级时间戳
	MsgBody string `json:"msg_body"`
	// 消息类型6
	SpiType int64 `json:"spi_type"`
	// 消息唯一id
	MsgId string `json:"msg_id"`
}
