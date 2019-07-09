package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/utils"
)

type BrandController struct {
	beego.Controller
}

func (this *BrandController) Brand_List() {

	page := this.GetString("page")
	size := this.GetString("size")

	var intsize int = 10
	if size != "" {
		intsize = utils.String2Int(size)
	}

	var intpage int = 1
	if page != "" {
		intpage = utils.String2Int(page)
	}

	o := orm.NewOrm()
	brandtable := new(models.NsGoodsBrand)
	var brands []orm.Params
	o.QueryTable(brandtable).Values(&brands, "brand_id", "brand_name", "brand_initial", "brand_pic")

	pagedata := utils.GetPageData(brands, intpage, intsize)

	utils.ReturnHTTPSuccess(&this.Controller, pagedata)
	this.ServeJSON()

}

type BrandDetailRtnJson struct {
	Data models.NsGoodsBrand
}

func (this *BrandController) Brand_Detail() {
	id := this.GetString("id")
	intid := utils.String2Int(id)

	o := orm.NewOrm()
	brandtable := new(models.NsGoodsBrand)
	var brand models.NsGoodsBrand

	o.QueryTable(brandtable).Filter("brand_id", intid).One(&brand)

	utils.ReturnHTTPSuccess(&this.Controller, BrandDetailRtnJson{brand})
	this.ServeJSON()
}
