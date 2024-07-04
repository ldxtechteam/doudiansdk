package sendHome_afterSale_closeRefundMsg_response

type SendHomeAfterSaleCloseRefundMsgResponse struct {
	doudian_sdk.BaseDoudianOpSpiResponse
	Data *SendHomeAfterSaleCloseRefundMsgData `json:"data"`
}

func (c *SendHomeAfterSaleCloseRefundMsgResponse) GetData() interface{} {
	return c.Data
}

type SendHomeAfterSaleCloseRefundMsgData struct {
	// 业务错误响应码。0代表成功。 【通用】： 0：业务处理成功 100001：验签失败 【消息领域错误】： 200001：未订阅/未实现/不支持的消息类型「也可以直接返回0:成功」 200002：频控，在一定的时间内用户发送了太多的请求，即超出了“频次限制” 200003：已收到该消息，重复消息 【服务商业务错误(3开头)业务错误。处理流程正常，但是业务逻辑不自洽导致的异常返回】： 300001：未知业务错误「最好将错误原因放在msg内」 【参数错误(4开头)】： 400001：未接入的商家/不支持的商家/商家不存在 400002：消息结构不符合预期「最好将错误的结构放在msg内」 400003：消息字段(值)不符合预期「最好将错误的消息字段放在msg内」 【服务商系统错误(5开头)处理流程正常，但是由于系统错误导致的异常返回】： 500001：未知系统错误「最好将错误原因放在msg内」
	Code int64 `json:"code"`
	// 异常附加信息
	Message string `json:"message"`
}
