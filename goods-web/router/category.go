package router

import (
	"apis/goods-web/api/category"
	"github.com/gin-gonic/gin"
)

func CategoryRouter(Router *gin.RouterGroup) {
	CategoryGroup := Router.Group("categorys")
	{
		CategoryGroup.GET("/list", category.List)     // 商品类别列表页
		CategoryGroup.DELETE("/:id", category.Delete) // 删除分类
		CategoryGroup.GET("/:id", category.Detail)    // 获取分类详情
		CategoryGroup.POST("/new", category.New)      //新建分类
		CategoryGroup.PUT("/:id", category.Update)    //修改分类信息
	}
}
