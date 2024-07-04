package sendHome_product_updateProductStatus_response

import (
	"github.com/ldxtechteam/doudiansdk/core"
)

type SendHomeProductUpdateProductStatusResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SendHomeProductUpdateProductStatusData `json:"data"`
}
type SendHomeProductUpdateProductStatusData struct {
}
