package models

import (
	"github.com/astaxie/beego/orm"
)

func ClearBuyGoods(userid int) {

	o := orm.NewOrm()

	carttable := new(NsCart)

	o.QueryTable(carttable).Filter("buyer_id", userid).Delete()

}
func DeleteBuyGoods(userid int, ids []int) {

	o := orm.NewOrm()

	carttable := new(NsCart)

	o.QueryTable(carttable).Filter("buyer_id", userid).Filter("cart_id__in", ids).Delete()

}
