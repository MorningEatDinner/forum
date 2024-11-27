package svc

import (
	"forum/app/comment/api/internal/config"
	"forum/app/comment/rpc/commentservice"
	"forum/app/post/rpc/postservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	PostRpc    postservice.PostService
	CommentRpc commentservice.CommentService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		PostRpc:    postservice.NewPostService(zrpc.MustNewClient(c.PostRpcConf)),
		CommentRpc: commentservice.NewCommentService(zrpc.MustNewClient(c.CommentRpcConf)),
	}
}
