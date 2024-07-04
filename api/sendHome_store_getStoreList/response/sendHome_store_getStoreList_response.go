package sendHome_store_getStoreList_response

type SendHomeStoreGetStoreListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SendHomeStoreGetStoreListData `json:"data"`
}
type OpenTime struct {
	// key:星期(1:周一、2:周二...7:周日）val:时间(多段用空格分隔，例"08:30-11:00 13:00-17:00")
	DayMap map[int64]string `json:"day_map"`
}
type Store struct {
	// 省份code
	ProvinceCode string `json:"province_code"`
	// 区
	District string `json:"district"`
	// 联系方式
	Contact string `json:"contact"`
	// 绑定成功的时间戳，即state变成2的时间
	BoundTime int64 `json:"bound_time"`
	// 门店id
	StoreId int64 `json:"store_id"`
	// 门店所属的总部Id
	ShopId int64 `json:"shop_id"`
	// 省
	Province string `json:"province"`
	// 营业时间
	OpenTime *OpenTime `json:"open_time"`
	// 记录更新时间
	UpdateTime int64 `json:"update_time"`
	// 门店名称
	Name string `json:"name"`
	// 城市code
	CityCode string `json:"city_code"`
	// 当前是否在营业时间中，如果不在则是打烊中
	IsOpenNow bool `json:"is_open_now"`
	// 绑定状态 0初始化;1绑定中;2绑定成功;3:资质验证中;4账户认证中;5正常营业(c端可下单);6暂停营业
	State int64 `json:"state"`
	// 停业类型:1商家主动操作;2平台处罚;3清退。state = 6的时候该字段有意义
	SuspendType int64 `json:"suspend_type"`
	// 维度
	Latitude string `json:"latitude"`
	// 经度
	Longitude string `json:"longitude"`
	// 市
	City string `json:"city"`
	// 区code
	DistrictCode string `json:"district_code"`
	// 4级地址 街道/镇
	Town string `json:"town"`
	// 4级地址code
	TownCode string `json:"town_code"`
	// 详细地址
	Address string `json:"address"`
	// 记录创建时间
	CreateTime int64 `json:"create_time"`
	// 外部门店编码
	OuterStoreId string `json:"outer_store_id"`
}
type StoreDetailListItem struct {
	// 门店信息
	Store *Store `json:"store"`
}
type SendHomeStoreGetStoreListData struct {
	// 总数
	Total int32 `json:"total"`
	// 是否还能查询更多
	HasMore bool `json:"has_more"`
	// 门店详情列表
	StoreDetailList []StoreDetailListItem `json:"store_detail_list"`
}
