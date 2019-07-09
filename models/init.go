package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/niushop?charset=utf8mb4", 30)

	// register model
	orm.RegisterModel(new(NsPlatformAdv))
	orm.RegisterModel(new(NsPlatformAdvPosition))
	orm.RegisterModel(new(NsMemberExpressAddress))
	orm.RegisterModel(new(SysUserAdmin))
	orm.RegisterModel(new(NsAttribute))
	orm.RegisterModel(new(NsAttributeValue))

	orm.RegisterModel(new(NsGoodsBrand))
	orm.RegisterModel(new(NsCart))
	orm.RegisterModel(new(NsGoodsCategory))

	orm.RegisterModel(new(SysWeixinMenu))
	orm.RegisterModel(new(NsMemberFavorites))
	orm.RegisterModel(new(NsGoodsComment))
	orm.RegisterModel(new(NsGoodsEvaluate))
	orm.RegisterModel(new(NsGoodsSku))
	orm.RegisterModel(new(NsCoupon))

	orm.RegisterModel(new(NsPlatformLink))
	orm.RegisterModel(new(NsGoods))

	orm.RegisterModel(new(NsGoodsAttribute))
	orm.RegisterModel(new(SysAlbumClass))
	orm.RegisterModel(new(SysAlbumPicture))
	orm.RegisterModel(new(NsGoodsSpec))

	orm.RegisterModel(new(NsGoodsSpecValue))
	orm.RegisterModel(new(NsOrder))

	orm.RegisterModel(new(NsOrderGoodsExpress))
	orm.RegisterModel(new(NsOrderGoods))

	orm.RegisterModel(new(SysArea))
	orm.RegisterModel(new(SysProvince))
	orm.RegisterModel(new(SysCity))
	orm.RegisterModel(new(SysDistrict))
	orm.RegisterModel(new(NsExpressCompany))

	orm.RegisterModel(new(NsShopAd))

	orm.RegisterModel(new(NsMember))
	orm.RegisterModel(new(SysUser))
}
