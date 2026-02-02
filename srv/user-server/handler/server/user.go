package server

import (
	"context"
	"errors"

	"github.com/yuhang-jieke/yuedemo2/srv/user-server/basic/config"
	__ "github.com/yuhang-jieke/yuedemo2/srv/user-server/handler/proto"
	"github.com/yuhang-jieke/yuedemo2/srv/user-server/model"
)

type Server struct {
	__.UnimplementedUserServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Register(_ context.Context, in *__.RegisterReq) (*__.RegisterResp, error) {
	user := model.User{
		Name:    in.Name,
		Age:     int(in.Age),
		Address: in.Address,
	}
	if err := user.Register(config.DB); err != nil {
		return nil, errors.New("注册失败")
	}
	config.Rdb.Set(context.Background(), "key"+in.Name, in.Name, 0)
	return &__.RegisterResp{
		Greet: "注册成功",
	}, nil
}
func (s *Server) Login(_ context.Context, in *__.LoginReq) (*__.LoginResp, error) {
	var user model.User
	if err := user.FindName(config.DB, in.Name); err != nil {
		return nil, errors.New("查询失败")
	}
	if in.Age != int64(user.Age) {
		return nil, errors.New("密码不正确")
	}
	return &__.LoginResp{
		UserId: int64(user.ID),
	}, nil
}
