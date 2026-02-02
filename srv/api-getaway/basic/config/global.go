package config

import (
	"github.com/redis/go-redis/v9"
	__ "github.com/yuhang-jieke/yuedemo2/srv/api-getaway/basic/proto"
)

var (
	UserClient __.UserClient
	Rdb        *redis.Client
)
