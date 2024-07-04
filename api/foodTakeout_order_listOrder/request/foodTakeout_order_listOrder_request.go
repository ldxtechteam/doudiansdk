package foodTakeout_order_listOrder_request

import (
	"doudian.com/open/sdk_golang/api/foodTakeout_order_listOrder/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type FoodTakeoutOrderListOrderRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *FoodTakeoutOrderListOrderParam 
}
func (c *FoodTakeoutOrderListOrderRequest) GetUrlPath() string{
	return "/foodTakeout/order/listOrder"
}


func New() *FoodTakeoutOrderListOrderRequest{
	request := &FoodTakeoutOrderListOrderRequest{
		Param: &FoodTakeoutOrderListOrderParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *FoodTakeoutOrderListOrderRequest) Execute(accessToken *doudian_sdk.AccessToken) (*foodTakeout_order_listOrder_response.FoodTakeoutOrderListOrderResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &foodTakeout_order_listOrder_response.FoodTakeoutOrderListOrderResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *FoodTakeoutOrderListOrderRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *FoodTakeoutOrderListOrderRequest) GetParams() *FoodTakeoutOrderListOrderParam{
	return c.Param
}


type Pagination struct {
	// 第几页(从1开始)
	PageNo *int32 `json:"page_no"`
	// 每页条数,最大为10
	PageSize *int32 `json:"page_size"`
}
type FoodTakeoutOrderListOrderParam struct {
	// 订单创建结束时间，时间戳（秒）
	CreateEndTime *int64 `json:"create_end_time"`
	// 查询状态。1:待支付 2：待接单 4:已关闭 5:已完成
	OrderSearchStatus *int32 `json:"order_search_status"`
	// 门店id
	StoreId *int64 `json:"store_id"`
	// 分页
	Pagination *Pagination `json:"pagination"`
	// 订单创建开始时间。时间戳（秒）
	CreateBeginTime *int64 `json:"create_begin_time"`
}
