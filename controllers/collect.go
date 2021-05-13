package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/utils"
)

// CollectController CollectController
type CollectController struct {
	beego.Controller
}

// CollectListRtnJSON CollectListRtnJSON
type CollectListRtnJSON struct {
	models.NsMemberFavorites
	ListPicURL  string  `json:"list_pic_url"`
	GoodsBrief  string  `json:"goods_brief"`
	RetailPrice float64 `json:"retail_price"`
}

// CollectList CollectList
func (c *CollectController) CollectList() {

	typeID := c.GetString("typeID")

	qb, _ := orm.NewQueryBuilder("mysql")
	var list []CollectListRtnJSON

	qb.Select("nc.*", "ng.name", "ng.list_pic_url", "ng.goods_brief", "ng.retail_price").
		From("ns_member_favorites nc").
		LeftJoin("ns_goods ng").
		On("nc.value_id = ng.id").
		Where("gc.user_id =" + utils.Int2String(getLoginUserID()) + "and gc.type_id = " + typeID)

	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&list)

	utils.ReturnHTTPSuccess(&c.Controller, list)
	c.ServeJSON()

}

// AddOrDeleteRtnJSON AddOrDeleteRtnJSON
type AddOrDeleteRtnJSON struct {
	HandleType string
}

// CollectAddOrDelete CollectAddOrDelete
func (c *CollectController) CollectAddOrDelete() {
	//获取参数
	FavType := c.GetString("favType")
	valueID := c.GetString("valueID")

	intValueID := utils.String2Int(valueID)

	o := orm.NewOrm()
	collectTable := new(models.NsMemberFavorites)
	qs := o.QueryTable(collectTable)

	var collect models.NsMemberFavorites
	var rvJSON AddOrDeleteRtnJSON

	err := qs.Filter("fav_type", FavType).Filter("value_id", intValueID).Filter("user_id", getLoginUserID()).One(&collect)

	if err == orm.ErrNoRows {
		_, err = o.Insert(models.NsMemberFavorites{
			FavType: FavType,
			ValueID: intValueID,
			UserID:  getLoginUserID(),
			FavTime: utils.GetTimestamp(),
		})
		rvJSON = AddOrDeleteRtnJSON{HandleType: "add"}

	} else {
		_, err = qs.Filter("id", collect.ID).Delete()
		rvJSON = AddOrDeleteRtnJSON{HandleType: "delete"}
	}

	if err != nil {
		c.Abort(err.Error())
	}

	utils.ReturnHTTPSuccess(&c.Controller, rvJSON)

	c.ServeJSON()

}
