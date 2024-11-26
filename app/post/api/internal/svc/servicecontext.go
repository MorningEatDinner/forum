package svc

import (
	"forum/app/community/rpc/communityservice"
	"forum/app/post/api/internal/config"
	"forum/app/post/rpc/postservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	PostRpc      postservice.PostService
	CommunityRpc communityservice.CommunityService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		PostRpc:      postservice.NewPostService(zrpc.MustNewClient(c.PostRpcConf)),
		CommunityRpc: communityservice.NewCommunityService(zrpc.MustNewClient(c.CommunityRpcConf)),
	}
}
