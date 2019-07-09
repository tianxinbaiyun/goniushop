package controllers

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/utils"
)

type GoodsController struct {
	beego.Controller
}

type SkuRtnJson struct {
	ProductList []models.NsGoodsSku `json:"productList"`
}

type DetailRtnJson struct {
	SkuRtnJson
	Goods     models.NsGoods           `json:"info"`
	Galleries []models.SysAlbumPicture `json:"gallery"`
	Attribute []Attribute              `json:"attribute"`
	// Issues         []models.NsGoodsIssue `json:"issue"`
	UserHasCollect int                 `json:"userHasCollect"`
	Comment        Comment             `json:"comment"`
	Brand          models.NsGoodsBrand `json:"brand"`
}

type CategoryRtnJson struct {
	CurCategory     models.NsGoodsCategory   `json:"currentCategory"`
	ParentCategory  models.NsGoodsCategory   `json:"parentCategory"`
	BrotherCategory []models.NsGoodsCategory `json:"brotherCategory"`
}

type Attribute struct {
	Value         string `json:"value"`
	AttrValueName string `json:"attr_value_name"`
}

type CommentUser struct {
	NickName string `json:"nick_name"`
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
}

type Comment struct {
	Count int64                  `json:"count"`
	Data  models.NsGoodsEvaluate `json:"data"`
}

type FilterCategory struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Checked bool   `json:"checked"`
}

type ListRtnJson struct {
	utils.PageData
	FilterCategories []FilterCategory `json:"filterCategory"`
	GoodsList        []orm.Params     `json:"goodsList"`
}

type Banner struct {
	Url     string `json:"url"`
	Name    string `json:"name"`
	Img_url string `json:"imgurl"`
}

type NewRtnJson struct {
	BannerInfo Banner `json:"bannerinfo"`
}

type HotRtnJson struct {
	BannerInfo Banner `json:"bannerinfo"`
}

type CountRtnJson struct {
	GoodsCount int64 `json:"goodsCount"`
}

func updateJsonKeysGoods(vals []orm.Params) {

	for _, val := range vals {
		for k, v := range val {
			switch k {
			case "Id":
				delete(val, k)
				val["id"] = v
			case "Name":
				delete(val, k)
				val["name"] = v
			case "ListPicUrl":
				delete(val, k)
				val["list_pic_url"] = v
			case "RetailPrice":
				delete(val, k)
				val["retail_price"] = v
			}
		}
	}
}

func (this *GoodsController) Goods_Index() {
	o := orm.NewOrm()

	var goods []models.NsGoods
	good := new(models.NsGoods)
	o.QueryTable(good).All(&goods)

	utils.ReturnHTTPSuccess(&this.Controller, goods)
	this.ServeJSON()
}

func (this *GoodsController) Goods_Sku() {

	goodsId := this.GetString("id")
	goodsId_int := utils.String2Int(goodsId)

	// plist := models.GetProductList(goodsId_int)
	plist := models.GetGoodSku(goodsId_int)

	utils.ReturnHTTPSuccess(&this.Controller, SkuRtnJson{ProductList: plist})
	this.ServeJSON()
}

func (this *GoodsController) Goods_Detail() {
	goodsId := this.GetString("id")
	intGoodsId := utils.String2Int(goodsId)

	o := orm.NewOrm()

	var goodone models.NsGoods
	good := new(models.NsGoods)
	o.QueryTable(good).Filter("goods_id", intGoodsId).One(&goodone)

	var galleries []models.SysAlbumPicture
	gallerie := new(models.SysAlbumPicture)
	imgIds := strings.Split(goodone.ImgIdArray, ",")
	o.QueryTable(gallerie).Filter("pic_id__in", imgIds).Limit(4).All(&galleries)

	qb, _ := orm.NewQueryBuilder("mysql")
	var attributes []Attribute
	qb.Select("a.value", "a.attr_value_name").
		From(" ns_attribute_value a").
		InnerJoin("ns_goods_attribute ga").On("ga.attr_value_id = a.attr_value_id").
		Where("ga.goods_id =" + goodsId).GroupBy("a.attr_value_id").Asc()
	sql := qb.String()
	o.Raw(sql).QueryRows(&attributes)

	var brandone models.NsGoodsBrand
	brand := new(models.NsGoodsBrand)
	o.QueryTable(brand).Filter("brand_id", goodone.BrandId).One(&brandone)

	comment := new(models.NsGoodsEvaluate)
	commentCount, _ := o.QueryTable(comment).Filter("goods_id", intGoodsId).Filter("is_show", 1).Count()
	var hotcommentone models.NsGoodsEvaluate
	o.QueryTable(comment).Filter("goods_id", intGoodsId).Filter("is_show", 1).One(&hotcommentone)
	commentval := Comment{Count: commentCount, Data: hotcommentone}
	loginuserid := getLoginUserId()

	userhascollect := models.IsUserHasCollect(loginuserid, `goods`, intGoodsId)

	// models.AddFootprint(loginuserid, intGoodsId)

	plist := models.GetGoodSku(intGoodsId)

	utils.ReturnHTTPSuccess(&this.Controller, DetailRtnJson{Goods: goodone, Galleries: galleries, Attribute: attributes,
		UserHasCollect: userhascollect, Comment: commentval, Brand: *brand,
		SkuRtnJson: SkuRtnJson{ProductList: plist}})
	this.ServeJSON()
}
func (this *GoodsController) Goods_Category() {

	goodsId := this.GetString("id")
	intgoogsid := utils.String2Int(goodsId)

	o := orm.NewOrm()
	var curcategory models.NsGoodsCategory
	var parentcategory models.NsGoodsCategory
	var brothercategory []models.NsGoodsCategory

	category := new(models.NsGoodsCategory)

	o.QueryTable(category).Filter("id", intgoogsid).One(&curcategory)
	o.QueryTable(category).Filter("id", curcategory.ParentId).One(&parentcategory)
	o.QueryTable(category).Filter("pid", curcategory.ParentId).All(&brothercategory)

	utils.ReturnHTTPSuccess(&this.Controller, CategoryRtnJson{CurCategory: curcategory,
		ParentCategory: parentcategory, BrotherCategory: brothercategory})
	this.ServeJSON()
}
func (this *GoodsController) Goods_List() {
	categoryId := this.GetString("categoryId")
	brandId := this.GetString("brandId")
	keyword := this.GetString("keyword")
	isNew := this.GetString("isNew")
	isHot := this.GetString("isHot")
	page := this.GetString("page")
	size := this.GetString("size")
	sort := this.GetString("sort")
	order := this.GetString("order")

	var intsize int = 10
	if size != "" {
		intsize = utils.String2Int(size)
	}

	var intpage int = 1
	if page != "" {
		intpage = utils.String2Int(page)
	}

	o := orm.NewOrm()
	var categoryids []orm.Params
	var list []orm.Params

	//构建查询语句
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("g.goods_id", "g.goods_name", "g.brand_id", "g.category_id", "g.price", "p.pic_cover", "p.pic_cover_mid", "p.pic_cover_big").
		From("ns_goods g").
		LeftJoin("sys_album_picture p").
		On("g.picture = p.pic_id").
		Where("g.state=1")
	goodstable := new(models.NsGoods)
	rs := o.QueryTable(goodstable)

	//参数处理
	if isNew != "" {
		qb.And(fmt.Sprintf("is_new=%v", isNew))
		rs = rs.Filter("is_new", isNew)
	}
	if isHot != "" {
		qb.And(fmt.Sprintf("is_hot=%v", isHot))
		rs = rs.Filter("is_hot", isHot)
	}
	keyword, _ = url.QueryUnescape(keyword)
	if keyword != "" {
		qb.And(fmt.Sprintf("goods_name like '{%v%}'", keyword))
		rs = rs.Filter("goods_name__icontains", keyword)
	}
	if brandId != "" {
		qb.And(fmt.Sprintf("brand_id = %v", brandId))
		rs = rs.Filter("brand_id", brandId)
	}

	//分类处理

	rs.Limit(10000).Values(&categoryids, "category_id")
	categoryintids := utils.ExactMapValues2Int64Array(categoryids, "CategoryId")
	var filterCategories = []FilterCategory{FilterCategory{Id: 0, Name: "全部", Checked: false}}
	if len(categoryintids) > 0 {
		var parentids []orm.Params
		categorytable := new(models.NsGoodsCategory)
		o.QueryTable(categorytable).Filter("category_id__in", categoryintids).Limit(10000).Values(&parentids, "pid")
		parentintids := utils.ExactMapValues2Int64Array(parentids, "ParentId")

		var parentcategories []orm.Params
		o.QueryTable(categorytable).Filter("category_id__in", parentintids).OrderBy("sort").Values(&parentcategories, "category_id", "category_name")

		for _, value := range parentcategories {
			id := value["Id"].(int64)
			checked := (categoryId == "" && id == 0)

			filterCategories = append(filterCategories, FilterCategory{Id: id, Name: value["Name"].(string), Checked: checked})
		}
	}

	//商品查询
	if categoryId != "" {
		intcategoryId := utils.String2Int(categoryId)
		if intcategoryId >= 0 {
			qb.And(fmt.Sprintf("category_id_1 = %v or category_id_2=%v or category_id_3 = %v", intcategoryId, intcategoryId, intcategoryId))
		}

	}
	if sort == "price" {
		orderstr := "price"
		if order == "desc" {
			orderstr = "-" + orderstr
		}
		qb.OrderBy(orderstr)
	} else {
		qb.OrderBy("goods_id")
	}
	sql := qb.String()
	o.Raw(sql).Values(&list)
	// logs.Debug("list success,sql:%v,list:%v", sql, list)

	pageData := utils.GetPageData(list, intpage, intsize)
	fmt.Println(pageData)

	utils.ReturnHTTPSuccess(&this.Controller, ListRtnJson{PageData: pageData, FilterCategories: filterCategories, GoodsList: pageData.Data.([]orm.Params)})
	this.ServeJSON()
}

func (this *GoodsController) Goods_Filter() {

	categoryId := this.GetString("categoryId")
	keyword := this.GetString("keyword")
	isNew := this.GetString("isNew")
	isHot := this.GetString("isHot")

	o := orm.NewOrm()
	goodstable := new(models.NsGoods)
	rs := o.QueryTable(goodstable)

	if categoryId != "" {
		intcategoryId := utils.String2Int(categoryId)
		rs = rs.Filter("category_id__in", models.GetChildCategoryId(intcategoryId))
	}
	if isNew != "" {
		rs = rs.Filter("is_new", isNew)
	}
	if isHot != "" {
		rs = rs.Filter("is_hot", isHot)
	}
	if keyword != "" {
		rs = rs.Filter("icontains", keyword)
	}

	var filterCategories = []FilterCategory{FilterCategory{Id: 0, Name: "全部"}}

	var categoryids []orm.Params
	rs.Limit(10000).Values(&categoryids, "category_id")
	categoryintids := utils.ExactMapValues2Int64Array(categoryids, "Id")

	if len(categoryintids) > 0 {

		var parentids []orm.Params
		categorytable := new(models.NsGoodsCategory)
		o.QueryTable(categorytable).Filter("id__in", categoryintids).Limit(10000).Values(&parentids, "parent_id")
		parentintids := utils.ExactMapValues2Int64Array(parentids, "ParentId")

		var parentcategories []orm.Params
		rs.OrderBy("sort_order").Filter("id__in", parentintids).Values(&parentcategories, "id", "name")

		for _, value := range parentcategories {
			id := value["id"].(int64)
			filterCategories = append(filterCategories, FilterCategory{Id: id, Name: value["name"].(string)})
		}
	}

	utils.ReturnHTTPSuccess(&this.Controller, filterCategories)
	this.ServeJSON()
}

func (this *GoodsController) Goods_New() {

	o := orm.NewOrm()
	var banners []models.NsPlatformAdv
	ad := new(models.NsPlatformAdv)
	o.QueryTable(ad).Filter("ap_id", 6667).All(&banners)
	utils.ReturnHTTPSuccess(&this.Controller, banners)
	this.ServeJSON()

}

func (this *GoodsController) Goods_Hot() {

	o := orm.NewOrm()
	var banners []models.NsPlatformAdv
	ad := new(models.NsPlatformAdv)
	o.QueryTable(ad).Filter("ap_id", 1165).All(&banners)
	utils.ReturnHTTPSuccess(&this.Controller, banners)
	this.ServeJSON()
}

func (this *GoodsController) Goods_Count() {

	o := orm.NewOrm()
	goodstable := new(models.NsGoods)

	count, _ := o.QueryTable(goodstable).Filter("state", 1).Count()

	utils.ReturnHTTPSuccess(&this.Controller, CountRtnJson{GoodsCount: count})
	this.ServeJSON()
}
