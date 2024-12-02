package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"forum/app/post/rpc/internal/svc"
	"forum/app/post/rpc/pb"
	"forum/app/user/rpc/userservice"
	"forum/common/globalkey"
	"forum/common/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const (
	PostDetailExpireTime = 5 * 60
)

func NewGetPostDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostDetailLogic {
	return &GetPostDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostDetailLogic) GetPostDetail(in *pb.GetPostDetailRequest) (*pb.GetPostDetailResponse, error) {
	key := fmt.Sprintf(globalkey.GetRedisKey(globalkey.PostDetailKey), in.PostId)
	// 先从redis中获取帖子的详细信息
	val, err := l.svcCtx.RedisClient.Get(key)
	if err != nil && err != redis.Nil {
		logx.WithContext(l.ctx).Errorf("get redis failed, err: %v", err)
	}
	if val != "" {
		logx.WithContext(l.ctx).Infof("get post detail from redis, postId: %v", in.PostId)
		resp := &pb.GetPostDetailResponse{}
		if err := json.Unmarshal([]byte(val), resp); err != nil {
			logx.WithContext(l.ctx).Errorf("unmarshal redis value failed, err: %v", err)
		} else {
			return resp, nil
		}
	}

	// 如果没有数据， 那么就使用singleflight 去获取
	v, err, _ := l.svcCtx.SingleFlightGroup.Do(key, func() (interface{}, error) {
		// 输入一个帖子id， 获取帖子的详细信息
		postResp, err := l.svcCtx.PostModel.FindOne(l.ctx, in.PostId)
		if err != nil {
			logx.WithContext(l.ctx).Errorf("get post detail failed, err: %v", err)
			return nil, errors.Wrapf(xerr.NewErrMsg("get post detail failed"), "get post detail failed, postId: %v", in.PostId)
		}

		// 获取用户的详细信息
		userResp, err := l.svcCtx.UserRpc.GetUserDetail(l.ctx, &userservice.UserInfoRequest{
			UserId: postResp.AuthorId,
		})
		if err != nil {
			logx.WithContext(l.ctx).Errorf("get user detail failed, err: %v", err)
			return nil, errors.Wrapf(xerr.NewErrMsg("get user detail failed"), "get user detail failed, userId: %v", postResp.AuthorId)
		}

		postRet := &pb.Post{}
		copier.Copy(postRet, postResp)
		userRet := &pb.Author{}
		copier.Copy(userRet, userResp.User)

		// 获取点赞数据
		idStr := strconv.FormatInt(postResp.PostId, 10)
		keyUp := globalkey.GetRedisKey(globalkey.PostUpCountKey)
		keyDown := globalkey.GetRedisKey(globalkey.PostDownCountKey)

		upCount, err := l.svcCtx.RedisClient.Zscore(keyUp, idStr)
		if err != nil && err != redis.Nil {
			logx.WithContext(l.ctx).Errorf("get up count failed: %v", err)
			upCount = 0
		}

		downCount, err := l.svcCtx.RedisClient.Zscore(keyDown, idStr)
		if err != nil && err != redis.Nil {
			logx.WithContext(l.ctx).Errorf("get down count failed: %v", err)
			downCount = 0
		}
		ret := &pb.GetPostDetailResponse{
			Post:      postRet,
			UserInfo:  userRet,
			UpCount:   int64(upCount),
			DownCount: int64(downCount),
		}

		return ret, nil
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("get post detail failed, err: %v", err)
		return nil, errors.Wrapf(xerr.NewErrMsg("get post detail failed"), "get post detail failed, postId: %v", in.PostId)
	}
	postRet := v.(*pb.GetPostDetailResponse)
	// 将数据存入redis
	valByte, err := json.Marshal(postRet)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("marshal post detail failed, err: %v", err)
		return nil, errors.Wrapf(xerr.NewErrMsg("marshal post detail failed"), "marshal post detail failed, postId: %v", in.PostId)
	}
	err = l.svcCtx.RedisClient.Setex(key, string(valByte), PostDetailExpireTime)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("set redis failed, err: %v", err)
		return nil, errors.Wrapf(xerr.NewErrMsg("set redis failed"), "set redis failed, postId: %v", in.PostId)
	}

	return postRet, nil
}
