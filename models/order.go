package models

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/utils"
)

// GetOrderStatusText GetOrderStatusText
func GetOrderStatusText(orderID int) string {

	o := orm.NewOrm()
	orderTable := new(NsOrder)
	var order NsOrder
	o.QueryTable(orderTable).Filter("id", orderID).One(&order)
	var statusText = "未付款"
	switch order.OrderStatus {
	case 0:
		statusText = "未付款"
	}
	return statusText
}

// OrderHandleOption OrderHandleOption
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

// GetOrderHandleOption GetOrderHandleOption
func GetOrderHandleOption(orderID int) OrderHandleOption {

	// 订单流程：下单成功－》支付订单－》发货－》收货－》评论
	// 订单相关状态字段设计，采用单个字段表示全部的订单状态
	// 1xx表示订单取消和删除等状态 0订单创建成功等待付款，101订单已取消，102订单已删除
	// 2xx表示订单支付状态,201订单已付款，等待发货
	// 3xx表示订单物流相关状态,300订单已发货，301用户确认收货
	// 4xx表示订单退换货相关的状态,401没有发货，退款402,已收货，退款退货
	// 如果订单已经取消或是已完成，则可删除和再次购买

	var handOption = OrderHandleOption{false, false, false, false, false, false, false, false}

	o := orm.NewOrm()
	orderTable := new(NsOrder)
	var order NsOrder
	o.QueryTable(orderTable).Filter("id", orderID).One(&order)

	switch order.OrderStatus {
	case 0:
		handOption.Cancel = true
		handOption.Pay = true
	case 101:
		handOption.Delete = true
		handOption.Buy = true
	case 201:
		handOption.Return = true
	case 300:
		handOption.Cancel = true
		handOption.Pay = true
		handOption.Return = true
	case 301:
		handOption.Delete = true
		handOption.Comment = true
		handOption.Buy = true
	}

	return handOption
}

// GenerateOrderNumber GenerateOrderNumber
func GenerateOrderNumber() string {

	year := time.Now().Year()     //年
	month := time.Now().Month()   //月
	day := time.Now().Day()       //日
	hour := time.Now().Hour()     //小时
	minute := time.Now().Minute() //分钟
	// second := time.Now().Second() //秒

	strYear := utils.Int2String(year)        //年
	strMonth := utils.Int2String(int(month)) //月
	strDay := utils.Int2String(day)          //日
	strHour := utils.Int2String(hour)        //小时
	strMinute := utils.Int2String(minute)    //分钟
	// strSecond := utils.Int2String(second)    //秒

	strMonth2 := fmt.Sprintf("%02s", strMonth)
	strDay2 := fmt.Sprintf("%02s", strDay)
	strHour2 := fmt.Sprintf("%02s", strHour)
	strMinute2 := fmt.Sprintf("%02s", strMinute)
	// strSecond2 := fmt.Sprintf("%02s", strSecond)

	timeStr := strYear + strMonth2 + strDay2 + strHour2 + strMinute2

	generateOrderNumber := GetLastOrderSn(timeStr)
	return generateOrderNumber
}

// GetLastOrderSn GetLastOrderSn
func GetLastOrderSn(timeStr string) (generateOrderNumber string) {
	o := orm.NewOrm()
	orderTable := new(NsOrder)
	var order NsOrder
	o.QueryTable(orderTable).OrderBy("order_id").One(&order)
	lastOrderSn := order.OrderSn
	if timeStr == lastOrderSn[0:12] {
		number, _ := strconv.Atoi(lastOrderSn[12:])
		number++
		generateOrderNumber = fmt.Sprintf("%v%4d", timeStr, number)
		return
	}
	generateOrderNumber = fmt.Sprintf("%v0001", timeStr)
	return

}

// CreateOutTradeNo CreateOutTradeNo
func CreateOutTradeNo() (outTradeNo string) {
	timeStr := utils.GetTimestamp()
	randStr := rand.Intn(8999) + 1000
	secondStr := rand.Intn(8999) + 1000
	outTradeNo = fmt.Sprintf("%v%v%v", timeStr, randStr, secondStr)
	return
}
