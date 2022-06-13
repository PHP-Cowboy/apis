package router

import (
	"github.com/gin-gonic/gin"
	"shop-api/goods-web/api/goods"
	"shop-api/goods-web/middlewares"
)

func GoodsRoute(g *gin.RouterGroup) {
	userGroup := g.Group("/goods")
	{
		userGroup.GET("/list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.GetList)
	}
}
