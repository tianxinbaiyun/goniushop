package models

type NsPlatformAdv struct {
	SlideSort  int    `orm:"column(slide_sort);null json:"slide_sort""`
	ClickNum   uint   `orm:"column(click_num)" description:"广告点击率" json:"click_num"`
	Background string `orm:"column(background);size(255)" description:"背景色" json:"background"`
	AdvCode    string `orm:"column(adv_code)" description:"广告代码" json:"adv_code"`

	AdPositionId uint32 `orm:"column(ap_id)" description:"广告位id" json:"ad_position_id"`
	Content      string `orm:"column(adv_title);size(255)" description:"广告内容描述" json:"content"`
	Id           int    `orm:"column(adv_id);auto" description:"广告自增标识编号" json:"id`
	ImageUrl     string `orm:"column(adv_image);size(1000)" description:"广告内容图片" json:"image_url"`
	Link         string `orm:"column(adv_url);null" json:"link"`
}

type NsPlatformAdvPosition struct {
	ApClass           uint16 `orm:"column(ap_class)" description:"广告类别：0图片1文字2幻灯3Flash" json:"ap_class"`
	ApDisplay         uint16 `orm:"column(ap_display)" description:"广告展示方式：0幻灯片1多广告展示2单广告展示" json:"ap_display"`
	IsUse             uint16 `orm:"column(is_use)" description:"广告位是否启用：0不启用1启用" json:"is_use"`
	ApPrice           int    `orm:"column(ap_price)" description:"广告位单价" json:"ap_price"`
	AdvNum            int    `orm:"column(adv_num)" description:"拥有的广告数" json:"adv_num"`
	ClickNum          int    `orm:"column(click_num)" description:"广告位点击率" json:"click_num"`
	DefaultContent    string `orm:"column(default_content);size(300);null" json:"default_content"`
	ApBackgroundColor string `orm:"column(ap_background_color);size(50)" description:"广告位背景色 默认白色" json:"ap_background_color"`
	Type              int    `orm:"column(type)" description:"广告位所在位置类型   1 pc端  2 手机端" json:"type"`
	InstanceId        int    `orm:"column(instance_id)" description:"店铺id" json:"instance_id"`
	IsDel             int    `orm:"column(is_del);null" json:"is_del"`
	ApKeyword         string `orm:"column(ap_keyword);size(100)" description:"广告位关键字" json:"ap_keyword"`

	Desc   string `orm:"column(ap_intro);size(255)" description:"广告位简介" json:"desc"`
	Height int    `orm:"column(ap_height)" description:"广告位高度" json:"height"`
	Id     int    `orm:"column(ap_id);auto" description:"广告位置id" json:"id"`
	Name   string `orm:"column(ap_name);size(100)" description:"广告位置名" json:"name"`
	Width  int    `orm:"column(ap_width)" description:"广告位宽度" json:"width"`
}

type NsMemberExpressAddress struct {
	Phone   string `orm:"column(phone);size(20)" description:"固定电话" json:"phone"`
	ZipCode string `orm:"column(zip_code);size(6)" description:"邮编" json:"zip_code"`
	Alias   string `orm:"column(alias);size(50)" description:"地址别名" json:"alias,omitempty"`

	Address    string `orm:"column(address);size(255)" description:"详细地址" json:"address"`
	CityId     int    `orm:"column(city)" description:"市" json:"city_id"`
	DistrictId int    `orm:"column(district)" description:"区县" json:"district_id"`
	Id         int    `orm:"column(id);auto" json:"id"`
	IsDefault  int    `orm:"column(is_default)" description:"默认收货地址" json:"is_default"`
	Mobile     string `orm:"column(mobile);size(11)" description:"手机" json:"mobile"`
	Name       string `orm:"column(consigner);size(255)" description:"收件人" json:"name"`
	ProvinceId int    `orm:"column(province)" description:"省" json:"province_id"`
	UserId     int    `orm:"column(uid)" description:"会员基本资料表ID" json:"user_id"`
}
type SysUserAdmin struct {
	IsAdmin     int    `orm:"column(is_admin)" description:"是否是系统管理员组"`
	AdminStatus int    `orm:"column(admin_status);null" description:"状态 默认为1"`
	Desc        string `orm:"column(desc);null" description:"附加信息"`
	AdminRoleId int    `orm:"column(group_id_array)" description:"系统用户组" json:"admin_role_id"`
	Id          int    `orm:"column(uid);pk" description:"user用户ID" json:"id"`
	Username    string `orm:"column(admin_name);size(50)" description:"用户姓名" json:"username"`
}

type NsAttribute struct {
	SpecIdArray string `orm:"column(spec_id_array);size(255)" description:"关联规格"`
	Sort        int    `orm:"column(sort);null"`
	CreateTime  int    `orm:"column(create_time);null" description:"创建时间"`
	ModifyTime  int    `orm:"column(modify_time);null" description:"修改时间"`

	Enabled int    `orm:"column(is_use)" description:"是否使用" json:"enabled"`
	Id      int    `orm:"column(attr_id);auto" description:"商品属性ID" json:"id"`
	Name    string `orm:"column(attr_name);size(255)" description:"属性名称" json:"name"`
}
type NsAttributeValue struct {
	AttrId    int    `orm:"column(attr_id)" description:"属性ID"`
	IsSearch  int    `orm:"column(is_search)" description:"是否使用"`
	Id        int    `orm:"column(attr_value_id);auto" description:"属性值ID" json:"id"`
	InputType int    `orm:"column(type)" description:"属性对应输入类型1.直接2.单选3.多选" json:"input_type"`
	Name      string `orm:"column(attr_value_name);size(50)" description:"属性值名称" json:"name"`
	SortOrder int    `orm:"column(sort)" description:"排序号" json:"sort_order"`
	Values    string `orm:"column(value);size(1000)" description:"属性对应相关数据" json:"values"`
}
type NsGoodsBrand struct {
	BrandInitial      string `orm:"column(brand_initial);size(1)" description:"品牌首字母" json:"brand_initial"`
	BrandCategoryName string `orm:"column(brand_category_name);size(50)" description:"类别名称" json:"brand_category_name"`
	BrandId           int    `orm:"column(brand_id);auto" description:"索引ID" json:"brand_id"`
	BrandRecommend    int    `orm:"column(brand_recommend)" description:"推荐，0为否，1为是，默认为0"  json:"brand_recommend"`
	BrandName         string `orm:"column(brand_name);size(100)" description:"品牌名称" json:"brand_name"`
	BrandPic          string `orm:"column(brand_pic);size(100)" description:"图片" json:"brand_pic"`
	Sort              int    `orm:"column(sort);null" json:"sort"`
}
type NsCart struct {
	SkuId        int    `orm:"column(sku_id)" description:"商品的skuid" json:"sku_id"`
	SkuName      string `orm:"column(sku_name);size(200)" description:"商品的sku名称" json:"sku_name,omitempty"`
	GoodsPicture int    `orm:"column(goods_picture)" description:"商品图片" json:"goods_picture"`
	BlId         int32  `orm:"column(bl_id)" description:"组合套装ID" json:"bl_id,omitempty"`

	GoodsId   int     `orm:"column(goods_id)" description:"商品id" json:"goods_id"`
	GoodsName string  `orm:"column(goods_name);size(200)" description:"商品名称" json:"goods_name"`
	Id        int     `orm:"column(cart_id);auto" description:"购物车id" json:"id"`
	Number    int     `orm:"column(num)" description:"购买商品数量" json:"number"`
	Price     float64 `orm:"column(price);digits(10);decimals(2)" description:"商品价格" json:"price"`
	UserId    int     `orm:"column(buyer_id)" description:"买家id" json:"user_id"`
}
type NsGoodsCategory struct {
	ShortName   string `orm:"column(short_name);size(50)" description:"商品分类简称 "`
	AttrId      int    `orm:"column(attr_id)" description:"关联商品类型ID"`
	AttrName    string `orm:"column(attr_name);size(255)" description:"关联类型名称"`
	Description string `orm:"column(description);size(255);null"`

	Id        int    `orm:"column(category_id);auto" json:"id"`
	ImgUrl    string `orm:"column(category_pic);size(255)" description:"商品分类图片" json:"img_url"`
	IsShow    int    `orm:"column(is_visible)" description:"是否显示  1 显示 0 不显示" json:"is_show"`
	Keywords  string `orm:"column(keywords);size(255)" json:"keywords"`
	Level     int    `orm:"column(level)" json:"level"`
	Name      string `orm:"column(category_name);size(50)" json:"name"`
	ParentId  int    `orm:"column(pid)" json:"parent_id"`
	SortOrder int    `orm:"column(sort);null" json:"sort_order"`
}
type SysWeixinMenu struct {
	InstanceId    int    `orm:"column(instance_id)" description:"店铺id"`
	Pid           int    `orm:"column(pid)" description:"父菜单"`
	MenuEventType int    `orm:"column(menu_event_type)" description:"1普通url 2 图文素材 3 功能"`
	MediaId       int    `orm:"column(media_id)" description:"图文消息ID"`
	MenuEventUrl  string `orm:"column(menu_event_url);size(255)" description:"菜单url"`
	Hits          int    `orm:"column(hits)" description:"触发数"`
	CreateDate    int    `orm:"column(create_date);null" description:"创建日期"`
	ModifyDate    int    `orm:"column(modify_date);null" description:"修改日期"`

	IconUrl   string `orm:"column(ico);size(32)" description:"菜图标单" json:"icon_url"`
	Id        int    `orm:"column(menu_id);auto" description:"主键" json:"id"`
	Name      string `orm:"column(menu_name);size(50)" description:"菜单名称" json:"name"`
	SortOrder int    `orm:"column(sort)" description:"排序" json:"sort_order"`
}

type NsMemberFavorites struct {
	ShopId     int     `orm:"column(shop_id)" description:"店铺ID" json:"shop_id,omitempty"`
	ShopName   string  `orm:"column(shop_name);size(20)" description:"店铺名称" json:"shop_name,omitempty"`
	ShopLogo   string  `orm:"column(shop_logo);size(255)" description:"店铺logo" json:"shop_logo,omitempty"`
	GoodsName  string  `orm:"column(goods_name);size(50)" description:"商品名称" json:"goods_name,omitempty"`
	GoodsImage string  `orm:"column(goods_image);size(300);null" json:"goods_image,omitempty"`
	LogPrice   float64 `orm:"column(log_price);digits(10);decimals(2)" description:"商品收藏时价格" json:"log_price,omitempty"`
	LogMsg     string  `orm:"column(log_msg);size(1000)" description:"收藏备注" json:"log_msg,omitempty"`

	FavTime int64  `orm:"column(fav_time);null" description:"收藏时间" json:"fav_time,omitempty"`
	Id      int    `orm:"column(log_id);auto" description:"记录ID" json:"id,omitempty"`
	FavType string `orm:"column(fav_type);size(20)" description:"类型:goods为商品,shop为店铺,默认为商品" json:"fav_type"`
	UserId  int    `orm:"column(uid)" description:"会员ID" json:"user_id"`
	ValueId int    `orm:"column(fav_id)" description:"商品或店铺ID" json:"value_id"`
}
type NsGoodsComment struct {
	OrderId int     `orm:"column(order_id)" description:"订单id" json:"order_id"`
	Number  float64 `orm:"column(number);digits(10);decimals(2)" description:"数量" json:"Number"`

	CreateTime int `orm:"column(create_time);null" description:"评论创建时间" json:"add_time" json:"create_time"`
	Id         int `orm:"column(id);auto" description:"主键id" json:"id" json:"id"`
	Status     int `orm:"column(status)" description:"评论状态 0未评论 1已评论" json:"status"`
	UserId     int `orm:"column(uid)" description:"用户id" json:"user_id"`
}
type NsGoodsEvaluate struct {
	Id           int     `orm:"column(id);auto" description:"评价ID"  json:"id"`
	OrderId      int     `orm:"column(order_id)" description:"订单ID"  json:"order_id"`
	OrderNo      string  `orm:"column(order_no)" description:"订单编号"  json:"order_no"`
	OrderGoodsId int     `orm:"column(order_goods_id)" description:"订单项ID"  json:"order_goods_id"`
	GoodsId      int     `orm:"column(goods_id)" description:"商品ID"  json:"goods_id"`
	GoodsName    string  `orm:"column(goods_name);size(100)" description:"商品名称"  json:"goods_name"`
	GoodsPrice   float64 `orm:"column(goods_price);digits(10);decimals(2)" description:"商品价格"  json:"goods_price"`
	GoodsImage   string  `orm:"column(goods_image);size(255)" description:"商品图片"  json:"goods_image"`
	ShopName     string  `orm:"column(shop_name);size(100)" description:"店铺名称"  json:"shop_name"`
	Content      string  `orm:"column(content);size(255)" description:"评价内容"  json:"content"`
	Image        string  `orm:"column(image);size(1000)" description:"评价图片"  json:"image"`
	ExplainFirst string  `orm:"column(explain_first);size(255)" description:"解释内容"  json:"explain_first"`
	MemberName   string  `orm:"column(member_name);size(100)" description:"评价人名称"  json:"member_name"`
	Uid          int     `orm:"column(uid)" description:"评价人编号"  json:"uid"`
	IsAnonymous  int     `orm:"column(is_anonymous)" description:"0表示不是 1表示是匿名评价"  json:"is_anonymous"`
	Scores       int     `orm:"column(scores)" description:"1-5分"  json:"scores"`
	AgainContent string  `orm:"column(again_content);size(255)" description:"追加评价内容"  json:"again_content"`
	AgainImage   string  `orm:"column(again_image);size(1000)" description:"追评评价图片"  json:"again_image"`
	AgainExplain string  `orm:"column(again_explain);size(255)" description:"追加解释内容"  json:"again_explain"`
	ExplainType  int     `orm:"column(explain_type);null" description:"1好评2中评3差评"  json:"explain_type"`
	IsShow       int     `orm:"column(is_show);null" description:"1显示 0隐藏"  json:"is_show"`
	Addtime      int64   `orm:"column(addtime);null" description:"评价时间"  json:"addtime"`
	AgainAddtime int64   `orm:"column(again_addtime);null" description:"追加评价时间" json:"again_addtime"`
}
type NsCoupon struct {
	Id            int     `orm:"column(coupon_id);auto" description:"优惠券id"`
	CouponTypeId  int     `orm:"column(coupon_type_id)" description:"优惠券类型id"`
	ShopId        int     `orm:"column(shop_id)" description:"店铺Id"`
	CouponCode    string  `orm:"column(coupon_code);size(255)" description:"优惠券编码"`
	Uid           int     `orm:"column(uid)" description:"领用人"`
	UseOrderId    int     `orm:"column(use_order_id)" description:"优惠券使用订单id"`
	CreateOrderId int     `orm:"column(create_order_id)" description:"创建订单id(优惠券只有是完成订单发放的优惠券时才有值)"`
	Money         float64 `orm:"column(money);digits(10);decimals(2)" description:"面额"`
	State         int     `orm:"column(state)" description:"优惠券状态 0未领用 1已领用（未使用） 2已使用 3已过期"`
	GetType       int     `orm:"column(get_type)" description:"获取方式1订单2.首页领取"`
	FetchTime     int     `orm:"column(fetch_time);null" description:"领取时间"`
	UseTime       int     `orm:"column(use_time);null" description:"使用时间"`
	StartTime     int     `orm:"column(start_time);null" description:"有效期开始时间"`
	EndTime       int     `orm:"column(end_time);null" description:"有效期结束时间"`
}
type NsCouponType struct {
	ShopId        int `orm:"column(shop_id)" description:"店铺ID"`
	Count         int `orm:"column(count)" description:"发放数量"`
	MaxFetch      int `orm:"column(max_fetch)" description:"每人最大领取个数 0无限制"`
	NeedUserLevel int `orm:"column(need_user_level)" description:"领取人会员等级"`
	RangeType     int `orm:"column(range_type)" description:"使用范围0部分产品使用 1全场产品使用"`
	IsShow        int `orm:"column(is_show)" description:"是否允许首页显示0不显示1显示"`
	CreateTime    int `orm:"column(create_time);null" description:"创建时间"`
	UpdateTime    int `orm:"column(update_time);null" description:"修改时间"`

	Id             int    `orm:"column(coupon_type_id);auto" description:"优惠券类型Id" json:"id"`
	MinGoodsAmount string `orm:"column(at_least);digits(10);decimals(2)" description:"满多少元使用 0代表无限制" json:"min_goods_amount"`
	Name           string `orm:"column(coupon_name);size(50)" description:"优惠券名称" json:"name"`
	TypeMoney      string `orm:"column(money);digits(10);decimals(2)" description:"发放面额" json:"type_money"`
	UseEndDate     int    `orm:"column(start_time);null" description:"有效日期开始时间" json:"use_end_date"`
	UseStartDate   int    `orm:"column(end_time);null" description:"有效日期结束时间" json:"use_start_date"`
}
type NsOrderRefund struct {
	Id             int    `orm:"column(id);auto" description:"id"`
	OrderGoodsId   int    `orm:"column(order_goods_id)" description:"订单商品表id"`
	RefundStatus   string `orm:"column(refund_status);size(255)" description:"操作状态"`
	Action         string `orm:"column(action);size(255)" description:"退款操作内容描述"`
	ActionWay      int    `orm:"column(action_way)" description:"操作方 1 买家 2 卖家"`
	ActionUserid   string `orm:"column(action_userid);size(255)" description:"操作人id"`
	ActionUsername string `orm:"column(action_username);size(255)" description:"操作人姓名"`
	ActionTime     int    `orm:"column(action_time);null" description:"操作时间"`
}
type NsPlatformLink struct {
	Id        int    `orm:"column(link_id);auto" description:"索引id"`
	LinkTitle string `orm:"column(link_title);size(100)" description:"标题"`
	LinkUrl   string `orm:"column(link_url);size(100)" description:"链接"`
	LinkPic   string `orm:"column(link_pic);size(100)" description:"图片"`
	IsBlank   int    `orm:"column(is_blank)" description:"是否新窗口打开 1.是 2.否"`
	IsShow    int    `orm:"column(is_show)" description:"是否显示 1.是 2.否"`
}
type NsGoods struct {
	PromotionPrice   string `orm:"column(promotion_price);digits(10);decimals(2)" description:"商品促销价格" json:"promotion_price"`
	Clicks           uint   `orm:"column(clicks)" description:"商品点击数量" json:"clicks"`
	MinStockAlarm    int    `orm:"column(min_stock_alarm)" description:"库存预警值" json:"min_stock_alarm"`
	Sales            uint   `orm:"column(sales)" description:"销售数量" json:"sales"`
	Collects         uint   `orm:"column(collects)" description:"收藏数量" json:"collects"`
	Star             uint   `orm:"column(star)" description:"好评星级" json:"star"`
	Evaluates        uint   `orm:"column(evaluates)" description:"评价数" json:"evaluates"`
	Shares           int    `orm:"column(shares)" description:"分享数" json:"shares"`
	ProvinceId       uint   `orm:"column(province_id)" description:"一级地区id" json:"province_id"`
	CityId           uint   `orm:"column(city_id)" description:"二级地区id" json:"city_id"`
	Picture          int    `orm:"column(picture)" description:"商品主图" json:"picture"`
	QRcode           string `orm:"column(QRcode);size(255)" description:"商品二维码" json:"QRcode"`
	IsStockVisible   int    `orm:"column(is_stock_visible)" description:"页面不显示库存" json:"is_stock_visible"`
	IsRecommend      int    `orm:"column(is_recommend)" description:"是否推荐" json:"is_recommend"`
	IsPreSale        int    `orm:"column(is_pre_sale);null" json:"is_pre_sale"`
	IsBill           int    `orm:"column(is_bill)" description:"是否开具增值税发票 1是，0否" json:"is_bill"`
	State            int    `orm:"column(state)" description:"商品状态 0下架，1正常，10违规（禁售）" json:"state"`
	Introduction     string `orm:"column(introduction);size(255)" description:"商品简介，促销语" json:"introduction"`
	ImgIdArray       string `orm:"column(img_id_array);size(1000);null" description:"商品图片序列" json:"img_id_array"`
	SkuImgArray      string `orm:"column(sku_img_array);size(1000);null" description:"商品sku应用图片列表  属性,属性值，图片ID" json:"sku_img_array"`
	GoodsAttributeId int    `orm:"column(goods_attribute_id)" description:"商品类型" json:"goods_attribute_id"`
	GoodsSpecFormat  string `orm:"column(goods_spec_format)" description:"商品规格" json:"goods_spec_format"`
	SupplierId       int    `orm:"column(supplier_id)" description:"供货商id" json:"supplier_id"`
	SaleDate         int    `orm:"column(sale_date);null" description:"上下架时间" json:"sale_date"`

	CreateTime  int    `orm:"column(create_time);null" description:"商品添加时间" json:"create_time"`
	MarketPrice string `orm:"column(market_price);digits(10);decimals(2)" description:"市场价" json:"market_price"`
	BrandId     int    `orm:"column(brand_id)" description:"品牌id" json:"brand_id"`
	CategoryId  int    `orm:"column(category_id)" description:"商品分类id" json:"category_id"`
	CostPrice   string `orm:"column(cost_price);digits(19);decimals(2)" description:"成本价" json:"cost_price"`
	ShippingFee string `orm:"column(shipping_fee);digits(10);decimals(2)" description:"运费 0为免运费" json:"shipping_fee"`
	GoodsDesc   string `orm:"column(description)" description:"商品详情" json:"goods_desc"`
	Stock       int    `orm:"column(stock)" description:"商品库存" json:"stock"`
	GoodsId     int    `orm:"column(goods_id);auto" description:"商品id(SKU)" json:"goods_id"`
	IsHot       int    `orm:"column(is_hot)" description:"是否热销商品" json:"is_hot"`
	IsNew       int    `orm:"column(is_new)" description:"是否新品" json:"is_new"`
	Keywords    string `orm:"column(keywords);size(255)" description:"商品关键词" json:"keywords"`
	GoodsName   string `orm:"column(goods_name);size(100)" description:"商品名称" json:"goods_name"`
	Price       string `orm:"column(price);digits(19);decimals(2)" description:"商品原价格" json:"price"`
	Sort        int    `orm:"column(sort)" description:"排序" json:"sort"`
}
type NsGoodsAttribute struct {
	Id            int    `orm:"column(attr_id);auto" json:"id"`
	GoodsId       int    `orm:"column(goods_id)" description:"商品ID" json:"goods_id"`
	AttributeId   int    `orm:"column(attr_value_id)" description:"属性值id" json:"attr_value_id"`
	Value         string `orm:"column(attr_value);size(255)" description:"属性值名称" json:"attr_value"`
	AttrValueName string `orm:"column(attr_value_name);size(255)" description:"属性值对应数据值" json:"attr_value_name"`
	Sort          int    `orm:"column(sort)" description:"排序" json:"sort"`
	CreateTime    int    `orm:"column(create_time);null" description:"创建时间" json:"create_time"`
}
type SysAlbumClass struct {
	Id         int    `orm:"column(album_id);auto" description:"相册id" json:"album_id"`
	Pid        int    `orm:"column(pid)" description:"上级相册ID" json:"pid"`
	AlbumName  string `orm:"column(album_name);size(100)" description:"相册名称" json:"album_name"`
	AlbumCover string `orm:"column(album_cover);size(255)" description:"相册封面" json:"album_cover"`
	IsDefault  uint   `orm:"column(is_default)" description:"是否为默认相册,1代表默认" json:"is_default"`
	Sort       int    `orm:"column(sort);null" json:"sort"`
	CreateTime int    `orm:"column(create_time);null" description:"创建时间" json:"create_time"`
}

type SysAlbumPicture struct {
	Id            int    `orm:"column(pic_id);auto" description:"相册图片表id" json:"create_time"`
	ShopId        uint   `orm:"column(shop_id);null" description:"所属实例id" json:"create_time"`
	AlbumId       uint   `orm:"column(album_id)" description:"相册id" json:"create_time"`
	IsWide        int    `orm:"column(is_wide)" description:"是否宽屏" json:"create_time"`
	PicName       string `orm:"column(pic_name);size(100)" description:"图片名称" json:"create_time"`
	PicTag        string `orm:"column(pic_tag);size(255)" description:"图片标签" json:"pic_tag"`
	PicCover      string `orm:"column(pic_cover);size(255)" description:"原图图片路径" json:"pic_cover"`
	PicSize       string `orm:"column(pic_size);size(255)" description:"原图大小" json:"pic_size"`
	PicSpec       string `orm:"column(pic_spec);size(100)" description:"原图规格" json:"pic_spec"`
	PicCoverBig   string `orm:"column(pic_cover_big);size(255)" description:"大图路径" json:"pic_cover_big"`
	PicSizeBig    string `orm:"column(pic_size_big);size(255)" description:"大图大小" json:"pic_size_big"`
	PicSpecBig    string `orm:"column(pic_spec_big);size(100)" description:"大图规格" json:"pic_spec_big"`
	PicCoverMid   string `orm:"column(pic_cover_mid);size(255)" description:"中图路径" json:"pic_cover_mid"`
	PicSizeMid    string `orm:"column(pic_size_mid);size(255)" description:"中图大小" json:"pic_size_mid"`
	PicSpecMid    string `orm:"column(pic_spec_mid);size(100)" description:"中图规格" json:"pic_spec_mid"`
	PicCoverSmall string `orm:"column(pic_cover_small);size(255)" description:"小图路径" json:"pic_cover_small"`
	PicSizeSmall  string `orm:"column(pic_size_small);size(255)" description:"小图大小" json:"pic_size_small"`
	PicSpecSmall  string `orm:"column(pic_spec_small);size(255)" description:"小图规格" json:"pic_spec_small"`
	PicCoverMicro string `orm:"column(pic_cover_micro);size(255)" description:"微图路径" json:"pic_cover_micro"`
	PicSizeMicro  string `orm:"column(pic_size_micro);size(255)" description:"微图大小" json:"pic_size_micro"`
	PicSpecMicro  string `orm:"column(pic_spec_micro);size(255)" description:"微图规格" json:"pic_spec_micro"`
	UploadTime    int    `orm:"column(upload_time);null" description:"图片上传时间" json:"upload_time"`
	UploadType    int    `orm:"column(upload_type);null" description:"图片外链" json:"upload_type"`
	Domain        string `orm:"column(domain);size(255);null" description:"图片外链" json:"domain"`
	Bucket        string `orm:"column(bucket);size(255);null" description:"存储空间名称" json:"bucket"`
}
type NsGoodsSpec struct {
	Id         int    `orm:"column(spec_id);auto" description:"属性ID"`
	ShopId     int    `orm:"column(shop_id)" description:"店铺ID"`
	SpecName   string `orm:"column(spec_name);size(255)" description:"属性名称"`
	IsVisible  int    `orm:"column(is_visible)" description:"是否可视"`
	Sort       int    `orm:"column(sort)" description:"排序"`
	ShowType   int    `orm:"column(show_type)" description:"展示方式 1 文字 2 颜色 3 图片"`
	CreateTime int    `orm:"column(create_time);null" description:"创建日期"`
	IsScreen   int    `orm:"column(is_screen)" description:"是否参与筛选 0 不参与 1 参与"`
}
type NsGoodsSpecValue struct {
	Id            int    `orm:"column(spec_value_id);auto" description:"商品属性值ID"`
	SpecId        int    `orm:"column(spec_id)" description:"商品属性ID"`
	SpecValueName string `orm:"column(spec_value_name);size(255)" description:"商品属性值名称"`
	SpecValueData string `orm:"column(spec_value_data);size(255)" description:"商品属性值数据"`
	IsVisible     int    `orm:"column(is_visible)" description:"是否可视"`
	Sort          int    `orm:"column(sort)" description:"排序"`
	CreateTime    int    `orm:"column(create_time);null"`
}
type NsOrder struct {
	OutTradeNo         string  `orm:"column(out_trade_no);size(100)" description:"外部交易号" json:"out_trade_no,omitempty"`
	OrderType          int     `orm:"column(order_type)" description:"订单类型" json:"order_type,omitempty"`
	PaymentType        int     `orm:"column(payment_type)" description:"支付类型。取值范围" json:"payment_type,omitempty"`
	ShippingType       int     `orm:"column(shipping_type)" description:"订单配送方式" json:"shipping_type,omitempty"`
	OrderFrom          string  `orm:"column(order_from);size(255)" description:"订单来源" json:"order_from,omitempty"`
	UserName           string  `orm:"column(user_name);size(50)" description:"买家会员名称" json:"user_name,omitempty"`
	BuyerIp            string  `orm:"column(buyer_ip);size(20)" description:"买家ip" json:"buyer_ip,omitempty"`
	BuyerMessage       string  `orm:"column(buyer_message);size(255)" description:"买家附言" json:"buyer_message,omitempty"`
	BuyerInvoice       string  `orm:"column(buyer_invoice);size(255)" description:"买家发票信息" json:"buyer_invoice,omitempty"`
	ReceiverZip        string  `orm:"column(receiver_zip);size(6)" description:"收货人邮编" json:"receiver_zip,omitempty"`
	ShopId             int     `orm:"column(shop_id)" description:"卖家店铺id" json:"shop_id,omitempty"`
	ShopName           string  `orm:"column(shop_name);size(100)" description:"卖家店铺名称" json:"shop_name,omitempty"`
	SellerStar         int     `orm:"column(seller_star)" description:"卖家对订单的标注星标" json:"seller_star,omitempty"`
	SellerMemo         string  `orm:"column(seller_memo);size(255)" description:"卖家对订单的备注" json:"seller_memo,omitempty"`
	ConsignTimeAdjust  int     `orm:"column(consign_time_adjust)" description:"卖家延迟发货时间" json:"consign_time_adjust,omitempty"`
	GoodsMoney         float64 `orm:"column(goods_money);digits(19);decimals(2)" description:"商品总价" json:"goods_money,omitempty"`
	Point              int     `orm:"column(point)" description:"订单消耗积分" json:"point,omitempty"`
	PointMoney         float64 `orm:"column(point_money);digits(10);decimals(2)" description:"订单消耗积分抵多少钱" json:"point_money,omitempty"`
	UserMoney          float64 `orm:"column(user_money);digits(10);decimals(2)" description:"订单余额支付金额" json:"user_money,omitempty"`
	UserPlatformMoney  float64 `orm:"column(user_platform_money);digits(10);decimals(2)" description:"用户平台余额支付" json:"user_platform_money,omitempty"`
	PromotionMoney     float64 `orm:"column(promotion_money);digits(10);decimals(2)" description:"订单优惠活动金额" json:"promotion_money,omitempty"`
	PayMoney           float64 `orm:"column(pay_money);digits(10);decimals(2)" description:"订单实付金额" json:"pay_money,omitempty"`
	RefundMoney        float64 `orm:"column(refund_money);digits(10);decimals(2)" description:"订单退款金额" json:"refund_money,omitempty"`
	CoinMoney          float64 `orm:"column(coin_money);digits(10);decimals(2)" description:"购物币金额" json:"coin_money,omitempty"`
	GivePoint          int     `orm:"column(give_point)" description:"订单赠送积分" json:"give_point,omitempty"`
	GiveCoin           float64 `orm:"column(give_coin);digits(10);decimals(2)" description:"订单成功之后返购物币" json:"give_coin,omitempty"`
	ReviewStatus       int     `orm:"column(review_status)" description:"订单评价状态" json:"review_status,omitempty"`
	FeedbackStatus     int     `orm:"column(feedback_status)" description:"订单维权状态" json:"feedback_status,omitempty"`
	IsEvaluate         int16   `orm:"column(is_evaluate)" description:"是否评价 0为未评价 1为已评价 2为已追评" json:"is_evaluate,omitempty"`
	TaxMoney           float64 `orm:"column(tax_money);digits(10);decimals(2)" json:"tax_money,omitempty"`
	ShippingCompanyId  int     `orm:"column(shipping_company_id)" description:"配送物流公司ID" json:"shipping_company_id,omitempty"`
	GivePointType      int     `orm:"column(give_point_type)" description:"积分返还类型 1 订单完成  2 订单收货 3  支付订单" json:"give_point_type,omitempty"`
	ShippingTime       int     `orm:"column(shipping_time);null" description:"买家要求配送时间" json:"shipping_time,omitempty"`
	SignTime           int     `orm:"column(sign_time);null" description:"买家签收时间" json:"sign_time,omitempty"`
	ConsignTime        int     `orm:"column(consign_time);null" description:"卖家发货时间" json:"consign_time,omitempty"`
	IsDeleted          int     `orm:"column(is_deleted)" description:"订单是否已删除" json:"is_deleted,omitempty"`
	OperatorType       int     `orm:"column(operator_type)" description:"操作人类型  1店铺  2用户" json:"operator_type,omitempty"`
	OperatorId         int     `orm:"column(operator_id)" description:"操作人id" json:"operator_id,omitempty"`
	RefundBalanceMoney float64 `orm:"column(refund_balance_money);digits(10);decimals(2)" description:"订单退款余额" json:"refund_balance_money,omitempty"`
	FixedTelephone     string  `orm:"column(fixed_telephone);size(50)" description:"固定电话" json:"fixed_telephone,omitempty" `

	CreateTime     int64   `orm:"column(create_time);null" description:"订单创建时间" json:"create_time,omitempty"`
	Address        string  `orm:"column(receiver_address);size(255)" description:"收货人详细地址" json:"address"`
	City           int     `orm:"column(receiver_city)" description:"收货人所在城市" json:"city"`
	ConfirmTime    int64   `orm:"column(finish_time);null" description:"订单完成时间" json:"confirm_time"`
	Consignee      string  `orm:"column(receiver_name);size(50)" description:"收货人姓名" json:"consignee"`
	CouponId       int     `orm:"column(coupon_id)" description:"订单代金券id" json:"coupon_id,omitempty"`
	CouponPrice    float64 `orm:"column(coupon_money);digits(10);decimals(2)" description:"订单代金券支付金额" json:"coupon_price,omitempty"`
	District       int     `orm:"column(receiver_district)" description:"收货人所在街道" json:"district"`
	Id             int     `orm:"column(order_id);auto" description:"订单id" json:"id"`
	Mobile         string  `orm:"column(receiver_mobile);size(11)" description:"收货人的手机号码" json:"mobile"`
	OrderMoney     float64 `orm:"column(order_money);digits(10);decimals(2)" description:"订单总价" json:"order_money"`
	OrderSn        string  `orm:"column(order_no);size(255);null" description:"订单编号" json:"order_sn"`
	OrderStatus    int     `orm:"column(order_status)" description:"订单状态" json:"order_status"`
	PayStatus      int     `orm:"column(pay_status)" description:"订单付款状态"`
	PayTime        int64   `orm:"column(pay_time);null" description:"订单付款时间" json:"pay_time"`
	Province       int     `orm:"column(receiver_province)" description:"收货人所在省" json:"province"`
	ShippingFee    float64 `orm:"column(shipping_money);digits(10);decimals(2)" description:"订单运费" json:"shipping_fee"`
	ShippingStatus int     `orm:"column(shipping_status)" description:"订单配送状态" json:"shipping_status"`
	UserId         int     `orm:"column(buyer_id)" description:"买家id" json:"user_id"`
}
type NsOrderGoodsExpress struct {
	OrderGoodsIdArray string `orm:"column(order_goods_id_array);size(255)" description:"订单项商品组合列表" json:"order_goods_id_array,omitempty"`
	ExpressName       string `orm:"column(express_name);size(50)" description:"包裹名称  （包裹- 1 包裹 - 2）" json:"express_name,omitempty"`
	ShippingType      int    `orm:"column(shipping_type)" description:"发货方式1 需要物流 0无需物流" json:"shipping_type,omitempty"`
	Uid               int    `orm:"column(uid)" description:"用户id" json:"uid,omitempty"`
	UserName          string `orm:"column(user_name);size(50)" description:"用户名" json:"user_name,omitempty"`
	Memo              string `orm:"column(memo);size(255)" description:"备注" json:"memo,omitempty"`
	ShippingTime      int    `orm:"column(shipping_time);null" description:"发货时间" json:"shipping_time,omitempty"`

	Id          int    `orm:"column(id);auto" json:"id" json:"id,omitempty"`
	OrderId     int    `orm:"column(order_id)" description:"订单id"  json:"order_id,omitempty"`
	ShipperCode string `orm:"column(express_no);size(50)" description:"运单编号" json:"shipper_code" `
	ShipperId   int    `orm:"column(express_company_id)" description:"快递公司id" json:"express_company_id,omitempty"`
	ShipperName string `orm:"column(express_company);size(255)" description:"物流公司名称" json:"shipper_name"`
}
type NsOrderGoods struct {
	SkuId                 int     `orm:"column(sku_id)" description:"skuID" json:"sku_id,omitempty"`
	SkuName               string  `orm:"column(sku_name);size(50)" description:"sku名称" json:"sku_name,omitempty"`
	Price                 float64 `orm:"column(price);digits(19);decimals(2)" description:"商品价格" json:"price,omitempty"`
	CostPrice             float64 `orm:"column(cost_price);digits(19);decimals(2)" description:"商品成本价" json:"cost_price,omitempty"`
	AdjustMoney           float64 `orm:"column(adjust_money);digits(10);decimals(2)" description:"调整金额" json:"adjust_money,omitempty"`
	GoodsMoney            float64 `orm:"column(goods_money);digits(10);decimals(2)" description:"商品总价" json:"goods_money,omitempty"`
	GoodsPicture          int     `orm:"column(goods_picture)" description:"商品图片" json:"goods_picture,omitempty"`
	ShopId                int     `orm:"column(shop_id)" description:"店铺ID" json:"shop_id,omitempty"`
	BuyerId               int     `orm:"column(buyer_id)" description:"购买人ID" json:"buyer_id,omitempty"`
	PointExchangeType     int     `orm:"column(point_exchange_type)" description:"积分兑换类型0.非积分兑换1.积分兑换" json:"point_exchange_type,omitempty"`
	GoodsType             int     `orm:"column(goods_type);size(255)" description:"商品类型" json:"goods_type,omitempty"`
	PromotionId           int     `orm:"column(promotion_id)" description:"促销ID" json:"promotion_id,omitempty"`
	PromotionTypeId       int     `orm:"column(promotion_type_id)" description:"促销类型" json:"promotion_type_id,omitempty"`
	OrderType             int     `orm:"column(order_type)" description:"订单类型" json:"order_type,omitempty"`
	OrderStatus           int     `orm:"column(order_status)" description:"订单状态" json:"order_status,omitempty"`
	GivePoint             int     `orm:"column(give_point)" description:"积分数量" json:"give_point,omitempty"`
	ShippingStatus        int     `orm:"column(shipping_status)" description:"物流状态" json:"shipping_status,omitempty"`
	RefundType            int     `orm:"column(refund_type)" description:"退款方式" json:"refund_type,omitempty"`
	RefundRequireMoney    float64 `orm:"column(refund_require_money);digits(10);decimals(2)" description:"退款金额" json:"refund_require_money,omitempty"`
	RefundReason          string  `orm:"column(refund_reason);size(255)" description:"退款原因" json:"refund_reason,omitempty"`
	RefundShippingCode    string  `orm:"column(refund_shipping_code);size(255)" description:"退款物流单号" json:"refund_shipping_code,omitempty"`
	RefundShippingCompany string  `orm:"column(refund_shipping_company);size(255)" description:"退款物流公司名称" json:"refund_shipping_company,omitempty"`
	RefundRealMoney       float64 `orm:"column(refund_real_money);digits(10);decimals(2)" description:"实际退款金额" json:"refund_real_money,omitempty"`
	RefundStatus          int     `orm:"column(refund_status)" description:"退款状态" json:"refund_status,omitempty"`
	Memo                  string  `orm:"column(memo);size(255)" description:"备注" json:"memo,omitempty"`
	IsEvaluate            int16   `orm:"column(is_evaluate)" description:"是否评价 0为未评价 1为已评价 2为已追评" json:"is_evaluate,omitempty"`
	RefundTime            int     `orm:"column(refund_time);null" description:"退款时间" json:"refund_time,omitempty"`
	RefundBalanceMoney    float64 `orm:"column(refund_balance_money);digits(10);decimals(2)" description:"订单退款余额" json:"refund_balance_money,omitempty"`
	TmpExpressCompany     string  `orm:"column(tmp_express_company);size(255)" description:"批量打印时添加的临时物流公司" json:"tmp_express_company,omitempty"`
	TmpExpressCompanyId   int     `orm:"column(tmp_express_company_id)" description:"批量打印时添加的临时物流公司id" json:"tmp_express_company_id,omitempty"`
	TmpExpressNo          string  `orm:"column(tmp_express_no);size(50)" description:"批量打印时添加的临时订单编号" json:"tmp_express_no,omitempty"`

	GoodsId   int    `orm:"column(goods_id)" description:"商品ID" json:"goods_id"`
	GoodsName string `orm:"column(goods_name);size(50)" description:"商品名称" json:"goods_name"`
	Id        int    `orm:"column(order_goods_id);auto" description:"订单项ID" json:"id"`
	Number    int    `orm:"column(num);size(255)" description:"购买数量" json:"number"`
	OrderId   int    `orm:"column(order_id)" description:"订单ID" json:"order_id"`
}

type NsExpressCompany struct {
	ShopId      int    `orm:"column(shop_id)" description:"商铺id"`
	ExpressNo   string `orm:"column(express_no);size(20)" description:"物流编号"`
	IsEnabled   int    `orm:"column(is_enabled)" description:"使用状态"`
	Image       string `orm:"column(image);size(255);null" description:"物流公司模版图片"`
	Phone       string `orm:"column(phone);size(50)" description:"联系电话"`
	Orders      int    `orm:"column(orders);null"`
	ExpressLogo string `orm:"column(express_logo);size(255);null" description:"公司logo"`
	IsDefault   int    `orm:"column(is_default)" description:"是否设置为默认 0未设置 1 默认"`

	Id   int    `orm:"column(co_id);auto" description:"表序号" json:"id"`
	Name string `orm:"column(company_name);size(50)" description:"物流公司名称" json:"name"`
}
type NsShopAd struct {
	Id         int    `orm:"column(id);auto"`
	ShopId     int    `orm:"column(shop_id)"`
	ItemPicUrl string `orm:"column(ad_image);size(255)" description:"广告图片" json:"item_pic_url"`
	LinkUrl    string `orm:"column(link_url);size(255)" description:"链接地址"`
	Sort       int    `orm:"column(sort)" description:"排序"`
	Type       int    `orm:"column(type)" description:"类型 0 -- pc端  1-- 手机端 "`
	Background string `orm:"column(background);size(255)" description:"背景色"`
}

type NsGoodsSku struct {
	Id                   int     `orm:"column(sku_id);auto" description:"表序号" json:"sku_id"`
	GoodsId              int     `orm:"column(goods_id)" description:"商品编号" json:"goods_id"`
	SkuName              string  `orm:"column(sku_name);size(500)" description:"SKU名称" json:"sku_name"`
	AttrValueItems       string  `orm:"column(attr_value_items);size(255)" description:"属性和属性值 id串 attribute + attribute value 表ID分号分隔" json:"attr_value_items"`
	AttrValueItemsFormat string  `orm:"column(attr_value_items_format);size(500)" description:"属性和属性值id串组合json格式" json:"attr_value_items_format"`
	MarketPrice          float64 `orm:"column(market_price);digits(10);decimals(2)" description:"市场价" json:"market_price"`
	Price                float64 `orm:"column(price);digits(10);decimals(2)" description:"价格" json:"price"`
	PromotePrice         float64 `orm:"column(promote_price);digits(10);decimals(2)" description:"促销价格" json:"promote_price"`
	CostPrice            float64 `orm:"column(cost_price);digits(19);decimals(2)" description:"成本价" json:"cost_price"`
	Stock                int     `orm:"column(stock)" description:"库存" json:"stock"`
	Picture              int     `orm:"column(picture)" description:"如果是第一个sku编码, 可以加图片" json:"picture"`
	Code                 string  `orm:"column(code);size(255)" description:"商家编码" json:"code"`
	QRcode               string  `orm:"column(QRcode);size(255)" description:"商品二维码" json:"QRcode"`
	CreateDate           int     `orm:"column(create_date);null" description:"创建时间" json:"create_date"`
	UpdateDate           int     `orm:"column(update_date);null" description:"修改时间" json:"update_date"`
}
type NsNewGoods struct {
	MarketPrice string `orm:"column(market_price);digits(10);decimals(2)" description:"市场价" json:"market_price"`
	BrandId     int    `orm:"column(brand_id)" description:"品牌id" json:"brand_id"`
	CategoryId  int    `orm:"column(category_id)" description:"商品分类id" json:"category_id"`
	stock       int    `orm:"column(stock)" description:"商品库存" json:"stock"`
	Id          int    `orm:"column(goods_id);auto" description:"商品id(SKU)" json:"id"`
	Name        string `orm:"column(goods_name);size(100)" description:"商品名称" json:"name"`
	RetailPrice string `orm:"column(price);digits(19);decimals(2)" description:"商品原价格" json:"price"`
	PicCoverBig string `json:"pic_cover_big"`
	PicCoverMid string `json:"pic_cover_mid"`
	PicCover    string `json:"pic_cover"`
}
type SysUser struct {
	Id               int    `orm:"column(uid);auto" description:"主键" json:"uid,omitempty"`
	InstanceId       int    `orm:"column(instance_id)" description:"实例信息" json:"instance_id,omitempty"`
	UserName         string `orm:"column(user_name);size(50)" description:"用户名" json:"user_name,omitempty"`
	UserPassword     string `orm:"column(user_password);size(255)" description:"用户密码（MD5）" json:"user_password,omitempty"`
	UserStatus       int    `orm:"column(user_status)" description:"用户状态  用户状态默认为1" json:"user_status,omitempty"`
	UserHeadimg      string `orm:"column(user_headimg);size(255)" description:"用户头像" json:"user_headimg,omitempty"`
	IsSystem         int    `orm:"column(is_system)" description:"是否是系统后台用户 0 不是 1 是" json:"is_system,omitempty"`
	IsMember         int    `orm:"column(is_member)" description:"是否是前台会员" json:"is_member,omitempty"`
	UserTel          string `orm:"column(user_tel);size(20)" description:"手机号" json:"user_tel,omitempty"`
	UserTelBind      int    `orm:"column(user_tel_bind)" description:"手机号是否绑定 0 未绑定 1 绑定 " json:"user_tel_bind,omitempty"`
	UserQq           string `orm:"column(user_qq);size(255);null" description:"qq号" json:"user_qq,omitempty"`
	QqOpenid         string `orm:"column(qq_openid);size(255)" description:"qq互联id" json:"qq_openid,omitempty"`
	QqInfo           string `orm:"column(qq_info);size(2000)" description:"qq账号相关信息" json:"qq_info,omitempty"`
	UserEmail        string `orm:"column(user_email);size(50)" description:"邮箱" json:"user_email,omitempty"`
	UserEmailBind    int    `orm:"column(user_email_bind)" description:"是否邮箱绑定" json:"user_email_bind,omitempty"`
	WxOpenid         string `orm:"column(wx_openid);size(255);null" description:"微信用户openid" json:"wx_openid,omitempty"`
	WxIsSub          int    `orm:"column(wx_is_sub)" description:"微信用户是否关注" json:"wx_is_sub,omitempty"`
	WxInfo           string `orm:"column(wx_info);size(2000);null" description:"微信用户信息" json:"wx_info,omitempty"`
	OtherInfo        string `orm:"column(other_info);size(255);null" description:"附加信息" json:"other_info,omitempty"`
	CurrentLoginIp   string `orm:"column(current_login_ip);size(255);null" description:"当前登录ip" json:"current_login_ip,omitempty"`
	CurrentLoginType int    `orm:"column(current_login_type);null" description:"当前登录的操作终端类型" json:"current_login_type,omitempty"`
	LastLoginIp      string `orm:"column(last_login_ip);size(255);null" description:"上次登录ip" json:"last_login_ip,omitempty"`
	LastLoginType    int    `orm:"column(last_login_type);null" description:"上次登录的操作终端类型" json:"last_login_type,omitempty"`
	LoginNum         int    `orm:"column(login_num)" description:"登录次数" json:"login_num,omitempty"`
	RealName         string `orm:"column(real_name);size(50);null" description:"真实姓名" json:"real_name,omitempty"`
	Sex              int16  `orm:"column(sex);null" description:"性别 0保密 1男 2女" json:"sex,omitempty"`
	Location         string `orm:"column(location);size(255);null" description:"所在地" json:"location,omitempty"`
	NickName         string `orm:"column(nick_name);size(50);null" description:"用户昵称" json:"nick_name,omitempty"`
	WxUnionid        string `orm:"column(wx_unionid);size(255)" description:"微信unionid" json:"wx_unionid,omitempty"`
	QrcodeTemplateId int    `orm:"column(qrcode_template_id)" description:"模板id" json:"qrcode_template_id,omitempty"`
	WxSubTime        int    `orm:"column(wx_sub_time);null" description:"微信用户关注时间" json:"wx_sub_time,omitempty"`
	WxNotsubTime     int    `orm:"column(wx_notsub_time);null" description:"微信用户取消关注时间" json:"wx_notsub_time,omitempty"`
	RegTime          int    `orm:"column(reg_time);null" description:"注册时间" json:"reg_time,omitempty"`
	CurrentLoginTime int    `orm:"column(current_login_time);null" description:"当前登录时间" json:"current_login_time,omitempty"`
	LastLoginTime    int    `orm:"column(last_login_time);null" description:"上次登录时间" json:"last_login_time,omitempty"`
	Birthday         int    `orm:"column(birthday);null" json:"birthday,omitempty"`
}

type SysArea struct {
	Id       int    `orm:"column(area_id);auto" json:"area_id,omitempty"`
	AreaName string `orm:"column(area_name);size(50)" json:"area_name,omitempty"`
	Sort     int    `orm:"column(sort);null" json:"sort,omitempty"`
}
type SysProvince struct {
	Id           int    `orm:"column(province_id);auto" json:"province_id,omitempty"`
	AreaId       int    `orm:"column(area_id)" json:"area_id,omitempty"`
	ProvinceName string `orm:"column(province_name);size(255)" json:"province_name,omitempty"`
	Sort         int    `orm:"column(sort)" json:"sort,omitempty"`
}
type SysCity struct {
	ProvinceId int    `orm:"column(province_id)" json:"province_id,omitempty"`
	Zipcode    string `orm:"column(zipcode);size(6)" json:"zipcode,omitempty"`
	Sort       int    `orm:"column(sort)" json:"sort,omitempty"`
	CityId     int    `orm:"column(city_id);auto"  json:"city_id,omitempty"`
	CityName   string `orm:"column(city_name);size(255)" json:"name" json:"city_name,omitempty"`
}
type SysDistrict struct {
	DistrictId   int    `orm:"column(district_id);auto" json:"district_id,omitempty"`
	CityId       int    `orm:"column(city_id);null" json:"city_id,omitempty"`
	DistrictName string `orm:"column(district_name);size(255)" json:"district_name,omitempty"`
	Sort         int    `orm:"column(sort)" json:"sort,omitempty"`
}
type NsMember struct {
	Uid         int    `orm:"column(uid);pk" description:"用户ID" json:"uid"`
	MemberName  string `orm:"column(member_name);size(50)" description:"前台用户名" json:"member_name"`
	MemberLevel int    `orm:"column(member_level)" description:"会员等级" json:"member_level"`
	Memo        string `orm:"column(memo);size(1000);null" description:"备注" json:"memo"`
	RegTime     int    `orm:"column(reg_time);null" description:"注册时间" json:"reg_time"`
}
