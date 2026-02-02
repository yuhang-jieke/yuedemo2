package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuhang-jieke/yuedemo2/srv/api-getaway/basic/config"
	"github.com/yuhang-jieke/yuedemo2/srv/api-getaway/pkg"
)

func AuthTokenh() gin.HandlerFunc {
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
		_, err := pkg.PersonToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		/*userIdStr, ok := personToken["userId"].(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "解析用户ID失败",
			})
			c.Abort()
			return
		}*/

		res, err := config.Rdb.SIsMember(context.Background(), "list:user", token).Result()
		if err != nil {
			log.Printf("Redis SIsMember 执行失败: %v, Key: list:user, UserID: %s", err, token)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "黑名单查询失败",
			})
			c.Abort()
			return
		}
		res, err = config.Rdb.SIsMember(context.Background(), "list:user", token).Result()
		if res {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "被加入黑名单",
			})
			c.Abort()
			return
		}
	}

}
