package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/services"
	"github.com/tianxinbaiyun/goniushop/utils"
)

// PayController PayController
type PayController struct {
	beego.Controller
}

// PayPrepay PayPrepay
func (c *PayController) PayPrepay() {
	orderID := c.GetString("orderId")
	intOrderID := utils.String2Int(orderID)

	o := orm.NewOrm()
	orderTable := new(models.NsOrder)
	var order models.NsOrder

	err := o.QueryTable(orderTable).Filter("id", intOrderID).One(&order)
	if err == orm.ErrNoRows {
		c.CustomAbort(400, "订单已取消")
	}

	if order.PayStatus != 0 {
		c.CustomAbort(400, "订单已支付，请不要重复操作")
	}

	userTable := new(models.SysUser)
	var user models.SysUser
	err = o.QueryTable(userTable).Filter("id", order.UserID).One(&user)

	if err != orm.ErrNoRows && user.WxOpenid == "" {
		c.Abort("微信支付失败")
	}

	payInfo := services.PayInfo{
		OpenID:     user.WxOpenid,
		Body:       "order NO: " + order.OrderSn,
		OutTradeNo: order.OrderSn,
		TotalFee:   int64(order.PayMoney * 100),
	}

	params, err := services.CreateUnifiedOrder(payInfo)

	if err != nil {
		c.Abort("微信支付失败")
	} else {

		utils.ReturnHTTPSuccess(&c.Controller, params)
		c.ServeJSON()
	}
}
