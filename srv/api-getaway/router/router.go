package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuhang-jieke/yuedemo2/srv/api-getaway/handler/server"
	"github.com/yuhang-jieke/yuedemo2/srv/api-getaway/middleware"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/register", server.Register)
	r.POST("/login", server.Login)
	r.POST("/token", middleware.AuthTokenh(), server.TokenCreate)
	//r.POST("/upload", middleware.AuthToken2(), server.Upload)
	return r

}
