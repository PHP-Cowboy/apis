package router

import (
	"apis/goods-web/api/goods"
	"apis/goods-web/middlewares"
	"github.com/gin-gonic/gin"
)

func GoodsRoute(g *gin.RouterGroup) {
	goodsGroup := g.Group("/goods")
	{
		goodsGroup.GET("/list", goods.GetList)
		goodsGroup.POST("/new", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.New)      //改接口需要管理员权限
		goodsGroup.GET("/:id", goods.Detail)                                                      //获取商品的详情
		goodsGroup.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Delete) //删除商品
		goodsGroup.GET("/stocks/:id", goods.Stocks)                                               //获取商品的库存

		goodsGroup.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Update)
		goodsGroup.PATCH("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.UpdateStatus)
	}
}
