package logic

import (
	"context"
	"fmt"
	"forum/common/globalkey"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"

	"forum/app/post/model"
	"forum/app/post/rpc/internal/svc"
	"forum/app/post/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePostLogic) CreatePost(in *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	post := &model.Posts{
		CommunityId: in.CommunityId,
		AuthorId:    in.AuthorId,
		Title:       in.Title,
		Content:     in.Content,
		CreateTime:  time.Now(),
		UpdatedTime: time.Now(),
	}
	insertRes, err := l.svcCtx.PostModel.Insert(l.ctx, post)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("create post failed, err: %v", err)
		return nil, err
	}
	id, _ := insertRes.LastInsertId()
	// 插入记录的同时， 在Redis中插入分数记录， 投票记录以及社区记录
	pipe, err := l.svcCtx.RedisClient.TxPipeline()
	if err != nil {
		logx.WithContext(l.ctx).Errorf("create post failed, err: %v", err)
		return nil, err
	}

	// TODO:  使用消息队列进行优化， 否则可能出现这样的状况， 就是MySQL创建成功了， 但是Redis没有创建成功。使用消息队列能够确保最终一致性。
	// 上面是存放的数据和当下存放在redis中的数据是不同的
	pipe.ZAdd(l.ctx, globalkey.GetRedisKey(globalkey.PostScoreKey), redis.Z{
		Member: id,
		Score:  float64(time.Now().Unix()),
	})
	pipe.ZAdd(l.ctx, globalkey.GetRedisKey(globalkey.PostUpCountKey), redis.Z{
		Member: id,
		Score:  0,
	})
	pipe.ZAdd(l.ctx, globalkey.GetRedisKey(globalkey.PostDownCountKey), redis.Z{
		Member: id,
		Score:  0,
	})
	pipe.SAdd(l.ctx, fmt.Sprintf(globalkey.GetRedisKey(globalkey.PostCommunityKey), in.CommunityId), id)

	_, err = pipe.Exec(l.ctx)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("create post failed, err: %v", err)
		return nil, err
	}
	return &pb.CreatePostResponse{
		PostId: id,
	}, nil
}
