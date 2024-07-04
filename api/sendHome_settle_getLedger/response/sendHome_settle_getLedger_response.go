package sendHome_settle_getLedger_response

type SendHomeSettleGetLedgerResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SendHomeSettleGetLedgerData `json:"data"`
}
type LedgerOrderItem struct {
	// 混资券
	MixedCouponShopCostApportAmount int64 `json:"mixed_coupon_shop_cost_apport_amount"`
	// 总公司收入（分）
	MainIncome int64 `json:"main_income"`
	// 商家全资券金额（分）
	ShopCouponAmount int64 `json:"shop_coupon_amount"`
	// 商品直降金额（分）
	ShopDeductionAmount int64 `json:"shop_deduction_amount"`
	// 打包费（分）
	PackingAmount int64 `json:"packing_amount"`
	// 平台商品补贴金额（分）
	PromotionAmount int64 `json:"promotion_amount"`
	// 达人补贴金额（分）
	AuthorCouponAmount int64 `json:"author_coupon_amount"`
	// 支出合计（分）
	OutcomeTotalAmount int64 `json:"outcome_total_amount"`
	// 运费（分）
	PostAmount int64 `json:"post_amount"`
	// 抖音月付营销补贴金额（分）
	ZrPayPromotionAmount int64 `json:"zr_pay_promotion_amount"`
	// 商家支出达人的佣金（分）
	AuthorCommision int64 `json:"author_commision"`
	// 免佣金标识
	FreeCommissionFlag int64 `json:"free_commission_flag"`
	// 管理单元收入（分）
	UnitIncome int64 `json:"unit_income"`
	// 商品总价（分）
	ProductPriceTotalAmount int64 `json:"product_price_total_amount"`
	// 用户以旧换新购买商品，旧机扣款，由二手商出资（分）
	RecyclerAmount int64 `json:"recycler_amount"`
	// 商家支出团长的佣金（分）
	ColonelCommission int64 `json:"colonel_commission"`
	// 抖音支付补贴金额（不含银行补贴金额）（分）
	ZtPayPromotionAmount int64 `json:"zt_pay_promotion_amount"`
	// 用户支付时，银行出资的补贴金额，算作商家的收入（分）
	BankPayPromotionAmount int64 `json:"bank_pay_promotion_amount"`
	// 收入合计（分）
	IncomeTotalAmount int64 `json:"income_total_amount"`
	// 商家支出都可的佣金（分）
	DouCustomerCommission int64 `json:"dou_customer_commission"`
	// 支付金额
	PayAmount int64 `json:"pay_amount"`
	// 平台运费补贴金额（分）
	PostPromotionAmount int64 `json:"post_promotion_amount"`
	// 商家支付服务商的佣金（分）
	PartnerCommission int64 `json:"partner_commission"`
	// 结算时间
	SettleTime string `json:"settle_time"`
	// 分店收入（分）
	StoreIncome int64 `json:"store_income"`
	// 平台服务费=（用户实付商品金额+用户实付运费及打包费金额+达人补贴金额+抖音支付补贴金额+抖音月付营销补贴金额+银行营销补贴金额+以旧换新抵扣金额）*平台服务费率*（-1）（分）
	PlatformServiceFee int64 `json:"platform_service_fee"`
	// 平台服务费率
	PlatformServiceFeeRate int64 `json:"platform_service_fee_rate"`
	// 结算状态1. 待结算 2. 已结算 3. 部分退款（结算前）4. 部分退款（结算后）5. 全额退款（结算前）6. 全额退款（结算后）
	OrderSettleStatus int64 `json:"order_settle_status"`
}
type LedgerDeliveryItem struct {
	// 付款方 0子店 1总店 2管理单元
	PayerRole int64 `json:"payer_role"`
	// 平台配送履约服务费
	ChargeFeeAmount int64 `json:"charge_fee_amount"`
	// 配送取消费
	CancelFeeAmount int64 `json:"cancel_fee_amount"`
	// 结算状态 0待结算 1结算中 2结算成功 3结算失败 4冻结（手动流转，不参与自动结算）
	DeliverySettleStatus int64 `json:"delivery_settle_status"`
}
type Income struct {
	// 总公司收入
	MainIncome int64 `json:"main_income"`
	// 管理单元收入
	UnitIncome int64 `json:"unit_income"`
	// 门店收入
	StoreIncome int64 `json:"store_income"`
	// 总收入
	TotalIncome int64 `json:"total_income"`
}
type LedgerItem struct {
	// 货款账单
	LedgerOrder []LedgerOrderItem `json:"ledger_order"`
	// 平台配送费账单
	LedgerDelivery []LedgerDeliveryItem `json:"ledger_delivery"`
	// 收入汇总
	Income *Income `json:"income"`
	// 订单id
	OrderId string `json:"order_id"`
}
type SendHomeSettleGetLedgerData struct {
	// 账单,key为订单ID,value为具体账单
	Ledger []LedgerItem `json:"ledger"`
}
