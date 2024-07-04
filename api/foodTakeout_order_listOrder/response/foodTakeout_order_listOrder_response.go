package foodTakeout_order_listOrder_response

import (
	"doudian.com/open/sdk_golang/core"
)

type FoodTakeoutOrderListOrderResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *FoodTakeoutOrderListOrderData `json:"data"`
}
type Pagination struct {
	// 第几页(从1开始)
	PageNo int32 `json:"page_no"`
	// 每页条数,最大为10
	PageSize int32 `json:"page_size"`
	// 总数
	Total int32 `json:"total"`
	// 是否还能查询更多
	HasMore bool `json:"has_more"`
}
type FoodTakeoutOrderListOrderData struct {
	// 订单ids
	OrderIds []string `json:"order_ids"`
	// 分页值
	Pagination *Pagination `json:"pagination"`
}
