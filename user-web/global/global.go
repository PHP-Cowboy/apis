package global

import (
	"apis/user-web/config"
	"apis/user-web/proto/proto"
)

var (
	ServerConfig = &config.ServerConfig{}
	NacosConfig  = &config.NacosConfig{}
	UserClient   proto.UserClient
)
