package sendHome_order_acceptOrder_response

import (
	"github.com/ldxtechteam/doudiansdk/core"
)

type SendHomeOrderAcceptOrderResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SendHomeOrderAcceptOrderData `json:"data"`
}
type SendHomeOrderAcceptOrderData struct {
}
