package sendHome_afterSale_operateAfterSale_request

import (
	"encoding/json"
)

type SendHomeAfterSaleOperateAfterSaleRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SendHomeAfterSaleOperateAfterSaleParam
}

func (c *SendHomeAfterSaleOperateAfterSaleRequest) GetUrlPath() string {
	return "/sendHome/afterSale/operateAfterSale"
}

func New() *SendHomeAfterSaleOperateAfterSaleRequest {
	request := &SendHomeAfterSaleOperateAfterSaleRequest{
		Param: &SendHomeAfterSaleOperateAfterSaleParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *SendHomeAfterSaleOperateAfterSaleRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sendHome_afterSale_operateAfterSale_response.SendHomeAfterSaleOperateAfterSaleResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sendHome_afterSale_operateAfterSale_response.SendHomeAfterSaleOperateAfterSaleResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *SendHomeAfterSaleOperateAfterSaleRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *SendHomeAfterSaleOperateAfterSaleRequest) GetParams() *SendHomeAfterSaleOperateAfterSaleParam {
	return c.Param
}

type EvidenceItem struct {
	// 操作凭证类型,凭证类型，1:图片，2:视频(暂不支持)，3:音频(暂不支持展示)，4:文字(暂不支持展示)。
	Type *int64 `json:"type"`
	// 操作凭证url
	Url *string `json:"url"`
	// 凭证描述
	Desc *string `json:"desc"`
}
type SendHomeAfterSaleOperateAfterSaleParam struct {
	// 拒绝原因，拒绝操作必填
	Reason *string `json:"reason"`
	// 操作评论，拒绝操作必填
	Remark *string `json:"remark"`
	// 操作凭证
	Evidence []EvidenceItem `json:"evidence"`
	// 门店ID
	StoreId int64 `json:"store_id"`
	// 操作类型,201同意仅退款202拒绝仅退款
	OperateType int32 `json:"operate_type"`
	// 拒绝原因code
	RejectReasonCode *int64 `json:"reject_reason_code"`
	// 售后单ID
	AfterSaleId int64 `json:"after_sale_id"`
}
