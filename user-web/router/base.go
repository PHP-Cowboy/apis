package router

import (
	"apis/user-web/api"
	"github.com/gin-gonic/gin"
)

func BaseRoute(g *gin.RouterGroup) {
	userGroup := g.Group("/base")
	{
		userGroup.GET("/captcha", api.GenerateCaptcha)
	}
}
