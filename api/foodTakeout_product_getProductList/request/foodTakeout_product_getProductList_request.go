package foodTakeout_product_getProductList_request

import (
	"encoding/json"
	"github.com/ldxtechteam/doudiansdk/api/foodTakeout_product_getProductList/response"
	"github.com/ldxtechteam/doudiansdk/core"
)

type FoodTakeoutProductGetProductListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *FoodTakeoutProductGetProductListParam
}

func (c *FoodTakeoutProductGetProductListRequest) GetUrlPath() string {
	return "/foodTakeout/product/getProductList"
}

func New() *FoodTakeoutProductGetProductListRequest {
	request := &FoodTakeoutProductGetProductListRequest{
		Param: &FoodTakeoutProductGetProductListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *FoodTakeoutProductGetProductListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*foodTakeout_product_getProductList_response.FoodTakeoutProductGetProductListResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &foodTakeout_product_getProductList_response.FoodTakeoutProductGetProductListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *FoodTakeoutProductGetProductListRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *FoodTakeoutProductGetProductListRequest) GetParams() *FoodTakeoutProductGetProductListParam {
	return c.Param
}

type FoodTakeoutProductGetProductListParam struct {
	// 抖店门店id
	StoreId *int64 `json:"store_id"`
	// 外部门店编码，查询门店品时，store_id和outer_store_id二选一传输
	OuterStoreId *string `json:"outer_store_id"`
	// 创建时间时间戳结束时间
	CreateEndTime *int64 `json:"create_end_time"`
	// 更新时间时间戳结束时间
	UpdateEndTime *int64 `json:"update_end_time"`
	// 商品名称，支持模糊查询
	ProductNameFuzzy *string `json:"product_name_fuzzy"`
	// 第几页(从1开始)
	PageNo int64 `json:"page_no"`
	// 每页条数,最大为50
	PageSize int64 `json:"page_size"`
	// 创建时间时间戳开始时间
	CreateStartTime *int64 `json:"create_start_time"`
	// 更新时间时间戳开始时间
	UpdateStartTime *int64 `json:"update_start_time"`
	// 底层接口默认创建时间排序 1-创建时间,2-更新时间
	SortType *int32 `json:"sort_type"`
	// 商品在店铺中状态：0售卖中，1已下架，2已封禁，3草稿箱，4审核中，5审核驳回；对应抖店列表tab
	ProductStatus int32 `json:"product_status"`
}
