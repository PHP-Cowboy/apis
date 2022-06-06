package initialize

import (
	"github.com/gin-gonic/gin"
	"shop-api/user-web/middlewares"
	"shop-api/user-web/router"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.Cors())

	group := r.Group("/v1")

	router.UserRoute(group)

	return r
}
