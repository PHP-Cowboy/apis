package initialize

import (
	"github.com/gin-gonic/gin"
	"shop-api/router"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	group := r.Group("/v1")

	router.UserRoute(group)

	return r
}
