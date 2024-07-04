package sendHome_product_batchUpdateStock_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SendHomeProductBatchUpdateStockResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SendHomeProductBatchUpdateStockData `json:"data"`
}
type SendHomeProductBatchUpdateStockData struct {
}
