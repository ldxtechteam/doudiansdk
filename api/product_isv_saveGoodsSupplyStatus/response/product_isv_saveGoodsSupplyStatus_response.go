package product_isv_saveGoodsSupplyStatus_response

import (
	"github.com/ldxtechteam/doudiansdk/core"
)

type ProductIsvSaveGoodsSupplyStatusResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductIsvSaveGoodsSupplyStatusData `json:"data"`
}
type ProductIsvSaveGoodsSupplyStatusData struct {
}
