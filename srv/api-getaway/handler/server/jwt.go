package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuhang-jieke/yuedemo2/srv/api-getaway/pkg"
)

func TokenCreate(c *gin.Context) {
	oldtoken := c.GetHeader("token")
	if oldtoken == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "不能为空",
		})
		return
	}
	newtoken, err := pkg.CreateToken(oldtoken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "生成新token失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "生成新token成功",
		"token": newtoken,
	})
	return

}
