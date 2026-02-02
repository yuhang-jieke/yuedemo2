package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuhang-jieke/yuedemo2/srv/api-getaway/pkg"
)

var (
	logfile *os.File
	lodmu   sync.Mutex
)

func LogInit() {
	var err error
	logfile, err = os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败")
	}
}
func TokenLog(path, userId, status string) {
	lodmu.Lock()
	defer lodmu.Unlock()
	data, _ := json.Marshal(map[string]string{
		"path":   path,
		"userId": userId,
		"status": status,
		"time":   time.Now().Format("2006-01-02-15"),
	})
	logfile.Write(append(data, '\n'))
}
func AuthTokenLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		userId := ""
		token := c.Request.Header.Get("token")
		if token == "" {
			TokenLog(path, userId, "未登录")
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "未登录",
			})
			c.Abort()
			return
		}
		personToken, err := pkg.PersonToken(token)
		if err != nil {
			TokenLog(path, userId, err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("userId", personToken["userId"])
		c.Next()
	}
}
