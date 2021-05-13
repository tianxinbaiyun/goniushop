package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/utils"
)

// BrandController BrandController
type BrandController struct {
	beego.Controller
}

// BrandList BrandList
func (c *BrandController) BrandList() {

	page := c.GetString("page")
	size := c.GetString("size")

	var intSize = 10
	if size != "" {
		intSize = utils.String2Int(size)
	}

	var intPage = 1
	if page != "" {
		intPage = utils.String2Int(page)
	}

	o := orm.NewOrm()
	brandTable := new(models.NsGoodsBrand)
	var brands []orm.Params
	_, _ = o.QueryTable(brandTable).Values(&brands, "brand_id", "brand_name", "brand_initial", "brand_pic")

	pageData := utils.GetPageData(brands, intPage, intSize)

	utils.ReturnHTTPSuccess(&c.Controller, pageData)
	c.ServeJSON()

}

// BrandDetailRtnJSON BrandDetailRtnJSON
type BrandDetailRtnJSON struct {
	Data models.NsGoodsBrand
}

// BrandDetail BrandDetail
func (c *BrandController) BrandDetail() {
	id := c.GetString("id")
	intID := utils.String2Int(id)

	o := orm.NewOrm()
	brandTable := new(models.NsGoodsBrand)
	var brand models.NsGoodsBrand

	_ = o.QueryTable(brandTable).Filter("brand_id", intID).One(&brand)

	utils.ReturnHTTPSuccess(&c.Controller, BrandDetailRtnJSON{brand})
	c.ServeJSON()
}
