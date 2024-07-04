package sendHome_settle_getLedger_request

import (
	"encoding/json"
)

type SendHomeSettleGetLedgerRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SendHomeSettleGetLedgerParam
}

func (c *SendHomeSettleGetLedgerRequest) GetUrlPath() string {
	return "/sendHome/settle/getLedger"
}

func New() *SendHomeSettleGetLedgerRequest {
	request := &SendHomeSettleGetLedgerRequest{
		Param: &SendHomeSettleGetLedgerParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *SendHomeSettleGetLedgerRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sendHome_settle_getLedger_response.SendHomeSettleGetLedgerResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sendHome_settle_getLedger_response.SendHomeSettleGetLedgerResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *SendHomeSettleGetLedgerRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *SendHomeSettleGetLedgerRequest) GetParams() *SendHomeSettleGetLedgerParam {
	return c.Param
}

type SendHomeSettleGetLedgerParam struct {
	// 订单ID
	OrderIds []string `json:"order_ids"`
}
