package foodTakeout_order_getOrder_response

import (
	"github.com/ldxtechteam/doudiansdk/core"
)

type FoodTakeoutOrderGetOrderResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *FoodTakeoutOrderGetOrderData `json:"data"`
}
type ShopInfo struct {
	// 门店id
	StoreId int64 `json:"store_id"`
	// 门店名称
	StoreName string `json:"store_name"`
	// 总户id
	ShopId int64 `json:"shop_id"`
}
type SkuSpecsItem struct {
	// 规格值 500
	Value string `json:"value"`
	// 规格单位名称 克
	Name string `json:"name"`
}
type PropertyDetailsItem struct {
	// 属性名称
	Name string `json:"name"`
}
type AfterSaleInfosItem struct {
	// 售后id
	AfterSaleId string `json:"after_sale_id"`
	// 退款状态
	RefundStatus int64 `json:"refund_status"`
	// 售后状态
	AfterSaleStatus int64 `json:"after_sale_status"`
	// 退款金额
	RefundAmount int64 `json:"refund_amount"`
	// 退款时间戳
	RefundTimeStamp int64 `json:"refund_time_stamp"`
}
type FoodTakeoutOrderGetOrderData struct {
	// 订单信息
	OpenOrderDetail []OpenOrderDetailItem `json:"open_order_detail"`
}
type UserCoordinate struct {
	// 经度
	Longitude string `json:"longitude"`
	// 维度
	Latitude string `json:"latitude"`
}
type DeliveryInfo struct {
	// 三方运单id
	DistributionDeliveryId string `json:"distribution_delivery_id"`
	// 三方物流公司名称
	DistributionCode string `json:"distribution_code"`
	// 运单id
	DeliveryId string `json:"delivery_id"`
	// 骑手名称
	RiderName string `json:"rider_name"`
	// 骑手手机号
	RiderPhone string `json:"rider_phone"`
	// 运费实付（商家） //c配
	PostAmount int64 `json:"post_amount"`
	// 运费优惠(商家)
	PostPromotionAmount int64 `json:"post_promotion_amount"`
	// 配送方式 //todo: 只区分平台配送和自配送
	DeliveryType int64 `json:"delivery_type"`
}
type IngredientsItem struct {
	// 商品id
	ProductId int64 `json:"product_id"`
	// 商品sku id
	SkuId int64 `json:"sku_id"`
	// 外部商品sku id
	OutSkuId string `json:"out_sku_id"`
	// 名称
	SkuName string `json:"sku_name"`
	// 图片
	ImgUrl string `json:"img_url"`
	// sku价格
	SkuPrice int64 `json:"sku_price"`
	// 同一个sku单下的item单数量,如果购买了2份奶茶,每份奶茶都加了一份珍珠,此处为2
	Nums int64 `json:"nums"`
	// 订单sku单id
	SkuOrderId string `json:"sku_order_id"`
}
type SkuOrderListItem struct {
	// 商品价格
	PayPrice int64 `json:"pay_price"`
	// sku单
	Sku []SkuItem `json:"sku"`
}
type BuyInfo struct {
	// 真实手机号
	Phone string `json:"phone"`
	// 虚拟号
	SecretNumber string `json:"secret_number"`
}
type OrderBase struct {
	// 支付优惠
	PayDiscountAmount int64 `json:"pay_discount_amount"`
	// 用户备注
	Remark string `json:"remark"`
	// 运费优惠(用户)
	PostPromotionAmount int64 `json:"post_promotion_amount"`
	// 用户预计送达时间开始时间戳
	UserExpectTimeBegin int64 `json:"user_expect_time_begin"`
	// 总打包费
	PackingPrice int64 `json:"packing_price"`
	// 流水号
	ShopNumber string `json:"shop_number"`
	// 订单id
	OrderId string `json:"order_id"`
	// 商家预计收入
	Income int64 `json:"income"`
	// 时间戳
	PayTimestamp int64 `json:"pay_timestamp"`
	// 是否为预约单
	IsBook int64 `json:"is_book"`
	// 运费实付（用户） //c配
	PostAmount int64 `json:"post_amount"`
	// 订单原始价
	OrderAmount int64 `json:"order_amount"`
	// 订单实付价
	PayAmount int64 `json:"pay_amount"`
	// 创建时间戳
	CreateTimestamp int64 `json:"create_timestamp"`
	// 订单状态
	OrderStatus int32 `json:"order_status"`
	// 用户预计送达时间结束时间戳
	UserExpectTimeEnd int64 `json:"user_expect_time_end"`
}
type ReceiveInfo struct {
	// 详细地址自配送有值平台配送没有值
	DetailAddress string `json:"detail_address"`
	// 手机号(加密)
	MaskPhone string `json:"mask_phone"`
	// 真实手机号
	Phone string `json:"phone"`
	// 收货人
	Name string `json:"name"`
	// 省
	Province string `json:"province"`
	// 市
	City string `json:"city"`
	// 区
	Town string `json:"town"`
	// 街道
	Street string `json:"street"`
	// 隐私号
	SecretNumber string `json:"secret_number"`
	// 用户经纬度
	UserCoordinate *UserCoordinate `json:"user_coordinate"`
}
type OpenOrderDetailItem struct {
	// 收货人信息
	ReceiveInfo *ReceiveInfo `json:"receive_info"`
	// 优惠信息
	Discount *Discount `json:"discount"`
	// 购买人信息
	BuyInfo *BuyInfo `json:"buy_info"`
	// 订单基础信息
	OrderBase *OrderBase `json:"order_base"`
	// 运单信息
	DeliveryInfo *DeliveryInfo `json:"delivery_info"`
	// 门店信息
	ShopInfo *ShopInfo `json:"shop_info"`
	// 总的商品信息
	SkuOrderList []SkuOrderListItem `json:"sku_order_list"`
	// 售后信息
	AfterSaleInfos []AfterSaleInfosItem `json:"after_sale_infos"`
}
type Discount struct {
	// 商家优惠金额
	ShopDiscountAmount int64 `json:"shop_discount_amount"`
	// 平台优惠金额
	PlatformDiscountAmount int64 `json:"platform_discount_amount"`
	// 达人优惠金额
	TalentDiscountAmount int64 `json:"talent_discount_amount"`
	// 优惠金额总额(不包含支付优惠、运费优惠)
	Amount int64 `json:"amount"`
}
type SkuItem struct {
	// 外部商品id
	OutProductId string `json:"out_product_id"`
	// 名称
	SkuName string `json:"sku_name"`
	// 小料单
	Ingredients []IngredientsItem `json:"ingredients"`
	// 规格名
	SkuSpecs []SkuSpecsItem `json:"sku_specs"`
	// 订单sku单id
	SkuOrderId string `json:"sku_order_id"`
	// 商品id
	ProductId int64 `json:"product_id"`
	// 商品sku id
	SkuId int64 `json:"sku_id"`
	// 外部商品sku id
	OutSkuId string `json:"out_sku_id"`
	// 图片
	ImgUrl string `json:"img_url"`
	// sku价格
	SkuPrice int64 `json:"sku_price"`
	// sku的数量
	Nums int64 `json:"nums"`
	// 属性名
	PropertyDetails []PropertyDetailsItem `json:"property_details"`
}
