package sendHome_product_batchUpdatePrice_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SendHomeProductBatchUpdatePriceResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SendHomeProductBatchUpdatePriceData `json:"data"`
}
type ErrListItem struct {
	// 修改错误对应的外部商品编码，如果请求参数有该字段，则返回该字段
	OuterProductId string `json:"outer_product_id"`
	// 修改错误对应的skuid，如果请求参数有该字段，则返回该字段
	SkuId int64 `json:"sku_id"`
	// 修改错误对应的外部sku编码，如果请求参数有该字段，则返回该字段
	OuterSkuId string `json:"outer_sku_id"`
	// 错误的原因
	ErrReason string `json:"err_reason"`
	// 修改错误对应的商品id，如果请求参数有该字段，则返回该字段
	ProductId int64 `json:"product_id"`
}
type SendHomeProductBatchUpdatePriceData struct {
	// 非全部成功时，对应的错误数据的原因
	ErrList []ErrListItem `json:"err_list"`
}
