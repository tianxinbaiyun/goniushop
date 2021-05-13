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

// GoodsController GoodsController
type GoodsController struct {
	beego.Controller
}

// SkuRtnJSON SkuRtnJSON
type SkuRtnJSON struct {
	ProductList []models.NsGoodsSku `json:"productList"`
}

// DetailRtnJSON DetailRtnJSON
type DetailRtnJSON struct {
	SkuRtnJSON
	Goods     models.NsGoods           `json:"info"`
	Galleries []models.SysAlbumPicture `json:"gallery"`
	Attribute []Attribute              `json:"attribute"`
	// Issues         []models.NsGoodsIssue `json:"issue"`
	UserHasCollect int                 `json:"userHasCollect"`
	Comment        Comment             `json:"comment"`
	Brand          models.NsGoodsBrand `json:"brand"`
}

// CategoryRtnJSON CategoryRtnJSON
type CategoryRtnJSON struct {
	CurCategory     models.NsGoodsCategory   `json:"currentCategory"`
	ParentCategory  models.NsGoodsCategory   `json:"parentCategory"`
	BrotherCategory []models.NsGoodsCategory `json:"brotherCategory"`
}

// Attribute Attribute
type Attribute struct {
	Value         string `json:"value"`
	AttrValueName string `json:"attr_value_name"`
}

// CommentUser CommentUser
type CommentUser struct {
	NickName string `json:"nick_name"`
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
}

// Comment Comment
type Comment struct {
	Count int64                  `json:"count"`
	Data  models.NsGoodsEvaluate `json:"data"`
}

// FilterCategory FilterCategory
type FilterCategory struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Checked bool   `json:"checked"`
}

// ListRtnJSON ListRtnJSON
type ListRtnJSON struct {
	utils.PageData
	FilterCategories []FilterCategory `json:"filterCategory"`
	GoodsList        []orm.Params     `json:"goodsList"`
}

// Banner Banner
type Banner struct {
	URL    string `json:"url"`
	Name   string `json:"name"`
	ImgURL string `json:"imgurl"`
}

// NewRtnJSON NewRtnJSON
type NewRtnJSON struct {
	BannerInfo Banner `json:"bannerinfo"`
}

// HotRtnJSON HotRtnJSON
type HotRtnJSON struct {
	BannerInfo Banner `json:"bannerinfo"`
}

// CountRtnJSON CountRtnJSON
type CountRtnJSON struct {
	GoodsCount int64 `json:"goodsCount"`
}

// updateJSONKeysGoods updateJSONKeysGoods
func updateJSONKeysGoods(values []orm.Params) {

	for _, val := range values {
		for k, v := range val {
			switch k {
			case "ID":
				delete(val, k)
				val["id"] = v
			case "Name":
				delete(val, k)
				val["name"] = v
			case "ListPicURL":
				delete(val, k)
				val["list_pic_url"] = v
			case "RetailPrice":
				delete(val, k)
				val["retail_price"] = v
			}
		}
	}
}

// GoodsIndex GoodsIndex
func (c *GoodsController) GoodsIndex() {
	o := orm.NewOrm()

	var goods []models.NsGoods
	good := new(models.NsGoods)
	_, _ = o.QueryTable(good).All(&goods)

	utils.ReturnHTTPSuccess(&c.Controller, goods)
	c.ServeJSON()
}

// GoodsSku GoodsSku
func (c *GoodsController) GoodsSku() {

	goodsID := c.GetString("id")
	goodsIDInt := utils.String2Int(goodsID)

	// plist := models.GetProductList(goodsIDInt)
	plist := models.GetGoodSku(goodsIDInt)

	utils.ReturnHTTPSuccess(&c.Controller, SkuRtnJSON{ProductList: plist})
	c.ServeJSON()
}

// GoodsDetail GoodsDetail
func (c *GoodsController) GoodsDetail() {
	goodsID := c.GetString("id")
	intGoodsID := utils.String2Int(goodsID)

	o := orm.NewOrm()

	var goodOne models.NsGoods
	good := new(models.NsGoods)
	_ = o.QueryTable(good).Filter("goods_id", intGoodsID).One(&goodOne)

	var galleries []models.SysAlbumPicture
	gallery := new(models.SysAlbumPicture)
	imgIDs := strings.Split(goodOne.ImgIDArray, ",")
	_, _ = o.QueryTable(gallery).Filter("pic_id__in", imgIDs).Limit(4).All(&galleries)

	qb, _ := orm.NewQueryBuilder("mysql")
	var attributes []Attribute
	qb.Select("a.value", "a.attr_value_name").
		From(" ns_attribute_value a").
		InnerJoin("ns_goods_attribute ga").On("ga.attr_value_id = a.attr_value_id").
		Where("ga.goods_id =" + goodsID).GroupBy("a.attr_value_id").Asc()
	sql := qb.String()
	_, _ = o.Raw(sql).QueryRows(&attributes)

	var brandOne models.NsGoodsBrand
	brand := new(models.NsGoodsBrand)
	_ = o.QueryTable(brand).Filter("brand_id", goodOne.BrandID).One(&brandOne)

	comment := new(models.NsGoodsEvaluate)
	commentCount, _ := o.QueryTable(comment).Filter("goods_id", intGoodsID).Filter("is_show", 1).Count()
	var hotCommentOne models.NsGoodsEvaluate
	_ = o.QueryTable(comment).Filter("goods_id", intGoodsID).Filter("is_show", 1).One(&hotCommentOne)
	commentVal := Comment{Count: commentCount, Data: hotCommentOne}
	loginUserID := getLoginUserID()

	userHasCollect := models.IsUserHasCollect(loginUserID, `goods`, intGoodsID)

	// models.AddFootprint(loginUserID, intGoodsID)

	plist := models.GetGoodSku(intGoodsID)

	utils.ReturnHTTPSuccess(&c.Controller, DetailRtnJSON{Goods: goodOne, Galleries: galleries, Attribute: attributes,
		UserHasCollect: userHasCollect, Comment: commentVal, Brand: *brand,
		SkuRtnJSON: SkuRtnJSON{ProductList: plist}})
	c.ServeJSON()
}

// GoodsCategory GoodsCategory
func (c *GoodsController) GoodsCategory() {

	goodsID := c.GetString("id")
	intGoodsID := utils.String2Int(goodsID)

	o := orm.NewOrm()
	var curCategory models.NsGoodsCategory
	var parentCategory models.NsGoodsCategory
	var brotherCategory []models.NsGoodsCategory

	category := new(models.NsGoodsCategory)

	_ = o.QueryTable(category).Filter("id", intGoodsID).One(&curCategory)
	_ = o.QueryTable(category).Filter("id", curCategory.ParentID).One(&parentCategory)
	_, _ = o.QueryTable(category).Filter("pid", curCategory.ParentID).All(&brotherCategory)

	utils.ReturnHTTPSuccess(&c.Controller, CategoryRtnJSON{CurCategory: curCategory,
		ParentCategory: parentCategory, BrotherCategory: brotherCategory})
	c.ServeJSON()
}

// GoodsList GoodsList
func (c *GoodsController) GoodsList() {
	categoryID := c.GetString("categoryID")
	brandID := c.GetString("brandId")
	keyword := c.GetString("keyword")
	isNew := c.GetString("isNew")
	isHot := c.GetString("isHot")
	page := c.GetString("page")
	size := c.GetString("size")
	sort := c.GetString("sort")
	order := c.GetString("order")

	var intSize = 10
	if size != "" {
		intSize = utils.String2Int(size)
	}

	var intPage = 1
	if page != "" {
		intPage = utils.String2Int(page)
	}

	o := orm.NewOrm()
	var categoryIDs []orm.Params
	var list []orm.Params

	//构建查询语句
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("g.goods_id", "g.goods_name", "g.brand_id", "g.category_id", "g.price", "p.pic_cover", "p.pic_cover_mid", "p.pic_cover_big").
		From("ns_goods g").
		LeftJoin("sys_album_picture p").
		On("g.picture = p.pic_id").
		Where("g.state=1")
	goodsTable := new(models.NsGoods)
	rs := o.QueryTable(goodsTable)

	//参数处理
	if isNew != "" {
		qb.And(fmt.Sprintf("is_new=%s", isNew))
		rs = rs.Filter("is_new", isNew)
	}
	if isHot != "" {
		qb.And(fmt.Sprintf("is_hot=%s", isHot))
		rs = rs.Filter("is_hot", isHot)
	}
	keyword, _ = url.QueryUnescape(keyword)
	if keyword != "" {
		qb.And(fmt.Sprintf("goods_name like '{%s}'", "%"+keyword+"%"))
		rs = rs.Filter("goods_name__icontains", keyword)
	}
	if brandID != "" {
		qb.And(fmt.Sprintf("brand_id = %v", brandID))
		rs = rs.Filter("brand_id", brandID)
	}

	//分类处理

	_, _ = rs.Limit(10000).Values(&categoryIDs, "category_id")
	categoryIntIDs := utils.ExactMapValues2Int64Array(categoryIDs, "CategoryID")
	var filterCategories = []FilterCategory{{ID: 0, Name: "全部", Checked: false}}
	if len(categoryIntIDs) > 0 {
		var parentIDs []orm.Params
		categoryTable := new(models.NsGoodsCategory)
		_, _ = o.QueryTable(categoryTable).Filter("category_id__in", categoryIntIDs).Limit(10000).Values(&parentIDs, "pid")
		parentIntIDs := utils.ExactMapValues2Int64Array(parentIDs, "ParentID")

		var parentCategories []orm.Params
		_, _ = o.QueryTable(categoryTable).Filter("category_id__in", parentIntIDs).OrderBy("sort").Values(&parentCategories, "category_id", "category_name")

		for _, value := range parentCategories {
			id := value["Id"].(int64)
			checked := categoryID == "" && id == 0

			filterCategories = append(filterCategories, FilterCategory{ID: id, Name: value["Name"].(string), Checked: checked})
		}
	}

	//商品查询
	if categoryID != "" {
		intCategoryID := utils.String2Int(categoryID)
		if intCategoryID >= 0 {
			qb.And(fmt.Sprintf("category_id_1 = %v or category_id_2=%v or category_id_3 = %v", intCategoryID, intCategoryID, intCategoryID))
		}

	}
	if sort == "price" {
		orderStr := "price"
		if order == "desc" {
			orderStr = "-" + orderStr
		}
		qb.OrderBy(orderStr)
	} else {
		qb.OrderBy("goods_id")
	}
	sql := qb.String()
	_, _ = o.Raw(sql).Values(&list)
	// logs.Debug("list success,sql:%v,list:%v", sql, list)

	pageData := utils.GetPageData(list, intPage, intSize)
	fmt.Println(pageData)

	utils.ReturnHTTPSuccess(&c.Controller, ListRtnJSON{PageData: pageData, FilterCategories: filterCategories, GoodsList: pageData.Data.([]orm.Params)})
	c.ServeJSON()
}

// GoodsFilter GoodsFilter
func (c *GoodsController) GoodsFilter() {

	categoryID := c.GetString("categoryID")
	keyword := c.GetString("keyword")
	isNew := c.GetString("isNew")
	isHot := c.GetString("isHot")

	o := orm.NewOrm()
	goodsTable := new(models.NsGoods)
	rs := o.QueryTable(goodsTable)

	if categoryID != "" {
		intCategoryID := utils.String2Int(categoryID)
		rs = rs.Filter("category_id__in", models.GetChildCategoryID(intCategoryID))
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

	var filterCategories = []FilterCategory{{ID: 0, Name: "全部"}}

	var categoryIDs []orm.Params
	_, _ = rs.Limit(10000).Values(&categoryIDs, "category_id")
	categoryIntIDs := utils.ExactMapValues2Int64Array(categoryIDs, "Id")

	if len(categoryIntIDs) > 0 {

		var parentIDs []orm.Params
		categoryTable := new(models.NsGoodsCategory)
		_, _ = o.QueryTable(categoryTable).Filter("id__in", categoryIntIDs).Limit(10000).Values(&parentIDs, "parent_id")
		parentIntIDs := utils.ExactMapValues2Int64Array(parentIDs, "ParentID")

		var parentCategories []orm.Params
		_, _ = rs.OrderBy("sort_order").Filter("id__in", parentIntIDs).Values(&parentCategories, "id", "name")

		for _, value := range parentCategories {
			id := value["id"].(int64)
			filterCategories = append(filterCategories, FilterCategory{ID: id, Name: value["name"].(string)})
		}
	}

	utils.ReturnHTTPSuccess(&c.Controller, filterCategories)
	c.ServeJSON()
}

// GoodsNew GoodsNew
func (c *GoodsController) GoodsNew() {

	o := orm.NewOrm()
	var banners []models.NsPlatformAdv
	ad := new(models.NsPlatformAdv)
	_, _ = o.QueryTable(ad).Filter("ap_id", 6667).All(&banners)
	utils.ReturnHTTPSuccess(&c.Controller, banners)
	c.ServeJSON()

}

// GoodsHot GoodsHot
func (c *GoodsController) GoodsHot() {

	o := orm.NewOrm()
	var banners []models.NsPlatformAdv
	ad := new(models.NsPlatformAdv)
	_, _ = o.QueryTable(ad).Filter("ap_id", 1165).All(&banners)
	utils.ReturnHTTPSuccess(&c.Controller, banners)
	c.ServeJSON()
}

// GoodsCount GoodsCount
func (c *GoodsController) GoodsCount() {

	o := orm.NewOrm()
	goodsTable := new(models.NsGoods)

	count, _ := o.QueryTable(goodsTable).Filter("state", 1).Count()

	utils.ReturnHTTPSuccess(&c.Controller, CountRtnJSON{GoodsCount: count})
	c.ServeJSON()
}
