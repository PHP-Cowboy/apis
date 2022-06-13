package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"shop-api/goods-web/global"
	"shop-api/goods-web/proto/proto"
)

func InitSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo

	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)

	client, err := api.NewClient(cfg)

	if err != nil {
		panic(err)
	}

	data, err := client.Agent().ServicesWithFilter(`Service == "goods-srv"`)
	if err != nil {
		return
	}

	goodsSrvHost := ""
	goodsSrvPort := 0
	if err != nil {
		panic(err)
	}
	for _, value := range data {
		goodsSrvHost = value.Address
		goodsSrvPort = value.Port
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", goodsSrvHost, goodsSrvPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("拨号失败")
	}
	global.GoodsClient = proto.NewGoodsClient(conn)
}
