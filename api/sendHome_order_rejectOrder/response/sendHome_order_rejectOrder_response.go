package sendHome_order_rejectOrder_response

import (
	"github.com/ldxtechteam/doudiansdk/core"
)

type SendHomeOrderRejectOrderResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SendHomeOrderRejectOrderData `json:"data"`
}
type SendHomeOrderRejectOrderData struct {
}
