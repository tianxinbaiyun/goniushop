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

// CartController CartController
type CartController struct {
	beego.Controller
}

// CartTotal CartTotal
type CartTotal struct {
	GoodsCount         int     `json:"goodsCount"`
	GoodsAmount        float64 `json:"goodsAmount"`
	CheckedGoodsCount  int     `json:"checkedGoodsCount"`
	CheckedGoodsAmount float64 `json:"checkedGoodsAmount"`
}

// GoodsCount GoodsCount
type GoodsCount struct {
	CartTotal CartTotal `json:"cartTotal"`
}

// IndexCartData IndexCartData
type IndexCartData struct {
	CartList  []newCart `json:"cartList"`
	CartTotal CartTotal `json:"cartTotal"`
}

// newCart newCart
type newCart struct {
	models.NsCart
	PicCoverBig string `json:"pic_cover_big,omitempty"`
	PicCoverMid string `json:"pic_cover_mid,omitempty"`
	PicCover    string `json:"pic_cover,omitempty"`
	SkuName     string `json:"sku_name,omitempty"`
}

// getCart getCart
func getCart() IndexCartData {
	intUserID := getLoginUserID()
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	var carts []newCart
	qb.Select("c.*", "p.pic_cover", "p.pic_cover_mid", "p.pic_cover_big", "s.sku_name").
		From(" ns_cart c").
		LeftJoin("sys_album_picture p").
		On("c.goods_picture = p.pic_id").
		LeftJoin("ns_goods_sku s").
		On("c.sku_id = s.sku_id").
		Where(fmt.Sprintf("c.buyer_id = %d", intUserID))

	sql := qb.String()
	num, err := o.Raw(sql).QueryRows(&carts)
	if err != nil {
		logs.Debug("get carts error,err:%v,num:%v", err, num)
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

// CartIndex CartIndex
func (c *CartController) CartIndex() {

	utils.ReturnHTTPSuccess(&c.Controller, getCart())
	c.ServeJSON()
}

// CartAddBody CartAddBody
type CartAddBody struct {
	GoodsID int    `json:"goodsId"`
	SpecIDs string `json:"specIds"`
	Number  int    `json:"number"`
}

// CartAdd CartAdd
func (c *CartController) CartAdd() {
	var (
		err error
		ab  CartAddBody
	)
	defer func() {
		if err != nil {
			logs.Debug("this has a err ,err:%v", err)
		}
		c.ServeJSON()
		return
	}()
	//输入检验
	body := c.Ctx.Input.RequestBody
	err = json.Unmarshal(body, &ab)
	if err != nil {
		logs.Debug("Unmarshal RequestBody ,err:%v", err)
		utils.ReturnHTTPError(&c.Controller, 400, "请求参数错误")
		return
	}
	intGoodsID := ab.GoodsID
	specIDs := ab.SpecIDs
	intNumber := ab.Number
	intUserID := getLoginUserID()
	o := orm.NewOrm()
	//查询商品
	goodsTable := new(models.NsGoods)
	var goods models.NsGoods
	err = o.QueryTable(goodsTable).Filter("goods_id", intGoodsID).One(&goods)
	if err == orm.ErrNoRows {
		logs.Debug("goods lose ,err:%v", err)
		utils.ReturnHTTPError(&c.Controller, 400, "商品已下架")
		return
	}
	//查询库存
	skuTable := new(models.NsGoodsSku)
	var sku models.NsGoodsSku
	err = o.QueryTable(skuTable).Filter("goods_id", intGoodsID).Filter("attr_value_items", specIDs).One(&sku)
	if err == orm.ErrNoRows || sku.Stock < intNumber {
		utils.ReturnHTTPError(&c.Controller, 400, "库存不足")
		return
	}
	//开始事物
	_ = o.Begin()
	defer func() {
		if err != nil {
			_ = o.Rollback()
			logs.Debug("this has a err ,err:%v", err)
			utils.ReturnHTTPError(&c.Controller, 400, "操作撤销")
			return
		}
		err = o.Commit()
		if err != nil {
			return
		}
		utils.ReturnHTTPSuccess(&c.Controller, getCart())
		return

	}()

	cartTable := new(models.NsCart)
	var cart models.NsCart
	err = o.QueryTable(cartTable).Filter("goods_id", intGoodsID).Filter("sku_id", sku.ID).
		Filter("buyer_id", intUserID).One(&cart)
	if err == orm.ErrNoRows {
		cartData := models.NsCart{
			UserID:       intUserID,
			GoodsID:      intGoodsID,
			GoodsName:    goods.GoodsName,
			SkuID:        sku.ID,
			SkuName:      sku.SkuName,
			Price:        sku.Price,
			Number:       intNumber,
			GoodsPicture: goods.Picture,
		}
		_, err = o.Insert(&cartData)
	} else {
		if sku.Stock < (intNumber + cart.Number) {
			utils.ReturnHTTPError(&c.Controller, 400, "库存不足")
			return
		}
		_, err = o.QueryTable(cartTable).Filter("cart_id", cart.ID).Update(orm.Params{"number": orm.ColValue(orm.ColAdd, intNumber)})
	}
	return
}

// CartUpdateBody CartUpdateBody
type CartUpdateBody struct {
	GoodsID int `json:"goodsId"`
	SkuID   int `json:"skuId"`
	Number  int `json:"number"`
	ID      int `json:"id"`
}

// CartUpdate CartUpdate
func (c *CartController) CartUpdate() {

	var ub CartUpdateBody
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &ub)
	logs.Debug(ub)
	if err != nil {
		logs.Debug("Unmarshal RequestBody ,err:%v", err)
	}

	intGoodsID := ub.GoodsID
	skuID := ub.SkuID
	intNumber := ub.Number
	intid := ub.ID
	o := orm.NewOrm()
	goodsTable := new(models.NsGoods)
	var goods models.NsGoods
	err = o.QueryTable(goodsTable).Filter("goods_id", intGoodsID).One(&goods)
	if err == orm.ErrNoRows {
		utils.ReturnHTTPError(&c.Controller, 400, "商品已下架")
		c.ServeJSON()
		return
	}

	skuTable := new(models.NsGoodsSku)
	var sku models.NsGoodsSku
	err = o.QueryTable(skuTable).Filter("goods_id", intGoodsID).Filter("sku_id", skuID).One(&sku)
	if err == orm.ErrNoRows || sku.Stock < intNumber {
		utils.ReturnHTTPError(&c.Controller, 400, "库存不足")
		c.ServeJSON()
		return
	}
	defer func() {
		if err != nil {
			_ = o.Rollback()
			logs.Debug("this has a err ,err:%v", err)
			utils.ReturnHTTPError(&c.Controller, 400, "操作撤销")
			c.ServeJSON()
			return
		}
		err = o.Commit()
		if err != nil {
			return
		}
		utils.ReturnHTTPSuccess(&c.Controller, getCart())
		c.ServeJSON()

	}()

	cartTable := new(models.NsCart)
	var cart models.NsCart
	err = o.QueryTable(cartTable).Filter("id", intid).One(&cart)
	if err != nil {
		logs.Info("select end,cart:%v,err:%v", cart, err)
		return
	}
	if sku.Stock < intNumber {
		utils.ReturnHTTPError(&c.Controller, 400, "库存不足")
		c.ServeJSON()
		return
	}
	//开始事物
	err = o.Begin()
	_, err = o.QueryTable(cartTable).Filter("id", intid).Update(orm.Params{
		"num":    intNumber,
		"price":  sku.Price,
		"sku_id": sku.ID,
	})

}

// CartCheckedBody CartCheckedBody
type CartCheckedBody struct {
	IsChecked int         `json:"isChecked"`
	IDs       interface{} `json:"ids"`
}

// CartChecked CartChecked
func (c *CartController) CartChecked() {

	var cb CartCheckedBody
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &cb)
	if err != nil {
		logs.Debug("json Unmarshal error，err:%v", err)
	}
	intisChecked := cb.IsChecked

	if cb.IDs == "" {
		c.Abort("删除出错")
	}
	var IDsArray []string
	switch val := cb.IDs.(type) {
	// 单选
	case float64:
		IDsArray = append(IDsArray, utils.Int2String(int(val)))
	//多选
	case string:
		IDsArray = strings.Split(val, ",")
	default:

	}

	o := orm.NewOrm()
	cartTable := new(models.NsCart)
	_, _ = o.QueryTable(cartTable).Filter("cart_id__in", IDsArray).Update(orm.Params{
		"checked": intisChecked,
	})

	utils.ReturnHTTPSuccess(&c.Controller, getCart())
	c.ServeJSON()
}

// CartDeleteBody CartDeleteBody
type CartDeleteBody struct {
	IDs string `json:"ids"`
}

// CartDelete CartDelete
func (c *CartController) CartDelete() {

	var input CartDeleteBody
	body := c.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &input)
	if err != nil {
		logs.Debug("json Unmarshal err，err:%v", err)
	}
	intUserID := getLoginUserID()

	idsArray := strings.Split(input.IDs, ",")

	o := orm.NewOrm()
	_ = o.Begin()
	cartTable := new(models.NsCart)
	_, err = o.QueryTable(cartTable).Filter("cart_id__in", idsArray).Filter("buyer_id", intUserID).Delete()
	defer func() {
		if err != nil {
			_ = o.Rollback()
			logs.Debug("this has a err ,err:%v", err)
			utils.ReturnHTTPError(&c.Controller, 400, "删除失败")
			c.ServeJSON()
			return
		}
		err = o.Commit()
		if err != nil {
			return
		}
		utils.ReturnHTTPSuccess(&c.Controller, getCart())
		c.ServeJSON()

	}()
	return
}

// CartGoodsCount CartGoodsCount
func (c *CartController) CartGoodsCount() {

	cartData := getCart()
	goodsCount := GoodsCount{CartTotal: CartTotal{GoodsCount: cartData.CartTotal.GoodsCount}}
	utils.ReturnHTTPSuccess(&c.Controller, goodsCount)
	c.ServeJSON()
}

// CartAddress CartAddress
type CartAddress struct {
	models.NsMemberExpressAddress
	ProvinceName string `json:"province_name"`
	CityName     string `json:"city_name"`
	DistrictName string `json:"district_name"`
	FullRegion   string `json:"full_region"`
}

// CheckoutRtnJSON CheckoutRtnJSON
type CheckoutRtnJSON struct {
	CheckedAddress   CartAddress `json:"checkedAddress"`
	FreightPrice     float64     `json:"freightPrice"`
	CouponPrice      float64     `json:"couponPrice"`
	CheckedGoodsList []newCart   `json:"checkedGoodsList"`
	GoodsTotalPrice  float64     `json:"goodsTotalPrice"`
	OrderTotalPrice  float64     `json:"orderTotalPrice"`
	ActualPrice      float64     `json:"actualPrice"`
}

// CheckoutInput CheckoutInput
type CheckoutInput struct {
	AddressID int    `json:"address_id"`
	CouponID  string `json:"coupon_id"`
}

// CartCheckout CartCheckout
func (c *CartController) CartCheckout() {
	var (
		goodsTotalPrice float64
		orderTotalPrice float64
		actualPrice     float64
		freightPrice    float64
	)
	ids := c.GetString("ids")
	// addressID := c.GetString("intAddressID")
	if ids == "" {
		utils.ReturnHTTPError(&c.Controller, 400, "商品未找到")
		c.ServeJSON()
		return
	}
	o := orm.NewOrm()
	//获取用户地址

	// intAddressID := utils.String2Int(addressID)
	addressTable := new(models.NsMemberExpressAddress)
	var myaddress models.NsMemberExpressAddress
	err := o.QueryTable(addressTable).Filter("uid", getLoginUserID()).Filter("is_default", 1).One(&myaddress)
	var customaddress CartAddress
	if err != orm.ErrNoRows {
		customaddress.NsMemberExpressAddress = myaddress
		customaddress.ProvinceName = models.GetProvinceName(myaddress.ProvinceID)
		customaddress.CityName = models.GetCityName(myaddress.CityID)
		customaddress.DistrictName = models.GetDistrictName(myaddress.DistrictID)
		customaddress.FullRegion = customaddress.ProvinceName + customaddress.CityName + customaddress.DistrictName
	}
	//获取购物车
	intUserID := getLoginUserID()

	qb, _ := orm.NewQueryBuilder("mysql")
	var carts []newCart
	qb.Select("c.*", "p.pic_cover", "p.pic_cover_mid", "p.pic_cover_big", "s.sku_name").
		From(" ns_cart c").
		LeftJoin("sys_album_picture p").
		On("c.goods_picture = p.pic_id").
		LeftJoin("ns_goods_sku s").
		On("c.sku_id = s.sku_id").
		Where(fmt.Sprintf("c.buyer_id = %d", intUserID)).
		And(fmt.Sprintf("c.cart_id IN (%v)", ids))
	sql := qb.String()
	num, err := o.Raw(sql).QueryRows(&carts)
	if err != nil {
		logs.Debug("get cart error，num:%v,err:%v", num, err)
	}
	// logs.Info(carts)
	for _, v := range carts {
		goodsTotalPrice += v.Price * float64(v.Number)
		orderTotalPrice += v.Price * float64(v.Number)
		actualPrice += v.Price * float64(v.Number)
	}
	utils.ReturnHTTPSuccess(&c.Controller, CheckoutRtnJSON{
		CheckedAddress:   customaddress,
		CheckedGoodsList: carts,
		GoodsTotalPrice:  goodsTotalPrice,
		OrderTotalPrice:  orderTotalPrice,
		ActualPrice:      actualPrice,
		FreightPrice:     freightPrice,
	})
	c.ServeJSON()
}
