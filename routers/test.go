package routers

import (
	"github.com/astaxie/beego"
	"github.com/tianxinbaiyun/goniushop/controllers"
)

func init() {

	beego.Router("test/index/index", &controllers.IndexController{}, "get:IndexIndex")

	beego.Router("test/catalog/index", &controllers.CatalogController{}, "get:CatalogIndex")
	beego.Router("test/catalog/current", &controllers.CatalogController{}, "get:CatalogCurrent")

	beego.Router("test/auth/loginByWeixin", &controllers.AuthController{}, "post:AuthLoginByWeixin")

	beego.Router("test/goods/count", &controllers.GoodsController{}, "get:GoodsCount")
	beego.Router("test/goods/list", &controllers.GoodsController{}, "get:GoodsList")
	beego.Router("test/goods/category", &controllers.GoodsController{}, "get:GoodsCategory")
	beego.Router("test/goods/detail", &controllers.GoodsController{}, "get:GoodsDetail")
	beego.Router("test/goods/new", &controllers.GoodsController{}, "get:GoodsNew")
	beego.Router("test/goods/hot", &controllers.GoodsController{}, "get:GoodsHot")
	// beego.Router("test/goods/related", &controllers.GoodsController{}, "get:Goods_Related")

	beego.Router("test/brand/list", &controllers.BrandController{}, "get:BrandList")
	beego.Router("test/brand/detail", &controllers.BrandController{}, "get:BrandDetail")

	beego.Router("test/cart/index", &controllers.CartController{}, "get:CartIndex")
	beego.Router("test/cart/add", &controllers.CartController{}, "post:CartAdd")
	beego.Router("test/cart/update", &controllers.CartController{}, "post:CartUpdate")
	beego.Router("test/cart/delete", &controllers.CartController{}, "post:CartDelete")
	beego.Router("test/cart/checked", &controllers.CartController{}, "post:CartChecked")
	beego.Router("test/cart/goodsCount", &controllers.CartController{}, "get:CartGoodsCount")
	beego.Router("test/cart/checkout", &controllers.CartController{}, "get:CartCheckout")

	beego.Router("test/pay/prepay", &controllers.PayController{}, "get:PayPrepay")

	beego.Router("test/collect/list", &controllers.CollectController{}, "get:CollectList")
	beego.Router("test/collect/addordelete", &controllers.CollectController{}, "post:CollectAddOrDelete")

	beego.Router("test/comment/list", &controllers.CommentController{}, "get:CommentList")
	beego.Router("test/comment/count", &controllers.CommentController{}, "get:CommentCount")
	beego.Router("test/comment/post", &controllers.CommentController{}, "post:CommentPost")

	beego.Router("test/address/list", &controllers.AddressController{}, "get:AddressList")
	beego.Router("test/address/detail", &controllers.AddressController{}, "get:AddressDetail")
	beego.Router("test/address/save", &controllers.AddressController{}, "post:AddressSave")
	beego.Router("test/address/delete", &controllers.AddressController{}, "post:AddressDelete")

	beego.Router("test/region/list", &controllers.RegionController{}, "get:RegionList")

	beego.Router("test/order/submit", &controllers.OrderController{}, "post:OrderSubmit")
	beego.Router("test/order/list", &controllers.OrderController{}, "get:OrderList")
	beego.Router("test/order/detail", &controllers.OrderController{}, "get:OrderDetail")
	//beego.Router("test/order/cancel", &controllers.OrderController{}, "post:AddressSave")
	beego.Router("test/order/express", &controllers.OrderController{}, "get:OrderExpress")

}
