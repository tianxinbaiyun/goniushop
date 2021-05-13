package controllers

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/services"
	"github.com/tianxinbaiyun/goniushop/utils"
)

// OrderController OrderController
type OrderController struct {
	beego.Controller
}

// GetOrderPageData It may need to be refactored.
func GetOrderPageData(rawData []models.NsOrder, page int, size int) utils.PageData {

	count := len(rawData)
	totalPages := (count + size - 1) / size
	var pageData []models.NsOrder

	for idx := (page - 1) * size; idx < page*size && idx < count; idx++ {
		pageData = append(pageData, rawData[idx])
	}

	return utils.PageData{NumsPerPage: size, CurrentPage: page, Count: count, TotalPages: totalPages, Data: pageData}
}

// OrderListRtnJSON OrderListRtnJSON
type OrderListRtnJSON struct {
	models.NsOrder
	GoodsList       []models.NsOrderGoods    `json:"goodList"`
	GoodsCount      int                      `json:"goodsCount"`
	OrderStatusText string                   `json:"order_status_text"`
	HandOption      models.OrderHandleOption `json:"handleOption"`
}

// OrderList OrderList
func (c *OrderController) OrderList() {
	//请求参数
	page := utils.String2Int(c.GetString("page"))
	pageSize := utils.String2Int(c.GetString("pageSize"))
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	o := orm.NewOrm()
	orderTable := new(models.NsOrder)
	var orders []models.NsOrder
	_, _ = o.QueryTable(orderTable).Filter("buyer_id", getLoginUserID()).All(&orders)

	firstPageOrders := GetOrderPageData(orders, page, pageSize)

	var rtnOrderList []OrderListRtnJSON
	orderGoodsTable := new(models.NsOrderGoods)
	var orderGoods []models.NsOrderGoods
	qsOrderGoods := o.QueryTable(orderGoodsTable)
	for _, val := range firstPageOrders.Data.([]models.NsOrder) {
		_, _ = qsOrderGoods.Filter("order_id", val.ID).All(&orderGoods)
		var goodsCount int
		for _, val := range orderGoods {
			goodsCount += val.Number
		}
		orderStatusText := models.GetOrderStatusText(val.ID)
		orderHandOption := models.GetOrderHandleOption(val.ID)
		var orderListRtn = OrderListRtnJSON{val, orderGoods, goodsCount, orderStatusText, orderHandOption}

		rtnOrderList = append(rtnOrderList, orderListRtn)

	}

	firstPageOrders.Data = rtnOrderList

	utils.ReturnHTTPSuccess(&c.Controller, firstPageOrders)
	c.ServeJSON()
}

// OrderInfo OrderInfo
type OrderInfo struct {
	models.NsOrder
	ProvinceName        string                  `json:"province_name"`
	CityName            string                  `json:"city_name"`
	DistrictName        string                  `json:"district_name"`
	FullRegion          string                  `json:"full_region"`
	Express             services.ExpressRtnInfo `json:"express"`
	OrderStatusText     string                  `json:"order_status_text"`
	FormatAddTime       string                  `json:"add_time"`
	FormatFinalPlayTime string                  `json:"final_pay_time"`
}

// OrderDetailRtnJSON OrderDetailRtnJSON
type OrderDetailRtnJSON struct {
	OrderInfo    OrderInfo                `json:"orderInfo"`
	OrderGoods   []models.NsOrderGoods    `json:"orderGoods"`
	HandleOption models.OrderHandleOption `json:"handleOption"`
}

// OrderDetail OrderDetail
func (c *OrderController) OrderDetail() {
	//请求参数
	orderID := c.GetString("orderId")
	intOrderID := utils.String2Int(orderID)

	//错误返回
	var err error
	defer func() {
		if err != nil {
			logs.Debug("this order id[%v] has a error ,err:%v", orderID, err)
		}
		c.ServeJSON()
		return
	}()
	o := orm.NewOrm()
	orderTable := new(models.NsOrder)
	var order models.NsOrder
	err = o.QueryTable(orderTable).Filter("id", intOrderID).Filter("buyer_id", getLoginUserID()).One(&order)

	if err == orm.ErrNoRows {
		logs.Debug("order lose ,err:%v", err)
		utils.ReturnHTTPError(&c.Controller, 400, "订单不存在")
		return
	}

	var orderInfo = OrderInfo{NsOrder: order}
	orderInfo.ProvinceName = models.GetProvinceName(order.Province)
	orderInfo.CityName = models.GetCityName(order.City)
	orderInfo.DistrictName = models.GetDistrictName(order.District)
	orderInfo.FullRegion = orderInfo.ProvinceName + orderInfo.CityName + orderInfo.DistrictName

	lastExpressInfo := models.GetLatestOrderExpress(intOrderID)
	orderInfo.Express = lastExpressInfo

	orderGoodsTable := new(models.NsOrderGoods)
	var orderGoods []models.NsOrderGoods

	_, _ = o.QueryTable(orderGoodsTable).Filter("order_id", intOrderID).All(&orderGoods)

	orderInfo.OrderStatusText = models.GetOrderStatusText(intOrderID)
	orderInfo.FormatAddTime = utils.FormatTimestamp(orderInfo.CreateTime, "2006-01-02 03:04:05 PM")
	orderInfo.FormatFinalPlayTime = utils.FormatTimestamp(1234, "04:05")

	if orderInfo.OrderStatus == 0 {
		//todo 订单超时逻辑
		logs.Debug("order timeout ,err:%v", err)
		utils.ReturnHTTPError(&c.Controller, 400, "订单超时")
		return
	}

	handleOption := models.GetOrderHandleOption(intOrderID)
	utils.ReturnHTTPSuccess(&c.Controller, OrderDetailRtnJSON{
		OrderInfo:    orderInfo,
		OrderGoods:   orderGoods,
		HandleOption: handleOption,
	})
	return
}

// SubmitOrder SubmitOrder
type SubmitOrder struct {
	AddressID int    `json:"addressId"`
	IDs       string `json:"ids"`
}

// OrderSubmit OrderSubmit
func (c *OrderController) OrderSubmit() {
	var err error
	defer func() {
		if err != nil {
			logs.Debug("this has a err ,err:%v", err)
		}
		c.ServeJSON()
		return
	}()
	var input SubmitOrder
	body := c.Ctx.Input.RequestBody
	err = json.Unmarshal(body, &input)
	if err != nil {
		logs.Debug("json Unmarshal err，err:%v", err)
		utils.ReturnHTTPError(&c.Controller, 400, "参数错误")
		return
	}
	intAddressID := input.AddressID
	ids := input.IDs
	if ids == "" {
		logs.Debug("input.ids err，err:%v", err)
		utils.ReturnHTTPError(&c.Controller, 400, "没有选择购物车商品")
		return
	}
	//查询用户
	o := orm.NewOrm()
	userTable := new(models.SysUser)
	var user models.SysUser
	err = o.QueryTable(userTable).Filter("uid", getLoginUserID()).One(&user)
	if err == orm.ErrNoRows {
		logs.Debug("no user，err:%v", err)
		utils.ReturnHTTPError(&c.Controller, 400, "用户未登陆")
	}

	//用户地址检查
	addressTable := new(models.NsMemberExpressAddress)
	var address models.NsMemberExpressAddress

	err = o.QueryTable(addressTable).Filter("id", intAddressID).One(&address)
	if err == orm.ErrNoRows {
		logs.Debug("no selectting address，err:%v", err)
		utils.ReturnHTTPError(&c.Controller, 400, "请选择收获地址")
	}
	//用户购物车检查
	idsArray := strings.Split(ids, ",")
	cartTable := new(models.NsCart)
	var carts []models.NsCart
	_, err = o.QueryTable(cartTable).Filter("buyer_id", getLoginUserID()).Filter("cart_id__in", idsArray).All(&carts)
	if err == orm.ErrNoRows {
		utils.ReturnHTTPError(&c.Controller, 400, "请选择商品")
	}
	//开启事务
	var (
		freightPrice    float64
		goodsTotalPrice float64
		orderTotalPrice float64
		actualprice     float64
		orderInfo       models.NsOrder
	)
	_ = o.Begin()
	defer func() {
		if err != nil {
			_ = o.Rollback()
			logs.Debug("this has a err ,err:%v", err)
			utils.ReturnHTTPError(&c.Controller, 400, "删除失败")
			return
		}
		err = o.Commit()
		utils.ReturnHTTPSuccess(&c.Controller, orderInfo)
		return

	}()

	for _, val := range carts {
		goodsTotalPrice += float64(val.Number) * val.Price
	}
	orderTotalPrice = goodsTotalPrice - freightPrice
	actualprice = orderTotalPrice
	orderInfo = models.NsOrder{
		OrderSn:      models.GenerateOrderNumber(),
		OutTradeNo:   models.CreateOutTradeNo(),
		OrderType:    1,
		PaymentType:  10,
		ShippingType: 1,
		OrderFrom:    "3",
		UserName:     user.UserName,
		UserID:       getLoginUserID(),
		Consignee:    address.Name,
		Mobile:       address.Mobile,
		Province:     address.ProvinceID,
		City:         address.CityID,
		District:     address.DistrictID,
		Address:      address.Address,
		ShippingFee:  freightPrice,
		GoodsMoney:   goodsTotalPrice,
		OrderMoney:   orderTotalPrice,
		PayMoney:     actualprice,
		CreateTime:   utils.GetTimestamp(),
	}

	orderID, err := o.Insert(&orderInfo)
	if err != nil {
		logs.Debug("订单提交失败,error:%v", err)
		return
	}
	orderInfo.ID = int(orderID)
	var checkedIDs []int
	for _, item := range carts {
		checkedIDs = append(checkedIDs, item.ID)
		ordergood := models.NsOrderGoods{
			OrderID:        int(orderID),
			GoodsID:        item.GoodsID,
			GoodsName:      item.GoodsName,
			SkuID:          item.SkuID,
			SkuName:        item.SkuName,
			Price:          item.Price,
			Number:         item.Number,
			GoodsMoney:     item.Price * float64(item.Number),
			GoodsPicture:   item.GoodsPicture,
			BuyerID:        item.UserID,
			GoodsType:      1,
			OrderType:      1,
			OrderStatus:    1,
			ShippingStatus: 0,
			RefundType:     1,
		}
		_, _ = o.Insert(&ordergood)
	}

	models.DeleteBuyGoods(getLoginUserID(), checkedIDs)
	utils.ReturnHTTPSuccess(&c.Controller, orderInfo)
}

// OrderExpress OrderExpress
func (c *OrderController) OrderExpress() {
	orderID := c.GetString("orderId")
	intOrderID := utils.String2Int(orderID)

	if orderID == "" {
		logs.Debug("order[%v] lose", intOrderID)
		utils.ReturnHTTPError(&c.Controller, 400, "订单不存在")
		c.ServeJSON()
		return
	}

	latestExpressInfo := models.GetLatestOrderExpress(intOrderID)

	utils.ReturnHTTPSuccess(&c.Controller, latestExpressInfo)
	c.ServeJSON()
}
