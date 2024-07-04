package sendHome_store_getStoreList_request

import (
	"doudian.com/open/sdk_golang/api/sendHome_store_getStoreList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SendHomeStoreGetStoreListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SendHomeStoreGetStoreListParam 
}
func (c *SendHomeStoreGetStoreListRequest) GetUrlPath() string{
	return "/sendHome/store/getStoreList"
}


func New() *SendHomeStoreGetStoreListRequest{
	request := &SendHomeStoreGetStoreListRequest{
		Param: &SendHomeStoreGetStoreListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SendHomeStoreGetStoreListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sendHome_store_getStoreList_response.SendHomeStoreGetStoreListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sendHome_store_getStoreList_response.SendHomeStoreGetStoreListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SendHomeStoreGetStoreListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SendHomeStoreGetStoreListRequest) GetParams() *SendHomeStoreGetStoreListParam{
	return c.Param
}


type SendHomeStoreGetStoreListParam struct {
	// 门店id集合，用于查询门店信息，最大长度50
	StoreIdList []int64 `json:"store_id_list"`
	// 外部门店编码集合，开发者自定义，最大长度50
	OuterStoreIdList []string `json:"outer_store_id_list"`
	// 绑定状态 1绑定中;2绑定成功;3:资质验证中;4账户认证中;5正常营业(c端可下单) 6暂停营业
	State *int64 `json:"state"`
	// 门店名称，支持模糊查询
	StoreNameFuzzy *string `json:"store_name_fuzzy"`
	// 第几页(从1开始)
	PageNo int32 `json:"page_no"`
	// 每页条数,最大为50
	PageSize int32 `json:"page_size"`
}
