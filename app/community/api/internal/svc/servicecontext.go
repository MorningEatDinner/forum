package svc

import (
	"forum/app/community/api/internal/config"
	"forum/app/community/rpc/communityservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	CommunityRpc communityservice.CommunityService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		CommunityRpc: communityservice.NewCommunityService(zrpc.MustNewClient(c.CommunityRpcConf)),
	}
}
