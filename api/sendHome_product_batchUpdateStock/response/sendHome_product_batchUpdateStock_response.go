package sendHome_product_batchUpdateStock_response

import (
	"github.com/ldxtechteam/doudiansdk/core"
)

type SendHomeProductBatchUpdateStockResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SendHomeProductBatchUpdateStockData `json:"data"`
}
type SendHomeProductBatchUpdateStockData struct {
}
