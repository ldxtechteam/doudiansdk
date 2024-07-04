package sendHome_product_batchUpdateStock_request

import (
	"encoding/json"
)

type SendHomeProductBatchUpdateStockRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SendHomeProductBatchUpdateStockParam
}

func (c *SendHomeProductBatchUpdateStockRequest) GetUrlPath() string {
	return "/sendHome/product/batchUpdateStock"
}

func New() *SendHomeProductBatchUpdateStockRequest {
	request := &SendHomeProductBatchUpdateStockRequest{
		Param: &SendHomeProductBatchUpdateStockParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *SendHomeProductBatchUpdateStockRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sendHome_product_batchUpdateStock_response.SendHomeProductBatchUpdateStockResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sendHome_product_batchUpdateStock_response.SendHomeProductBatchUpdateStockResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *SendHomeProductBatchUpdateStockRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *SendHomeProductBatchUpdateStockRequest) GetParams() *SendHomeProductBatchUpdateStockParam {
	return c.Param
}

type SkuStocksItem struct {
	// 商品规格ID；一次请求sku维度最大支持50个sku更新
	SkuId *int64 `json:"sku_id"`
	// 外部skuid；
	OuterSkuId *string `json:"outer_sku_id"`
	// 现货库存数
	StockNum int64 `json:"stock_num"`
}
type UpdateItemsItem struct {
	// 例：批量更新10个skuid库存，其中一个skuid信息不正确，这样整个请求都会失败，10个skuid都未更新成功。商品ID，商品id和外部商品id二选一传输，优先使用product_id。注：抖音侧的商品id需要配合抖音侧的skuId使用
	ProductId *int64 `json:"product_id"`
	// 外部商品id。注：外部的商品id需要配合外部的skuId使用
	OuterProductId *string `json:"outer_product_id"`
	// 要修改的规格id对应的库存数
	SkuStocks []SkuStocksItem `json:"sku_stocks"`
}
type SendHomeProductBatchUpdateStockParam struct {
	// 操作视角，2和3时，store_id和outer_store_id选填
	StoreId *int64 `json:"store_id"`
	// 相对于store_id 是平台侧的门店id，outer_store_id是外部门店编码，使用上，选择其中一个，优先使用 store_id
	OuterStoreId *string `json:"outer_store_id"`
	// 库存更新方式，必填字段；true-增量更新；false-全量更新
	Incremental bool `json:"incremental"`
	// 幂等ID，incremental=true时必须传递。请保证24小时内在店铺下唯一性，建议使用年月日时分秒+随机数生成
	IdempotentId *string `json:"idempotent_id"`
	// 操作视角。2:子品id操作子品;3:主品操作子品。注：普通商品需要传递2
	OperatePerspective int32 `json:"operate_perspective"`
	// 修改库存的信息列表
	UpdateItems []UpdateItemsItem `json:"update_items"`
}
