package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yuhang-jieke/yuedemo2/srv/api-getaway/basic/config"
	"github.com/yuhang-jieke/yuedemo2/srv/api-getaway/basic/inits"
	__ "github.com/yuhang-jieke/yuedemo2/srv/api-getaway/basic/proto"
	"github.com/yuhang-jieke/yuedemo2/srv/api-getaway/handler/request"
	"github.com/yuhang-jieke/yuedemo2/srv/api-getaway/pkg"
)

func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "注册失败",
		})
		return
	}
	_, err := config.UserClient.Register(c, &__.RegisterReq{
		Name:    form.Name,
		Age:     int64(form.Age),
		Address: form.Address,
	})
	result, _ := config.Rdb.Get(context.Background(), "user"+form.Name).Result()
	if form.Name == result {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "用户已经注册",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "注册失败",
		})
		return
	}
	config.Rdb.Set(context.Background(), "user"+form.Name, form.Name, 0)
	go func() {
		usermap := map[string]string{
			"name":    form.Name,
			"age":     strconv.Itoa(form.Age),
			"address": form.Address,
		}
		_, err2 := inits.ElasticClient.Index().Index("usermap").BodyJson(usermap).Do(context.Background())
		if err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
				"msg":  "es同步失败",
			})
			return
		}
	}()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
	return
}
func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "登录失败",
		})
		return
	}
	r, err := config.UserClient.Login(c, &__.LoginReq{
		Name: form.Name,
		Age:  int64(form.Age),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "登录失败",
		})
		return
	}

	userIdStr := strconv.FormatInt(r.UserId, 10)
	token, _ := pkg.TokenHandler(userIdStr)
	config.Rdb.SAdd(context.Background(), "list:user", token, 1122)
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "登录成功",
		"token": token,
	})
	return
}
