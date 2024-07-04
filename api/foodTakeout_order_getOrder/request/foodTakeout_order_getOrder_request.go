package foodTakeout_order_getOrder_request

import (
	"encoding/json"
	"github.com/ldxtechteam/doudiansdk/api/foodTakeout_order_getOrder/response"
	"github.com/ldxtechteam/doudiansdk/core"
)

type FoodTakeoutOrderGetOrderRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *FoodTakeoutOrderGetOrderParam
}

func (c *FoodTakeoutOrderGetOrderRequest) GetUrlPath() string {
	return "/foodTakeout/order/getOrder"
}

func New() *FoodTakeoutOrderGetOrderRequest {
	request := &FoodTakeoutOrderGetOrderRequest{
		Param: &FoodTakeoutOrderGetOrderParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *FoodTakeoutOrderGetOrderRequest) Execute(accessToken *doudian_sdk.AccessToken) (*foodTakeout_order_getOrder_response.FoodTakeoutOrderGetOrderResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &foodTakeout_order_getOrder_response.FoodTakeoutOrderGetOrderResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *FoodTakeoutOrderGetOrderRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *FoodTakeoutOrderGetOrderRequest) GetParams() *FoodTakeoutOrderGetOrderParam {
	return c.Param
}

type FoodTakeoutOrderGetOrderParam struct {
	// 订单idlist
	OrderIds []string `json:"order_ids"`
}
