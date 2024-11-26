package svc

import (
	"forum/app/community/rpc/communityservice"
	"forum/app/post/api/internal/config"
	"forum/app/post/rpc/postservice"
	"forum/app/user/rpc/userservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	PostRpc      postservice.PostService
	CommunityRpc communityservice.CommunityService
	UserRpc      userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		PostRpc:      postservice.NewPostService(zrpc.MustNewClient(c.PostRpcConf)),
		CommunityRpc: communityservice.NewCommunityService(zrpc.MustNewClient(c.CommunityRpcConf)),
		UserRpc:      userservice.NewUserService(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
