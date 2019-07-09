package models

import (
	"github.com/astaxie/beego/orm"
)

func IsUserHasCollect(userId int, favType string, valueId int) int {

	o := orm.NewOrm()

	var collect NsMemberFavorites
	collecttable := new(NsMemberFavorites)

	err := o.QueryTable(collecttable).Filter("fav_type", favType).Filter("fav_id", valueId).Filter("uid", userId).One(&collect)

	if err == nil {
		return 1
	} else {
		return 0
	}

}
