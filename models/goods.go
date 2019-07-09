package models

import (
	"github.com/astaxie/beego/orm"
)

type SpecificationItem struct {
	list []NsGoodsSku
}

func GetProductList(goodsId int) []NsGoods {
	o := orm.NewOrm()

	var products []NsGoods
	product := new(NsGoods)

	o.QueryTable(product).Filter("goods_id", goodsId).All(&products)

	return products

}
func GetGoodSku(goodsId int) []NsGoodsSku {
	var skuList []NsGoodsSku
	o := orm.NewOrm()
	sku := new(NsGoodsSku)
	o.QueryTable(sku).Filter("goods_id", goodsId).All(&skuList)
	return skuList

}

func GetGoodsList(andWhere string, limit int) (list []NsNewGoods) {
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("g.goods_id", "g.goods_name", "g.brand_id", "g.category_id", "g.price", "p.pic_cover", "p.pic_cover_mid", "p.pic_cover_big").
		From("ns_goods g").
		LeftJoin("sys_album_picture p").
		On("g.picture = p.pic_id").
		Where(andWhere).
		Limit(limit)

	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql).QueryRows(&list)
	return
}
