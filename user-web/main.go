package main

import (
	"apis/user-web/global"
	"apis/user-web/initialize"
	"apis/user-web/utils/register/consul"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	initialize.InitLogger()

	initialize.InitConfig()

	initialize.InitSrvConn()

	serverConfig := global.ServerConfig

	zap.S().Info(serverConfig)

	initialize.InitValidator()

	g := initialize.InitRouter()

	serviceId := fmt.Sprintf("%s", uuid.NewV4())

	registryClient := consul.NewRegistryClient(serverConfig.ConsulInfo.Host, serverConfig.ConsulInfo.Port)

	err := registryClient.Register(serverConfig.Host, serverConfig.Port, serverConfig.Name, serverConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panicf("服务注册失败:", err.Error())
	}

	zap.S().Info("服务启动中,端口:", serverConfig.Port)

	go func() {
		err := g.Run(fmt.Sprintf(":%d", serverConfig.Port))
		if err != nil {
			zap.S().Panicf("启动失败:", err.Error())
		}
	}()

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = registryClient.ServiceDeregister(serviceId); err != nil {
		zap.S().Info("注销失败:", err.Error())
	} else {
		zap.S().Info("注销成功:")
	}
}
