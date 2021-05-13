package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/services"
)

// GetLatestOrderExpress 获取最新的订单物流信息
func GetLatestOrderExpress(orderID int) services.ExpressRtnInfo {
	var expressInfo = services.ExpressRtnInfo{
		ShipperCode:  "",
		ShipperName:  "",
		LogisticCode: "",
		IsFinish:     0,
		RequestTime:  0,
		Traces:       make([]services.Traces, 0),
	}
	o := orm.NewOrm()

	//获取订单快递单
	orderExpressTable := new(NsOrderGoodsExpress)
	var orderExpress NsOrderGoodsExpress
	err := o.QueryTable(orderExpressTable).Filter("order_id", orderID).One(&orderExpress)
	if err == orm.ErrNoRows {
		return expressInfo
	}

	if orderExpress.ShipperCode == "" {
		return expressInfo
	}
	// expressInfo.ShipperCode = orderExpress.ShipperCode
	// expressInfo.ShipperName = orderExpress.ShipperName

	// expressInfo.RequestTime = utils.GetTimestamp()

	// expressServiceRes := services.QueryExpress(expressInfo.ShipperCode, expressInfo.LogisticCode, "")
	// nowTime := utils.GetTimestamp()

	// if expressServiceRes.Success {
	// 	expressInfo.Traces = expressServiceRes.Traces
	// 	expressInfo.IsFinish = expressServiceRes.IsFinish
	// 	expressInfo.RequestTime = nowTime
	// }

	// traces, _ := json.Marshal(expressInfo.Traces)

	// o.QueryTable(orderExpressTable).Filter("id", orderExpress.ID).Update(orm.Params{
	// 	"request_time":  nowTime,
	// 	"update_time":   nowTime,

	// 	"request_count": orm.ColValue(orm.ColAdd, 1)})

	return expressInfo
}
