package router

import (
	"github.com/gin-gonic/gin"
	"shop-api/user-web/api"
)

func BaseRoute(g *gin.RouterGroup) {
	userGroup := g.Group("/base")
	{
		userGroup.GET("/captcha", api.GenerateCaptcha)
	}
}
