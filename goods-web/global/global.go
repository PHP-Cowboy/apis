package global

import (
	"shop-api/goods-web/config"
	"shop-api/goods-web/proto/proto"
)

var (
	ServerConfig = &config.ServerConfig{}
	NacosConfig  = &config.NacosConfig{}
	GoodsClient  proto.GoodsClient
)
