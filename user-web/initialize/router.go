package initialize

import (
	"github.com/gin-gonic/gin"
	"shop-api/user-web/router"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	group := r.Group("/v1")

	router.UserRoute(group)

	return r
}
