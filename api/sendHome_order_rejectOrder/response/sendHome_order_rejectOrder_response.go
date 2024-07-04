package sendHome_order_rejectOrder_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SendHomeOrderRejectOrderResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SendHomeOrderRejectOrderData `json:"data"`
}
type SendHomeOrderRejectOrderData struct {
}
