package models

import (
	"github.com/astaxie/beego/orm"
)

// ClearBuyGoods ClearBuyGoods
func ClearBuyGoods(userid int) {

	o := orm.NewOrm()

	cartTable := new(NsCart)

	o.QueryTable(cartTable).Filter("buyer_id", userid).Delete()

}

// DeleteBuyGoods DeleteBuyGoods
func DeleteBuyGoods(userid int, ids []int) {

	o := orm.NewOrm()

	cartTable := new(NsCart)

	o.QueryTable(cartTable).Filter("buyer_id", userid).Filter("cart_id__in", ids).Delete()

}
