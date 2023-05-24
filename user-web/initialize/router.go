package initialize

import (
	"apis/user-web/middlewares"
	"apis/user-web/router"
	"github.com/gin-gonic/gin"
	"net/http"
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

	group := r.Group("/u/v1")

	router.UserRoute(group)

	router.BaseRoute(group)

	return r
}
