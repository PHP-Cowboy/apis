package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	initialize2 "shop-api/user-web/initialize"
)

func main() {
	port := flag.Int("port", 8021, "端口号")

	initialize2.InitLogger()

	g := initialize2.InitRouter()

	zap.S().Info("服务启动中,端口:", *port)
	if err := g.Run(fmt.Sprintf(":%d", *port)); err != nil {
		zap.S().Panicf("启动失败:", err.Error())
	}
}
