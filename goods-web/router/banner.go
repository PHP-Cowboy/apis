package router

import (
	"github.com/gin-gonic/gin"
	"shop-api/goods-web/api/banners"
	"shop-api/goods-web/middlewares"
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
