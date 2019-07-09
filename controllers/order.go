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

type OrderController struct {
	beego.Controller
}

//It may need to be refactored.
func GetOrderPageData(rawData []models.NsOrder, page int, size int) utils.PageData {

	count := len(rawData)
	totalpages := (count + size - 1) / size
	var pagedata []models.NsOrder

	for idx := (page - 1) * size; idx < page*size && idx < count; idx++ {
		pagedata = append(pagedata, rawData[idx])
	}

	return utils.PageData{NumsPerPage: size, CurrentPage: page, Count: count, TotalPages: totalpages, Data: pagedata}
}

type OrderListRtnJson struct {
	models.NsOrder
	GoodsList       []models.NsOrderGoods    `json:"goodList"`
	GoodsCount      int                      `json:"goodsCount"`
	OrderStatusText string                   `json:"order_status_text"`
	HandOption      models.OrderHandleOption `json:"handleOption"`
}

func (this *OrderController) Order_List() {
	//请求参数
	page := utils.String2Int(this.GetString("page"))
	pageSize := utils.String2Int(this.GetString("pageSize"))
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	o := orm.NewOrm()
	ordertable := new(models.NsOrder)
	var orders []models.NsOrder
	o.QueryTable(ordertable).Filter("buyer_id", getLoginUserId()).All(&orders)

	firstpagedorders := GetOrderPageData(orders, page, pageSize)

	var rtnorderlist []OrderListRtnJson
	ordergoodstable := new(models.NsOrderGoods)
	var ordergoods []models.NsOrderGoods
	qsordergoods := o.QueryTable(ordergoodstable)
	for _, val := range firstpagedorders.Data.([]models.NsOrder) {
		qsordergoods.Filter("order_id", val.Id).All(&ordergoods)
		var goodscount int
		for _, val := range ordergoods {
			goodscount += val.Number
		}
		orderstatustext := models.GetOrderStatusText(val.Id)
		orderhandoption := models.GetOrderHandleOption(val.Id)
		var orderlistrtn OrderListRtnJson = OrderListRtnJson{val, ordergoods, goodscount, orderstatustext, orderhandoption}

		rtnorderlist = append(rtnorderlist, orderlistrtn)

	}

	firstpagedorders.Data = rtnorderlist

	utils.ReturnHTTPSuccess(&this.Controller, firstpagedorders)
	this.ServeJSON()
}

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

type OrderDetailRtnJson struct {
	OrderInfo    OrderInfo                `json:"orderInfo"`
	OrderGoods   []models.NsOrderGoods    `json:"orderGoods"`
	HandleOption models.OrderHandleOption `json:"handleOption"`
}

func (this *OrderController) Order_Detail() {
	//请求参数
	orderId := this.GetString("orderId")
	intorderId := utils.String2Int(orderId)

	//错误返回
	var err error
	defer func() {
		if err != nil {
			logs.Debug("this order id[%v] has a error ,err:%v", orderId, err)
		}
		this.ServeJSON()
		return
	}()
	o := orm.NewOrm()
	ordertable := new(models.NsOrder)
	var order models.NsOrder
	err = o.QueryTable(ordertable).Filter("id", intorderId).Filter("buyer_id", getLoginUserId()).One(&order)

	if err == orm.ErrNoRows {
		logs.Debug("order losed ,err:%v", err)
		utils.ReturnHTTPError(&this.Controller, 400, "订单不存在")
		return
	}

	var orderinfo OrderInfo = OrderInfo{NsOrder: order}
	orderinfo.ProvinceName = models.GetProvinceName(order.Province)
	orderinfo.CityName = models.GetCityName(order.City)
	orderinfo.DistrictName = models.GetDistrictName(order.District)
	orderinfo.FullRegion = orderinfo.ProvinceName + orderinfo.CityName + orderinfo.DistrictName

	lastestexpressinfo := models.GetLatestOrderExpress(intorderId)
	orderinfo.Express = lastestexpressinfo

	ordergoodstable := new(models.NsOrderGoods)
	var ordergoods []models.NsOrderGoods

	o.QueryTable(ordergoodstable).Filter("order_id", intorderId).All(&ordergoods)

	orderinfo.OrderStatusText = models.GetOrderStatusText(intorderId)
	orderinfo.FormatAddTime = utils.FormatTimestamp(orderinfo.CreateTime, "2006-01-02 03:04:05 PM")
	orderinfo.FormatFinalPlayTime = utils.FormatTimestamp(1234, "04:05")

	if orderinfo.OrderStatus == 0 {
		//todo 订单超时逻辑
		logs.Debug("order timeout ,err:%v", err)
		utils.ReturnHTTPError(&this.Controller, 400, "订单超时")
		return
	}

	handleoption := models.GetOrderHandleOption(intorderId)
	utils.ReturnHTTPSuccess(&this.Controller, OrderDetailRtnJson{
		OrderInfo:    orderinfo,
		OrderGoods:   ordergoods,
		HandleOption: handleoption,
	})
	return
}

type SubmitOrder struct {
	AddressId int    `json:"addressId"`
	Ids       string `json:"ids"`
}

func (this *OrderController) Order_Submit() {
	var err error
	defer func() {
		if err != nil {
			logs.Debug("this has a err ,err:%v", err)
		}
		this.ServeJSON()
		return
	}()
	var input SubmitOrder
	body := this.Ctx.Input.RequestBody
	err = json.Unmarshal(body, &input)
	if err != nil {
		logs.Debug("json Unmarshal err，err:%v", err)
		utils.ReturnHTTPError(&this.Controller, 400, "参数错误")
		return
	}
	intaddressId := input.AddressId
	ids := input.Ids
	if ids == "" {
		logs.Debug("input.ids err，err:%v", err)
		utils.ReturnHTTPError(&this.Controller, 400, "没有选择购物车商品")
		return
	}
	//查询用户
	o := orm.NewOrm()
	usertable := new(models.SysUser)
	var user models.SysUser
	err = o.QueryTable(usertable).Filter("uid", getLoginUserId()).One(&user)
	if err == orm.ErrNoRows {
		logs.Debug("no user，err:%v", err)
		utils.ReturnHTTPError(&this.Controller, 400, "用户未登陆")
	}

	//用户地址检查
	addresstable := new(models.NsMemberExpressAddress)
	var address models.NsMemberExpressAddress

	err = o.QueryTable(addresstable).Filter("id", intaddressId).One(&address)
	if err == orm.ErrNoRows {
		logs.Debug("no selectting address，err:%v", err)
		utils.ReturnHTTPError(&this.Controller, 400, "请选择收获地址")
	}
	//用户购物车检查
	idsArray := strings.Split(ids, ",")
	carttable := new(models.NsCart)
	var carts []models.NsCart
	_, err = o.QueryTable(carttable).Filter("buyer_id", getLoginUserId()).Filter("cart_id__in", idsArray).All(&carts)
	if err == orm.ErrNoRows {
		utils.ReturnHTTPError(&this.Controller, 400, "请选择商品")
	}
	//开启事务
	var (
		freightPrice    float64
		goodstotalprice float64
		ordertotalprice float64
		actualprice     float64
		orderinfo       models.NsOrder
	)
	o.Begin()
	defer func() {
		if err != nil {
			o.Rollback()
			logs.Debug("this has a err ,err:%v", err)
			utils.ReturnHTTPError(&this.Controller, 400, "删除失败")
			return
		} else {
			err = o.Commit()
			utils.ReturnHTTPSuccess(&this.Controller, orderinfo)
			return
		}
	}()

	for _, val := range carts {
		goodstotalprice += float64(val.Number) * val.Price
	}
	ordertotalprice = goodstotalprice - freightPrice
	actualprice = ordertotalprice
	orderinfo = models.NsOrder{
		OrderSn:      models.GenerateOrderNumber(),
		OutTradeNo:   models.CreateOutTradeNo(),
		OrderType:    1,
		PaymentType:  10,
		ShippingType: 1,
		OrderFrom:    "3",
		UserName:     user.UserName,
		UserId:       getLoginUserId(),
		Consignee:    address.Name,
		Mobile:       address.Mobile,
		Province:     address.ProvinceId,
		City:         address.CityId,
		District:     address.DistrictId,
		Address:      address.Address,
		ShippingFee:  freightPrice,
		GoodsMoney:   goodstotalprice,
		OrderMoney:   ordertotalprice,
		PayMoney:     actualprice,
		CreateTime:   utils.GetTimestamp(),
	}

	orderid, err := o.Insert(&orderinfo)
	if err != nil {
		logs.Debug("订单提交失败,error:%v", err)
		return
	}
	orderinfo.Id = int(orderid)
	var checkedIds []int
	for _, item := range carts {
		checkedIds = append(checkedIds, item.Id)
		ordergood := models.NsOrderGoods{
			OrderId:        int(orderid),
			GoodsId:        item.GoodsId,
			GoodsName:      item.GoodsName,
			SkuId:          item.SkuId,
			SkuName:        item.SkuName,
			Price:          item.Price,
			Number:         item.Number,
			GoodsMoney:     item.Price * float64(item.Number),
			GoodsPicture:   item.GoodsPicture,
			BuyerId:        item.UserId,
			GoodsType:      1,
			OrderType:      1,
			OrderStatus:    1,
			ShippingStatus: 0,
			RefundType:     1,
		}
		o.Insert(&ordergood)
	}

	models.DeleteBuyGoods(getLoginUserId(), checkedIds)
	utils.ReturnHTTPSuccess(&this.Controller, orderinfo)
}

func (this *OrderController) Order_Express() {
	orderId := this.GetString("orderId")
	intorderId := utils.String2Int(orderId)

	if orderId == "" {
		logs.Debug("order[%v] losed", intorderId)
		utils.ReturnHTTPError(&this.Controller, 400, "订单不存在")
		this.ServeJSON()
		return
	}

	latestexpressinfo := models.GetLatestOrderExpress(intorderId)

	utils.ReturnHTTPSuccess(&this.Controller, latestexpressinfo)
	this.ServeJSON()
}
