package initialize

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"shop-api/user-web/global"
)

func InitConfig() {
	v := viper.New()

	v.SetConfigFile("config.yaml")

	err := v.ReadInConfig()
	if err != nil {
		zap.S().Panicf("读取配置文件失败:", err.Error())
	}

	err = v.Unmarshal(global.ServerConfig)

	if err != nil {
		zap.S().Panicf("解析配置文件到结构体出错:", err.Error())
	}

}
