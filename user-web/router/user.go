package router

import (
	"github.com/gin-gonic/gin"
	"shop-api/user-web/api"
	"shop-api/user-web/middlewares"
)

func UserRoute(g *gin.RouterGroup) {
	userGroup := g.Group("/user")
	{
		userGroup.GET("/list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		userGroup.POST("/pwd_login", api.PasswordLogin)
	}
}
