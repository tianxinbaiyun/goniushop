package models

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/utils"
)

func GetOrderStatusText(orderid int) string {

	o := orm.NewOrm()
	ordertable := new(NsOrder)
	var order NsOrder
	o.QueryTable(ordertable).Filter("id", orderid).One(&order)
	var statustext string = "未付款"
	switch order.OrderStatus {
	case 0:
		statustext = "未付款"
	}
	return statustext
}

type OrderHandleOption struct {
	Cancel   bool `json:"cancel"`
	Delete   bool `json:"delete"`
	Pay      bool `json:"pay"`
	Comment  bool `json:"comment"`
	Delivery bool `json:"delivery"`
	Confirm  bool `json:"confirm"`
	Return   bool `json:"return"`
	Buy      bool `json:"buy"`
}

func GetOrderHandleOption(orderid int) OrderHandleOption {

	// 订单流程：下单成功－》支付订单－》发货－》收货－》评论
	// 订单相关状态字段设计，采用单个字段表示全部的订单状态
	// 1xx表示订单取消和删除等状态 0订单创建成功等待付款，101订单已取消，102订单已删除
	// 2xx表示订单支付状态,201订单已付款，等待发货
	// 3xx表示订单物流相关状态,300订单已发货，301用户确认收货
	// 4xx表示订单退换货相关的状态,401没有发货，退款402,已收货，退款退货
	// 如果订单已经取消或是已完成，则可删除和再次购买

	var handoption OrderHandleOption = OrderHandleOption{false, false, false, false, false, false, false, false}

	o := orm.NewOrm()
	ordertable := new(NsOrder)
	var order NsOrder
	o.QueryTable(ordertable).Filter("id", orderid).One(&order)

	switch order.OrderStatus {
	case 0:
		handoption.Cancel = true
		handoption.Pay = true
	case 101:
		handoption.Delete = true
		handoption.Buy = true
	case 201:
		handoption.Return = true
	case 300:
		handoption.Cancel = true
		handoption.Pay = true
		handoption.Return = true
	case 301:
		handoption.Delete = true
		handoption.Comment = true
		handoption.Buy = true
	}

	return handoption
}

func GenerateOrderNumber() string {

	year := time.Now().Year()     //年
	month := time.Now().Month()   //月
	day := time.Now().Day()       //日
	hour := time.Now().Hour()     //小时
	minute := time.Now().Minute() //分钟
	// second := time.Now().Second() //秒

	stryear := utils.Int2String(year)        //年
	strmonth := utils.Int2String(int(month)) //月
	strday := utils.Int2String(day)          //日
	strhour := utils.Int2String(hour)        //小时
	strminute := utils.Int2String(minute)    //分钟
	// strsecond := utils.Int2String(second)    //秒

	strmonth2 := fmt.Sprintf("%02s", strmonth)
	strday2 := fmt.Sprintf("%02s", strday)
	strhour2 := fmt.Sprintf("%02s", strhour)
	strminute2 := fmt.Sprintf("%02s", strminute)
	// strsecond2 := fmt.Sprintf("%02s", strsecond)

	timeStr := stryear + strmonth2 + strday2 + strhour2 + strminute2

	generateOrderNumber := GetLastOrderSn(timeStr)
	return generateOrderNumber
}
func GetLastOrderSn(timeStr string) (generateOrderNumber string) {
	o := orm.NewOrm()
	ordertable := new(NsOrder)
	var order NsOrder
	o.QueryTable(ordertable).OrderBy("order_id").One(&order)
	lastOrderSn := order.OrderSn
	if timeStr == lastOrderSn[0:12] {
		number, _ := strconv.Atoi(lastOrderSn[12:])
		number++
		generateOrderNumber = fmt.Sprintf("%v%4d", timeStr, number)
		return
	} else {
		generateOrderNumber = fmt.Sprintf("%v0001", timeStr)
		return
	}
}
func CreateOutTradeNo() (outTradeNo string) {
	timeStr := utils.GetTimestamp()
	randStr := rand.Intn(8999) + 1000
	secondStr := rand.Intn(8999) + 1000
	outTradeNo = fmt.Sprintf("%v%v%v", timeStr, randStr, secondStr)
	return
}
