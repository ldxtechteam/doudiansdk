package sendHome_product_productStatusChange_request

import (
	"doudian.com/open/sdk_golang/core"
	"doudian.com/open/sdk_golang/spi/sendHome_product_productStatusChange/response"
	"doudian.com/open/sdk_golang/utils"
)

type SendHomeProductProductStatusChangeRequest struct {
	doudian_sdk.BaseDoudianOpSpiRequest
	Param *SendHomeProductProductStatusChangeParam 
	Response *sendHome_product_productStatusChange_response.SendHomeProductProductStatusChangeResponse 
}
func (c *SendHomeProductProductStatusChangeRequest) GetParamJsonObject() interface{}{
	return c.Param
}


func (c *SendHomeProductProductStatusChangeRequest) GetResponseObject() interface{}{
	return c.Response
}


func (c *SendHomeProductProductStatusChangeRequest) Execute() (interface{}, error){
	return c.GetClient().Request(c)
}


func (c *SendHomeProductProductStatusChangeRequest) ResponseJson() (string, error){
	responseObj, err := c.Execute()
	if err != nil {
		return "", err
	}
	return utils.MarshalNoErr(responseObj), nil
}


func New() *SendHomeProductProductStatusChangeRequest{
	
	request := new(SendHomeProductProductStatusChangeRequest)
	request.SetClient(doudian_sdk.DefaultDoudianOpSpiClient)
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetSpiParam(new(doudian_sdk.DoudianOpSpiParam))
	request.Param = new(SendHomeProductProductStatusChangeParam)
	response := new(sendHome_product_productStatusChange_response.SendHomeProductProductStatusChangeResponse)
	response.Data = new(sendHome_product_productStatusChange_response.SendHomeProductProductStatusChangeData)
	request.Response = response
	return request
}


type SendHomeProductProductStatusChangeParam struct {
	// SPI类型，服务商可根据type做不同的回调逻辑商品状态变化类型：21
	SpiType int64 `json:"spi_type"`
	// SPI具体请求内容，具体信息。Event：「商品变更事件，具体枚举和解释见：：https://op.jinritemai.com/docs/question-docs/99/2354」ProductId：「商品id」outer_product_id：「外部商品id」shop_id：「总部id」store_id：「门店id」update_time：「事件发生时间，时间戳，秒」main_product：「主品信息，如果当前商品是子品」main_product.outer_product_id：「主商品的外部商品ID」main_product.product_id：「主商品的外部商品ID」
	MsgBody string `json:"msg_body"`
	// 消息唯一ID，可根据该信息进行幂等处理
	MsgId string `json:"msg_id"`
}
