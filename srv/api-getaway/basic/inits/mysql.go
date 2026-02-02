package inits

import (
	"fmt"

	"github.com/gospacex/gospacex/core/storage/conf"
	"github.com/gospacex/gospacex/core/storage/db/mysql"
	"github.com/yuhang-jieke/yuedemo2/srv/user-server/basic/config"
	"github.com/yuhang-jieke/yuedemo2/srv/user-server/model"
)

func MysqlInit() {
	var err error
	config.DB, err = mysql.Init(true, "debug", conf.Cfg.Mysql)
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	fmt.Println("数据库连接成功")
	err = config.DB.AutoMigrate(&model.User{}, &model.ImgContent{}, &model.TitleContent{})
	if err != nil {
		fmt.Println("数据表迁移失败")
		return
	}
	fmt.Println("数据表迁移成功")
}
