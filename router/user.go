package router

import (
	"github.com/gin-gonic/gin"
	"shop-api/api"
)

func UserRoute(g *gin.RouterGroup) {
	userGroup := g.Group("/user")
	{
		userGroup.GET("/list", api.GetUserList)
	}
}
