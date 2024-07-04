package sendHome_delivery_syncSelfDeliveryLocation_request

import (
	"encoding/json"
	"github.com/ldxtechteam/doudiansdk/api/sendHome_delivery_syncSelfDeliveryLocation/response"
	"github.com/ldxtechteam/doudiansdk/core"
)

type SendHomeDeliverySyncSelfDeliveryLocationRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SendHomeDeliverySyncSelfDeliveryLocationParam
}

func (c *SendHomeDeliverySyncSelfDeliveryLocationRequest) GetUrlPath() string {
	return "/sendHome/delivery/syncSelfDeliveryLocation"
}

func New() *SendHomeDeliverySyncSelfDeliveryLocationRequest {
	request := &SendHomeDeliverySyncSelfDeliveryLocationRequest{
		Param: &SendHomeDeliverySyncSelfDeliveryLocationParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *SendHomeDeliverySyncSelfDeliveryLocationRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sendHome_delivery_syncSelfDeliveryLocation_response.SendHomeDeliverySyncSelfDeliveryLocationResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sendHome_delivery_syncSelfDeliveryLocation_response.SendHomeDeliverySyncSelfDeliveryLocationResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *SendHomeDeliverySyncSelfDeliveryLocationRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *SendHomeDeliverySyncSelfDeliveryLocationRequest) GetParams() *SendHomeDeliverySyncSelfDeliveryLocationParam {
	return c.Param
}

type SendHomeDeliverySyncSelfDeliveryLocationParam struct {
	// 抖音电商shop order单id
	OrderId int64 `json:"order_id"`
	// 骑手手机号类型，0是真实号，1是隐私号
	RiderPhoneType *int32 `json:"rider_phone_type"`
	// 骑手位置上报时间戳，传入骑手坐标时必填
	ReportTime int64 `json:"report_time"`
	// 坐标系 0 高德 1 百度
	RiderCoordinateType int32 `json:"rider_coordinate_type"`
	// 三方运力服务商，枚举： SF:顺丰同城 DADA:达达 FENG_NIAO:蜂鸟配送 MEI_TUAN:美团配送 SHAN_SONG:闪送 DIAN_WO_DA:点我达 UU:UU跑腿 XIN_TIAN_WENG:信天翁 SAN_LIU_WU:365跑腿 CAO_CAO:曹操跑腿 AI_PAO_TUI:爱跑腿 KUAI_PAO_ZHE:快跑者 JI_KE_KUAI_SONG:极客快送 TONG_DA:同达 SHENG_HUO_BAN_JING:生活半径 LIN_QU:临趣 QU_SONG:趣送 KUAI_FU_WU:快服务 CAI_NIAO_XIN_LIAN_MENG:菜鸟新联盟 FENG_XIAN_SHENG:风先生 LAI_DA_PEI_SONG:来答配送 HAO_JI_PAO_TUI:好急跑腿 SONG_GE_DONG_XI_PAO_TUI:送个东西跑腿 KAO_PU_SONG:靠谱送 KUAI_NAN_PAO_TUI:快男跑腿 GUO_XIAO_DI:裹小递 MERCHANT:商家自配送,使用此枚举需要提前和平台沟通
	DistributionCode string `json:"distribution_code"`
	// 三方运单号，与发货运单号保持一致
	DistributionDeliveryId string `json:"distribution_delivery_id"`
	// 骑手姓名； 骑手已接单/到店/取货/送达状态时必传；
	RiderName *string `json:"rider_name"`
	// 骑手手机号； 骑手已接单/到店/取货/送达状态时必传；
	RiderPhone *string `json:"rider_phone"`
	// 骑手经度；骑手已接单/到店/取货/送达状态时必传；高德坐标系
	RiderLongitude string `json:"rider_longitude"`
	// 骑手纬度；骑手已接单/到店/取货/送达状态时必传；高德坐标系
	RiderLatitude string `json:"rider_latitude"`
}
