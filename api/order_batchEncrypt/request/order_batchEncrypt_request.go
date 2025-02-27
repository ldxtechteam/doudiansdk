package order_batchEncrypt_request

import (
	"encoding/json"
	"github.com/ldxtechteam/doudiansdk/api/order_batchEncrypt/response"
	"github.com/ldxtechteam/doudiansdk/core"
)

type OrderBatchEncryptRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderBatchEncryptParam
}

func (c *OrderBatchEncryptRequest) GetUrlPath() string {
	return "/order/batchEncrypt"
}

func New() *OrderBatchEncryptRequest {
	request := &OrderBatchEncryptRequest{
		Param: &OrderBatchEncryptParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *OrderBatchEncryptRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_batchEncrypt_response.OrderBatchEncryptResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_batchEncrypt_response.OrderBatchEncryptResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *OrderBatchEncryptRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *OrderBatchEncryptRequest) GetParams() *OrderBatchEncryptParam {
	return c.Param
}

type BatchEncryptListItem struct {
	// 明文
	PlainText string `json:"plain_text"`
	// 业务标识，value为抖音订单号；若是三方订单，可用三方订单号作为标识或自定义标识
	AuthId string `json:"auth_id"`
	// 是否支持密文索引
	IsSupportIndex bool `json:"is_support_index"`
	// 加密类型；1地址加密 2姓名加密 3电话加密
	SensitiveType int16 `json:"sensitive_type"`
}
type OrderBatchEncryptParam struct {
	// 待加密列表
	BatchEncryptList []BatchEncryptListItem `json:"batch_encrypt_list"`
	// 加密场景：OrderCode代表订单 AftersaleCode代表售后单
	SensitiveScene *string `json:"sensitive_scene"`
}
