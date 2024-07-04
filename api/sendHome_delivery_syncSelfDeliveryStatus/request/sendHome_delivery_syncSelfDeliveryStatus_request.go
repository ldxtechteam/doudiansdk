package sendHome_delivery_syncSelfDeliveryStatus_request

import (
	"encoding/json"
	"github.com/ldxtechteam/doudiansdk/api/sendHome_delivery_syncSelfDeliveryStatus/response"
	"github.com/ldxtechteam/doudiansdk/core"
)

type SendHomeDeliverySyncSelfDeliveryStatusRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SendHomeDeliverySyncSelfDeliveryStatusParam
}

func (c *SendHomeDeliverySyncSelfDeliveryStatusRequest) GetUrlPath() string {
	return "/sendHome/delivery/syncSelfDeliveryStatus"
}

func New() *SendHomeDeliverySyncSelfDeliveryStatusRequest {
	request := &SendHomeDeliverySyncSelfDeliveryStatusRequest{
		Param: &SendHomeDeliverySyncSelfDeliveryStatusParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *SendHomeDeliverySyncSelfDeliveryStatusRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sendHome_delivery_syncSelfDeliveryStatus_response.SendHomeDeliverySyncSelfDeliveryStatusResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sendHome_delivery_syncSelfDeliveryStatus_response.SendHomeDeliverySyncSelfDeliveryStatusResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *SendHomeDeliverySyncSelfDeliveryStatusRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *SendHomeDeliverySyncSelfDeliveryStatusRequest) GetParams() *SendHomeDeliverySyncSelfDeliveryStatusParam {
	return c.Param
}

type SendHomeDeliverySyncSelfDeliveryStatusParam struct {
	// 是否需要重新发货，默认为否；当订单发生变更时商家自行传入是否需要重新发货
	NeedResend *bool `json:"need_resend"`
	// 三方运力服务商，枚举： SF:顺丰同城 DADA:达达 FENG_NIAO:蜂鸟配送 MEI_TUAN:美团配送 SHAN_SONG:闪送 DIAN_WO_DA:点我达 UU:UU跑腿 XIN_TIAN_WENG:信天翁 SAN_LIU_WU:365跑腿 CAO_CAO:曹操跑腿 AI_PAO_TUI:爱跑腿 KUAI_PAO_ZHE:快跑者 JI_KE_KUAI_SONG:极客快送 TONG_DA:同达 SHENG_HUO_BAN_JING:生活半径 LIN_QU:临趣 QU_SONG:趣送 KUAI_FU_WU:快服务 CAI_NIAO_XIN_LIAN_MENG:菜鸟新联盟 FENG_XIAN_SHENG:风先生 LAI_DA_PEI_SONG:来答配送 HAO_JI_PAO_TUI:好急跑腿 SONG_GE_DONG_XI_PAO_TUI:送个东西跑腿 KAO_PU_SONG:靠谱送 KUAI_NAN_PAO_TUI:快男跑腿 GUO_XIAO_DI:裹小递 MERCHANT:商家自配送,使用此枚举需要提前和平台沟通
	DistributionCode string `json:"distribution_code"`
	// 三方运单号
	DistributionDeliveryId string `json:"distribution_delivery_id"`
	// 状态更新时间戳
	UpdateTime int64 `json:"update_time"`
	// todo 这个接口code支持哪些取消编码，枚举： 201:超出配送范围 202:骑手原因无法完成配送 203:运力紧张无可接单骑手 204:骑手联系不上收货人 206:物品超长超重无法完成配送 208:商家未营业 209:商家无法出餐 210:商家地址错误 211:用户拒收 222:用户取消订单 200:其他原因
	CancelCode *int32 `json:"cancel_code"`
	// 坐标系 0 高德 1 百度
	RiderCoordinateType *int32 `json:"rider_coordinate_type"`
	// 骑手姓名； 骑手已接单/到店/取货/送达状态时必传；
	RiderName *string `json:"rider_name"`
	// 骑手手机号类型，0是真实号，1是隐私号
	RiderPhoneType *int32 `json:"rider_phone_type"`
	// 骑手经度；骑手已接单/到店/取货/送达状态时必传；高德坐标系
	RiderLongitude *string `json:"rider_longitude"`
	// 骑手纬度；骑手已接单/到店/取货/送达状态时必传；高德坐标系
	RiderLatitude *string `json:"rider_latitude"`
	// 骑手位置上报时间戳，传入骑手坐标时必填
	ReportTime *int64 `json:"report_time"`
	// 取消原因
	CancelReason *string `json:"cancel_reason"`
	// 抖音电商shop order单id
	OrderId int64 `json:"order_id"`
	// 骑手手机号； 骑手已接单/到店/取货/送达状态时必传；
	RiderPhone *string `json:"rider_phone"`
	// 运力状态运力状态，枚举： 101:待骑手接单 102:骑手已接单 103:骑手已到店 104:骑手已取货 200:已送达 300:已取消，枚举： 101:待骑手接单 102:骑手已接单 103:骑手已到店 104:骑手已取货 200:已送达 300:已取消
	Status int32 `json:"status"`
}
