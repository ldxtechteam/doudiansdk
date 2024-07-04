package rights_info_request

import (
	"encoding/json"
	"github.com/ldxtechteam/doudiansdk/api/rights_info/response"
	"github.com/ldxtechteam/doudiansdk/core"
)

type RightsInfoRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *RightsInfoParam
}

func (c *RightsInfoRequest) GetUrlPath() string {
	return "/rights/info"
}

func New() *RightsInfoRequest {
	request := &RightsInfoRequest{
		Param: &RightsInfoParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *RightsInfoRequest) Execute(accessToken *doudian_sdk.AccessToken) (*rights_info_response.RightsInfoResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &rights_info_response.RightsInfoResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *RightsInfoRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *RightsInfoRequest) GetParams() *RightsInfoParam {
	return c.Param
}

type RightsInfoParam struct {
	// 授权主体类型，不传参默认查询店铺，2-ebill用户；3-供应商
	BizType *int32 `json:"biz_type"`
	// 外部业务ID(模板市场为模板code)，非必传，不可与service_id同传
	OuterBizId string `json:"outer_biz_id"`
	// 抖店服务市场服务ID，非必传，不可与outer_biz_id同传
	ServiceId int64 `json:"service_id"`
}
