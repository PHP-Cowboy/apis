package initialize

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	myValidator "shop-api/goods-web/validator"
)

func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("mobile", myValidator.Mobile); err != nil {
			zap.S().Panicf("mobile自定义验证器加载失败:", err.Error())
		}
	}
}
