package sendHome_order_acceptOrder_request

import (
	"doudian.com/open/sdk_golang/api/sendHome_order_acceptOrder/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SendHomeOrderAcceptOrderRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SendHomeOrderAcceptOrderParam 
}
func (c *SendHomeOrderAcceptOrderRequest) GetUrlPath() string{
	return "/sendHome/order/acceptOrder"
}


func New() *SendHomeOrderAcceptOrderRequest{
	request := &SendHomeOrderAcceptOrderRequest{
		Param: &SendHomeOrderAcceptOrderParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SendHomeOrderAcceptOrderRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sendHome_order_acceptOrder_response.SendHomeOrderAcceptOrderResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sendHome_order_acceptOrder_response.SendHomeOrderAcceptOrderResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SendHomeOrderAcceptOrderRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SendHomeOrderAcceptOrderRequest) GetParams() *SendHomeOrderAcceptOrderParam{
	return c.Param
}


type SendHomeOrderAcceptOrderParam struct {
	// 订单id
	OrderId string `json:"order_id"`
	// 门店id
	StoreId int64 `json:"store_id"`
}
