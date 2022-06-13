package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop-api/user-web/middlewares"
	"shop-api/user-web/router"
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

	group := r.Group("/v1")

	router.UserRoute(group)

	router.BaseRoute(group)

	return r
}
