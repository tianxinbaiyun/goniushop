package models

import (
	"github.com/astaxie/beego/orm"
)

// IsUserHasCollect IsUserHasCollect
func IsUserHasCollect(userID int, favType string, valueID int) int {

	o := orm.NewOrm()

	var collect NsMemberFavorites
	collectTable := new(NsMemberFavorites)

	err := o.QueryTable(collectTable).Filter("fav_type", favType).Filter("fav_id", valueID).Filter("uid", userID).One(&collect)

	if err == nil {
		return 1
	}
	return 0

}
