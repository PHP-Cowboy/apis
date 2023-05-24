package router

import (
	"apis/user-web/api"
	"apis/user-web/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoute(g *gin.RouterGroup) {
	userGroup := g.Group("/user")
	{
		userGroup.GET("/list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		userGroup.POST("/pwd_login", api.PasswordLogin)
		userGroup.POST("/register", api.Register)
		userGroup.POST("/detail", api.GetUserDetail)
		userGroup.POST("/update", api.UpdateUser)
	}
}
