package sendHome_afterSale_getRejectReasonCodeList_request

import (
	"doudian.com/open/sdk_golang/api/sendHome_afterSale_getRejectReasonCodeList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SendHomeAfterSaleGetRejectReasonCodeListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SendHomeAfterSaleGetRejectReasonCodeListParam 
}
func (c *SendHomeAfterSaleGetRejectReasonCodeListRequest) GetUrlPath() string{
	return "/sendHome/afterSale/getRejectReasonCodeList"
}


func New() *SendHomeAfterSaleGetRejectReasonCodeListRequest{
	request := &SendHomeAfterSaleGetRejectReasonCodeListRequest{
		Param: &SendHomeAfterSaleGetRejectReasonCodeListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SendHomeAfterSaleGetRejectReasonCodeListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sendHome_afterSale_getRejectReasonCodeList_response.SendHomeAfterSaleGetRejectReasonCodeListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sendHome_afterSale_getRejectReasonCodeList_response.SendHomeAfterSaleGetRejectReasonCodeListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SendHomeAfterSaleGetRejectReasonCodeListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SendHomeAfterSaleGetRejectReasonCodeListRequest) GetParams() *SendHomeAfterSaleGetRejectReasonCodeListParam{
	return c.Param
}


type SendHomeAfterSaleGetRejectReasonCodeListParam struct {
	// 门店ID
	StoreId *int64 `json:"store_id"`
	// 售后单ID
	AfterSaleId *int64 `json:"after_sale_id"`
}
