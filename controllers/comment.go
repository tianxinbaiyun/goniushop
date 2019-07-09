package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/utils"
)

type CommentController struct {
	beego.Controller
}

func (this *CommentController) Comment_Post() {
	//获取参数
	orderId := this.GetString("orderId")
	goodsId := this.GetString("goodsId")
	isAnonymous := this.GetString("isAnonymous")
	content := this.GetString("content")
	scores := this.GetString("scores")
	explainType := this.GetString("explainType")

	intorderId := utils.String2Int(orderId)
	intgoodsId := utils.String2Int(goodsId)
	intisAnonymous := utils.String2Int(isAnonymous)
	intscores := utils.String2Int(scores)
	intexplainType := utils.String2Int(explainType)
	//错误返回
	var err error
	defer func() {
		if err != nil {
			logs.Debug("this order id[%v] has a error ,err:%v", err)
		}
		this.ServeJSON()
		return
	}()

	o := orm.NewOrm()
	//查询订单
	ordertable := new(models.NsOrder)
	var order models.NsOrder
	err = o.QueryTable(ordertable).Filter("id", intorderId).Filter("buyer_id", getLoginUserId()).One(&order)
	if err != nil {
		utils.ReturnHTTPError(&this.Controller, 400, "订单不存在")
		logs.Debug("get order data error,err:%v", err)
		return
	}
	//查询订单商品
	ordergoodstable := new(models.NsOrderGoods)
	var ordergoods models.NsOrderGoods
	err = o.QueryTable(ordergoodstable).Filter("goods_id", intgoodsId).Filter("orderId", intorderId).One(&ordergoods)
	if err != nil {
		utils.ReturnHTTPError(&this.Controller, 400, "评论商品不存在")
		logs.Debug("get order goods data error,err:%v", err)
		return
	}
	//查询图片
	picture := new(models.SysAlbumPicture)
	var picturestable models.SysAlbumPicture
	err = o.QueryTable(picturestable).Filter("goods_id", intgoodsId).Filter("pic_id", ordergoods.GoodsPicture).One(&picture)
	if err != nil {
		utils.ReturnHTTPError(&this.Controller, 400, "商品图片不存在")
		logs.Debug("get picture data error,err:%v", err)
		return
	}
	//评论保存
	var comment models.NsGoodsEvaluate = models.NsGoodsEvaluate{
		OrderId:      order.Id,
		OrderNo:      order.OrderSn,
		OrderGoodsId: ordergoods.Id,
		GoodsId:      ordergoods.GoodsId,
		GoodsName:    ordergoods.GoodsName,
		GoodsPrice:   ordergoods.Price,
		GoodsImage:   picture.PicCover,
		Content:      utils.Base64Encode(content),
		MemberName:   order.UserName,
		Uid:          getLoginUserId(),
		IsAnonymous:  intisAnonymous,
		Scores:       intscores,
		ExplainType:  intexplainType,
		IsShow:       1,
		Addtime:      utils.GetTimestamp(),
	}

	_, err = o.Insert(&comment)
	if err != nil {
		this.Abort("添加评论成功")
	} else {
		this.Abort("评论保存失败")
	}

}

type CommentCountRtnJson struct {
	AllCount    int64
	HasPicCount int
}

func (this *CommentController) Comment_Count() {

	typeId := this.GetString("typeId")
	valueId := this.GetString("valueId")
	inttypeId := utils.String2Int(typeId)
	intvalueId := utils.String2Int(valueId)

	o := orm.NewOrm()
	commenttable := new(models.NsGoodsEvaluate)
	allcount, _ := o.QueryTable(commenttable).Filter("type_id", inttypeId).Filter("value_id", intvalueId).Count()

	qb, _ := orm.NewQueryBuilder("mysql")
	var list []models.NsGoodsEvaluate

	qb.Select("nc.*").
		From("ns_goods_evaluate nc").
		Where("nc.type_id =" + typeId + "and nc.value_id = " + valueId)

	sql := qb.String()
	o.Raw(sql).QueryRows(&list)
	haspiccount := len(list)

	utils.ReturnHTTPSuccess(&this.Controller, CommentCountRtnJson{allcount, haspiccount})
	this.ServeJSON()
}

//It may need to be refactored.
func GetCommentPageData(rawData []models.NsGoodsEvaluate, page int, size int) utils.PageData {

	count := len(rawData)
	totalpages := (count + size - 1) / size
	var pagedata []models.NsGoodsEvaluate

	for idx := (page - 1) * size; idx < page*size && idx < count; idx++ {
		pagedata = append(pagedata, rawData[idx])
	}

	return utils.PageData{NumsPerPage: size, CurrentPage: page, Count: count, TotalPages: totalpages, Data: pagedata}
}

type CommenListtRtnJson struct {
	Comment  string
	TypeId   int
	ValueId  int
	Id       int
	AddTime  string
	UserInfo orm.Params
	// PicList  []models.NsGoodsEvaluatePicture
}

func (this *CommentController) Comment_List() {

	typeId := this.GetString("typeId")
	valueId := this.GetString("valueId")
	page := this.GetString("page")
	size := this.GetString("size")
	showType := this.GetString("showType")
	inttypeId := utils.String2Int(typeId)
	intvalueId := utils.String2Int(valueId)

	intshowtype := utils.String2Int(showType)

	var intsize int = 10
	if size != "" {
		intsize = utils.String2Int(size)
	}

	var intpage int = 1
	if page != "" {
		intpage = utils.String2Int(page)
	}

	o := orm.NewOrm()
	commenttable := new(models.NsGoodsEvaluate)
	var pagedata utils.PageData
	var comments []models.NsGoodsEvaluate
	if intshowtype != 1 {
		o.QueryTable(commenttable).Filter("type_id", inttypeId).Filter("value_id", intvalueId).All(&comments)

	} else {
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("nc.*").
			From("ns_goods_evaluate nc").
			Where("c.type_id =" + typeId + "and c.value_id = " + valueId)

		sql := qb.String()
		o := orm.NewOrm()
		o.Raw(sql).QueryRows(&comments)
	}

	pagedata = GetCommentPageData(comments, intpage, intsize)

	// var rtncomments []CommenListtRtnJson
	// usertable := new(models.SysUser)

	// for _, val := range pagedata.Data.([]models.NsGoodsEvaluate) {

	// 	var users []orm.Params
	// 	o.QueryTable(usertable).Filter("id", val.UserId).Values(&users, "user_name", "user_headimg", "nick_name")
	// 	rtncomments = append(rtncomments, CommenListtRtnJson{
	// 		Comment:  val.Content,
	// 		TypeId:   val.TypeId,
	// 		ValueId:  val.ValueId,
	// 		Id:       val.Id,
	// 		AddTime:  utils.FormatTimestamp(val.AddTime, "2006-01-02 03:04:05 PM"),
	// 		UserInfo: users[0],
	// 		PicList:  commentpictures,
	// 	})

	// }
	pagedata.Data = comments

	utils.ReturnHTTPSuccess(&this.Controller, pagedata)
	this.ServeJSON()

}
