package router

import (
	"apis/goods-web/api/banners"
	"apis/goods-web/middlewares"
	"github.com/gin-gonic/gin"
)

func BannerRoute(g *gin.RouterGroup) {
	bannerGroup := g.Group("/banners")
	{
		bannerGroup.GET("/list", banners.List)
		bannerGroup.POST("/new", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banners.New)      //改接口需要管理员权限
		bannerGroup.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banners.Delete) //删除
		bannerGroup.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banners.Update)    //修改轮播图信息
	}
}
