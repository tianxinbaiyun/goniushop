package routers

import (
	"github.com/astaxie/beego"
	"github.com/tianxinbaiyun/goniushop/controllers"
)

func init() {

	beego.Router("test/index/index", &controllers.IndexController{}, "get:Index_Index")

	beego.Router("test/catalog/index", &controllers.CatalogController{}, "get:Catalog_Index")
	beego.Router("test/catalog/current", &controllers.CatalogController{}, "get:Catalog_Current")

	beego.Router("test/auth/loginByWeixin", &controllers.AuthController{}, "post:Auth_LoginByWeixin")

	beego.Router("test/goods/count", &controllers.GoodsController{}, "get:Goods_Count")
	beego.Router("test/goods/list", &controllers.GoodsController{}, "get:Goods_List")
	beego.Router("test/goods/category", &controllers.GoodsController{}, "get:Goods_Category")
	beego.Router("test/goods/detail", &controllers.GoodsController{}, "get:Goods_Detail")
	beego.Router("test/goods/new", &controllers.GoodsController{}, "get:Goods_New")
	beego.Router("test/goods/hot", &controllers.GoodsController{}, "get:Goods_Hot")
	// beego.Router("test/goods/related", &controllers.GoodsController{}, "get:Goods_Related")

	beego.Router("test/brand/list", &controllers.BrandController{}, "get:Brand_List")
	beego.Router("test/brand/detail", &controllers.BrandController{}, "get:Brand_Detail")

	beego.Router("test/cart/index", &controllers.CartController{}, "get:Cart_Index")
	beego.Router("test/cart/add", &controllers.CartController{}, "post:Cart_Add")
	beego.Router("test/cart/update", &controllers.CartController{}, "post:Cart_Update")
	beego.Router("test/cart/delete", &controllers.CartController{}, "post:Cart_Delete")
	beego.Router("test/cart/checked", &controllers.CartController{}, "post:Cart_Checked")
	beego.Router("test/cart/goodscount", &controllers.CartController{}, "get:Cart_GoodsCount")
	beego.Router("test/cart/checkout", &controllers.CartController{}, "get:Cart_Checkout")

	beego.Router("test/pay/prepay", &controllers.PayController{}, "get:Pay_Prepay")

	beego.Router("test/collect/list", &controllers.CollectController{}, "get:Collect_List")
	beego.Router("test/collect/addordelete", &controllers.CollectController{}, "post:Collect_AddorDelete")

	beego.Router("test/comment/list", &controllers.CommentController{}, "get:Comment_List")
	beego.Router("test/comment/count", &controllers.CommentController{}, "get:Comment_Count")
	beego.Router("test/comment/post", &controllers.CommentController{}, "post:Comment_Post")

	beego.Router("test/address/list", &controllers.AddressController{}, "get:Address_List")
	beego.Router("test/address/detail", &controllers.AddressController{}, "get:Address_Detail")
	beego.Router("test/address/save", &controllers.AddressController{}, "post:Address_Save")
	beego.Router("test/address/delete", &controllers.AddressController{}, "post:Address_Delete")

	beego.Router("test/region/list", &controllers.RegionController{}, "get:Region_List")

	beego.Router("test/order/submit", &controllers.OrderController{}, "post:Order_Submit")
	beego.Router("test/order/list", &controllers.OrderController{}, "get:Order_List")
	beego.Router("test/order/detail", &controllers.OrderController{}, "get:Order_Detail")
	//beego.Router("test/order/cancel", &controllers.OrderController{}, "post:Address_Save")
	beego.Router("test/order/express", &controllers.OrderController{}, "get:Order_Express")

}
