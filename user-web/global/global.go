package global

import (
	"shop-api/user-web/config"
	"shop-api/user-web/proto/proto"
)

var (
	ServerConfig = &config.ServerConfig{}
	UserClient   proto.UserClient
)
