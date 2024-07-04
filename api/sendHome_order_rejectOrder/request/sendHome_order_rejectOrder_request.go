package sendHome_order_rejectOrder_request

import (
	"encoding/json"
)

type SendHomeOrderRejectOrderRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SendHomeOrderRejectOrderParam
}

func (c *SendHomeOrderRejectOrderRequest) GetUrlPath() string {
	return "/sendHome/order/rejectOrder"
}

func New() *SendHomeOrderRejectOrderRequest {
	request := &SendHomeOrderRejectOrderRequest{
		Param: &SendHomeOrderRejectOrderParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *SendHomeOrderRejectOrderRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sendHome_order_rejectOrder_response.SendHomeOrderRejectOrderResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sendHome_order_rejectOrder_response.SendHomeOrderRejectOrderResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *SendHomeOrderRejectOrderRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *SendHomeOrderRejectOrderRequest) GetParams() *SendHomeOrderRejectOrderParam {
	return c.Param
}

type SendHomeOrderRejectOrderParam struct {
	// 门店id
	StoreId int64 `json:"store_id"`
	// 具体错误枚举[    {        "code":12,        "reason":"联系不上用户"    },    {        "code":72,        "reason":"商品已售完"    },    {        "code":73,        "reason":"商家已打烊"    },    {        "code":74,        "reason":"商家订单太多忙不过来"    },    {        "code":75,        "reason":"超出配送范围"    },    {        "code":76,        "reason":"没有配送员接单"    },    {        "code":77,        "reason":"商家超时未接单，订单已被取消"    }]
	RejectCode int64 `json:"reject_code"`
	// 拒绝原因
	RejectMsg *string `json:"reject_msg"`
	// 订单id
	OrderId string `json:"order_id"`
}
