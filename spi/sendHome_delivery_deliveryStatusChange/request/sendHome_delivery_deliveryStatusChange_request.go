package sendHome_delivery_deliveryStatusChange_request

import (
	"doudian.com/open/sdk_golang/core"
	"doudian.com/open/sdk_golang/spi/sendHome_delivery_deliveryStatusChange/response"
	"doudian.com/open/sdk_golang/utils"
)

type SendHomeDeliveryDeliveryStatusChangeRequest struct {
	doudian_sdk.BaseDoudianOpSpiRequest
	Param *SendHomeDeliveryDeliveryStatusChangeParam 
	Response *sendHome_delivery_deliveryStatusChange_response.SendHomeDeliveryDeliveryStatusChangeResponse 
}
func (c *SendHomeDeliveryDeliveryStatusChangeRequest) GetParamJsonObject() interface{}{
	return c.Param
}


func (c *SendHomeDeliveryDeliveryStatusChangeRequest) GetResponseObject() interface{}{
	return c.Response
}


func (c *SendHomeDeliveryDeliveryStatusChangeRequest) Execute() (interface{}, error){
	return c.GetClient().Request(c)
}


func (c *SendHomeDeliveryDeliveryStatusChangeRequest) ResponseJson() (string, error){
	responseObj, err := c.Execute()
	if err != nil {
		return "", err
	}
	return utils.MarshalNoErr(responseObj), nil
}


func New() *SendHomeDeliveryDeliveryStatusChangeRequest{
	
	request := new(SendHomeDeliveryDeliveryStatusChangeRequest)
	request.SetClient(doudian_sdk.DefaultDoudianOpSpiClient)
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetSpiParam(new(doudian_sdk.DoudianOpSpiParam))
	request.Param = new(SendHomeDeliveryDeliveryStatusChangeParam)
	response := new(sendHome_delivery_deliveryStatusChange_response.SendHomeDeliveryDeliveryStatusChangeResponse)
	response.Data = new(sendHome_delivery_deliveryStatusChange_response.SendHomeDeliveryDeliveryStatusChangeData)
	request.Response = response
	return request
}


type SendHomeDeliveryDeliveryStatusChangeParam struct {
	// SPI类型，服务商可根据type做不同的回调逻辑运力状态变化对应枚举：31
	SpiType int64 `json:"spi_type"`
	// SPI具体请求内容，具体信息。order_id：「订单id」rider_name：「骑手姓名； 骑手已接单/到店/取货/送达状态时存在；」rider_phone：「骑手电话； 骑手已接单/到店/取货/送达状态时存在；」operate_time：「状态变更时间戳，秒」store_id：「门店id」shop_id：「总部id」status：「运力状态，101:待骑手接单 102:骑手已接单 103:骑手已到店 104:骑手已取货 200:已送达 300:已取消」cancel_code：「取消原因，100:其他原因取消(商家发起) 201:超出配送范围 202:骑手原因无法完成配送 203:运力紧张无可接单骑手 204:骑手联系不上收货人 206:物品超长超重无法完成配送 208:商家未营业 209:商家无法出餐 210:商家地址错误 211:用户拒收 222:用户自身原因要求取消订单 200:其他原因。注：若非上述取消码，详情含义请联系平台」delivery_service_type：「配送模式；1-平台运力配送;2-自配送」
	MsgBody string `json:"msg_body"`
	// 消息唯一ID，可根据该信息进行幂等处理
	MsgId string `json:"msg_id"`
}
