package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop-api/goods-web/middlewares"
	"shop-api/goods-web/router"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	r.Use(middlewares.Cors())

	group := r.Group("/g/v1")

	router.GoodsRoute(group)     //商品
	router.BannerRoute(group)    //banner
	router.BrandRouter(group)    //品牌
	router.CategoryRouter(group) //分类

	return r
}
