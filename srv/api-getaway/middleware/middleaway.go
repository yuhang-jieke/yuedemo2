package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuhang-jieke/yuedemo2/srv/api-getaway/pkg"
)

func AuthToken2() gin.HandlerFunc {
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
		if c.Writer.Status() == http.StatusOK {
			c.Writer.Header().Set("Cache-Control", "ok,600")
		}
		c.Set("userId", personToken["userId"])
		c.Next()
	}
}
