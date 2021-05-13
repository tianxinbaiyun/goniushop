package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/utils"
)

// CommentController CommentController
type CommentController struct {
	beego.Controller
}

// CommentPost CommentPost
func (c *CommentController) CommentPost() {
	//获取参数
	orderID := c.GetString("orderId")
	goodsID := c.GetString("goodsId")
	isAnonymous := c.GetString("isAnonymous")
	content := c.GetString("content")
	scores := c.GetString("scores")
	explainType := c.GetString("explainType")

	intOrderID := utils.String2Int(orderID)
	intGoodsID := utils.String2Int(goodsID)
	intIsAnonymous := utils.String2Int(isAnonymous)
	intScores := utils.String2Int(scores)
	intExplainType := utils.String2Int(explainType)
	//错误返回
	var err error
	defer func() {
		if err != nil {
			logs.Debug("this order id[%v] has a error ,err:%v", err)
		}
		c.ServeJSON()
		return
	}()

	o := orm.NewOrm()
	//查询订单
	orderTable := new(models.NsOrder)
	var order models.NsOrder
	err = o.QueryTable(orderTable).Filter("id", intOrderID).Filter("buyer_id", getLoginUserID()).One(&order)
	if err != nil {
		utils.ReturnHTTPError(&c.Controller, 400, "订单不存在")
		logs.Debug("get order data error,err:%v", err)
		return
	}
	//查询订单商品
	orderGoodsTable := new(models.NsOrderGoods)
	var orderGoods models.NsOrderGoods
	err = o.QueryTable(orderGoodsTable).Filter("goods_id", intGoodsID).Filter("orderID", intOrderID).One(&orderGoods)
	if err != nil {
		utils.ReturnHTTPError(&c.Controller, 400, "评论商品不存在")
		logs.Debug("get order goods data error,err:%v", err)
		return
	}
	//查询图片
	picture := new(models.SysAlbumPicture)
	var picturestable models.SysAlbumPicture
	err = o.QueryTable(picturestable).Filter("goods_id", intGoodsID).Filter("pic_id", orderGoods.GoodsPicture).One(&picture)
	if err != nil {
		utils.ReturnHTTPError(&c.Controller, 400, "商品图片不存在")
		logs.Debug("get picture data error,err:%v", err)
		return
	}
	//评论保存
	var comment models.NsGoodsEvaluate = models.NsGoodsEvaluate{
		OrderID:      order.ID,
		OrderNo:      order.OrderSn,
		OrderGoodsID: orderGoods.ID,
		GoodsID:      orderGoods.GoodsID,
		GoodsName:    orderGoods.GoodsName,
		GoodsPrice:   orderGoods.Price,
		GoodsImage:   picture.PicCover,
		Content:      utils.Base64Encode(content),
		MemberName:   order.UserName,
		UID:          getLoginUserID(),
		IsAnonymous:  intIsAnonymous,
		Scores:       intScores,
		ExplainType:  intExplainType,
		IsShow:       1,
		Addtime:      utils.GetTimestamp(),
	}

	_, err = o.Insert(&comment)
	if err != nil {
		c.Abort("添加评论成功")
	} else {
		c.Abort("评论保存失败")
	}

}

// CommentCountRtnJSON CommentCountRtnJSON
type CommentCountRtnJSON struct {
	AllCount    int64
	HasPicCount int
}

// CommentCount CommentCount
func (c *CommentController) CommentCount() {

	typeID := c.GetString("typeId")
	valueID := c.GetString("valueId")
	intTypeID := utils.String2Int(typeID)
	intValueID := utils.String2Int(valueID)

	o := orm.NewOrm()
	commentTable := new(models.NsGoodsEvaluate)
	allCount, _ := o.QueryTable(commentTable).Filter("type_id", intTypeID).Filter("value_id", intValueID).Count()

	qb, _ := orm.NewQueryBuilder("mysql")
	var list []models.NsGoodsEvaluate

	qb.Select("nc.*").
		From("ns_goods_evaluate nc").
		Where("nc.type_id =" + typeID + "and nc.value_id = " + valueID)

	sql := qb.String()
	o.Raw(sql).QueryRows(&list)
	hasPicCount := len(list)

	utils.ReturnHTTPSuccess(&c.Controller, CommentCountRtnJSON{allCount, hasPicCount})
	c.ServeJSON()
}

// GetCommentPageData It may need to be refactored.
func GetCommentPageData(rawData []models.NsGoodsEvaluate, page int, size int) utils.PageData {

	count := len(rawData)
	totalPages := (count + size - 1) / size
	var pageData []models.NsGoodsEvaluate

	for idx := (page - 1) * size; idx < page*size && idx < count; idx++ {
		pageData = append(pageData, rawData[idx])
	}

	return utils.PageData{NumsPerPage: size, CurrentPage: page, Count: count, TotalPages: totalPages, Data: pageData}
}

// CommentListtRtnJSON CommentListtRtnJSON
type CommentListtRtnJSON struct {
	Comment  string
	TypeID   int
	ValueID  int
	ID       int
	AddTime  string
	UserInfo orm.Params
	// PicList  []models.NsGoodsEvaluatePicture
}

// CommentList CommentList
func (c *CommentController) CommentList() {

	typeID := c.GetString("typeId")
	valueID := c.GetString("valueId")
	page := c.GetString("page")
	size := c.GetString("size")
	showType := c.GetString("showType")
	intTypeID := utils.String2Int(typeID)
	intValueID := utils.String2Int(valueID)

	intShowType := utils.String2Int(showType)

	var intSize = 10
	if size != "" {
		intSize = utils.String2Int(size)
	}

	var intPage = 1
	if page != "" {
		intPage = utils.String2Int(page)
	}

	o := orm.NewOrm()
	commentTable := new(models.NsGoodsEvaluate)
	var pageData utils.PageData
	var comments []models.NsGoodsEvaluate
	if intShowType != 1 {
		o.QueryTable(commentTable).Filter("type_id", intTypeID).Filter("value_id", intValueID).All(&comments)

	} else {
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("nc.*").
			From("ns_goods_evaluate nc").
			Where("c.type_id =" + typeID + "and c.value_id = " + valueID)

		sql := qb.String()
		o := orm.NewOrm()
		o.Raw(sql).QueryRows(&comments)
	}

	pageData = GetCommentPageData(comments, intPage, intSize)

	// var rtncomments []CommentListtRtnJSON
	// userTable := new(models.SysUser)

	// for _, val := range pageData.Data.([]models.NsGoodsEvaluate) {

	// 	var users []orm.Params
	// 	o.QueryTable(userTable).Filter("id", val.UserID).Values(&users, "user_name", "user_headimg", "nick_name")
	// 	rtncomments = append(rtncomments, CommentListtRtnJSON{
	// 		Comment:  val.Content,
	// 		TypeID:   val.TypeID,
	// 		ValueID:  val.ValueID,
	// 		ID:       val.ID,
	// 		AddTime:  utils.FormatTimestamp(val.AddTime, "2006-01-02 03:04:05 PM"),
	// 		UserInfo: users[0],
	// 		PicList:  commentpictures,
	// 	})

	// }
	pageData.Data = comments

	utils.ReturnHTTPSuccess(&c.Controller, pageData)
	c.ServeJSON()

}
