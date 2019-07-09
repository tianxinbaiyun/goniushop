package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/services"
	"github.com/tianxinbaiyun/goniushop/utils"
)

type PayController struct {
	beego.Controller
}

func (this *PayController) Pay_Prepay() {
	orderid := this.GetString("orderId")
	intorderid := utils.String2Int(orderid)

	o := orm.NewOrm()
	ordertable := new(models.NsOrder)
	var order models.NsOrder

	err := o.QueryTable(ordertable).Filter("id", intorderid).One(&order)
	if err == orm.ErrNoRows {
		this.CustomAbort(400, "订单已取消")
	}

	if order.PayStatus != 0 {
		this.CustomAbort(400, "订单已支付，请不要重复操作")
	}

	usertable := new(models.SysUser)
	var user models.SysUser
	err = o.QueryTable(usertable).Filter("id", order.UserId).One(&user)

	if err != orm.ErrNoRows && user.WxOpenid == "" {
		this.Abort("微信支付失败")
	}

	payinfo := services.PayInfo{
		OpenId:     user.WxOpenid,
		Body:       "order NO: " + order.OrderSn,
		OutTradeNo: order.OrderSn,
		TotalFee:   int64(order.PayMoney * 100),
	}

	params, err := services.CreateUnifiedOrder(payinfo)

	if err != nil {
		this.Abort("微信支付失败")
	} else {

		utils.ReturnHTTPSuccess(&this.Controller, params)
		this.ServeJSON()
	}
}
