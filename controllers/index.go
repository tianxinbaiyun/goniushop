package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/utils"
)

type IndexController struct {
	beego.Controller
}

type newCategoryList struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"`
	GoodsList []orm.Params `json:"goodsList"`
}

type IndexRtnJson struct {
	Banners      []models.NsPlatformAdv `json:"banner"`
	Newgoods     []models.NsNewGoods    `json:"newGoodsList"`
	Hotgoods     []models.NsNewGoods    `json:"hotGoodsList"`
	BrandList    []models.NsGoodsBrand  `json:"brandList"`
	CategoryList []newCategoryList      `json:"categoryList"`
}

func updateJsonKeysIndex(vals []orm.Params) {

	for _, val := range vals {
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

func (this *IndexController) Index_Index() {
	o := orm.NewOrm()

	var banners []models.NsPlatformAdv
	ad := new(models.NsPlatformAdv)
	o.QueryTable(ad).Filter("ap_id", 1105).All(&banners)

	var channels []models.SysWeixinMenu
	channel := new(models.SysWeixinMenu)
	o.QueryTable(channel).OrderBy("sort").All(&channels)

	newgoods := models.GetGoodsList("g.is_new=1", 4)
	hotgoods := models.GetGoodsList("g.is_hot=1", 3)

	var brandList []models.NsGoodsBrand
	brand := new(models.NsGoodsBrand)
	o.QueryTable(brand).Filter("brand_recommend", 1).OrderBy("sort").Limit(4).All(&brandList)

	var categoryList []models.NsGoodsCategory
	category := new(models.NsGoodsCategory)
	o.QueryTable(category).Filter("pid", 0).Exclude("category_name", "推荐").All(&categoryList)

	var newList []newCategoryList

	for _, categoryItem := range categoryList {
		var mapids []orm.Params
		o.QueryTable(category).Filter("pid", categoryItem.Id).Values(&mapids, "category_id")

		// var valIds []int64
		// for _, value := range mapids {
		// 	valIds = append(valIds, value["Id"].(int64))
		// }

		valIds := utils.ExactMapValues2Int64Array(mapids, "Id")
		goods := new(models.NsGoods)
		var categorygoods []orm.Params
		o.QueryTable(goods).Filter("category_id__in", valIds).Limit(7).Values(&categorygoods, "goods_id", "goods_name", "picture", "price")
		updateJsonKeysIndex(categorygoods)
		newList = append(newList, newCategoryList{categoryItem.Id, categoryItem.Name, categorygoods})
	}

	utils.ReturnHTTPSuccess(&this.Controller, IndexRtnJson{banners, newgoods, hotgoods, brandList, newList})

	this.ServeJSON()

}
