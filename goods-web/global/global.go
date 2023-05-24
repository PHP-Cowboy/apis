package global

import (
	"apis/goods-web/config"
	"apis/goods-web/proto/proto"
)

var (
	ServerConfig = &config.ServerConfig{}
	NacosConfig  = &config.NacosConfig{}
	GoodsClient  proto.GoodsClient
)
