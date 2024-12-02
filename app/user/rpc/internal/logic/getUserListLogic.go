package logic

import (
	"context"
	"forum/common/xerr"
	"github.com/pkg/errors"

	"forum/app/user/rpc/internal/svc"
	"forum/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {
	userList, err := l.svcCtx.UserModel.FindUserList(l.ctx, in.LastUserId, int(in.BatchSize))
	if err != nil {
		logx.WithContext(l.ctx).Errorf("failed to get user list: %v", err)
		return nil, errors.Wrapf(xerr.NewErrMsg("failed to get user list"), "failed to get user list")
	}
	lastId := userList[len(userList)-1].UserId
	users := make([]*pb.UserInfo, in.BatchSize)
	for i, user := range userList {
		users[i] = &pb.UserInfo{
			Username: user.Username,
			Email:    user.Email,
		}
	}

	return &pb.GetUserListResponse{
		Users:      users,
		LastUserId: lastId,
	}, nil
}
