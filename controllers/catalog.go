package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/utils"
)

type CatalogController struct {
	beego.Controller
}

type CurCategory struct {
	models.NsGoodsCategory
	SubCategoryList []models.NsGoodsCategory `json:"subCategoryList"`
}

type CateLogIndexRtnJson struct {
	CategoryList    []models.NsGoodsCategory `json:"categoryList"`
	CurrentCategory CurCategory              `json:"currentCategory"`
	CategoryAd      models.NsPlatformAdv     `json:"categoryAd"`
}

func (this *CatalogController) Catalog_Index() {

	categoryId := this.GetString("id")
	// orm.Debug = true
	o := orm.NewOrm()
	//获取广告
	var ad models.NsPlatformAdv
	adtable := new(models.NsPlatformAdv)
	o.QueryTable(adtable).Filter("ap_id", 1162).One(&ad)

	//获取分类
	var categories []models.NsGoodsCategory
	categorytable := new(models.NsGoodsCategory)
	o.QueryTable(categorytable).Filter("pid", 0).Limit(10).All(&categories)

	var currentCategory *models.NsGoodsCategory = nil
	if categoryId != "" {
		o.QueryTable(categorytable).Filter("category_id", categoryId).One(currentCategory)
	}
	if currentCategory == nil {
		currentCategory = &categories[0]
	}
	curCategory := new(CurCategory)
	if currentCategory != nil && currentCategory.Id > 0 {
		var subCategories []models.NsGoodsCategory
		o.QueryTable(categorytable).Filter("pid", currentCategory.Id).All(&subCategories)
		curCategory.SubCategoryList = subCategories
		curCategory.NsGoodsCategory = *currentCategory
	}
	utils.ReturnHTTPSuccess(&this.Controller, CateLogIndexRtnJson{categories, *curCategory, ad})
	this.ServeJSON()
}

type CateLogCurRtnJson struct {
	CurrentCategory CurCategory `json:"currentCategory"`
}

func (this *CatalogController) Catalog_Current() {

	categoryId := this.GetString("id")

	o := orm.NewOrm()
	categorytable := new(models.NsGoodsCategory)
	currentCategory := new(models.NsGoodsCategory)
	if categoryId != "" {
		o.QueryTable(categorytable).Filter("category_id", categoryId).One(currentCategory)
	}

	curCategory := new(CurCategory)
	if currentCategory != nil && currentCategory.Id > 0 {
		var subCategories []models.NsGoodsCategory
		o.QueryTable(categorytable).Filter("pid", currentCategory.Id).All(&subCategories)
		curCategory.SubCategoryList = subCategories
		curCategory.NsGoodsCategory = *currentCategory
	}

	utils.ReturnHTTPSuccess(&this.Controller, CateLogCurRtnJson{*curCategory})
	this.ServeJSON()

}
