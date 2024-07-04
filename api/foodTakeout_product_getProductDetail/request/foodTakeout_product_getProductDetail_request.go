package foodTakeout_product_getProductDetail_request

import (
	"encoding/json"
)

type FoodTakeoutProductGetProductDetailRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *FoodTakeoutProductGetProductDetailParam
}

func (c *FoodTakeoutProductGetProductDetailRequest) GetUrlPath() string {
	return "/foodTakeout/product/getProductDetail"
}

func New() *FoodTakeoutProductGetProductDetailRequest {
	request := &FoodTakeoutProductGetProductDetailRequest{
		Param: &FoodTakeoutProductGetProductDetailParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *FoodTakeoutProductGetProductDetailRequest) Execute(accessToken *doudian_sdk.AccessToken) (*foodTakeout_product_getProductDetail_response.FoodTakeoutProductGetProductDetailResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &foodTakeout_product_getProductDetail_response.FoodTakeoutProductGetProductDetailResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *FoodTakeoutProductGetProductDetailRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *FoodTakeoutProductGetProductDetailRequest) GetParams() *FoodTakeoutProductGetProductDetailParam {
	return c.Param
}

type ProductListItem struct {
	// 商品ID和外部商品id二选一传输，优先选择product_id
	ProductId *int64 `json:"product_id"`
	// 外部商品id
	OuterProductId *string `json:"outer_product_id"`
}
type FoodTakeoutProductGetProductDetailParam struct {
	// 商品查询列表
	ProductList []ProductListItem `json:"product_list"`
	// 抖店门店id 查询子品时需传入
	StoreId *int64 `json:"store_id"`
	// 外部门店编码，store_id和outer_store_id二选一传输
	OuterStoreId *string `json:"outer_store_id"`
	// 操作视角。1:主品id操作主品;2:子品id操作子品;3:主品操作子品
	OperatePerspective int32 `json:"operate_perspective"`
}
