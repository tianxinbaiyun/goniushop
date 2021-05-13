package routers

import (
	"github.com/astaxie/beego"
	"github.com/tianxinbaiyun/goniushop/controllers"
)

func init() {

	beego.Router("api/index/index", &controllers.IndexController{}, "get:IndexIndex")

	beego.Router("api/catalog/index", &controllers.CatalogController{}, "get:CatalogIndex")
	beego.Router("api/catalog/current", &controllers.CatalogController{}, "get:CatalogCurrent")

	beego.Router("api/auth/loginByWeixin", &controllers.AuthController{}, "post:AuthLoginByWeixin")

	beego.Router("api/goods/count", &controllers.GoodsController{}, "get:GoodsCount")
	beego.Router("api/goods/list", &controllers.GoodsController{}, "get:GoodsList")
	beego.Router("api/goods/category", &controllers.GoodsController{}, "get:GoodsCategory")
	beego.Router("api/goods/detail", &controllers.GoodsController{}, "get:GoodsDetail")
	beego.Router("api/goods/new", &controllers.GoodsController{}, "get:GoodsNew")
	beego.Router("api/goods/hot", &controllers.GoodsController{}, "get:GoodsHot")
	// beego.Router("api/goods/related", &controllers.GoodsController{}, "get:Goods_Related")

	beego.Router("api/brand/list", &controllers.BrandController{}, "get:BrandList")
	beego.Router("api/brand/detail", &controllers.BrandController{}, "get:BrandDetail")

	beego.Router("api/cart/index", &controllers.CartController{}, "get:CartIndex")
	beego.Router("api/cart/add", &controllers.CartController{}, "post:CartAdd")
	beego.Router("api/cart/update", &controllers.CartController{}, "post:CartUpdate")
	beego.Router("api/cart/delete", &controllers.CartController{}, "post:CartDelete")
	beego.Router("api/cart/checked", &controllers.CartController{}, "post:CartChecked")
	beego.Router("api/cart/goodsCount", &controllers.CartController{}, "get:CartGoodsCount")
	beego.Router("api/cart/checkout", &controllers.CartController{}, "get:CartCheckout")

	beego.Router("api/pay/prepay", &controllers.PayController{}, "get:PayPrepay")

	beego.Router("api/collect/list", &controllers.CollectController{}, "get:CollectList")
	beego.Router("api/collect/addordelete", &controllers.CollectController{}, "post:CollectAddOrDelete")

	beego.Router("api/comment/list", &controllers.CommentController{}, "get:CommentList")
	beego.Router("api/comment/count", &controllers.CommentController{}, "get:CommentCount")
	beego.Router("api/comment/post", &controllers.CommentController{}, "post:CommentPost")

	beego.Router("api/address/list", &controllers.AddressController{}, "get:AddressList")
	beego.Router("api/address/detail", &controllers.AddressController{}, "get:AddressDetail")
	beego.Router("api/address/save", &controllers.AddressController{}, "post:AddressSave")
	beego.Router("api/address/delete", &controllers.AddressController{}, "post:AddressDelete")

	beego.Router("api/region/list", &controllers.RegionController{}, "get:RegionList")

	beego.Router("api/order/submit", &controllers.OrderController{}, "post:OrderSubmit")
	beego.Router("api/order/list", &controllers.OrderController{}, "get:OrderList")
	beego.Router("api/order/detail", &controllers.OrderController{}, "get:OrderDetail")
	//beego.Router("api/order/cancel", &controllers.OrderController{}, "post:AddressSave")
	beego.Router("api/order/express", &controllers.OrderController{}, "get:OrderExpress")

}
