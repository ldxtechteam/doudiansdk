package sendHome_product_batchUpdatePrice_request

import (
	"doudian.com/open/sdk_golang/api/sendHome_product_batchUpdatePrice/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SendHomeProductBatchUpdatePriceRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SendHomeProductBatchUpdatePriceParam 
}
func (c *SendHomeProductBatchUpdatePriceRequest) GetUrlPath() string{
	return "/sendHome/product/batchUpdatePrice"
}


func New() *SendHomeProductBatchUpdatePriceRequest{
	request := &SendHomeProductBatchUpdatePriceRequest{
		Param: &SendHomeProductBatchUpdatePriceParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SendHomeProductBatchUpdatePriceRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sendHome_product_batchUpdatePrice_response.SendHomeProductBatchUpdatePriceResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sendHome_product_batchUpdatePrice_response.SendHomeProductBatchUpdatePriceResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SendHomeProductBatchUpdatePriceRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SendHomeProductBatchUpdatePriceRequest) GetParams() *SendHomeProductBatchUpdatePriceParam{
	return c.Param
}


type SkuPricesItem struct {
	// 商品规格ID；一次请求sku维度最大支持20个sku更新
	SkuId *int64 `json:"sku_id"`
	// 外部skuid；
	OuterSkuId *string `json:"outer_sku_id"`
	// 价格, 单位分
	Price int64 `json:"price"`
}
type UpdateItemsItem struct {
	// 商品ID，商品id和外部商品id二选一传输，优先使用product_id。注：抖音侧的商品id需要配合抖音侧的skuId使用
	ProductId *int64 `json:"product_id"`
	// 外部商品id。注：外部的商品id需要配合外部的skuId使用
	OuterProductId *string `json:"outer_product_id"`
	// 要修改的规格id对应的价格,单位分
	SkuPrices []SkuPricesItem `json:"sku_prices"`
}
type SendHomeProductBatchUpdatePriceParam struct {
	// 修改价格的信息列表
	UpdateItems []UpdateItemsItem `json:"update_items"`
	// 操作视角，2和3时，store_id和outer_store_id选填
	StoreId *int64 `json:"store_id"`
	// 相对于store_id 是平台侧的门店id，outer_store_id是外部门店编码，使用上，选择其中一个，优先使用 store_id
	OuterStoreId *string `json:"outer_store_id"`
	// 操作视角。2:子品id操作子品;3:主品操作子品。注：普通商品需要传递2
	OperatePerspective int32 `json:"operate_perspective"`
}
