package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/utils"
)

// CatalogController CatalogController
type CatalogController struct {
	beego.Controller
}

// CurCategory CurCategory
type CurCategory struct {
	models.NsGoodsCategory
	SubCategoryList []models.NsGoodsCategory `json:"subCategoryList"`
}

// CateLogIndexRtnJSON CateLogIndexRtnJSON
type CateLogIndexRtnJSON struct {
	CategoryList    []models.NsGoodsCategory `json:"categoryList"`
	CurrentCategory CurCategory              `json:"currentCategory"`
	CategoryAd      models.NsPlatformAdv     `json:"categoryAd"`
}

// CatalogIndex CatalogIndex
func (c *CatalogController) CatalogIndex() {

	categoryID := c.GetString("id")
	// orm.Debug = true
	o := orm.NewOrm()
	//获取广告
	var ad models.NsPlatformAdv
	adTable := new(models.NsPlatformAdv)
	o.QueryTable(adTable).Filter("ap_id", 1162).One(&ad)

	//获取分类
	var categories []models.NsGoodsCategory
	categoryTable := new(models.NsGoodsCategory)
	o.QueryTable(categoryTable).Filter("pid", 0).Limit(10).All(&categories)

	var currentCategory *models.NsGoodsCategory
	if categoryID != "" {
		o.QueryTable(categoryTable).Filter("category_id", categoryID).One(currentCategory)
	}

	currentCategory = &categories[0]

	curCategory := new(CurCategory)
	if currentCategory != nil && currentCategory.ID > 0 {
		var subCategories []models.NsGoodsCategory
		o.QueryTable(categoryTable).Filter("pid", currentCategory.ID).All(&subCategories)
		curCategory.SubCategoryList = subCategories
		curCategory.NsGoodsCategory = *currentCategory
	}
	utils.ReturnHTTPSuccess(&c.Controller, CateLogIndexRtnJSON{categories, *curCategory, ad})
	c.ServeJSON()
}

// CateLogCurRtnJSON CateLogCurRtnJSON
type CateLogCurRtnJSON struct {
	CurrentCategory CurCategory `json:"currentCategory"`
}

// CatalogCurrent CatalogCurrent
func (c *CatalogController) CatalogCurrent() {

	categoryID := c.GetString("id")

	o := orm.NewOrm()
	categoryTable := new(models.NsGoodsCategory)
	currentCategory := new(models.NsGoodsCategory)
	if categoryID != "" {
		o.QueryTable(categoryTable).Filter("category_id", categoryID).One(currentCategory)
	}

	curCategory := new(CurCategory)
	if currentCategory.ID > 0 {
		var subCategories []models.NsGoodsCategory
		o.QueryTable(categoryTable).Filter("pid", currentCategory.ID).All(&subCategories)
		curCategory.SubCategoryList = subCategories
		curCategory.NsGoodsCategory = *currentCategory
	}

	utils.ReturnHTTPSuccess(&c.Controller, CateLogCurRtnJSON{*curCategory})
	c.ServeJSON()

}
