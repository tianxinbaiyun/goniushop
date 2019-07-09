package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/utils"
)

type CollectController struct {
	beego.Controller
}

type CollectListRtnJson struct {
	models.NsMemberFavorites
	ListPicUrl  string  `json:"list_pic_url"`
	GoodsBrief  string  `json:"goods_brief"`
	RetailPrice float64 `json:"retail_price"`
}

func (this *CollectController) Collect_List() {

	typeId := this.GetString("typeId")

	qb, _ := orm.NewQueryBuilder("mysql")
	var list []CollectListRtnJson

	qb.Select("nc.*", "ng.name", "ng.list_pic_url", "ng.goods_brief", "ng.retail_price").
		From("ns_member_favorites nc").
		LeftJoin("ns_goods ng").
		On("nc.value_id = ng.id").
		Where("gc.user_id =" + utils.Int2String(getLoginUserId()) + "and gc.type_id = " + typeId)

	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&list)

	utils.ReturnHTTPSuccess(&this.Controller, list)
	this.ServeJSON()

}

type AddorDeleteRtnJson struct {
	HandleType string
}

func (this *CollectController) Collect_AddorDelete() {
	//获取参数
	FavType := this.GetString("favType")
	valueId := this.GetString("valueId")

	intvalueId := utils.String2Int(valueId)

	o := orm.NewOrm()
	collecttable := new(models.NsMemberFavorites)
	qs := o.QueryTable(collecttable)

	var collect models.NsMemberFavorites
	var rvjson AddorDeleteRtnJson

	err := qs.Filter("fav_type", FavType).Filter("value_id", intvalueId).Filter("user_id", getLoginUserId()).One(&collect)

	if err == orm.ErrNoRows {
		_, err = o.Insert(models.NsMemberFavorites{
			FavType: FavType,
			ValueId: intvalueId,
			UserId:  getLoginUserId(),
			FavTime: utils.GetTimestamp(),
		})
		rvjson = AddorDeleteRtnJson{HandleType: "add"}

	} else {
		_, err = qs.Filter("id", collect.Id).Delete()
		rvjson = AddorDeleteRtnJson{HandleType: "delete"}
	}

	if err != nil {
		this.Abort(err.Error())
	}

	utils.ReturnHTTPSuccess(&this.Controller, rvjson)

	this.ServeJSON()

}
