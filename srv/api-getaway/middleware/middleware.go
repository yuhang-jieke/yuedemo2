package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuhang-jieke/yuedemo2/srv/api-getaway/basic/config"
	"github.com/yuhang-jieke/yuedemo2/srv/api-getaway/pkg"
)

func AuthTokenc() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "未登录",
			})
			c.Abort()
			return
		}
		personToken, err := pkg.PersonToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		config.Rdb.Exists(context.Background(), "key")
		result, _ := config.Rdb.Incr(context.Background(), "key").Result()
		if result > 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "请勿重复操作",
			})
			c.Abort()
			return
		}
		c.Set("userId", personToken["userId"])
		c.Next()
	}
}
