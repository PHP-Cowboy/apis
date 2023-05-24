package router

import (
	"apis/goods-web/api/brands"
	"github.com/gin-gonic/gin"
)

func BrandRouter(Router *gin.RouterGroup) {
	brandGroup := Router.Group("brands")
	{
		brandGroup.GET("", brands.BrandList)          // 品牌列表页
		brandGroup.DELETE("/:id", brands.DeleteBrand) // 删除品牌
		brandGroup.POST("", brands.NewBrand)          //新建品牌
		brandGroup.PUT("/:id", brands.UpdateBrand)    //修改品牌信息
	}

	categoryBrandGroup := Router.Group("category_brands")
	{
		categoryBrandGroup.GET("", brands.CategoryBrandList)          // 类别品牌列表页
		categoryBrandGroup.DELETE("/:id", brands.DeleteCategoryBrand) // 删除类别品牌
		categoryBrandGroup.POST("", brands.NewCategoryBrand)          //新建类别品牌
		categoryBrandGroup.PUT("/:id", brands.UpdateCategoryBrand)    //修改类别品牌
		categoryBrandGroup.GET("/:id", brands.GetCategoryBrandList)   //获取分类的品牌
	}
}
