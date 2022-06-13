package initialize

import (
	"github.com/gin-gonic/gin"
	"shop-api/goods-web/middlewares"
	"shop-api/goods-web/router"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.Cors())

	group := r.Group("/v1")

	router.GoodsRoute(group)

	return r
}
