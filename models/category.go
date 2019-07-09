package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/utils"
)

func GetChildCategoryId(categoryid int) []int64 {

	o := orm.NewOrm()
	categorytable := new(NsGoodsCategory)
	var childids []orm.Params
	o.QueryTable(categorytable).Filter("pid", categoryid).Limit(10000).Values(&childids, "category_id")
	childintids := utils.ExactMapValues2Int64Array(childids, "Id")
	return childintids
}

func GetCategoryWhereIn(categoryid int) []int64 {

	childintids := GetChildCategoryId(categoryid)
	childintids = append(childintids, int64(categoryid))
	return childintids
}
