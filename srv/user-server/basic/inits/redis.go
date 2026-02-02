package inits

import (
	"fmt"

	"github.com/gospacex/gospacex/core/storage/cache/redis"
	"github.com/gospacex/gospacex/core/storage/conf"
	"github.com/yuhang-jieke/yuedemo2/srv/user-server/basic/config"
)

func RedisInit() {
	err := redis.Init(true, conf.Cfg.Redis)
	if err != nil {
		fmt.Println("redis连接失败")
		return
	}
	config.Rdb = redis.RC
	fmt.Println("redis连接成功")
}
