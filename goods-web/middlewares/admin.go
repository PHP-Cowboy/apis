package middlewares

import (
	"apis/goods-web/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsAdminAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		claims, ok := c.Get("claims")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "claims获取失败",
			})
			c.Abort()
			return
		}

		userInfo := claims.(*models.CustomClaims)

		if userInfo.AuthorityId != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "无权限",
			})
			c.Abort()
			return
		}

		c.Next()
	}

}
