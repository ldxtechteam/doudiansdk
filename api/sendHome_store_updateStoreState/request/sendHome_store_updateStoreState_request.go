package sendHome_store_updateStoreState_request

import (
	"encoding/json"
	"github.com/ldxtechteam/doudiansdk/api/sendHome_store_updateStoreState/response"
	"github.com/ldxtechteam/doudiansdk/core"
)

type SendHomeStoreUpdateStoreStateRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SendHomeStoreUpdateStoreStateParam
}

func (c *SendHomeStoreUpdateStoreStateRequest) GetUrlPath() string {
	return "/sendHome/store/updateStoreState"
}

func New() *SendHomeStoreUpdateStoreStateRequest {
	request := &SendHomeStoreUpdateStoreStateRequest{
		Param: &SendHomeStoreUpdateStoreStateParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *SendHomeStoreUpdateStoreStateRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sendHome_store_updateStoreState_response.SendHomeStoreUpdateStoreStateResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sendHome_store_updateStoreState_response.SendHomeStoreUpdateStoreStateResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *SendHomeStoreUpdateStoreStateRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *SendHomeStoreUpdateStoreStateRequest) GetParams() *SendHomeStoreUpdateStoreStateParam {
	return c.Param
}

type SendHomeStoreUpdateStoreStateParam struct {
	// 门店id
	StoreId *int64 `json:"store_id"`
	// 外部门店编码，store_id和outer_store_id二选一传输
	OuterStoreId *string `json:"outer_store_id"`
	// 接单状态操作类型，1恢复、2关闭
	OperateType int32 `json:"operate_type"`
	// 恢复/关闭营业原因
	Reason *string `json:"reason"`
}
