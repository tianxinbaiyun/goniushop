package controllers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/tianxinbaiyun/goniushop/models"
	"github.com/tianxinbaiyun/goniushop/utils"
)

type CartController struct {
	beego.Controller
}

type CartTotal struct {
	GoodsCount         int     `json:"goodsCount"`
	GoodsAmount        float64 `json:"goodsAmount"`
	CheckedGoodsCount  int     `json:"checkedGoodsCount"`
	CheckedGoodsAmount float64 `json:"checkedGoodsAmount"`
}
type GoodsCount struct {
	CartTotal CartTotal `json:"cartTotal"`
}
type IndexCartData struct {
	CartList  []newCart `json:"cartList"`
	CartTotal CartTotal `json:"cartTotal"`
}
type newCart struct {
	models.NsCart
	PicCoverBig string `json:"pic_cover_big,omitempty"`
	PicCoverMid string `json:"pic_cover_mid,omitempty"`
	PicCover    string `json:"pic_cover,omitempty"`
	SkuName     string `json:"sku_name,omitempty"`
}

func getCart() IndexCartData {
	intuserId := getLoginUserId()
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	var carts []newCart
	qb.Select("c.*", "p.pic_cover", "p.pic_cover_mid", "p.pic_cover_big", "s.sku_name").
		From(" ns_cart c").
		LeftJoin("sys_album_picture p").
		On("c.goods_picture = p.pic_id").
		LeftJoin("ns_goods_sku s").
		On("c.sku_id = s.sku_id").
		Where(fmt.Sprintf("c.buyer_id = %d", intuserId))

	sql := qb.String()
	num, err := o.Raw(sql).QueryRows(&carts)
	if err != nil {
		logs.Debug("get carts erroor,err:%v,num:%v", err, num)
	}
	var goodsCount int
	var goodsAmount float64
	var checkedGoodsCount int
	var checkedGoodsAmount float64

	for _, val := range carts {
		goodsCount += val.Number
		goodsAmount += float64(val.Number) * val.Price
		checkedGoodsCount += val.Number
		checkedGoodsAmount += float64(val.Number) * val.Price
	}

	return IndexCartData{carts, CartTotal{goodsCount, goodsAmount, checkedGoodsCount, checkedGoodsAmount}}
}

func (this *CartController) Cart_Index() {

	utils.ReturnHTTPSuccess(&this.Controller, getCart())
	this.ServeJSON()
}

type CartAddBody struct {
	GoodsId int    `json:"goodsId"`
	SpecIds string `json:"specIds"`
	Number  int    `json:"number"`
}

func (this *CartController) Cart_Add() {
	var (
		err error
		ab  CartAddBody
	)
	defer func() {
		if err != nil {
			logs.Debug("this has a err ,err:%v", err)
		}
		this.ServeJSON()
		return
	}()
	//输入检验
	body := this.Ctx.Input.RequestBody
	err = json.Unmarshal(body, &ab)
	if err != nil {
		logs.Debug("Unmarshal RequestBody ,err:%v", err)
		utils.ReturnHTTPError(&this.Controller, 400, "请求参数错误")
		return
	}
	intgoodsId := ab.GoodsId
	specIds := ab.SpecIds
	intnumber := ab.Number
	intuserId := getLoginUserId()
	o := orm.NewOrm()
	//查询商品
	goodstable := new(models.NsGoods)
	var goods models.NsGoods
	err = o.QueryTable(goodstable).Filter("goods_id", intgoodsId).One(&goods)
	if err == orm.ErrNoRows {
		logs.Debug("goods losed ,err:%v", err)
		utils.ReturnHTTPError(&this.Controller, 400, "商品已下架")
		return
	}
	//查询库存
	skutable := new(models.NsGoodsSku)
	var sku models.NsGoodsSku
	err = o.QueryTable(skutable).Filter("goods_id", intgoodsId).Filter("attr_value_items", specIds).One(&sku)
	if err == orm.ErrNoRows || sku.Stock < intnumber {
		utils.ReturnHTTPError(&this.Controller, 400, "库存不足")
		return
	}
	//开始事物
	o.Begin()
	defer func() {
		if err != nil {
			o.Rollback()
			logs.Debug("this has a err ,err:%v", err)
			utils.ReturnHTTPError(&this.Controller, 400, "操作撤销")
			return
		} else {
			err = o.Commit()
			utils.ReturnHTTPSuccess(&this.Controller, getCart())
			return
		}
	}()

	carttable := new(models.NsCart)
	var cart models.NsCart
	err = o.QueryTable(carttable).Filter("goods_id", intgoodsId).Filter("sku_id", sku.Id).
		Filter("buyer_id", intuserId).One(&cart)
	if err == orm.ErrNoRows {
		cartData := models.NsCart{
			UserId:       intuserId,
			GoodsId:      intgoodsId,
			GoodsName:    goods.GoodsName,
			SkuId:        sku.Id,
			SkuName:      sku.SkuName,
			Price:        sku.Price,
			Number:       intnumber,
			GoodsPicture: goods.Picture,
		}
		_, err = o.Insert(&cartData)
	} else {
		if sku.Stock < (intnumber + cart.Number) {
			utils.ReturnHTTPError(&this.Controller, 400, "库存不足")
			return
		}
		_, err = o.QueryTable(carttable).Filter("cart_id", cart.Id).Update(orm.Params{"number": orm.ColValue(orm.ColAdd, intnumber)})
	}
	return
}

type CartUpdateBody struct {
	GoodsId int `json:"goodsId"`
	SkuId   int `json:"skuId"`
	Number  int `json:"number"`
	Id      int `json:"id"`
}

func (this *CartController) Cart_Update() {

	var ub CartUpdateBody
	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &ub)
	logs.Debug(ub)
	if err != nil {
		logs.Debug("Unmarshal RequestBody ,err:%v", err)
	}

	intgoodsId := ub.GoodsId
	skuId := ub.SkuId
	intnumber := ub.Number
	intid := ub.Id
	o := orm.NewOrm()
	goodstable := new(models.NsGoods)
	var goods models.NsGoods
	err = o.QueryTable(goodstable).Filter("goods_id", intgoodsId).One(&goods)
	if err == orm.ErrNoRows {
		utils.ReturnHTTPError(&this.Controller, 400, "商品已下架")
		this.ServeJSON()
		return
	}

	skutable := new(models.NsGoodsSku)
	var sku models.NsGoodsSku
	err = o.QueryTable(skutable).Filter("goods_id", intgoodsId).Filter("sku_id", skuId).One(&sku)
	if err == orm.ErrNoRows || sku.Stock < intnumber {
		utils.ReturnHTTPError(&this.Controller, 400, "库存不足")
		this.ServeJSON()
		return
	}
	defer func() {
		if err != nil {
			o.Rollback()
			logs.Debug("this has a err ,err:%v", err)
			utils.ReturnHTTPError(&this.Controller, 400, "操作撤销")
			this.ServeJSON()
			return
		} else {
			err = o.Commit()
			utils.ReturnHTTPSuccess(&this.Controller, getCart())
			this.ServeJSON()
		}
	}()

	carttable := new(models.NsCart)
	var cart models.NsCart
	err = o.QueryTable(carttable).Filter("id", intid).One(&cart)
	if err != nil {
		logs.Info("select end,cart:%v,err:%v", cart, err)
		return
	}
	if sku.Stock < intnumber {
		utils.ReturnHTTPError(&this.Controller, 400, "库存不足")
		this.ServeJSON()
		return
	}
	//开始事物
	err = o.Begin()
	_, err = o.QueryTable(carttable).Filter("id", intid).Update(orm.Params{
		"num":    intnumber,
		"price":  sku.Price,
		"sku_id": sku.Id,
	})

}

type CartCheckedBody struct {
	IsChecked int         `json:"isChecked"`
	Ids       interface{} `json:"ids"`
}

func (this *CartController) Cart_Checked() {

	var cb CartCheckedBody
	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &cb)
	if err != nil {
		logs.Debug("json Unmarshal error，err:%v", err)
	}
	intisChecked := cb.IsChecked

	if cb.Ids == "" {
		this.Abort("删除出错")
	}
	var IdsArray []string
	switch val := cb.Ids.(type) {
	// 单选
	case float64:
		IdsArray = append(IdsArray, utils.Int2String(int(val)))
	//多选
	case string:
		IdsArray = strings.Split(val, ",")
	default:

	}

	o := orm.NewOrm()
	carttable := new(models.NsCart)
	o.QueryTable(carttable).Filter("cart_id__in", IdsArray).Update(orm.Params{
		"checked": intisChecked,
	})

	utils.ReturnHTTPSuccess(&this.Controller, getCart())
	this.ServeJSON()
}

type CartDeleteBody struct {
	Ids string `json:"ids"`
}

func (this *CartController) Cart_Delete() {

	var input CartDeleteBody
	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &input)
	if err != nil {
		logs.Debug("json Unmarshal err，err:%v", err)
	}
	intuserId := getLoginUserId()

	idsArray := strings.Split(input.Ids, ",")

	o := orm.NewOrm()
	o.Begin()
	carttable := new(models.NsCart)
	_, err = o.QueryTable(carttable).Filter("cart_id__in", idsArray).Filter("buyer_id", intuserId).Delete()
	defer func() {
		if err != nil {
			o.Rollback()
			logs.Debug("this has a err ,err:%v", err)
			utils.ReturnHTTPError(&this.Controller, 400, "删除失败")
			this.ServeJSON()
			return
		} else {
			err = o.Commit()
			utils.ReturnHTTPSuccess(&this.Controller, getCart())
			this.ServeJSON()
		}
	}()
	return
}

func (this *CartController) Cart_GoodsCount() {

	cartData := getCart()
	goodscount := GoodsCount{CartTotal: CartTotal{GoodsCount: cartData.CartTotal.GoodsCount}}
	utils.ReturnHTTPSuccess(&this.Controller, goodscount)
	this.ServeJSON()
}

type CartAddress struct {
	models.NsMemberExpressAddress
	ProvinceName string `json:"province_name"`
	CityName     string `json:"city_name"`
	DistrictName string `json:"district_name"`
	FullRegion   string `json:"full_region"`
}

type CheckoutRtnJson struct {
	CheckedAddress   CartAddress `json:"checkedAddress"`
	FreightPrice     float64     `json:"freightPrice"`
	CouponPrice      float64     `json:"couponPrice"`
	CheckedGoodsList []newCart   `json:"checkedGoodsList"`
	GoodsTotalPrice  float64     `json:"goodsTotalPrice"`
	OrderTotalPrice  float64     `json:"orderTotalPrice"`
	ActualPrice      float64     `json:"actualPrice"`
}
type CheckoutInput struct {
	AddressId int    `json:"address_id"`
	CouponId  string `json:"coupon_id"`
}

func (this *CartController) Cart_Checkout() {
	var (
		goodstotalprice float64
		ordertotalprice float64
		actualPrice     float64
		freightPrice    float64
	)
	ids := this.GetString("ids")
	// addressId := this.GetString("intaddressid")
	if ids == "" {
		utils.ReturnHTTPError(&this.Controller, 400, "商品未找到")
		this.ServeJSON()
		return
	}
	o := orm.NewOrm()
	//获取用户地址

	// intaddressid := utils.String2Int(addressId)
	addresstable := new(models.NsMemberExpressAddress)
	var myaddress models.NsMemberExpressAddress
	err := o.QueryTable(addresstable).Filter("uid", getLoginUserId()).Filter("is_default", 1).One(&myaddress)
	var customaddress CartAddress
	if err != orm.ErrNoRows {
		customaddress.NsMemberExpressAddress = myaddress
		customaddress.ProvinceName = models.GetProvinceName(myaddress.ProvinceId)
		customaddress.CityName = models.GetCityName(myaddress.CityId)
		customaddress.DistrictName = models.GetDistrictName(myaddress.DistrictId)
		customaddress.FullRegion = customaddress.ProvinceName + customaddress.CityName + customaddress.DistrictName
	}
	//获取购物车
	intuserId := getLoginUserId()

	qb, _ := orm.NewQueryBuilder("mysql")
	var carts []newCart
	qb.Select("c.*", "p.pic_cover", "p.pic_cover_mid", "p.pic_cover_big", "s.sku_name").
		From(" ns_cart c").
		LeftJoin("sys_album_picture p").
		On("c.goods_picture = p.pic_id").
		LeftJoin("ns_goods_sku s").
		On("c.sku_id = s.sku_id").
		Where(fmt.Sprintf("c.buyer_id = %d", intuserId)).
		And(fmt.Sprintf("c.cart_id IN (%v)", ids))
	sql := qb.String()
	num, err := o.Raw(sql).QueryRows(&carts)
	if err != nil {
		logs.Debug("get cart error，num:%v,err:%v", num, err)
	}
	// logs.Info(carts)
	for _, v := range carts {
		goodstotalprice += v.Price * float64(v.Number)
		ordertotalprice += v.Price * float64(v.Number)
		actualPrice += v.Price * float64(v.Number)
	}
	utils.ReturnHTTPSuccess(&this.Controller, CheckoutRtnJson{
		CheckedAddress:   customaddress,
		CheckedGoodsList: carts,
		GoodsTotalPrice:  goodstotalprice,
		OrderTotalPrice:  ordertotalprice,
		ActualPrice:      actualPrice,
		FreightPrice:     freightPrice,
	})
	this.ServeJSON()
}
