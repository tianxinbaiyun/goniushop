package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/utils"
)

// IndexController IndexController
type IndexController struct {
	beego.Controller
}

// newCategoryList newCategoryList
type newCategoryList struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	GoodsList []orm.Params `json:"goodsList"`
}

// IndexRtnJSON IndexRtnJSON
type IndexRtnJSON struct {
	Banners      []models.NsPlatformAdv `json:"banner"`
	NewGoods     []models.NsNewGoods    `json:"newGoodsList"`
	HotGoods     []models.NsNewGoods    `json:"hotGoodsList"`
	BrandList    []models.NsGoodsBrand  `json:"brandList"`
	CategoryList []newCategoryList      `json:"categoryList"`
}

// updateJSONKeysIndex updateJSONKeysIndex
func updateJSONKeysIndex(values []orm.Params) {

	for _, val := range values {
		for k, v := range val {
			switch k {
			case "Id":
				delete(val, k)
				val["id"] = v
			case "Name":
				delete(val, k)
				val["name"] = v
			case "Picture":
				delete(val, k)
				val["picture"] = v
			case "RetailPrice":
				delete(val, k)
				val["retail_price"] = v
			}
		}
	}
}

// IndexIndex IndexIndex
func (c *IndexController) IndexIndex() {
	o := orm.NewOrm()

	var banners []models.NsPlatformAdv
	ad := new(models.NsPlatformAdv)
	_, _ = o.QueryTable(ad).Filter("ap_id", 1105).All(&banners)

	var channels []models.SysWeixinMenu
	channel := new(models.SysWeixinMenu)
	_, _ = o.QueryTable(channel).OrderBy("sort").All(&channels)

	newGoods := models.GetGoodsList("g.is_new=1", 4)
	hotGoods := models.GetGoodsList("g.is_hot=1", 3)

	var brandList []models.NsGoodsBrand
	brand := new(models.NsGoodsBrand)
	_, _ = o.QueryTable(brand).Filter("brand_recommend", 1).OrderBy("sort").Limit(4).All(&brandList)

	var categoryList []models.NsGoodsCategory
	category := new(models.NsGoodsCategory)
	_, _ = o.QueryTable(category).Filter("pid", 0).Exclude("category_name", "推荐").All(&categoryList)

	var newList []newCategoryList

	for _, categoryItem := range categoryList {
		var mapIDs []orm.Params
		_, _ = o.QueryTable(category).Filter("pid", categoryItem.ID).Values(&mapIDs, "category_id")

		// var valIDs []int64
		// for _, value := range mapIDs {
		// 	valIDs = append(valIDs, value["ID"].(int64))
		// }

		valIDs := utils.ExactMapValues2Int64Array(mapIDs, "ID")
		goods := new(models.NsGoods)
		var categoryGoods []orm.Params
		_, _ = o.QueryTable(goods).Filter("category_id__in", valIDs).Limit(7).Values(&categoryGoods, "goods_id", "goods_name", "picture", "price")
		updateJSONKeysIndex(categoryGoods)
		newList = append(newList, newCategoryList{categoryItem.ID, categoryItem.Name, categoryGoods})
	}

	utils.ReturnHTTPSuccess(&c.Controller, IndexRtnJSON{banners, newGoods, hotGoods, brandList, newList})

	c.ServeJSON()

}
