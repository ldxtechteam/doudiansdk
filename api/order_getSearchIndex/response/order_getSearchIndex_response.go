package order_getSearchIndex_response

import (
	"github.com/ldxtechteam/doudiansdk/core"
)

type OrderGetSearchIndexResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderGetSearchIndexData `json:"data"`
}
type OrderGetSearchIndexData struct {
	// 索引串
	EncryptIndexText string `json:"encrypt_index_text"`
}
