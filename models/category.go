package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/utils"
)

// GetChildCategoryID GetChildCategoryID
func GetChildCategoryID(categoryID int) []int64 {

	o := orm.NewOrm()
	categoryTable := new(NsGoodsCategory)
	var childids []orm.Params
	_, _ = o.QueryTable(categoryTable).Filter("pid", categoryID).Limit(10000).Values(&childids, "category_id")
	childIntIDs := utils.ExactMapValues2Int64Array(childids, "Id")
	return childIntIDs
}

// GetCategoryWhereIn GetCategoryWhereIn
func GetCategoryWhereIn(categoryID int) []int64 {

	childIntIDs := GetChildCategoryID(categoryID)
	childIntIDs = append(childIntIDs, int64(categoryID))
	return childIntIDs
}
