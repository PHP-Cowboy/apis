package initialize

import (
	myValidator "apis/goods-web/validator"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("mobile", myValidator.Mobile); err != nil {
			zap.S().Panicf("mobile自定义验证器加载失败:", err.Error())
		}
	}
}
