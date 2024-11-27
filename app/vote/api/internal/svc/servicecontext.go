package svc

import (
	"forum/app/vote/api/internal/config"
	"forum/app/vote/rpc/voteservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	VoteRpc voteservice.VoteService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		VoteRpc: voteservice.NewVoteService(zrpc.MustNewClient(c.VoteRpcConf)),
	}
}
