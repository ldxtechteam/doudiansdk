package foodTakeout_product_getProductList_response

type FoodTakeoutProductGetProductListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *FoodTakeoutProductGetProductListData `json:"data"`
}
type CategoryDetail struct {
	// 四级类目id
	FourthCid int64 `json:"fourth_cid"`
	// 一级类目名
	FirstCname string `json:"first_cname"`
	// 二级类目名
	SecondCname string `json:"second_cname"`
	// 三级类目名
	ThirdCname string `json:"third_cname"`
	// 四级类目名
	FourthCname string `json:"fourth_cname"`
	// 一级类目id
	FirstCid int64 `json:"first_cid"`
	// 二级类目id
	SecondCid int64 `json:"second_cid"`
	// 三级类目id
	ThirdCid int64 `json:"third_cid"`
}
type PicListItem struct {
	// 图片url
	PicUrl string `json:"pic_url"`
}
type PropertyDetailsItem struct {
	// 属性名称
	Name string `json:"name"`
}
type ProductIngredientDetailsItem struct {
	// 组套商品id
	ProductId int64 `json:"product_id"`
	// 组套skuid 若绑定类型为商品，则本值为0
	SkuId int64 `json:"sku_id"`
	// 组套商品名称
	ProductName string `json:"product_name"`
	// 组套价格
	Price int64 `json:"price"`
	// 组套库存
	Stock int64 `json:"stock"`
}
type ValueListItem struct {
	// 属性值单位
	Value string `json:"value"`
	// 属性值名称
	Unit string `json:"unit"`
}
type SkusItem struct {
	// 规格属性
	QuantityInfo *QuantityInfo `json:"quantity_info"`
	// 抖音skuId
	SkuId int64 `json:"sku_id"`
	// 第三方skuid
	OuterSkuId string `json:"outer_sku_id"`
	// sku名称
	SkuName string `json:"sku_name"`
	// sku状态 true上架 false下架
	SkuStatus bool `json:"sku_status"`
	// 实际支付价格，单位：分
	ActualPrice int64 `json:"actual_price"`
	// 库存信息
	Stock *Stock `json:"stock"`
}
type ProductIngredientItem struct {
	// 组套包id
	PackageId int64 `json:"package_id"`
	// 组套包名称
	PackageName string `json:"package_name"`
	// 组套列表
	ProductIngredientDetails []ProductIngredientDetailsItem `json:"product_ingredient_details"`
}
type ProductListItem struct {
	// 外部商家编码，商家自定义字段，支持最多 255个字符
	OuterProductId string `json:"outer_product_id"`
	// 商品标题，规则：至少输入8个字（16个字符）以上~输入30个字（60个字符）以内。；标题不规范会引起商品下架，影响您的正常销售，详见商品发布规范：https://school.jinritemai.com/doudian/web/article/101800?from=shop_article
	ProductName string `json:"product_name"`
	// 商品审核状态: 1-未提交；2-待审核；3-审核通过；4-审核未通过；5-封禁；7-审核通过待上架；详见商品状态机：https://op.jinritemai.com/docs/question-docs/92/2070
	CheckStatus int64 `json:"check_status"`
	// 总部id
	ShopId int64 `json:"shop_id"`
	// 菜品原料详情
	MaterialInfo []MaterialInfoItem `json:"material_info"`
	// 抖店门店id，只有子品有对应信息，主品为空
	StoreId int64 `json:"store_id"`
	// 商品白底图
	ImgUrl string `json:"img_url"`
	// sku列表
	Skus []SkusItem `json:"skus"`
	// 是否绑定商家连锁菜品: 普通品：False && 有门店id;正常主品：False && 没有门店id；正常子品：True&&有门店id
	IsBindMerchant bool `json:"is_bind_merchant"`
	// 绑定的商家连锁菜品id
	MainProductId int64 `json:"main_product_id"`
	// 草稿状态；0-无草稿,1-未提审,2-待审核,3-审核通过,4-审核未通过。详见：https://op.jinritemai.com/docs/question-docs/92/2070
	DraftStatus int64 `json:"draft_status"`
	// 商品在店铺中状态: 0-在线；1-下线；2-删除；详见商品状态机：https://op.jinritemai.com/docs/question-docs/92/2070
	OnlineStatus int64 `json:"online_status"`
	// 创建时间时间戳
	CreateTime int64 `json:"create_time"`
	// 退款规则
	RefundTips string `json:"refund_tips"`
	// 商品ID，抖店系统生成，店铺下唯一；长度19位。
	ProductId int64 `json:"product_id"`
	// 类目详情；商品类目id可使用【/shop/getShopCategory】查询，叶子类目到第三级或者第四级
	CategoryDetail *CategoryDetail `json:"category_detail"`
	// 商品分组信息
	ProductGroups []ProductGroupsItem `json:"product_groups"`
	// 外部门店编码，只有子品有对应信息，主品为空
	OuterStoreId string `json:"outer_store_id"`
	// 更新时间时间戳
	UpdateTime int64 `json:"update_time"`
	// 商品主图；最多支持5张图片；仅支持png，jpg，jpeg格式，宽高比例为1:1（至少600*600px），大小5M内
	PicList []PicListItem `json:"pic_list"`
	// 商品描述文案
	ProductDescText string `json:"product_desc_text"`
	// 定制属性列表
	CustomPropertyGroups []CustomPropertyGroupsItem `json:"custom_property_groups"`
	// 组套包信息，对应加料组列表
	ProductIngredient []ProductIngredientItem `json:"product_ingredient"`
}
type MaterialInfoItem struct {
	// 对应的值
	ValueList []ValueListItem `json:"value_list"`
	// 原料key信息
	Key string `json:"key"`
	// 原料key描述名称
	Name string `json:"name"`
}
type CustomPropertyGroupsItem struct {
	// 属性组列表
	GroupName string `json:"group_name"`
	// 属性列表
	PropertyDetails []PropertyDetailsItem `json:"property_details"`
}
type ProductGroupsItem struct {
	// 商品分组id
	GroupId int64 `json:"group_id"`
	// 商品分组名称
	GroupName string `json:"group_name"`
}
type FoodTakeoutProductGetProductListData struct {
	// 总数
	Total int32 `json:"total"`
	// 是否还能查询更多
	HasMore bool `json:"has_more"`
	// 菜品详情列表
	ProductList []ProductListItem `json:"product_list"`
}
type QuantityInfo struct {
	// 规格值 500
	Value int64 `json:"value"`
	// 规格单位名称 克
	UnitName string `json:"unit_name"`
	// 规格单位名称前缀，约
	UnitPrefix string `json:"unit_prefix"`
}
type Stock struct {
	// 现货库存
	StockNum int64 `json:"stock_num"`
}
