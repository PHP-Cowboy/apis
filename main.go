package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"

	"shop-api/initialize"
)

func main() {
	port := flag.Int("port", 8021, "端口号")

	g := initialize.InitRouter()

	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)

	zap.S().Info("服务启动中,端口:", *port)
	if err := g.Run(fmt.Sprintf(":%d", *port)); err != nil {
		zap.S().Panicf("启动失败:", err.Error())
	}
}
