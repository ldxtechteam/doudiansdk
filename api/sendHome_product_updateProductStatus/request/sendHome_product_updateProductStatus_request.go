package sendHome_product_updateProductStatus_request

import (
	"doudian.com/open/sdk_golang/api/sendHome_product_updateProductStatus/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SendHomeProductUpdateProductStatusRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SendHomeProductUpdateProductStatusParam 
}
func (c *SendHomeProductUpdateProductStatusRequest) GetUrlPath() string{
	return "/sendHome/product/updateProductStatus"
}


func New() *SendHomeProductUpdateProductStatusRequest{
	request := &SendHomeProductUpdateProductStatusRequest{
		Param: &SendHomeProductUpdateProductStatusParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SendHomeProductUpdateProductStatusRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sendHome_product_updateProductStatus_response.SendHomeProductUpdateProductStatusResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sendHome_product_updateProductStatus_response.SendHomeProductUpdateProductStatusResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SendHomeProductUpdateProductStatusRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SendHomeProductUpdateProductStatusRequest) GetParams() *SendHomeProductUpdateProductStatusParam{
	return c.Param
}


type SendHomeProductUpdateProductStatusParam struct {
	// 抖店门店id
	StoreId *int64 `json:"store_id"`
	// 外部门店编码，store_id和outer_store_id二选一传输 todo 对外文档详细描述
	OuterStoreId *string `json:"outer_store_id"`
	// 上下架操作类型，1下架、2上架
	OperateType int32 `json:"operate_type"`
	// 商品ID和外部商品id二选一传输，优先选择product_id
	ProductId *int64 `json:"product_id"`
	// 外部商品id，outer_product_id操作子品必传store_id/outer_store_id
	OuterProductId *string `json:"outer_product_id"`
	// 操作视角。1:主品id操作主品;2:子品id操作子品;3:主品操作子品
	OperatePerspective int32 `json:"operate_perspective"`
}
