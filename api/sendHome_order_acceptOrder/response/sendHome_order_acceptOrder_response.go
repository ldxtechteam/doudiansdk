package sendHome_order_acceptOrder_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SendHomeOrderAcceptOrderResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SendHomeOrderAcceptOrderData `json:"data"`
}
type SendHomeOrderAcceptOrderData struct {
}
